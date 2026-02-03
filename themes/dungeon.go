package themes

import (
	"fmt"
	"strings"
)

// DungeonTheme 地牢火把風格
type DungeonTheme struct{}

func init() {
	RegisterTheme(&DungeonTheme{})
}

func (t *DungeonTheme) Name() string {
	return "dungeon"
}

func (t *DungeonTheme) Description() string {
	return "地牢：石牆火把照明，黑暗冒險氛圍"
}

const (
	DunStone     = "\033[38;2;105;105;105m"
	DunDarkStone = "\033[38;2;64;64;64m"
	DunTorch     = "\033[38;2;255;147;41m"
	DunFlame     = "\033[38;2;255;100;0m"
	DunGold      = "\033[38;2;255;215;0m"
	DunRed       = "\033[38;2;178;34;34m"
	DunGreen     = "\033[38;2;34;139;34m"
	DunBlue      = "\033[38;2;70;130;180m"
	DunPurple    = "\033[38;2;138;43;226m"
	DunBone      = "\033[38;2;255;250;240m"
	DunShadow    = "\033[38;2;25;25;25m"
	DunMoss      = "\033[38;2;85;107;47m"
)

func (t *DungeonTheme) Render(data StatusData) string {
	var sb strings.Builder

	// Fixed width: 80 characters
	width := 80

	// Stone wall top with torches (pure ASCII)
	// ###(*)###...###(*)###
	torchDecor := DunDarkStone + "###" + DunTorch + "(" + DunFlame + "*" + DunTorch + ")" + DunDarkStone + "###" + Reset
	torchLen := 9 // ###(*)###
	middleLen := width - torchLen*2
	sb.WriteString(torchDecor + DunDarkStone + strings.Repeat("#", middleLen) + Reset + torchDecor + "\n")

	// Chamber name
	modelColor, _ := GetModelConfig(data.ModelType)
	chamberName := "The Dark Chamber"
	if data.ModelType == "Opus" {
		chamberName = "The Arcane Sanctum"
	} else if data.ModelType == "Haiku" {
		chamberName = "The Monk's Cell"
	}

	update := ""
	if data.UpdateAvailable {
		update = DunGold + " *" + Reset
	}

	line1 := fmt.Sprintf("%s#%s %s%s%s  %s~ %s ~%s%s  %s%s",
		DunDarkStone, Reset,
		modelColor, data.ModelName, Reset,
		DunTorch, chamberName, Reset, update,
		DunStone, data.Version)
	sb.WriteString(dunPadLine(line1, width, DunDarkStone+"#"+Reset))

	// Quest scroll
	gitStr := ""
	if data.GitBranch != "" {
		gitStr = fmt.Sprintf("  %s<%s>%s", DunMoss, data.GitBranch, Reset)
		if data.GitStaged > 0 {
			gitStr += fmt.Sprintf(" %s+%d%s", DunGreen, data.GitStaged, Reset)
		}
		if data.GitDirty > 0 {
			gitStr += fmt.Sprintf(" %s*%d%s", DunRed, data.GitDirty, Reset)
		}
	}

	line2 := fmt.Sprintf("%s#%s %sScroll:%s %s%s",
		DunDarkStone, Reset,
		DunBone, Reset, ShortenPath(data.ProjectPath, 30), gitStr)
	sb.WriteString(dunPadLine(line2, width, DunDarkStone+"#"+Reset))

	// Stone separator
	sb.WriteString(DunDarkStone + "#" + DunStone + strings.Repeat("=", width-2) + DunDarkStone + "#" + Reset + "\n")

	// Stats as dungeon items
	line3 := fmt.Sprintf("%s#%s %sSwd%s %-6s %sShd%s %-3d %sTime%s %-6s %sSkul%s %-6s %sGem%s %s",
		DunDarkStone, Reset,
		DunRed, Reset, FormatTokens(data.TokenCount),
		DunBlue, Reset, data.MessageCount,
		DunStone, Reset, data.SessionTime,
		DunPurple, Reset, FormatCostShort(data.BurnRate),
		DunGold, Reset, FormatCostShort(data.DayCost))
	sb.WriteString(dunPadLine(line3, width, DunDarkStone+"#"+Reset))

	// Health/Mana pools
	hp := 100 - data.ContextPercent
	hpColor := DunGreen
	if hp <= 20 {
		hpColor = DunRed
	} else if hp <= 50 {
		hpColor = DunTorch
	}

	hpBar := t.generateDungeonBar(hp, 10, hpColor)
	mpBar := t.generateDungeonBar(100-data.API5hrPercent, 8, DunBlue)
	xpBar := t.generateDungeonBar(100-data.API7dayPercent, 8, DunPurple)

	line4 := fmt.Sprintf("%s#%s %sHP%s%s%s%3d%s %sMP%s%s%s%3d%s %sXP%s%s%s%3d%s",
		DunDarkStone, Reset,
		DunRed, Reset, hpBar, hpColor, hp, Reset,
		DunBlue, Reset, mpBar, DunBlue, 100-data.API5hrPercent, Reset,
		DunPurple, Reset, xpBar, DunPurple, 100-data.API7dayPercent, Reset)
	sb.WriteString(dunPadLine(line4, width, DunDarkStone+"#"+Reset))

	// Treasure info
	line5 := fmt.Sprintf("%s#%s %sGold%s %s ses  %sPotn%s %d%% hit  %sLeft%s %s / %s",
		DunDarkStone, Reset,
		DunGold, Reset, FormatCostShort(data.SessionCost),
		DunGreen, Reset, data.CacheHitRate,
		DunStone, Reset, data.API5hrTimeLeft, data.API7dayTimeLeft)
	sb.WriteString(dunPadLine(line5, width, DunDarkStone+"#"+Reset))

	// Stone wall bottom with torches
	sb.WriteString(torchDecor + DunDarkStone + strings.Repeat("#", middleLen) + Reset + torchDecor + "\n")

	return sb.String()
}

func dunPadLine(line string, targetWidth int, suffix string) string {
	visible := dunVisibleLen(line)
	suffixLen := dunVisibleLen(suffix)
	padding := targetWidth - visible - suffixLen
	if padding < 0 {
		padding = 0
	}
	return line + strings.Repeat(" ", padding) + suffix + "\n"
}

func dunVisibleLen(s string) int {
	inEscape := false
	count := 0
	for _, r := range s {
		if r == '\033' {
			inEscape = true
		} else if inEscape {
			if r == 'm' {
				inEscape = false
			}
		} else {
			count++
		}
	}
	return count
}

func (t *DungeonTheme) generateDungeonBar(percent, width int, color string) string {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	filled := percent * width / 100
	empty := width - filled

	var bar strings.Builder
	bar.WriteString(DunShadow + "[" + Reset)
	if filled > 0 {
		bar.WriteString(color)
		bar.WriteString(strings.Repeat("#", filled))
		bar.WriteString(Reset)
	}
	if empty > 0 {
		bar.WriteString(DunShadow)
		bar.WriteString(strings.Repeat("-", empty))
		bar.WriteString(Reset)
	}
	bar.WriteString(DunShadow + "]" + Reset)
	return bar.String()
}
