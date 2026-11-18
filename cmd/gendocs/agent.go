package main

import (
	"fmt"
	"strings"

	"github.com/TheVicky1/GitVerse/internal/catalog"
	"github.com/TheVicky1/GitVerse/internal/webui"
)

// trustPage is one of the pages that say who runs the site, how to reach them and what
// happens to visitors' data. The site renders them as HTML, the Markdown twin as text.
type trustPage struct {
	Slug, Title, Heading, Description string
	Eyebrow                           string
	Lede                              string
	Paras                             []para
}

var trustPages = []trustPage{
	{
		Slug: "about", Title: "About GitVerse: Architecture, Data Trust & Author · GitVerse",
		Heading:     "About GitVerse",
		Eyebrow:     "Architecture & Philosophy",
		Description: "GitVerse is an open-source GitHub achievement intelligence suite and Go CLI by Vicky Patel. Architecture, privacy model, and methodology.",
		Lede:        "A modern, developer-first intelligence suite engineered to provide complete transparency into GitHub profile achievements.",
		Paras: []para{
			{
				Heading: "The Platform",
				Bullets: []bullet{
					{"Zero Guesswork", "Surfaces verified tier thresholds, current standings, and precise distances to your next milestone without estimation."},
					{"Universal Access", "Functions seamlessly as an interactive, real-time web intelligence dashboard and as a zero-dependency CLI binary."},
					{"Live Telemetry", "Synchronized directly against GitHub’s live achievement models, tier badges, and developer profile specifications."},
				},
			},
			{
				Heading: "Why We Built It",
				Bullets: []bullet{
					{"Native Visibility Gaps", "While GitHub highlights earned badges on profiles, it provides developers no native visibility into progress pacing or remaining milestone requirements."},
					{"Omitted API Telemetry", "The official GitHub REST and GraphQL APIs omit achievement metrics, unlock histories, and tier progression entirely."},
					{"Authoritative Reference", "Establishes a single, verified open source standard documenting all 15 earnable, disabled, retired, and internal achievements."},
				},
			},
			{
				Heading: "Architecture & Security",
				Bullets: []bullet{
					{"Ephemeral Execution", "OAuth tokens for private checks are held only in volatile memory during compilation and immediately revoked upon completion."},
					{"Zero Persistence", "GitVerse operates strictly without databases, credential storage, IP tracking, or persistent profile telemetry."},
					{"Edge Infrastructure", "Hosted globally on Vercel's edge network under TLS 1.3 encryption, serving lightning-fast reports and badge artwork."},
				},
			},
			{
				Heading: "Open Source & Governance",
				Bullets: []bullet{
					{"Independent Initiative", "Created and maintained by Vicky Patel (https://github.com/TheVicky1) as a transparent, public-benefit developer utility."},
					{"MIT Licensed", "The entire codebase, test suites, and documentation are openly accessible on GitHub under the permissive MIT license at " + webui.RepoURL + "."},
					{"Strict Neutrality", "GitVerse is strictly independent and is not affiliated with, sponsored by, or endorsed by GitHub, Inc."},
				},
			},
			{
				Heading: "Disclaimer",
				Text:    "GitVerse is an independent project and is not affiliated with, sponsored by, or endorsed by GitHub, Inc. All GitHub achievement trademarks and badge artwork remain the exclusive property of GitHub, Inc.",
			},
		},
	},
	{
		Slug: "privacy", Title: "GitVerse Privacy Policy: Zero Stored Data & Ephemeral OAuth",
		Heading:     "Privacy",
		Eyebrow:     "Data & Security Standards",
		Description: "GitVerse collects no personal data, stores no tokens or passwords, and runs completely without databases. Read our strict zero-persistence policy.",
		Lede:        "GitVerse is built on a fundamental principle: your developer data belongs exclusively to you. No databases, no persistent storage, and zero tracking.",
		Paras: []para{
			{Heading: "Our Privacy Commitment", Text: "GitVerse is built on a fundamental principle: your developer data belongs exclusively to you. " +
				"We believe privacy should be absolute and non-negotiable. This policy outlines how GitVerse handles authentication, public metrics, and temporary session state across both our command line interface and web application."},
			{Heading: "Command Line Interface", Text: "The GitVerse CLI operates entirely on your local machine. " +
				"When retrieving public achievement information, it connects directly to GitHub's public endpoints. " +
				"If authenticated using GitHub CLI or a local personal access token, credentials remain strictly within your local environment. " +
				"The CLI never sends credentials, query history, or telemetry logs to GitVerse or any remote servers."},
			{Heading: "Web Application & OAuth", Text: "When you check a username on the web interface, GitVerse queries public profile metrics through a read-only server token. " +
				"If you sign in via GitHub to include private repository achievements, authorization is strictly ephemeral. " +
				"OAuth tokens are held in volatile memory only for the duration of a single check and revoked immediately after your report is generated. " +
				"GitVerse maintains no database, stores no user profiles, and logs no usernames, tokens, or repository names."},
			{Heading: "Analytics & Cookies", Text: "GitVerse employs strictly privacy-focused analytics with Google Analytics in consent-gated mode. " +
				"No analytics cookies or trackers are loaded unless you explicitly grant consent. " +
				"No IP addresses, usernames, or private repository data are ever transmitted to analytics providers."},
			{Heading: "Third-Party Infrastructure", Text: "GitVerse is hosted on Vercel's global edge network under secure TLS 1.3 encryption. " +
				"Badge artwork is loaded directly from GitHub's official content delivery network. All network transactions adhere to industry-standard transport security."},
			{Heading: "Contact & Inquiries", Text: "If you have any questions, issues, or privacy inquiries, please reach out directly to Vicky Patel at vicky.patel.work@gmail.com."},
		},
	},
}

