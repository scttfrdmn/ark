#!/usr/bin/env bash
set -euo pipefail

# Go CLI Project Setup Script
# Companion to CLAUDE.md
# Version: 2.0 (December 2025)

echo "=== Go CLI Project Setup ==="
echo ""

# Check prerequisites
if ! command -v go &> /dev/null; then
  echo "Error: go not found. Install Go 1.23+ first."
  exit 1
fi

if ! command -v gh &> /dev/null; then
  echo "Error: gh CLI not found. Install from https://cli.github.com"
  exit 1
fi

# Gather project info
read -rp "Project name (e.g., myctl): " PROJECT_NAME
read -rp "GitHub owner (user/org): " GITHUB_OWNER
read -rp "Short description: " DESCRIPTION
read -rp "Initial version [v0.1.0]: " VERSION
VERSION=${VERSION:-v0.1.0}
read -rp "Create minimal README? [Y/n]: " CREATE_README
CREATE_README=${CREATE_README:-y}
read -rp "Create GitHub Actions CI? [Y/n]: " CREATE_CI
CREATE_CI=${CREATE_CI:-y}
read -rp "Create GitHub Project board? [y/N]: " CREATE_PROJECT
read -rp "Private repo? [y/N]: " PRIVATE_REPO

MODULE="github.com/${GITHUB_OWNER}/${PROJECT_NAME}"

echo ""
echo "Creating: ${MODULE}"
echo ""

# Create directory structure
mkdir -p "${PROJECT_NAME}"/{cmd/"${PROJECT_NAME}",internal,pkg,testdata}
cd "${PROJECT_NAME}"

# Initialize Go module
go mod init "${MODULE}"

# Create main.go
cat > "cmd/${PROJECT_NAME}/main.go" <<EOF
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:     "${PROJECT_NAME}",
	Short:   "${DESCRIPTION}",
	Version: version,
}
EOF

# Create Makefile
cat > Makefile <<'EOF'
.PHONY: check test integration build clean install-tools

# Variables can be overridden: make build VERSION=v1.0.0
VERSION ?= dev
PROJECT_NAME ?= $(shell basename $(CURDIR))

# Fast pre-commit checks
check:
	gofmt -l -w .
	go vet ./...
	staticcheck ./...
	go test -short -race ./...

# Full unit tests with coverage
test:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out

# Integration/e2e tests
integration:
	go test -race -tags=integration ./...

# Build binary with version
build:
	go build -ldflags="-s -w -X main.version=$(VERSION)" -o bin/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✓ Tools installed"

clean:
	rm -rf bin/ coverage.out
EOF

# Create pre-commit hook
mkdir -p .githooks
cat > .githooks/pre-commit <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
make check
EOF
chmod +x .githooks/pre-commit

# Create .gitignore
cat > .gitignore <<'EOF'
bin/
coverage.out
*.exe
*.test
.DS_Store
*.log
EOF

# Create .golangci.yml
cat > .golangci.yml <<'EOF'
linters:
  enable:
    - gofmt
    - govet
    - errcheck
    - staticcheck
    - unused
    - gosimple
    - ineffassign
    - gosec
    - bodyclose
run:
  timeout: 5m
EOF

# Create README if requested
if [[ "${CREATE_README,,}" != "n" ]]; then
  cat > README.md <<EOF
# ${PROJECT_NAME}

${DESCRIPTION}

## Installation

\`\`\`bash
go install ${MODULE}/cmd/${PROJECT_NAME}@latest
\`\`\`

## Usage

\`\`\`bash
${PROJECT_NAME} --help
\`\`\`

## Development

\`\`\`bash
# Install development tools
make install-tools

# Run pre-commit checks (fast)
make check

# Run full test suite
make test

# Build binary
make build
\`\`\`

## License

Apache 2.0
EOF
fi

# Create GitHub Actions CI if requested
if [[ "${CREATE_CI,,}" != "n" ]]; then
  mkdir -p .github/workflows
  cat > .github/workflows/ci.yml <<'EOFCI'
name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: '1.23'
          cache: true

      - name: Install tools
        run: make install-tools

      - name: Run checks
        run: make check

      - name: Run tests
        run: make test

      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out
EOFCI
fi

# Install dependencies
echo "Installing dependencies..."
go get github.com/spf13/cobra
go mod tidy

# Install development tools
echo "Installing development tools..."
if ! command -v staticcheck &> /dev/null; then
  go install honnef.co/go/tools/cmd/staticcheck@latest
fi
if ! command -v golangci-lint &> /dev/null; then
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

# Initialize git
git init -q
git config core.hooksPath .githooks

# Create GitHub repo
echo "Creating GitHub repository..."
VISIBILITY="--public"
[[ "${PRIVATE_REPO,,}" == "y" ]] && VISIBILITY="--private"

gh repo create "${GITHUB_OWNER}/${PROJECT_NAME}" \
  --description "${DESCRIPTION}" \
  ${VISIBILITY} \
  --source . \
  --remote origin

# Create initial labels
echo "Creating labels..."
gh label create "priority:high" --color "D93F0B" --force 2>/dev/null || true
gh label create "priority:low" --color "0E8A16" --force 2>/dev/null || true
gh label create "type:bug" --color "D73A4A" --force 2>/dev/null || true
gh label create "type:feature" --color "0075CA" --force 2>/dev/null || true
gh label create "type:refactor" --color "CFD3D7" --force 2>/dev/null || true

# Create milestone
echo "Creating milestone ${VERSION}..."
gh api repos/"${GITHUB_OWNER}/${PROJECT_NAME}"/milestones \
  -f title="${VERSION}" \
  -f description="Initial release" \
  -f state="open" 2>/dev/null || true

# Create GitHub Project if requested
if [[ "${CREATE_PROJECT,,}" == "y" ]]; then
  echo "Creating GitHub Project..."
  if ! gh project create --owner "${GITHUB_OWNER}" --title "${PROJECT_NAME}" 2>/dev/null; then
    echo "⚠️  Could not create project automatically."
    echo "   Create manually: https://github.com/${GITHUB_OWNER}/${PROJECT_NAME}/projects/new"
  fi
fi

# Initial commit
echo "Creating initial commit..."
git add -A
git commit -m "feat: initial project scaffold

- Set up Go module structure
- Add Cobra CLI framework
- Configure linters and pre-commit hooks
- Add Makefile with check/test/build targets
$([ "${CREATE_CI,,}" != "n" ] && echo "- Configure GitHub Actions CI")
$([ "${CREATE_README,,}" != "n" ] && echo "- Add README with usage instructions")"

git push -u origin main

# Create and push version tag
echo "Creating version tag ${VERSION}..."
git tag "${VERSION}"
git push origin "${VERSION}"

echo ""
echo "=== Setup Complete ==="
echo ""
echo "cd ${PROJECT_NAME}"
echo "make check  # verify setup"
echo ""
echo "Repository: https://github.com/${GITHUB_OWNER}/${PROJECT_NAME}"
echo ""
echo "Next steps:"
echo "  1. Create issues for initial features"
echo "  2. Update README with project-specific details"
echo "  3. Add your first command: cobra-cli add <command>"
echo ""
