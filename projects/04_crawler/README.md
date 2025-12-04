To be able to design crawlers/scrapers “from first principles” (and not just copy snippets), you want four pillars in your head:

1. **How the web actually delivers content (HTTP + browsers)**
2. **How pages are structured (HTML/CSS/DOM) + how to select data**
3. **How crawling works at scale (frontier, dedupe, scheduling, politeness)**
4. **How to turn scraped pages into reliable data (parsing, validation, storage, QA)**

Below is a comprehensive concept map + a roadmap you can follow.

---

## Concept map: what you need to understand (and why)

### A) Web fundamentals (HTTP, not magic)

**Goal:** predict what your code will receive and how servers will react.

* **Request/response lifecycle**

  * Methods: `GET`, `POST` (plus `HEAD` can be useful)
  * Status codes: `200`, `301/302`, `403`, `404`, `429`, `5xx`
  * Redirect behavior and how it affects canonical URLs / dedupe
* **Headers that matter in scraping**

  * `User-Agent`, `Accept`, `Accept-Language`, `Referer`
  * `Cookie` (sessions), `Authorization` (APIs)
  * Caching headers: `ETag`, `If-Modified-Since` (for incremental crawls)
* **Content types**

  * HTML vs JSON vs XML (sitemap, RSS) vs files (PDF)
* **Encodings**

  * UTF-8 (usually), sometimes legacy encodings
  * Gzip/Brotli compression (client usually handles, but good to know)
* **Server-side defenses & “politeness signals”**

  * Rate limits (`429`), WAF behavior, blocks (`403`)
  * Why concurrency + delays must be controlled

**If you nail only one thing here:** understand how to reproduce a browser request (URL + headers + cookies) and interpret the response reliably.

---

### B) HTML / DOM essentials (how pages are “shaped”)

**Goal:** be able to look at page source / DevTools and identify stable extraction points.

* **HTML structure**

  * Tags, attributes (`href`, `src`, `data-*`)
  * Forms, tables, lists, semantic sections
* **DOM vs “View Source”**

  * The DOM may differ after JavaScript runs
* **CSS selectors**

  * `div.class`, `#id`, descendant `A B`, direct child `A > B`
  * Attribute selectors: `a[href*="foo"]`, `img[src^="https"]`
  * Why IDs can be unstable, and `data-*` often is more stable
* **Text extraction pitfalls**

  * Hidden text, whitespace normalization, nested nodes
* **Relative vs absolute URLs**

  * `/path` vs `https://domain/path` and how to resolve properly

**Small superpower:** open DevTools → inspect element → craft a selector that survives small layout changes.

---

### C) JavaScript + dynamic sites (when HTML isn’t “there” yet)

**Goal:** know when Colly alone is enough vs when you need a browser.

* **Two categories of websites**

  1. **Server-rendered HTML** (Colly is perfect)
  2. **Client-rendered (SPA)** where content loads via XHR/fetch → HTML has placeholders
* **Strategies**

  * Prefer: scrape the underlying JSON endpoints (network tab → XHR) rather than rendering
  * If unavoidable: use a headless browser (in Go: `chromedp`), then parse DOM

**Rule of thumb:** if “View Source” doesn’t contain your data, it’s probably rendered or fetched dynamically.

---

### D) Crawling fundamentals (the “frontier” mindset)

**Goal:** control exploration so you don’t loop forever or crawl garbage.

* **Crawler anatomy**

  * **Seeds** → **Frontier (queue)** → **Fetcher** → **Parser/Extractor** → **Storage**
* **Traversal strategies**

  * BFS vs DFS (queue naturally gives BFS-ish behavior)
  * Depth limits
* **URL normalization**

  * Remove fragments `#...`
  * Handle trailing slash, default ports, lowercasing host
  * Decide policy for query params (keep, drop, whitelist)
* **Deduplication**

  * “Visited URLs” set (in-memory / persistent)
  * Content dedupe (hash pages) when URLs differ but content same
* **Scope control**

  * Allowed domains
  * URL allow/deny patterns (regex or string rules)
  * “Page type routing” (list pages vs detail pages)

---

### E) Reliability & production concerns (what makes it not fall apart)

**Goal:** keep it stable on bad networks, weird HTML, and partial failures.

* **Timeouts everywhere**

  * HTTP client timeouts (connect, TLS, response headers, overall)
* **Retries with backoff**

  * Retry only on transient errors (`5xx`, `429`, network issues)
  * Respect `Retry-After` if provided
