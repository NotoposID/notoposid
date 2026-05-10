# Deployment Guide - VPS Ubuntu

This guide covers deploying NOTOPOS AI to a fresh Ubuntu 22.04+ VPS.

## 1. Initial Server Setup

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install Docker Compose
sudo apt install docker-compose-plugin -y
```

## 2. Environment Configuration

Create the `.env` files for both apps based on the examples provided in the repository.

```bash
# Backend .env
PORT=8000
DATABASE_URL=postgres://user:pass@db:5432/notopos
REDIS_URL=redis:6379
JWT_SECRET=your_jwt_secret
OPENAI_API_KEY=your_openai_key
```

## 3. Deploy with Docker Compose

```bash
# Clone the repo
git clone https://github.com/your-org/notopos.git
cd notopos

# Build and start services
docker compose -f docker-compose.yml up -d --build
```

## 4. Nginx Reverse Proxy Setup

Install Nginx:
```bash
sudo apt install nginx -y
```

Create a configuration file `/etc/nginx/sites-available/notopos`:

```nginx
server {
    listen 80;
    server_name api.yourdomain.com;

    location / {
        proxy_pass http://localhost:8000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}

server {
    listen 80;
    server_name yourdomain.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

Enable the site and restart Nginx:
```bash
sudo ln -s /etc/nginx/sites-available/notopos /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

## 5. SSL with Certbot

```bash
sudo apt install certbot python3-certbot-nginx -y
sudo certbot --nginx -d yourdomain.com -d api.yourdomain.com
```

## 6. Monitoring & Logs

View logs:
```bash
docker compose logs -f api
```
