package devops

const NginxYaml = `  nginx:
    image: nginx:1.23-alpine
    container_name: nginx
    restart: unless-stopped
    ports:
      - 80:80
      - 443:443
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/conf.d/nginx.conf:ro
      - ./.lego/certificates:/certs:ro
    depends_on:
      - gitea
      - drone
      - pocketbase
`

const NginxConf = `
server {
    listen 80;
    listen 443 ssl;
    server_name example.com;
    ssl_certificate /certs/example.com.crt;
    ssl_certificate_key /certs/example.com.key;
    client_max_body_size 500M;
    location / {
        proxy_pass http://gitea/;
    }
}

server {
    listen 80;
    listen 443 ssl;
    server_name sub.example.com;
    ssl_certificate /certs/sub.example.com.crt;
    ssl_certificate_key /certs/sub.example.com.key;
    location / {
        proxy_pass http://sub/;
    }
}

`

const LegoSh = `go install github.com/go-acme/lego/v4/cmd/lego@latest
sudo lego --email="user@example.com" --domains="domain.com" --http run
sudo lego --email="user@example.com" --domains="domain.com" --http run
sudo lego --email="user@example.com" --domains="domain.com" --http run
sudo chown -R user:user .lego
`
