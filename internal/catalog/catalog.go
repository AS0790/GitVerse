// Package catalog is the single source of truth for every known GitHub achievement.
//
// GitHub publishes no official list any more. This one follows the community lists
// (Schweinepriester/github-profile-achievements, drknzz/GitHub-Achievements) and was
// checked against live profiles. A badge GitHub adds later still shows up in the CLI,
// because earned badges are read from the profile, not from this table.
package catalog

// LastVerified is the date the catalog was last checked against live profiles.
const LastVerified = "2026-10-17"

// Kind says whether an achievement can be earned today.
type Kind string

const (
	Earnable Kind = "earnable"
	Disabled Kind = "disabled" // tested by GitHub, switched off, not earnable now
	Retired  Kind = "retired"  // a one-off event that is over
	Internal Kind = "internal" // GitHub staff only
)

// Counter names how progress towards the next tier is measured.
type Counter string

const (
	CounterNone      Counter = "none"       // official tier only
	CounterMergedPRs Counter = "merged-prs" // Pull Shark
	CounterStars     Counter = "stars"      // Starstruck
	CounterAnswers   Counter = "answers"    // Galaxy Brain
)

// QA is one FAQ entry. Answers state only what is verified.
type QA struct {
	Q string `json:"q"`
	A string `json:"a"`
}

// Achievement describes one badge.
type Achievement struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	Kind Kind   `json:"kind"`
	// Tiers are the counts for x1, x2, x3, x4. A single-tier badge has one entry.
	// 0 marks a threshold nobody has measured yet.
	Tiers     []int   `json:"tiers"`
	Unit      string  `json:"unit,omitempty"`
	HowToEarn string  `json:"howToEarn"`
	Details   string  `json:"details"`
	Counter   Counter `json:"counter"`
	// NoCounterReason says why a tiered badge shows its tier only.
	NoCounterReason string `json:"noCounterReason,omitempty"`
	FAQ             []QA   `json:"faq,omitempty"`
}

// IsTiered reports whether the badge has more than one tier.
func (a Achievement) IsTiered() bool { return len(a.Tiers) > 1 }

// Highlight is a profile "Highlights" entry, shown under the achievements.
type Highlight struct {
	Name string `json:"name"`
	// GraphQLField is the User field that reports it, when GitHub's API has one.
	GraphQLField string `json:"graphqlField,omitempty"`
	HowToEarn    string `json:"howToEarn"`
}

