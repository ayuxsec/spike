# Spike Roadmap

> **Status:** Planning / Non-binding  
> This roadmap captures current ideas, experiments, and long-term direction.
> Features may change, be delayed, or be dropped entirely.

## 🧠 Philosophy

Spike is built as a **DB-backed, resumable recon engine** first not just a tool wrapper.
The roadmap prioritizes:
- correctness over speed
- composability over monolithic scans
- automation-first workflows for real bug bounty usage

## 🧩 Core Improvements

- [ ] HTML report output
- [ ] Parallel domain scanning
- [ ] ClickHouse backend support
- [ ] Automatic nuclei template updates
- [ ] Better error aggregation & per-step failure visibility
- [ ] Spike Shell REPL-like workflow Support (Interactive DB CLI)

## 📊 UI & Reporting

- [ ] Dashboard UI
- [ ] Advanced reporting formats  
  - JSON
  - HTML
  - API / pipe-friendly mode
- [ ] Per-domain scan summaries
- [ ] Historical scan comparison

## 🔧 Tooling & Pipelines

- [ ] Integration of additional recon & vulnerability tools (ongoing)
- [ ] More scanning pipelines:
  - JavaScript-focused recon
  - Port scanning
  - Cloud asset checks
- [ ] More built-in custom nuclei templates
- [ ] Optional plugin system for community modules


- [ ] **Spike Shell (Interactive DB CLI)**

```console
spike shell "example.db"
```

Planned capabilities:
- View all domains
- Inspect scan status
- Run subcommands (list, filter, describe, purge)
- Query tool outputs (subfinder, httpx, katana, nuclei, etc.)
- Reset or re-scan domains
- Power-user friendly, REPL-like workflow