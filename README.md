# Ark: AWS Research Kit for UCLA
### Integrated Cloud Security Training & Tooling for Academic Research

**⚠️ PROTOTYPE/WORK IN PROGRESS** - This document outlines a proposed solution. Feedback welcome!

---

## The Problem

UCLA researchers need AWS access for computationally intensive research, but face critical challenges:

- **Security incidents**: Exposed credentials, misconfigured S3 buckets, unencrypted sensitive data
- **Compliance gaps**: HIPAA, CUI, FERPA, and UC data classification (P1-P4) violations
- **Cost overruns**: Forgotten instances, orphaned resources, lack of budget controls
- **Training disconnect**: Generic AWS training doesn't translate to research workflows
- **Support burden**: Repetitive questions, preventable mistakes, reactive firefighting

**Current approach**: Separate training courses + generic AWS tools = knowledge doesn't transfer to practice.

---

## The Solution: Training-as-Tool

**Ark** is a unified command-line tool that simultaneously trains researchers and provides production security tooling.

### How It Works: First-Time User Experience

**Scenario**: A new researcher receives notification that their UCLA AWS account is ready.

**Prerequisites**: 
- macOS, Linux, or Windows computer
- Internet connection for initial setup
- AWS CLI v2.15+ (Ark will help install if missing)

#### Step 1: Installation (2 minutes)

```bash
$ curl -sSL https://ark.ucla.edu/install.sh | bash

╔══════════════════════════════════════════════════════════════╗
║  🚀 Installing Ark - AWS Research Kit                        ║
╚══════════════════════════════════════════════════════════════╝

→ Detecting system... macOS (arm64)
→ Downloading ark v1.2.0... ✓
→ Installing to /usr/local/bin/ark... ✓

✅ Ark installed successfully!

Next: ark init --institution ucla
```

#### Step 2: Configuration (5 minutes)

```bash
$ ark init --institution ucla

╔══════════════════════════════════════════════════════════════╗
║  🎓 UCLA AWS Research Tool Setup                             ║
╚══════════════════════════════════════════════════════════════╝

→ Loading UCLA configuration...
  ✓ Configuration loaded

→ Required training modules:
  1. AWS Basics for Researchers (35 min)
  2. IAM & Identity Management (25 min)
  3. UC Data Classification (P1-P4) (25 min)
  4. S3 Storage Security (35 min)
  
  📚 Total: ~120 minutes (can pause and resume)

→ Downloading training content... ✓

Next: ark setup wizard
```

#### Step 3: AWS Authentication (3 minutes)

```bash
$ ark setup wizard

╔══════════════════════════════════════════════════════════════╗
║  🔐 AWS Authentication Setup                                 ║
╚══════════════════════════════════════════════════════════════╝

Let's connect you to UCLA's AWS environment.

→ Checking AWS CLI... ✓ AWS CLI v2.15.2 detected

📖 You'll log in with your UCLA credentials + DUO (two-factor authentication).
   Uses SSO (Single Sign-On) - no API keys to manage. More secure!

Ready? [Y/n]: y

→ Executing: aws login

Opening browser for authentication...
[Browser opens for UCLA SSO login with DUO]

✅ Authentication successful!

Account: 123456789012 (UCLA Research)
User: sarah.chen@ucla.edu

Would you like to start training now? [Y/n]: y
```

#### Step 4: Training-as-You-Go

Ark uses **progressive training** - you only complete modules when you need them for specific operations.

Basic AWS commands trigger Module 1, but creating storage buckets requires understanding data classification (Module 3) and storage security (Module 4) first. Module 2 (IAM) becomes required when managing users and permissions.

After basic setup, when the researcher tries to create a storage bucket:

```bash
# Example: Trying to create a bucket for internal research data
# (Note: P2 = "Internal" classification - you'll learn this in Module 3)
$ ark bucket create --name my-research-data --classification P2

╔═══════════════════════════════════════════════════════════╗
║  🎓 Training Required                                     ║
╠═══════════════════════════════════════════════════════════╣
║  Before creating buckets, complete:                       ║
║    • Module 3: UC Data Classification (25 min)           ║
║    • Module 4: S3 Storage Security (35 min)              ║
║                                                           ║
║  You'll learn this command while completing training!     ║
╚═══════════════════════════════════════════════════════════╝

Start Module 3 now? [Y/n]: y
```

**After training**, the command executes with built-in security controls:
- ✓ Encryption at rest (AES-256)
- ✓ Encryption in transit (TLS 1.3)
- ✓ Versioning and access logging
- ✓ Block all public access
- ✓ Cost monitoring enabled

**Key insight**: Training isn't a separate hurdle - it's embedded in the workflow. Researchers learn by doing, using the actual production tool.

**📖 Complete Journey**: See Appendix A for a detailed day-in-the-life walkthrough following Dr. Sarah Chen from installation through productive AWS usage (includes all training modules totaling ~2 hours, real commands, troubleshooting, and week 2 self-sufficiency). See Appendix B for module template structure and customization options.

---

## Key Features

### 🎓 **Progressive Training**
- Just-in-time learning when attempting new operations
- Interactive tutorials embedded in actual commands
- Quiz checkpoints ensure comprehension
- Completion tracking and certification generation (PDF certificates with cryptographic proof, recognized by institutional compliance offices for audit purposes)

### 🔒 **Built-in Compliance**
- UC P1-P4 data classification validation
- HIPAA, CUI, FERPA requirement enforcement
- Pre-approved policy templates
- Automatic security best practices

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

### 📊 **Institutional Oversight**
- Centralized completion tracking
- Security posture dashboards
- Audit trail integration with CloudTrail
- Customizable training content per department

---

## Implementation Approach

### Phase 1: Core Tool (Months 1-2)
- IAM user/group management with MFA enforcement
- S3 bucket creation with classification-based security
- EC2 instance lifecycle management
- Cost monitoring and alerting

### Phase 2: Training Integration (Months 2-3)
- 4 required modules (AWS basics, IAM, data classification, S3 security)
- Interactive checkpoints and quizzes
- Completion verification system
- Certificate generation

### Phase 3: Institutional Deployment (Month 4)
- UCLA-specific configuration (SSO, policies, support contacts)
- Integration with existing identity management
- Training content customization for departments
- Admin dashboards and reporting

---

## Benefits

### For Researchers
✓ **One tool to learn** - Training and production tooling unified  
✓ **Faster onboarding** - 2 hours to full AWS competency (vs weeks with traditional training)  
✓ **Confidence** - Can't make critical security mistakes (built-in guardrails)  
✓ **Self-service** - Standard operations (buckets, instances, databases) don't require approval
   (Note: P4 data and specialized resources still require institutional review)

### For IT Security
✓ **Enforced compliance** - Can't skip security controls  
✓ **Reduced incidents** - Built-in guardrails prevent common mistakes (target: 80% reduction)  
✓ **Audit trails** - Complete logging of training and operations  
✓ **Scalable** - Minimal support burden as researchers self-serve

### For UCLA
✓ **Risk reduction** - Systematic security control enforcement  
✓ **Cost control** - Automated budget monitoring and alerting (reduce unexpected costs by 90%, save ~$200k/year in support)  
✓ **Compliance** - Demonstrable training and audit trails for regulators  
✓ **Competitive advantage** - Enables cutting-edge research safely

---

## Technology

- **Language**: Go (single binary, cross-platform, fast)
- **AWS SDK**: Official AWS SDK v2 for Go
- **Distribution**: GitHub releases, institutional package repos
- **Configuration**: YAML-based, remotely updatable training content
- **Authentication**: AWS SSO with new `aws login` command support

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

---

## Next Steps

