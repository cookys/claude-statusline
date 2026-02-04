# Claude Statusline

为 Claude Code 打造的自定义状态栏，使用 Go 语言编写。显示模型信息、Git 状态、API 使用量、Token 消耗、成本指标等。

## 安装

### 需求

- Go 1.18+
- macOS 或 Linux

### 步骤

```bash
# 克隆项目
git clone https://github.com/kevinlincg/claude-statusline.git ~/.claude/statusline-go

# 编译
cd ~/.claude/statusline-go
go build -o statusline statusline.go

# 配置 Claude Code (~/.claude/settings.json)
{
  "statusLine": {
    "type": "command",
    "command": "/path/to/.claude/statusline-go/statusline"
  }
}
```

## 主题

编辑 `~/.claude/statusline-go/config.json` 来自定义状态栏外观：

```json
{
  "theme": "classic_framed"
}
```

### 可用主题

| 主题 | 说明 |
|------|------|
| `classic` | 原版经典：保持原有布局风格 |
| `classic_framed` | 经典树状+框线：左侧文字信息，右侧光棒垂直对齐 |
| `minimal` | 简洁树状：无外框，树状结构显示信息 |
| `compact` | 精简三行：最小高度，信息完整 |
| `boxed` | 框线整齐：完整框线包围，左右对称分区 |
| `zen` | 禅风：极简留白，宁静淡雅 |
| `hud` | 科幻 HUD：未来感界面，角括号标签 |
| `cyberpunk` | 赛博朋克：霓虹双色框线 |
| `synthwave` | 合成波：霓虹日落渐变，80年代复古未来 |
| `matrix` | 矩阵黑客：绿色终端机风格 |
| `glitch` | 故障风：数字错位，赛博朋克破碎美学 |
| `ocean` | 深海：海洋波浪渐变，宁静蓝调 |
| `pixel` | 像素风：8-bit 复古游戏，方块字符 |
| `retro_crt` | 复古 CRT：绿色磷光屏幕，扫描线效果 |
| `steampunk` | 蒸汽朋克：维多利亚黄铜齿轮，工业美学 |
| `htop` | htop：经典系统监视器，彩色进度条风格 |
| `btop` | btop：现代系统监视器，渐变色彩与圆角框风格 |
| `gtop` | gtop：简约系统监视器，火花图与干净排版 |
| `stui` | s-tui：CPU 压力测试监视器，频率温度图风格 |
| `bbs` | BBS：经典电子布告栏 ANSI 艺术风格 |
| `lord` | LORD：红龙传说 BBS 经典文字游戏风格 |
| `tradewars` | Trade Wars：太空贸易游戏，星舰控制台风格 |
| `nethack` | NetHack：经典 Roguelike 地牢探索风格 |
| `dungeon` | 地牢：石墙火把照明，黑暗冒险氛围 |
| `mud_rpg` | MUD RPG：经典文字冒险游戏角色状态界面 |

## 显示信息

### 第一行：基本信息
- **模型**：当前使用的 Claude 模型（Opus/Sonnet/Haiku）
- **项目**：当前工作目录名称
- **Git 分支**：分支名称与状态（+已暂存/~未暂存）
- **Context**：Context Window 使用量进度条
- **每日工时**：今日累计工作时间

### 第二行：API 限制
- **Session**：5 小时内 API 使用率与重置时间
- **Week**：7 天内 API 使用率与重置时间

进度条颜色：绿色 (<50%) → 黄色 (50-75%) → 橙色 (75-90%) → 红色 (>90%)

### 第三行：Session 统计
- **Token**：本次 Session 累计使用的 Token 数量
- **成本**：本次 Session 的预估成本 (USD)
- **时长**：Session 持续时间
- **消息数**：对话消息数量
- **烧钱速度**：每小时花费
- **今日/周成本**：累计成本
- **Cache 命中率**：Cache read 比例（绿色 ≥70% / 黄色 40-70% / 橙色 <40%）

## 定价

每百万 Token（2026 年 1 月）：

| 模型 | 输入 | 输出 | Cache 读取 | Cache 写入 |
|------|------|------|------------|------------|
| Opus 4.5 | $5 | $25 | $0.50 | $6.25 |
| Sonnet 4/4.5 | $3 | $15 | $0.30 | $3.75 |
| Haiku 4.5 | $1 | $5 | $0.10 | $1.25 |

## 数据存储

统计数据保存于 `~/.claude/session-tracker/`：
- `sessions/` - 单个 Session 数据
- `stats/` - 每日与每周 Token 统计

## 许可证

MIT License
