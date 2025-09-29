---
title: "git - revert, reset, undo last, second to last commits"
description: "Brief_description_Fildsfe_adfdf_adsfdsf"
publishDate: 2025-09-29
tags: ["git", "revert", "reset"]
draft: false
---
Here's a comprehensive guide to reverting commits in Git:

## 🔄 **Git Revert vs Reset vs Checkout**

### **Key Differences:**
- **`git revert`**: Creates NEW commit that undoes changes (safe, keeps history)
- **`git reset`**: Moves HEAD back (rewrites history - dangerous if pushed)
- **`git checkout`**: Switches branches or restores files (non-destructive)

---

## 🎯 **Git Revert (Recommended - Safe)**

### **Revert Last Commit**
```bash
# Revert the most recent commit (creates new commit)
git revert HEAD

# With commit message editor
git revert HEAD

# Skip commit message editor
git revert HEAD --no-edit

# Revert but don't commit yet (stage changes)
git revert HEAD --no-commit
```

### **Revert Second to Last Commit**
```bash
# Revert the commit before the last one
git revert HEAD~1

# Or use commit hash
git revert a1b2c3d
```

### **Revert Multiple Commits**
```bash
# Revert last 3 commits (creates 3 new commits)
git revert HEAD HEAD~1 HEAD~2

# Revert range of commits
git revert HEAD~2..HEAD

# Revert multiple but create single commit
git revert --no-commit HEAD~2..HEAD
git commit -m "Revert last 3 commits"
```

---

## ⚡ **Git Reset (Destructive - Use Carefully)**

### **Reset to Last Commit (Undo Changes)**
```bash
# Soft reset - keep changes in staging area
git reset --soft HEAD

# Mixed reset - keep changes in working directory (default)
git reset HEAD
# or
git reset --mixed HEAD

# Hard reset - DELETE all changes (DANGEROUS!)
git reset --hard HEAD
```

### **Reset to Second to Last Commit**
```bash
# Soft - keep changes staged
git reset --soft HEAD~1

# Mixed - keep changes unstaged (default)
git reset HEAD~1

# Hard - delete all changes (DANGEROUS!)
git reset --hard HEAD~1
```

### **Reset to Specific Commit**
```bash
# Reset to specific commit by hash
git reset --hard a1b2c3d

# Reset to specific commit, keep changes
git reset --soft a1b2c3d
```

---

## 📋 **Complete Examples**

### **Example 1: Revert Last Commit (Safe)**
```bash
# Current history:
# C (HEAD) - "Add feature X"
# B - "Update README"
# A - "Initial commit"

# Revert last commit
git revert HEAD

# New history:
# D (HEAD) - "Revert 'Add feature X'"
# C - "Add feature X"
# B - "Update README"  
# A - "Initial commit"
```

### **Example 2: Revert Second to Last Commit**
```bash
# Current history:
# C (HEAD) - "Fix bug"
# B - "Add broken feature"  ← We want to revert this
# A - "Initial commit"

# Revert second to last commit
git revert HEAD~1

# New history:
# D (HEAD) - "Revert 'Add broken feature'"
# C - "Fix bug"
# B - "Add broken feature"
# A - "Initial commit"
```

### **Example 3: Reset to Last Commit (Undo Local Changes)**
```bash
# You made changes but haven't committed
echo "bad changes" >> file.txt
git add file.txt

# Undo everything - back to last commit
git reset --hard HEAD

# file.txt is back to last committed state
```

### **Example 4: Reset to Second to Last Commit**
```bash
# Current history:
# C (HEAD) - "Oops, bad commit"  ← Remove this
# B - "Good commit"  ← Go back here
# A - "Initial commit"

# Remove last commit, keep changes
git reset --soft HEAD~1

# Or remove last commit and all changes
git reset --hard HEAD~1

# New history:
# B (HEAD) - "Good commit"
# A - "Initial commit"
```

---

## 🎭 **Visual Reference Guide**

### **Revert (Safe - Adds New Commit):**
```
Before:
A---B---C (HEAD)

After git revert HEAD:
A---B---C---D (HEAD)
            └─ Undoes C
```

### **Reset (Dangerous - Moves HEAD):**
```
Before:
A---B---C (HEAD)

After git reset --hard HEAD~1:
A---B (HEAD)
    └─ C is gone!
```

---

## 🔍 **Checking Commit History**

### **View Commits**
```bash
# See last 5 commits with hashes
git log --oneline -5

# Output example:
# a1b2c3d (HEAD) Fix bug
# e4f5g6h Add feature
# i7j8k9l Update docs
# m1n2o3p Initial commit
```

### **Reference Commits**
```bash
HEAD      # Current commit
HEAD~1    # 1 commit before HEAD (second to last)
HEAD~2    # 2 commits before HEAD (third to last)
HEAD~3    # 3 commits before HEAD
HEAD^     # Same as HEAD~1
HEAD^^    # Same as HEAD~2

# By hash
a1b2c3d   # Specific commit
```

---

## 🛠️ **Practical Scenarios**

### **Scenario 1: Undo Last Commit (Not Pushed)**
```bash
# Made a commit but want to modify it
git reset --soft HEAD~1
# Make changes
git add .
git commit -m "Better commit message"
```

### **Scenario 2: Undo Last Commit (Already Pushed)**
```bash
# Safe way - use revert
git revert HEAD
git push origin main

# Dangerous way - force push (only if no one else pulled)
git reset --hard HEAD~1
git push --force origin main
```

### **Scenario 3: Undo Multiple Commits**
```bash
# Undo last 3 commits safely
git revert HEAD HEAD~1 HEAD~2
git push origin main

# Or create single revert commit
git revert --no-commit HEAD HEAD~1 HEAD~2
git commit -m "Revert last 3 commits"
git push origin main
```

### **Scenario 4: Temporarily Go Back in Time**
```bash
# Look at second to last commit (doesn't change anything)
git checkout HEAD~1

# Look around, test things...

# Return to latest commit
git checkout main
```

---

## 🚨 **Safety Checklist**

### **Before Using Reset:**
```bash
# 1. Check if commits are pushed
git log --oneline origin/main..HEAD

# 2. Create backup branch
git branch backup-branch

# 3. Then safely reset
git reset --hard HEAD~1

# If you mess up, restore from backup:
git reset --hard backup-branch
```

### **Safe Revert Workflow:**
```bash
# Always safe to use revert
git revert HEAD          # Revert last commit
git revert HEAD~1        # Revert second to last
git revert a1b2c3d       # Revert specific commit
```

---

## 📊 **Quick Reference Table**

| Command | What it Does | Safe? | Use When |
|---------|--------------|-------|----------|
| `git revert HEAD` | Undo last commit (new commit) | ✅ Yes | Commits already pushed |
| `git revert HEAD~1` | Undo 2nd to last (new commit) | ✅ Yes | Commits already pushed |
| `git reset --soft HEAD~1` | Undo commit, keep changes staged | ⚠️ Careful | Before pushing |
| `git reset --mixed HEAD~1` | Undo commit, keep changes unstaged | ⚠️ Careful | Before pushing |
| `git reset --hard HEAD~1` | Undo commit, DELETE changes | ❌ Dangerous | Local work only |

---

## 💡 **Golden Rules**

✅ **Use `git revert`** for published commits (pushed to remote)  
✅ **Use `git reset`** for local commits only (not pushed)  
✅ **Always backup** before using `--hard`  
✅ **Check history** with `git log` before reverting  
⚠️ **Never force push** to shared branches without team agreement  

These commands give you full control over your commit history! 🚀