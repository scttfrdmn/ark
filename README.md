# Ark: AWS Research Kit
### Integrated Cloud Security Training & Tooling for Academic Research

**Open Source** | Apache 2.0 License | Institutional-Agnostic Framework

**⚠️ PROTOTYPE/WORK IN PROGRESS** - This document outlines a proposed solution. Feedback welcome!

---

## Overview

**Ark** is an open-source framework that provides research institutions with integrated AWS training and security tooling. It combines progressive, just-in-time training with production-grade security guardrails, ensuring researchers can use AWS safely and compliantly from day one.

**Key Innovation**: Training-as-tool - security education is embedded directly into the workflow, not delivered separately.

---

## The Problem

Research institutions face critical challenges when enabling AWS access for their researchers:

- **Security incidents**: Exposed credentials, misconfigured S3 buckets, unencrypted sensitive data
- **Compliance gaps**: HIPAA, CUI, FERPA, and institutional data classification violations
- **Cost overruns**: Forgotten instances, orphaned resources, lack of budget controls
- **Training disconnect**: Generic AWS training doesn't translate to research workflows
- **Support burden**: Repetitive questions, preventable mistakes, reactive firefighting

**Current approach**: Separate training courses + generic AWS tools = knowledge doesn't transfer to practice.

---

## The Solution: Training-as-Tool

**Ark** provides both a command-line interface (CLI) and web interface that simultaneously train researchers and provide production security tooling. Institutions can customize branding, policies, and training content to match their specific requirements.

### Dual Interface Approach

**CLI** - For technical users and automation:
- Command-line tool for scriptable, repeatable workflows
- Rich terminal UI for interactive training
- Ideal for power users, reproducible research, CI/CD integration

**Web Interface** - For visual learners and administrators:
- Browser-based application using AWS Cloudscape design system
- Streamlined, curated AWS console experience
- Interactive training with visual examples and simulations
- Administrative dashboards for institutional oversight

Both interfaces share the same backend, ensuring consistent training gates, security policies, and audit trails.

### Architecture Overview

Ark consists of three components:

```
┌─────────────┐          ┌──────────────┐         ┌─────────────────┐
│   CLI Tool  │          │  Web App     │         │ Institutional   │
│             │          │  (Browser)   │         │ Backend         │
│  (Go CLI)   │          │ (Vue/Cloud-  │         │ (Go API)        │
│             │          │  scape)      │         │                 │
└──────┬──────┘          └──────┬───────┘         └────────┬────────┘
       │                        │                          │
       └─────────┬──────────────┘                          │
                 │                                         │
          ┌──────▼──────┐                                 │
          │ Local Agent │◄────────────────────────────────┘
          │ (Go proxy)  │  Training state, policy checks
          │ localhost   │  Audit logging
          └──────┬──────┘
                 │
                 │ AWS credentials (never leave machine)
                 │
          ┌──────▼──────┐
          │  AWS APIs   │
          └─────────────┘
```

**Local Agent**: Runs on user's machine, manages AWS credentials securely, enforces policies, caches training content for offline use

**Institutional Backend**: Centralized API for training progress, policy definitions, audit logging, compliance reporting, user provisioning

**CLI & Web**: User interfaces that communicate through the local agent, ensuring consistent experience and security

---

## How It Works: First-Time User Experience

**Scenario**: A new researcher receives notification that their institutional AWS account is ready.

**Example**: We'll use "{INSTITUTION}" as a placeholder - institutions customize this during deployment.

**Prerequisites**:
- macOS, Linux, or Windows computer
- Internet connection for initial setup
- AWS CLI v2.15+ (Ark will help install if missing)

### Step 1: Installation (2 minutes)

```bash
$ curl -sSL https://ark.{INSTITUTION}.edu/install.sh | bash

╔══════════════════════════════════════════════════════════════╗
║  🚀 Installing Ark - AWS Research Kit                        ║
╚══════════════════════════════════════════════════════════════╝

→ Detecting system... macOS (arm64)
→ Downloading ark v1.2.0... ✓
→ Installing ark CLI to /usr/local/bin/ark... ✓
→ Installing ark-agent service... ✓
→ Starting ark-agent on localhost:8737... ✓

✅ Ark installed successfully!

Both CLI and web interface are now available:
  • CLI: ark --help
  • Web: https://ark.{INSTITUTION}.edu

Next: ark init --institution {INSTITUTION}
```

The installation sets up both the CLI tool and the local agent service. The web interface is accessed through your browser and requires no additional installation.

### Step 2: Configuration (5 minutes)

