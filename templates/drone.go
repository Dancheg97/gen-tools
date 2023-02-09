package templates

import (
	"strings"

	"dancheg97.ru/templates/gen-tools/templates/utils"
)

func GenerateDroneYml(giteaurl string) {
	utils.WriteFile(".drone.yml", strings.ReplaceAll(DroneYml, `DMN`, giteaurl))
}

const DroneYml = `kind: pipeline
name: default
type: docker

trigger:
  branch:
    - main

volumes:
  - name: docker
    host:
      path: /var/run/docker.sock

steps:
  - name: build
    image: docker
    volumes:
      - name: docker
        path: /var/run/docker.sock
    commands:
      - docker build -t DMN/$DRONE_REPO:latest -t DMN/$DRONE_REPO:$(date +"%m-%d-%y") .

  - name: push
    image: docker
    environment:
      PASS:
        from_secret: PASS
    volumes:
      - name: docker
        path: /var/run/docker.sock
    commands:
      - docker login -u login -p $PASS DMN
      - docker push DMN/$DRONE_REPO:latest
      - docker push DMN/$DRONE_REPO:$(date +"%m-%d-%y")
`
