For long-term use on macOS, I recommend **Docker** for these reasons:

**Why Docker is best for long-term:**
- **Reliability:** Containers are isolated and consistent
- **Easy updates:** Just pull new image versions
- **No system conflicts:** Won't interfere with macOS updates
- **Easy backup/restore:** Container volumes can be easily backed up
- **Portability:** If you switch computers, just move the container
- **Automatic restarts:** Can configure to restart after system reboots

**Docker setup for long-term use:**
```bash
# Create persistent data directory
mkdir ~/mumble-data

# Run with persistent storage and auto-restart
docker run -d \
  --name mumble-server \
  --restart unless-stopped \
  -p 64738:64738/tcp \
  -p 64738:64738/udp \
  -v ~/mumble-data:/data \
  mumblevoip/mumble-server
```

**Set admin password:**
```bash
docker exec -it mumble-server murmurd -supw YOUR_PASSWORD
```

**To start on boot:**
Docker Desktop can be set to start automatically, and the `--restart unless-stopped` flag will restart your Mumble server.

**Alternative for maximum simplicity:** If Docker seems too complex, the **official .dmg installer** is also fine for long-term use, but you'll need to manually manage updates and ensure it starts on boot.

For a family setup that "just works" long-term, Docker is worth the small initial learning curve.


You need to start Docker first. Here's how to fix this:

**Install and Start Docker Desktop:**

1. **Install Docker Desktop:**
   ```bash
   brew install --cask docker
   ```

2. **Start Docker Desktop:**
   - Open Applications folder and launch "Docker"
   - Or from Terminal: `open -a Docker`
   - You'll see a whale icon in your menu bar when it's running

3. **Wait for Docker to fully start:**
   - The whale icon will stop animating when ready
   - This can take 1-2 minutes on first launch

4. **Verify it's working:**
   ```bash
   docker --version
   docker ps
   ```

**Then run your Mumble server:**
```bash
mkdir ~/mumble-data

docker run -d \
  --name mumble-server \
  --restart unless-stopped \
  -p 64738:64738/tcp \
  -p 64738:64738/udp \
  -v ~/mumble-data:/data \
  mumblevoip/mumble-server
```

**Auto-start Docker:**
- In Docker Desktop preferences, check "Start Docker Desktop when you log in"

**Alternative if you don't want Docker Desktop:**
You could use the native macOS installer from mumble.info instead - it's simpler but requires more manual management for long-term use.

Which approach would you prefer?