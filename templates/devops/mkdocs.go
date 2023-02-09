package devops

import (
	"fmt"

	"dancheg97.ru/templates/gen-tools/templates/utils"
)

func GenerateMkdocs(mail string, domain string) {
	utils.AppendToCompose(MkDocsCompose)
	utils.WriteFile(`mkdocs/mkdocs.yml`, MkDocsConfigYaml)
	utils.WriteFile(`mkdocs/docs/stylesheets/extra.css`, MkDocsCss)
	utils.AppendToNginx(fmt.Sprintf(MkDocsNginx, domain, domain, domain))
	utils.AppendToCerts(mail, "docs."+domain)
}

const MkDocsCompose = `
  docs:
    image: squidfunk/mkdocs-material
    container_name: docs
    volumes:
      - ./mkdocs:/docs

`

const MkDocsNginx = `
server {
    listen 80;
    listen 443 ssl;
    server_name docs.%s;
    ssl_certificate /certs/docs.%s.crt;
    ssl_certificate_key /certs/docs.%s.key;
    location / {
        proxy_pass http://docs:8000/;
    }
}
`

const MkDocsConfigYaml = `site_name: Example name

repo_url: https://gitrepo.com/example
repo_name: RepoName

theme:
  name: material
  logo: assets/logo.png

  icon:
    edit: material/pencil
    view: material/eye

  features:
    - announce.dismiss
    - content.action.edit
    - git-revision-date-localized:
      enable_creation_date: true

  palette:
    scheme: slate
    accent: light blue

plugins:
  - search

extra:
  alternate:
    - name: English
      link: /en/
      lang: en
    - name: Russian
      link: /ru/
      lang: ru

extra_css:
  - stylesheets/extra.css
`

const MkDocsCss = `:root {
--md-primary-fg-color: #1d1f23;
--md-sidebar-primary: #1d1f23;
}

[data-md-color-scheme="slate"] {
--md-typeset-a-color: #1f7295;
}

body {
background: #24262b;
}

.md-nav--secondary .md-nav__title {
background: #24262b;
box-shadow: 0 0 0.4rem 0.4rem #24262b;
position: -webkit-sticky;
position: sticky;
top: 0;
z-index: 1;
}

.md-nav--primary .md-nav__title {
background: #24262b;
box-shadow: 0 0 0.4rem 0.4rem #24262b;
position: -webkit-sticky;
position: sticky;
top: 0;
z-index: 1;
}

.md-nav__link {
color: white;
align-items: center;
cursor: pointer;
display: flex;
justify-content: space-between;
margin-top: 0.625em;
overflow: hidden;
scroll-snap-align: start;
text-overflow: ellipsis;
transition: color 125ms;
}

.md-nav__link md-nav__link--active {
color: white;
align-items: center;
cursor: pointer;
display: flex;
justify-content: space-between;
margin-top: 0.625em;
overflow: hidden;
scroll-snap-align: start;
text-overflow: ellipsis;
transition: color 125ms;
}`
