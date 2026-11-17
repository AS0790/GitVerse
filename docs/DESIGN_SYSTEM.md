# GitVerse Design System Specification

GitVerse features a bespoke **Space Aurora** minimalist dark theme designed for executive developer clarity, calm focus, and premium tactile interaction.

---

## 1. Palette & Surface Foundations

The design strictly eliminates bright blinding canvas surfaces in favor of deep OLED obsidian tones and subtle auroral contour lines:

| Token | Hex Value | Purpose |
| :--- | :--- | :--- |
| `--bg` | `#0a0a0b` | Primary OLED canvas background |
| `--surface` | `#121214` | High-elevation card surfaces |
| `--surface-2` | `#18181b` | Secondary card elements, nested shelves |
| `--surface-3` | `#27272a` | Hover states, button surfaces, borders |
| `--ink` | `#f4f4f5` | High-contrast primary typography |
| `--ink-2` | `#a1a1aa` | Muted secondary copy, labels, hints |
| `--ink-3` | `#71717a` | Tertiary captions, subtle placeholders |
| `--line` | `rgba(255, 255, 255, 0.08)` | Ultra-fine hairline border dividers |

---

## 2. Metallic Tier Accent System

Achievement tiers reflect authentic precious metal finishes:

| Tier | Metallic Accent | Hex Value | CSS Class |
| :---: | :--- | :--- | :--- |
| **x4** | 🥇 **Gold** | `#c9971f` | `.badge-card--gold` |
| **x3** | 🥈 **Silver** | `#e2e8f0` | `.badge-card--silver` |
| **x2** | 🥉 **Bronze** | `#d97706` | `.badge-card--bronze` |
| **x1** | 🔘 **Base** | `#71717a` | `.badge-card--base` |

---

## 3. Topographic Orbital Backdrop

Every page features an inline SVG contour backdrop representing mathematical orbital coordinates:

```css
body {
  background-color: var(--bg);
  background-image: url("data:image/svg+xml,...");
  background-attachment: fixed;
  background-size: cover;
}
```

---

## 4. Components & Micro-Interactions

### A. Badge Shelves (`.result-shelf-card`)
Badges are displayed on sleek, illuminated horizontal vitrine shelves with subtle specular edge highlights:
* Border: `1px solid rgba(255, 255, 255, 0.06)`
* Hover Transition: `transform: translateY(-2px)`, border illumination.

### B. Progress Meters (`.badge-progress`)
* Height: `4px`
* Background: `rgba(255, 255, 255, 0.08)`
* Fill: Solid warm gold (`#c9971f`) with a soft glow filter.

### C. Interactive Button Spinners (`.btn.is-loading`)
* Form submission triggers a smooth inline spinner (`.btn-spinner`).
* Locked minimum width prevents layout shift.
* `aria-busy="true"` guarantees screen reader accessibility.

