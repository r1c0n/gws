@echo off
cd /d ".\bin"
::start "" "gws.exe"
if "%1" == "-h" (
    start "" "gwsvc.exe"
) else (
    start "" "gws.exe"
)