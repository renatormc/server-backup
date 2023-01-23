$CurrentDir = Get-Location
$srcFolder = "."
Set-Location $srcFolder
go build -o ".\dist\server-backup.exe"
if($?){
    Set-Location $CurrentDir
    .\dist\server-backup.bat $args
}

