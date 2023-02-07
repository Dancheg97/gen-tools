package templates

const Makefile = `check:
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest gofumpt -l -w .
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest golangci-lint run
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf lint

.PHONY: gen
gen:
	sudo chmod a+rwx -R .
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf format -w
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf generate
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest swag fmt
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest swag init \
	 --parseDependency --parseInternal --parseDepth 1 -o . -ot yaml

.PHONY: cert
cert:
	go install github.com/go-acme/lego/v4/cmd/lego@latest
	sudo lego --email="name@example.com" --domains="sub.example.com" --http run
	sudo chown a+rwx -R .lego
 
`
