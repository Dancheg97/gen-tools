package devops

const PacmanYaml = `  pacman:
image: gitea.dancheg97.ru/dancheg97/go-pacman:latest
container_name: pacman
command: run
volumes:
  - ./go-pacman:/var/cache/pacman/pkg
environment:
  REPO: pacman.dancheg97.ru
  INIT_PKGS: ttf-droid adw-gtk-theme gnome-browser-connector onlyoffice-bin
ports:
  - 9080:9080
`
