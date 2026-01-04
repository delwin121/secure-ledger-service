# Firewall Rules (UFW)

To secure your deployment and ensure Grafana is only accessible via the Nginx proxy (which handles authentication), follow these steps.

**⚠️ WARNING: incorrect firewall rules can lock you out of your server (SSH). Ensure you allow SSH first!**

## 1. Install UFW
```bash
sudo apt-get install ufw  # Debian/Ubuntu
# sudo pacman -S ufw      # Arch
```

## 2. Default Policies
Deny incoming by default, allow outgoing.
```bash
sudo ufw default deny incoming
sudo ufw default allow outgoing
```

## 3. Allow SSH (Critical)
```bash
sudo ufw allow ssh
# OR specific port if changed
# sudo ufw allow 2222/tcp
```

## 4. Allow Nginx Proxy
Allow traffic to the secure proxy port (8081).
```bash
sudo ufw allow 8081/tcp
```

## 5. Allow Kubernetes/App Ports (Internal/Local only)
If you need to access the Go app (8080) or Grafana (3000) directly from `localhost` but block external:
```bash
sudo ufw allow from 127.0.0.1 to any port 8080
sudo ufw allow from 127.0.0.1 to any port 3000
```

## 6. Enable Firewall
```bash
sudo ufw enable
```

## Verification
Check status:
```bash
sudo ufw status verbose
```
