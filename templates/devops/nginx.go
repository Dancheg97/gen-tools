package devops

import "dancheg97.ru/templates/gen-tools/templates/utils"

func GenerateNginx() {
	utils.AppendToCompose(NginxYaml)
}

const NginxYaml = `
  nginx:
    image: nginx:1.23-alpine
    container_name: nginx
    restart: unless-stopped
    ports:
      - 80:80
      - 443:443
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/conf.d/nginx.conf:ro
      - ./.lego/certificates:/certs:ro

`