**Via CLI:**
```bash
$ ark init --institution {INSTITUTION}

╔══════════════════════════════════════════════════════════════╗
║  🎓 {INSTITUTION} AWS Research Tool Setup                    ║
╚══════════════════════════════════════════════════════════════╝

→ Loading {INSTITUTION} configuration...
  ✓ Configuration loaded from https://ark.{INSTITUTION}.edu/config

→ Required training modules:
  1. AWS Basics for Researchers (35 min)
  2. IAM & Identity Management (25 min)
  3. Institutional Data Classification (25 min)
  4. S3 Storage Security (35 min)

  📚 Total: ~120 minutes (can pause and resume)

→ Downloading training content... ✓
  (Available offline after download)

✅ Setup complete!

Next steps:
  • CLI: ark login
  • Web: Visit https://ark.{INSTITUTION}.edu
```

**Or via Web:**
*[Placeholder: Screenshot would show web interface with setup wizard, {INSTITUTION} branding, and progress indicators using Cloudscape design components]*

The web interface provides the same configuration experience with a visual wizard, progress indicators, and institutional branding.

### Step 3: Authentication (3 minutes)

**Single Sign-On for Both Interfaces:**

```bash
$ ark login

╔══════════════════════════════════════════════════════════════╗
║  🔐 Authentication                                           ║
╚══════════════════════════════════════════════════════════════╝

Opening browser for authentication...
→ https://ark.{INSTITUTION}.edu/auth/login

[Browser opens for institutional SSO + MFA]

✅ Authentication successful!

Account: 123456789012 ({INSTITUTION} Research)
User: researcher@{INSTITUTION}.edu

Credentials valid for both:
  • CLI (ark commands)
  • Web (https://ark.{INSTITUTION}.edu)

Would you like to start training now? [Y/n]: y
```

**Key Feature**: Login once, works everywhere. After CLI login, visiting the web interface shows you're already authenticated. The local agent manages secure token storage and refresh.

### Step 4: Progressive Training

Ark uses **progressive training** - you only complete modules when you need them for specific operations.

**Example: Researcher wants to create an S3 bucket**

**Via CLI:**
```bash
$ ark bucket create --name my-research-data --classification P2

╔═══════════════════════════════════════════════════════════╗
║  🎓 Training Required                                     ║
╠═══════════════════════════════════════════════════════════╣
║  Before creating buckets, complete:                       ║
║    • Module 3: Data Classification (25 min)              ║
║    • Module 4: S3 Storage Security (35 min)              ║
║                                                           ║
║  You'll learn this command while completing training!     ║
╚═══════════════════════════════════════════════════════════╝

Start Module 3 now? [Y/n]: y
```

**Via Web:**
*[Placeholder: Screenshot would show a modal dialog with Cloudscape Alert component explaining training requirement, progress bars showing 2/4 modules complete, and a "Start Training" button. The create bucket form would be visible but disabled in the background.]*

The web interface presents the same training gate with visual progress indicators, but in a graphical modal dialog using Cloudscape design components.

**After training**, the command executes with built-in security controls:
- ✓ Encryption at rest (AES-256)
- ✓ Encryption in transit (TLS 1.3)
- ✓ Versioning and access logging
- ✓ Block all public access
- ✓ Cost monitoring enabled

### Training Content Delivery

**CLI Training Experience:**
- Rich terminal UI with interactive elements
- Text-based quizzes with immediate feedback
- Code examples and command demonstrations
- Progress saved automatically
- Works offline (cached content)

**Web Training Experience:**
*[Placeholder: Screenshot series would show:
1. Training module layout with Cloudscape Container components, navigation sidebar showing progress
2. Interactive quiz with visual question formats (multiple choice with radio buttons, drag-and-drop classification exercises)
3. Visual diagrams explaining S3 bucket structure, data flows
4. Video embedding for complex topics
5. Certificate download page with PDF preview]*

The web interface offers enhanced interactivity:
- Embedded videos and animations
- Interactive simulations (sandbox environments)
- Visual quiz formats (drag-and-drop, image selection)
- Progress dashboards with charts
- Social features (see peer progress, anonymized)

**Both interfaces track progress to the same institutional backend**, so you can start training in the CLI and continue in the web browser, or vice versa.

---

## Key Features

### 🎓 **Progressive Training**
- Just-in-time learning when attempting new operations
- Interactive tutorials embedded in actual commands/workflows
- Quiz checkpoints ensure comprehension
- Completion tracking and certificate generation (cryptographically signed PDFs)
- Cross-interface sync (start CLI, finish web, or vice versa)

### 🔒 **Built-in Compliance**
- Institutional data classification validation (customizable taxonomy)
- HIPAA, CUI, FERPA requirement enforcement
- Pre-approved policy templates
- Automatic security best practices
- Cannot disable critical security controls

