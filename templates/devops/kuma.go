package devops

import (
	"fmt"

	"dancheg97.ru/templates/gen-tools/templates/utils"
)

func GenerateUptimeKuma(mail string, domain string) {
	utils.AppendToCompose(UptimeKumaCompose)
	utils.AppendToNginx(fmt.Sprintf(UptimeKumaNginx, domain, domain, domain))
	utils.AppendToCerts(mail, "kuma."+domain)
}

const UptimeKumaCompose = `
  kuma:
    image: louislam/uptime-kuma:1
    container_name: kuma
    volumes:
      - ./kuma:/app/data
`

const UptimeKumaNginx = `
server {
    listen 80;
    listen 443 ssl;
    server_name kuma.%s;
    ssl_certificate /certs/kuma.%s.crt;
    ssl_certificate_key /certs/kuma.%s.key;
    location / {
        proxy_pass http://kuma:3001/;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $host;
    }
}
`
