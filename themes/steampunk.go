package themes

import (
	"fmt"
	"strings"
)

// SteampunkTheme 蒸汽龐克風格
type SteampunkTheme struct{}

func init() {
	RegisterTheme(&SteampunkTheme{})
}

func (t *SteampunkTheme) Name() string {
	return "steampunk"
}

func (t *SteampunkTheme) Description() string {
	return "蒸汽龐克：維多利亞黃銅齒輪，工業美學"
}

const (
	SteamBrass   = "\033[38;2;205;165;85m"
	SteamCopper  = "\033[38;2;184;115;51m"
	SteamBronze  = "\033[38;2;150;116;68m"
	SteamGold    = "\033[38;2;255;215;0m"
	SteamRust    = "\033[38;2;140;80;60m"
	SteamIvory   = "\033[38;2;255;255;240m"
	SteamDark    = "\033[38;2;60;50;40m"
	SteamGear    = "\033[38;2;120;100;80m"
	SteamGreen   = "\033[38;2;100;140;100m"
	SteamRed     = "\033[38;2;180;80;60m"
	SteamBgBrass = "\033[48;2;40;35;25m"
)

func (t *SteampunkTheme) Render(data StatusData) string {
	var sb strings.Builder

	// Top border with gear decoration
	sb.WriteString(SteamDark + "╔" + SteamBrass + "⚙" + SteamDark + "═══════════════════════════════════════════════════════════════════════════════" + SteamBrass + "⚙" + SteamDark + "╗" + Reset)
	sb.WriteString("\n")

	// Model + Version + Path
	modelColor, modelIcon := GetModelConfig(data.ModelType)
	update := ""
	if data.UpdateAvailable {
		update = fmt.Sprintf(" %s⚡%s", SteamGold, Reset)
	}

	line1 := fmt.Sprintf("%s║%s %s⚙%s %s%s%s%s%s %s%s%s%s  %s│%s  %s⚙%s %s%s%s",
		SteamDark, Reset,
		SteamBrass, Reset,
		modelColor, Bold, modelIcon, data.ModelName, Reset,
		SteamGear, data.Version, Reset, update,
		SteamDark, Reset,
		SteamCopper, Reset,
		SteamIvory, ShortenPath(data.ProjectPath, 20), Reset)
	if data.GitBranch != "" {
		line1 += fmt.Sprintf("  %s◈%s%s", SteamBronze, data.GitBranch, Reset)
		if data.GitStaged > 0 {
			line1 += fmt.Sprintf(" %s+%d%s", SteamGreen, data.GitStaged, Reset)
		}
		if data.GitDirty > 0 {
			line1 += fmt.Sprintf(" %s~%d%s", SteamRust, data.GitDirty, Reset)
		}
	}
	// Pad and close
	sb.WriteString(PadRight(line1, 85))
	sb.WriteString(SteamDark + "║" + Reset)
	sb.WriteString("\n")

	// Separator with pipes
	sb.WriteString(SteamDark + "╠" + SteamGear + "═══════════════════════════════════" + SteamBrass + "◈" + SteamGear + "═══════════════════════════════════════════" + SteamDark + "╣" + Reset)
	sb.WriteString("\n")

	// Stats with gauge style
	line2 := fmt.Sprintf("%s║%s %s⊚%s%s tok  %s⊛%s%d msg  %s⊙%s%s  %s│%s  %s%s%s  %s%s%s  %s%s/h%s  %s%d%%hit%s",
		SteamDark, Reset,
		SteamGear, SteamBrass, FormatTokens(data.TokenCount),
		SteamGear, SteamCopper, data.MessageCount,
		SteamGear, SteamBronze, data.SessionTime,
		SteamDark, Reset,
		SteamGreen, FormatCostShort(data.SessionCost), Reset,
		SteamGold, FormatCostShort(data.DayCost), Reset,
		SteamRed, FormatCostShort(data.BurnRate), Reset,
		SteamGreen, data.CacheHitRate, Reset)
	sb.WriteString(PadRight(line2, 85))
	sb.WriteString(SteamDark + "║" + Reset)
	sb.WriteString("\n")

	// Pressure gauges (progress bars)
	ctxBar := t.generateGaugeBar(data.ContextPercent, 14)
	bar5 := t.generateGaugeBar(data.API5hrPercent, 10)
	bar7 := t.generateGaugeBar(data.API7dayPercent, 10)

	ctxColor := SteamGreen
	if data.ContextPercent >= 80 {
		ctxColor = SteamRed
	} else if data.ContextPercent >= 60 {
		ctxColor = SteamGold
	}

	line3 := fmt.Sprintf("%s║%s %sCTX%s%s%s%3d%%%s  %s5HR%s%s%s%3d%%%s %s%s%s  %s7DY%s%s%s%3d%%%s %s%s%s",
		SteamDark, Reset,
		SteamGear, Reset, ctxBar, ctxColor, data.ContextPercent, Reset,
		SteamGear, Reset, bar5, SteamBrass, data.API5hrPercent, Reset,
		SteamGear, data.API5hrTimeLeft, Reset,
		SteamGear, Reset, bar7, SteamCopper, data.API7dayPercent, Reset,
		SteamGear, data.API7dayTimeLeft, Reset)
	sb.WriteString(PadRight(line3, 85))
	sb.WriteString(SteamDark + "║" + Reset)
	sb.WriteString("\n")

	// Bottom border with gear decoration
	sb.WriteString(SteamDark + "╚" + SteamBrass + "⚙" + SteamDark + "═══════════════════════════════════════════════════════════════════════════════" + SteamBrass + "⚙" + SteamDark + "╝" + Reset)
	sb.WriteString("\n")

	return sb.String()
}

func (t *SteampunkTheme) generateGaugeBar(percent, width int) string {
	filled := percent * width / 100
	if filled > width {
		filled = width
	}
	empty := width - filled

	var bar strings.Builder
	bar.WriteString(SteamDark + "〔" + Reset)

	// Pressure gauge style
	if filled > 0 {
		bar.WriteString(SteamBgBrass)
		bar.WriteString(SteamBrass)
		bar.WriteString(strings.Repeat("▰", filled))
		bar.WriteString(Reset)
	}
	if empty > 0 {
		bar.WriteString(SteamDark)
		bar.WriteString(strings.Repeat("▱", empty))
		bar.WriteString(Reset)
	}
	bar.WriteString(SteamDark + "〕" + Reset)
	return bar.String()
}
