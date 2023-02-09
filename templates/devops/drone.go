package devops

import (
	"fmt"

	"dancheg97.ru/templates/gen-tools/templates/utils"
)

func GenerateDrone(mail string, domain string) {
	utils.AppendToCompose(fmt.Sprintf(DroneYaml, domain, domain))
	utils.AppendToNginx(fmt.Sprintf(DroneNginx, domain, domain, domain))
	utils.AppendToCerts(mail, "drone."+domain)
}

const DroneYaml = `  
  drone:
    image: drone/drone:2.15
    container_name: drone
    restart: unless-stopped
    environment:
      DRONE_DATABASE_DRIVER: sqlite3
      DRONE_DATABASE_DATASOURCE: /data/database.sqlite
      DRONE_GITEA_SERVER: https://gitea.%s/
      DRONE_GIT_ALWAYS_AUTH: false
      DRONE_RPC_SECRET: ce126d3fb16f10fe78073199630f158b
      DRONE_SERVER_PROTO: https
      DRONE_SERVER_HOST: drone.%s
      DRONE_TLS_AUTOCERT: false
      DRONE_GITEA_CLIENT_ID: 6c582a78-65df-4d84-a9e4-aacaa25d222a
      DRONE_GITEA_CLIENT_SECRET: gto_2xmk3zney35mleko5liwpjc74s2h7r5a6ngiyo63jtfcezkza5ra
    volumes:
      - ./drone:/data
      - /var/run/docker.sock:/var/run/docker.sock
    
  drone-runner:
    image: drone/drone-runner-docker:1.8.2
    container_name: droner
    restart: unless-stopped
    environment:
      DRONE_RPC_PROTO: http
      DRONE_RPC_HOST: drone
      DRONE_RPC_SECRET: ce126d3fb16f10fe78073199630f158b
      DRONE_RUNNER_NAME: drone-runner
      DRONE_RUNNER_CAPACITY: 2
      DRONE_RUNNER_NETWORKS: composer_default
      DRONE_DEBUG: false
      DRONE_TRACE: false
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    depends_on:
      - drone

`
const DroneNginx = `
server {
    listen 80;
    listen 443 ssl;
    server_name drone.%s;
    ssl_certificate /certs/drone.%s.crt;
    ssl_certificate_key /certs/drone.%s.key;
    location / {
        proxy_pass http://drone/;
    }
}
`
