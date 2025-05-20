# Ensure Go bin is on PATH
$env:PATH += ";$env:USERPROFILE\go\bin"

# Function to find the Go module root by walking up from this script's directory
function Find-GoModRoot {
    $dir = Split-Path -Parent $MyInvocation.MyCommand.Path
    while ($dir -ne [System.IO.Path]::GetPathRoot($dir)) {
        if (Test-Path "$dir\go.mod") {
            return $dir
        }
        $dir = Split-Path -Parent $dir
    }
    Write-Host "❌ go.mod not found. Are you in a Go module?" -ForegroundColor Red
    exit 1
}

$moduleRoot = Find-GoModRoot
Set-Location $moduleRoot

Write-Host "🧠 Running go vet..."
go vet ./...
if ($LASTEXITCODE -ne 0) {
    Write-Host "`n❌ go vet failed. Fix issues before continuing." -ForegroundColor Red
    exit 1
}

Write-Host "🕵️ Running staticcheck..."
if (Get-Command staticcheck -ErrorAction SilentlyContinue) {
    staticcheck ./...
    if ($LASTEXITCODE -ne 0) {
        Write-Host "`n❌ staticcheck failed. Fix issues before continuing." -ForegroundColor Red
        exit 1
    }
} else {
    Write-Host "⚠️ staticcheck not found. Skipping."
}

Write-Host "🔎 Running tests with coverage..."
$coverageFile = "coverage.out"
go test ./... -coverprofile=$coverageFile -json 2>&1 | gotestfmt
$testExitCode = $LASTEXITCODE

if (Test-Path $coverageFile) {
    Write-Host "`n📊 Test Coverage Summary:"
    go tool cover -func=$coverageFile | Select-String "total:"
    Write-Host ""
}

if ($testExitCode -eq 0) {
    Write-Host "`n✅ ✅ ✅  ALL TESTS PASSED ✅ ✅ ✅" -ForegroundColor Green
} else {
    Write-Host "`n❌ ❌ ❌  SOME TESTS FAILED ❌ ❌ ❌" -ForegroundColor Red
}

exit $testExitCode
