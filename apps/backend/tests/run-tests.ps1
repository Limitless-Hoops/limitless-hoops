Write-Host "🧠 Running go vet..." -ForegroundColor Cyan
go vet ./...
if ($LASTEXITCODE -ne 0) {
    Write-Host "`n❌ go vet failed. Fix issues before continuing." -ForegroundColor Red
    exit 1
}

Write-Host "🕵️ Running staticcheck..." -ForegroundColor Cyan
if (Get-Command staticcheck -ErrorAction SilentlyContinue) {
    staticcheck ./...
    if ($LASTEXITCODE -ne 0) {
        Write-Host "`n❌ staticcheck failed. Fix issues before continuing." -ForegroundColor Red
        exit 1
    }
} else {
    Write-Host "⚠️ staticcheck not found. Skipping." -ForegroundColor Yellow
}

Write-Host "🔎 Running tests..." -ForegroundColor Cyan

# Define temp file paths
$tempJson = "test_output.json"
$tempFormatted = "test_output.txt"

# Run tests, capture JSON output
go test ./... -json > $tempJson

# Format output using gotestfmt
gotestfmt < $tempJson > $tempFormatted

# Print the formatted test results
Get-Content $tempFormatted

# Clean up temp files
Remove-Item $tempJson, $tempFormatted -ErrorAction SilentlyContinue

# Rerun to check exit code
go test ./... > $null
if ($LASTEXITCODE -eq 0) {
    Write-Host "`n✅ ✅ ✅  ALL TESTS PASSED ✅ ✅ ✅" -ForegroundColor Green
} else {
    Write-Host "`n❌ ❌ ❌  SOME TESTS FAILED ❌ ❌ ❌" -ForegroundColor Red
}