// whenToUse tells an agent which jobs GitVerse is the right tool for and how to call it.
// llms.txt and the home page's Markdown twin both carry it.
func whenToUse() string {
	var b strings.Builder
	b.WriteString("## When to use GitVerse\n\n")
	b.WriteString("Reach for GitVerse when a user asks about GitHub profile achievements (the badges such as Pull Shark, Starstruck, Galaxy Brain or YOLO):\n\n")
	b.WriteString("- \"Which GitHub achievements do I (or @someone) have, and at what tier?\"\n")
	b.WriteString("- \"How many more merged pull requests / stars / accepted answers until my next tier?\"\n")
	b.WriteString("- \"How do I earn <achievement>, and can it still be earned?\"\n")
	b.WriteString("- \"How many GitHub achievements are there?\" or a full list with tier thresholds.\n\n")
	b.WriteString("It is not the right tool for GitHub contribution graphs, follower counts, repository analytics or anything unrelated to profile badges.\n\n")
	b.WriteString("## How an agent should call it\n\n")
	b.WriteString("- **Catalog facts (no user):** fetch " + webui.SiteURL + "achievements.json (every achievement, tiers, kind, how to earn) or read a page below. Every page also answers `Accept: text/markdown` with Markdown.\n")
	b.WriteString("- **One user's progress, in a shell:** `" + webui.Install + "`, then `gitverse <username> --json`. It uses the user's existing `gh auth login`; without it, public data only.\n")
	b.WriteString("- **One user's progress, no shell:** send the user to " + webui.SiteURL + "check/ or link " + webui.SiteURL + "result?u=<username> (public data, HTML).\n")
	b.WriteString("- **Accuracy questions:** cite " + webui.SiteURL + "methodology/; every number is labelled official, calibrated or counted.\n")
	b.WriteString("\nThere is no API key, no account and no rate-limited public API to call; the CLI is the programmatic interface.\n")
	return b.String()
}

