# Ark: Gaps Analysis & Recommendations
**Analysis Date**: December 2025
**Project**: Ark AWS Research Kit (Open Source)
**License**: Apache 2.0

---

## Intended Audience for This Document

This analysis is intended for:
- **Technical architects** designing and implementing Ark
- **Project managers** planning deployment timeline and resources
- **Security officers** evaluating risk and compliance implications
- **Institutional decision-makers** assessing feasibility and ROI
- **CISO Office staff** ensuring alignment with security policies
- **Open source contributors** understanding architecture and design decisions

This document is **NOT** for:
- End users (researchers) - see the main Ark overview (README.md) instead
- General stakeholders requiring executive summary only
- Marketing/communications teams

**Prerequisites for readers**: Understanding of AWS services, institutional IT infrastructure, academic research workflows, and web/CLI application architecture.

---

## Executive Summary

This document analyzes the Ark proposal to identify gaps, risks, and areas requiring further consideration before implementation. The analysis is organized by category with specific recommendations for each gap.

**Key Updates**:
- **Dual Interface**: Both CLI and web (Vue/Cloudscape) developed simultaneously
- **Open Source**: Apache 2.0 license, community-driven
- **Brand Flexibility**: Institutional-agnostic framework with customization
- **Architecture**: Local agent + institutional backend model
- **Timeline**: See [ROADMAP.md](ROADMAP.md) for detailed implementation schedule

---

## 1. Technical Architecture & Implementation

### Architecture Overview

Ark consists of three primary components:

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

**Key Design Principles**:
1. **Security boundary**: AWS credentials never leave user's machine
2. **Unified backend**: CLI and web share same institutional API
3. **Offline-capable**: Training content cached locally
4. **Brand-flexible**: Institutions customize branding, policies, content

---

### Priority Definitions

Before diving into specific gaps, understand how priorities are assigned:

**CRITICAL (Must Address Before Launch)**
- **Criteria**: Will cause system failure, security breach, or compliance violation
- **Impact**: Complete blocker for production deployment
- **Timeline**: Must be resolved during pilot phase
- **Examples**: Data architecture, anti-bypass measures, audit logging, dual interface consistency

**HIGH PRIORITY (Address in First 6 Months)**
- **Criteria**: Will cause poor user experience or significant operational burden
- **Impact**: Limits adoption, increases support costs, reduces effectiveness
- **Timeline**: Must be resolved before broad rollout
- **Examples**: Identity integration, accessibility (WCAG 2.1 AA), SIEM integration, web UI polish

**MEDIUM PRIORITY (Address in First Year)**
- **Criteria**: Would enhance usability or reduce long-term maintenance burden
- **Impact**: Affects scalability and sustainability
- **Timeline**: Should be resolved before declaring "production ready"
- **Examples**: Spaced repetition, multi-language support, advanced web features

**FUTURE CONSIDERATIONS**
- **Criteria**: Nice-to-have enhancements for future versions
- **Impact**: Minimal on core functionality
- **Timeline**: Version 2.0 or later
- **Examples**: Mobile app, VR training integration, AI-powered assistance

---

### Implementation Dependency Map

Many gaps depend on each other. Understanding these dependencies is critical for project planning.

```
CRITICAL PATH (Sequential Dependencies):

Phase 1: Foundation
├─ 1.1 Data Architecture → Everything depends on this
├─ 1.5 Agent Architecture → Required for both CLI and web
├─ 1.6 Web Framework Setup → Vue/Cloudscape foundation
├─ 4.1 Identity Integration → Required for user tracking
└─ 2.1 Audit Logging → Required for compliance proof

Phase 2: Core Security
├─ 1.2 Anti-Bypass Measures → Depends on 1.1, 1.5
├─ 2.2 Incident Response → Depends on 2.1
├─ 4.2 AWS Organizations → Depends on 4.1
└─ 1.7 Web Security → XSS, CSRF protection

Phase 3: User Experience
├─ 3.1 Support Model → Can start after Phase 1
├─ 3.3 Accessibility → Critical for web, important for CLI
├─ 1.8 Cross-Interface Sync → CLI ↔ web consistency
└─ 5.1 Learning Science → Depends on pilot feedback

Phase 4: Operations
├─ 6.1 Success Metrics → Depends on 2.1
├─ 3.2 Cost Model → Depends on usage data
├─ 6.2 Feedback Loops → Depends on user base
└─ 1.9 E2E Testing → Playwright across both interfaces

PARALLEL TRACKS (Can be developed concurrently):
• Backend API development
• CLI development
• Web development (Vue/Cloudscape)
• Content & Pedagogy (Gap 5.x)
• Documentation & Training materials
```

**Key Insight**: Backend API is the critical dependency. Once stable, CLI and web can develop in parallel. Testing strategy must cover both interfaces.

---

### Gap 1.1: Progress Tracking Implementation Details

**Missing:**
- How is training progress stored? (Local files, S3, DynamoDB, institutional database?)
- What happens if progress data is lost or corrupted?
- How do you handle users on multiple machines?
- How does CLI and web sync progress in real-time?
- Sync mechanism for offline/online transitions

**Recommendation:**
```yaml
Proposed Architecture:

Primary Storage: Institutional Backend (DynamoDB)
  Table: training_progress
    - Partition key: user_id
    - Sort key: module_id
    - Attributes:
        status: "not_started" | "in_progress" | "completed"
        score: 0-100
        attempts: number
        time_spent_seconds: number
        completed_at: ISO timestamp
        certificate_url: S3 URL
        last_synced: ISO timestamp

Local Cache: Agent (SQLite)
  Location: ~/.ark/cache.db
  Tables:
    - training_progress (mirror of backend)
    - modules (content cache)
    - sync_queue (pending operations)

Sync Strategy:
  - Agent queries backend every 5 minutes (polling)
  - Web uses GraphQL subscriptions (real-time)
  - CLI triggers sync on command execution
  - Offline: Queue operations, sync when reconnected
  - Conflict resolution: Backend always wins (last-write-wins with timestamp)

Cross-Interface Flow:
  1. User starts training in CLI
  2. Agent writes to local cache + queues backend update
  3. Agent syncs to backend (async)
  4. Web subscribes to user's progress (GraphQL subscription)
  5. Web receives update via WebSocket
  6. User sees updated progress in web dashboard

Backup: S3 with versioning
  - Full audit trail
  - Recovery capability
  - Compliance evidence
```

