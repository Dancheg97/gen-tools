
check:
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest gofumpt -l -w .
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest golangci-lint run
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest buf lint

.PHONY: gen
gen:
	sudo chmod a+rwx -R .
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest buf format -w
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest buf generate
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest swag fmt
	docker run --rm -v ${pwd}:/src -w /src gitea.dancheg97.ru/templates/golden-go:latest swag init --parseDependency --parseInternal --parseDepth 1 -o . -ot yaml