### 🛡️ **Bulletproof Operations**
- Automatic retry with exponential backoff
- Handles AWS eventual consistency
- Transaction rollback on failures
- Idempotent operations (safe to re-run)

### 💰 **Cost Protection**
- Mandatory billing alerts
- Auto-shutdown for compute instances
- Orphaned resource detection
- Budget enforcement hooks
- Real-time cost visualization (web interface)

### 📊 **Institutional Oversight**
- Centralized completion tracking
- Security posture dashboards (web-based)
- Audit trail integration with CloudTrail
- Customizable training content per department
- Role-based access (researcher, PI, admin, CISO views)

### 🌐 **Dual Interface Choice**
- **CLI**: Scriptable, automatable, efficient for power users
- **Web**: Visual, interactive, accessible, great for learning
- **Unified**: Same training, same policies, same audit trail
- **User preference**: Choose the interface that fits your workflow

---

## Open Source & Institutional Flexibility

### Licensing
- **Apache 2.0 License**: Free to use, modify, and distribute
- **Open Core**: All components are open source
- **Community-driven**: Contributions welcome via GitHub

### Institutional Customization

Institutions can customize Ark to their specific needs:

**Branding:**
- Logo, colors, institution name throughout UI
- Custom domain (ark.{your-institution}.edu)
- Configurable terminology (adapt to local naming conventions)

**Policy Framework:**
- Define your own data classification system (or use templates)
- Map to regulatory frameworks (HIPAA, FISMA, etc.)
- Set budget limits and approval workflows
- Customize security baselines

**Training Content:**
- Modify existing modules to match institutional policies
- Add institution-specific case studies
- Create domain-specific modules (genomics, HPC, etc.)
- Multiple language support

**Integration:**
- SSO with any SAML/OIDC identity provider
- LDAP/Active Directory sync
- SCIM provisioning
- SIEM integration (Splunk, ELK, AWS Security Hub)
- Ticketing systems (ServiceNow, Jira)

**Example Configuration:**
```yaml
# ark-config.yaml
institution:
  name: "University of Research"
  short_name: "UResearch"
  domain: "ark.uresearch.edu"

branding:
  logo: "https://cdn.uresearch.edu/logo.png"
  primary_color: "#003366"

data_classification:
  framework: "custom"  # or "nist", "iso27001", "uc-p1-p4"
  levels:
    - id: "public"
      name: "Public"
      encryption_required: false
    - id: "internal"
      name: "Internal"
      encryption_required: true
    - id: "confidential"
      name: "Confidential"
      encryption_required: true
      mfa_required: true
    - id: "restricted"
      name: "Restricted"
      encryption_required: true
      mfa_required: true
      approval_required: true

training:
  modules:
    - id: "aws-basics"
      required: true
      duration_minutes: 35
      content_url: "https://content.uresearch.edu/modules/aws-basics.md"
    # ... additional modules

identity:
  sso_provider: "saml"
  idp_url: "https://sso.uresearch.edu"
  attributes:
    uid: "eduPersonPrincipalName"
    email: "mail"
    department: "ou"
```

---

## Implementation Approach

### Phase 1: Core Foundation (Months 1-2)
- Local agent (Go) with AWS credential management
- Institutional backend API (Go) with REST + GraphQL
- CLI tool (Go, Cobra) with basic commands
- Web application (Vue 3, Cloudscape, TypeScript)
- Authentication (SSO integration)

### Phase 2: Training Integration (Months 2-3)
- 4 core training modules (customizable templates)
- Interactive checkpoints and quizzes (CLI + Web)
- Training gate enforcement in agent
- Certificate generation with cryptographic proof
- Progress sync across interfaces

### Phase 3: Institutional Deployment (Month 4)
- Institution-specific configuration system
- Integration with existing identity management
- Training content customization tools
- Admin dashboards and reporting (web-based)
- SIEM integration

**For detailed implementation timeline, see [ROADMAP.md](ROADMAP.md)**

---

## Technology Stack

**Backend:**
- Language: Go (single binary, cross-platform, fast)
- AWS SDK: Official AWS SDK v2 for Go
- Database: DynamoDB (progress, audit), S3 (content, logs)
- API: REST + GraphQL (for web subscriptions)

**Local Agent:**
- Language: Go
- Server: HTTP proxy on localhost:8737
- Storage: SQLite (local cache), secure keychain (credentials)

**CLI:**
- Language: Go
- Framework: Cobra (commands), bubbletea (interactive UI)
- Distribution: GitHub releases, package managers (brew, apt, chocolatey)

**Web Application:**
- Framework: Vue 3 with TypeScript
- Design System: AWS Cloudscape (consistent with AWS console)
- Build: Vite
- State: Pinia
- Testing: Playwright (E2E), Vitest (unit)

