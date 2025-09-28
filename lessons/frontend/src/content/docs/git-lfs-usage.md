---
title: "git LFS usage"
description: "How to use LFS when a file is larger than 100 MB"
publishDate: 2025-09-28
tags: ["git", "lfs", "100MB", "find"]
draft: false
---
# Check current directory file sizes
find . -name "*.exe" -o -name "*.gz" | xargs ls -lh

# Check files over 50MB
find . -size +50M -name "*.exe" -o -size +50M -name "*.gz"

# Check files over 100MB (GitHub limit)
find . -size +100M
Here's how to use Git LFS (Large File Storage) for your 100+ MB file:

Great question! You have a mix of file sizes, so here's the **best selective approach** for LFS:

## 🎯 **Recommended Strategy: Size-Based LFS**

### **Option 1: Track Only Large Files (>50MB)**
```bash
# Initialize LFS
git lfs install

# Track only the large exe file specifically
git lfs track "COMPUTER_BASICS-8.25.2025.exe"

# Or use a size-based approach (if your Git version supports it)
git lfs track "*.exe" --above=50MB
```

### **Option 2: Selective Tracking by Name Pattern**
```bash
# Track only files with specific patterns that tend to be large
git lfs track "*COMPUTER_BASICS*.exe"
git lfs track "*course-server*.exe"

# Leave small utilities in regular Git
# (gcf-browser.exe at 12M can stay in regular Git)
```

## 📋 **File Analysis & Recommendations**

**Your files:**
- ✅ **COMPUTER_BASICS-8.25.2025.exe (572M)** → **MUST use LFS** (exceeds 100MB GitHub limit)
- ⚠️ **course-server-2.*.exe (13-16M)** → **Optional LFS** (small enough for regular Git)
- ✅ **gcf-browser.exe (12M)** → **Regular Git** (small, no LFS needed)

## 🚀 **Recommended Implementation**

### **Step 1: Configure LFS for Large Files Only**
```bash
cd your-repo

# Initialize LFS
git lfs install

# Track only files over 50MB
# Method 1: Specific large file
git lfs track "COMPUTER_BASICS-8.25.2025.exe"

# Method 2: Pattern for large course files (optional)
git lfs track "*COMPUTER_BASICS*.exe"

# Don't track small exe files - keep in regular Git
```

### **Step 2: Verify Tracking**
```bash
# Check what's being tracked
git lfs track

# Should show only your large files
```

### **Step 3: Add Files Appropriately**
```bash
# Add LFS configuration
git add .gitattributes

# Add your files (LFS will handle the large ones automatically)
git add *.exe

# Commit
git commit -m "Add exe files - large ones via LFS, small ones regular Git"

# Push
git push origin main
```

## 💡 **Why This Approach is Best**

### **Benefits:**
✅ **Cost-effective** - Only large files use LFS bandwidth  
✅ **Performance** - Small files clone faster (no LFS overhead)  
✅ **Selective** - Each file type handled optimally  
✅ **Future-proof** - New large files auto-tracked by pattern  

### **File Handling:**
- **572M file** → LFS (required, exceeds GitHub limit)
- **13-16M files** → Regular Git (optional, your choice)
- **12M file** → Regular Git (efficient, no LFS needed)

## 🔄 **Alternative Approaches**

### **Option A: Track All EXE Files**
```bash
# Simple but less efficient
git lfs track "*.exe"
```
**Pros:** Simple, consistent  
**Cons:** Wastes LFS quota on small files

### **Option B: Size-Based Threshold**
```bash
# Track only files over certain size (if supported)
git lfs track "*.exe" --size-threshold=50MB
```

### **Option C: Manual Selection**
```bash
# Track specific large files only
git lfs track "COMPUTER_BASICS-*.exe"
git lfs track "*course-server-2.1[6-9].*.exe"  # Future versions
```

## 📝 **Your .gitattributes File Should Look Like:**

```
# Large educational software (>50MB)
COMPUTER_BASICS-*.exe filter=lfs diff=lfs merge=lfs -text

# Optional: Future large course-server versions
*course-server-2.1[6-9].*.exe filter=lfs diff=lfs merge=lfs -text
```

## 🎯 **Final Recommendation**

```bash
# Best approach for your use case:
git lfs install
git lfs track "COMPUTER_BASICS-*.exe"  # Only the large educational software
git add .gitattributes
git add *.exe
git commit -m "Add exe files: large ones via LFS, small ones regular Git"
git push origin main
```

This gives you:
- **Required LFS** for the 572M file
- **Regular Git** for smaller files (saves LFS quota)
- **Flexible pattern** for future large educational software
- **Cost-effective** approach

The smaller exe files (12-16M) work perfectly fine in regular Git and don't need the LFS overhead! 🚀

## 🚀 **Quick Setup for Git LFS**

