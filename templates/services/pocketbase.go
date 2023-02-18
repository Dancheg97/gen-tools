package services

import (
	"fmt"

	"dancheg97.ru/dancheg97/gen-tools/utils"
)

func GeneratePocketbase(mail string, domain string) {
	utils.AppendToCompose(PocketbaseYaml)
	utils.AppendToNginx(fmt.Sprintf(PocketbaseNginx, domain, domain, domain))
	utils.AppendToCerts(mail, "pocketbase."+domain)
}

const PocketbaseYaml = `
  pocketbase:
    image: ghcr.io/muchobien/pocketbase:latest
    restart: unless-stopped
    volumes:
      - ./pocketbase:/pb_data
    healthcheck:
      test: wget --no-verbose --tries=1 --spider http://localhost:8090/api/health || exit 1
      interval: 5s
      timeout: 5s
      retries: 5

`

const PocketbaseNginx = `
server {
    listen 80;
    listen 443 ssl;
    server_name pocketbase.%s;
    ssl_certificate /certs/pocketbase.%s.crt;
    ssl_certificate_key /certs/pocketbase.%s.key;
    location / {
        proxy_pass http://pocketbase:8090/;
    }
}
`
