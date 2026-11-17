# Windows Development & Contribution Guide

This guide provides practical recommendations and commands for developers working on **Windows 10 / 11** with PowerShell or Command Prompt.

---

## 🛠️ Recommended Tools for Windows

* **Terminal:** [Windows Terminal](https://aka.ms/terminal) (highly recommended for modern Unicode and ANSI color rendering).
* **Shell:** PowerShell 7+ or Windows PowerShell 5.1.
* **Editor:** Visual Studio Code with the Go and Prettier extensions.
* **Runtime:** Node.js 20+ (LTS) & Go 1.25+.

---

## ⚙️ Environment Setup

### 1. Line Endings (CRLF vs. LF)
GitVerse uses LF line endings. Configure Git to automatically handle line endings safely:

```powershell
git config --global core.autocrlf true
```

### 2. Execution Policy
If PowerShell blocks running local script utilities, set your execution policy to `RemoteSigned`:

```powershell
Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned
```

---

## 🚀 Running GitVerse Locally

```powershell
# 1. Clone repository
git clone https://github.com/TheVicky1/GitVerse.git
cd GitVerse

# 2. Launch local preview server (runs on port 5173)
node scripts/serve-preview.cjs

# 3. In another terminal window, run the automated test suite
node --test internal/webui/jstest/*.test.mjs
```

---

## ⚡ PowerShell Troubleshooting

### Port 5173 already in use
If another process is using port 5173:
```powershell
Get-NetTCPConnection -LocalPort 5173 -ErrorAction SilentlyContinue | Select-Object OwningProcess
Stop-Process -Id <PID> -Force
```

### Go not recognized in PowerShell
If `go` commands fail, ensure Go's binary path is added to your Environment Variables:
```powershell
$env:Path += ";C:\Program Files\Go\bin;$HOME\go\bin"
```

