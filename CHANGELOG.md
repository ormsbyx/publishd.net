# Publishd - Development Changelog

All notable changes to the Publishd project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- ✅ Go project structure with standard layout
- ✅ Git repository initialization with proper .gitignore
- ✅ Gin web framework integration
- ✅ Basic HTTP server foundation

---

## Project Initialization - 2025-09-09

### Added
- ✅ Project planning document with technical architecture
- ✅ Development phases breakdown (8-week timeline)
- ✅ To-do list with organized milestones
- ✅ Technology stack finalization:
  - Backend: Go (new learning opportunity)
  - Database: PostgreSQL
  - Frontend: Vue 3
  - Hosting: Render (free tier)
- ✅ Multi-tenancy strategy:
  - Subdomains: user.publishd.net
  - Custom domains: user-domain.com
  - Single database with tenant isolation

### Decisions Made
- **Framework Choice:** Go (over Node.js) for learning and performance
- **Hosting Choice:** Render (over A2Hosting) for modern deployment and easy SSL
- **Domain Strategy:** Support both subdomains and custom domains from day one
- **Database Schema:** Single PostgreSQL instance with proper tenant isolation

### Next Steps
- [ ] Initialize Go project structure
- [ ] Set up Git repository
- [ ] Choose Go web framework (Gin recommended)
- [ ] Create basic HTTP server with tenant detection
- [ ] Set up Render deployment configuration

---

## Template for Future Entries

### [Version] - YYYY-MM-DD

#### Added
- New features

#### Changed
- Changes in existing functionality

#### Deprecated
- Soon-to-be removed features

#### Removed
- Now removed features

#### Fixed
- Bug fixes

#### Security
- Vulnerability fixes

---

**Changelog Guidelines:**
- Keep entries concise but descriptive
- Group changes by type (Added, Changed, Fixed, etc.)
- Include relevant commit hashes when applicable
- Update this file with every significant change
- Use present tense for descriptions


## Scope Refinement - 2025-09-09

### Changed
- ✅ Refined project scope from multi-tenant platform to single-author reading platform
- ✅ Shifted focus from Substack-like newsletter platform to Kindle-like reading experience
- ✅ Simplified initial scope: single writer (expandable to pen names later)
- ✅ Content focus: short stories, essays, and articles (not novels initially)
- ✅ Payment model: monthly subscriptions + pay-per-story options
- ✅ Removed email/newsletter functionality from initial scope
- ✅ Removed social features, discovery, and cross-promotion
- ✅ Emphasized mobile-first, clean reading experience

### Decisions Made
- **Content Strategy:** Focus on premium reading experience over platform features
- **Initial User:** Single author (you) to start, pen names can be added later
- **Reader Experience:** Kindle-like interface with mobile-first responsive design
- **Monetization:** Dual model - monthly subscriptions and individual story purchases
- **Technical Simplification:** Removed multi-tenancy complexity for initial version