1. **Pilot program** with 2-3 research labs (Month 1)
   - Owner: IT Security Team with Solutions Architecture support
   - Target: 50-100 users
2. **Refinement** based on researcher feedback (Month 2)
   - Owner: Product team + CISO Office
3. **Broader rollout** to departments with AWS needs (Month 3-4)
   - Owner: IT Leadership
4. **Mandatory requirement** for new AWS account requests (Month 5+)
   - Owner: Institutional Policy

---

## Appendix A: First-Time User Walkthrough

### Scenario: Dr. Sarah Chen, Postdoc in Computational Biology

**Background**: Sarah needs to analyze 500GB of genomic data. She's comfortable with Python and the command line but has never used AWS. Her PI just got her an AWS account through UCLA.

---

### Day 1, 9:00 AM - Installation

Sarah receives an email from IT:

> Your UCLA AWS account is ready!  
> Install Ark to get started: https://ark.ucla.edu/install

```bash
$ curl -sSL https://ark.ucla.edu/install.sh | bash

╔══════════════════════════════════════════════════════════════╗
║  🚀 Installing Ark - AWS Research Kit                        ║
╚══════════════════════════════════════════════════════════════╝

→ Detecting system... macOS (arm64)
→ Checking for AWS CLI... Not found
  Installing AWS CLI v2.15.2... ✓
→ Downloading ark v1.2.0... ✓
→ Installing to /usr/local/bin/ark... ✓
→ Verifying installation... ✓

✅ Ark installed successfully!

Next steps:
  1. Run: ark init --institution ucla
  2. Complete setup: ark setup wizard
  3. Start training: ark learn start

Need help? Visit https://ark.ucla.edu/docs
```

---

### 9:02 AM - Initial Configuration

```bash
$ ark init --institution ucla

╔══════════════════════════════════════════════════════════════╗
║  🎓 UCLA AWS Research Tool Setup                             ║
╚══════════════════════════════════════════════════════════════╝

→ Loading UCLA configuration...
  📥 Downloading from: https://ucla-aws-training.s3.amazonaws.com/config/ucla.yaml
  ✓ Configuration loaded

Institution: UCLA
Support Email: your institutional AWS support
Documentation: https://it.ucla.edu/aws

→ Required training modules:
  1. AWS Basics for Researchers (35 min)
  2. IAM & Identity Management (25 min)
  3. UC Data Classification (P1-P4) (25 min)
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

Next: ark setup wizard
```

---

### 9:05 AM - AWS Authentication Setup

```bash
$ ark setup wizard

╔══════════════════════════════════════════════════════════════╗
║  🔐 AWS Authentication Setup                                 ║
╚══════════════════════════════════════════════════════════════╝

Let's connect you to UCLA's AWS environment.

→ Checking AWS CLI installation...
  ✓ AWS CLI v2.15.2 detected
  ✓ Supports new 'aws login' command

📖 About AWS Single Sign-On (SSO)
   You'll log in with your UCLA credentials + DUO (two-factor authentication).
   No API keys to manage - more secure and simpler!
   
   The new 'aws login' command (AWS CLI v2.15+) simplifies
   authentication compared to the older 'aws sso login' method.

Ready to authenticate? [Y/n]: y

→ Executing: aws login

Opening your browser to authenticate...
  🌐 https://ucla.awsapps.com/start

[Browser opens, Sarah logs in with UCLA credentials and DUO]
[After successful login, browser shows: "You may now close this window"]

→ Waiting for authentication... ✓

✅ Authentication successful!

→ Verifying credentials...
  Account: 123456789012 (UCLA Research)
  User: AIDAI...XYZ (sarah.chen@ucla.edu)
  ✓ Credentials verified

→ Checking your permissions...
  ✓ S3 access: Read/Write
  ✓ EC2 access: Launch instances
  ✓ IAM access: Limited (read-only)
  ✓ Cost Explorer: View own usage

💡 Your permissions follow the "UCLA Researcher" policy.
   This gives you access to core services while maintaining security.

✅ All systems ready!

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
  • UCLA's AWS setup and support resources
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

Real UCLA example:
  Dr. Martinez (Neuroscience) analyzed 10TB of fMRI data using 
  100 EC2 instances for 8 hours. Cost: $240. 
  
  Buying equivalent hardware: ~$50,000 + maintenance.

Press ENTER to continue...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Section 1.2: Security First - Why This Matters

🚨 REAL INCIDENTS FROM ACADEMIC INSTITUTIONS (2023-2024):

❌ Incident 1: Public S3 Bucket
   University: Major Research Institution (anonymized)
   What happened: Researcher made bucket public to "share with collaborator"
   Result: 2TB of patient genomic data exposed for 6 months
   Impact: $4.2M HIPAA fine, IRB suspension, lawsuits
   
   Prevention: Ark BLOCKS public access by default for sensitive data

❌ Incident 2: Exposed AWS Keys in GitHub
   University: West Coast R1 Institution
   What happened: Student committed AWS keys to public GitHub repo
   Result: Cryptominers used account, $62,000 bill in 3 days
   Impact: Lab funding exhausted, student's PhD delayed
   
   Prevention: Ark uses SSO - no long-term keys to expose

❌ Incident 3: Forgotten EC2 Instance
   University: Midwest Research Lab
   What happened: Postdoc left institution, instance kept running
   Result: $18,000 over 14 months, GPUs sitting idle
   Impact: PI had to return grant funds to cover costs
   
   Prevention: Ark requires auto-shutdown configuration

⚠️  These aren't rare - they happen weekly across academia.

🛡️  YOUR RESPONSIBILITY:
   As a researcher with AWS access, you are:
   • A steward of sensitive research data
   • A guardian of lab/grant funding
   • A representative of UCLA's security posture
   
   This training ensures you don't become a cautionary tale.

Press ENTER to continue...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Section 1.3: The Shared Responsibility Model

┌─────────────────────────────────────────────────┐
│           AWS Responsibility                    │
│  (Security OF the cloud)                        │
│                                                 │
│  • Physical data centers                        │
│  • Hardware infrastructure                      │
│  • Network infrastructure                       │
│  • Virtualization layer                         │
└─────────────────────────────────────────────────┘
                       ↓
┌─────────────────────────────────────────────────┐
│          YOUR Responsibility                    │
│  (Security IN the cloud)                        │
│                                                 │
│  • Data encryption ← YOU must enable            │
│  • Access controls ← YOU must configure         │
│  • Credential management ← YOU must protect     │
│  • Network configuration ← YOU must secure      │
│  • Cost management ← YOU must monitor           │
└─────────────────────────────────────────────────┘

Think of it like a safe deposit box:
  • Bank secures the building (AWS's job)
  • You must lock your box and guard your key (YOUR job)

Ark helps you fulfill YOUR responsibilities correctly.

Press ENTER to continue...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Section 1.4: Core AWS Services

Think of AWS services as tools in a toolkit:

🗄️  S3 (Simple Storage Service)
   Like Dropbox, but for research data.
   • Store files from bytes to terabytes
   • Automatic redundancy (your data is safe)
   • Access from anywhere
   
   Example: Store your genomic sequence files
   
   🚨 Security consideration:
      Default S3 buckets are PRIVATE, but one wrong setting
      makes them PUBLIC. Ark prevents this mistake.

💻 EC2 (Elastic Compute Cloud)
   Rent virtual computers by the hour.
   • From small (2 CPUs) to huge (hundreds of CPUs)
   • GPU instances for machine learning
   • Run any software you need
   
   Example: Process 500GB of data in parallel
   
   🚨 Security consideration:
      Forgotten instances = wasted money. Always set auto-shutdown.

🔐 IAM (Identity & Access Management)
   Control who can access what.
   • Create users for lab members
   • Set permissions carefully
   • Enable multi-factor authentication
   
   Example: Give your student read-only access to data
   
   🚨 Security consideration:
      Over-privileged users = biggest risk. Follow "least privilege."

📊 CloudWatch
   Monitor costs and usage.
   • Set billing alarms
   • Track resource usage
   • Get alerts before overspending
   
   🚨 Security consideration:
      You MUST set billing alarms - treat it like a fume hood alarm.

Press ENTER to continue...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Section 1.5: The 5 Golden Rules of AWS Security

Remember these ALWAYS:

1️⃣  NEVER share credentials
   • Not with lab mates
   • Not via email
   • Not in Slack/Teams
   • Not in code repositories
   
   If someone needs access → Create them an account

2️⃣  ALWAYS enable MFA (multi-factor authentication)
   • Prevents 99.9% of account compromises
   • Takes 2 minutes to set up
   • Required by UCLA policy
   
   Your password alone is NOT enough

3️⃣  ENCRYPT everything sensitive
   • P3/P4 data MUST be encrypted
   • Encryption at rest + in transit
   • Don't assume it's automatic
   
   Ark handles this for you when you classify correctly

4️⃣  MONITOR your costs daily
   • Set billing alarms FIRST
   • Check costs at end of each day
   • Investigate unusual spikes immediately
   
   Financial responsibility = security responsibility

5️⃣  AUDIT regularly
   • What resources are running?
   • Who has access to what?
   • Are security settings still correct?
   
   Use: ark audit scan (weekly recommended)

These aren't suggestions - they're requirements.

Press ENTER to continue...

[Training continues through sections on costs, billing, support...]

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎯 Checkpoint Quiz

Let's check your understanding!

Q1: What's the main advantage of AWS for researchers?
  a) It's free
  b) Pay for what you use, scale as needed
  c) Faster than local computers
  d) Automatic data analysis

Your answer: b

✅ Correct! The elasticity and pay-as-you-go model means you can 
   access massive compute resources without capital investment.

Q2: Which service would you use to store 200GB of sequencing data?
  a) EC2
  b) IAM
  c) S3
  d) CloudWatch

Your answer: c

✅ Exactly! S3 is designed for data storage at any scale.

Q3: What should you set up to avoid surprise AWS bills?
  a) CloudWatch billing alarms
  b) Nothing - AWS is always cheap
  c) Automatic shutdowns only
  d) IAM policies

Your answer: a

✅ Perfect! Always set billing alarms before using AWS.

Q4: 🔒 SECURITY QUESTION: You need to share AWS access with a 
    visiting collaborator. What should you do?

  a) Share your username and password
  b) Create them their own IAM user account
  c) Give them your laptop
  d) Email them your access keys

Your answer: b

✅ CORRECT! Never share credentials. Always create separate accounts.
   This ensures:
   • Accountability (know who did what)
   • Revocable access (can remove when they leave)
   • Audit trails (CloudTrail logs their actions)

Q5: 🔒 SECURITY QUESTION: You find your AWS access key accidentally
    committed to a public GitHub repo. What do you do?

  a) Delete the GitHub commit and hope no one saw it
  b) Immediately rotate the key and contact security
  c) Wait and see if anything bad happens
  d) Change your GitHub password

Your answer: b

✅ CRITICAL! Exposed keys = compromised account. Always:
   1. Rotate keys immediately (Ark will help)
   2. Contact your institutional AWS support
   3. Check CloudTrail for unauthorized usage
   4. Document the incident
   
   Keys can be scraped in minutes by bots.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎓 Hands-On Exercise: Set Up Your First Security Control

Now let's actually set up a billing alarm!

This is a real operation - we'll create an actual alarm on your account.

Why this matters:
  💰 Prevents cost overruns
  🚨 Early warning system
  📧 Alerts you before problems grow

→ Creating billing alarm...
  Name: sarah-chen-monthly-budget
  Threshold: $100/month
  Alert: sarah.chen@ucla.edu

Execute this operation? [Y/n]: y

→ Calling AWS CloudWatch API...
  ✓ Alarm created

→ Sending test notification...
  ✓ Check your email for confirmation

💡 You'll receive an email if spending exceeds $100/month.
   Adjust anytime with: ark cost alert update
   
   🔒 Security tip: Set alarms at multiple thresholds:
      • $50 - Advisory notice
      • $100 - Warning
      • $200 - Critical alert

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
```

