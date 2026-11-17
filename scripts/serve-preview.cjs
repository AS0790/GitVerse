const http = require("http");
const https = require("https");
const fs = require("fs");
const path = require("path");

const root = path.resolve(__dirname, "..", "site");
const catalogPath = path.resolve(__dirname, "..", "site", "achievements.json");
const catalog = fs.existsSync(catalogPath) ? JSON.parse(fs.readFileSync(catalogPath, "utf8")) : { achievements: [] };

const types = {
  ".css": "text/css; charset=utf-8",
  ".html": "text/html; charset=utf-8",
  ".js": "application/javascript; charset=utf-8",
  ".json": "application/json; charset=utf-8",
  ".png": "image/png",
  ".svg": "image/svg+xml",
  ".woff2": "font/woff2",
  ".ttf": "font/ttf",
  ".txt": "text/plain; charset=utf-8"
};

const port = 5173;

// Cache to prevent GitHub rate-limiting during local testing
const profileCache = new Map();
const CACHE_TTL_MS = 60 * 1000; // 60 seconds

const communityArt = "https://raw.githubusercontent.com/Schweinepriester/github-profile-achievements/549ba7291ccb68efcff5daf8af8c2c0ac6817b6f/images/";

const badgeImages = {
  "heart-on-your-sleeve": [
    communityArt + "heart-on-your-sleeve-default.png",
    communityArt + "tiers/heart-on-your-sleeve-bronze.png",
    communityArt + "tiers/heart-on-your-sleeve-silver.png",
    communityArt + "tiers/heart-on-your-sleeve-gold.png"
  ],
  "open-sourcerer": [
    communityArt + "open-sourcerer-default.png",
    communityArt + "tiers/open-sourcerer-bronze.png",
    communityArt + "tiers/open-sourcerer-silver.png",
    communityArt + "tiers/open-sourcerer-gold.png"
  ],
  "pull-shark": [
    "https://github.githubassets.com/assets/pull-shark-default-498c279a747d.png",
    "https://github.githubassets.com/assets/pull-shark-bronze-a37accb528d1.png",
    "https://github.githubassets.com/assets/pull-shark-silver-0643f87ac9fd.png",
    "https://github.githubassets.com/assets/pull-shark-gold-90985540b385.png"
  ],
  "starstruck": [
    "https://github.githubassets.com/assets/starstruck-default-b6610abad518.png",
    "https://github.githubassets.com/assets/starstruck-bronze-b71ec4d7766a.png",
    "https://github.githubassets.com/assets/starstruck-silver-275c101223c0.png",
    "https://github.githubassets.com/assets/starstruck-gold-a35e077693cc.png"
  ],
  "pair-extraordinaire": [
    "https://github.githubassets.com/assets/pair-extraordinaire-default-579438a20e01.png",
    "https://github.githubassets.com/assets/pair-extraordinaire-bronze-797214f41e7e.png",
    "https://github.githubassets.com/assets/pair-extraordinaire-silver-a0b2b472348c.png",
    "https://github.githubassets.com/assets/pair-extraordinaire-gold-5335a3b8ae7d.png"
  ],
  "galaxy-brain": [
    "https://github.githubassets.com/assets/galaxy-brain-default-847262c21056.png",
    "https://github.githubassets.com/assets/galaxy-brain-bronze-da3c4829fd22.png",
    "https://github.githubassets.com/assets/galaxy-brain-silver-2ce8da282a85.png",
    "https://github.githubassets.com/assets/galaxy-brain-gold-9c2c03fd1acf.png"
  ],
  "finishing-touch": [
    "https://github.githubassets.com/assets/finishing-touch-default-764b3a996d62.png",
    "https://github.githubassets.com/assets/finishing-touch-bronze-7c49185628f3.png",
    "https://github.githubassets.com/assets/finishing-touch-silver-75ab9a5c22ba.png",
    "https://github.githubassets.com/assets/finishing-touch-gold-e19e1fcbfee2.png"
  ],
  "quickdraw": ["https://github.githubassets.com/assets/quickdraw-default-39c6aec8ff89.png"],
  "yolo": ["https://github.githubassets.com/assets/yolo-default-be0bbff04951.png"],
  "public-sponsor": ["https://github.githubassets.com/assets/public-sponsor-default-9fa68986b057.png"],
  "arctic-code-vault-contributor": ["https://github.githubassets.com/assets/arctic-code-vault-contributor-default-df8d74122a06.png"],
  "mars-2020-contributor": ["https://github.githubassets.com/assets/mars-2020-contributor-default-c9cad8495acf.png"],
  "proxima-pioneer": ["https://github.githubassets.com/assets/proxima-pioneer-default-51cd4e4969d5.png"],
  "proxima-staffshipper": ["https://github.githubassets.com/assets/proxima-staffshipper-default-84f658288021.png"],
  "proxima-staffuser": ["https://github.githubassets.com/assets/proxima-staffuser-default-02bf6163ea31.png"]
};

