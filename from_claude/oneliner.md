---
---
# Oneliners

## Using gitbash unzip
```sh
mkdir -p en4 && for i in *.zip; do mkdir -p "en4/${i%.zip}" && unzip -q "$i" -d "en4/${i%.zip}"; done
```
### 💡 Explanation

1. && chains commands (only runs next if previous succeeds)
1. ${i%.zip} removes .zip extension from filename
1. -q flag makes unzip quiet (less output)
1. All commands run in sequence on one line

## Using powershell
```ps
ls *.zip | % { Expand-Archive $_ -DestinationPath "en\$($_.BaseName)" }
```

## Other options
1. One 
```sh
# Create en folder if it doesn't exist
mkdir -p en

# Extract each zip to en/filename folder
for i in *.zip; do 
    folder_name="${i%.zip}"
    mkdir -p "en/$folder_name"
    unzip "$i" -d "en/$folder_name"
done
```
1. Two
```sh
# Navigate to your folder
cd "/c/Users/Public/Downloads/2025/edu.gcfglobal.org-2025"

# Create en folder
mkdir -p en

# Extract all ZIP files, skip existing folders
for i in *.zip; do 
    if [ -f "$i" ]; then
        folder_name="${i%.zip}"
        if [ ! -d "en/$folder_name" ]; then
            echo "Extracting $i to en/$folder_name"
            mkdir -p "en/$folder_name"
            unzip "$i" -d "en/$folder_name"
        else
            echo "Skipping $i - en/$folder_name already exists"
        fi
    fi
done

echo "Extraction complete!"
```
1. Using 7z
```sh
for i in *.zip; do 
    folder_name="${i%.zip}"
    7z x "$i" -o"en/$folder_name"
done
```
1. Using 7z
```sh
7z x file.zip -oen/folder_name
```
1. Powershell
```ps
for %f in (*.zip) do powershell -command "Expand-Archive -Path '%f' -DestinationPath 'en\%~nf'"
```
1. powershell
```ps
Get-ChildItem *.zip | ForEach-Object { Expand-Archive $_.FullName -DestinationPath "en\$($_.BaseName)" }
```

## Bat file with issues
setlocal and enabl... needed to be set following VS Code suggestions
errorlevel kept giving warnings
It stopped after trying to created the folder
On the command line powershell command worked for one line but not the whole script

```bat
@echo off
setlocal enabledelayedexpansion

REM Windows Batch Script to Extract Multiple ZIP Files
REM Extracts all ZIP files from source directory to en\ subdirectory
REM Skips extraction if folder already exists

echo ========================================
echo ZIP File Batch Extractor
echo ========================================
echo.

REM Set source and destination paths
set "SOURCE_PATH=C:\Users\Public\Downloads\2025\edu.gcfglobal.org-2025"
set "DEST_PATH=%SOURCE_PATH%\en"

REM Display paths
echo Source directory: %SOURCE_PATH%
echo Destination directory: %DEST_PATH%
echo.

REM Check if source directory exists
if not exist "%SOURCE_PATH%" (
    echo ERROR: Source directory does not exist!
    echo Path: %SOURCE_PATH%
    pause
    exit /b 1
)

REM Create destination directory if it doesn't exist
if not exist "%DEST_PATH%" (
    echo Creating destination directory: %DEST_PATH%
    mkdir "%DEST_PATH%"
    if errorlevel 1 (
        echo ERROR: Failed to create destination directory!
        pause
        exit /b 1
    )
)

REM Change to source directory
cd /d "%SOURCE_PATH%"

REM Count total ZIP files
set ZIP_COUNT=0
for %%f in (*.zip) do (
    set /a ZIP_COUNT+=1
)

if %ZIP_COUNT%==0 (
    echo No ZIP files found in source directory.
    pause
    exit /b 0
)

echo Found %ZIP_COUNT% ZIP file(s) to process.
echo.

REM Initialize counters
set PROCESSED=0
set SKIPPED=0
set ERRORS=0

REM Process each ZIP file
for %%f in (*.zip) do (
    set "ZIP_FILE=%%f"
    set "ZIP_NAME=%%~nf"
    
    echo Processing: !ZIP_FILE!
    
    REM Check if folder already exists in destination
    if exist "%DEST_PATH%\!ZIP_NAME!" (
        echo   - SKIPPED: Folder !ZIP_NAME! already exists
        set /a SKIPPED+=1
    ) else (
        echo   - EXTRACTING: !ZIP_FILE! to en\!ZIP_NAME!\
        
        REM Create specific folder for this ZIP
        mkdir "%DEST_PATH%\!ZIP_NAME!"
        
        REM Extract ZIP file using PowerShell (built into Windows 10)
        powershell -command "Expand-Archive -Path '!ZIP_FILE!' -DestinationPath '%DEST_PATH%\!ZIP_NAME!' -Force" 2>nul
        
        if errorlevel 1 (
            echo   - ERROR: Failed to extract !ZIP_FILE!
            set /a ERRORS+=1
            REM Remove empty folder if extraction failed
            rmdir "%DEST_PATH%\!ZIP_NAME!" 2>nul
        ) else (
            echo   - SUCCESS: Extracted to en\!ZIP_NAME!\
            set /a PROCESSED+=1
        )
    )
    echo.
)

REM Display summary
echo ========================================
echo EXTRACTION SUMMARY
echo ========================================
echo Total ZIP files found: %ZIP_COUNT%
echo Successfully extracted: %PROCESSED%
echo Skipped (already exists): %SKIPPED%
echo Errors: %ERRORS%
echo.

REM List extracted folders
if %PROCESSED% gtr 0 (
    echo Extracted folders in en\:
    dir "%DEST_PATH%" /b /ad
    echo.
)

REM Handle MHTML files (optional cleanup)
set MHTML_COUNT=0
for %%f in (*.mhtml) do (
    set /a MHTML_COUNT+=1
)

if %MHTML_COUNT% gtr 0 (
    echo Found %MHTML_COUNT% MHTML file(s) - these were ignored as requested.
    echo MHTML files:
    for %%f in (*.mhtml) do (
        echo   - %%f
    )
    echo.
)

if %ERRORS% gtr 0 (
    echo WARNING: %ERRORS% file(s) failed to extract. Check if they are valid ZIP files.
)

echo Script completed.
pause
```