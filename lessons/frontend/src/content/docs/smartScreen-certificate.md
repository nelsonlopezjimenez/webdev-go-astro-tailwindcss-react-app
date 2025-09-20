---
title: "smarScreen certificate generation"
description: "To sign your Go executable for Windows, you have several options. Here's a comprehensive guide:"
publishDate: 2025-09-18
tags: ["smartScreen","cert","certificate"]
draft: false
---

To sign your Go executable for Windows, you have several options. Here's a comprehensive guide:

## Option 1: Self-Signed Certificate (Free, for Testing/Internal Use)

### Step 1: Create a Self-Signed Certificate

```powershell
# Run in PowerShell as Administrator
# Create a self-signed certificate
$cert = New-SelfSignedCertificate -Subject "CN=YourCompanyName" -Type CodeSigning -KeyUsage DigitalSignature -FriendlyName "Code Signing Certificate" -CertStoreLocation Cert:\CurrentUser\My -HashAlgorithm SHA256

# Export certificate to file (optional, for distribution)
Export-Certificate -Cert $cert -FilePath "C:\temp\codecert.crt"
```

### Step 2: Sign Your Executable

```powershell
# Using PowerShell
Set-AuthenticodeSignature -FilePath "course-server-2.12.0.exe" -Certificate $cert

# Or using signtool (if you have Windows SDK)
signtool sign /n "YourCompanyName" /t http://timestamp.digicert.com /fd sha256 course-server-2.12.0.exe
```

## Option 2: Commercial Code Signing Certificate (Recommended for Distribution)

### Popular Certificate Authorities:
- **DigiCert** (~$474/year)
- **Sectigo/Comodo** (~$200/year)  
- **GlobalSign** (~$249/year)
- **Entrust** (~$699/year)

### Steps:
1. **Purchase certificate** from a CA
2. **Verify your identity** (company verification process)
3. **Download certificate** (.p12 or .pfx file)
4. **Sign your executable**

```powershell
# Sign with commercial certificate
signtool sign /f "your-certificate.pfx" /p "password" /t http://timestamp.digicert.com /fd sha256 course-server-2.12.0.exe

# Or with certificate in store
signtool sign /n "Your Company Name" /t http://timestamp.digicert.com /fd sha256 course-server-2.12.0.exe
```

## Option 3: Windows Store Certificate (Free, for Store Apps)

If distributing through Microsoft Store, use their signing process.

## Option 4: GitHub Actions with Azure Code Signing (Easiest)

For automated signing in CI/CD:

```yaml
# .github/workflows/build-and-sign.yml
name: Build and Sign
on: [push, pull_request]

jobs:
  build:
    runs-on: windows-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Build
      run: go build -o course-server-2.12.0.exe main.go
    
    - name: Sign executable
      uses: dlemstra/code-sign-action@v1
      with:
        certificate: '${{ secrets.CERTIFICATE }}'
        password: '${{ secrets.CERTIFICATE_PASSWORD }}'
        folder: '.'
        recursive: false
```

## Immediate Solution for Testing

For quick testing, use the self-signed approach:## Usage Instructions:

1. **Save the script** as `sign-executable.ps1`
2. **Run as Administrator**:
   ```powershell
   .\sign-executable.ps1 -ExecutablePath "course-server-2.12.0.exe" -CompanyName "Your Course Platform"
   ```

## Quick One-Liner:

```powershell
# Create cert and sign in one command (run as Admin)
$cert = New-SelfSignedCertificate -Subject "CN=CourseServer" -Type CodeSigning -KeyUsage DigitalSignature -CertStoreLocation Cert:\CurrentUser\My; Set-AuthenticodeSignature -FilePath "course-server-2.12.0.exe" -Certificate $cert
```

## Recommendations:

**For Development/Testing**: Use self-signed certificate (Option 1)  
**For Internal Company Use**: Use self-signed + install in company's trusted root store  
**For Public Distribution**: Purchase commercial certificate (Option 2)  
**For Continuous Deployment**: Use GitHub Actions with Azure Code Signing

## Note About Windows SmartScreen:

Even with a valid signature, new executables may trigger SmartScreen warnings until they build reputation. Commercial certificates from established CAs reduce this issue.

The self-signed approach will immediately solve your current signing need and can be upgraded to commercial certificates later.


