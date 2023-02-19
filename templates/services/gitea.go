package services

import (
	"fmt"

	"dancheg97.ru/dancheg97/gen-tools/utils"
	"github.com/sirupsen/logrus"
)

func GenerateGitea(mail string, domain string, logo string) {
	utils.WriteFile(`gitea/gitea/templates/home.tmpl`, GiteaHomeTmpl)
	utils.WriteFile(`gitea/gitea/templates/custom/body_outer_pre.tmpl`, GiteaThemeParkTmpl)
	utils.WriteFile(`gitea/gitea/public/css/theme-earl-grey.css`, GiteaEarlGrayCss)
	utils.AppendToCompose(fmt.Sprintf(GiteaYaml, domain, domain, domain))
	utils.AppendToNginx(fmt.Sprintf(GiteaNginx, domain, domain, domain))
	utils.AppendToCerts(mail, "gitea."+domain)
	GenerateGiteaLogo(logo)
}

func GenerateGiteaLogo(logo string) {
	if logo == `` {
		logrus.Info(`no gitea logo provided, using default logo`)
		return
	}
	err := utils.CopyFile(logo, `logo.svg`)
	utils.CheckErr(err)

	utils.WriteFile(`logo-generator-gitea.js`, GiteaLogoGenerator)

	err = utils.SystemCall(`node logo-generator-gitea.js`)
	utils.CheckErr(err)
}

const GiteaNginx = `
server {
    listen 80;
    listen 443 ssl;
    server_name gitea.%s;
    ssl_certificate /certs/gitea.%s.crt;
    ssl_certificate_key /certs/gitea.%s.key;
    client_max_body_size 500M;
    location / {
        proxy_pass http://gitea/;
    }
}
`

const GiteaYaml = `
  gitea:
    image: gitea/gitea:1.17.3
    container_name: gitea
    restart: unless-stopped
    environment:
      GITEA__server__DOMAIN: gitea.%s
      GITEA__server__SSH_DOMAIN: gitea.%s
      GITEA__server__HTTP_PORT: 80
      GITEA__server__ROOT_URL: https://gitea.%s/
      GITEA__ui__THEMES: gitea,arc-green,plex,aquamarine,dark,dracula,hotline,organizr,space-gray,hotpink,onedark,overseerr,nord,earl-grey
      GITEA__ui__DEFAULT_THEME: earl-grey
      GITEA__service_DEFAULT_USER_IS_RESTRICTED: true
    volumes:
      - ./gitea:/data
      - /etc/timezone:/etc/timezone:ro
      - /etc/localtime:/etc/localtime:ro

`

const GiteaThemeParkTmpl = `{{ if .IsSigned }}
{{ if and (ne .SignedUser.Theme "gitea") (ne .SignedUser.Theme "arc-green") }}
  <link rel="stylesheet" href="https://theme-park.dev/css/base/gitea/{{.SignedUser.Theme}}.css">
{{end}}
{{ else if and (ne DefaultTheme "gitea") (ne DefaultTheme "arc-green") }}
<link rel="stylesheet" href="https://theme-park.dev/css/base/gitea/{{DefaultTheme}}.css">
{{end}}
`