function getBadgeImage(slug, tier = 1) {
  const list = badgeImages[slug] || [];
  if (list.length === 0) return "";
  const idx = Math.min(Math.max((tier || 1) - 1, 0), list.length - 1);
  return list[idx];
}

function tierClass(tier) {
  if (tier === 4) return 'gold';
  if (tier === 3) return 'silver';
  if (tier === 2) return 'bronze';
  return 'default';
}

function tierText(tiers) {
  if (!tiers || tiers.length < 2) return "";
  return tiers.map(t => (t > 0 ? t : "?")).join(" · ");
}

function stripTags(html) {
  return html.replace(/<[^>]+>/g, ' ').replace(/\s+/g, ' ').trim();
}

function fetchUrl(url, headers = {}) {
  return new Promise((resolve, reject) => {
    const req = https.get(url, {
      headers: {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36',
        ...headers
      }
    }, res => {
      let data = '';
      res.on('data', chunk => data += chunk);
      res.on('end', () => resolve({ status: res.statusCode, body: data, headers: res.headers }));
    });
    req.on('error', reject);
    req.setTimeout(15000, () => {
      req.destroy();
      reject(new Error('Request timeout: ' + url));
    });
  });
}

async function fetchUserProfile(username) {
  // Validate username
  if (!/^[A-Za-z0-9](?:[A-Za-z0-9]|-(?=[A-Za-z0-9])){0,38}$/.test(username)) {
    return { error: 'invalid', message: 'That is not a GitHub username.' };
  }

  const cached = profileCache.get(username.toLowerCase());
  if (cached && (Date.now() - cached.timestamp < CACHE_TTL_MS)) {
    return cached.data;
  }

  // 1. Fetch achievements tab
  const tabRes = await fetchUrl(`https://github.com/${username}?tab=achievements`);
  let isHidden = false;
  let tabHtml = '';

  if (tabRes.status === 404) {
    // Check if user exists on GitHub
    const profileRes = await fetchUrl(`https://github.com/${username}`);
    if (profileRes.status === 404) {
      return { error: 'not_found', message: `GitHub has no user called ${username}.` };
    }
    isHidden = true;
    tabHtml = profileRes.body;
  } else if (tabRes.status === 200) {
    tabHtml = tabRes.body;
  } else {
    return { error: 'github', message: 'GitHub did not answer. Please try again in a minute.' };
  }

  // Extract earned badge slugs
  const earnedSlugs = [];
  const badgeRegex = /data-hovercard-url="\/users\/[^\/"]+\/achievements\/([a-z0-9-]+)\/detail/g;
  let m;
  while ((m = badgeRegex.exec(tabHtml)) !== null) {
    if (!earnedSlugs.includes(m[1])) {
      earnedSlugs.push(m[1]);
    }
  }

  // Extract highlights
  const highlightsList = [
    { name: "Developer Program Member", earned: false },
    { name: "Security Bug Bounty Hunter", earned: false },
    { name: "GitHub Campus Expert", earned: false },
    { name: "GitHub Star", earned: false },
    { name: "Pro", earned: false },
    { name: "Security advisory credit", earned: false }
  ];

  const highlightsBlock = tabHtml.match(/>Highlights<\/h2>([\s\S]*?)<\/ul>/i);
  if (highlightsBlock) {
    const rawItems = highlightsBlock[1].match(/<li[^>]*>([\s\S]*?)<\/li>/gi) || [];
    const foundText = rawItems.map(item => stripTags(item).toLowerCase());
    highlightsList.forEach(h => {
      const target = h.name.toLowerCase();
      h.earned = foundText.some(item => item.includes(target) || (target === 'pro' && /\bpro\b/.test(item)));
    });
  }

  // Fetch details for each earned badge
  const details = {};
  for (const slug of earnedSlugs) {
    try {
      const detailRes = await fetchUrl(`https://github.com/users/${username}/achievements/${slug}/detail`, {
        'X-Requested-With': 'XMLHttpRequest'
      });
      if (detailRes.status === 200) {
        const dHtml = detailRes.body;
        const tierMatch = dHtml.match(/tier-count="(\d+)"/);
        const tier = tierMatch ? parseInt(tierMatch[1], 10) : 1;
        const unlockedMatch = dHtml.match(/datetime="(\d{4}-\d{2}-\d{2})/);
        const unlocked = unlockedMatch ? unlockedMatch[1] : '';

        // Extract milestones
        const milestones = [];
        const historyMatches = dHtml.matchAll(/achievement-history-tier[\s\S]*?TimelineItem-body[^>]*>([\s\S]*?)<\/div>/g);
        for (const hm of historyMatches) {
          const body = hm[1];
          const prMatch = body.match(/href="\/([^\/"]+\/[^\/"]+)\/pull\/(\d+)"/);
          const repoMatch = body.match(/href="\/([^\/"]+\/[^\/"]+)"/);
          let link = '';
          if (prMatch) {
            link = `https://github.com/${prMatch[1]}/pull/${prMatch[2]}`;
          } else if (repoMatch) {
            link = `https://github.com/${repoMatch[1]}`;
          }
          milestones.push({
            label: stripTags(body),
            url: link
          });
        }

        details[slug] = {
          tier,
          unlocked,
          milestones
        };
      }
    } catch (e) {
      console.error(`Error fetching detail for ${slug}:`, e.message);
    }
  }

  const result = {
    username,
    isHidden,
    earnedSlugs,
    details,
    highlights: highlightsList
  };

  profileCache.set(username.toLowerCase(), {
    timestamp: Date.now(),
    data: result
  });

  return result;
}

