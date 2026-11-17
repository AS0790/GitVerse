# GitVerse Automated Local Verification Script for Windows
$ErrorActionPreference = "Stop"
Write-Host "Running go vet..." -ForegroundColor Cyan
go vet ./...
Write-Host "Running tests..." -ForegroundColor Cyan
go test ./...
Write-Host "Running gendocs check..." -ForegroundColor Cyan
go run ./cmd/gendocs -check
Write-Host "Running JavaScript unit tests..." -ForegroundColor Cyan
node --test "internal/webui/jstest/*.test.mjs"
Write-Host "All local checks passed successfully!" -ForegroundColor Green