const GiteaEarlGrayCss = `.chroma .hl {
  background-color: #3f424d;
}
.chroma .ln,
.chroma .lnt {
  color: #7f7f7f;
}
.chroma .k {
  color: #f63;
}
.chroma .kc {
  color: #fa1;
}
.chroma .kd {
  color: #9daccc;
}
.chroma .kn {
  color: #fa1;
}
.chroma .kp {
  color: #5f8700;
}
.chroma .kr {
  color: #f63;
}
.chroma .kt {
  color: #9daccc;
}
.chroma .na {
  color: #8a8a8a;
}
.chroma .bp,
.chroma .nb {
  color: #9daccc;
}
.chroma .nc,
.chroma .no {
  color: #fa1;
}
.chroma .nd {
  color: #9daccc;
}
.chroma .ni {
  color: #fa1;
}
.chroma .ne {
  color: #af8700;
}
.chroma .nf {
  color: #9daccc;
}
.chroma .nl,
.chroma .nn {
  color: #fa1;
}
.chroma .nt,
.chroma .nv,
.chroma .nx {
  color: #9daccc;
}
.chroma .vc {
  color: #f81;
}
.chroma .vg,
.chroma .vi {
  color: #fa1;
}
.chroma .s,
.chroma .sa {
  color: #1af;
}
.chroma .sb {
  color: #40AAFF;
}
.chroma .dl,
.chroma .sc {
  color: #1af;
}
.chroma .sd {
  color: #6a737d;
}
.chroma .s2 {
  color: #40AAFF;
}
.chroma .se {
  color: #f63;
}
.chroma .sh {
  color: #1af;
}
.chroma .si,
.chroma .sx {
  color: #fa1;
}
.chroma .sr {
  color: #97c;
}
.chroma .s1 {
  color: #40AAFF;
}
.chroma .ss {
  color: #fa1;
}
.chroma .il,
.chroma .m,
.chroma .mb,
.chroma .mf,
.chroma .mh,
.chroma .mi,
.chroma .mo {
  color: #1af;
}
.chroma .o {
  color: #f63;
}
.chroma .ow {
  color: #5f8700;
}
.chroma .c,
.chroma .c1,
.chroma .ch,
.chroma .cm {
  color: #6a737d;
}
.chroma .cs {
  color: #637d;
}
.chroma .cp,
.chroma .cpf {
  color: #fc6;
}
.chroma .gd {
  color: #fff;
  background-color: #5f3737;
}
.chroma .ge {
  color: #ef5;
}
.chroma .gr {
  color: #f33;
}
.chroma .gh {
  color: #fa1;
}
.chroma .gi {
  color: #fff;
  background-color: #3a523a;
}
.chroma .go {
  color: #888888;
}
.chroma .gp {
  color: #555555;
}
.chroma .gu {
  color: #9daccc;
}
.chroma .gt {
  color: #f63;
}
.chroma .w {
  color: #bbbbbb;
}
:root {
  --color-primary: #3D84E7;
  --color-primary-dark-1: #739cb3;
  --color-primary-dark-2: #40AAFF;
  --color-primary-dark-3: #92b4c4;
  --color-primary-dark-4: #a1bbcd;
  --color-primary-dark-5: #cfddc1;
  --color-primary-dark-6: #e7eee0;
  --color-primary-dark-7: #f8faf6;
  --color-primary-light-1: #3D84E7;
  --color-primary-light-2: #437aad;
  --color-primary-light-3: #415b8b;
  --color-primary-light-4: #25425a;
  --color-primary-light-5: #223546;
  --color-primary-light-6: #131923;
  --color-primary-light-7: #06090b;
  --color-primary-alpha-10: #3683C019;
  --color-primary-alpha-20: #3683C033;
  --color-primary-alpha-30: #3683C04b;
  --color-primary-alpha-40: #3683C066;
  --color-primary-alpha-50: #3683C080;
  --color-primary-alpha-60: #3683C099;
  --color-primary-alpha-70: #3683C0b3;
  --color-primary-alpha-80: #3683C0cc;
  --color-primary-alpha-90: #3683C0e1;
  --color-secondary: #2C2F35;
  --color-secondary-dark-1: #505665;
  --color-secondary-dark-2: #5b6273;
  --color-secondary-dark-3: #71798e;
  --color-secondary-dark-4: #7f8699;
  --color-secondary-dark-5: #8c93a4;
  --color-secondary-dark-6: #9aa0af;
  --color-secondary-dark-7: #a8adba;
  --color-secondary-dark-7: #b6bac5;
  --color-secondary-dark-8: #c4c7d0;
  --color-secondary-dark-8: #d2d4db;
  --color-secondary-dark-9: #dfe1e6;
  --color-secondary-dark-10: #edeef1;
  --color-secondary-dark-11: #fbfbfc;
  --color-secondary-light-1: #373b46;
  --color-secondary-light-2: #292c34;
  --color-secondary-light-3: #1c1e23;
  --color-secondary-light-4: #0e0f11;
  --color-secondary-alpha-10: #2C2F35;
  --color-secondary-alpha-20: #2C2F3533;
  --color-secondary-alpha-30: #2C2F354b;
  --color-secondary-alpha-40: #2C2F3566;
  --color-secondary-alpha-50: #2C2F3580;
  --color-secondary-alpha-60: #2C2F3599;
  --color-secondary-alpha-70: #2C2F35b3;
  --color-secondary-alpha-80: #2C2F35cc;
  --color-secondary-alpha-90: #2C2F35e1;
  --color-red: #da3737;
  --color-orange: #f17a2b;
  --color-yellow: #f3c640;
  --color-olive: #c8df36;
  --color-green: #3bc75b;
  --color-teal: #69d4cf;
  --color-blue: #2d9ff7;
  --color-violet: #754ad3;
  --color-purple: #b65dd4;
  --color-pink: #e04b9f;
  --color-brown: #a86d45;
  --color-grey: #797c85;
  --color-black: #141516;
  --color-gold: #d4b74c;
  --color-white: #ffffff;
  --color-diff-removed-word-bg: #6f3333;
  --color-diff-added-word-bg: #3c653c;
  --color-diff-removed-row-bg: #3c2626;
  --color-diff-added-row-bg: #283e2d;
  --color-diff-removed-row-border: #634343;
  --color-diff-added-row-border: #314a37;
  --color-diff-inactive: #1D1F23;
  --color-body: #1D1F23;
  /* main body color */
  --color-box-header: #1D1F23;
  --color-box-body: #1D1F23;
  --color-text-dark: #dbe0ea;
  --color-text: #bbc0ca;
  --color-text-light: #a6aab5;
  --color-text-light-2: #8a8e99;
  --color-text-light-3: #707687;
  --color-footer: #1D1F23;
  --color-timeline: #4c525e;
  --color-input-text: #d5dbe6;
  --color-input-background: #2C2F35;
  --color-input-border: #2C2F35;
  --color-input-border-hover: #505667;
  --color-navbar: #24262B;
  --color-light: #00000028;
  --color-light-border: #ffffff28;
  --color-hover: #ffffff10;
  --color-active: #ffffff16;
  --color-menu: #1D1F23;
  --color-card: #1D1F23;
  --color-markdown-table-row: #ffffff06;
  --color-markdown-code-block: #2C2F35;
  --color-button: #1D1F23;
  --color-code-bg: #1D1F23;
  --color-shadow: #00000060;
  --color-secondary-bg: #2C2F35;
  --color-text-focus: #fff;
  --color-expand-button: #2C2F35;
  --color-placeholder-text: #6a737d;
  --color-editor-line-highlight: var(--color-primary-light-5);
  --color-project-board-bg: var(--color-secondary-light-2);
}
::-webkit-calendar-picker-indicator {
  filter: invert(0.8);
}
.ui.horizontal.segments > .segment {
  background-color: #2C2F35;
}
.repository .segment.reactions .ui.label.basic.blue {
  background: var(--color-primary-alpha-20) !important;
}
[data-tooltip]:after,
[data-tooltip]:before {
  background: #1b1c1d !important;
  border-color: #1b1c1d !important;
  color: #dbdbdb !important;
}
[data-tooltip]:before {
  box-shadow: 1px 1px 0 0 #1b1c1d !important;
}
.ui.green.progress .bar {
  background-color: #668844;
}
.ui.progress.success .bar {
  background-color: #7b9e57 !important;
}
.following.bar.light {
  background: #1D1F23;
  border-color: var(--color-secondary-alpha-40);
}
.following.bar .top.menu a.item:hover {
  color: #fff;
}
.feeds .list ul li.private {
  background: #1D1F23;
}
.ui.link.list .item,
.ui.link.list .item a:not(.ui),
.ui.link.list a.item {
  color: #dbdbdb;
}
.ui.red.label,
.ui.red.labels .label {
  background-color: #E14C4C !important;
  border-color: #8a2121 !important;
}
.ui.yellow.label,
.ui.yellow.labels .label {
  border-color: #664d02 !important;
  background-color: #936e00 !important;
}
.ui.accordion .title:not(.ui) {
  color: #dbdbdb;
}
.ui.basic.green.label,
.ui.green.label,
.ui.green.labels .label {
  background-color: #0060AC !important;
  border-color: #0060AC !important;
}
.ui.basic.green.labels a.label:hover,
.ui.green.labels a.label:hover,
a.ui.basic.green.label:hover,
a.ui.ui.ui.green.label:hover {
  background-color: #3d794b !important;
  border-color: #3d794b !important;
  color: #fff !important;
}
.ui.divider:not(.vertical):not(.horizontal) {
  border-bottom-color: var(--color-secondary);
  border-top-color: transparent;
}
.form .help {
  color: #7f8699;
}
.ui .text.light.grey {
  color: #7f8699 !important;
}
.ui.form .field.error input:not([type]),
.ui.form .field.error input[type=date],
.ui.form .field.error input[type=datetime-local],
.ui.form .field.error input[type=email],
.ui.form .field.error input[type=file],
.ui.form .field.error input[type=number],
.ui.form .field.error input[type=password],
.ui.form .field.error input[type=search],
.ui.form .field.error input[type=tel],
.ui.form .field.error input[type=text],
.ui.form .field.error input[type=time],
.ui.form .field.error input[type=url],
.ui.form .field.error select,
.ui.form .field.error textarea,
.ui.form .fields.error .field input:not([type]),
.ui.form .fields.error .field input[type=date],
.ui.form .fields.error .field input[type=datetime-local],
.ui.form .fields.error .field input[type=email],
.ui.form .fields.error .field input[type=file],
.ui.form .fields.error .field input[type=number],
.ui.form .fields.error .field input[type=password],
.ui.form .fields.error .field input[type=search],
.ui.form .fields.error .field input[type=tel],
.ui.form .fields.error .field input[type=text],
.ui.form .fields.error .field input[type=time],
.ui.form .fields.error .field input[type=url],
.ui.form .fields.error .field select,
.ui.form .fields.error .field textarea {
  background-color: #522;
  border: 1px solid #E14C4C;
  color: #f9cbcb;
}
.ui.form .field.error input:not([type]):focus,
.ui.form .field.error input[type=date]:focus,
.ui.form .field.error input[type=datetime-local]:focus,
.ui.form .field.error input[type=email]:focus,
.ui.form .field.error input[type=file]:focus,
.ui.form .field.error input[type=number]:focus,
.ui.form .field.error input[type=password]:focus,
.ui.form .field.error input[type=search]:focus,
.ui.form .field.error input[type=tel]:focus,
.ui.form .field.error input[type=text]:focus,
.ui.form .field.error input[type=time]:focus,
.ui.form .field.error input[type=url]:focus,
.ui.form .field.error select:focus {
  background-color: #522;
  border: 1px solid #a04141;
  color: #f9cbcb;
}
.ui.green.button,
.ui.green.buttons .button {
  background-color: #3683C0;
}
.ui.green.button:hover,
.ui.green.buttons .button:hover {
  background-color: #40AAFF;
}
.ui.search > .results {
  background: #1D1F23;
  border-color: var(--color-secondary);
}
.ui.category.search > .results .category .result:hover,
.ui.search > .results .result:hover {
  background: var(--color-secondary);
}
.ui.search > .results .result .title {
  color: #dbdbdb;
}
.ui.table > thead > tr > th {
  background: var(--color-secondary);
  color: #dbdbdb !important;
}
.repository.file.list #repo-files-table tr {
  background: #1D1F23;
}
.repository.file.list #repo-files-table tr:hover {
  background-color: #24262B !important;
}
.repository.file.editor.edit + .editor-preview-side,
.repository.file.editor.edit .editor-preview,
.repository.file.editor.edit .editor-preview-side,
.repository.wiki.new .CodeMirror + .editor-preview-side,
.repository.wiki.new .CodeMirror .editor-preview,
.repository.wiki.new .CodeMirror .editor-preview-side {
  background: #24262B;
}
.repository.file.editor.edit + .editor-preview-side .markdown:not(code).ui.segment,
.repository.file.editor.edit .editor-preview-side .markdown:not(code).ui.segment,
.repository.file.editor.edit .editor-preview .markdown:not(code).ui.segment,
.repository.wiki.new .CodeMirror + .editor-preview-side .markdown:not(code).ui.segment,
.repository.wiki.new .CodeMirror .editor-preview-side .markdown:not(code).ui.segment,
.repository.wiki.new .CodeMirror .editor-preview .markdown:not(code).ui.segment {
  border-width: 0;
}
.overflow.menu .items .item {
  color: #9d9d9d;
}
.overflow.menu .items .item:hover {
  color: #dbdbdb;
}
.ui.list > .item > .content {
  color: var(--color-secondary-dark-6) !important;
}
.ui.active.button,
.ui.active.button:active,
.ui.button:active,
.ui.button:focus {
  background-color: #2C2F35;
  color: #dbdbdb;
}
.ui.green.button:active, .ui.green.buttons .button:active {
    background-color: #40AAFF;
    color: #FFFFFF;
    text-shadow: none;
}
.ui.green.button:active, .ui.green.buttons .button:active {
    background-color: #40AAFF;
    color: #FFFFFF;
    text-shadow: none;
}
.ui.green.button:focus, .ui.green.buttons .button:focus {
    background-color: #40AAFF;
    color: #FFFFFF;
    text-shadow: none;
}
.ui.active.button:hover {
  background-color: #474B51;
  color: #dbdbdb;
}
.repository .navbar .active.item,
.repository .navbar .active.item:hover {
  border-color: transparent !important;
}
.ui .info.segment.top {
  background-color: var(--color-secondary) !important;
}
.repository .diff-stats li {
  border-color: var(--color-secondary);
}
.tag-code,
.tag-code td {
  background: #24262B !important;
}
.tag-code td.lines-num {
  background-color: #3a3e4c !important;
}
.tag-code td.lines-type-marker,
td.blob-hunk {
  color: #dbdbdb !important;
}
.ui.attached.info.message,
.ui.info.message {
  box-shadow: inset 0 0 0 1px #4b5e71, 0 0 0 0 transparent;
}
.ui.bottom.attached.message {
  background-color: #2c662d;
  color: #3683C0;
}
.ui.bottom.attached.message .pull-right {
  color: #3683C0;
}
.ui.info.message {
  background-color: #2c3b4a;
  color: #9ebcc5;
}
.ui .warning.header,
.ui.warning.message {
  background-color: #542 !important;
  border-color: #ec8;
}
.ui.warning.message {
  color: #ec8;
  box-shadow: 0 0 0 1px #ec8;
}
.ui.warning.segment {
  border-color: #ec8;
}
.ui.error.message,
.ui.red.message {
  background-color: #522;
  color: #f9cbcb;
  box-shadow: inset 0 0 0 1px #a04141;
}
.ui .error.header,
.ui.error.message {
  background-color: #522 !important;
  border-color: #a04141;
}
.ui.error.segment {
  border-color: #a04141;
}
.ui.red.button,
.ui.red.buttons .button {
  background-color: #E14C4C;
}
.ui.red.button:hover,
.ui.red.buttons .button:hover {
  background-color: #984646;
}
.ui.positive.message {
  background-color: #48915A;
  color: #FFFFFF;
  box-shadow: inset 0 0 0 1px #48915A, 0 0 0 0 transparent;
}
.ui.negative.message {
  background-color: #E14C4C;
  color: #FFFFFF;
  box-shadow: inset 0 0 0 1px #E14C4C, 0 0 0 0 transparent;
}
.ui.list .list > .item .header,
.ui.list > .item .header {
  color: #dedede;
}
.ui.list .list > .item .description,
.ui.list > .item .description {
  color: var(--color-secondary-dark-6);
}
.repository.file.list #repo-files-table tbody .svg.octicon-file-directory,
.repository.file.list #repo-files-table tbody .svg.octicon-file-submodule {
  color: #9AA0AF;
}
.repository.labels .ui.basic.black.label {
  background-color: #bbbbbb !important;
}
.blame .lines-num,
.lines-commit {
  background: #1D1F23 !important;
}
.lines-num {
  color: var(--color-secondary-dark-6) !important;
  border-color: var(--color-secondary) !important;
}
td.blob-excerpt {
  background-color: rgba(0, 0, 0, 0.15);
}
.lines-code.active,
.lines-code .active {
  background: #534d1b !important;
}
.ui.ui.table td.active,
.ui.ui.ui.ui.table tr.active {
  color: #dbdbdb;
}
.ui.active.label {
  background: #393d4a;
  border-color: #393d4a;
  color: #dbdbdb;
}
.repository .ui.attached.message.isSigned.isVerified {
  background-color: #394829;
  color: var(--color-secondary-dark-6);
}
.repository .ui.attached.message.isSigned.isVerified.message {
  color: #3683C0;
}
.repository .ui.attached.message.isSigned.isVerified.message .ui.text {
  color: var(--color-secondary-dark-6);
}
.repository .ui.attached.message.isSigned.isVerified.message .pull-right {
  color: #3683C0;
}
.repository .ui.attached.message.isSigned.isVerifiedUntrusted {
  background-color: #4a3903;
  color: var(--color-secondary-dark-6);
}
.repository .ui.attached.message.isSigned.isVerifiedUntrusted.message {
  color: #c2c193;
}
.repository .ui.attached.message.isSigned.isVerifiedUntrusted.message .ui.text {
  color: var(--color-secondary-dark-6);
}
.repository .ui.attached.message.isSigned.isVerifiedUntrusted.message a {
  color: #c2c193;
}
.repository .ui.attached.message.isSigned.isVerifiedUnmatched {
  background-color: #4e3321;
  color: var(--color-secondary-dark-6);
}
.repository .ui.attached.message.isSigned.isVerifiedUnmatched.message {
  color: #c2a893;
}
.repository .ui.attached.message.isSigned.isVerifiedUnmatched.message .ui.text {
  color: var(--color-secondary-dark-6);
}
.repository .ui.attached.message.isSigned.isVerifiedUnmatched.message a {
  color: #c2a893;
}
.repository .ui.attached.message.isSigned.isWarning {
  background-color: rgba(80, 23, 17, 0.6);
}
.repository .ui.attached.message.isSigned.isWarning.message,
.repository .ui.attached.message.isSigned.isWarning.message .ui.text {
  color: #d07d7d;
}
.ui.header .sub.header {
  color: var(--color-secondary-dark-6);
}
.ui.dividing.header {
  border-bottom: 1px solid var(--color-secondary);
}
.ui.modal > .header {
  background: var(--color-secondary);
  color: #dbdbdb;
}
.ui.modal > .actions {
  background: var(--color-secondary);
  border-color: var(--color-secondary);
}
.ui.modal > .content {
  background: #1D1F23;
}
.minicolors-panel {
  background: var(--color-secondary) !important;
  border-color: #6a737d !important;
}
.emoji[aria-label="check mark"],
.emoji[aria-label="curly loop"],
.emoji[aria-label="currency exchange"],
.emoji[aria-label="double curly loop"],
.emoji[aria-label="END arrow"],
.emoji[aria-label="heavy dollar sign"],
.emoji[aria-label="musical note"],
.emoji[aria-label="musical notes"],
.emoji[aria-label="ON! arrow"],
.emoji[aria-label="paw prints"],
.emoji[aria-label="SOON arrow"],
.emoji[aria-label="TOP arrow"],
.emoji[aria-label="trade mark"],
.emoji[aria-label="wavy dash"],
.emoji[aria-label=copyright],
.emoji[aria-label=divide],
.emoji[aria-label=minus],
.emoji[aria-label=multiply],
.emoji[aria-label=plus],
.emoji[aria-label=registered] {
  filter: invert(100%);
}
.edit-diff > div > .ui.table {
  border-left-color: var(--color-secondary) !important;
  border-right-color: var(--color-secondary) !important;
}
.CodeMirror.cm-s-default .cm-property,
.CodeMirror.cm-s-paper .cm-property {
  color: #40AAFF;
}
.CodeMirror.cm-s-default .cm-header,
.CodeMirror.cm-s-paper .cm-header {
  color: #9daccc;
}
.CodeMirror.cm-s-default .cm-quote,
.CodeMirror.cm-s-paper .cm-quote {
  color: #009900;
}
.CodeMirror.cm-s-default .cm-keyword,
.CodeMirror.cm-s-paper .cm-keyword {
  color: #cc8a61;
}
.CodeMirror.cm-s-default .cm-atom,
.CodeMirror.cm-s-paper .cm-atom {
  color: #ef5e77;
}
.CodeMirror.cm-s-default .cm-number,
.CodeMirror.cm-s-paper .cm-number {
  color: #ff5656;
}
.CodeMirror.cm-s-default .cm-def,
.CodeMirror.cm-s-paper .cm-def {
  color: #e4e4e4;
}
.CodeMirror.cm-s-default .cm-variable-2,
.CodeMirror.cm-s-paper .cm-variable-2 {
  color: #00bdbf;
}
.CodeMirror.cm-s-default .cm-variable-3,
.CodeMirror.cm-s-paper .cm-variable-3 {
  color: #008855;
}
.CodeMirror.cm-s-default .cm-comment,
.CodeMirror.cm-s-paper .cm-comment {
  color: #8e9ab3;
}
.CodeMirror.cm-s-default .cm-string,
.CodeMirror.cm-s-paper .cm-string {
  color: #a77272;
}
.CodeMirror.cm-s-default .cm-string-2,
.CodeMirror.cm-s-paper .cm-string-2 {
  color: #ff5500;
}
.CodeMirror.cm-s-default .cm-meta,
.CodeMirror.cm-s-default .cm-qualifier,
.CodeMirror.cm-s-paper .cm-meta,
.CodeMirror.cm-s-paper .cm-qualifier {
  color: #ffb176;
}
.CodeMirror.cm-s-default .cm-builtin,
.CodeMirror.cm-s-paper .cm-builtin {
  color: #b7c951;
}
.CodeMirror.cm-s-default .cm-bracket,
.CodeMirror.cm-s-paper .cm-bracket {
  color: #999977;
}
.CodeMirror.cm-s-default .cm-tag,
.CodeMirror.cm-s-paper .cm-tag {
  color: #f1d273;
}
.CodeMirror.cm-s-default .cm-attribute,
.CodeMirror.cm-s-paper .cm-attribute {
  color: #bfcc70;
}
.CodeMirror.cm-s-default .cm-hr,
.CodeMirror.cm-s-paper .cm-hr {
  color: #999999;
}
.CodeMirror.cm-s-default .cm-url,
.CodeMirror.cm-s-paper .cm-url {
  color: #c5cfd0;
}
.CodeMirror.cm-s-default .cm-link,
.CodeMirror.cm-s-paper .cm-link {
  color: #d8c792;
}
.CodeMirror.cm-s-default .cm-error,
.CodeMirror.cm-s-paper .cm-error {
  color: #dbdbeb;
}
footer .container .links > * {
  border-left-color: #888;
}
.repository.file.list #repo-files-table tbody .svg {
  color: var(--color-secondary-dark-6);
}
.repository.release #release-list > li .detail .dot {
  background-color: #505667;
  border-color: #1D1F23;
}
.tribute-container {
  box-shadow: 0 0.25rem 0.5rem rgba(0, 0, 0, 0.6);
}
.repository .repo-header .ui.huge.breadcrumb.repo-title .repo-header-icon .avatar {
  color: #1D1F23;
}
img[src$="/img/matrix.svg"] {
  filter: invert(80%);
}
#git-graph-container li .time {
  color: #6a737d;
}
#git-graph-container.monochrome #rel-container .flow-group {
  stroke: dimgrey;
  fill: dimgrey;
}
#git-graph-container.monochrome #rel-container .flow-group.highlight {
  stroke: darkgrey;
  fill: darkgrey;
}
#git-graph-container:not(.monochrome) #rel-container .flow-group.flow-color-16-5 {
  stroke: #5543b1;
  fill: #5543b1;
}
#git-graph-container:not(.monochrome) #rel-container .flow-group.highlight.flow-color-16-5 {
  stroke: #7058e6;
  fill: #7058e6;
}
#git-graph-container #rev-list li.highlight.hover {
  background-color: rgba(255, 255, 255, 0.1);
}
#git-graph-container .ui.buttons button#flow-color-monochrome.ui.button {
  border-left: 1px solid #4c505c;
}
.mermaid-chart {
  filter: invert(84%) hue-rotate(180deg);
}
.is-loading:after {
  border-color: #4a4c58 #4a4c58 #d7d7da #d7d7da;
}
.markdown-block-error {
  border: 1px solid rgba(121, 71, 66, 0.5) !important;
  border-bottom: none !important;
}

.monaco-editor .view-lines {
	filter: invert(1) !important;
}
`

