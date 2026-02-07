package themes

import (
	"strings"
	"testing"
)

func TestFormatTokens(t *testing.T) {
	tests := []struct {
		tokens   int64
		expected string
	}{
		{0, "0"},
		{500, "500"},
		{999, "999"},
		{1000, "1.0k"},
		{1500, "1.5k"},
		{10000, "10.0k"},
		{999999, "1000.0k"},
		{1000000, "1.0M"},
		{1500000, "1.5M"},
		{10000000, "10.0M"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := FormatTokens(tt.tokens)
			if result != tt.expected {
				t.Errorf("FormatTokens(%d) = %q, want %q", tt.tokens, result, tt.expected)
			}
		})
	}
}

func TestFormatCost(t *testing.T) {
	tests := []struct {
		cost     float64
		expected string
	}{
		{0, "$0.00"},
		{0.01, "$0.01"},
		{0.99, "$0.99"},
		{1.00, "$1.00"},
		{5.50, "$5.50"},
		{10.00, "$10.0"},
		{99.99, "$100.0"},
		{100.00, "$100"},
		{999.99, "$1000"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := FormatCost(tt.cost)
			if result != tt.expected {
				t.Errorf("FormatCost(%v) = %q, want %q", tt.cost, result, tt.expected)
			}
		})
	}
}

func TestVisibleWidth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty", "", 0},
		{"ascii", "hello", 5},
		{"with ansi", "\033[31mred\033[0m", 3},
		{"cjk", "中文", 4},
		{"mixed", "a中b", 4},
		{"complex ansi", "\033[38;2;100;100;100mtext\033[0m", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := VisibleWidth(tt.input)
			if result != tt.expected {
				t.Errorf("VisibleWidth(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRuneWidth(t *testing.T) {
	tests := []struct {
		name     string
		r        rune
		expected int
	}{
		{"ascii", 'a', 1},
		{"digit", '0', 1},
		{"cjk", '中', 2},
		{"emoji", '😀', 2},
		{"variation selector", '\uFE0F', 0},
		{"zero width", '\u200B', 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RuneWidth(tt.r)
			if result != tt.expected {
				t.Errorf("RuneWidth(%q) = %d, want %d", tt.r, result, tt.expected)
			}
		})
	}
}

func TestGetTheme(t *testing.T) {
	// Test that default theme exists
	theme, ok := GetTheme("classic_framed")
	if !ok {
		t.Error("classic_framed theme should exist")
	}
	if theme == nil {
		t.Error("theme should not be nil")
	}

	// Test non-existent theme
	_, ok = GetTheme("non_existent_theme_12345")
	if ok {
		t.Error("non-existent theme should return false")
	}
}

func TestListThemes(t *testing.T) {
	themes := ListThemes()
	if len(themes) == 0 {
		t.Error("ListThemes should return at least one theme")
	}

	// Verify each theme has required methods
	for _, theme := range themes {
		if theme.Name() == "" {
			t.Error("theme name should not be empty")
		}
		if theme.Description() == "" {
			t.Error("theme description should not be empty")
		}
	}
}

func TestGenerateBar(t *testing.T) {
	tests := []struct {
		name        string
		percent     int
		width       int
		filledChar  string
		emptyChar   string
		filledColor string
		emptyColor  string
	}{
		{"empty bar", 0, 10, "█", "░", "", ""},
		{"full bar", 100, 10, "█", "░", "", ""},
		{"half bar", 50, 10, "█", "░", "", ""},
		{"over 100", 150, 10, "█", "░", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateBar(tt.percent, tt.width, tt.filledChar, tt.emptyChar, tt.filledColor, tt.emptyColor)
			if result == "" {
				t.Error("GenerateBar should not return empty string")
			}
		})
	}
}

func TestPadFunctions(t *testing.T) {
	t.Run("PadLeft", func(t *testing.T) {
		result := PadLeft("test", 10)
		if len(result) != 10 {
			t.Errorf("PadLeft result length = %d, want 10", len(result))
		}
		if !strings.HasSuffix(result, "test") {
			t.Errorf("PadLeft should end with 'test': %q", result)
		}
	})

	t.Run("PadRight", func(t *testing.T) {
		result := PadRight("test", 10)
		if len(result) != 10 {
			t.Errorf("PadRight result length = %d, want 10", len(result))
		}
		if !strings.HasPrefix(result, "test") {
			t.Errorf("PadRight should start with 'test': %q", result)
		}
	})

	t.Run("PadCenter", func(t *testing.T) {
		result := PadCenter("test", 10)
		if len(result) != 10 {
			t.Errorf("PadCenter result length = %d, want 10", len(result))
		}
		if !strings.Contains(result, "test") {
			t.Errorf("PadCenter should contain 'test': %q", result)
		}
	})
}

func TestGetModelConfig(t *testing.T) {
	tests := []struct {
		modelType     string
		expectedColor string
	}{
		{"Opus", ColorGold},
		{"Sonnet", ColorCyan},
		{"Haiku", ColorPink},
		{"Unknown", ColorCyan},
	}

	for _, tt := range tests {
		t.Run(tt.modelType, func(t *testing.T) {
			color, icon := GetModelConfig(tt.modelType)
			if color != tt.expectedColor {
				t.Errorf("GetModelConfig(%q) color = %q, want %q", tt.modelType, color, tt.expectedColor)
			}
			if icon == "" {
				t.Errorf("GetModelConfig(%q) icon should not be empty", tt.modelType)
			}
		})
	}
}

func TestGetBarColor(t *testing.T) {
	tests := []struct {
		percent       int
		expectedColor string
	}{
		{0, ColorBrightGreen},
		{25, ColorBrightGreen},
		{49, ColorBrightGreen},
		{50, ColorBrightYellow},
		{74, ColorBrightYellow},
		{75, ColorRed},
		{100, ColorRed},
	}

	for _, tt := range tests {
		t.Run(string(rune(tt.percent)), func(t *testing.T) {
			color, _ := GetBarColor(tt.percent)
			if color != tt.expectedColor {
				t.Errorf("GetBarColor(%d) = %q, want %q", tt.percent, color, tt.expectedColor)
			}
		})
	}
}

func TestGetContextColor(t *testing.T) {
	tests := []struct {
		percent       int
		expectedColor string
	}{
		{0, ColorCtxGreen},
		{59, ColorCtxGreen},
		{60, ColorCtxGold},
		{79, ColorCtxGold},
		{80, ColorCtxRed},
		{100, ColorCtxRed},
	}

	for _, tt := range tests {
		t.Run(string(rune(tt.percent)), func(t *testing.T) {
			color := GetContextColor(tt.percent)
			if color != tt.expectedColor {
				t.Errorf("GetContextColor(%d) = %q, want %q", tt.percent, color, tt.expectedColor)
			}
		})
	}
}
