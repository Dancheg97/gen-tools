package services

import (
	"fmt"

	"dancheg97.ru/dancheg97/gen-tools/img"
	"dancheg97.ru/dancheg97/gen-tools/utils"
)

func GenerateMkdocs(mail string, domain string, logo string) {
	utils.AppendToCompose(MkDocsCompose)
	utils.WriteFile(`mkdocs/mkdocs.yml`, MkDocsConfigYaml)
	utils.WriteFile(`mkdocs/docs/index.md`, MkDocsPageExample)
	utils.AppendToNginx(fmt.Sprintf(MkDocsNginx, domain, domain, domain))
	utils.AppendToCerts(mail, "docs."+domain)
	if logo != `` {
		GenerateMkdocsLogo(logo)
	}
}

func GenerateMkdocsLogo(logo string) {
	utils.PrepareDir(MkDocsAssetsDir + `.gitkeep`)
	img.SvgToPng(logo, MkDocsAssetsDir+`logo.png`, 64)
	img.SvgToPng(logo, MkDocsAssetsDir+`favicon.png`, 16)
}

const MkDocsAssetsDir = `mkdocs/assets/`

const MkDocsPageExample = `# Page name example

Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem 
Ipsum has been the industry's standard dummy text ever since the 1500s, when an 
unknown printer took a galley of type and scrambled it to make a type specimen 
book. It has survived not only five centuries, but also the leap into electronic 
typesetting, remaining essentially unchanged. It was popularised in the 1960s 
with the release of Letraset sheets containing Lorem Ipsum passages, and more 
recently with desktop publishing software like Aldus PageMaker including 
versions of Lorem Ipsum.

`

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

repo_url: %s
repo_name: %s

theme:
  name: material
  logo: assets/logo.png
  favicon: assets/favicon.png

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

`
