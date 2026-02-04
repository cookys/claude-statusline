# Claude Statusline

[![CI](https://github.com/kevinlincg/claude-statusline/actions/workflows/ci.yml/badge.svg)](https://github.com/kevinlincg/claude-statusline/actions/workflows/ci.yml)
[![Release](https://github.com/kevinlincg/claude-statusline/actions/workflows/release.yml/badge.svg)](https://github.com/kevinlincg/claude-statusline/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinlincg/claude-statusline)](https://goreportcard.com/report/github.com/kevinlincg/claude-statusline)
[![GitHub release](https://img.shields.io/github/v/release/kevinlincg/claude-statusline)](https://github.com/kevinlincg/claude-statusline/releases/latest)
[![SLSA 3](https://slsa.dev/images/gh-badge-level3.svg)](https://slsa.dev)

Claude Code 用のカスタムステータスライン。Go 言語で作成。モデル情報、Git ステータス、API 使用量、トークン消費、コスト指標などを表示します。

## インストール

### 必要条件

- Go 1.18+
- macOS または Linux

### 手順

```bash
# リポジトリをクローン
git clone https://github.com/kevinlincg/claude-statusline.git ~/.claude/statusline-go

# ビルド
cd ~/.claude/statusline-go
go build -o statusline statusline.go

# Claude Code を設定 (~/.claude/settings.json)
{
  "statusLine": {
    "type": "command",
    "command": "/path/to/.claude/statusline-go/statusline"
  }
}
```

## テーマ

`~/.claude/statusline-go/config.json` を編集してステータスラインの外観をカスタマイズ：

```json
{
  "theme": "classic_framed"
}
```

### 利用可能なテーマ

| テーマ | 説明 |
|--------|------|
| `classic` | オリジナルレイアウトスタイル |
| `classic_framed` | ツリー構造＋フレーム、整列されたプログレスバー |
| `minimal` | シンプルなツリー構造、ボーダーなし |
| `compact` | 最小限の高さ、完全な情報 |
| `boxed` | 完全なボーダーフレーム、対称的なセクション |
| `zen` | ミニマリストの余白、穏やかでエレガント |
| `hud` | SF HUD インターフェース、角括弧ラベル |
| `cyberpunk` | ネオンデュアルカラーボーダー |
| `synthwave` | ネオンサンセットグラデーション、80年代レトロフューチャー |
| `matrix` | グリーンターミナルハッカースタイル |
| `glitch` | デジタル歪み、サイバーパンク破壊美学 |
| `ocean` | 深海の波グラデーション、穏やかなブルートーン |
| `pixel` | 8-bit レトロゲーム、ブロック文字 |
| `retro_crt` | グリーン蛍光スクリーン、スキャンライン効果 |
| `steampunk` | ヴィクトリア朝の真鍮歯車、工業美学 |
| `htop` | クラシックシステムモニター、カラフルなプログレスバー |
| `btop` | モダンシステムモニター、グラデーションカラーと丸角フレーム |
| `gtop` | ミニマルシステムモニター、スパークラインとクリーンなレイアウト |
| `stui` | CPU ストレステストモニター、周波数/温度スタイル |
| `bbs` | クラシック BBS ANSI アートスタイル |
| `lord` | Legend of the Red Dragon BBS テキストゲームスタイル |
| `tradewars` | Trade Wars 宇宙貿易ゲーム、宇宙船コンソール |
| `nethack` | クラシック Roguelike ダンジョン探索スタイル |
| `dungeon` | 松明に照らされた石壁、ダークアドベンチャー雰囲気 |
| `mud_rpg` | クラシック MUD テキストアドベンチャーキャラクターステータス |

## 表示情報

### 1行目：基本情報
- **モデル**：現在の Claude モデル（Opus/Sonnet/Haiku）
- **プロジェクト**：現在の作業ディレクトリ名
- **Git ブランチ**：ブランチ名とステータス（+ステージ済み/~未ステージ）
- **Context**：Context Window 使用量プログレスバー
- **日次作業時間**：今日の合計作業時間

### 2行目：API 制限
- **Session**：5時間の API 使用率とリセット時間
- **Week**：7日間の API 使用率とリセット時間

プログレスバーの色：緑 (<50%) → 黄 (50-75%) → オレンジ (75-90%) → 赤 (>90%)

### 3行目：Session 統計
- **トークン**：このセッションで使用したトークン総数
- **コスト**：推定セッションコスト (USD)
- **時間**：セッション長
- **メッセージ**：メッセージ数
- **消費速度**：時間あたりのコスト
- **日次/週次コスト**：累積コスト
- **キャッシュヒット**：キャッシュ読み取り比率（緑 ≥70% / 黄 40-70% / オレンジ <40%）

## 価格

100万トークンあたり（2026年1月現在）：

| モデル | 入力 | 出力 | キャッシュ読取 | キャッシュ書込 |
|--------|------|------|----------------|----------------|
| Opus 4.5 | $5 | $25 | $0.50 | $6.25 |
| Sonnet 4/4.5 | $3 | $15 | $0.30 | $3.75 |
| Haiku 4.5 | $1 | $5 | $0.10 | $1.25 |

## データ保存

統計は `~/.claude/session-tracker/` に保存されます：
- `sessions/` - 個別セッションデータ
- `stats/` - 日次・週次トークン統計

## ライセンス

MIT License
