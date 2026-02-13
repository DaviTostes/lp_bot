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

const lpPromptTemplate = `
Return ONLY raw TSX code. No markdown. No backticks. No explanations. No comments.

Tkch stack:
- Next.js page.tsx with "use client" directive
- Tailwind CSS only (no external dependencies, no component libraries)
- TypeScript strict

Design requirements:
- Modern, visually striking layout with generous whitespace and clear visual hierarchy
- Smooth CSS animations and transitions (hover effects, fade-ins, subtle micro-interactions)
- Fully responsive: mobile-first, adapts seamlessly to tablet and desktop
- Use CSS gradients, shadows, and backdrop-blur for depth and polish
- Professional typography: varied font sizes, weights, and letter-spacing for contrast
- Consistent color palette with strong primary CTA colors and neutral backgrounds
- Sections: hero with compelling headline, features/benefits grid, social proof/stats, and a clear call-to-action
- Accessible: semantic HTML, sufficient color contrast, focus-visible states
- Use Tailwind arbitrary values when needed for fine-tuned spacing or colors

Do NOT:
- Import or use any external package
- Add code comments
- Read or edit any files
- Output anything other than the TSX code itself

Content/context for the landing page:
%s
	`

func GenLP(companyInfo string) ([]byte, error) {
	prompt := fmt.Sprintf(lpPromptTemplate, companyInfo)

	cmd := exec.Command("claude", "-p", "--model", "claude-opus-4-6", string(prompt))
	output, err := cmd.Output()
	return output, err
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
