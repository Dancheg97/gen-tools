package templates

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
      - docker build -t gitea.example.org/$DRONE_REPO:latest -t gitea.example.org/$DRONE_REPO:$(date +"%m-%d-%y") .

  - name: push
    image: docker
    environment:
      PASS:
        from_secret: PASS
    volumes:
      - name: docker
        path: /var/run/docker.sock
    commands:
      - docker login -u login -p $PASS gitea.example.org
      - docker push gitea.example.org/$DRONE_REPO:latest
      - docker push gitea.example.org/$DRONE_REPO:$(date +"%m-%d-%y")
`
