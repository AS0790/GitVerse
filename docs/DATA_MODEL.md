# GitVerse Data Model & Provenance Specification

This document defines the data structures, provenance classification, and JSON output contracts used throughout GitVerse.

---

## 1. Provenance Hierarchy

Every metric reported by GitVerse carries an explicit provenance tag to provide absolute transparency into data certainty:

```text
                     +---------------------------+
                     |  Provenance Classification |
                     +---------------------------+
                                   |
         +-------------------------+-------------------------+
         |                         |                         |
         v                         v                         v
   [ official ]              [ calibrated ]             [ counted ]
Highest Certainty         Milestone Correlated       Telemetry Computed
Extracted directly        API count verified         Live API query
from GitHub's own         against historical         without named
profile HTML documents.   unlock records.            reference milestone.
```

### Classification Rules
1. **`official`**: Data sourced directly from GitHub markup (`data-hovercard-url`, `tier-count`, `datetime`). Zero calculation required.
2. **`calibrated`**: Counted via the GitHub REST/GraphQL API and matched against explicit milestone numbers named in GitHub's badge history records (e.g. Pull Shark, Starstruck).
3. **`counted`**: Derived purely from API telemetry (e.g. Galaxy Brain discussion answers) when GitHub does not expose a named milestone target.

---

## 2. Report Schema (JSON Output)

Running `gitverse --json <username>` outputs a deterministic JSON document matching the following schema:

```json
{
  "user": "TheVicky1",
  "achievements_hidden": false,
  "authenticated": true,
  "catalog_verified": "2026-10-17",
  "achievements": [
    {
      "slug": "pull-shark",
      "name": "Pull Shark",
      "kind": "earnable",
      "status": "earned",
      "tier": 2,
      "unlocked": "2024-03-12",
      "count": 80,
      "target": 128,
      "source": "calibrated",
      "counter": "merged_prs",
      "notes": []
    },
    {
      "slug": "starstruck",
      "name": "Starstruck",
      "kind": "earnable",
      "status": "earned",
      "tier": 1,
      "unlocked": "2023-08-20",
      "count": 42,
      "target": 128,
      "source": "calibrated",
      "counter": "stars",
      "notes": []
    }
  ],
  "highlights": [
    {
      "name": "Pro",
      "earned": true
    },
    {
      "name": "Developer Program Member",
      "earned": false
    }
  ]
}
```

---

## 3. Calibration Safeguards

To prevent negative or contradictory progress indicators, GitVerse enforces **Calibration Floors**:

$$\text{Displayed Count} = \max(\text{Live Count}, \text{Tier Floor})$$

If GitHub's search API lags or fails to index recent private merges, a developer who holds a Silver tier (x3) Pull Shark badge will never be shown a count lower than the 128 PR floor required for that tier.

