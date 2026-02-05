# Claude Statusline

[![CI](https://github.com/kevinlincg/claude-statusline/actions/workflows/ci.yml/badge.svg)](https://github.com/kevinlincg/claude-statusline/actions/workflows/ci.yml)
[![Release](https://github.com/kevinlincg/claude-statusline/actions/workflows/release.yml/badge.svg)](https://github.com/kevinlincg/claude-statusline/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinlincg/claude-statusline)](https://goreportcard.com/report/github.com/kevinlincg/claude-statusline)
[![Go Reference](https://pkg.go.dev/badge/github.com/kevinlincg/claude-statusline.svg)](https://pkg.go.dev/github.com/kevinlincg/claude-statusline)
[![GitHub release](https://img.shields.io/github/v/release/kevinlincg/claude-statusline)](https://github.com/kevinlincg/claude-statusline/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![SLSA 3](https://slsa.dev/images/gh-badge-level3.svg)](https://slsa.dev)

[English](README.md) | [繁體中文](README.zh-TW.md) | [简体中文](README.zh-CN.md) | [日本語](README.ja.md)

A custom status line for Claude Code written in Go. Displays model info, Git status, API usage, token consumption, cost metrics, and more.

<p align="center">
  <img src="assets/intro.gif" alt="Claude Statusline Theme Demo" width="750">
</p>

## Installation

### Option 1: Download Binary (Recommended)

Download the latest release for your platform from [GitHub Releases](https://github.com/kevinlincg/claude-statusline/releases/latest):

```bash
# macOS (Apple Silicon)
curl -L https://github.com/kevinlincg/claude-statusline/releases/latest/download/claude-statusline_darwin_arm64.tar.gz | tar xz
mv statusline ~/.claude/statusline-go/

# macOS (Intel)
curl -L https://github.com/kevinlincg/claude-statusline/releases/latest/download/claude-statusline_darwin_amd64.tar.gz | tar xz
mv statusline ~/.claude/statusline-go/

# Linux (x64)
curl -L https://github.com/kevinlincg/claude-statusline/releases/latest/download/claude-statusline_linux_amd64.tar.gz | tar xz
mv statusline ~/.claude/statusline-go/

# Linux (ARM64)
curl -L https://github.com/kevinlincg/claude-statusline/releases/latest/download/claude-statusline_linux_arm64.tar.gz | tar xz
mv statusline ~/.claude/statusline-go/
```

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/kevinlincg/claude-statusline.git ~/.claude/statusline-go

# Build
cd ~/.claude/statusline-go
go build -o statusline .
```

### Configure Claude Code

Add to `~/.claude/settings.json`:

```json
{
  "statusLine": {
    "type": "command",
    "command": "~/.claude/statusline-go/statusline"
  }
}
```

## Themes

### Interactive Theme Selector

Use the interactive menu to preview and select themes:

```bash
./statusline --menu
```

Use arrow keys (or h/l) to browse themes, Enter to confirm, q to cancel.

### Command Line Options

```bash
./statusline --list-themes      # List all available themes
./statusline --preview <name>   # Preview a specific theme
./statusline --set-theme <name> # Set theme directly
./statusline --menu             # Interactive theme selector
./statusline --version          # Show version information
```

### Manual Configuration

Edit `~/.claude/statusline-go/config.json`:

```json
{
  "theme": "classic_framed"
}
```

### Available Themes

#### Classic & Minimal

<details>
<summary><b>classic</b> - Original layout style</summary>
<img src="assets/themes/classic.png" alt="classic theme">
</details>

<details>
<summary><b>classic_framed</b> - Tree structure with frame, aligned progress bars (Default)</summary>
<img src="assets/themes/classic_framed.png" alt="classic_framed theme">
</details>

<details>
<summary><b>minimal</b> - Clean tree structure, no borders</summary>
<img src="assets/themes/minimal.png" alt="minimal theme">
</details>

<details>
<summary><b>compact</b> - Minimal height, complete info</summary>
<img src="assets/themes/compact.png" alt="compact theme">
</details>

<details>
<summary><b>boxed</b> - Full border frame, symmetrical sections</summary>
<img src="assets/themes/boxed.png" alt="boxed theme">
</details>

<details>
<summary><b>zen</b> - Minimalist whitespace, calm and elegant</summary>
<img src="assets/themes/zen.png" alt="zen theme">
</details>

#### Sci-Fi & Cyberpunk

<details>
<summary><b>hud</b> - Sci-fi HUD interface with angle bracket labels</summary>
<img src="assets/themes/hud.png" alt="hud theme">
</details>

<details>
<summary><b>cyberpunk</b> - Neon dual-color borders</summary>
<img src="assets/themes/cyberpunk.png" alt="cyberpunk theme">
</details>

<details>
<summary><b>synthwave</b> - Neon sunset gradient, 80s retro-future</summary>
<img src="assets/themes/synthwave.png" alt="synthwave theme">
</details>

<details>
<summary><b>matrix</b> - Green terminal hacker style</summary>
<img src="assets/themes/matrix.png" alt="matrix theme">
</details>

<details>
<summary><b>glitch</b> - Digital distortion, cyberpunk broken aesthetic</summary>
<img src="assets/themes/glitch.png" alt="glitch theme">
</details>

#### Nature & Aesthetic

<details>
<summary><b>ocean</b> - Deep sea wave gradient, calm blue tones</summary>
<img src="assets/themes/ocean.png" alt="ocean theme">
</details>

<details>
<summary><b>steampunk</b> - Victorian brass gears, industrial aesthetic</summary>
<img src="assets/themes/steampunk.png" alt="steampunk theme">
</details>

#### System Monitor Style

<details>
<summary><b>htop</b> - Classic system monitor, colorful progress bars</summary>
<img src="assets/themes/htop.png" alt="htop theme">
</details>

<details>
<summary><b>btop</b> - Modern system monitor, gradient colors and rounded frames</summary>
<img src="assets/themes/btop.png" alt="btop theme">
</details>

<details>
<summary><b>gtop</b> - Minimal system monitor, sparklines and clean layout</summary>
<img src="assets/themes/gtop.png" alt="gtop theme">
</details>

<details>
<summary><b>stui</b> - CPU stress test monitor, frequency/temperature style</summary>
<img src="assets/themes/stui.png" alt="stui theme">
</details>

#### Retro & Gaming

<details>
<summary><b>pixel</b> - 8-bit retro game, block characters</summary>
<img src="assets/themes/pixel.png" alt="pixel theme">
</details>

<details>
<summary><b>retro_crt</b> - Green phosphor screen, scanline effect</summary>
<img src="assets/themes/retro_crt.png" alt="retro_crt theme">
</details>

<details>
<summary><b>bbs</b> - Classic BBS ANSI art style</summary>
<img src="assets/themes/bbs.png" alt="bbs theme">
</details>

<details>
<summary><b>lord</b> - Legend of the Red Dragon BBS text game style</summary>
<img src="assets/themes/lord.png" alt="lord theme">
</details>

<details>
<summary><b>tradewars</b> - Trade Wars space trading game, starship console</summary>
<img src="assets/themes/tradewars.png" alt="tradewars theme">
</details>

<details>
<summary><b>nethack</b> - Classic Roguelike dungeon exploration style</summary>
<img src="assets/themes/nethack.png" alt="nethack theme">
</details>

<details>
<summary><b>dungeon</b> - Stone walls with torch lighting, dark adventure atmosphere</summary>
<img src="assets/themes/dungeon.png" alt="dungeon theme">
</details>

<details>
<summary><b>mud_rpg</b> - Classic MUD text adventure character status interface</summary>
<img src="assets/themes/mud_rpg.png" alt="mud_rpg theme">
</details>

## Display Information

### Line 1: Basic Info
- **Model**: Current Claude model (Opus/Sonnet/Haiku)
- **Project**: Current working directory name
- **Git Branch**: Branch name and status (+staged/~dirty)
- **Context**: Context window usage with progress bar
- **Daily Hours**: Total work time today

### Line 2: API Limits
- **Session**: 5-hour API usage rate and reset time
- **Week**: 7-day API usage rate and reset time

Progress bar colors: Green (<50%) → Yellow (50-75%) → Orange (75-90%) → Red (>90%)

### Line 3: Session Stats
- **Tokens**: Total tokens used this session
- **Cost**: Estimated session cost (USD)
- **Duration**: Session length
- **Messages**: Message count
- **Burn Rate**: Hourly cost rate
- **Daily/Weekly Cost**: Accumulated costs
- **Cache Hit**: Cache read ratio (Green ≥70% / Yellow 40-70% / Orange <40%)

## Pricing

Per million tokens (as of Jan 2026):

| Model | Input | Output | Cache Read | Cache Write |
|-------|-------|--------|------------|-------------|
| Opus 4.5 | $5 | $25 | $0.50 | $6.25 |
| Sonnet 4/4.5 | $3 | $15 | $0.30 | $3.75 |
| Haiku 4.5 | $1 | $5 | $0.10 | $1.25 |

## Data Storage

Stats are saved in `~/.claude/session-tracker/`:
- `sessions/` - Individual session data
- `stats/` - Daily and weekly token statistics

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Security

Release artifacts are signed and include SLSA provenance. See [SECURITY.md](SECURITY.md) for verification instructions.

## License

[MIT License](LICENSE)