// Achievements lists every known badge, in display order.
var Achievements = []Achievement{
	{
		Slug: "pull-shark", Name: "Pull Shark", Kind: Earnable,
		Tiers: []int{2, 16, 128, 1024}, Unit: "merged pull requests", Counter: CounterMergedPRs,
		HowToEarn: "Open pull requests that get merged",
		Details: "Pull Shark recognizes developers who open pull requests that get successfully merged across public or private repositories. " +
			"GitHub awards progressive tier medals at 2, 16, 128, and 1024 merged pull requests. " +
			"GitHub's underlying achievement history logs the exact pull request and repository associated with each milestone unlock.",
		FAQ: []QA{
			{"How many merged pull requests do you need for Pull Shark?",
				"2 merged pull requests for the base badge (x1), 16 for bronze (x2), 128 for silver (x3), and 1024 for gold (x4)."},
			{"Do pull requests in private repositories count towards Pull Shark?",
				"Yes. GitHub's Pull Shark history records private repository milestones as \"inaccessible\", confirming they are counted even when private from public visitors."},
			{"Can you earn Pull Shark by merging pull requests into your own repositories?",
				"Yes. Merging pull requests on personal or organization repositories that you own officially increments your Pull Shark counter."},
		},
	},
	{
		Slug: "starstruck", Name: "Starstruck", Kind: Earnable,
		Tiers: []int{16, 128, 512, 4096}, Unit: "stars on one repository", Counter: CounterStars,
		HowToEarn: "Create a repository that gets stars",
		Details: "Starstruck celebrates repository creators who build open source projects that earn community stars. " +
			"Thresholds are evaluated per individual repository rather than aggregated across multiple projects. " +
			"GitHub awards tiers at 16, 128, 512, and 4096 stargazers on a single repository.",
		FAQ: []QA{
			{"How many stars do you need for Starstruck?",
				"A single repository requires 16 stars for base (x1), 128 for bronze (x2), 512 for silver (x3), and 4096 for gold (x4)."},
			{"Do stars across several repositories add up for Starstruck?",
				"No. The thresholds apply strictly to a single repository. GitHub's badge history records one specific repository per threshold unlock."},
			{"Do stars on forked repositories count toward Starstruck?",
				"No. GitHub only evaluates stargazers on original, non-forked repositories that your account created."},
		},
	},
	{
		Slug: "pair-extraordinaire", Name: "Pair Extraordinaire", Kind: Earnable,
		Tiers: []int{1, 10, 24, 48}, Unit: "coauthored merged pull requests", Counter: CounterNone,
		NoCounterReason: "GitHub's counting rule is not public: tier only, no counter",
		HowToEarn:       "Coauthor commits on merged pull requests",
		Details: "Pair Extraordinaire honors collaborative programming through Git co-authorship. " +
			"It is unlocked by authoring or co-authoring commits included in merged pull requests. " +
			"GitHub logs the merged pull request and collaborator identities for milestones at 1, 10, 24, and 48 co-authored PRs.",
		FAQ: []QA{
			{"How many coauthored pull requests do you need for Pair Extraordinaire?",
				"1 co-authored merged PR for base (x1), 10 for bronze (x2), 24 for silver (x3), and 48 for gold (x4)."},
			{"What Git syntax is required to credit a co-author?",
				"Add a commit trailer formatted as `Co-authored-by: Name <email@example.com>` at the bottom of the Git commit message."},
			{"Why does GitVerse show no counter for Pair Extraordinaire?",
				"GitHub's internal qualification logic for co-authored commits is unexposed and differs from public commit search. GitVerse surfaces the verified official tier to prevent misleading estimates."},
		},
	},
	{
		Slug: "galaxy-brain", Name: "Galaxy Brain", Kind: Earnable,
		Tiers: []int{2, 8, 16, 32}, Unit: "accepted answers", Counter: CounterAnswers,
		HowToEarn: "Have discussion answers accepted",
		Details: "Galaxy Brain recognizes developers who provide helpful technical mentorship by having their answers marked as accepted in GitHub Discussions. " +
			"Milestone tiers are reached at 2, 8, 16, and 32 accepted solutions.",
		FAQ: []QA{
			{"How many accepted answers do you need for Galaxy Brain?",
				"2 accepted answers for base (x1), 8 for bronze (x2), 16 for silver (x3), and 32 for gold (x4)."},
			{"Do answers in GitHub Community Discussions count for Galaxy Brain?",
				"Not since February 6, 2024, when GitHub paused achievement qualification in github.com/orgs/community to deter automated spam. Answers in all other open-source repository discussions count normally."},
			{"Who can mark an answer as accepted in GitHub Discussions?",
				"The original discussion author or any repository maintainer with triage or write permissions can click \"Mark as answer\"."},
		},
	},
	{
		Slug: "quickdraw", Name: "Quickdraw", Kind: Earnable, Tiers: []int{1}, Counter: CounterNone,
		HowToEarn: "Close an issue or pull request within 5 minutes of opening it",
		Details: "Quickdraw (\"Gitty up!\") is a binary achievement awarded for closing an issue or pull request within 5 minutes (300 seconds) of its creation. " +
			"It is commonly unlocked during rapid triage, accidental issue duplication, or automated workflow fixes.",
		FAQ: []QA{
			{"Can you get Quickdraw by closing your own issue?",
				"Yes. Opening an issue or pull request and manually closing it within 5 minutes triggers the achievement."},
			{"Does Quickdraw have multiple tiers or multipliers?",
				"No. Quickdraw is a single-tier binary unlock with no progressive tier rings."},
		},
	},
	{
		Slug: "yolo", Name: "YOLO", Kind: Earnable, Tiers: []int{1}, Counter: CounterNone,
		HowToEarn: "Merge a pull request without a review",
		Details: "YOLO (\"You want it? You merge it.\") is a binary achievement granted when you merge a pull request into a default branch without any approving reviews. " +
			"It reflects direct developer autonomy on repositories where branch protection rules do not enforce peer review.",
		FAQ: []QA{
			{"How do you unlock the YOLO badge on GitHub?",
				"Open a pull request on a repository where branch protection rules allow merging without reviews, and merge it directly with 0 approvals."},
			{"Does having an approving review prevent unlocking YOLO?",
				"Yes. If a pull request receives an approving code review prior to merge, it qualifies as reviewed and will not award the YOLO badge."},
		},
	},
	{
		Slug: "public-sponsor", Name: "Public Sponsor", Kind: Earnable, Tiers: []int{1}, Counter: CounterNone,
		HowToEarn: "Sponsor someone publicly through GitHub Sponsors",
		Details: "Public Sponsor recognizes community members who financially back open source software creators via GitHub Sponsors. " +
			"Sponsorship must have public visibility enabled on your profile to unlock this recognition.",
		FAQ: []QA{
			{"What qualifies for the Public Sponsor badge?",
				"Any active public monthly or one-time sponsorship to an open-source maintainer or organization through GitHub Sponsors."},
			{"Do private sponsorships unlock the Public Sponsor badge?",
				"No. Private sponsorships intentionally withhold sponsorship visibility and do not display this public profile achievement."},
		},
	},
	{
		Slug: "finishing-touch", Name: "Finishing Touch", Kind: Earnable,
		Tiers: []int{1, 4, 16, 32}, Unit: "merged UI quality pull requests", Counter: CounterNone,
		NoCounterReason: "GitHub decides which pull requests count as UI quality: tier only, no counter",
		HowToEarn:       "Get pull requests that improve GitHub's UI quality merged",
		Details: "Finishing Touch counts merged pull requests classified by GitHub as improving UI and frontend quality. " +
			"Milestone rings are awarded at 1, 4, 16, and 32 merged contributions that enhance interface polishing.",
		FAQ: []QA{
			{"How many pull requests do you need for Finishing Touch?",
				"1 merged UI quality pull request for base (x1), 4 for bronze (x2), 16 for silver (x3), and 32 for gold (x4)."},
			{"How does GitHub identify UI quality pull requests?",
				"GitHub uses internal heuristics and commit metadata evaluating frontend enhancements, CSS modifications, and linked UI issue resolutions."},
		},
	},
	{
		Slug: "heart-on-your-sleeve", Name: "Heart On Your Sleeve", Kind: Disabled,
		Tiers: []int{0, 16, 128, 0}, Unit: "❤️ reactions", Counter: CounterNone,
		HowToEarn: "React to something on GitHub with a ❤️ emoji",
		Details: "Heart On Your Sleeve was an experimental achievement celebrating positive community interactions through emoji reactions. " +
			"GitHub switched it off to prevent automated reaction spam; existing earners retain their badge on profile pages.",
		FAQ: []QA{
			{"Can you still earn the Heart On Your Sleeve badge?",
				"No. GitHub concluded testing for this achievement and deactivated it. Profiles that earned it previously keep it on display."},
			{"What were the known milestone thresholds for Heart On Your Sleeve?",
				"Verified community records document 16 ❤️ reactions for Tier 2 (bronze) and 128 for Tier 3 (silver)."},
		},
	},
	{
		Slug: "open-sourcerer", Name: "Open Sourcerer", Kind: Disabled,
		Tiers: []int{0, 8, 16, 64}, Unit: "merged open source pull requests", Counter: CounterNone,
		HowToEarn: "Have pull requests merged in several public repositories",
		Details: "Open Sourcerer rewarded developers who contributed across multiple diverse open source repositories. " +
			"GitHub tested the badge during early gamification pilots and currently has it switched off for new contributions.",
		FAQ: []QA{
			{"Can new contributors unlock Open Sourcerer today?",
				"No. Open Sourcerer is currently inactive and cannot be unlocked by new pull requests."},
			{"What were the historical thresholds for Open Sourcerer?",
				"Community documentation verified milestones at 8 merged PRs for Tier 2, 16 for Tier 3, and 64 for Tier 4."},
		},
	},
	{
		Slug: "arctic-code-vault-contributor", Name: "Arctic Code Vault Contributor", Kind: Retired,
		Tiers: []int{1}, Counter: CounterNone,
		HowToEarn: "Code in a repository of the 2020 GitHub Archive Program",
		Details: "Arctic Code Vault Contributor commemorates developers whose public code was captured during the GitHub Archive Program snapshot on February 2, 2020. " +
			"The archive is preserved on silver halide film in a decommissioned coal mine in Svalbard, Norway, designed to last 1,000 years.",
		FAQ: []QA{
			{"Is Arctic Code Vault Contributor still earnable?",
				"No. The historical snapshot was taken on 02/02/2020 and cannot be retroactively added to new accounts."},
			{"How was eligibility determined for the Arctic Code Vault badge?",
				"Any developer who authored at least one commit in the default branch of an active public repository captured in the 2020 archive received the badge."},
		},
	},
	{
		Slug: "mars-2020-contributor", Name: "Mars 2020 Contributor", Kind: Retired,
		Tiers: []int{1}, Counter: CounterNone,
		HowToEarn: "Code in a repository used by the Mars 2020 Helicopter Mission",
		Details: "Mars 2020 Contributor honors developers who authored code or contributed to open source libraries utilized by NASA Jet Propulsion Laboratory (JPL) " +
			"in the Ingenuity Mars Helicopter mission.",
		FAQ: []QA{
			{"Can you still get the Mars 2020 Contributor badge?",
				"No. The contribution manifest was finalized following the NASA Perseverance / Ingenuity Mars mission deployment in 2020."},
			{"Which open source projects were included in the Mars 2020 mission?",
				"NASA JPL cited thousands of open source dependencies across Linux, Python, C++, and orbital mechanics packages cataloged in the mission software registry."},
		},
	},
	{
		Slug: "proxima-pioneer", Name: "Proxima Pioneer", Kind: Internal, Tiers: []int{1}, Counter: CounterNone,
		HowToEarn: "GitHub staff: contributed to the Proxima proof of concept",
		Details: "Proxima Pioneer is an exclusive internal achievement awarded to GitHub staff who participated in early architecture and proof-of-concept development for Proxima.",
		FAQ: []QA{
			{"Can public developers earn Proxima Pioneer?",
				"No. Proxima Pioneer is strictly an internal employee badge issued exclusively to GitHub staff."},
			{"What is Proxima at GitHub?",
				"Proxima refers to internal developer infrastructure and workstation initiatives within GitHub's core engineering organization."},
		},
	},
	{
		Slug: "proxima-staffshipper", Name: "Proxima Staffshipper", Kind: Internal, Tiers: []int{1}, Counter: CounterNone,
		HowToEarn: "GitHub staff: shipped the Proxima staffship instance",
		Details: "Proxima Staffshipper is an internal achievement awarded to GitHub employees who successfully deployed production services into the Proxima staffship environment.",
		FAQ: []QA{
			{"Can public contributors earn Proxima Staffshipper?",
				"No. Proxima Staffshipper is reserved strictly for internal GitHub engineering teams."},
		},
	},
	{
		Slug: "proxima-staffuser", Name: "Proxima Staffuser", Kind: Internal, Tiers: []int{1}, Counter: CounterNone,
		HowToEarn: "GitHub staff: your team was onboarded to Proxima",
		Details: "Proxima Staffuser is an internal employee recognition badge granted to GitHub staff members whose engineering teams were onboarded into Proxima.",
		FAQ: []QA{
			{"Can external developers unlock Proxima Staffuser?",
				"No. It is an internal employee milestone restricted to verified GitHub staff accounts."},
		},
	},
}