---

### 9:45 AM - Sarah Takes a Coffee Break

She's learned the basics and completed the security foundations. Now she wants to actually upload her genomic data.

---

### 9:50 AM - Trying to Use S3 (Training Gate)

```bash
$ ark bucket create --name sarah-genomics-data --classification P2

╔══════════════════════════════════════════════════════════════╗
║  ⚠️  Training Required                                       ║
╠══════════════════════════════════════════════════════════════╣
║  Before creating S3 buckets, you must complete:             ║
║                                                              ║
║  Module 3: UC Data Classification ................ ✗         ║
║    (~15 min - learn P1-P4 levels)                            ║
║                                                              ║
║  Module 4: S3 Storage Security ................... ✗         ║
║    (~30 min - encryption, access control)                    ║
║                                                              ║
║  Why? Creating buckets incorrectly is a top security risk.  ║
║  These modules ensure you protect your research data.        ║
╚══════════════════════════════════════════════════════════════╝

Start Module 3 now? [Y/n]: y
```

---

### 10:00 AM - Module 3: UC Data Classification

```bash
╔══════════════════════════════════════════════════════════════╗
║  📚 Module 3: UC Data Classification (P1-P4)                 ║
╚══════════════════════════════════════════════════════════════╝

Understanding data sensitivity is CRITICAL for compliance and security.

⚠️  Getting this wrong has serious consequences:
   • Federal fines ($100k - $50M+ per incident)
   • Loss of grant funding
   • IRB suspension
   • Legal liability
   • Reputational damage to UCLA

This module ensures you classify and protect data correctly.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

UC Protection Levels - The Framework:

UCLA follows University of California data protection standards.
Every piece of data falls into one of four categories:

┌─────────────────────────────────────────────────────────────┐
│ P1 - PUBLIC INFORMATION                                     │
├─────────────────────────────────────────────────────────────┤
│ What: Information intended for public distribution          │
│                                                             │
│ Examples:                                                   │
│   ✓ Published research papers                              │
│   ✓ Public course catalogs                                 │
│   ✓ Campus directory information                           │
│   ✓ Marketing materials                                    │
│                                                             │
│ Requirements: None (already public)                         │
│ AWS: Standard S3, no special controls needed               │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ P2 - INTERNAL INFORMATION                                   │
├─────────────────────────────────────────────────────────────┤
│ What: Information for UCLA use only                         │
│                                                             │
│ Examples:                                                   │
│   ✓ Unpublished research data (no PII)                     │
│   ✓ Grant proposals (pre-submission)                       │
│   ✓ Internal reports and memos                             │
│   ✓ Non-sensitive lab data                                 │
│                                                             │
│ Requirements:                                               │
│   • Access limited to UCLA affiliates                      │
│   • Basic access controls                                  │
│   • Encryption recommended but not required                │
│                                                             │
│ AWS: Private S3 bucket, encryption enabled                  │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ P3 - PROTECTED INFORMATION  ← MOST COMMON FOR RESEARCH      │
├─────────────────────────────────────────────────────────────┤
│ What: Sensitive data requiring protection                   │
│                                                             │
│ Examples:                                                   │
│   ✓ Personal Identifiable Information (PII)                │
│     - Names, addresses, phone numbers                      │
│     - Email addresses, student IDs                         │
│     - Birth dates, driver's license numbers                │
│   ✓ Student records (FERPA protected)                      │
│   ✓ De-identified health data (not full PHI)              │
│   ✓ Research data with confidentiality agreements          │
│   ✓ Export-controlled research data                        │
│   ✓ Proprietary business information                       │
│                                                             │
│ Legal Frameworks:                                           │
│   • FERPA (Family Educational Rights and Privacy Act)      │
│   • PII protection laws (CCPA, GDPR if applicable)         │
│   • Contractual confidentiality obligations                │
│                                                             │
│ Requirements:                                               │
│   • ✓ Encryption at rest (REQUIRED)                        │
│   • ✓ Encryption in transit (REQUIRED)                     │
│   • ✓ Access logging for audits                            │
│   • ✓ Strong access controls                               │
│   • ✓ MFA for administrators                               │
│   • ✓ Incident response plan                               │
│   • ✓ Regular access reviews                               │
│                                                             │
│ AWS: Ark P3 configuration enforces ALL requirements         │
│                                                             │
│ 🚨 Common Mistake: "It's de-identified so it's fine"       │
│    Even de-identified data can often be re-identified!     │
│    When in doubt, treat as P3.                             │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ P4 - HIGHLY RESTRICTED INFORMATION                          │
├─────────────────────────────────────────────────────────────┤
│ What: Extremely sensitive data with strict regulations      │
│                                                             │
│ Examples:                                                   │
│   ✓ Protected Health Information (PHI) - HIPAA            │
│   ✓ Social Security Numbers                               │
│   ✓ Financial account numbers                             │
│   ✓ Controlled Unclassified Information (CUI)             │
│   ✓ ITAR/EAR controlled technical data                    │
│   ✓ Credit card numbers (PCI DSS)                         │
│                                                             │
│ Legal Frameworks:                                           │
│   • HIPAA (Health Insurance Portability Act)               │
│   • NIST 800-171 (CUI protection)                          │
│   • CMMC (DoD cybersecurity)                               │
│   • ITAR (International Traffic in Arms)                   │
│   • PCI DSS (Payment Card Industry)                        │
│                                                             │
│ Requirements:                                               │
│   • ✓ All P3 requirements PLUS:                            │
│   • ✓ Pre-approved AWS account configuration               │
│   • ✓ Business Associate Agreement (BAA) for HIPAA        │
│   • ✓ Enhanced monitoring and alerting                     │
│   • ✓ Dedicated security review                            │
│   • ✓ Compliance officer approval                          │
│   • ✓ Annual audits                                        │
│   • ✓ Incident notification within 24-72 hours            │
│                                                             │
│ AWS: Requires CISO office approval BEFORE use               │
│      Contact: your institutional HIPAA compliance office                     │
│                                                             │
│ ⚠️  DO NOT store P4 data without explicit approval!        │
└─────────────────────────────────────────────────────────────┘

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🔒 REAL WORLD: Classification Failures

Case Study 1: The "Anonymous" Survey
  ✗ Scenario: Researcher collected "anonymous" health surveys
  ✗ Reality: Included zip code + age + gender
  ✗ Problem: This combination can identify ~87% of US population
  ✗ Classification error: Treated as P2, actually P3 (maybe P4!)
  ✗ Consequence: Data breach notification to 1,200 participants
  
  Lesson: Combinations of "non-sensitive" data = sensitive data

Case Study 2: The Collaboration Mistake  
  ✗ Scenario: Shared student performance data with external partner
  ✗ Reality: Didn't get data sharing agreement
  ✗ Problem: FERPA violation (student data improperly disclosed)
  ✗ Consequence: $50,000 fine, IRB investigation
  
  Lesson: P3 data sharing requires agreements, even with collaborators

Case Study 3: The De-identification Assumption
  ✗ Scenario: Published "de-identified" genomic sequences
  ✗ Reality: Sequences + public genealogy DB = re-identification
  ✗ Problem: Participants identified, privacy violated
  ✗ Consequence: Study retracted, lawsuits filed
  
  Lesson: De-identification is harder than you think

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📋 Decision Tree: When In Doubt

Start here: Does your data contain ANY of the following?

  ┌─ Names, email addresses, phone numbers?
  │  └─ YES → At least P3
  │
  ┌─ Student records or grades?
  │  └─ YES → P3 (FERPA applies)
  │
  ┌─ Health information (even de-identified)?
  │  └─ YES → At least P3, possibly P4 if identifiable
  │
  ┌─ Financial data, SSNs, credit cards?
  │  └─ YES → P4 (stop, contact CISO office)
  │
  ┌─ Under confidentiality agreement?
  │  └─ YES → Read agreement, probably P3
  │
  ┌─ Export controlled (ITAR/EAR)?
  │  └─ YES → P4 (stop, contact export control office)
  │
  ┌─ Will be published/public eventually?
  │  └─ YES but not yet → P2 until published
  │
  └─ None of the above?
     └─ Probably P1 or P2, but verify with PI

🆘 Still unsure? That's OK!
   Contact: your institutional data classification office
   They'll help you classify correctly (better safe than sorry)

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎯 Interactive Exercise: Classify Sarah's Data

Your research scenario:
  • Genomic sequences from Drosophila (fruit flies)
  • No human subjects
  • No personally identifiable information
  • Funded by NSF grant
  • Will be published when analysis complete
  • No confidentiality agreements

What classification level? [P1/P2/P3/P4]: P2

✅ Correct! This is P2 (Internal) because:
   • Not yet published (so not P1)
   • No PII or regulated data (so not P3/P4)
   • Internal research data until publication
   • Non-human subject research

When you publish, you can reclassify to P1.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎯 Checkpoint Quiz (Higher Stakes)

Q1: You have a dataset with: age (binned in 5-year ranges), 
    zip code, and diagnosis. No names. What level?

  a) P1 - It's de-identified
  b) P2 - Internal use only
  c) P3 - Can be re-identified
  d) P4 - Contains health info

Your answer: c

✅ CORRECT! Even without names, this is P3 because:
   • Age + zip code + diagnosis = potentially identifiable
   • Health information requires protection even when de-identified
   • Could violate HIPAA if re-identified
   
   This is called "quasi-identifiers" - seemingly anonymous
   data that can be combined to identify individuals.

Q2: Your collaborator at Stanford needs access to your 
    P3 research data. What do you need?

  a) Just share an S3 link
  b) Data sharing agreement + BAA if needed
  c) Their email address
  d) Nothing special - they're at a university

Your answer: b

✅ PERFECT! For P3 data sharing, you need:
   1. Data Sharing Agreement (legal framework)
   2. Business Associate Agreement if health data (HIPAA)
   3. Document what data is shared and why
   4. Time-limited access (not permanent)
   5. UCLA IRB approval if human subjects
   
   Contact: your institutional data sharing office for templates

Q3: 🚨 COMPLIANCE SCENARIO: You discover you've been storing
    what you thought was P2 data, but it actually contains 
    email addresses (P3). What do you do?

  a) Delete the emails and move on
  b) Immediately report to CISO, re-classify, audit access
  c) Just fix it going forward
  d) Hope no one noticed

Your answer: b

✅ CRITICAL! When you discover a classification error:
   
   IMMEDIATE actions:
   1. Stop any current data sharing
   2. Email: your institutional security incident response team
   3. Document: What data? How long misclassified? Who had access?
   
   CISO will help you:
   • Re-classify correctly
   • Audit who accessed the data
   • Implement proper controls
   • Determine if breach notification needed
   
   🎯 Key principle: It's never wrong to report. It IS wrong to hide.

Q4: Can you mix P2 and P3 data in the same S3 bucket?

  a) Yes, it's fine
  b) Yes, but separate folders
  c) No, always use separate buckets
  d) Only with special permission

Your answer: c

✅ CORRECT! Best practice: Separate buckets per classification.
   
   Why?
   • Bucket-level encryption settings differ
   • Access controls are simpler
   • Audit logging is clearer
   • Reduces accidental exposure risk
   • Compliance audits are easier
   
   If you MUST mix (rare cases):
   • Get CISO approval
   • Use highest classification's controls (P3)
   • Document exception clearly
   • More frequent audits required

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📋 UCLA Data Classification Quick Reference

Save this! You'll need it.

┌──────────┬────────────┬─────────────┬──────────────┬──────────────┐
│ Question │     P1     │     P2      │      P3      │      P4      │
├──────────┼────────────┼─────────────┼──────────────┼──────────────┤
│ Contains │ None       │ None        │ Names, email │ SSN, PHI,    │
│ PII?     │            │             │ phone, DOB   │ financials   │
├──────────┼────────────┼─────────────┼──────────────┼──────────────┤
│ Encrypt  │ Optional   │ Recommended │ REQUIRED     │ REQUIRED +   │
│ at rest? │            │             │              │ key mgmt     │
├──────────┼────────────┼─────────────┼──────────────┼──────────────┤
│ Access   │ Public     │ UCLA only   │ Authorized   │ Minimal,     │
│ control? │            │             │ users only   │ documented   │
├──────────┼────────────┼─────────────┼──────────────┼──────────────┤
│ Logging  │ Optional   │ Recommended │ REQUIRED     │ REQUIRED +   │
│ required?│            │             │              │ monitoring   │
├──────────┼────────────┼─────────────┼──────────────┼──────────────┤
│ Breach   │ None       │ CISO notice │ CISO + OCR   │ CISO + OCR   │
│ notify?  │            │             │ if PII       │ within 72h   │
├──────────┼────────────┼─────────────┼──────────────┼──────────────┤
│ UCLA     │ None       │ None        │ Training     │ CISO pre-    │
│ approval?│            │             │ + compliance │ approval     │
└──────────┴────────────┴─────────────┴──────────────┴──────────────┘

Download full guide: ark classify --download-guide

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

✅ Module 3 Complete! (25 minutes)

🔒 Security Concepts Learned:
  ✓ UC P1-P4 classification framework
  ✓ Legal frameworks (FERPA, HIPAA, CUI)
  ✓ Re-identification risks
  ✓ Data sharing requirements
  ✓ Incident response for classification errors
  ✓ Compliance requirements per level

Progress: ▓▓▓▓▓▓░░░░ 2/4 modules (50%)

🎓 You now understand UCLA's data protection standards!

Commands Unlocked:
  ✓ ark classify --help    - Classification helper tool
  ✓ ark bucket create      - Create buckets (with classification)

Continue to Module 4: S3 Storage Security? [Y/n]: y
```