**Authentication:**
- SSO: SAML 2.0, OAuth 2.0/OIDC
- Tokens: JWT (access + refresh)
- MFA: Institutional (DUO, etc.)

---

## Benefits

### For Researchers
✓ **One tool to learn** - Training and production tooling unified
✓ **Interface choice** - Use CLI or web based on preference
✓ **Faster onboarding** - 2 hours to full AWS competency
✓ **Confidence** - Can't make critical security mistakes
✓ **Self-service** - Standard operations don't require approval
✓ **Works offline** - Training content cached locally

### For IT Security
✓ **Enforced compliance** - Can't skip security controls
✓ **Reduced incidents** - Built-in guardrails prevent common mistakes (target: 80% reduction)
✓ **Audit trails** - Complete logging of training and operations
✓ **Scalable** - Minimal support burden as researchers self-serve
✓ **Visibility** - Web dashboards for institutional oversight

### For Institutions
✓ **Risk reduction** - Systematic security control enforcement
✓ **Cost control** - Automated budget monitoring (reduce surprises by 90%)
✓ **Compliance** - Demonstrable training and audit trails for regulators
✓ **Customizable** - Adapt to institutional policies and branding
✓ **Open source** - No vendor lock-in, community-driven improvements
✓ **Multi-interface** - Supports diverse user preferences and accessibility needs

---

## Success Metrics

**Measurement Period**: Evaluated at 6 and 12 months post-deployment
**Baseline**: 6 months prior to Ark deployment
**Reporting**: Ongoing dashboard with quarterly reviews

- **Training completion rate**: Target 95% within 30 days of AWS access
- **Security incidents**: Reduce by 80% compared to baseline
- **Cost incidents**: Reduce surprise bills >$1000 by 90%
- **Support tickets**: Reduce AWS-related tickets by 60%
- **Time to productivity**: <2 hours from account creation to first resource deployed
- **User satisfaction**: >4.0/5.0 for both CLI and web interfaces
- **Accessibility**: WCAG 2.1 AA compliance for web interface

---

## Getting Started

### For Institutions
1. **Review documentation**: Understand architecture and requirements
2. **Customize configuration**: Define data classifications, branding, policies
3. **Deploy backend**: Set up institutional backend API and database
4. **Integrate identity**: Connect to SSO provider (SAML/OIDC)
5. **Customize training**: Adapt modules to institutional policies
6. **Pilot program**: Start with 50-100 users in 2-3 research labs
7. **Iterate**: Gather feedback and refine
8. **Broad rollout**: Deploy institution-wide

### For Researchers
1. **Install Ark**: Follow institutional installation guide
2. **Authenticate**: Login with institutional credentials
3. **Complete training**: ~2 hours of progressive, interactive training
4. **Start using AWS**: Create resources with built-in security
5. **Choose your interface**: Use CLI for scripts, web for exploration

### For Developers
1. **Clone repository**: `git clone https://github.com/aws-research-kit/ark`
2. **Read CONTRIBUTING.md**: Guidelines for contributions
3. **Set up dev environment**: Follow setup instructions
4. **Pick an issue**: Check GitHub issues for good first contributions
5. **Submit PR**: Follow contribution guidelines

---

## Community & Support

