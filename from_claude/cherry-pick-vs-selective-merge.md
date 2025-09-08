# Cherry-pick when some files from advanced history need to be back in history in new branch

## Cherry-pick Strategy - Most Precise Method
```sh
# 1. First, identify the commits that added weeks 13-19
git log --oneline development --grep="week1[3-9]" --grep="Section2" --grep="javascript"

# 2. Or look at commits by examining the files
git log --oneline development -- lessons/section2-javascript/

# 3. Switch to the target branch (addfour)
git checkout addfour

# 4. Cherry-pick the specific commits (replace with actual commit hashes)
# Based on your log, these seem to be the relevant commits:
git cherry-pick 2e3b42b  # "Section2 javascript added week13,14,15,16"
git cherry-pick d50527c  # "Section2 JS week17, 18,19 probably incomplete added"

# 5. If there are conflicts, resolve them:
git status  # Check for conflicts
# Edit conflicted files, then:
git add .
git cherry-pick --continue

# 6. Verify the files were added correctly
ls lessons/section2-javascript/
git log --oneline -5  # Check recent commits
```

## Manual File Copy Strategy - Safest Method
```sh
# 1. First, create a temporary backup of the files you want
mkdir /tmp/js_weeks_backup
git checkout development
cp lessons/section2-javascript/week1[3-9].md /tmp/js_weeks_backup/ 2>/dev/null || true

# 2. List what files actually exist
ls lessons/section2-javascript/week*.md

# 3. Copy the specific files you want
cp lessons/section2-javascript/week13.md /tmp/js_weeks_backup/
cp lessons/section2-javascript/week14.md /tmp/js_weeks_backup/
cp lessons/section2-javascript/week15.md /tmp/js_weeks_backup/
cp lessons/section2-javascript/week16.md /tmp/js_weeks_backup/
cp lessons/section2-javascript/week17.md /tmp/js_weeks_backup/
cp lessons/section2-javascript/week18.md /tmp/js_weeks_backup/
cp lessons/section2-javascript/week19.md /tmp/js_weeks_backup/

# 4. Switch to target branch
git checkout addfour

# 5. Create the directory if it doesn't exist
mkdir -p lessons/section2-javascript

# 6. Copy the files to the target branch
cp /tmp/js_weeks_backup/week*.md lessons/section2-javascript/

# 7. Add and commit the new files
git add lessons/section2-javascript/week1[3-9].md
git status  # Verify what you're about to commit
git commit -m "Add JavaScript weeks 13-19 from development branch"

# 8. Clean up temporary files
rm -rf /tmp/js_weeks_backup
```


## Selective Merge Strategy - Advanced Method
```sh
# 1. Switch to target branch
git checkout addfour

# 2. Start a merge but don't commit
git merge --no-ff --no-commit development

# 3. Reset everything except the files you want
git reset HEAD

# 4. Add only the specific JavaScript week files
git add lessons/section2-javascript/week13.md
git add lessons/section2-javascript/week14.md
git add lessons/section2-javascript/week15.md
git add lessons/section2-javascript/week16.md
git add lessons/section2-javascript/week17.md
git add lessons/section2-javascript/week18.md
git add lessons/section2-javascript/week19.md

# 5. Check what you're about to commit
git status
git diff --cached

# 6. Commit only these specific files
git commit -m "Add JavaScript weeks 13-19 from development branch"

# 7. Clean up any remaining merge state
git merge --abort 2>/dev/null || true
```


## Clean Up Messy Merge State - Step by Step
```sh
# 1. First, check your current status
git status

# 2. If there are staged files you DON'T want to commit:
git reset HEAD .

# 3. If there are untracked files you want to remove:
git clean -f    # Remove untracked files
git clean -fd   # Remove untracked files and directories
git clean -fx   # Remove untracked AND ignored files (be careful!)

# 4. If there are modified files you want to discard:
git checkout -- .

# 5. If you want to keep ONLY specific files (the JavaScript weeks):
# First, stash everything
git stash push -u -m "backup before cleanup"

# Then selectively restore just what you want
git stash show -p | grep "lessons/section2-javascript/week1[3-9].md" | git apply

# 6. Or the nuclear option - reset everything to clean state:
git reset --hard HEAD
git clean -fd

# 7. Verify you're back to a clean state:
git status

# Should show: "nothing to commit, working tree clean"
```
## After cleaning up the messy merge, use this SAFE method
```sh
# 1. Make sure you're on a clean addfour branch
git checkout addfour
git status  # Should be clean now

# 2. Create a backup branch (just in case)
git branch addfour-backup

# 3. Use the manual file copy method (safest)
# First, extract the files from development branch to temp location
git show development:lessons/section2-javascript/week13.md > /tmp/week13.md 2>/dev/null || echo "week13.md not found"
git show development:lessons/section2-javascript/week14.md > /tmp/week14.md 2>/dev/null || echo "week14.md not found"
git show development:lessons/section2-javascript/week15.md > /tmp/week15.md 2>/dev/null || echo "week15.md not found"
git show development:lessons/section2-javascript/week16.md > /tmp/week16.md 2>/dev/null || echo "week16.md not found"
git show development:lessons/section2-javascript/week17.md > /tmp/week17.md 2>/dev/null || echo "week17.md not found"
git show development:lessons/section2-javascript/week18.md > /tmp/week18.md 2>/dev/null || echo "week18.md not found"
git show development:lessons/section2-javascript/week19.md > /tmp/week19.md 2>/dev/null || echo "week19.md not found"

# 4. Create the directory structure if needed
mkdir -p lessons/section2-javascript

# 5. Copy only the files that actually exist
for week in 13 14 15 16 17 18 19; do
  if [ -f "/tmp/week${week}.md" ] && [ -s "/tmp/week${week}.md" ]; then
    cp "/tmp/week${week}.md" "lessons/section2-javascript/"
    echo "Copied week${week}.md"
  else
    echo "week${week}.md does not exist or is empty"
  fi
done

# 6. Check what was actually copied
ls -la lessons/section2-javascript/

# 7. Add and commit only the files that were successfully copied
git add lessons/section2-javascript/week*.md
git status  # Review what you're about to commit
git commit -m "Add JavaScript weeks 13-19 from development branch"

# 8. Clean up temp files
rm -f /tmp/week*.md

# 9. Verify the result
git log --oneline -3
git status
```