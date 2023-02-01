<p align="center">
<img style="align: center; padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="238px" height="238px" src="https://gitea.dancheg97.ru/repo-avatars/67-4297f15da3e76c29478ec89973007622" />
</p>

<h2 align="center">Go tools</h2>

[![Generic badge](https://img.shields.io/badge/LICENSE-GPLv3-red.svg)](https://gitea.dancheg97.ru/templates/go-tools/src/branch/main/LICENSE)
[![Generic badge](https://img.shields.io/badge/GITEA-REPO-orange.svg)](https://gitea.dancheg97.ru/templates/go-tools)
[![Generic badge](https://img.shields.io/badge/DOCKER-CONTAINER-blue.svg)](https://gitea.dancheg97.ru/templates/-/packages/container/go-tools/latest)
[![Build Status](https://drone.dancheg97.ru/api/badges/templates/go-tools/status.svg)](https://drone.dancheg97.ru/templates/go-tools)

Project goal is to simplify process of

---

You can pull docker container via command:

```sh
docker pull gitea.dancheg97.ru/templates/go-tools:latest
```

---

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