const GiteaHomeTmpl = `{{template "base/head" .}}

<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Sofia+Sans:wght@300&display=swap" rel="stylesheet">
<p class="features">Random message</p>
<div id="light">
    <div id="lineh1"></div>
</div>
<img class="logo_main" src="https://backstage.io/animations/backstage-techdocs-icon-1.gif" alt="logo of website">

<p class="features_two">Welcome to gitea!
    <br> Example message
</p>

<div class="arrow-8"></div>

<div class="line"></div>

<div class="choices">

    <h1 class="dbHeadName">Some custom badges</h1>
    <div class="choicesTable">
        <div class="img-badge">
            <a href='https://www.postgresql.org/'>
                <img
                    src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/29/Postgresql_elephant.svg/1200px-Postgresql_elephant.svg.png" />
            </a>
            <h2>PostgreSQL</h2>
            <p>Advanced, enterprise class open source relational database that supports both SQL (relational) and JSON
                (non-relational) querying.</p>
        </div>
        <div class="img-badge">
            <a href='https://redis.io/'>
                <img
                    src="https://camo.githubusercontent.com/4050472d0036e02ed3805e8329474f062eac6ae847ca0ac107d4889fa778711a/68747470733a2f2f6973332d73736c2e6d7a7374617469632e636f6d2f696d6167652f7468756d622f507572706c653132342f76342f31372f63642f61322f31376364613261302d623634312d633364302d336432322d3134313730346134306565662f49636f6e2e706e672f313230307836333062622e706e67" />
            </a>
            <h2>Redis</h2>
            <p>Open source (BSD licensed), in-memory data structure store, used as a database, cache, and
                message broke</p>
        </div>
        <div class="img-badge">
            <a href='https://nats.io/'>
                <img
                    src="https://raw.githubusercontent.com/docker-library/docs/ad703934a62fabf54452755c8486698ff6fc5cc2/nats/logo.png" />
            </a>
            <h2>Nats</h2>
            <p>Neural Autonomic Transport System. Derek Collison conceived NATS as a messaging platform that
                functions like a central nervous system.</p>
        </div>
        <div class="img-badge">
            <a href='https://github.com/google/leveldb'>
                <img src="https://cdn.freebiesupply.com/logos/large/2x/leveldb-logo-png-transparent.png" />
            </a>
            <h2>LevelDB</h2>
            <p>Fast key-value storage library written at Google that provides an ordered mapping from string
                keys to string values</p>
        </div>
    </div>

    <div class="line2"></div>
</div>


<div class="area">
    <ul class="circles">
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
    </ul>
</div>

<style>
    @keyframes lineH {
        0% {
            width: 0%;
        }

        100% {
            width: 100%;
            opacity: 0;
        }
    }

    .name_main {
        text-align: center;
        font-size: 32px;
        margin-bottom: 0px;
        margin-top: 10px;
    }

    .logo_main {
        display: block;
        margin-left: auto;
        margin-right: auto;
        margin-bottom: 40px;
        margin-top: 40px;
    }

    .features {
        text-align: center;
        font-size: 80px;
        margin-bottom: 0px;
        font-family: "Sofia Sans", sans-serif;
    }

    .features_two {
        color: rgb(167, 167, 167);
        text-align: center;
        font-size: 28px;
        margin-bottom: 60px;
        font-family: "Sofia Sans", sans-serif;
    }

    .arrow-8 {
        position: relative;
        width: 100px;
        height: 100px;
        margin: 30px auto 80px auto;
    }

    .arrow-8:before,
    .arrow-8:after {
        content: '';
        position: absolute;
        box-sizing: border-box;
        width: 100%;
        height: 100%;
        border-left: 26px solid #7cf1df;
        border-bottom: 26px solid #7cf1df;
        animation: arrow-8 3s linear infinite;
    }

    .arrow-8:after {
        animation: arrow-8 3s linear infinite -1.5s;
    }

    @keyframes arrow-8 {
        0% {
            opacity: 0;
            transform: translate(0, -53px) rotate(-45deg);
        }

        10%,
        90% {
            opacity: 0;
        }

        50% {
            opacity: 1;
            transform: translate(0, 0) rotate(-45deg);
        }

        100% {
            opacity: 0;
            transform: translate(0, 53px) rotate(-45deg);
        }
    }

    .line {
        border: 0;
        height: 1px;
        background-image: -webkit-linear-gradient(left, #7cf1df, #1d1f23, #7cf1df);
        position: relative;
        top: 35px;
        animation: lineH 5s 0s infinite linear;
    }

    .line2 {
        border: 0;
        height: 1px;
        background-image: -webkit-linear-gradient(left, #7cf1df, #1d1f23, #7cf1df);
        position: relative;
        top: 10px;
        animation: lineH 5s 0s infinite linear;
    }


    .choices .dbHeadName {
        margin-top: 100px;
    }

    .choices {
        text-align: center;
        font-family: "Sofia Sans", sans-serif;
    }

    .choicesTable {
        display: flex;
        justify-content: space-around;
        flex-wrap: wrap;
    }

    .choices p {
        text-align: center;
        font-size: 40px;
        margin-bottom: 0px;
        color: rgb(167, 167, 167);
        font-size: 25px;
    }

    .choices h2 {
        font-size: 40px;
        font-weight: normal;
        color: white;
        font-family: "Sofia Sans", sans-serif;
    }

    .choices img {
        width: 68px;
        height: 68px;
    }

    .choices img:hover {
        transition: all 0.1s ease-in-out;
        transform: scale(1.2);
    }

    #light {
        position: relative;
        top: 150px;
    }

    h1 {
        color: white;
        font-family: "Sofia Sans", sans-serif;
        font-size: 60px;
        font-weight: normal;
    }

    .name_main {
        text-align: center;
        font-size: 100px;
        margin-bottom: 0px;
        margin-top: 20px;
    }

    @keyframes show {
        from {
            opacity: 0;
        }

        to {
            opacity: 1;
        }
    }

    .img-badge {
        margin-bottom: 20px;
        max-width: 300px;
    }


    #lineh1 {
        position: absolute;
        left: 0;
        bottom: 160px;
        height: 1px;
        background-image: -webkit-linear-gradient(left, #7cf1df, #1d1f23, #7cf1df);
        animation: lineH 5s 0s linear;
    }

    /* Анимация квадратов в начале экрана */
    .circles {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow: hidden;
        z-index: -1
    }

    .circles li {
        position: absolute;
        display: block;
        list-style: none;
        width: 20px;
        height: 20px;
        background: rgba(255, 255, 255, 0.048);
        animation: animate 25s linear infinite;
        bottom: -150px;

    }

    .circles li:nth-child(1) {
        left: 25%;
        width: 80px;
        height: 80px;
        animation-delay: 0s;
    }


    .circles li:nth-child(2) {
        left: 10%;
        width: 20px;
        height: 20px;
        animation-delay: 2s;
        animation-duration: 12s;
    }

    .circles li:nth-child(3) {
        left: 70%;
        width: 20px;
        height: 20px;
        animation-delay: 4s;
    }

    .circles li:nth-child(4) {
        left: 40%;
        width: 60px;
        height: 60px;
        animation-delay: 0s;
        animation-duration: 18s;
    }

    .circles li:nth-child(5) {
        left: 65%;
        width: 20px;
        height: 20px;
        animation-delay: 0s;
    }

    .circles li:nth-child(6) {
        left: 75%;
        width: 110px;
        height: 110px;
        animation-delay: 3s;
    }

    .circles li:nth-child(7) {
        left: 35%;
        width: 150px;
        height: 150px;
        animation-delay: 7s;
    }

    .circles li:nth-child(8) {
        left: 50%;
        width: 25px;
        height: 25px;
        animation-delay: 15s;
        animation-duration: 45s;
    }

    .circles li:nth-child(9) {
        left: 20%;
        width: 15px;
        height: 15px;
        animation-delay: 2s;
        animation-duration: 35s;
    }

    .circles li:nth-child(10) {
        left: 85%;
        width: 150px;
        height: 150px;
        animation-delay: 0s;
        animation-duration: 11s;
    }



    @keyframes animate {

        0% {
            transform: translateY(0) rotate(0deg);
            opacity: 1;
            border-radius: 0;
        }

        100% {
            transform: translateY(-1000px) rotate(720deg);
            opacity: 0;
            border-radius: 50%;
        }
    }
</style>


{{template "base/footer" .}}`

