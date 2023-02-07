
check:
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest gofumpt -l -w .
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest golangci-lint run
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf lint

.PHONY: gen
gen:
	sudo chmod a+rwx -R .
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf format -w
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf generate
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest swag fmt
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest swag init --parseDependency --parseInternal --parseDepth 1 -o . -ot yaml
