package devops

import (
	"fmt"

	"dancheg97.ru/templates/gen-tools/templates/utils"
)

func GenerateDozzle(mail string, domain string, user string, pass string) {
	utils.AppendToCompose(fmt.Sprintf(DozzleYaml, user, pass))
	utils.AppendToNginx(fmt.Sprintf(DozzleNginx, domain, domain, domain))
	utils.AppendToCerts(mail, "drone."+domain)
}

const DozzleNginx = `
server {
    listen 80;
    listen 443 ssl;
    server_name dozzle.%s;
    ssl_certificate /certs/dozzle.%s.crt;
    ssl_certificate_key /certs/dozzle.%s.key;
    location / {
        proxy_pass http://dozzle:8080/;
    }
}

`

const DozzleYaml = `
  dozzle:
    container_name: dozzle
    image: amir20/dozzle:latest
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      DOZZLE_USERNAME: %s
      DOZZLE_PASSWORD: %s

`
