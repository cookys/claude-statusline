# Claude Statusline (Go 版本)

這是一個用於 Claude Code 自定義狀態欄的 Go 程式。它可以顯示目前的模型、Git 分支、專案名稱、Session 持續時間以及 API 使用量統計。

## 功能特點

- 💠 **模型顯示**：顯示目前使用的 Claude 模型 (Opus, Sonnet, Haiku)。
- ⚡ **Git 整合**：顯示目前的 Git 分支名稱。
- ⏱️ **時間追蹤**：追蹤每日與每個 Session 的工作時間。
- 💰 **成本估算**：即時計算目前的 Token 消耗與預估成本。
- 📊 **API 限制監控**：顯示 5 小時與 7 天的 API 使用量進度條與重置時間。
- 💬 **訊息統計**：顯示目前 Session 的對話訊息數量。

## 安裝與編譯

### 1. 編譯程式

在 `~/.claude/statusline-go` 目錄下執行：

```bash
go build -o statusline statusline.go
```

### 2. 配置 Claude Code

編輯你的 `~/.claude/settings.json`（或是在 Claude Code 中輸入 `/config`），加入以下設定：

```json
{
  "customStatuslineCommand": "/Users/your-username/.claude/statusline-go/statusline"
}
```

> [!IMPORTANT]
> 請確保將 `/Users/your-username/` 替換為你實際的使用者目錄路徑。

## 設定檔說明 (`settings.json`)

如果你想要在 Claude Code 中使用此自定義狀態欄，你需要確保 `customStatuslineCommand` 指向編譯後的執行檔。

此程式會從 `stdin` 讀取 Claude Code 傳入的 JSON 資訊，格式如下：

```json
{
  "model": { "display_name": "Claude 3.5 Sonnet" },
  "session_id": "...",
  "workspace": { "current_dir": "..." },
  "transcript_path": "..."
}
```

並輸出三行資訊供 Claude Code 顯示。

## 本地資料儲存

本程式會在以下路徑儲存統計資料：
- `~/.claude/session-tracker/sessions/`: 個別 Session 的時間與資訊。
- `~/.claude/session-tracker/stats/`: 每日與每週的 Token 使用統計。

## 授權

MIT License
