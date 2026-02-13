package gen

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/gin-gonic/gin"
)

func GenLP(companyInfo, templateDir string) ([]byte, error) {
	prompt := fmt.Sprintf(lpPromptTemplate, companyInfo)

	// cmd := exec.Command("claude", "-p", "--model", "claude-opus-4-6", string(prompt))
	cmd := exec.Command("opencode", "run", "--agent", "bot", string(prompt))
	cmd.Dir = templateDir
	output, err := cmd.Output()
	return output, err
}

func GenCompanyInfo(ctx *gin.Context, htmlContent string) (string, error) {
	g := genkit.Init(ctx,
		genkit.WithPlugins(&openai.OpenAI{
			APIKey: os.Getenv("OPENAI_API_KEY"),
		}),
		genkit.WithDefaultModel("openai/gpt-5-nano"),
	)

	resp, err := genkit.Generate(ctx, g,
		ai.WithSystem(companyPromptTemplate),
		ai.WithPrompt(htmlContent),
	)
	return resp.Text(), err
}

const companyPromptTemplate = `
Analyze the following Instagram page HTML content and extract a comprehensive company/brand description.

Return ONLY a concise, well-structured company description in plain text (max 3 paragraphs). No markdown. No backticks. No labels. No prefixes like "Company Description:".

Extract and synthesize from the HTML:
- Brand/company name
- Industry/niche
- Products or services offered
- Brand tone and personality (infer from bio, captions, hashtags)
- Target audience (infer from content style and language)
- Value proposition or unique selling points
- Location if available

If certain details are not present in the HTML, infer what you can from available context and omit what cannot be reasonably determined.

Write the description in a professional, third-person tone suitable for use as input context when generating a landing page for this brand.

Instagram page HTML content:
%s
`

const lpPromptTemplate = `
Go to @app/page.tsx and create a Landing Page for the Company.
Don't read or edit another files, all the new content needs to be on page.tsx

Tech stack:
- Next.js page.tsx with "use client" directive
- Tailwind CSS only (no external dependencies, no component libraries)
- TypeScript strict

---

## Design Philosophy

You are a world-class UI designer, not a template engine. Every landing page you create must feel like it was hand-crafted by a top design agency — NOT like a generic SaaS template.

**CRITICAL: Avoid the "AI Landing Page" look.** Do NOT fall into these patterns:
- Generic hero → features grid → stats → CTA footer (the same skeleton every time)
- Perfectly symmetrical 3-column feature cards with icons
- "Get Started" / "Learn More" as the only CTA text
- Gradient blobs as the only decorative element
- The same blue/purple/indigo color palette every time

Instead, approach each landing page as a **unique design challenge**. Before writing code, decide on:

1. **A visual concept/theme** that matches the company's personality. Examples:
   - A fintech company → sharp, editorial, number-heavy with a dark theme and accent green
   - A wellness brand → organic shapes, warm earth tones, flowing layouts with lots of breathing room
   - A developer tool → terminal-inspired aesthetics, monospace type, dark mode with neon accents
   - A luxury product → extreme minimalism, serif typography, black-and-white with one gold accent
   - A creative agency → bold asymmetric layouts, oversized typography, vivid clashing colors
   - A food/restaurant → rich photography placeholders, warm tones, textured backgrounds

2. **A unique layout structure.** Mix and match from these patterns — never use the same combination twice:
   - Split-screen hero (text left, visual right — or inverted)
   - Full-bleed hero with overlapping elements
   - Asymmetric grid with varied card sizes
   - Horizontal scrolling section (overflow-x)
   - Bento grid layout (mixed-size tiles like Apple's marketing pages)
   - Staggered/masonry-style content blocks
   - Large statement text interspersed with small detail blocks
   - Sticky/parallax-inspired scroll sections
   - Editorial/magazine-style layout with pull quotes and sidebars
   - Single long-scroll narrative with alternating left-right content
   - Card stack or overlapping panels
   - Timeline or step-based visual flow

3. **A distinct typographic personality.** Vary these across projects:
   - Oversized display headings (clamp-based fluid type, 4xl to 8xl)
   - Mix of serif-feel (tracking-tight, font-light) vs sans-serif bold
   - All-caps micro-labels with wide letter-spacing for section headers
   - Large pull-quotes or statement text between sections
   - Varying text sizes within the same section for visual rhythm

4. **A unique color strategy.** Rotate between approaches:
   - Dark mode with a single vivid accent color
   - Light, airy palette with pastel accents
   - Monochromatic (shades of one color) with a complementary pop
   - High contrast black/white with one bold brand color
   - Earth tones and warm neutrals
   - Bold, saturated multi-color palette
   - Muted/desaturated tones with vibrant CTAs

5. **Distinctive decorative elements.** Choose 1-2, never all:
   - CSS-drawn geometric shapes (circles, triangles, lines) as accents
   - Dot grids or line patterns as subtle backgrounds
   - Diagonal or angled section dividers (clip-path or transforms)
   - Glassmorphism cards (backdrop-blur + transparency)
   - Neumorphism-inspired subtle depth
   - Border-based decorations (partial borders, corner accents)
   - Animated gradient meshes or shifting backgrounds
   - Noise/grain texture overlays (CSS or SVG filter)
   - Outlined/stroke text mixed with filled text
   - Large watermark-style background numbers or letters

---

## Sections & Content Strategy

Include 4-7 sections, but **vary their order, emphasis, and presentation**. Not every page needs every section. Choose what fits the brand:

- **Hero**: Always present, but vary the approach (minimal single-line, bold full-screen, split layout, video-placeholder hero, text-only editorial, etc.)
- **Value proposition / Features**: Can be a grid, a bento layout, alternating rows, a single large showcase, icon-less text blocks, or interactive tabs
- **Social proof**: Can be logos, testimonial cards, a single large quote, animated counters, case study snippets, or press mentions
- **How it works**: Numbered steps, timeline, animated flow, or skip entirely if it doesn't fit
- **Pricing or plans**: Only if relevant to the company type
- **CTA section**: Vary between minimal one-liner, full-width bold banner, or integrated into the last content section
- **Footer**: Keep it functional but styled to match the page personality

---

## Animation & Interaction Requirements

- Use CSS @keyframes and Tailwind's animate utilities for entrance animations
- Implement scroll-triggered animations using IntersectionObserver (create a simple useEffect hook)
- Add meaningful hover states that go beyond opacity changes (scale, translate, color shifts, border reveals, shadow changes)
- Include at least one subtle continuous animation (floating element, gradient shift, pulsing glow)
- Stagger animation delays on grouped elements for a polished sequential reveal
- All animations should be smooth (use transform and opacity for performance)
- Respect prefers-reduced-motion

---

## Technical Quality

- Semantic HTML5 elements (header, main, section, footer, nav)
- All interactive elements must have focus-visible styles
- Color contrast must meet WCAG AA minimum
- Responsive breakpoints: mobile-first, sm, md, lg, xl
- Use CSS custom properties (via Tailwind arbitrary values) for the color palette so it feels cohesive
- No placeholder images — use CSS gradients, patterns, or geometric shapes as visual stand-ins
- No Lorem Ipsum — write realistic, compelling copy that matches the company's voice

---

## Do NOT:
- Import or use any external package, library, or icon set
- Create the same layout structure you used in previous generations
- Use generic tech-startup copy ("Revolutionize", "Supercharge", "Seamless")
- Default to the same blue/purple gradient palette
- Make every section a centered container with symmetric columns
- Use placeholder text or icons like "Feature 1", "Feature 2"

---

## Content/context for the landing page:
%s
	`