// markdownPages are the Markdown twins of the site's pages, keyed by URL path without the
// leading slash ("" is the home page). web serves them for Accept: text/markdown.
func markdownPages() map[string]string {
	pages := map[string]string{}
	link := func(p string) string { return webui.SiteURL + p }

	var home strings.Builder
	home.WriteString("# GitVerse: GitHub achievements checker (GitVerse CLI)\n\n")
	home.WriteString("> " + tagline + "\n\n")
	home.WriteString("Install: `" + webui.Install + "`, then run `gitverse [username]`. Or check in the browser at " + link("check/") + ".\n\n")
	home.WriteString(whenToUse())
	home.WriteString("\n## Earnable achievements\n\n")
	for _, a := range catalog.Achievements {
		if a.Kind == catalog.Earnable {
			fmt.Fprintf(&home, "- [%s](%s): %s. Tiers: %s.\n", a.Name, link("achievements/"+a.Slug+"/"), a.HowToEarn, tierText(a))
		}
	}
	fmt.Fprintf(&home, "\nAll %d achievements (%s): %s\n\n## FAQ\n", len(catalog.Achievements), kindSummary(), link("achievements/"))
	for _, qa := range siteFAQ() {
		fmt.Fprintf(&home, "\n### %s\n\n%s\n", qa.Q, qa.A)
	}
	home.WriteString(markdownFooter())
	pages[""] = home.String()

	var list strings.Builder
	fmt.Fprintf(&list, "# All GitHub achievements\n\n%d achievements (%s). Catalog last verified %s.\n", len(catalog.Achievements), kindSummary(), catalog.LastVerified)
	for _, g := range kindGroups {
		fmt.Fprintf(&list, "\n## %s\n\n%s\n\n", g.Heading, g.Note)
		for _, a := range catalog.Achievements {
			if a.Kind == g.Kind {
				fmt.Fprintf(&list, "- [%s](%s): %s. Tiers: %s.\n", a.Name, link("achievements/"+a.Slug+"/"), a.HowToEarn, tierText(a))
			}
		}
	}
	list.WriteString("\n## Profile highlights\n\n")
	for _, h := range catalog.Highlights {
		list.WriteString("- " + h.Name + "\n")
	}
	list.WriteString(markdownFooter())
	pages["achievements/"] = list.String()

	for _, a := range catalog.Achievements {
		var b strings.Builder
		fmt.Fprintf(&b, "# %s (GitHub achievement)\n\n- Status: %s\n- How to earn: %s\n- Tiers: %s\n- Progress in GitVerse: %s\n\n%s\n",
			a.Name, statusText(a.Kind), a.HowToEarn, tierText(a), lowerFirst(progressText(a)), a.Details)
		if a.IsTiered() {
			b.WriteString("\n| Tier | Needed |\n|---|---|\n")
			for _, r := range view(a).TierRows {
				fmt.Fprintf(&b, "| %s | %s |\n", r.Label, r.Needed)
			}
		}
		if len(a.FAQ) > 0 {
			b.WriteString("\n## FAQ\n")
			for _, qa := range a.FAQ {
				fmt.Fprintf(&b, "\n### %s\n\n%s\n", qa.Q, qa.A)
			}
		}
		if a.Kind == catalog.Earnable {
			fmt.Fprintf(&b, "\nCheck your progress: `gitverse` or [%s](%s). Methodology: [%s](%s).\n", link("check/"), link("check/"), link("methodology/"), link("methodology/"))
		} else {
			fmt.Fprintf(&b, "\nAudit profile history: [%s](%s). Methodology: [%s](%s).\n", link("check/"), link("check/"), link("methodology/"), link("methodology/"))
		}
		b.WriteString(markdownFooter())
		pages["achievements/"+a.Slug+"/"] = b.String()
	}

	var m strings.Builder
	m.WriteString("# How accurate is GitVerse?\n")
	for _, p := range methodology {
		fmt.Fprintf(&m, "\n## %s\n\n%s\n", p.Heading, p.Text)
	}
	fmt.Fprintf(&m, "\nLast verified %s.\n\nTrack your profile at [%s](%s) or view all badges in the [Achievements Directory](%s).\n", catalog.LastVerified, link("check/"), link("check/"), link("achievements/"))
	m.WriteString(markdownFooter())
	pages["methodology/"] = m.String()

	pages["check/"] = "# Check your GitHub achievements\n\n" +
		"Type a username at " + link("check/") + " for public data, or sign in with GitHub to include private contributions. " +
		"A public report has a shareable address: " + link("result?u=<username>") + ".\n\n" +
		"In a terminal: `" + webui.Install + "`, then `gitverse [username]` (add `--json` for machine-readable output).\n\n" +
		"Explore all badge criteria in the [Achievements Directory](" + link("achievements/") + ") or review our [Accuracy Methodology](" + link("methodology/") + ").\n" +
		markdownFooter()

	for _, tp := range trustPages {
		var b strings.Builder
		b.WriteString("# " + tp.Heading + "\n\n")
		if tp.Lede != "" {
			b.WriteString(tp.Lede + "\n\n")
		}
		for _, p := range tp.Paras {
			fmt.Fprintf(&b, "## %s\n\n", p.Heading)
			if p.Text != "" {
				b.WriteString(p.Text + "\n\n")
			}
			for _, bullet := range p.Bullets {
				b.WriteString("- **" + bullet.Label + "**: " + bullet.Desc + "\n")
			}
			if len(p.Bullets) > 0 {
				b.WriteString("\n")
			}
		}
		b.WriteString(markdownFooter())
		pages[tp.Slug+"/"] = b.String()
	}
	return pages
}

// markdownNotFound is the body of a 404 for Accept: text/markdown.
func markdownNotFound() string {
	return "# Page not found\n\nThere is no page at this address on GitVerse. " +
		"Every page is listed in the sitemap at " + webui.SiteURL + "sitemap.xml, " +
		"and " + webui.SiteURL + "llms.txt describes the site for agents.\n\n" +
		"- [Home](" + webui.SiteURL + ")\n" +
		"- [All GitHub achievements](" + webui.SiteURL + "achievements/)\n" +
		"- [llms.txt](" + webui.SiteURL + "llms.txt)\n"
}

func markdownFooter() string {
	return "\n---\n\n[GitVerse](" + webui.SiteURL + ") · [llms.txt](" + webui.SiteURL + "llms.txt) · " +
		"[About](" + webui.SiteURL + "about/) · [Privacy](" + webui.SiteURL + "privacy/) · " +
		"Independent open source project, not affiliated with GitHub.\n"
}