**Action Items:**
- [ ] Define complete data schema for progress tracking
- [ ] Implement sync protocol (agent ↔ backend)
- [ ] Add GraphQL subscriptions for web real-time updates
- [ ] Design conflict resolution strategy (with edge cases)
- [ ] Plan for data retention and privacy compliance
- [ ] Create backup and disaster recovery plan
- [ ] Test multi-device scenarios
- [ ] Test CLI ↔ web sync latency (<5 seconds target)

---

### Gap 1.2: Training Bypass Prevention

**Missing:**
- What prevents users from manipulating local progress files?
- Can users share completion certificates?
- How do you verify CloudTrail logs aren't forged?
- What about time-based attacks (system clock manipulation)?
- Can users inspect agent traffic and replay requests?

**Recommendation:**
```
Anti-Bypass Measures (Defense in Depth):

Layer 1: Server-Side Validation (CRITICAL)
  - All progress stored server-side with cryptographic signing
  - Local cache is read-only display, not source of truth
  - Agent validates all operations with backend
  - Backend checks training completion before allowing commands
  - Certificates include HMAC signature (signed by backend private key)

Layer 2: CloudTrail Cross-Reference
  - Backend verifies API calls with CloudTrail digest files
  - Check IP addresses and user agents for anomalies
  - Require recent AWS activity (within last 7 days of training completion)
  - Detect if user is creating resources without Ark (bypass attempt)

Layer 3: Behavioral Analysis
  - Flag suspiciously fast completions (<50% of expected time)
  - Detect answer patterns (all correct in <1 second = bot)
  - Require minimum time per section (e.g., video must play to completion)
  - Mouse/keyboard interaction patterns (web only)
  - Random challenge questions (not in training material)

Layer 4: Agent Security
  - Agent validates backend responses with signature verification
  - HTTPS only communication (TLS 1.3)
  - Certificate pinning for institutional backend
  - Rate limiting on agent API (prevent brute force)
  - Request nonces (prevent replay attacks)

Layer 5: Certificate Binding
  - Certificates contain:
      * User ID (eduPersonPrincipalName)
      * AWS Account ID
      * Completion timestamp
      * Module checksums
      * Cryptographic signature (RSA-4096 or Ed25519)
  - Can't be transferred or reused
  - Verifiable by third parties (public key published)
  - Includes institutional seal/logo (visual verification)

Layer 6: Audit Trail
  - All training attempts logged (immutable)
  - Failed bypass attempts flagged for manual review
  - Impossible to delete logs (write-only from user perspective)
  - Security team can review suspicious patterns

Example Flow (S3 Bucket Creation):
  User: ark bucket create --name test
  CLI → Agent: POST /api/bucket/create
  Agent → Backend: GET /api/policy/can-create-bucket?user={id}&operation=s3:CreateBucket
  Backend checks:
    1. Is training completed? (query training_progress table)
    2. Is certificate valid? (verify signature)
    3. Is training recent? (<90 days, or requires refresher)
    4. Any anomalies? (check behavioral flags)
    5. Sign response with HMAC
  Backend → Agent: { allowed: true, signature: "..." }
  Agent validates signature
  Only if valid: Agent → AWS: CreateBucket API call
  Agent → Backend: POST /api/audit/event (log operation)
```

**Action Items:**
- [ ] Implement server-side progress verification
- [ ] Add cryptographic signing to all backend responses
- [ ] Implement certificate generation with crypto signatures
- [ ] Add tamper detection to local storage
- [ ] Create anomaly detection rules (ML-based for v2)
- [ ] Design certificate validation system
- [ ] Test bypass attempts (red team exercise)
- [ ] Document security model for auditors

---

### Gap 1.3: Offline Functionality Scope

**Missing:**
- Which features work offline vs require network?
- How long can users work offline before sync required?
- What happens when AWS APIs are unavailable?
- Can training modules be completed fully offline?
- How does web app behave offline? (Service worker?)

**Recommendation:**
```
Offline Capability Matrix:

FULLY OFFLINE (CLI):
  ✓ Read training content (cached in ~/.ark/cache/)
  ✓ Take quizzes (local validation, queued for backend confirmation)
  ✓ View previous progress (from local cache)
  ✓ Read help documentation
  ✓ View certificates (previously downloaded)

REQUIRES NETWORK (CLI):
  ✗ Download new training modules (initial download or updates)
  ✗ Final progress validation (backend must confirm quiz answers)
  ✗ Validate AWS credentials
  ✗ Execute actual AWS operations (create buckets, launch instances)
  ✗ Generate certificates (requires backend signing)

WEB OFFLINE CAPABILITIES:
  ✓ View cached dashboard (Service Worker)
  ✓ Read previously viewed training modules (IndexedDB)
  ✓ Limited navigation (cached routes)

  ✗ Live updates (requires WebSocket connection)
  ✗ Most write operations
  ✗ Real-time sync

Graceful Degradation Strategy:
  - Queue operations for later sync (show "Pending Sync" badge)
  - Show clear offline indicators in UI (both CLI and web)
    CLI: "⚠️  Offline Mode - Changes will sync when reconnected"
    Web: Cloudscape Banner component with warning
  - Warn about operations requiring network before attempting
  - Cache last 30 days of content automatically
  - Progressive Web App (PWA) for web interface:
      * Service Worker for offline HTML/CSS/JS
      * IndexedDB for training content
      * Background sync when reconnected

Example User Flow (Offline):
  1. User completes Module 3 on airplane (offline)
  2. CLI saves to local SQLite: status="completed_offline"
  3. CLI shows: "✓ Module 3 complete (will sync when online)"
  4. User lands, connects to WiFi
  5. Agent detects connectivity, starts sync
  6. Agent → Backend: POST /api/training/sync (batch of operations)
  7. Backend validates quiz answers
  8. Backend responds: { status: "confirmed", certificate_url: "..." }
  9. Agent updates local cache: status="completed"
  10. Certificate auto-downloads
  11. CLI shows: "✓ Module 3 completion confirmed - Certificate ready"
```

