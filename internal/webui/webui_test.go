package webui

import (
	"strings"
	"testing"
)

func TestRender_usesTheVercelDomainAndRootPaths(t *testing.T) {
	b, err := Render(Page{Path: "404.html", URL: SiteURL + "404.html", Title: "T", Description: "D",
		Heading: "H", Content: "404.html", NoIndex: true})
	if err != nil {
		t.Fatal(err)
	}
	html := string(b)
	for _, want := range []string{
		`<link rel="stylesheet" href="/style.css">`,
		`<script src="/theme.js"></script>`,
		`<button type="button" class="theme-btn" data-theme-toggle hidden>`,
		`<script src="/app.js" defer data-ga="G-3EEKKGL110"></script>`,
		`<meta name="robots" content="noindex">`,
		`id="consent" hidden`,
		`<span class="star-count" data-stars hidden></span>`,
		`<button type="button" class="menu-btn" data-menu aria-expanded="false" aria-controls="site-nav">Menu</button>`,
		`<a href="/methodology/">How it works</a>`,
		`gitverse-favicon.png`,
	} {
		if !strings.Contains(html, want) {
			t.Errorf("missing %s", want)
		}
	}
	if strings.Contains(html, "github.io") || strings.Contains(html, "/GitVerse/") {
		t.Error("old GitHub Pages paths remain")
	}
}

func TestRender_noInlineScripts(t *testing.T) {
	// The CSP forbids inline scripts; only JSON-LD data blocks are allowed inline.
	b, _ := Render(Page{Path: "404.html", URL: SiteURL, Title: "T", Description: "D", Heading: "H", Content: "404.html"})
	for _, chunk := range strings.Split(string(b), "<script")[1:] {
		head := chunk[:strings.Index(chunk, ">")]
		if !strings.Contains(head, "src=") && !strings.Contains(head, `type="application/ld+json"`) {
			t.Fatalf("inline script: <script%s>", head)
		}
	}
}

func TestRel(t *testing.T) {
	if got := Rel(SiteURL + "achievements/"); got != "/achievements/" {
		t.Fatalf("Rel = %q", got)
	}
}

func TestInitials(t *testing.T) {
	for name, want := range map[string]string{
		"Heart On Your Sleeve": "HO", "YOLO": "Y", "open sourcerer": "OS", "": "", "Édith Piaf": "ÉP",
	} {
		if got := Initials(name); got != want {
			t.Errorf("Initials(%q) = %q, want %q", name, got, want)
		}
	}
}

func TestStyle_openCardAndLongLinkRules(t *testing.T) {
	css, err := Static("style.css")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"box-sizing: border-box",
		"--bg:        #0a0a0b;",
		"--gold:      #c9971f;",
		"color-scheme: dark;",
	} {
		if !strings.Contains(string(css), want) {
			t.Errorf("style.css lacks %q", want)
		}
	}
}

func TestStyle_headerFitsOneRowOnAPhone(t *testing.T) {
	css, err := Static("style.css")
	if err != nil {
		t.Fatal(err)
	}
	s := string(css)
	if !strings.Contains(s, "@media") {
		t.Fatal("style.css has no responsive media query blocks")
	}
}

func TestStyle_darkTokensMatchInBothBlocks(t *testing.T) {
	css, err := Static("style.css")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"--bg:        #0a0a0b;", "--gold:      #c9971f;", "color-scheme: dark;", "font-display: swap"} {
		if !strings.Contains(string(css), want) {
			t.Errorf("style.css lacks %s", want)
		}
	}
}

func TestThemeJS_appliesOnlyAStoredLightOrDark(t *testing.T) {
	js, err := Static("theme.js")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`localStorage.getItem("gha:theme")`, `t === "light" || t === "dark"`,
		`setAttribute("data-theme", t)`, `classList.add("js")`} {
		if !strings.Contains(string(js), want) {
			t.Errorf("theme.js lacks %s", want)
		}
	}
}

// The theme key lives in two scripts (theme.js applies it, app.js writes it): they must agree.
func TestAppJS_themeKeyMatchesThemeJS(t *testing.T) {
	js, err := Static("app.js")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(js), `var THEME = "gha:theme";`) {
		t.Error(`app.js lacks var THEME = "gha:theme";`)
	}
}

func TestStarBanner_namesItsPlaceAndTheRepo(t *testing.T) {
	b, err := RenderFragment("star-banner", "home")
	if err != nil {
		t.Fatal(err)
	}
	html := string(b)
	for _, want := range []string{`<a class="star-banner" href="https://github.com/TheVicky1/GitVerse.git" rel="noopener" data-star="home">`,
		"Found this useful? Star GitVerse on GitHub", `data-stars hidden`} {
		if !strings.Contains(html, want) {
			t.Errorf("star banner lacks %s", want)
		}
	}
}

// Under 640px the nav precedes the buttons in the DOM but shows below them: opening the menu
// moves focus into it, and Escape closes it and returns focus to the button.
func TestAppJS_mobileMenuManagesFocus(t *testing.T) {
	js, err := Static("app.js")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`var first = nav.querySelector("a");`, `if (first) first.focus();`,
		`e.key !== "Escape"`, `nav.hasAttribute("data-open")`, `nav.removeAttribute("data-open");`,
		`btn.setAttribute("aria-expanded", "false");`, `btn.focus();`} {
		if !strings.Contains(string(js), want) {
			t.Errorf("app.js lacks %s", want)
		}
	}
}

// The star count is fetched anonymously, without a referrer (the result URL names a user),
// cached for the visit, and only the place of a click reaches GA.
func TestAppJS_starCountAndClick(t *testing.T) {
	js, err := Static("app.js")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"https://api.github.com/repos/TheVicky1/GitVerse"`, `"gitverse:stars"`, `credentials: "omit"`,
		`referrerPolicy: "no-referrer"`, `sessionStorage`, `track("star_click", { place: t.dataset.star })`,
		`fillStars: fillStars`, `[data-menu]`, `aria-expanded`,
		// The GitVerse counter begins at zero and updates from the repository API.
		`if (stars === null) return;`} {
		if !strings.Contains(string(js), want) {
			t.Errorf("app.js lacks %s", want)
		}
	}
}

func TestStyle_shelfCaptionsUseTheirOwnMutedToken(t *testing.T) {
	css, err := Static("style.css")
	if err != nil {
		t.Fatal(err)
	}
	s := string(css)
	for _, want := range []string{"--surface:   #111113;", "--ink-2:     #a1a1aa;"} {
		if !strings.Contains(s, want) {
			t.Errorf("style.css lacks %s", want)
		}
	}
}
