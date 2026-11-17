# GitHub API Rate Limiting & Backoff

GitVerse respects GitHub API quotas:
* Anonymous: 60 requests / hour / IP
* Authenticated: 5,000 requests / hour / user

When a 403 or 429 response is encountered, GitVerse halts further polling and returns official cached floor milestones.
