package gogen

import "dancheg97.ru/templates/gen-tools/templates/utils"

func GenerateBuf(generate bool) {
	utils.WriteFile("buf.yaml", BufYaml)
	utils.WriteFile("buf.gen.yaml", BufGenYaml)
	utils.WriteFile("proto/v1/example.proto", GrpcProto)
	utils.AppendToMakefile(BufMake)
	if generate {
		utils.SystemCall("buf generate")
	}
}

const BufGenYaml = `version: v1
plugins:
  - plugin: go
    out: gen/pb
    opt: paths=source_relative
  - plugin: go-grpc
    out: gen/pb
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
`

const BufYaml = `version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT`

const GrpcProto = `syntax = "proto3";
package proto.v1;

option go_package = "gen/go/pb";

// Service description example.
service ExampleService {
  // Method description example.
  rpc Add(AddRequest) returns (AddResponse);
}

message AddRequest {
  string example = 1;
}

message AddResponse {}

`

const BufMake = `
buf:
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf lint
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf format -w
	docker run --rm -v ${pwd}:/src -w /src dancheg97.ru/templates/gen-tools:latest buf generate

`
