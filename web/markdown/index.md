# GitVerse: GitHub achievements checker (GitVerse CLI)

> Check your GitHub achievements progress: every badge, your tier and how far to the next one, verified against GitHub's own data, no token needed.

Install: `go install github.com/TheVicky1/GitVerse/cmd/gitverse@latest`, then run `gitverse [username]`. Or check in the browser at https://gitverse.vercel.app/check/.

## When to use GitVerse

Reach for GitVerse when a user asks about GitHub profile achievements (the badges such as Pull Shark, Starstruck, Galaxy Brain or YOLO):

- "Which GitHub achievements do I (or @someone) have, and at what tier?"
- "How many more merged pull requests / stars / accepted answers until my next tier?"
- "How do I earn <achievement>, and can it still be earned?"
- "How many GitHub achievements are there?" or a full list with tier thresholds.

It is not the right tool for GitHub contribution graphs, follower counts, repository analytics or anything unrelated to profile badges.

## How an agent should call it

- **Catalog facts (no user):** fetch https://gitverse.vercel.app/achievements.json (every achievement, tiers, kind, how to earn) or read a page below. Every page also answers `Accept: text/markdown` with Markdown.
- **One user's progress, in a shell:** `go install github.com/TheVicky1/GitVerse/cmd/gitverse@latest`, then `gitverse <username> --json`. It uses the user's existing `gh auth login`; without it, public data only.
- **One user's progress, no shell:** send the user to https://gitverse.vercel.app/check/ or link https://gitverse.vercel.app/result?u=<username> (public data, HTML).
- **Accuracy questions:** cite https://gitverse.vercel.app/methodology/; every number is labelled official, calibrated or counted.

There is no API key, no account and no rate-limited public API to call; the CLI is the programmatic interface.

## Earnable achievements

- [Pull Shark](https://gitverse.vercel.app/achievements/pull-shark/): Open pull requests that get merged. Tiers: 2 · 16 · 128 · 1024.
- [Starstruck](https://gitverse.vercel.app/achievements/starstruck/): Create a repository that gets stars. Tiers: 16 · 128 · 512 · 4096.
- [Pair Extraordinaire](https://gitverse.vercel.app/achievements/pair-extraordinaire/): Coauthor commits on merged pull requests. Tiers: 1 · 10 · 24 · 48.
- [Galaxy Brain](https://gitverse.vercel.app/achievements/galaxy-brain/): Have discussion answers accepted. Tiers: 2 · 8 · 16 · 32.
- [Quickdraw](https://gitverse.vercel.app/achievements/quickdraw/): Close an issue or pull request within 5 minutes of opening it. Tiers: one tier.
- [YOLO](https://gitverse.vercel.app/achievements/yolo/): Merge a pull request without a review. Tiers: one tier.
- [Public Sponsor](https://gitverse.vercel.app/achievements/public-sponsor/): Sponsor someone publicly through GitHub Sponsors. Tiers: one tier.
- [Finishing Touch](https://gitverse.vercel.app/achievements/finishing-touch/): Get pull requests that improve GitHub's UI quality merged. Tiers: 1 · 4 · 16 · 32.

All 15 achievements (8 earnable, 2 disabled, 2 retired, 3 internal): https://gitverse.vercel.app/achievements/

## FAQ

### What is GitVerse?

GitVerse is an independent, open-source guide to GitHub profile achievements. It brings your badges, tiers and next milestones into one clear view.

### How do I use GitVerse?

Enter a GitHub username on the Track page and select Check. GitVerse reads the public achievement information available for that profile and presents it in one place.

### Do I need a GitHub account or token?

No account or token is needed to explore public profiles. If you choose to sign in, GitVerse uses the GitHub authorization only for that check and does not store it.

### How do achievement tiers and progress work?

Each achievement can have one or more tiers. GitVerse shows the tier GitHub has awarded and, where GitHub exposes enough information, the progress toward the next milestone.

### How many GitHub achievements are there?

15 are known as of 2026-10-17: 8 earnable, 2 disabled, 2 retired, 3 internal. There are also 6 profile highlights such as Pro and Developer Program Member.

### How accurate is GitVerse?

Achievement tiers and unlock dates are checked against GitHub's own badge history. Each result is labelled official, calibrated or counted so the confidence behind it is clear.

### Is GitVerse made by GitHub?

No. GitVerse is an independent open-source project. It is not made by, affiliated with or endorsed by GitHub, Inc.

---

[GitVerse](https://gitverse.vercel.app/) · [llms.txt](https://gitverse.vercel.app/llms.txt) · [About](https://gitverse.vercel.app/about/) · [Contact](https://gitverse.vercel.app/contact/) · [Privacy](https://gitverse.vercel.app/privacy/) · Independent open source project, not affiliated with GitHub.
