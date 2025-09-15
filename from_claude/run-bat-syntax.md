Here's a batch file that checks if ports are listening before running the corresponding scripts:**Alternative version with synchronous execution** (waits for each service to start):**Key differences between the versions:**

**Version 1 (Recommended):**
- Uses `start` command to run batch files in **background/parallel**
- Services start simultaneously and don't block each other
- Uses `/min` to minimize windows
- Better for web servers or long-running services

**Version 2 (Synchronous):**
- Uses `call` command to run batch files **sequentially**
- Waits for each script to complete before moving to the next
- Better if scripts need to run in specific order

**How the port checking works:**
- `netstat -an` lists all network connections
- `findstr ":4200 "` searches for the specific port
- `findstr "LISTENING"` confirms the port is actively listening
- `>nul 2>&1` suppresses output
- `errorlevel` is 0 if port is found (listening), 1 if not found (available)

**Usage:**
1. Save as `run-4200-22025.bat`
2. Place in same directory as `run-4200.bat` and `run-22025.bat`
3. Run it - it will automatically check ports and start only the needed services

This prevents conflicts and duplicate services while ensuring both ports are covered.