**Action Items:**
- [ ] Define complete offline functionality matrix
- [ ] Implement operation queue system (SQLite)
- [ ] Design offline UI indicators (CLI and web)
- [ ] Implement Service Worker for web PWA
- [ ] Add background sync API (web)
- [ ] Test with flaky networks (throttled, intermittent)
- [ ] Test complete offline → online transition
- [ ] Document offline limitations clearly

---

### Gap 1.4: Update and Rollback Strategy

**Missing:**
- How are tool updates deployed at scale?
- What if an update breaks existing workflows?
- Can institutions pin to specific versions?
- How do you communicate breaking changes?
- Do CLI and web need to stay version-synchronized?

**Recommendation:**
```
Update Strategy:

Versioning (Semantic):
  - MAJOR.MINOR.PATCH (e.g., 1.2.3)
  - MAJOR: Breaking changes (rare, requires migration)
  - MINOR: New features (backward compatible)
  - PATCH: Bug fixes

  - LTS channels: 1.x, 2.x (2-year support each)
  - Beta channel for early adopters
  - Nightly builds for developers

Component Version Compatibility:
  - Agent version: X.Y.Z
  - CLI version: X.Y.Z (must match agent MAJOR.MINOR)
  - Web version: X.Y.Z (must match agent MAJOR.MINOR)
  - Backend API version: vX (supports multiple client versions)

  Example:
    - Agent 1.2.5 + CLI 1.2.3 = ✓ Compatible
    - Agent 1.2.5 + CLI 1.3.0 = ✗ Minor mismatch warning
    - Agent 1.2.5 + CLI 2.0.0 = ✗ Blocked (incompatible)
    - Backend API v1 supports: Agent 1.x, CLI 1.x, Web 1.x
    - Backend API v2 supports: Agent 2.x, CLI 2.x, Web 2.x

Deployment Strategy:
  CLI:
    - Auto-update with user consent (prompt on launch if outdated)
    - Manual update: ark update
    - Package managers: brew upgrade ark, apt update ark
    - Institutions can disable auto-update (config setting)

  Web:
    - Auto-deploy on visit (SPA, always latest)
    - Version displayed in footer
    - Browser caching with cache-busting
    - Graceful degradation if agent is outdated

  Agent:
    - Auto-update with user consent (requires restart)
    - Manual update: ark agent update
    - Packaged with CLI installer

  Backend:
    - Blue-green deployment (zero downtime)
    - API versioning (support N-1 versions)
    - Gradual rollout (canary → 10% → 50% → 100%)

Admin-Controlled Rollout:
  - Institutions can configure:
      auto_update: false  # Require manual updates
      version_pin: "1.2.x"  # Allow only patch updates
      update_channel: "stable"  # vs "beta" or "nightly"
  - Admin dashboard shows version distribution
  - Force update capability (for security patches)

Rollback Strategy:
  - Keep last 3 versions cached locally (~/.ark/versions/)
  - One-command rollback: ark version rollback
  - Rollback preserves user data (backwards compatible)
  - Server-side kill switch for broken versions:
      Backend API: /api/version/compatibility
      Response: {
        "1.5.0": { "status": "deprecated", "message": "Critical bug, please update" },
        "1.5.1": { "status": "active" }
      }
  - Agent checks on launch, warns or blocks if version deprecated

Communication:
  - Release notes in-tool:
      CLI: ark changelog
      Web: "What's New" modal on first launch after update
  - Email notifications for breaking changes (90 days advance)
  - Deprecation warnings (90 days before removal)
  - In-app banners: "New version available: see what's new"
  - Community forum: Release announcements
  - GitHub releases: Detailed changelogs

Breaking Change Example:
  Version 1.x: Data classification levels: P1, P2, P3, P4
  Version 2.0: Data classification levels: Public, Internal, Confidential, Restricted

  Migration:
    1. Announce 90 days in advance
    2. Version 1.9.0: Add migration tool (ark migrate to-v2)
    3. Version 1.9.0: Dual support (both naming schemes work)
    4. Version 2.0.0: Drop old naming scheme
    5. Version 2.0.0: Auto-migrate on first launch (with user confirmation)
```

**Action Items:**
- [ ] Design update distribution system (CDN, GitHub releases)
- [ ] Implement version compatibility checking
- [ ] Create version compatibility testing matrix
- [ ] Build rollback mechanism (agent + CLI)
- [ ] Design kill switch API (backend)
- [ ] Establish change communication process
- [ ] Create migration tool framework
- [ ] Test update flows (success and failure scenarios)

---

### Gap 1.5: Local Agent Architecture

**New Gap (Dual Interface Requirement):**
- Agent must serve both CLI and web reliably
- How to handle concurrent requests from both interfaces?
- Security model for localhost API?
- How to prevent other local processes from accessing agent API?

