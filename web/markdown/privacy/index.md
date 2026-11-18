# Privacy

GitVerse is built on a fundamental principle: your developer data belongs exclusively to you. No databases, no persistent storage, and zero tracking.

## Our Privacy Commitment

GitVerse is built on a fundamental principle: your developer data belongs exclusively to you. We believe privacy should be absolute and non-negotiable. This policy outlines how GitVerse handles authentication, public metrics, and temporary session state across both our command line interface and web application.

## Command Line Interface

The GitVerse CLI operates entirely on your local machine. When retrieving public achievement information, it connects directly to GitHub's public endpoints. If authenticated using GitHub CLI or a local personal access token, credentials remain strictly within your local environment. The CLI never sends credentials, query history, or telemetry logs to GitVerse or any remote servers.

## Web Application & OAuth

When you check a username on the web interface, GitVerse queries public profile metrics through a read-only server token. If you sign in via GitHub to include private repository achievements, authorization is strictly ephemeral. OAuth tokens are held in volatile memory only for the duration of a single check and revoked immediately after your report is generated. GitVerse maintains no database, stores no user profiles, and logs no usernames, tokens, or repository names.

## Analytics & Cookies

GitVerse employs strictly privacy-focused analytics with Google Analytics in consent-gated mode. No analytics cookies or trackers are loaded unless you explicitly grant consent. No IP addresses, usernames, or private repository data are ever transmitted to analytics providers.

## Third-Party Infrastructure

GitVerse is hosted on Vercel's global edge network under secure TLS 1.3 encryption. Badge artwork is loaded directly from GitHub's official content delivery network. All network transactions adhere to industry-standard transport security.

## Contact & Inquiries

If you have any questions, issues, or privacy inquiries, please reach out directly to Vicky Patel at vicky.patel.work@gmail.com.


---

[GitVerse](https://mygitverse.vercel.app/) · [llms.txt](https://mygitverse.vercel.app/llms.txt) · [About](https://mygitverse.vercel.app/about/) · [Privacy](https://mygitverse.vercel.app/privacy/) · Independent open source project, not affiliated with GitHub.
