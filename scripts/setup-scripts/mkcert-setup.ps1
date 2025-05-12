$certDir = "nginx\certs"
$domain = "limitlesshoops.dev"

Write-Output "🔐 Setting up mkcert certificates for $domain..."

# Check mkcert
if (-not (Get-Command mkcert -ErrorAction SilentlyContinue)) {
    Write-Output "📦 Installing mkcert..."

    Invoke-WebRequest -Uri "https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-windows-amd64.exe" -OutFile "mkcert.exe"
    Move-Item "mkcert.exe" "$env:ProgramFiles\mkcert.exe"
    $env:Path += ";$env:ProgramFiles"
}

# Initialize CA
& mkcert.exe -install

# Ensure cert directory
if (-not (Test-Path $certDir)) {
    New-Item -ItemType Directory -Path $certDir | Out-Null
}

# Generate certs
& mkcert.exe -key-file "$certDir\$domain.key" -cert-file "$certDir\$domain.crt" $domain

Write-Output "✅ Certificates created:"
Write-Output "  $certDir\$domain.crt"
Write-Output "  $certDir\$domain.key"
