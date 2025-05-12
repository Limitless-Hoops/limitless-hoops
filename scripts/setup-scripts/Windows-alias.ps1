$hostsPath = "$env:SystemRoot\System32\drivers\etc\hosts"
$entry = "127.0.0.1 limitlesshoops.dev www.limitlesshoops.dev"

# Read the file content
$hostsContent = Get-Content -Path $hostsPath -ErrorAction Stop

# Check if the entry already exists
if ($hostsContent -match "limitlesshoops\.dev") {
    Write-Output "Entry already exists in $hostsPath"
} else {
    Write-Output "Adding entry to $hostsPath"
    Add-Content -Path $hostsPath -Value $entry
    Write-Output "Entry added successfully."
}

