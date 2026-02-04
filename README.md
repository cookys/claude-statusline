# Claude Statusline

[![CI](https://github.com/kevinlincg/claude-statusline/actions/workflows/ci.yml/badge.svg)](https://github.com/kevinlincg/claude-statusline/actions/workflows/ci.yml)
[![Release](https://github.com/kevinlincg/claude-statusline/actions/workflows/release.yml/badge.svg)](https://github.com/kevinlincg/claude-statusline/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinlincg/claude-statusline)](https://goreportcard.com/report/github.com/kevinlincg/claude-statusline)
[![GitHub release](https://img.shields.io/github/v/release/kevinlincg/claude-statusline)](https://github.com/kevinlincg/claude-statusline/releases/latest)
[![SLSA 3](https://slsa.dev/images/gh-badge-level3.svg)](https://slsa.dev)

[English](README.md) | [繁體中文](README.zh-TW.md) | [简体中文](README.zh-CN.md) | [日本語](README.ja.md)

A custom status line for Claude Code written in Go. Displays model info, Git status, API usage, token consumption, cost metrics, and more.

## Installation

### Requirements

- Go 1.18+
- macOS or Linux

### Steps

```bash
# Clone the repository
git clone https://github.com/kevinlincg/claude-statusline.git ~/.claude/statusline-go

# Build
cd ~/.claude/statusline-go
go build -o statusline statusline.go

# Configure Claude Code (~/.claude/settings.json)
{
  "statusLine": {
    "type": "command",
    "command": "/path/to/.claude/statusline-go/statusline"
  }
}
```

## Themes

Customize your status line appearance by editing `~/.claude/statusline-go/config.json`:

```json
{
  "theme": "classic_framed"
}
```

### Available Themes

| Theme | Description |
|-------|-------------|
| `classic` | Original layout style |
| `classic_framed` | Tree structure with frame, aligned progress bars |
| `minimal` | Clean tree structure, no borders |
| `compact` | Minimal height, complete info |
| `boxed` | Full border frame, symmetrical sections |
| `zen` | Minimalist whitespace, calm and elegant |
| `hud` | Sci-fi HUD interface with angle bracket labels |
| `cyberpunk` | Neon dual-color borders |
| `synthwave` | Neon sunset gradient, 80s retro-future |
| `matrix` | Green terminal hacker style |
| `glitch` | Digital distortion, cyberpunk broken aesthetic |
| `ocean` | Deep sea wave gradient, calm blue tones |
| `pixel` | 8-bit retro game, block characters |
| `retro_crt` | Green phosphor screen, scanline effect |
| `steampunk` | Victorian brass gears, industrial aesthetic |
| `htop` | Classic system monitor, colorful progress bars |
| `btop` | Modern system monitor, gradient colors and rounded frames |
| `gtop` | Minimal system monitor, sparklines and clean layout |
| `stui` | CPU stress test monitor, frequency/temperature style |
| `bbs` | Classic BBS ANSI art style |
| `lord` | Legend of the Red Dragon BBS text game style |
| `tradewars` | Trade Wars space trading game, starship console |
| `nethack` | Classic Roguelike dungeon exploration style |
| `dungeon` | Stone walls with torch lighting, dark adventure atmosphere |
| `mud_rpg` | Classic MUD text adventure character status interface |

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

## License

MIT License