- **GitHub**: [github.com/aws-research-kit/ark](https://github.com/aws-research-kit/ark)
- **Documentation**: [docs.ark-aws.org](https://docs.ark-aws.org)
- **Community Forum**: [community.ark-aws.org](https://community.ark-aws.org)
- **Slack**: [ark-aws.slack.com](https://ark-aws.slack.com)

**Institutional Support**: Each institution provides first-line support for their users. Community provides shared knowledge base and collaboration.

---

## License

Apache License 2.0 - See [LICENSE](LICENSE) file for details.

Copyright © 2025 Ark Contributors

---

## Appendix A: Detailed User Walkthrough

### Scenario: Dr. Sarah Chen, Postdoc in Computational Biology

**Background**: Sarah needs to analyze 500GB of genomic data. She's comfortable with Python and the command line but has never used AWS. Her PI just got her an AWS account through her institution.

---

### Day 1, 9:00 AM - Installation

Sarah receives an email from IT:

> Your institutional AWS account is ready!
> Install Ark to get started: https://ark.{INSTITUTION}.edu/install

**Sarah chooses CLI (she's comfortable with terminal):**

```bash
$ curl -sSL https://ark.{INSTITUTION}.edu/install.sh | bash

╔══════════════════════════════════════════════════════════════╗
║  🚀 Installing Ark - AWS Research Kit                        ║
╚══════════════════════════════════════════════════════════════╝

→ Detecting system... macOS (arm64)
→ Checking for AWS CLI... Not found
  Installing AWS CLI v2.15.2... ✓
→ Downloading ark v1.2.0... ✓
→ Installing to /usr/local/bin/ark... ✓
→ Installing ark-agent service... ✓
→ Starting ark-agent... ✓
→ Verifying installation... ✓

✅ Ark installed successfully!

Next steps:
  1. Run: ark init --institution {INSTITUTION}
  2. Complete setup: ark login
  3. Start training: ark learn start

Need help? Visit https://ark.{INSTITUTION}.edu/docs
```

---

### 9:02 AM - Initial Configuration

```bash
$ ark init --institution {INSTITUTION}

╔══════════════════════════════════════════════════════════════╗
║  🎓 {INSTITUTION} AWS Research Tool Setup                    ║
╚══════════════════════════════════════════════════════════════╝

→ Loading {INSTITUTION} configuration...
  📥 Downloading from: https://ark.{INSTITUTION}.edu/config
  ✓ Configuration loaded

Institution: {INSTITUTION}
Support Email: aws-support@{INSTITUTION}.edu
Documentation: https://it.{INSTITUTION}.edu/aws

→ Required training modules:
  1. AWS Basics for Researchers (35 min)
  2. IAM & Identity Management (25 min)
  3. Data Classification (25 min)
  4. S3 Storage Security (35 min)

  📚 Total estimated time: 120 minutes
  💡 You can pause and resume anytime!

→ Downloading training content...
  Module 1/4: AWS Basics... ✓
  Module 2/4: IAM & Identity... ✓
  Module 3/4: Data Classification... ✓
  Module 4/4: S3 Security... ✓

✅ Setup complete!

Your Progress: ░░░░░░░░░░ 0/4 modules (0%)

Next: ark login
```

---

### 9:05 AM - Authentication

```bash
$ ark login

╔══════════════════════════════════════════════════════════════╗
║  🔐 Authentication                                           ║
╚══════════════════════════════════════════════════════════════╝

Opening browser for authentication...
  🌐 https://ark.{INSTITUTION}.edu/auth/login

[Browser opens, Sarah logs in with institutional credentials and MFA]
[After successful login, browser shows: "You may now close this window"]

→ Waiting for authentication... ✓

✅ Authentication successful!

→ Verifying credentials...
  Account: 123456789012 ({INSTITUTION} Research)
  User: AIDAI...XYZ (sarah.chen@{INSTITUTION}.edu)
  ✓ Credentials verified

→ Checking your permissions...
  ✓ S3 access: Read/Write
  ✓ EC2 access: Launch instances
  ✓ IAM access: Limited (read-only)
  ✓ Cost Explorer: View own usage

💡 Your permissions follow the "{INSTITUTION} Researcher" policy.

✅ All systems ready!

💡 Web interface also ready: https://ark.{INSTITUTION}.edu
   (You're already logged in)

Would you like to start training now? [Y/n]: y
```

---

### 9:10 AM - Module 1: AWS Basics

```bash
╔══════════════════════════════════════════════════════════════╗
║  📚 Module 1: AWS Basics for Researchers                     ║
║  Duration: ~35 minutes                                       ║
╚══════════════════════════════════════════════════════════════╝

Welcome, Sarah! 👋

This module covers:
  • What is AWS and why researchers use it
  • Key services: S3 (storage), EC2 (computing), IAM (security)
  • {INSTITUTION}'s AWS setup and support resources
  • How costs work and how to avoid surprises
  • **CRITICAL: Common security threats and how to prevent them**

Press ENTER to begin...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Section 1.1: What is AWS?

AWS (Amazon Web Services) is like renting lab equipment, but for computing.
Instead of buying servers, you rent what you need, when you need it.

Why researchers love AWS:
  ✓ Scale up for big analyses, scale down when done
  ✓ Pay only for what you use
  ✓ Access to powerful GPUs without buying hardware
  ✓ Collaborate by sharing data securely
  ✓ 99.99% uptime - more reliable than local servers

Real example from {INSTITUTION}:
  Dr. Martinez (Neuroscience) analyzed 10TB of fMRI data using
  100 EC2 instances for 8 hours. Cost: $240.

  Buying equivalent hardware: ~$50,000 + maintenance.

Press ENTER to continue...
```

[Training continues through all sections, including security fundamentals, interactive quizzes, and hands-on exercises]

```bash
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

✅ Module 1 Complete!

Time: 35 minutes
Score: 100% (5/5 correct)

🔒 Security Concepts Learned:
  ✓ Shared responsibility model
  ✓ Real-world incident examples
  ✓ The 5 Golden Rules
  ✓ Credential protection
  ✓ Incident response basics

Progress: ▓▓▓░░░░░░░ 1/4 modules (25%)

Commands Unlocked:
  ✓ ark cost alert    - Manage billing alerts
  ✓ ark cost report   - View spending breakdown
  ✓ ark audit basics  - Check account security

Continue to Module 2: IAM & Identity Management? [Y/n]: n

No problem! Resume anytime with: ark learn continue

Your progress is saved automatically.
💡 You can also continue training in the web interface:
   https://ark.{INSTITUTION}.edu/training
```

---

### 9:45 AM - Sarah Takes a Break, Checks Web Interface

*[Placeholder: Screenshot would show web dashboard with:
- Progress ring chart showing 25% complete (1/4 modules)
- Module 1 marked complete with green checkmark and certificate icon
- Modules 2-4 shown as "Not Started" with locked icons
- "Continue Training" button for Module 2
- Sidebar showing her AWS resources (currently empty)
- Cost widget showing $0.00 this month]*

Sarah opens her browser to see what the web interface offers. She's logged in automatically (same session as CLI). The dashboard shows her training progress and will eventually show her AWS resources.

---

### 9:50 AM - Trying to Use S3 (Training Gate)

Sarah decides to try creating a bucket for her genomic data:

```bash
$ ark bucket create --name sarah-genomics-data --classification internal

╔══════════════════════════════════════════════════════════════╗
║  ⚠️  Training Required                                       ║
╠══════════════════════════════════════════════════════════════╣
║  Before creating S3 buckets, you must complete:             ║
║                                                              ║
║  Module 3: Data Classification ................... ✗         ║
║    (~25 min - learn sensitivity levels)                      ║
║                                                              ║
║  Module 4: S3 Storage Security ................... ✗         ║
║    (~35 min - encryption, access control)                    ║
║                                                              ║
║  Why? Creating buckets incorrectly is a top security risk.  ║
║  These modules ensure you protect your research data.        ║
╚══════════════════════════════════════════════════════════════╝

Start Module 3 now? [Y/n]: y
```

**Same experience in web interface:**
*[Placeholder: Screenshot would show:
- Create Bucket form with fields filled in (bucket name, classification dropdown)
- Modal overlay with Cloudscape Alert component
- Title: "Training Required"
- Alert type: "warning"
- Two module cards showing Module 3 and 4 requirements
- Progress indicators showing "Not Started"
- Large "Start Training" button
- "Learn More" link
- Background form is dimmed/disabled]*

---

### 10:00 AM - Module 3: Data Classification

Sarah completes Module 3, learning about institutional data sensitivity levels, then Module 4 about S3 security. [Content similar to original walkthrough]

---

### 11:10 AM - Creating Her First Bucket

```bash
$ ark bucket create --name sarah-genomics-data --classification internal

╔══════════════════════════════════════════════════════════════╗
║  🗄️  S3 Bucket Creation                                      ║
╚══════════════════════════════════════════════════════════════╝

→ Validating bucket name...
  ✓ Name is valid and available

→ Confirming classification: Internal

  For Internal data, we'll configure:
    • ✓ Encryption at rest (AES-256)
    • ✓ TLS required for uploads
    • ✓ Versioning (protects from accidents)
    • ✓ Access logging
    • ✓ Block public access
    • ✗ MFA delete (not required for Internal)

  This configuration meets {INSTITUTION} policy for Internal data.

Proceed? [Y/n]: y

→ Creating S3 bucket...
  ✓ Bucket created: sarah-genomics-data

→ Enabling encryption...
  ✓ Encryption enabled

→ Enabling versioning...
  ✓ Versioning enabled

→ Configuring access logging...
  ✓ Access logging enabled

→ Blocking public access...
  ✓ Public access blocked

→ Adding cost allocation tags...
  ✓ Tagged: PI=chen-lab, Department=biology, Project=genomics

╔══════════════════════════════════════════════════════════════╗
║  ✅ Bucket Created Successfully!                             ║
╠══════════════════════════════════════════════════════════════╣
║  Name: sarah-genomics-data                                   ║
║  Classification: Internal                                    ║
║  Region: us-west-2                                           ║
║  Encryption: ✓ | Versioning: ✓ | Logging: ✓                 ║
║  Public Access: ✗ (Blocked)                                  ║
║  Estimated Cost: ~$12/month for 500GB                        ║
╚══════════════════════════════════════════════════════════════╝

Next steps:
  1️⃣  Upload your data:
     ark bucket upload --name sarah-genomics-data --file sequences.tar.gz

  2️⃣  Share with collaborators:
     ark bucket share --name sarah-genomics-data --with colleague@{INSTITUTION}.edu

  3️⃣  Monitor costs:
     ark cost report --bucket sarah-genomics-data

⏱️  Total time: 47 seconds

💡 View this bucket in web interface:
   https://ark.{INSTITUTION}.edu/buckets/sarah-genomics-data
```

**Checking the web interface:**
*[Placeholder: Screenshot would show:
- Buckets list page with Cloudscape Table component
- One row showing sarah-genomics-data bucket
- Columns: Name, Classification (badge), Region, Size (0 GB), Cost ($0.00), Created (timestamp)
- Status column with green "Active" badge
- Actions menu with: Upload Files, Share, View Details, Delete
- Top of page has metrics cards showing: Total Buckets (1), Total Storage (0 GB), Monthly Cost ($0.00)
- "Create Bucket" button in top-right]*

---

### Week 2 - Sarah is Self-Sufficient

Over the past two weeks, Sarah has:
- ✓ Completed all 4 training modules (~2 hours total)
- ✓ Uploaded 487 GB of genomic data to S3 (via web interface's drag-and-drop)
- ✓ Launched EC2 instances for analysis (via CLI for scripting)
- ✓ Stayed within her $150/month budget
- ✓ Shared data with two external collaborators securely
- ✓ Uses both CLI (for scripts) and web (for monitoring)

**Using web interface for monitoring:**
*[Placeholder: Screenshot would show main dashboard with:
- Header with {INSTITUTION} logo and "Welcome back, Sarah Chen"
- Cloudscape SpaceBetween layout with multiple Container components
- KPI cards showing: 2 Buckets, 3 Instances (1 running, 2 stopped), $127.43 this month
- Line chart showing cost trend over past 30 days
- Table showing active resources with status badges
- Alert banner: "Budget Alert: You've used 85% of your $150 monthly budget"
- Quick actions: Create Bucket, Launch Instance, View Training Certificate
- Security score widget: 94/100 with "Excellent" badge]*

**Running security audit via CLI:**
```bash
$ ark audit scan

Running security audit on all your resources...

✅ Overall Security Score: 94/100

Findings:

✓ S3 Buckets (2)
  • sarah-genomics-data: Perfect ✓
  • sarah-results: Perfect ✓

✓ EC2 Instances (1 running)
  • i-0123456789abcdef: Auto-shutdown enabled ✓

⚠️  IAM
  • MFA not enabled on your user account
    Fix: ark iam mfa enable

💰 Cost Optimization
  • You could save $45/month by switching idle instances to t3.medium
    Current spend: $127/month (on track)
    Budget: $150/month

✓ Compliance
  • All resources properly tagged ✓
  • Access logging enabled ✓
  • Encryption at rest enabled ✓

📊 Full report also available at:
   https://ark.{INSTITUTION}.edu/audit/latest
```

---

## Appendix B: Module Template Structure

Training modules in Ark are defined using YAML configuration and Markdown content. This allows institutions to customize modules to their specific policies and requirements.

### Module Configuration File: `module.yaml`

```yaml
# Module metadata and configuration
module:
  id: "03-data-classification"
  version: "2.1.0"
  name: "Data Classification"
  short_description: "Understanding data sensitivity and protection requirements"
  estimated_minutes: 25

  # What this module teaches
  learning_objectives:
    - "Classify data using institutional framework"
    - "Identify sensitive data types"
    - "Understand legal frameworks (FERPA, HIPAA, CUI)"
    - "Recognize re-identification risks"
    - "Apply appropriate security controls per classification"

  # Required before this module
  prerequisites:
    - "01-aws-basics"

  # What becomes available after completion
  unlocks:
    commands:
      - "ark bucket create"
      - "ark classify"
    next_modules:
      - "04-s3-security"

  # Passing requirements
  completion_criteria:
    quiz_passing_score: 85  # Higher for compliance content
    hands_on_required: true
    scenario_passing_score: 90

# Content sources (can be remote URLs or local files)
content:
  tutorial: "https://training.{INSTITUTION}.edu/modules/03/tutorial.md"
  quiz: "https://training.{INSTITUTION}.edu/modules/03/quiz.yaml"
  scenarios: "https://training.{INSTITUTION}.edu/modules/03/scenarios.yaml"

# Institution-specific customization
customization:
  institution: "{INSTITUTION}"

  # Institution-specific sections to inject
  custom_sections:
    - position: "after_intro"
      content_url: "https://training.{INSTITUTION}.edu/modules/03/policy.md"
      title: "{INSTITUTION}-Specific Requirements"

  # Override default contact information
  contacts:
    questions: "data-classification@{INSTITUTION}.edu"
    incidents: "security@{INSTITUTION}.edu"
    compliance: "compliance@{INSTITUTION}.edu"
```

Institutions can host their own training content, customizing examples, case studies, and policies while maintaining the core Ark framework.

---

## Appendix C: Architecture Deep Dive

### Component Interaction Flow

**Example: User creates S3 bucket via CLI**

```
1. User executes command:
   $ ark bucket create --name my-data --classification internal

2. CLI validates local syntax, sends to agent:
   POST http://localhost:8737/api/bucket/create
   Headers: Authorization: Bearer {local_jwt}
   Body: { name: "my-data", classification: "internal" }

3. Agent checks training status:
   GET https://ark.{INSTITUTION}.edu/api/user/training-status
   Response: { "s3_security": "completed", "data_classification": "completed" }

4. Agent validates policy locally:
   - Internal data requires encryption ✓
   - User has completed training ✓
   - Bucket name is valid ✓

5. Agent makes AWS API call:
   s3.CreateBucket(bucket_name)
   s3.PutBucketEncryption(...)
   s3.PutBucketVersioning(...)
   s3.PutPublicAccessBlock(...)

6. Agent logs to institutional backend (async):
   POST https://ark.{INSTITUTION}.edu/api/audit/events
   Body: { user, action: "s3:CreateBucket", resource: "my-data", ... }

7. Agent returns success to CLI:
   Response: { status: "success", bucket: { name, arn, region, ... } }

8. CLI displays result to user
```

**Same flow for web interface:**

```
1. User clicks "Create Bucket" in web UI

2. Web app sends to agent:
   POST http://localhost:8737/api/bucket/create
   Headers: Authorization: Bearer {local_jwt}
   Body: { name: "my-data", classification: "internal" }

3-7. [Identical to CLI flow]

8. Web app receives response, updates UI:
   - Shows success notification (Cloudscape Flash)
   - Adds bucket to table
   - Updates cost estimate
```

### Security Boundaries

```
┌─────────────────────────────────────────────────────────────┐
│ User's Machine (TRUSTED)                                    │
│                                                             │
│  ┌──────────┐                                              │
│  │ AWS Creds│ ← Never leave this machine                   │
│  │ (SSO)    │                                              │
│  └────┬─────┘                                              │
│       │                                                     │
│  ┌────▼─────────────┐       ┌─────────────────┐           │
│  │  Local Agent     │       │  CLI / Web      │           │
│  │  (Go service)    │◄──────┤  (User Interface)│          │
│  │  localhost:8737  │       │                 │           │
│  └────┬─────────────┘       └─────────────────┘           │
│       │                                                     │
│       │ ① AWS API calls (with creds)                       │
│       │ ② Policy checks (training, compliance)             │
│       │                                                     │
└───────┼─────────────────────────────────────────────────────┘
        │
        │ ② HTTPS (no creds, only user actions)
        │
┌───────▼─────────────────────────────────────────────────────┐
│ Institutional Backend (UNTRUSTED for AWS creds)             │
│                                                             │
│  • Training progress                                        │
│  • Policy definitions                                       │
│  • Audit logs                                              │
│  • User provisioning                                        │
│  • Certificates                                            │
│  • Reporting                                               │
│                                                             │
│  Never sees AWS credentials ✓                              │
└─────────────────────────────────────────────────────────────┘
```

**Key Security Principle**: AWS credentials never leave the user's machine. The institutional backend never has access to AWS credentials, only to training state and audit logs.

### Offline Capability

```
User working offline (airplane, field site):

1. Training content cached locally:
   ~/.ark/cache/modules/*.md

2. Can complete training offline:
   - Read content from cache
   - Take quizzes (validated locally)
   - Progress saved to local SQLite

3. When back online:
   - Agent syncs progress to backend
   - Backend validates quiz answers
   - Certificates generated
   - Audit logs uploaded
```

### Cross-Interface Sync

```
Scenario: User starts training on CLI, continues on web

1. CLI: User starts Module 3
   - Progress: { module: "03", status: "in_progress", page: 5 }
   - Saved to: Backend + Local cache

2. Web: User visits training page
   - Fetches progress from backend via agent
   - Shows Module 3 at page 5 (resume point)
   - User continues in web browser

3. Web: User completes Module 3
   - Progress: { module: "03", status: "completed", score: 90 }
   - Saved to: Backend + Local cache

4. CLI: User runs command requiring Module 3
   - Agent checks backend
   - Module 3 is completed ✓
   - Command proceeds
```

---

**For implementation details and timeline, see [ROADMAP.md](ROADMAP.md)**
**For technical gaps and recommendations, see [ark-gaps-and-suggestions.md](ark-gaps-and-suggestions.md)**