function renderBadgeCard(item) {
  const isLocked = !item.earned;
  const isMultiTier = item.tiers && item.tiers.length > 1;
  const tier = item.tier || 1;
  const tierCls = tierClass(tier);
  const tierPill = item.earned ? (isMultiTier ? `x${tier}` : '✓') : '';
  const imgUrl = getBadgeImage(item.slug, item.earned ? tier : 1);
  const altText = item.earned ? (isMultiTier ? `${item.name}, ${tierCls} tier (x${tier})` : `${item.name} badge`) : `${item.name} badge, not earned`;

  let line = '';
  if (item.earned) {
    if (isMultiTier) {
      if (tier >= item.tiers.length) {
        line = 'Top tier';
      } else {
        line = `x${tier}, next tier at ${item.tiers[tier]}`;
      }
    } else {
      const yearMatch = (item.unlocked || '').match(/\b(19|20)\d{2}\b/);
      line = yearMatch ? `Since ${yearMatch[0]}` : 'Unlocked';
    }
  } else {
    if (item.kind !== 'earnable') {
      line = item.howToEarn || 'Archived';
    } else if (isMultiTier) {
      line = `First tier at ${item.tiers[0]}`;
    } else {
      line = item.howToEarn || 'Available';
    }
  }

  let milestonesHtml = '';
  if (item.milestones && item.milestones.length > 0) {
    const listItems = item.milestones.map(m => {
      if (m.url) {
        return `<li><a href="${m.url}" target="_blank" rel="noopener">${m.label}</a></li>`;
      }
      return `<li>${m.label}</li>`;
    }).join('');
    milestonesHtml = `<ol class="milestones">${listItems}</ol>`;
  }

  const aboutLink = `/achievements/${item.slug}/`;

  return `
    <details class="badge-card${isLocked ? ' locked' : ' earned'}" data-badge="${item.slug}">
      <summary>
        ${imgUrl ? `<img class="badge-img" src="${imgUrl}" width="96" height="96" alt="${altText}" loading="lazy">` : `<span class="badge-img placeholder" aria-hidden="true">${item.name.slice(0, 2).toUpperCase()}</span>`}
        <span class="badge-name">${item.name}${tierPill ? ` <span class="tier-pill tier-${tierCls}">${tierPill}</span>` : ''}</span>
        <span class="badge-line">${line}</span>
      </summary>
      <div class="badge-detail">
        <p>${item.howToEarn || item.name}.${tierText(item.tiers) ? ` Tiers: ${tierText(item.tiers)}.` : ''}${item.unlocked ? ` First unlocked ${item.unlocked}.` : ''}</p>
        <p class="meta"><strong>official</strong>: Read from GitHub's own profile pages.</p>
        ${milestonesHtml}
        <p><a href="${aboutLink}" class="badge-detail-link">About ${item.name} →</a></p>
      </div>
    </details>
  `;
}