---

### 10:30 AM - Module 4: S3 Storage Security

```bash
╔══════════════════════════════════════════════════════════════╗
║  📚 Module 4: S3 Storage Security                            ║
║  Duration: ~35 minutes                                       ║
╚══════════════════════════════════════════════════════════════╝

This module covers:
  • Encryption at rest and in transit
  • Bucket policies and access controls
  • Versioning and lifecycle management
  • Access logging and monitoring
  • Common misconfigurations and how to avoid them

[Sarah completes Module 4, learning about:]
  • S3 encryption options (SSE-S3, SSE-KMS, SSE-C)
  • How bucket policies differ from IAM policies
  • Setting up lifecycle rules to reduce costs
  • Enabling access logging for audit trails
  • Preventing accidental public exposure
  • MFA Delete for critical data

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

✅ Module 4 Complete! (35 minutes)

Progress: ▓▓▓▓▓▓▓▓░░ 3/4 modules (75%)

🔒 Security Concepts Learned:
  ✓ S3 encryption methods and when to use each
  ✓ Bucket policy design for least privilege
  ✓ Lifecycle policies for cost optimization
  ✓ Access logging configuration
  ✓ Public access blocking (mandatory for P3/P4)
  ✓ MFA Delete protection

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎉 Great progress! You can now:
  ✓ Create S3 buckets with proper security
  ✓ Upload and manage research data
  ✓ Control access for collaborators

Commands Unlocked:
  ✓ ark bucket create
  ✓ ark bucket upload
  ✓ ark bucket share
  ✓ ark bucket audit

**Note on Module 2 (IAM)**: Sarah hasn't completed the IAM & Identity Management module yet. This is intentional - she'll complete it later when she needs to create users for lab members or manage permissions. Progressive training means you learn what you need, when you need it.

Ready to create your bucket now? [Y/n]: y
```

