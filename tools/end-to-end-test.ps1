$ErrorActionPreference = 'Stop'
Set-StrictMode -Version 3

# cd to location of script
$script_path = $MyInvocation.MyCommand.Path
$working_dir = Split-Path $script_path
Push-Location $working_dir

if (-not (Test-Path -Path '..\windows_exporter.exe')) {
    Write-Output "..\windows_exporter.exe not found. Consider running \`go build\` first"
}

$temp_dir = Join-Path $env:TEMP $([guid]::newguid()) | ForEach-Object { mkdir $_ }

# Create temporary directory for textfile collector
$textfile_dir = "$($temp_dir)/textfile"
mkdir $textfile_dir | Out-Null
Copy-Item 'e2e-textfile.prom' -Destination "$($textfile_dir)/e2e-textfile.prom"

# Omit dynamic collector information that will change after each run
$skip_re = "^(go_|windows_exporter_build_info|windows_exporter_collector_duration_seconds|windows_exporter_scrape_duration_seconds|process_|windows_textfile_mtime_seconds|windows_cpu|windows_cache|windows_pagefile|windows_logical_disk|windows_physical_disk|windows_memory|windows_net|windows_os|windows_process|windows_service_process|windows_printer|windows_udp|windows_tcp|windows_system|windows_time|windows_session|windows_performancecounter|windows_performancecounter|windows_textfile_mtime_seconds)"

# Start process in background, awaiting HTTP requests.
# Use default collectors, port and address: http://localhost:9182/metrics
$exporter_proc = Start-Process `
    -PassThru `
    -FilePath ..\windows_exporter.exe `
    -ArgumentList "--log.level=debug","--web.disable-exporter-metrics","--collectors.enabled=[defaults],cpu_info,textfile,process,pagefile,performancecounter,scheduled_task,tcp,udp,time,system,service,logical_disk,os,net,memory,cache","--collector.process.include=explorer.exe","--collector.scheduled_task.include=.*GAEvents","--collector.service.include=Themes","--collector.textfile.directories=$($textfile_dir)",@"
--collector.performancecounter.objects="[{\"name\":\"cpu\",\"object\":\"Processor Information\",\"instances\":[\"*\"],\"instance_label\":\"core\",\"counters\":[{\"name\":\"% Processor Time\",\"metric\":\"windows_performancecounter_processor_information_processor_time\",\"labels\":{\"state\":\"active\"}},{\"name\":\"% Idle Time\",\"metric\":\"windows_performancecounter_processor_information_processor_time\",\"labels\":{\"state\":\"idle\"}}]},{\"name\":\"memory\",\"object\":\"Memory\",\"counters\":[{\"name\":\"Cache Faults/sec\",\"type\":\"counter\"}]}]"
"@ `
    -WindowStyle Hidden `
    -RedirectStandardOutput "$($temp_dir)/windows_exporter.log" `
    -RedirectStandardError "$($temp_dir)/windows_exporter_error.log"

# Exporter can take some time to start
for ($i=1; $i -le 1; $i++) {
    Start-Sleep 10

    $netstat_output = netstat -anp tcp | Select-String 'listening'
    if ($netstat_output -like '*:9182*') {
        break
    }
    Write-Host "Waiting for exporter to start"
}

try {
    $response = Invoke-WebRequest -UseBasicParsing -URI http://127.0.0.1:9182/metrics
} catch {
    Write-Host "STDOUT"
    Get-Content "$($temp_dir)/windows_exporter.log"
    Write-Host "STDERR"
    Get-Content "$($temp_dir)/windows_exporter_error.log"

    throw $_
}
# Response output must be split and saved as UTF-8.
$response.content -split "[`r`n]"| Select-String -NotMatch $skip_re | Set-Content -Encoding utf8 "$($temp_dir)/e2e-output.txt"
try {
    Stop-Process -Id $exporter_proc.Id
} catch {
    Write-Host "STDOUT"
    Get-Content "$($temp_dir)/windows_exporter.log"
    Write-Host "STDERR"
    Get-Content "$($temp_dir)/windows_exporter_error.log"

    throw $_
}

# Compare the expected and actual output
$output_diff = Compare-Object (Get-Content 'e2e-output.txt' | Where-Object { $_ -ne "" }) (Get-Content "$($temp_dir)/e2e-output.txt" | Where-Object { $_ -ne "" })

# Fail if differences in output are detected
if (-not ($null -eq $output_diff)) {
    $output_diff | Format-Table -AutoSize | Out-String -Width 10000

    Write-Host "STDOUT"
    Get-Content "$($temp_dir)/windows_exporter.log"
    Write-Host "----------------------------------------"
    Write-Host "STDERR"
    Get-Content "$($temp_dir)/windows_exporter_error.log"

    (Get-Content "$($temp_dir)/e2e-output.txt") | Set-Content -Encoding utf8 "e2e-output.txt"

    exit 1
}