**Recommendation:**
```
Agent Architecture:

Technology:
  - Language: Go (single binary, cross-platform)
  - HTTP server: net/http (stdlib, battle-tested)
  - Port: 8737 (fixed, localhost only)
  - Storage: SQLite (~/.ark/cache.db)
  - Configuration: YAML (~/.ark/config.yaml)

Security Model (localhost-only):
  - Bind to 127.0.0.1:8737 (not 0.0.0.0)
  - Require authentication token in headers:
      Authorization: Bearer {token}
  - Token stored in ~/.ark/credentials (mode 0600, encrypted at rest)
  - Token rotates every 24 hours (refresh flow)
  - CORS: Only localhost origins allowed
  - CSRF protection: Custom headers required

API Design:
  RESTful with clear endpoints:

  Authentication:
    POST /api/auth/login → Redirect to institutional SSO
    POST /api/auth/refresh → Refresh tokens
    POST /api/auth/logout → Invalidate tokens

  Training:
    GET /api/training/modules → List available modules
    GET /api/training/progress → User's progress
    POST /api/training/start → Begin module
    POST /api/training/complete → Submit answers
    GET /api/training/certificate/:id → Download certificate

  AWS Operations:
    POST /api/aws/bucket/create → Create S3 bucket (with policy checks)
    POST /api/aws/instance/launch → Launch EC2 instance
    GET /api/aws/resources → List user's resources

  Audit:
    GET /api/audit/logs → View audit logs

  System:
    GET /api/system/health → Agent health check
    GET /api/system/version → Version info
    POST /api/system/sync → Force sync with backend

Concurrent Request Handling:
  - Go's goroutines handle concurrency naturally
  - Request queue with priority:
      High: Auth, health checks
      Medium: AWS operations
      Low: Background syncs
  - Rate limiting per endpoint (prevent abuse)
  - Graceful degradation under load

State Management:
  - Stateless agent (all state in SQLite or backend)
  - Connection pooling for backend API
  - Caching with TTL (reduce backend load)

Lifecycle:
  Start:
    - CLI: ark agent start (if not running)
    - Installer: Register as system service (systemd, launchd, Windows Service)
    - Auto-start on boot (optional, user configurable)

  Stop:
    - CLI: ark agent stop
    - Graceful shutdown (finish pending requests, sync to backend)

  Restart:
    - After updates
    - On configuration changes
    - On crash (auto-restart with exponential backoff)

Logging:
  - Agent logs: ~/.ark/logs/agent.log
  - Rotate daily, keep 30 days
  - Levels: DEBUG, INFO, WARN, ERROR
  - Structured logging (JSON for parsing)
  - Redact sensitive data (credentials, PII)

Monitoring:
  - Health endpoint: GET /api/system/health
      Response: { "status": "healthy", "uptime": 86400, "version": "1.2.0" }
  - Metrics: Prometheus endpoint (optional)
  - Web dashboard: Shows agent status
```

**Action Items:**
- [ ] Implement HTTP server with security headers
- [ ] Design token authentication system
- [ ] Build SQLite cache layer
- [ ] Implement rate limiting
- [ ] Create system service installers (all platforms)
- [ ] Add health check endpoint
- [ ] Test concurrent CLI + web usage
- [ ] Document agent API (OpenAPI spec)

---

### Gap 1.6: Web Framework Setup (Vue/Cloudscape)

**New Gap (Web Interface):**
- How to integrate Cloudscape with Vue 3?
- Component library organization?
- State management strategy?
- Build and deployment approach?
- Testing strategy for Vue components?

**Recommendation:**
```
Web Application Stack:

Core Technologies:
  - Framework: Vue 3 (Composition API)
  - UI Library: AWS Cloudscape Design System
  - Language: TypeScript (type safety)
  - Build Tool: Vite (fast dev server, optimized builds)
  - State Management: Pinia (Vue's official state library)
  - Router: Vue Router 4
  - HTTP Client: Axios (agent API communication)
  - Testing: Vitest (unit), Playwright (E2E)
  - Linting: ESLint + Prettier
  - Package Manager: pnpm (fast, disk-efficient)

Project Structure:
```
web/
├── src/
│   ├── assets/          # Static assets (images, fonts)
│   ├── components/      # Reusable Vue components
│   │   ├── training/    # Training-specific components
│   │   ├── resources/   # AWS resource components
│   │   └── common/      # Shared components
│   ├── composables/     # Vue composables (reusable logic)
│   │   ├── useAgent.ts  # Agent API client
│   │   ├── useAuth.ts   # Authentication logic
│   │   └── useTraining.ts # Training state
│   ├── layouts/         # Page layouts
│   ├── pages/           # Route pages
│   ├── stores/          # Pinia stores
│   │   ├── auth.ts      # Auth state
│   │   ├── training.ts  # Training progress
│   │   └── resources.ts # AWS resources
│   ├── router/          # Vue Router configuration
│   ├── styles/          # Global styles
│   ├── types/           # TypeScript types
│   ├── utils/           # Utility functions
│   ├── App.vue          # Root component
│   └── main.ts          # Entry point
├── public/              # Static files (served as-is)
├── tests/
│   ├── unit/            # Vitest unit tests
│   └── e2e/             # Playwright E2E tests
├── playwright.config.ts
├── vite.config.ts
├── tsconfig.json
└── package.json
```

Cloudscape Integration:
  Install:
    npm install @cloudscape-design/components
    npm install @cloudscape-design/global-styles

  Usage in Vue:
```typescript
<script setup lang="ts">
import {
  Button,
  Container,
  Header,
  SpaceBetween,
  Table
} from '@cloudscape-design/components'
import { useTraining } from '@/composables/useTraining'

const { modules, isLoading } = useTraining()
</script>

<template>
  <Container :header="Header({ title: 'Training Modules' })">
    <SpaceBetween size="l">
      <Table
        :items="modules"
        :loading="isLoading"
        :columnDefinitions="[
          { id: 'name', header: 'Module', cell: item => item.name },
          { id: 'status', header: 'Status', cell: item => item.status },
          { id: 'score', header: 'Score', cell: item => item.score }
        ]"
      />
      <Button @click="startTraining">Start Training</Button>
    </SpaceBetween>
  </Container>
</template>
```

State Management (Pinia):
```typescript
// stores/training.ts
import { defineStore } from 'pinia'
import { agentAPI } from '@/composables/useAgent'