---

### 11:10 AM - Creating Her First Bucket (For Real)

```bash
$ ark bucket create --name sarah-genomics-data --classification P2

╔══════════════════════════════════════════════════════════════╗
║  🗄️  S3 Bucket Creation                                      ║
╚══════════════════════════════════════════════════════════════╝

→ Validating bucket name...
  ✓ Name is valid and available

→ Confirming classification: P2 (Internal Information)
  
  For P2 data, we'll configure:
    • ✓ Encryption at rest (AES-256)
    • ✓ TLS required for uploads
    • ✓ Versioning (protects from accidents)
    • ✓ Access logging
    • ✓ Block public access
    • ✗ MFA delete (not required for P2)

  This configuration meets UCLA policy for P2 data.

Proceed? [Y/n]: y

→ Creating S3 bucket...
  ⠋ Calling S3 CreateBucket API...
  ✓ Bucket created: sarah-genomics-data

→ Enabling encryption...
  ⠙ Configuring AES-256-GCM...
  ✓ Encryption enabled
  
  💡 Files are automatically encrypted when uploaded!

→ Enabling versioning...
  ⠹ Configuring version control...
  ✓ Versioning enabled
  
  💡 Previous versions kept for 90 days (configurable)

→ Configuring access logging...
  ⠸ Creating audit log bucket...
  ⠼ Enabling logging...
  ✓ Access logging enabled
  
  💡 All access is logged for security audits

→ Blocking public access...
  ⠧ Applying block public access...
  ✓ Public access blocked
  
  💡 This bucket is private - only you can access it

→ Setting lifecycle policy...
  ⠇ Transitioning old data to cheaper storage...
  ✓ Lifecycle policy applied
  
  💡 Data >90 days old moves to Glacier (90% cheaper!)

→ Adding cost allocation tags...
  ✓ Tagged: PI=chen-lab, Department=biology, Project=drosophila

╔══════════════════════════════════════════════════════════════╗
║  ✅ Bucket Created Successfully!                             ║
╠══════════════════════════════════════════════════════════════╣
║  Name: sarah-genomics-data                                   ║
║  Classification: P2 (Internal)                               ║
║  Region: us-west-2                                           ║
║  Encryption: ✓ | Versioning: ✓ | Logging: ✓                 ║
║  Public Access: ✗ (Blocked)                                  ║
║  Estimated Cost: ~$12/month for 500GB                        ║
╚══════════════════════════════════════════════════════════════╝

Next steps:

  1️⃣  Upload your data:
     ark bucket upload --name sarah-genomics-data --file sequences.tar.gz
     
  2️⃣  Share with collaborators:
     ark bucket share --name sarah-genomics-data --with colleague@ucla.edu
     
  3️⃣  Monitor costs:
     ark cost report --bucket sarah-genomics-data

📖 Quick tips:
   • Upload large files (>100MB) in parts for reliability
   • Use folders to organize: sequences/, results/, metadata/
   • Set up notifications: ark bucket notify --on-upload

⏱️  Total time: 47 seconds
```

---

### 12:00 PM - Uploading Data

