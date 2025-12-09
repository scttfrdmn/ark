# CLAUDE.md

## Response Style
- Concise by default. No explanations unless asked.
- No file creation unless explicitly requested.
- Fix bugs silently unless cause is non-obvious.

## Go Standards
- Go 1.23+ with modules
- `gofmt`, `goimports` on all code
- Pass `go vet`, `staticcheck`, `golangci-lint` before done
- Godoc comments on all exported identifiers
- No `panic` except unrecoverable init failures

## Code Style
- Idiomatic short names: `r` for reader, `ctx` for context, `err` for error
- Wrap errors with `fmt.Errorf("operation: %w", err)`
- Return early on errors; avoid deep nesting
- Prefer standard library over dependencies
- Group imports: stdlib, external, internal

## CLI Patterns
- Use `cobra` for CLI structure
- Flags over args when >1 input
- Exit codes: 0=success, 1=error, 2=usage error
- Stderr for errors/logs, stdout for output
- Support `--json` output where applicable

## AWS SDK
- Use `aws-sdk-go-v2`
- Load config with `config.LoadDefaultConfig(ctx)`
- Always pass context for cancellation
- Wrap SDK errors with operation context
- Use pagination helpers for list operations

## Testing
- Minimum 60% coverage, target 80%+
- Table-driven tests as default
- Use `t.Helper()` in test helpers
- Mock AWS with interfaces, not SDK mocks
- Test error paths, not just happy path
- Use `testdata/` for fixtures
- Golden files for complex output verification

## Security
- Never log credentials or tokens
- Use `golang.org/x/crypto` for cryptographic operations
- Validate all external inputs
- Sanitize before logging user-provided data

## Project Setup
- Use `setup-go-project.sh` for new projects
- Standard layout: cmd/, internal/, pkg/

## Git & GitHub
- Use `gh` CLI for all GitHub operations
- Conventional commits: `feat:`, `fix:`, `refactor:`, `test:`, `docs:`
- Branch naming: `feat/`, `fix/`, `refactor/` prefixes
- PR per feature/fix; link to issue

## Pre-commit Checks
- Run before every commit: `gofmt`, `go vet`, `staticcheck`
- Smoke tests: `go test -short ./...`
- Use pre-commit hook or `make check`

## Testing Workflow
- `make check` — fast: fmt, vet, lint, short tests
- `make test` — full: all unit tests with coverage
- `make integration` — slow: integration/e2e tests
- `make build` — build binary with version
- Run `make check` before every commit

## Project Tracking
- Track all work via GitHub Issues
- Use GitHub Projects for planning/status
- Use Milestones for releases/versions
- Create labels as needed if no good existing match
- Close issues via commit message: `Fixes #123`

## Do Not
- Create README, docs, or configs unless asked
- Add dependencies without justification
- Use `interface{}` or `any` without reason
- Ignore returned errors
- Use global state
