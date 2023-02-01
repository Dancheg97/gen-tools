<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="https://gitea.dancheg97.ru/repo-avatars/67-4297f15da3e76c29478ec89973007622" />
</p>

<h2 align="center">Go tools</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-GPLv3-red.svg)](https://gitea.dancheg97.ru/templates/go-tools/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-orange.svg)](https://gitea.dancheg97.ru/templates/go-tools)
[![Generic badge](https://img.shields.io/badge/DOCKER-CONTAINER-blue.svg)](https://gitea.dancheg97.ru/templates/-/packages/container/go-tools/latest)
[![Build Status](https://drone.dancheg97.ru/api/badges/templates/go-tools/status.svg)](https://drone.dancheg97.ru/templates/go-tools)

🧰 CLI Tool for generating templates of go code for interaction with different
infrastructural elements.

📃 Options you can specify under 'gen' command:

This tool allows to generate prepared go code for interaction with following
infrastructure components (go-tools gen [options]):

- **cli** - includes cobra and viper
- **drone** - includes drone template for CI-CD
- **lint** - includes golanglint-ci linter for go code
- **grpc** - includes proto and buf files for generation
- **docker** - includes 2 stage dockerfile and compose for ease of development
- **pg** - includes pgx module in porstgres, sqlc for generation and goose for migrations
- **redis** - includes redis template
- **nats** - includes consumer and producer nats template
- **license** - adds GPLv3 license to project

### Installation:

- docker

```
docker pull gitea.dancheg97.ru/templates/go-tools:latest
```

- go

```
go install gitea.dancheg97.ru/templates/go-tools@latest
```

### Examples:

- [go-tools](README.md) - tool for generating go project templates

```sh
go-tools gen cli lint docker makefile gpl
```

Collection includes following tooling:

- [gofumpt](https://github.com/mvdan/gofumpt) - tool for formatting go code

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/go-tools:latest gofumpt --help
```

- [golanglint-ci](https://golangci-lint.run/) - tool for linting go code, [config template](.golangci.yml)

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/go-tools:latest golanglint-ci --help
```

- [buf](https://docs.buf.build/introduction) - tool for helping with protocol buffers and gRPC, [buf example](buf.yaml), [buf gen example](buf.gen.yaml)

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/go-tools:latest buf --help
```

- [sqlc](https://docs.sqlc.dev/en/stable) - tool for generating type-safe go code from sql queries, [sqlc.sql example](sqlc.sql), config - [sqlc.yaml example](sqlc.yaml)

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/go-tools:latest sqlc --help
```

- [go-swag](https://github.com/swaggo/swag) - tool for generating `swagger.yaml` from code annotations.

```sh
docker run --rm -it -v ${pwd}:/wd -w /wd gitea.dancheg97.ru/templates/go-tools:latest swag --help
```
