# Ensure we're in the right project directory
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

# Rerun to check exit code (cached, so fast)
go test ./... > $null
if ($LASTEXITCODE -eq 0) {
Write-Host "`n✅ ✅ ✅  ALL TESTS PASSED ✅ ✅ ✅" -ForegroundColor Green
} else {
Write-Host "`n❌ ❌ ❌  SOME TESTS FAILED ❌ ❌ ❌" -ForegroundColor Red
}