* **Rate limiting / politeness**

  * Concurrency controls + delays per domain
* **Observability**

  * Logs: request URL, status, elapsed time
  * Metrics: pages fetched, errors by type, queue size
* **Incremental crawling**

  * Resume after crash: persistent queue + visited set
  * Re-crawl policies: only new/changed pages

---

### F) Data extraction as “engineering” (not just hacking selectors)

**Goal:** produce clean datasets.

* **Schema**

  * Decide what fields you need, types, constraints
* **Validation**

  * Required fields present? parseable dates? ranges?
* **Normalization**

  * Clean whitespace, parse numbers/currency, unify units
* **Storage**

  * JSONL for quick wins, SQLite/Postgres for serious projects
* **Reproducibility**

  * Save raw HTML snapshots for debugging extraction regressions

---

### G) Ethics + legality (don’t skip this)

**Goal:** don’t get blocked or cause trouble.

* Respect robots where appropriate (and at least understand it)
* Avoid aggressive rates
* Don’t scrape personal data you don’t have a right to use
* Prefer official APIs if terms demand it

---

## Roadmap: from zero-to-“I can build one solo”

This is ordered so each step unlocks the next. Each phase ends with a deliverable you can actually implement.

### Phase 1 — Web + HTML essentials (1–2 weeks)

**Learn**

* HTTP basics (requests, headers, status codes, redirects)
* HTML structure + CSS selectors + relative/absolute URLs

**Deliverables**

* A tiny Go program using `net/http` that fetches a page and prints:

  * status code, final URL after redirects, key headers
* A “selector notebook”: for 10 pages, write 2–3 selectors each that extract something real (title, breadcrumbs, table rows)

---

### Phase 2 — Scraping single pages robustly (1 week)

**Learn**

* Parsing HTML (e.g., goquery via Colly’s callbacks)
* Clean text extraction + normalization
* Basic error handling

**Deliverable**

* “Single-page scraper” that outputs structured JSON for one page type (e.g., article page → `{title, date, body}`)

---

### Phase 3 — Your first crawler (2 weeks)

**Learn**

* Frontier (queue), dedupe, URL normalization
* Scope rules (allowed domains + URL patterns)
* Page-type routing

**Deliverable**

* “Two page types” crawler:

  * List page discovers detail links
  * Detail page extracts data
  * Writes JSONL
  * Has depth limit + visited set

**Structure pattern to aim for**

* `main.go` wires everything
* `rules.go` contains URL allow/deny + page-type detection
* `extract/*.go` contains extractor functions per page type
* `store/*.go` contains JSONL/DB writer

---

### Phase 4 — Dynamic data & “real websites” (1–2 weeks)

**Learn**

* Detect client-rendered content
* Find JSON endpoints in Network tab
* Use headless browser only when needed

**Deliverable**

* For one SPA-like site:

  * Either scrape its JSON endpoint directly **or**
  * Use a headless browser to get rendered HTML then extract

---

### Phase 5 — Reliability (2 weeks)

**Learn**

* Timeouts, retries with backoff, handling `429`
* Better dedupe (canonical URLs, query policy)
* Resuming: persistent visited set (SQLite or local file)

**Deliverable**

* Crawler can be stopped and resumed without losing progress
* Produces a crawl report: pages visited, errors by type, top slow URLs

---

### Phase 6 — “Production-style” architecture (ongoing)

**Learn**

* Concurrency patterns in Go: worker pools, channels, context cancellation
* Modular design: fetcher/parser/storage separation
* Testing with saved HTML fixtures

**Deliverables**

* Unit tests for your extractors using stored HTML samples
* Config-driven crawler (YAML/JSON): domain, seeds, rate limits, selectors

---

## A practical checklist for *structuring* any crawler

When you start a new target site, answer these in order:

1. **Is my data in “View Source”?** If yes → Colly only. If no → check XHR JSON endpoints or headless browser.
2. **What are my page types?** list/detail/search/pagination.
3. **What are stable selectors or stable JSON fields?** Avoid brittle class names if possible.
4. **What’s my scope?** domains + allow/deny URL patterns + depth.
5. **How do I avoid duplicates?** URL normalization + visited store.
6. **How do I stay polite?** rate limits + parallelism + retries.
7. **What’s my schema and validation?** decide now, not later.
8. **How do I debug?** save raw HTML snapshots for failed pages.