function renderResultHtml(data) {
  const { username, isHidden, earnedSlugs, details, highlights } = data;

  // Classify achievements from catalog
  const displayItems = [];
  const collectItems = [];
  const archiveItems = [];

  let earnedEarnable = 0;
  let tierTotal = 0;

  for (const ach of catalog.achievements) {
    const isEarned = earnedSlugs.includes(ach.slug);
    const detail = details[ach.slug] || {};
    const item = {
      ...ach,
      earned: isEarned,
      tier: detail.tier || (isEarned ? 1 : 0),
      unlocked: detail.unlocked || '',
      milestones: detail.milestones || []
    };

    if (isEarned) {
      displayItems.push(item);
      tierTotal += (item.tier || 1);
      if (ach.kind === 'earnable') {
        earnedEarnable++;
      }
    } else {
      if (ach.kind === 'earnable') {
        collectItems.push(item);
      } else {
        archiveItems.push(item);
      }
    }
  }

  // Also include any earned badges from GitHub that might not be in our catalog
  for (const slug of earnedSlugs) {
    if (!catalog.achievements.some(a => a.slug === slug)) {
      const detail = details[slug] || {};
      const item = {
        slug,
        name: slug.split('-').map(w => w.charAt(0).toUpperCase() + w.slice(1)).join(' '),
        kind: 'other',
        tiers: [1],
        earned: true,
        tier: detail.tier || 1,
        unlocked: detail.unlocked || '',
        milestones: detail.milestones || []
      };
      displayItems.push(item);
      tierTotal += item.tier;
    }
  }

  // Determine closest next unlock
  let nextUnlockText = '';
  // Check if any earned multi-tier badge can advance
  for (const item of displayItems) {
    if (item.kind === 'earnable' && item.tiers && item.tiers.length > 1 && item.tier < item.tiers.length) {
      nextUnlockText = `${item.name}: tier x${item.tier + 1} at ${item.tiers[item.tier]}`;
      break;
    }
  }
  if (!nextUnlockText && collectItems.length > 0) {
    const nextAch = collectItems[0];
    nextUnlockText = `${nextAch.name}: ${nextAch.howToEarn || 'Next unlock'}`;
  }

  const summaryText = `${earnedEarnable} of 8 earnable badges · ${tierTotal} tiers unlocked`;

  return `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>@${username}'s GitHub achievements · GitVerse</title>
<meta name="description" content="Every GitHub achievement, @${username}'s tier and how far to the next one.">
<meta name="robots" content="noindex">
<meta name="color-scheme" content="dark">
<meta property="og:type" content="website">
<meta property="og:site_name" content="GitVerse">
<meta property="og:title" content="@${username}'s GitHub achievements · GitVerse">
<meta property="og:description" content="Every GitHub achievement, @${username}'s tier and how far to the next one.">
<meta property="og:image" content="https://github.com/${username}.png?size=640">
<link rel="icon" type="image/png" href="/gitverse-favicon.png">
<script src="/theme.js"></script>
<link rel="stylesheet" href="/style.css">
</head>
<body>
<header class="top">
  <div class="top-in">
    <a class="brand" href="/"><img class="brand-wordmark" src="/gitverse-wordmark-transparent.png" alt="GitVerse"></a>
    <nav id="site-nav" aria-label="Site">
      <a href="/check/">Track</a>
      <a href="/achievements/">Explore</a>
      <a href="/about/">About</a>
      <a href="https://github.com/TheVicky1/GitVerse.git" rel="noopener">GitHub</a>
    </nav>
    <a href="/check/" class="nav-cta">Check Profile</a>
    <button type="button" class="menu-btn" data-menu aria-expanded="false" aria-controls="site-nav">Menu</button>
  </div>
</header>
<main>

<section class="result" data-report data-scope="public" data-earned="${earnedEarnable}" data-user="${username}">

  <!-- Profile header -->
  <div class="result-hero">
    <img class="result-avatar" src="https://github.com/${username}.png?size=160" width="80" height="80" alt="@${username} GitHub avatar">
    <div class="result-hero-info">
      <div class="hero-eyebrow">GitHub Profile</div>
      <h1 class="result-user">@${username}</h1>
      <p class="result-sum">${summaryText}</p>
      ${nextUnlockText ? `<p class="result-next">Next unlock — <strong>${nextUnlockText}</strong></p>` : ''}
    </div>
  </div>

  <p class="result-scope">Showing public data only. <a href="/auth/login" data-signin>Sign in with GitHub</a> to include private repository progress.</p>
  ${isHidden ? `<div class="notice" style="margin-bottom: 32px; padding: 14px 20px; border: 1px solid var(--line); border-radius: var(--radius-md); background: var(--surface); color: var(--ink-2); font-size: 0.9rem;">This profile has no public achievements on display or achievements are hidden in GitHub profile settings.</div>` : ''}

  <!-- On display -->
  <div class="result-section">
    <div class="result-section-label">On display</div>
    ${displayItems.length > 0 ? `
      <p class="result-section-sub">Badges you've earned, with progress toward the next tier.</p>
      <div class="result-shelf">
        ${displayItems.map(renderBadgeCard).join('')}
      </div>
    ` : `
      <p class="result-section-sub">Nothing here yet — start working toward the badges below.</p>
    `}
  </div>

  <!-- Still to collect -->
  ${collectItems.length > 0 ? `
  <div class="result-section">
    <div class="result-section-label">Still to collect</div>
    <p class="result-section-sub">Badges you can still earn. Tap any to see exactly how.</p>
    <div class="result-shelf">
      ${collectItems.map(renderBadgeCard).join('')}
    </div>
  </div>
  ` : ''}

  <!-- Archive -->
  ${archiveItems.length > 0 ? `
  <details class="result-archive">
    <summary>Archive <span class="result-archive-hint">— ${archiveItems.length} badges you can no longer earn (switched off, retired or staff only)</span></summary>
    <div class="result-shelf result-shelf--archive">
      ${archiveItems.map(renderBadgeCard).join('')}
    </div>
  </details>
  ` : ''}

  <!-- Highlights -->
  <div class="result-section">
    <div class="result-section-label">Profile highlights</div>
    <ul class="highlights">
      ${highlights.map(h => `
        <li class="${h.earned ? 'on' : 'off'}">
          <span class="${h.earned ? 'highlight-check' : 'highlight-dot'}">${h.earned ? '✓' : '·'}</span> ${h.name}
        </li>
      `).join('')}
    </ul>
  </div>

  <!-- Actions -->
  <div class="result-actions">
    <a class="star-banner" href="https://github.com/TheVicky1/GitVerse.git" rel="noopener" data-star="result">
      <strong>If GitVerse saved you time, a star on GitHub goes a long way</strong>
      <span class="star-cta"><span aria-hidden="true">★</span> Star <span class="star-count" data-stars>0</span></span>
    </a>
    <div class="actions">
      <button type="button" class="btn" data-share onclick="navigator.clipboard.writeText(location.href); this.textContent='Copied!'; setTimeout(()=>this.textContent='Copy share link',2000)">Copy share link</button>
      <a href="/check/" class="btn primary">Check another profile</a>
    </div>
    <p class="meta result-meta">
      <strong>official</strong> = read from GitHub's own pages ·
      <strong>calibrated</strong> = matched against named milestones ·
      <strong>counted</strong> = computed from activity.
      Want private counts? <code>gh extension install nmloc/GitVerse</code>
      <button type="button" class="btn small" onclick="navigator.clipboard.writeText('gh extension install nmloc/GitVerse'); this.textContent='Copied!'; setTimeout(()=>this.textContent='Copy',2000)">Copy</button>
    </p>
  </div>

</section>

</main>
<footer class="aurora-footer">
  <div class="aurora-footer-top">
    <a href="/privacy/" class="footer-privacy-link">Privacy Policy</a>
    <span class="footer-copyright">© 2026 GITVERSE · GITHUB ACHIEVEMENT INTELLIGENCE</span>
  </div>
  <div class="aurora-footer-watermark" aria-hidden="true">
    <span>GITVERSE</span>
  </div>
</footer>
</body>
</html>`;
}