export const useTrainingStore = defineStore('training', () => {
  const modules = ref<TrainingModule[]>([])
  const currentModule = ref<TrainingModule | null>(null)
  const loading = ref(false)

  async function fetchModules() {
    loading.value = true
    try {
      const response = await agentAPI.get('/api/training/modules')
      modules.value = response.data
    } finally {
      loading.value = false
    }
  }

  async function startModule(moduleId: string) {
    currentModule.value = modules.value.find(m => m.id === moduleId) || null
    await agentAPI.post(`/api/training/start`, { moduleId })
  }

  return { modules, currentModule, loading, fetchModules, startModule }
})
```

Routing:
```typescript
// router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: () => import('@/pages/Dashboard.vue') },
    { path: '/login', component: () => import('@/pages/Login.vue') },
    { path: '/training', component: () => import('@/pages/Training.vue') },
    { path: '/training/:id', component: () => import('@/pages/TrainingModule.vue') },
    { path: '/buckets', component: () => import('@/pages/Buckets.vue') },
    { path: '/instances', component: () => import('@/pages/Instances.vue') },
    { path: '/audit', component: () => import('@/pages/Audit.vue') },
    { path: '/settings', component: () => import('@/pages/Settings.vue') },
  ]
})

router.beforeEach((to, from, next) => {
  const auth = useAuthStore()
  if (to.path !== '/login' && !auth.isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})

export default router
```

Agent API Client (Composable):
```typescript
// composables/useAgent.ts
import axios from 'axios'

const AGENT_URL = 'http://localhost:8737'

export const agentAPI = axios.create({
  baseURL: AGENT_URL,
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Request interceptor: Add auth token
agentAPI.interceptors.request.use(config => {
  const token = localStorage.getItem('ark_token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Response interceptor: Handle errors
agentAPI.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      // Token expired, redirect to login
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export function useAgent() {
  return {
    get: agentAPI.get,
    post: agentAPI.post,
    put: agentAPI.put,
    delete: agentAPI.delete
  }
}
```

Build & Deployment:
  Development:
    npm run dev  # Vite dev server on http://localhost:5173
    # Proxies API requests to localhost:8737 (agent)

  Production Build:
    npm run build  # Creates dist/ with optimized bundles
    # Output: HTML, CSS, JS (minified, tree-shaken, code-split)

  Deployment Options:
    1. Static hosting (S3 + CloudFront, Netlify, Vercel)
    2. Institutional web server (Apache, Nginx)
    3. Bundled with backend (single domain)

  URL Structure:
    Development: http://localhost:5173
    Production: https://ark.{INSTITUTION}.edu
```

**Action Items:**
- [ ] Set up Vue 3 + Vite project
- [ ] Install and configure Cloudscape
- [ ] Create base layout and navigation
- [ ] Implement authentication flow
- [ ] Build agent API client
- [ ] Set up Pinia stores
- [ ] Create routing structure
- [ ] Implement dark mode support
- [ ] Set up TypeScript strict mode
- [ ] Configure ESLint + Prettier

---

### Gap 1.7: Web Security

**New Gap (Web-Specific):**
- XSS prevention?
- CSRF protection?
- Content Security Policy?
- Secure token storage?
- Clickjacking protection?

**Recommendation:**
```
Web Security Measures:

XSS Prevention:
  - Vue 3's template syntax escapes HTML by default
  - Use v-html sparingly, only with sanitized content:
      import DOMPurify from 'dompurify'
      const clean = DOMPurify.sanitize(userInput)
  - Content Security Policy (CSP) headers:
      Content-Security-Policy:
        default-src 'self';
        script-src 'self';
        style-src 'self' 'unsafe-inline';
        img-src 'self' data: https:;
        font-src 'self';
        connect-src 'self' http://localhost:8737;
        frame-ancestors 'none';

CSRF Protection:
  - Custom header requirement:
      X-Ark-Request: true
  - Agent validates header presence
  - SameSite cookie attribute:
      Set-Cookie: ark_session=...; SameSite=Strict; Secure; HttpOnly

Token Storage:
  - Access token: localStorage (short-lived, 1 hour)
  - Refresh token: HttpOnly cookie (long-lived, 30 days)
  - Never store in sessionStorage or regular cookies
  - Encrypt tokens at rest (Web Crypto API)

Authentication Flow:
  1. User clicks "Login" in web app
  2. Web redirects to: https://ark.{INSTITUTION}.edu/auth/login
  3. Backend redirects to institutional IdP (SSO)
  4. User authenticates with IdP
  5. IdP redirects back to backend with SAML assertion
  6. Backend validates, creates session
  7. Backend sets tokens:
      - Access token → response body (web stores in localStorage)
      - Refresh token → HttpOnly cookie
  8. Web redirects to dashboard

Clickjacking Protection:
  - X-Frame-Options: DENY
  - frame-ancestors 'none' (in CSP)

HTTPS Enforcement:
  - All requests use HTTPS (except localhost dev)
  - Strict-Transport-Security: max-age=31536000; includeSubDomains

Subresource Integrity (SRI):
  - For CDN-loaded resources (if any)
  - Vite generates integrity hashes

Secure Headers (Backend serves web app):
  X-Content-Type-Options: nosniff
  X-XSS-Protection: 1; mode=block
  Referrer-Policy: strict-origin-when-cross-origin
  Permissions-Policy: geolocation=(), microphone=(), camera=()

Rate Limiting:
  - Agent API: 100 requests/minute per user
  - Login attempts: 5 failures → 15 minute lockout
  - Web app: Client-side throttling (debounce inputs)

Input Validation:
  - Client-side: Vue form validation (immediate feedback)
  - Agent-side: Validate all inputs (never trust client)
  - Backend-side: Final validation (defense in depth)
  - Sanitize before display (DOMPurify)

Dependency Security:
  - npm audit (run weekly)
  - Dependabot (automated PRs for vulnerabilities)
  - Snyk integration (continuous monitoring)
  - Lock file (package-lock.json committed)
```

**Action Items:**
- [ ] Configure CSP headers
- [ ] Implement CSRF protection
- [ ] Set up secure token storage
- [ ] Add input sanitization (DOMPurify)
- [ ] Enable all security headers
- [ ] Set up npm audit automation
- [ ] Penetration test web app
- [ ] Document security model

---

### Gap 1.8: Cross-Interface Sync & Consistency

**New Gap (Critical for Dual Interface):**
- How to ensure CLI and web show same state?
- Real-time updates when training completed in CLI?
- What if user performs operation in both simultaneously?
- How to handle race conditions?

**Recommendation:**
```
Sync Architecture:

State Sources:
  1. Backend (source of truth): DynamoDB
  2. Agent cache (fast reads): SQLite
  3. CLI display (transient): Terminal buffer
  4. Web display (transient): Vue reactive state

Sync Mechanisms:

CLI → Backend:
  - Immediate writes for critical operations
  - Batched writes for non-critical (every 30 seconds)
  - Optimistic updates (show success, sync in background)

Web → Backend:
  - GraphQL subscriptions (real-time updates)
  - WebSocket connection maintained
  - Automatic reconnection on disconnect
  - Offline queue (IndexedDB) syncs when reconnected

Backend → CLI:
  - Poll every 5 minutes: GET /api/training/progress
  - On-demand: User runs ark sync
  - Before critical operations (bucket create, etc.)

Backend → Web:
  - GraphQL subscription:
      subscription {
        userProgressUpdated(userId: "user123") {
          moduleId
          status
          score
        }
      }
  - Push notifications (Web Push API)
  - Polling fallback (if WebSocket unavailable)

Example: User completes Module 3 in CLI
  1. CLI: User answers final quiz question
  2. CLI → Agent: POST /api/training/complete
  3. Agent → Backend: POST /api/training/complete (with answers)
  4. Backend: Validates answers, updates DynamoDB
  5. Backend → Web: GraphQL subscription notification
  6. Web: Receives update via WebSocket
  7. Web: Pinia store updates: modules[2].status = "completed"
  8. Web: UI reactively updates (green checkmark, unlock bucket creation)
  9. Time elapsed: <2 seconds

Conflict Resolution:
  Scenario: User clicks "Start Module 3" in web, while CLI is completing it

  1. Web → Agent: POST /api/training/start (moduleId: 3)
  2. Agent → Backend: POST /api/training/start
  3. Backend checks: Module 3 status = "completed" (from CLI)
  4. Backend → Agent: { error: "Module already completed" }
  5. Agent → Web: Error response
  6. Web shows notification: "Module 3 already completed!"
  7. Web fetches latest: GET /api/training/progress
  8. Web updates UI to reflect completed status

  Rule: Backend is always right (last-write-wins with timestamp check)

Race Condition Handling:
  Use distributed locks (Redis) for critical operations:

  Example: Both CLI and web try to create same bucket
  1. CLI → Agent: POST /api/aws/bucket/create?name=test
  2. Web → Agent: POST /api/aws/bucket/create?name=test (0.1s later)
  3. Agent acquires lock: bucket:create:test
  4. Agent processes CLI request first (FIFO queue)
  5. CLI request succeeds
  6. Agent releases lock
  7. Agent processes web request
  8. Web request fails: "Bucket already exists"
  9. Web shows notification with option to view existing bucket

Real-Time Indicators:
  Web UI shows live status:
  - "Syncing..." (during background sync)
  - "Up to date" (green checkmark)
  - "Offline" (red badge)
  - "Last synced: 2 minutes ago"

  CLI shows on-demand:
  - ark status → Shows sync status
  - Warning if >1 hour since last sync

Consistency Testing:
  Playwright E2E test:
  ```typescript
  test('CLI and web sync training progress', async ({ page }) => {
    // Complete training in CLI
    await exec('ark learn complete data-classification')

    // Open web app
    await page.goto('http://localhost:5173/training')

    // Expect web to show completion within 5 seconds
    await expect(page.locator('[data-module="data-classification"]'))
      .toHaveAttribute('data-status', 'completed', { timeout: 5000 })

    // Verify certificate downloadable in web
    await page.click('[data-testid="download-certificate"]')
    await expect(page.locator('[data-testid="certificate-pdf"]')).toBeVisible()
  })
  ```
```

**Action Items:**
- [ ] Implement GraphQL subscriptions (backend)
- [ ] Add WebSocket support (web client)
- [ ] Build sync status indicators (CLI + web)
- [ ] Implement conflict resolution logic
- [ ] Add distributed locking (Redis or DynamoDB)
- [ ] Create sync monitoring dashboard
- [ ] Write comprehensive E2E sync tests
- [ ] Document sync behavior for users

---

### Gap 1.9: End-to-End Testing Strategy

**New Gap (Dual Interface Testing):**
- How to test both CLI and web together?
- How to test CLI commands from Playwright?
- How to mock agent responses?
- How to test offline scenarios?

**Recommendation:**
```
Testing Strategy:

Test Pyramid:

Unit Tests (70%):
  Backend (Go):
    - Framework: Go testing package
    - Coverage target: >80%
    - Run on: Every commit (GitHub Actions)
    - Example:
        func TestCreateBucket(t *testing.T) {
          // Mock DynamoDB, S3
          // Test bucket creation logic
        }

  Web (Vue):
    - Framework: Vitest
    - Coverage target: >75%
    - Test: Components, composables, stores
    - Example:
        import { mount } from '@vue/test-utils'
        import BucketList from '@/components/BucketList.vue'

        test('displays buckets', () => {
          const wrapper = mount(BucketList, {
            props: { buckets: mockBuckets }
          })
          expect(wrapper.findAll('.bucket-item')).toHaveLength(3)
        })

Integration Tests (20%):
  Backend API:
    - Test API endpoints with real database (test env)
    - Test agent ↔ backend communication
    - Test SSO integration (with mock IdP)

  Agent:
    - Test agent ↔ AWS SDK (with LocalStack)
    - Test agent ↔ backend API
    - Test caching logic

E2E Tests (10% but most critical):
  Framework: Playwright
  Test both CLI and web in same scenarios

  Test Structure:
```typescript
// tests/e2e/training-flow.spec.ts
import { test, expect } from '@playwright/test'
import { exec } from 'child_process'
import util from 'util'

const execAsync = util.promisify(exec)

test.describe('Training Flow - CLI and Web', () => {
  test('complete training in CLI, verify in web', async ({ page }) => {
    // Start training via CLI
    const { stdout } = await execAsync('ark learn start data-classification')
    expect(stdout).toContain('Module started')

    // Open web app
    await page.goto('http://localhost:5173/training')

    // Check web shows in-progress
    await expect(page.locator('[data-module="data-classification"]'))
      .toHaveAttribute('data-status', 'in_progress')

    // Complete in CLI
    await execAsync('ark learn complete data-classification --answers answers.json')

    // Web should auto-update (within 5 seconds)
    await expect(page.locator('[data-module="data-classification"]'))
      .toHaveAttribute('data-status', 'completed', { timeout: 5000 })

    // Verify certificate available in web
    await page.click('[data-testid="download-certificate"]')
    const download = await page.waitForEvent('download')
    expect(download.suggestedFilename()).toMatch(/certificate.*\.pdf/)
  })

  test('training gate blocks operation in both interfaces', async ({ page }) => {
    // Reset training state
    await execAsync('ark admin reset-training test-user@example.com')

    // Try to create bucket via CLI - should block
    const { stdout: cliOutput } = await execAsync(
      'ark bucket create --name test --classification internal',
      { encoding: 'utf8' }
    ).catch(e => e)
    expect(cliOutput).toContain('Training Required')

    // Try via web - should also block
    await page.goto('http://localhost:5173/buckets/create')
    await page.fill('[name="bucketName"]', 'test')
    await page.selectOption('[name="classification"]', 'internal')
    await page.click('button:has-text("Create")')

    // Should show training gate modal
    await expect(page.locator('[data-testid="training-gate-modal"]')).toBeVisible()
    await expect(page.locator('text=Training Required')).toBeVisible()
  })

  test('simultaneous operations handled gracefully', async ({ page }) => {
    // User attempts same operation in both CLI and web
    const cliPromise = execAsync('ark bucket create --name duplicate-test')

    await page.goto('http://localhost:5173/buckets/create')
    await page.fill('[name="bucketName"]', 'duplicate-test')
    await page.click('button:has-text("Create")')

    // One should succeed, one should fail with helpful error
    const results = await Promise.allSettled([
      cliPromise,
      page.waitForSelector('[data-testid="success-notification"]', { timeout: 5000 })
        .catch(() => page.waitForSelector('[data-testid="error-notification"]'))
    ])

    // Exactly one should succeed
    const successes = results.filter(r => r.status === 'fulfilled')
    expect(successes).toHaveLength(1)
  })
})

// tests/e2e/offline-mode.spec.ts
test('offline training completion', async ({ page, context }) => {
  // Go offline
  await context.setOffline(true)

  // CLI should work with cached content
  const { stdout } = await execAsync('ark learn start aws-basics')
  expect(stdout).toContain('Offline Mode')
  expect(stdout).toContain('Using cached content')

  // Web should show offline banner
  await page.goto('http://localhost:5173')
  await expect(page.locator('[data-testid="offline-banner"]')).toBeVisible()

  // Can still view cached training
  await page.goto('http://localhost:5173/training/aws-basics')
  await expect(page.locator('article')).toBeVisible() // Content loads

  // Go online
  await context.setOffline(false)

  // Progress should sync within 10 seconds
  await page.waitForTimeout(10000)
  await expect(page.locator('[data-testid="offline-banner"]')).not.toBeVisible()
  await expect(page.locator('[data-testid="sync-status"]')).toHaveText('Up to date')
})
```

Test Environments:
  Development: Runs locally with LocalStack (mock AWS)
  CI/CD: GitHub Actions with Docker Compose
  Staging: Real AWS account (test/dev environment)
  Production: Smoke tests only (non-destructive)

Mocking Strategy:
  LocalStack: Mock AWS services
  Mock IdP: For SSO testing
  Wiremock: Mock backend API (for isolated web tests)

  Example Docker Compose:
```yaml
version: '3'
services:
  localstack:
    image: localstack/localstack
    ports:
      - "4566:4566"
    environment:
      - SERVICES=s3,ec2,iam,dynamodb

  mock-idp:
    image: kristophjunge/test-saml-idp
    ports:
      - "8080:8080"

  agent:
    build: ../agent
    ports:
      - "8737:8737"
    environment:
      - ARK_BACKEND_URL=http://backend:8080
      - AWS_ENDPOINT=http://localstack:4566

  backend:
    build: ../backend
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://test:test@db:5432/ark_test

  web:
    build: ../web
    ports:
      - "5173:5173"
    environment:
      - VITE_AGENT_URL=http://localhost:8737
```

CI/CD Pipeline (GitHub Actions):
```yaml
name: E2E Tests

on: [push, pull_request]

jobs:
  e2e:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Start services
        run: docker-compose up -d

      - name: Install CLI
        run: |
          cd cli
          go build -o ark
          sudo mv ark /usr/local/bin/

      - name: Install web dependencies
        run: |
          cd web
          npm install
          npx playwright install --with-deps

      - name: Wait for services
        run: ./scripts/wait-for-services.sh

      - name: Run E2E tests
        run: |
          cd web
          npm run test:e2e

      - name: Upload test results
        if: always()
        uses: actions/upload-artifact@v3
        with:
          name: playwright-report
          path: web/playwright-report/
```

Performance Testing:
  Tool: k6 (load testing)
  Scenarios:
    - 100 concurrent users completing training
    - 1000 users creating buckets simultaneously
    - 10,000 API requests/second

  Targets:
    - P95 latency: <500ms
    - Error rate: <0.1%
    - Agent uptime: 99.9%

Visual Regression Testing:
  Tool: Percy or Chromatic
  Test: Web UI screenshots across:
    - Different browsers (Chrome, Firefox, Safari)
    - Different viewport sizes (mobile, tablet, desktop)
    - Dark/light modes
    - Different states (loading, error, success)
```

**Action Items:**
- [ ] Set up Playwright for E2E testing
- [ ] Create test fixtures and helpers
- [ ] Implement CLI command mocking
- [ ] Set up LocalStack for AWS mocking
- [ ] Create Docker Compose for test environment
- [ ] Configure GitHub Actions CI/CD
- [ ] Write 50+ E2E test scenarios
- [ ] Set up visual regression testing
- [ ] Implement performance testing (k6)
- [ ] Create test data generators

---

## 2. Security & Compliance

[Continue with existing security gaps, updated for dual interface...]

### Gap 2.1: Audit Trail Completeness

**Missing:**
- What exact data is logged for compliance?
- How long are logs retained?
- Who has access to audit logs?
- How do you prove training completion to auditors?
- Do web and CLI operations log identically?

**Recommendation:**
```
Comprehensive Audit Logging:

What to Log (all operations, both CLI and web):
  ✓ Training events:
      - Module start/complete times (UTC)
      - Quiz attempts and scores
      - Time spent per section
      - Failed checkpoint attempts
      - Certificate generation
      - Interface used (CLI or web)

  ✓ AWS operations:
      - Commands executed (with parameters)
      - Resources created (ARNs)
      - Resource modifications
      - Resource deletions
      - Failed operations (why did they fail?)
      - Interface used (CLI or web)

  ✓ Security events:
      - Login/logout
      - Failed authentication attempts
      - MFA events
      - Permission changes
      - Policy violations
      - Training bypass attempts

  ✓ Admin actions:
      - User management (create, disable, reset)
      - Policy updates
      - Training resets
      - Overrides (with justification)

  ✓ System events:
      - Agent start/stop
      - Version updates
      - Configuration changes
      - Errors and exceptions

Event Schema:
```json
{
  "event_id": "uuid",
  "timestamp": "2025-12-09T10:30:00.000Z",
  "user_id": "user@institution.edu",
  "aws_account_id": "123456789012",
  "event_type": "s3:CreateBucket",
  "interface": "cli",  // or "web"
  "result": "success",  // or "failure"
  "resource": {
    "type": "s3:bucket",
    "name": "my-bucket",
    "arn": "arn:aws:s3:::my-bucket",
    "classification": "internal"
  },
  "metadata": {
    "command": "ark bucket create --name my-bucket --classification internal",
    "agent_version": "1.2.0",
    "cli_version": "1.2.0",
    "ip_address": "192.168.1.100",
    "user_agent": "Ark CLI/1.2.0 (darwin; arm64)",
    "session_id": "uuid",
    "training_completed": true,
    "policy_checks": ["encryption", "classification", "training"]
  },
  "context": {
    "department": "biology",
    "pi": "jane.doe@institution.edu",
    "grant": "NIH R01 123456"
  }
}
```

Storage Architecture:
  Tier 1: Real-time (CloudWatch Logs)
    - All events streamed immediately
    - Retention: 90 days
    - Purpose: Monitoring, alerting, debugging
    - Searchable via CloudWatch Insights

  Tier 2: Compliance (DynamoDB)
    - All events indexed
    - Retention: 7 years
    - Purpose: Queries, reporting, audits
    - Partition key: user_id
    - Sort key: timestamp
    - GSI: event_type, resource_type

  Tier 3: Archive (S3)
    - Daily batches exported
    - Retention: Indefinite (with S3 Glacier)
    - Purpose: Long-term compliance, disaster recovery
    - Immutable (Object Lock enabled)
    - Encrypted (SSE-KMS)

Access Control:
  - Logs are write-only for users
  - Users can read own logs only
  - Admins have read-only access (no modification)
  - Audit officers have full access
  - All access to logs is itself logged (audit the auditors)

Compliance Reports:
  Automated Monthly:
    - Training completion rates
    - Security violations
    - Cost anomalies
    - Resource inventory
    - Policy compliance score
    - Export: PDF, CSV, JSON

  On-Demand:
    - Per-user compliance report
    - Per-department dashboard
    - Certification evidence package (for HIPAA, FISMA audits)
    - Time-range filtered reports

  Web Interface:
    - Admin can view/search logs (Cloudscape Table with filters)
    - Export filtered results
    - Visualize trends (charts)
    - Drill-down to individual events

Example Queries:
  "Show all S3 bucket creations in last 30 days"
  "List users who haven't completed training"
  "Find all operations on P4 data"
  "Show failed authentication attempts for user X"
  "Export audit trail for grant ABC for fiscal year 2025"
```

**Action Items:**
- [ ] Define complete audit event schema
- [ ] Implement CloudWatch Logs streaming
- [ ] Set up DynamoDB audit table (with indexes)
- [ ] Configure S3 archival (with Object Lock)
- [ ] Create compliance report generator
- [ ] Design audit log access controls
- [ ] Build web UI for log viewing (Cloudscape)
- [ ] Implement log export functionality
- [ ] Test log retention and retrieval
- [ ] Document audit trail for compliance officers

---

[Due to length limitations, I'll provide a summary of remaining updates needed for the document. Would you like me to continue with the full updated version?]

**Remaining sections to update:**
- Gap 2.2: SIEM Integration (add web dashboard for security events)
- Gap 2.3: Data Sovereignty & Privacy (GDPR for both interfaces)
- Gap 3.1: Support Model (add web-specific support)
- Gap 3.2: Cost Model (add web infrastructure costs)
- Gap 3.3: Accessibility (major section on WCAG 2.1 AA for web)
- Gap 4.1: Identity Integration (SSO for both CLI and web)
- Gap 4.2: AWS Organizations (unchanged)
- Gap 5.1-5.3: Content gaps (add web training delivery)
- Gap 6.1-6.2: Metrics and feedback (add web analytics)
- Risk Register (add web-specific risks)
- Open Questions (add decisions about web deployment)

This is a large document. Would you like me to:
1. Continue with the full comprehensive update (may need multiple responses)
2. Provide a summary of key changes
3. Update specific sections you're most interested in
4. Create this as multiple files for manageability?
