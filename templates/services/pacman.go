package services

import (
	"fmt"

	"dancheg97.ru/dancheg97/gen-tools/utils"
)

func GeneratePacman(mail string, domain string) {
	utils.AppendToCompose(fmt.Sprintf(PacmanYaml, domain))
	utils.AppendToNginx(fmt.Sprintf(PacmanNginx, domain, domain, domain))
	utils.AppendToCerts(mail, "pacman."+domain)
}

const PacmanYaml = `
  pacman:
    image: gitea.dancheg97.ru/dancheg97/go-pacman:latest
    container_name: pacman
    command: run
    volumes:
      - ./go-pacman:/var/cache/pacman/pkg
    environment:
      REPO: %s
      INIT_PKGS: ttf-droid adw-gtk-theme yay
    ports:
      - 9080:9080

`

const PacmanNginx = `
server {
    listen 80;
    listen 443 ssl;
    server_name pacman.%s;
    ssl_certificate /certs/pacman.%s.crt;
    ssl_certificate_key /certs/pacman.%s.key;
    client_max_body_size 1500M;
    location / {
        proxy_pass http://pacman:8080/;
    }
}

`
