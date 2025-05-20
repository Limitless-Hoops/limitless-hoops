# Requires: Admin privileges for Docker install

$ErrorActionPreference = "Stop"

# ==== CONFIG ====
$goVersion = "1.24.2"
$dockerDesktopVersion = "4.30.0"
$goInstallerUrl = "https://go.dev/dl/go$goVersion.windows-amd64.msi"
$dockerInstallerUrl = "https://desktop.docker.com/win/main/$dockerDesktopVersion/Docker%20Desktop%20Installer.exe"
$gotestfmtVersion = "v1.3.0"
# =================

Write-Host "🔧 Starting tool installation for Windows..."

function Install-Go {
    Write-Host "`n📦 Checking Go installation..."

    try {
        $currentGo = & go version
        if ($currentGo -like "*$goVersion*") {
            Write-Host "✅ Go $goVersion is already installed"
            return
        }
    } catch {
        Write-Host "ℹ️ Go is not currently installed"
    }

    Write-Host "⬇️ Downloading Go $goVersion installer..."
    $goInstallerPath = "$env:TEMP\go$goVersion.msi"
    Invoke-WebRequest -Uri $goInstallerUrl -OutFile $goInstallerPath

    Write-Host "🛠 Installing Go..."
    Start-Process msiexec.exe -Wait -ArgumentList "/i `"$goInstallerPath`" /quiet /norestart"
    Remove-Item $goInstallerPath

    Write-Host "✅ Go $goVersion installed successfully"
}

function Install-Docker {
    Write-Host "`n🐳 Checking Docker Desktop installation..."

    try {
        $dockerVersion = & docker --version
        if ($dockerVersion -like "*28.1.1*") {
            Write-Host "✅ Docker 28.1.1 is already installed"
            return
        }
    } catch {
        Write-Host "ℹ️ Docker is not currently installed"
    }

    Write-Host "⬇️ Downloading Docker Desktop $dockerDesktopVersion installer..."
    $dockerInstallerPath = "$env:TEMP\DockerDesktopInstaller.exe"
    Invoke-WebRequest -Uri $dockerInstallerUrl -OutFile $dockerInstallerPath

    Write-Host "🛠 Installing Docker Desktop (this may take a few minutes)..."
    Start-Process -FilePath $dockerInstallerPath -Wait -ArgumentList "install", "--quiet"
    Remove-Item $dockerInstallerPath

    Write-Host "✅ Docker Desktop $dockerDesktopVersion installed"
}

function Install-Gotestfmt {
    Write-Host "`n🧪 Installing gotestfmt $gotestfmtVersion..."
    & "$env:USERPROFILE\go\bin\go.exe" install "github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@$gotestfmtVersion"
    Write-Host "✅ gotestfmt installed"
}

function Install-Staticcheck {
    Write-Host "`n🕵️ Installing staticcheck..."
    & "$env:USERPROFILE\go\bin\go.exe" install "honnef.co/go/tools/cmd/staticcheck@latest"
    Write-Host "✅ staticcheck installed"
}

Install-Go
Install-Docker
Install-Gotestfmt
Install-Staticcheck

# Ensure Go bin is in PATH for future use
if (-not ($env:PATH -like "*$env:USERPROFILE\go\bin*")) {
    Write-Host "🔧 Adding Go bin path to user PATH..."

    $newPath = "$env:USERPROFILE\go\bin"
    $currentUserPath = [Environment]::GetEnvironmentVariable("PATH", "User")

    if ($currentUserPath -notlike "*$newPath*") {
        [Environment]::SetEnvironmentVariable("PATH", "$currentUserPath;$newPath", "User")
        Write-Host "✅ Go bin path added to user PATH (restart terminal to apply)"
    }
}

Write-Host "`n✅ All tools installed successfully!"

Write-Host "`n🔍 Verifying installed tool versions..."

Write-Host "📦 Go version:"
& go version

Write-Host "`n🧪 gotestfmt version:"
& gotestfmt --version

Write-Host "`n🕵️ staticcheck version:"
& staticcheck --version

Write-Host "`n🐳 Docker version:"
& docker --version