function renderErrorHtml(status, title, message) {
  return `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>${title} · GitVerse</title>
<meta name="color-scheme" content="dark">
<link rel="icon" type="image/png" href="/gitverse-favicon.png">
<script src="/theme.js"></script>
<link rel="stylesheet" href="/style.css">
</head>
<body>
<header class="top">
  <div class="top-in">
    <a class="brand" href="/"><img class="brand-wordmark" src="/gitverse-wordmark-transparent.png" alt="GitVerse"></a>
    <nav id="site-nav" aria-label="Site">
      <a href="/check/">Track</a>
      <a href="/achievements/">Explore</a>
      <a href="/about/">About</a>
      <a href="https://github.com/TheVicky1/GitVerse.git" rel="noopener">GitHub</a>
    </nav>
    <a href="/check/" class="nav-cta">Check Profile</a>
  </div>
</header>
<main>
  <div class="check-hero" style="padding: clamp(60px, 12vw, 120px) 0; text-align: center;">
    <div class="hero-eyebrow">${status === 404 ? 'Not Found' : 'Error'}</div>
    <h1 style="font-size: clamp(2rem, 5vw, 3.2rem); font-weight: 800; margin: 16px 0;">${title}</h1>
    <p class="lede" style="color: var(--ink-2); max-width: 520px; margin: 0 auto 36px;">${message}</p>
    <div style="display: flex; justify-content: center; gap: 12px;">
      <a href="/check/" class="btn primary">Check another profile</a>
      <a href="/" class="btn">Return home</a>
    </div>
  </div>
</main>
<footer class="aurora-footer">
  <div class="aurora-footer-top">
    <a href="/privacy/" class="footer-privacy-link">Privacy Policy</a>
    <span class="footer-copyright">© 2026 GITVERSE · GITHUB ACHIEVEMENT INTELLIGENCE</span>
  </div>
  <div class="aurora-footer-watermark" aria-hidden="true">
    <span>GITVERSE</span>
  </div>
</footer>
</body>
</html>`;
}

