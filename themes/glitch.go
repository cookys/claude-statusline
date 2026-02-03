package themes

import (
	"fmt"
	"strings"
)

// GlitchTheme 數位故障風格
type GlitchTheme struct{}

func init() {
	RegisterTheme(&GlitchTheme{})
}

func (t *GlitchTheme) Name() string {
	return "glitch"
}

func (t *GlitchTheme) Description() string {
	return "故障風：數位錯位，賽博龐克破碎美學"
}

const (
	GlitchRed    = "\033[38;2;255;0;60m"
	GlitchCyan   = "\033[38;2;0;255;240m"
	GlitchWhite  = "\033[38;2;255;255;255m"
	GlitchGray   = "\033[38;2;80;80;80m"
	GlitchDim    = "\033[38;2;50;50;50m"
	GlitchPink   = "\033[38;2;255;60;150m"
	GlitchBlue   = "\033[38;2;30;144;255m"
	GlitchBg     = "\033[48;2;10;10;15m"
)

func (t *GlitchTheme) Render(data StatusData) string {
	var sb strings.Builder

	// Glitch top border (broken/offset)
	sb.WriteString(GlitchDim + "▓▒░" + GlitchRed + "█" + GlitchDim + "░▒▓")
	sb.WriteString(GlitchGray + strings.Repeat("▀", 30))
	sb.WriteString(GlitchCyan + "█")
	sb.WriteString(GlitchGray + strings.Repeat("▀", 30))
	sb.WriteString(GlitchDim + "▓▒░" + GlitchRed + "█" + GlitchDim + "░▒▓")
	sb.WriteString(Reset + "\n")

	// Model info with glitch offset effect
	modelColor, modelIcon := GetModelConfig(data.ModelType)
	update := ""
	if data.UpdateAvailable {
		update = fmt.Sprintf(" %s[!]%s", GlitchRed, Reset)
	}

	// Chromatic aberration style (red/cyan offset)
	line1 := fmt.Sprintf(" %s▌%s%s%s%s%s %s%s%s%s  %s▐%s%s▌%s  %s%s%s",
		GlitchRed, Reset,
		modelColor, Bold, modelIcon, data.ModelName, Reset,
		GlitchGray, data.Version, Reset, update,
		GlitchCyan, GlitchRed, Reset,
		GlitchWhite, ShortenPath(data.ProjectPath, 22), Reset)
	if data.GitBranch != "" {
		line1 += fmt.Sprintf("  %s◄%s%s►%s", GlitchCyan, data.GitBranch, GlitchRed, Reset)
		if data.GitStaged > 0 {
			line1 += fmt.Sprintf(" %s+%d%s", GlitchCyan, data.GitStaged, Reset)
		}
		if data.GitDirty > 0 {
			line1 += fmt.Sprintf(" %s~%d%s", GlitchRed, data.GitDirty, Reset)
		}
	}
	sb.WriteString(line1)
	sb.WriteString("\n")

	// Stats with corrupted look
	line2 := fmt.Sprintf(" %s▌%s %s%s%s  %s%d%s msg  %s%s%s  %s▐%s%s▌%s  %s%s%s  %s%s%s  %s%s/h%s  %s%d%%hit%s",
		GlitchRed, Reset,
		GlitchPink, FormatTokens(data.TokenCount), Reset,
		GlitchCyan, data.MessageCount, Reset,
		GlitchGray, data.SessionTime, Reset,
		GlitchCyan, GlitchRed, GlitchCyan, Reset,
		GlitchCyan, FormatCostShort(data.SessionCost), Reset,
		GlitchWhite, FormatCostShort(data.DayCost), Reset,
		GlitchRed, FormatCostShort(data.BurnRate), Reset,
		GlitchCyan, data.CacheHitRate, Reset)
	sb.WriteString(line2)
	sb.WriteString("\n")

	// Glitchy progress bars
	ctxBar := t.generateGlitchBar(data.ContextPercent, 14)
	bar5 := t.generateGlitchBar(data.API5hrPercent, 10)
	bar7 := t.generateGlitchBar(data.API7dayPercent, 10)

	ctxColor := GlitchCyan
	if data.ContextPercent >= 80 {
		ctxColor = GlitchRed
	} else if data.ContextPercent >= 60 {
		ctxColor = GlitchPink
	}

	line3 := fmt.Sprintf(" %s▌%s %sCTX%s%s%s%3d%%%s  %s5HR%s%s%s%3d%%%s %s%s%s  %s7DY%s%s%s%3d%%%s %s%s%s",
		GlitchRed, Reset,
		GlitchGray, Reset, ctxBar, ctxColor, data.ContextPercent, Reset,
		GlitchGray, Reset, bar5, GlitchCyan, data.API5hrPercent, Reset,
		GlitchGray, data.API5hrTimeLeft, Reset,
		GlitchGray, Reset, bar7, GlitchPink, data.API7dayPercent, Reset,
		GlitchGray, data.API7dayTimeLeft, Reset)
	sb.WriteString(line3)
	sb.WriteString("\n")

	// Glitch bottom border
	sb.WriteString(GlitchDim + "▓▒░" + GlitchCyan + "█" + GlitchDim + "░▒▓")
	sb.WriteString(GlitchGray + strings.Repeat("▄", 30))
	sb.WriteString(GlitchRed + "█")
	sb.WriteString(GlitchGray + strings.Repeat("▄", 30))
	sb.WriteString(GlitchDim + "▓▒░" + GlitchCyan + "█" + GlitchDim + "░▒▓")
	sb.WriteString(Reset + "\n")

	return sb.String()
}

func (t *GlitchTheme) generateGlitchBar(percent, width int) string {
	filled := percent * width / 100
	if filled > width {
		filled = width
	}
	empty := width - filled

	var bar strings.Builder
	bar.WriteString(GlitchDim + "【" + Reset)

	// Glitchy filled part with occasional "corruption"
	if filled > 0 {
		for i := 0; i < filled; i++ {
			if i == filled/2 && filled > 3 {
				// Add a glitch in the middle
				bar.WriteString(GlitchRed + "█" + Reset)
			} else {
				bar.WriteString(GlitchCyan + "▓" + Reset)
			}
		}
	}
	if empty > 0 {
		bar.WriteString(GlitchDim)
		bar.WriteString(strings.Repeat("░", empty))
		bar.WriteString(Reset)
	}
	bar.WriteString(GlitchDim + "】" + Reset)
	return bar.String()
}