```bash
$ ark bucket upload --name sarah-genomics-data --file drosophila-sequences.tar.gz

╔══════════════════════════════════════════════════════════════╗
║  📤 S3 Upload                                                ║
╚══════════════════════════════════════════════════════════════╝

File: drosophila-sequences.tar.gz
Size: 487 GB
Destination: s3://sarah-genomics-data/

→ Analyzing file...
  ✓ File is >100MB - will use multipart upload
  ✓ Splitting into 1,000 parts of 500MB each

→ Starting upload...
  ⚠️  This will take approximately 2.5 hours on your connection
  
  💡 You can close this terminal - upload will continue!
     Check status: ark bucket upload status

  Progress: ▓▓▓▓▓▓▓▓░░░░░░░░░░░░ 40% (195 GB / 487 GB)
  Speed: 21 MB/s
  ETA: 1h 52m
  
  Parts completed: 402/1000
  Failed/retrying: 3 (automatically retrying)

[Upload continues in background...]

✅ Upload complete! (2h 18m)

→ Verifying upload integrity...
  ✓ All parts uploaded successfully
  ✓ File checksum matches

Your data is now securely stored and encrypted in AWS!

Access: s3://sarah-genomics-data/drosophila-sequences.tar.gz
URL: https://sarah-genomics-data.s3.us-west-2.amazonaws.com/...

Estimated monthly cost: $11.80 (based on 487GB + 10% growth)
```

---

### Day 2 - Sarah Needs Compute

```bash
$ ark instance launch --type m5.xlarge --purpose "Sequence alignment"

╔══════════════════════════════════════════════════════════════╗
║  ⚠️  Training Required                                       ║
╠══════════════════════════════════════════════════════════════╣
║  Before launching EC2 instances, complete:                   ║
║                                                              ║
║  Module 5: Compute Security & Cost Control ...... ✗          ║
║    (~25 min)                                                 ║
║                                                              ║
║  This module covers:                                         ║
║    • Choosing the right instance type                        ║
║    • Auto-shutdown to prevent waste                          ║
║    • Security groups and network access                      ║
║    • Monitoring and cost optimization                        ║
╚══════════════════════════════════════════════════════════════╝

Start Module 5 now? [Y/n]: y

[Sarah completes Module 5 - 23 minutes]

✅ All Training Complete! 🎉

Progress: ▓▓▓▓▓▓▓▓▓▓ 4/4 modules (100%)

Total time: 88 minutes
Overall score: 97%

Generating certificate...
✓ Certificate saved: ~/Downloads/ucla-aws-training-sarah-chen.pdf

You now have full access to all Ark commands!
```

---

### Launching Her First Instance

```bash
$ ark instance launch --type m5.xlarge --purpose "Sequence alignment"

╔══════════════════════════════════════════════════════════════╗
║  💻 EC2 Instance Launch                                      ║
╚══════════════════════════════════════════════════════════════╝

→ Instance configuration:
  Type: m5.xlarge (4 vCPUs, 16GB RAM)
  Cost: $0.192/hour (~$140/month if running 24/7)
  
  ⚠️  COST WARNING: This will cost $4.61 per day if left running!

→ Recommended: Set auto-shutdown
  Shut down after idle for: [1h/4h/8h/12h/24h/never]: 4h

  ✓ Instance will auto-shutdown after 4 hours of inactivity
  💡 This could save you ~$100/month!

→ Selecting AMI (operating system)...
  ✓ Using: UCLA Bio-Linux 2024 (pre-configured for genomics)
  
  Includes: BWA, BLAST, samtools, Python 3.11, R 4.3

→ Configuring security group...
  ✓ SSH access from your IP only
  ✓ No public internet access (uses UCLA VPN)

→ Creating SSH key pair...
  ✓ Key saved: ~/.ssh/sarah-genomics-key.pem
  
  ⚠️  Keep this file safe! It's your password.

→ Launching instance...
  ⠋ Requesting capacity...
  ⠙ Instance starting...
  ⠹ Waiting for instance to be ready...
  
  ✓ Instance running: i-0123456789abcdef

╔══════════════════════════════════════════════════════════════╗
║  ✅ Instance Ready!                                          ║
╠══════════════════════════════════════════════════════════════╣
║  Instance ID: i-0123456789abcdef                             ║
║  Public IP: 34.216.45.123                                    ║
║  Status: Running                                             ║
║  Auto-shutdown: After 4h idle                                ║
╚══════════════════════════════════════════════════════════════╝

Connect via SSH:

  ssh -i ~/.ssh/sarah-genomics-key.pem ec2-user@34.216.45.123

Or use Ark's built-in SSH:

  ark instance connect --id i-0123456789abcdef

💡 Your S3 bucket is already mounted at /mnt/sarah-genomics-data/

⏱️  Total time: 3m 12s
```

---

### Week 2 - Sarah is Self-Sufficient

**What happened between Day 1 and Week 2:**

Over the past two weeks, Sarah has:
- ✓ Completed **Module 2 (IAM & Identity Management)** when she needed to give her grad student read-only access
- ✓ Uploaded 487 GB of genomic sequencing data to S3
- ✓ Launched and terminated 5 compute instances for different analyses
- ✓ Set up multiple billing alerts and stayed within her $150/month budget
- ✓ Shared data securely with two external collaborators
- ✓ Run weekly security audits (recommended practice)

**She's now using Ark routinely as part of her research workflow:**

```bash
$ ark audit scan

Running security audit on all your resources...

✅ Overall Security Score: 94/100

Findings:

✓ S3 Buckets (2)
  • sarah-genomics-data: Perfect ✓
  • sarah-results: Perfect ✓

✓ EC2 Instances (1)
  • i-0123456789abcdef: Auto-shutdown enabled ✓
  
⚠️  IAM
  • MFA not enabled on your user account
    Fix: ark iam mfa enable
    
💰 Cost Optimization
  • You could save $45/month by switching idle instances to t3.medium
    Current spend: $127/month (on track)
    Billing alarm: $100/month (triggered - check email!)

✓ Compliance
  • All resources properly tagged ✓
  • Access logging enabled ✓
  • Encryption at rest enabled ✓

📊 Full report: ~/ark-audit-report-2025-12-15.pdf
```

---

### Summary: Sarah's Journey

**Total Time Investment:**
- Installation & setup: 15 minutes
- Training: 120 minutes (2 hours, spaced over 2 days)
- First bucket creation: 5 minutes
- First instance launch: 10 minutes

**Total: ~2.5 hours** from zero to productive AWS researcher

**Training Breakdown:**
- Module 1: AWS Basics (35 min)
- Module 2: IAM & Identity Management (25 min) - completed Day 2
- Module 3: UC Data Classification (25 min)
- Module 4: S3 Storage Security (35 min)

**What Sarah Can Now Do:**
✓ Store and share research data securely  
✓ Launch compute instances for analysis  
✓ Monitor and control costs  
✓ Understand compliance requirements  
✓ Self-audit security posture  
✓ Know when to ask for help

**Security Incidents Prevented:**
- ✗ Unencrypted sensitive data
- ✗ Publicly accessible buckets  
- ✗ Runaway compute costs
- ✗ Missing audit trails
- ✗ Non-compliant data handling

**IT Support Tickets Avoided:**
- "How do I use AWS?" → Trained via tool
- "I forgot to shut down my instance!" → Auto-shutdown
- "My bill is $5000!" → Billing alarms + cost education
- "Is my data secure?" → Built-in compliance

---

## Appendix B: Module Template Structure

This appendix shows how training modules are structured and customized. Institutional administrators can modify these templates to meet specific requirements.

---

### Module Configuration File: `module.yaml`