// Highlights lists the profile highlights.
var Highlights = []Highlight{
	{Name: "Developer Program Member", GraphQLField: "isDeveloperProgramMember", HowToEarn: "Register for the GitHub Developer Program"},
	{Name: "Security Bug Bounty Hunter", GraphQLField: "isBountyHunter", HowToEarn: "Report a valid vulnerability through GitHub Security Bug Bounty"},
	{Name: "GitHub Campus Expert", GraphQLField: "isCampusExpert", HowToEarn: "Join the GitHub Campus Expert program"},
	{Name: "GitHub Star", GraphQLField: "isGitHubStar", HowToEarn: "Be selected for the GitHub Stars program"},
	{Name: "Pro", HowToEarn: "Subscribe to GitHub Pro"},
	{Name: "Security advisory credit", HowToEarn: "Have a security advisory you submitted accepted into the GitHub Advisory Database"},
}

// BySlug returns the catalog entry for slug.
func BySlug(slug string) (Achievement, bool) {
	for _, a := range Achievements {
		if a.Slug == slug {
			return a, true
		}
	}
	return Achievement{}, false
}

// KindTitle is the section heading for a kind.
func KindTitle(k Kind) string {
	switch k {
	case Earnable:
		return "EARNABLE"
	case Disabled:
		return "DISABLED · tested by GitHub, not earnable now"
	case Retired:
		return "RETIRED · can no longer be earned"
	case Internal:
		return "INTERNAL · GitHub staff only"
	}
	return string(k)
}

// Kinds is the display order of sections.
var Kinds = []Kind{Earnable, Disabled, Retired, Internal}
