# Publishd - Project Planning Document

## 🎯 Project Overview

**Name:** Publishd  
**Domain:** publishd.net  
**Concept:** Premium reading platform for short stories, essays, and articles
**Initial Scope:** Single writer (you) with clean, Kindle-like reading experience

## 🏗️ Technical Architecture

### Core Technology Stack
- **Backend:** Go (learning opportunity)
- **Database:** PostgreSQL
- **Frontend:** Vue 3
- **Hosting:** Render (free tier with easy SSL)
- **Domain:** publishd.net

### Simplified Architecture
- **Single Author Platform:** You as the primary writer
- **Custom Domain:** publishd.net with potential for custom domain later
- **Database:** Single PostgreSQL instance with simple schema
- **File Storage:** Basic file organization for stories and images

### Core Data Models
```sql
stories (id, title, content, excerpt, price, published_at, created_at)
users (id, email, password_hash, subscription_active, created_at)
purchases (user_id, story_id, purchased_at)
subscriptions (user_id, stripe_subscription_id, active, expires_at)
```

## 🔧 Development Phases

### Phase 1: Foundation (Week 1-2)
**Goal:** Single-author reading platform with basic functionality
- Set up Go project with simple routing
- Create story CRUD operations  
- Basic reading interface
- Deploy to Render

### Phase 2: Payment & Access (Week 3-4)
**Goal:** Subscription and pay-per-story functionality
- Stripe integration for payments
- User authentication and access control
- Monthly subscription logic
- Individual story purchases

### Phase 3: Reading Experience (Week 5-6)
**Goal:** Polish the core reading experience
- Mobile-first responsive design
- Kindle-like typography and layout
- Reading progress tracking
- Favorites and bookmarks

### Phase 4: Content Management (Week 7-8)
**Goal:** Easy content creation and organization
- Markdown import functionality
- Rich text editor for writing/editing
- Story categorization and tagging
- Draft/published workflow

## 🎨 User Experience Flow

### For Content Creators (You):
1. **Write** stories in markdown or rich text editor
2. **Set pricing** (free preview, subscription-only, or pay-per-story)
3. **Publish** with clean, mobile-optimized layout
4. **Track** reader engagement and revenue

### For Readers:
1. **Browse** story collection on publishd.net
2. **Read previews** for free
3. **Subscribe monthly** or **buy individual stories**
4. **Enjoy** distraction-free, Kindle-like reading experience

## 💰 Monetization Strategy

### Payment Model:
- **Free Preview** - First paragraph of each story
- **Monthly Subscription** - Access to everything ($5-10/month)
- **Pay-per-Story** - Individual story purchases ($1-3 each)
- **Subscribers get discount** on individual purchases

### Content Strategy:
- **Short Stories** - Main focus, perfect for mobile reading
- **Essays** - Personal and analytical pieces
- **Articles** - Informational content
- **No Novels** - Keep content digestible for initial scope

## 🔐 Security Considerations
- JWT-based authentication
- SQL injection prevention
- Secure payment processing via Stripe
- HTTPS everywhere (automatic via Render)
- Basic rate limiting

## 📊 Success Metrics
- **MVP Success:** 10 published stories with functional payment system
- **Phase 1 Success:** Basic reading platform deployed and working
- **Growth Success:** 50+ paying readers within 3 months

## 🚨 Potential Challenges
1. **Go Learning Curve:** First time using Go
2. **Payment Integration:** Stripe setup and testing
3. **Mobile Reading Experience:** Getting typography and layout right
4. **Content Migration:** Importing existing stories efficiently

## 🛠️ Development Tools & Workflow
- **Version Control:** Git with feature branches
- **Code Editor:** VS Code with Go extensions
- **Database Tool:** TablePlus or pgAdmin
- **Payment Testing:** Stripe test environment
- **Deployment:** Automatic via Git push to Render

## 📈 Future Enhancements (Post-MVP)
- **Multi-pen-name support:** Add author field for different writing identities
- **Reading analytics:** Track reading time, completion rates
- **Social features:** Reader comments and ratings
- **Email notifications:** New story alerts for subscribers
- **Advanced search:** Filter by genre, length, publication date
- **Export functionality:** PDF/EPUB downloads for subscribers

---
**Last Updated:** Scope refinement session - 2025-09-09  
**Next Review:** After Phase 1 completion