```yaml
# Module metadata and configuration
module:
  id: "03-data-classification"
  version: "2.1.0"
  name: "UC Data Classification (P1-P4)"
  short_description: "Understanding data sensitivity and protection requirements"
  estimated_minutes: 25
  
  # What this module teaches
  learning_objectives:
    - "Classify data using UC P1-P4 framework"
    - "Identify PII and regulated data types"
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
  tutorial: "https://ucla-training.s3.amazonaws.com/modules/03/tutorial.md"
  quiz: "https://ucla-training.s3.amazonaws.com/modules/03/quiz.yaml"
  scenarios: "https://ucla-training.s3.amazonaws.com/modules/03/scenarios.yaml"
  resources:
    - name: "UC Data Classification Policy"
      url: "https://policy.ucop.edu/doc/7000543/BFB-IS-3"
    - name: "FERPA Quick Reference"
      url: "https://it.ucla.edu/security/ferpa"
    - name: "HIPAA Compliance Guide"
      url: "https://hipaa.ucla.edu/guide"

# Institution-specific customization
customization:
  institution: "UCLA"
  ciso_sponsor: true  # Indicates CISO office reviewed this module
  
  # Institution-specific sections to inject
  custom_sections:
    - position: "after_intro"
      content_url: "https://ucla-training.s3.amazonaws.com/modules/03/ucla-specific.md"
      title: "UCLA-Specific Requirements"
    
    - position: "before_quiz"
      content_url: "https://ucla-training.s3.amazonaws.com/modules/03/case-studies.md"
      title: "Real UCLA Incidents (Anonymized)"
  
  # Override default contact information
  contacts:
    questions: "your institutional data classification office"
    incidents: "your institutional security incident response team"
    compliance: "your institutional HIPAA compliance office"

# Analytics and tracking
tracking:
  record_time_spent: true
  record_quiz_attempts: true
  record_common_mistakes: true
  send_completion_to: "https://ucla-training.ucla.edu/api/completion"

# Compliance attestation
compliance:
  reviewed_by: "UCLA CISO Office"
  review_date: "2025-11-15"
  next_review: "2026-05-15"
  frameworks_covered:
    - "UC BFB-IS-3 (Data Classification)"
    - "FERPA"
    - "HIPAA"
    - "NIST 800-171 (CUI)"
```

---

### Tutorial Content: `tutorial.md`

Tutorial content is written in enhanced Markdown with special syntax for interactive elements:

```markdown
# Module 3: UC Data Classification

## Section 1: Introduction

:::info
This module is sponsored by the UCLA CISO Office to ensure 
all researchers understand data protection requirements.
:::

Understanding data sensitivity is CRITICAL for compliance and security.

:::warning title="Getting This Wrong Has Consequences"
- Federal fines ($100k - $50M+ per incident)
- Loss of grant funding
- IRB suspension
- Legal liability
- Reputational damage to UCLA
:::

---

## Section 2: The Four Protection Levels

:::classification level="P1"
### P1 - Public Information

**Definition**: Information intended for public distribution

**Examples**:
- Published research papers
- Public course catalogs
- Campus maps

**Requirements**: None (already public)

**AWS Configuration**: Standard S3, no special controls
:::

:::classification level="P3" highlight="true"
### P3 - Protected Information ⭐ MOST COMMON

**Definition**: Sensitive data requiring protection

**Examples**:
- Personal Identifiable Information (PII)
  - Names, addresses, phone numbers
  - Email addresses, student IDs
- Student records (FERPA protected)
- De-identified health data
- Research data under confidentiality agreements

**Legal Frameworks**:
- FERPA (Family Educational Rights and Privacy Act)
- PII protection laws (CCPA, GDPR if applicable)
- Contractual confidentiality obligations

**Requirements**:
✓ Encryption at rest (REQUIRED)
✓ Encryption in transit (REQUIRED)
✓ Access logging for audits
✓ Strong access controls
✓ MFA for administrators

**AWS Configuration**: Ark P3 configuration enforces ALL requirements

:::alert type="danger"
**Common Mistake**: "It's de-identified so it's fine"

Even de-identified data can often be re-identified! When in doubt, treat as P3.
:::
:::

---

## Section 3: Real World Examples

:::case-study severity="high"
### Case Study: The "Anonymous" Survey Breach

**Institution**: Major Research University (2024)

**Scenario**: Researcher collected "anonymous" health surveys

**What they included**:
- Zip code
- Age (exact)
- Gender
- Medical condition

**The Problem**: These 4 data points can identify ~87% of the US population

**Classification Error**: Treated as P2, actually P3 (possibly P4!)

**Consequence**: 
- Data breach notification to 1,200 participants
- $250,000 fine
- IRB investigation
- 6-month research suspension

**Lesson**: Combinations of "non-sensitive" data = sensitive data
:::

---

## Section 4: Interactive Exercise

:::interactive type="classification-exercise"
**Exercise**: Classify this dataset

You have a dataset containing:
- Genomic sequences from fruit flies (Drosophila)
- No human subjects
- No personally identifiable information
- Funded by NSF grant
- Will be published when analysis complete
- No confidentiality agreements

What classification level?

[P1] [P2] [P3] [P4]

:::feedback correct="P2"
✅ **Correct!** This is P2 (Internal) because:

- Not yet published (so not P1)
- No PII or regulated data (so not P3/P4)
- Internal research data until publication
- Non-human subject research

**When you publish, you can reclassify to P1.**

:::tip
Use this command to reclassify later:
```bash
ark bucket reclassify --name my-bucket --from P2 --to P1
```
:::
:::

---

## Section 5: Decision Tree

:::decision-tree
# Data Classification Decision Tree

**Start**: Does your data contain ANY of the following?

- Names, email addresses, phone numbers?
  → YES: **At least P3**
  
- Student records or grades?
  → YES: **P3 (FERPA applies)**
  
- Health information (even de-identified)?
  → YES: **At least P3, possibly P4 if identifiable**
  
- Financial data, SSNs, credit cards?
  → YES: **P4 (stop, contact CISO office)**
  
- Under confidentiality agreement?
  → YES: **Read agreement, probably P3**
  
- Export controlled (ITAR/EAR)?
  → YES: **P4 (stop, contact export control office)**
  
- Will be published/public eventually?
  → YES but not yet: **P2 until published**
  
- None of the above?
  → **Probably P1 or P2, but verify with PI**

:::help
**Still unsure?** That's OK!

Contact: your institutional data classification office

Better to ask than to misclassify.
:::
:::

---

## Section 6: UCLA-Specific Requirements

<!-- This section is injected from custom_sections in module.yaml -->

{{% custom_section position="after_intro" %}}

---

## Checkpoint Quiz

You must score 85% or higher to proceed.

{{% quiz source="quiz.yaml" %}}
```

---

### Quiz Definition: `quiz.yaml`