### **1. Install Git LFS**

**Windows:**
```bash
# If you have Git for Windows, LFS is usually included
git lfs version

# If not installed, download from: https://git-lfs.github.io/
```

**macOS:**
```bash
brew install git-lfs
```

**Linux:**
```bash
sudo apt install git-lfs    # Ubuntu/Debian
sudo yum install git-lfs    # CentOS/RHEL
```

### **2. Initialize LFS in Your Repository**
```bash
# Navigate to your repository
cd your-repo

# Initialize Git LFS
git lfs install
```

## 📁 **Track Your Large File**

### **3. Add File Type to LFS Tracking**
```bash
# Track specific file extension (recommended)
git lfs track "*.vcf.gz"
git lfs track "*.fa.gz" 
git lfs track "*.bam"

# Or track specific file
git lfs track "your-large-file.vcf.gz"

# Or track files over certain size
git lfs track "*.zip"
```

### **4. Verify LFS Tracking**
```bash
# Check what's being tracked
git lfs track

# Check LFS status
git lfs ls-files
```

## 💾 **Add and Commit Your Large File**

### **5. Stage and Commit**
```bash
# Add the .gitattributes file (created by git lfs track)
git add .gitattributes

# Add your large file
git add your-large-file.vcf.gz

# Commit
git commit -m "Add large VCF file using LFS"

# Push to GitHub
git push origin main
```

## 🔍 **Complete Example Workflow**

```bash
# 1. Initialize LFS
git lfs install

# 2. Track large files by extension
git lfs track "*.vcf.gz"
git lfs track "*.fa.gz"
git lfs track "*.bam"

# 3. Add the tracking configuration
git add .gitattributes

# 4. Add your large file
git add chr22_variants.vcf.gz

# 5. Commit everything
git commit -m "Add LFS tracking and large genomics files"

# 6. Push to GitHub
git push origin main
```

## ✅ **Verify LFS is Working**

### **Check File Status:**
```bash
# Should show LFS files with pointer info
git lfs ls-files

# Check file is tracked by LFS
git lfs status
```

### **In GitHub Interface:**
- LFS files will show as "Git LFS" in the repository
- File size will show as "123 MB" instead of actual content

## 📋 **Common File Types for LFS**

**For Bioinformatics/Genomics:**
```bash
git lfs track "*.vcf.gz"
git lfs track "*.bam"
git lfs track "*.fastq.gz"
git lfs track "*.fa.gz"
git lfs track "*.sam"
git lfs track "*.bcf"
```

**General Large Files:**
```bash
git lfs track "*.zip"
git lfs track "*.tar.gz"
git lfs track "*.mp4"
git lfs track "*.pdf"
git lfs track "*.psd"
```

## 🔧 **Existing Repository with Large Files**

### **If File Already in Git:**
```bash
# Remove file from Git history but keep locally
git rm --cached your-large-file.vcf.gz

# Track with LFS
git lfs track "*.vcf.gz"

# Re-add the file
git add .gitattributes
git add your-large-file.vcf.gz

# Commit
git commit -m "Move large file to LFS"

# Push
git push origin main
```

## 💰 **GitHub LFS Pricing**

**Free Tier:**
- 1 GB storage
- 1 GB bandwidth per month

**Paid Storage:**
- $5/month for 50 GB storage + 50 GB bandwidth
- Additional packs available

## 🚨 **Important Notes**

### **File Size Limits:**
- **GitHub**: 100 MB hard limit without LFS
- **Git LFS**: Individual files up to 2 GB
- **Recommended**: Keep files under 1 GB for best performance

### **Best Practices:**
```bash
# Track by file type, not individual files
git lfs track "*.vcf.gz"  ✅
git lfs track "my-file.vcf.gz"  ❌

# Commit .gitattributes first
git add .gitattributes
git commit -m "Add LFS tracking"
```

### **Troubleshooting:**
```bash
# If push fails due to file size
git lfs migrate import --include="*.vcf.gz"

# Check LFS configuration
cat .gitattributes

# Verify LFS is working
git lfs env
```

## 📝 **Your .gitattributes File**

After running `git lfs track`, you'll see:
```
*.vcf.gz filter=lfs diff=lfs merge=lfs -text
*.fa.gz filter=lfs diff=lfs merge=lfs -text
*.bam filter=lfs diff=lfs merge=lfs -text
```

## 🎯 **Quick Checklist**

- [ ] Install Git LFS
- [ ] Run `git lfs install`
- [ ] Track file types with `git lfs track "*.extension"`
- [ ] Add `.gitattributes` file
- [ ] Add and commit your large file
- [ ] Push to GitHub
- [ ] Verify files show as "Git LFS" on GitHub

This will handle your 100+ MB file efficiently and keep your repository performant! 🚀