http.createServer(async (request, response) => {
  const url = new URL(request.url, `http://localhost:${port}`);
  const requestPath = decodeURIComponent(url.pathname);

  // Result view for profile check
  if (requestPath === "/result") {
    const user = (url.searchParams.get("u") || "").trim();
    if (!user) {
      response.writeHead(302, { Location: "/check/" });
      return response.end();
    }

    try {
      const data = await fetchUserProfile(user);
      if (data.error) {
        const status = data.error === 'not_found' ? 404 : 400;
        response.writeHead(status, { "Content-Type": "text/html; charset=utf-8" });
        return response.end(renderErrorHtml(status, data.error === 'not_found' ? `@${user} not found` : 'Invalid request', data.message));
      }

      const html = renderResultHtml(data);
      response.writeHead(200, { "Content-Type": "text/html; charset=utf-8" });
      return response.end(html);
    } catch (err) {
      console.error("Error serving /result:", err);
      response.writeHead(502, { "Content-Type": "text/html; charset=utf-8" });
      return response.end(renderErrorHtml(502, "GitHub Service Error", "GitHub did not answer or rate-limited the request. Please try again shortly."));
    }
  }

  let file = path.resolve(root, "." + (requestPath === "/" ? "/index.html" : requestPath));
  if (!file.startsWith(root)) return response.writeHead(403).end("Forbidden");

  if (fs.existsSync(file) && fs.statSync(file).isDirectory()) {
    file = path.join(file, "index.html");
  } else if (!fs.existsSync(file) && fs.existsSync(file + ".html")) {
    file = file + ".html";
  }

  fs.readFile(file, (error, data) => {
    if (error) {
      const notFound = path.join(root, "404.html");
      if (fs.existsSync(notFound)) {
        response.writeHead(404, { "Content-Type": "text/html; charset=utf-8" });
        return fs.createReadStream(notFound).pipe(response);
      }
      return response.writeHead(404).end("Not found");
    }
    response.writeHead(200, { "Content-Type": types[path.extname(file)] || "application/octet-stream" });
    response.end(data);
  });
}).listen(port, "127.0.0.1", () => console.log(`GitVerse preview: http://localhost:${port}`));