```yaml
# Quiz configuration
quiz:
  id: "03-data-classification-quiz"
  passing_score: 85
  randomize_questions: true
  randomize_answers: true
  max_attempts: 3
  show_correct_answers_after: 2  # After 2 attempts, show correct answers

questions:
  - id: "q1-identify-p3"
    type: "multiple_choice"
    points: 20
    question: |
      You have a dataset with: age (binned in 5-year ranges), 
      zip code, and diagnosis. No names. What classification level?
    
    options:
      - id: "a"
        text: "P1 - It's de-identified"
        correct: false
      
      - id: "b"
        text: "P2 - Internal use only"
        correct: false
      
      - id: "c"
        text: "P3 - Can be re-identified"
        correct: true
      
      - id: "d"
        text: "P4 - Contains health info"
        correct: false
    
    feedback:
      correct: |
        ✅ CORRECT! Even without names, this is P3 because:
        
        - Age + zip code + diagnosis = potentially identifiable
        - Health information requires protection even when de-identified
        - Could violate HIPAA if re-identified
        
        This is called "quasi-identifiers" - seemingly anonymous
        data that can be combined to identify individuals.
      
      incorrect: |
        ❌ Not quite. Consider:
        
        - Can these data points identify someone?
        - What if combined with public databases?
        - Does it contain health information?
        
        Think about re-identification risk.
    
    resources:
      - "https://www.hhs.gov/hipaa/for-professionals/privacy/special-topics/de-identification/"

  - id: "q2-data-sharing"
    type: "multiple_choice"
    points: 20
    question: |
      Your collaborator at Stanford needs access to your P3 research data.
      What do you need?
    
    options:
      - id: "a"
        text: "Just share an S3 link"
        correct: false
      
      - id: "b"
        text: "Data sharing agreement + BAA if needed"
        correct: true
      
      - id: "c"
        text: "Their email address"
        correct: false
      
      - id: "d"
        text: "Nothing special - they're at a university"
        correct: false
    
    feedback:
      correct: |
        ✅ PERFECT! For P3 data sharing, you need:
        
        1. Data Sharing Agreement (legal framework)
        2. Business Associate Agreement if health data (HIPAA)
        3. Document what data is shared and why
        4. Time-limited access (not permanent)
        5. UCLA IRB approval if human subjects
        
        Contact: your institutional data sharing office for templates

  - id: "q3-classification-error"
    type: "multiple_choice"
    points: 20
    question: |
      🚨 COMPLIANCE SCENARIO: You discover you've been storing
      what you thought was P2 data, but it actually contains 
      email addresses (P3). What do you do?
    
    options:
      - id: "a"
        text: "Delete the emails and move on"
        correct: false
      
      - id: "b"
        text: "Immediately report to CISO, re-classify, audit access"
        correct: true
      
      - id: "c"
        text: "Just fix it going forward"
        correct: false
      
      - id: "d"
        text: "Hope no one noticed"
        correct: false
    
    feedback:
      correct: |
        ✅ CRITICAL! When you discover a classification error:
        
        IMMEDIATE actions:
        1. Stop any current data sharing
        2. Email: your institutional security incident response team
        3. Document: What data? How long? Who had access?
        
        The CISO office will help you:
        - Re-classify correctly
        - Audit who accessed the data
        - Implement proper controls
        - Determine if breach notification needed
        
        🎯 Key principle: It's never wrong to report.
    
    tags: ["incident-response", "compliance", "critical"]

  - id: "q4-bucket-mixing"
    type: "multiple_choice"
    points: 20
    question: |
      Can you mix P2 and P3 data in the same S3 bucket?
    
    options:
      - id: "a"
        text: "Yes, it's fine"
        correct: false
      
      - id: "b"
        text: "Yes, but in separate folders"
        correct: false
      
      - id: "c"
        text: "No, always use separate buckets"
        correct: true
      
      - id: "d"
        text: "Only with special permission"
        correct: false
    
    feedback:
      correct: |
        ✅ CORRECT! Best practice: Separate buckets per classification.
        
        Why?
        - Bucket-level encryption settings differ
        - Access controls are simpler
        - Audit logging is clearer
        - Reduces accidental exposure risk
        - Compliance audits are easier

  - id: "q5-scenario"
    type: "scenario"
    points: 20
    question: |
      **Scenario**: You're analyzing survey data that includes:
      - Participant ID (non-identifiable code)
      - County of residence
      - Year of birth
      - Political affiliation
      - Voting history
      
      This data will inform policy recommendations. How do you classify it?
    
    correct_classification: "P3"
    
    reasoning_required: true
    min_reasoning_length: 50
    
    sample_reasoning: |
      This should be classified as P3 because:
      
      1. Demographic data (county, year of birth) combined with
         sensitive information (political affiliation, voting history)
         could potentially identify individuals
      
      2. Political information is sensitive even when aggregated
      
      3. While participant IDs are non-identifiable, the combination
         of other factors creates re-identification risk
      
      4. Policy recommendations may involve sensitive populations
    
    grading_rubric:
      - keyword: "re-identification"
        points: 5
      - keyword: "sensitive"
        points: 3
      - keyword: "combination"
        points: 3
      - mentions_demographics: true
        points: 4
      - mentions_political_sensitivity: true
        points: 5

# Post-quiz feedback
post_quiz:
  pass:
    message: |
      🎉 Excellent work! You've demonstrated strong understanding
      of UCLA's data classification framework.
      
      Score: {score}%
      
      You can now create S3 buckets with proper classification.
    
    next_steps:
      - "Continue to Module 4: S3 Storage Security"
      - "Download classification quick reference"
      - "Bookmark: your institutional data classification office"
  
  fail:
    message: |
      You scored {score}% (need 85% to pass).
      
      Don't worry! Data classification is complex.
      
      Review these sections:
      {weak_areas}
      
      You have {attempts_remaining} attempt(s) remaining.
    
    resources:
      - "Review the decision tree in Section 5"
      - "Read the case studies in Section 3"
      - "Contact your institutional data classification office for help"
```

---

### Scenarios: `scenarios.yaml`

```yaml
# Interactive scenario-based learning
scenarios:
  - id: "scenario-classify-research-data"
    title: "Classify Your Research Data"
    description: |
      Walk through a realistic data classification scenario
    
    steps:
      - prompt: "What type of research data do you work with?"
        options:
          - "Human subjects research"
          - "Animal research"
          - "Computational/modeling"
          - "Materials science"
          - "Other"
        branch_on_selection: true
      
      - prompt: "Does your data contain any identifiable information about individuals?"
        type: "yes_no"
        if_yes:
          guidance: "This likely requires P3 or P4 classification"
          next: "identify-pii-types"
        if_no:
          next: "check-other-sensitive"
      
      # ... more scenario steps ...
    
    completion:
      requires_correct_classification: true
      provides_certificate: true
```

---

### Customization Points for Institutions

Institutions can customize:

1. **Content URLs**: Host training materials on their own infrastructure
2. **Custom sections**: Inject institution-specific content anywhere
3. **Contact information**: Override default support contacts
4. **Passing scores**: Adjust based on risk tolerance
5. **Case studies**: Add institution-specific incidents (anonymized)
6. **Legal frameworks**: Emphasize relevant compliance requirements
7. **Quiz questions**: Add institution-specific scenarios
8. **Resources**: Link to internal policies and procedures
9. **CISO sponsorship**: Mark modules as officially reviewed
10. **Compliance tracking**: Send completion data to institutional systems

---

### Example: Adapting for Different Institutions

**UCLA Version** (Current):
- UC P1-P4 classification
- FERPA, HIPAA, CUI emphasis
- UCLA CISO contacts
- UCLA-specific case studies

**MIT Version** (Hypothetical):
```yaml
customization:
  institution: "MIT"
  classification_system: "TLP"  # Traffic Light Protocol
  frameworks:
    - "ITAR/EAR" # Export control emphasis
    - "CMMC Level 2"
    - "DFARS"
  contacts:
    questions: "institutional information security"
    incidents: "institutional CERT team"
```

**NIH Intramural** (Hypothetical):
```yaml
customization:
  institution: "NIH"
  classification_system: "NIH-specific"
  frameworks:
    - "FISMA High"
    - "FedRAMP"
    - "HIPAA (strict)"
  additional_requirements:
    - "All data is P3 minimum"
    - "Requires ISSO approval"
```

---

### Benefits of This Template System

**For Ark Tool:**
- Consistent learning experience
- Programmatic content validation
- Automated progress tracking
- Easy A/B testing of content

**For Institutions:**
- Full content control
- Rapid deployment of updates
- Compliance audit trail
- Multi-tenancy support

**For Learners:**
- Always current content
- Institution-relevant examples
- Consistent UI/UX
- Offline capability (cached content)

---