## 📍 **.gitattributes Should Be at the Repository Root**

### **Correct Location:**
```
your-repository/
├── .gitattributes          ← HERE (at root)
├── .git/
├── README.md
├── src/
│   ├── code files...
│   └── large-files/
│       ├── COMPUTER_BASICS-8.25.2025.exe
│       └── data.vcf.gz
└── other folders...
```

## 🎯 **Why at the Root?**

### **Git Attributes Inheritance:**
- ✅ **Root .gitattributes** → Applies to **entire repository**
- ✅ **Subfolder .gitattributes** → Applies only to **that folder and subfolders**
- ✅ **Git LFS needs global tracking** → Must be at root

### **Example Structure:**
```bash
# Repository layout
my-project/
├── .gitattributes                    # ← Main LFS configuration
├── docs/
├── src/
├── data/
│   ├── large-files/
│   │   ├── genome.fa.gz             # Tracked by root .gitattributes
│   │   └── variants.vcf.gz          # Tracked by root .gitattributes
│   └── small-files/
│       └── config.txt
└── executables/
    ├── COMPUTER_BASICS-8.25.2025.exe  # Tracked by root .gitattributes
    └── small-app.exe
```

## 📝 **What Goes in Root .gitattributes:**

```bash
# Global LFS tracking rules
*.exe filter=lfs diff=lfs merge=lfs -text
*COMPUTER_BASICS*.exe filter=lfs diff=lfs merge=lfs -text
*.vcf.gz filter=lfs diff=lfs merge=lfs -text
*.fa.gz filter=lfs diff=lfs merge=lfs -text
*.bam filter=lfs diff=lfs merge=lfs -text

# File-specific rules (if needed)
data/large-dataset.zip filter=lfs diff=lfs merge=lfs -text
```

## 🔄 **Can You Use Multiple .gitattributes?**

### **Yes, but with hierarchy:**

```
project/
├── .gitattributes              # Global rules
├── data/
│   ├── .gitattributes         # Additional rules for data/ folder
│   ├── genomics/
│   │   ├── .gitattributes     # Specific rules for genomics/
│   │   └── large-genome.fa.gz
│   └── results/
└── src/
```

### **Hierarchy Example:**
```bash
# Root .gitattributes
*.exe filter=lfs diff=lfs merge=lfs -text

# data/.gitattributes (additional rules)
*.vcf.gz filter=lfs diff=lfs merge=lfs -text
*.bam filter=lfs diff=lfs merge=lfs -text

# data/genomics/.gitattributes (more specific)
*.fa filter=lfs diff=lfs merge=lfs -text  # Even uncompressed FASTA files
```

## 🚀 **Best Practice Setup**

### **Step 1: Create at Repository Root**
```bash
# Navigate to repository root (where .git folder is)
cd your-repository

# Verify you're at root
ls -la | grep ".git"

# Initialize LFS and create .gitattributes at root
git lfs install
git lfs track "*.exe"
git lfs track "*COMPUTER_BASICS*.exe"
```

### **Step 2: Verify Location**
```bash
# Check .gitattributes is at root
ls -la .gitattributes

# Check contents
cat .gitattributes
```

### **Step 3: Should Look Like:**
```bash
# Contents of .gitattributes at root
*.exe filter=lfs diff=lfs merge=lfs -text
*COMPUTER_BASICS*.exe filter=lfs diff=lfs merge=lfs -text
```

## ⚠️ **Common Mistakes**

### **❌ Wrong:**
```
project/
├── executables/
│   ├── .gitattributes          # ← Only affects this folder
│   └── large-app.exe
└── data/
    ├── .gitattributes          # ← Only affects this folder  
    └── large-data.gz
```

### **✅ Correct:**
```
project/
├── .gitattributes              # ← Affects entire repository
├── executables/
│   └── large-app.exe          # ← Tracked by root .gitattributes
└── data/
    └── large-data.gz          # ← Tracked by root .gitattributes
```

## 🎯 **For Your Specific Case**

```bash
# Navigate to repository root
cd your-repository

# Create/edit .gitattributes at root
git lfs track "COMPUTER_BASICS-*.exe"

# Verify it's at root level
ls -la .gitattributes

# Add and commit
git add .gitattributes
git commit -m "Add LFS tracking configuration"

# Now add your large files from anywhere in the repository
git add executables/COMPUTER_BASICS-8.25.2025.exe
git add data/large-file.gz
git commit -m "Add large files via LFS"
```

## 💡 **Key Points**

✅ **Always start with .gitattributes at repository root**  
✅ **It will track patterns anywhere in the repository**  
✅ **You can add folder-specific .gitattributes later if needed**  
✅ **Git LFS works best with global configuration at root**  

The root location ensures your LFS tracking works for large files **anywhere** in your repository! 🚀