const GiteaLogoGenerator = `
#!/usr/bin/env node
import imageminZopfli from 'imagemin-zopfli';
import {optimize} from 'svgo';
import {fabric} from 'fabric';
import {readFile, writeFile} from 'node:fs/promises';

function exit(err) {
  if (err) console.error(err);
  process.exit(err ? 1 : 0);
}

function loadSvg(svg) {
  return new Promise((resolve) => {
    fabric.loadSVGFromString(svg, (objects, options) => {
      resolve({objects, options});
    });
  });
}

async function generate(svg, path, {size, bg}) {
  const outputFile = new URL(path, import.meta.url);

  if (String(outputFile).endsWith('.svg')) {
    const {data} = optimize(svg, {
      plugins: [
        'preset-default',
        'removeDimensions',
        {
          name: 'addAttributesToSVGElement',
          params: {attributes: [{width: size}, {height: size}]}
        },
      ],
    });
    await writeFile(outputFile, data);
    return;
  }

  const {objects, options} = await loadSvg(svg);
  const canvas = new fabric.Canvas();
  canvas.setDimensions({width: size, height: size});
  const ctx = canvas.getContext('2d');
  ctx.scale(options.width ? (size / options.width) : 1, options.height ? (size / options.height) : 1);

  if (bg) {
    canvas.add(new fabric.Rect({
      left: 0,
      top: 0,
      height: size * (1 / (size / options.height)),
      width: size * (1 / (size / options.width)),
      fill: 'white',
    }));
  }

  canvas.add(fabric.util.groupSVGElements(objects, options));
  canvas.renderAll();

  let png = Buffer.from([]);
  for await (const chunk of canvas.createPNGStream()) {
    png = Buffer.concat([png, chunk]);
  }

  png = await imageminZopfli({more: true})(png);
  await writeFile(outputFile, png);
}

async function main() {
  const gitea = process.argv.slice(2).includes('gitea');
  const logoSvg = await readFile(new URL('logo.svg', import.meta.url), 'utf8');

  await Promise.all([
    generate(logoSvg, 'gitea/gitea/public/img/logo.svg', {size: 32}),
    generate(logoSvg, 'gitea/gitea/public/img/logo.png', {size: 512}),
    generate(logoSvg, 'gitea/gitea/public/img/favicon.svg', {size: 32}),
    generate(logoSvg, 'gitea/gitea/public/img/favicon.png', {size: 180}),
    generate(logoSvg, 'gitea/gitea/public/img/avatar_default.png', {size: 200}),
    generate(logoSvg, 'gitea/gitea/public/img/apple-touch-icon.png', {size: 180, bg: true}),
    gitea && generate(logoSvg, 'gitea/gitea/public/img/gitea.svg', {size: 32}),
  ]);
}

main().then(exit).catch(exit);
`