```sh
# Run this PowerShell script as Administrator
# sign-executable.ps1

param(
    [Parameter(Mandatory=$true)]
    [string]$ExecutablePath,
    [string]$CompanyName = "Your Company Name"
)

# Check if running as admin
if (-NOT ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Error "This script must be run as Administrator"
    exit 1
}

Write-Host "Creating self-signed certificate..." -ForegroundColor Green

# Create self-signed certificate
$cert = New-SelfSignedCertificate -Subject "CN=$CompanyName" -Type CodeSigning -KeyUsage DigitalSignature -FriendlyName "Code Signing Certificate" -CertStoreLocation Cert:\CurrentUser\My -HashAlgorithm SHA256

Write-Host "Certificate created with thumbprint: $($cert.Thumbprint)" -ForegroundColor Yellow

# Sign the executable
Write-Host "Signing executable: $ExecutablePath" -ForegroundColor Green
$result = Set-AuthenticodeSignature -FilePath $ExecutablePath -Certificate $cert

if ($result.Status -eq "Valid") {
    Write-Host "✅ Successfully signed $ExecutablePath" -ForegroundColor Green
    Write-Host "Certificate Subject: $($cert.Subject)" -ForegroundColor Gray
    Write-Host "Valid Until: $($cert.NotAfter)" -ForegroundColor Gray
} else {
    Write-Error "❌ Failed to sign executable: $($result.StatusMessage)"
    exit 1
}

# Verify signature
Write-Host "Verifying signature..." -ForegroundColor Blue
$signature = Get-AuthenticodeSignature -FilePath $ExecutablePath
Write-Host "Signature Status: $($signature.Status)" -ForegroundColor $(if ($signature.Status -eq "Valid") { "Green" } else { "Red" })

Write-Host @"

📋 IMPORTANT NOTES:
- This is a self-signed certificate for testing purposes
- Windows will show a warning about "unknown publisher"
- For production distribution, purchase a commercial certificate
- Users can bypass warnings by right-clicking and selecting "Run anyway"

🔧 To install certificate in Trusted Root (removes warnings):
1. Double-click the signed .exe
2. Click "More info" → "Run anyway"
   OR
3. Import certificate to Trusted Root Certification Authorities store

"@ -ForegroundColor Cyan
```

**Azure Code Signing itself is FREE**, but there are some associated costs. Let me break this down:

## Azure Code Signing Costs:

### ✅ **FREE Components:**
- **Azure Code Signing service** - No charge for the signing operations
- **GitHub Actions** - Free for public repositories (2,000 minutes/month for private repos)
- **Basic Azure account** - Free tier available

### 💰 **PAID Components:**

#### 1. **Certificate Authority (CA) Certificate** - **Required**
- **Azure Key Vault** stores your certificate: ~$0.03 per 10,000 operations
- **CA Certificate Purchase**: $200-500/year (same cost as commercial certificates)
  - DigiCert: ~$474/year
  - Sectigo: ~$200/year  
  - GlobalSign: ~$249/year

#### 2. **Azure Storage (minimal)**
- **Key Vault storage**: ~$0.03/month per certificate
- **Blob storage** (if used): ~$0.02/GB/month

## **Total Annual Cost Estimate:**
- **Certificate**: $200-500/year (one-time annual purchase)
- **Azure services**: $1-5/year (very minimal usage)
- **GitHub Actions**: $0 (unless heavy CI/CD usage)

**Total: ~$200-505/year**

## Alternative FREE Options:

### 1. **Self-Signed Certificate** (Completely Free)
```powershell
# Free but users see "unknown publisher" warnings
$cert = New-SelfSignedCertificate -Subject "CN=YourApp" -Type CodeSigning
Set-AuthenticodeSignature -FilePath "app.exe" -Certificate $cert
```

### 2. **Open Source Signing** (Free for OSS projects)
- Some CAs offer free certificates for open source projects
- **SignPath.io** offers free code signing for open source

### 3. **GitHub Actions + Self-Signed** (Free)
```yaml
# .github/workflows/build.yml
- name: Sign with self-signed cert
  run: |
    $cert = New-SelfSignedCertificate -Subject "CN=MyApp" -Type CodeSigning -CertStoreLocation Cert:\CurrentUser\My
    Set-AuthenticodeSignature -FilePath "app.exe" -Certificate $cert
```

## **My Recommendation:**

### **For Your Use Case:**
1. **Start with self-signed** (free) for development/testing
2. **Upgrade to commercial certificate** ($200-500/year) when ready for wider distribution
3. **Use GitHub Actions** for automated signing (mostly free)

### **Skip Azure Code Signing if:**
- You're just getting started
- Budget is a concern
- Internal/personal use only

### **Use Azure Code Signing if:**
- You want enterprise-grade certificate management
- You're already using Azure ecosystem
- You need centralized certificate storage and access control

## **Bottom Line:**
Azure Code Signing isn't free due to the required CA certificate purchase (~$200-500/year), but it's not more expensive than buying certificates directly. The Azure services themselves add minimal cost (~$1-5/year).

**For most developers: Start with self-signed certificates, then buy a commercial certificate directly when needed.**