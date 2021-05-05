# Copyright 2016 Michal Witkowski. All Rights Reserved.
# See LICENSE for licensing terms.

mkfile_dir = "$(dir $(abspath $(lastword $(MAKEFILE_LIST))))"

ifdef GOBIN
extra_path = "$(mkfile_dir)deps/bin:$(GOBIN)"
else
extra_path = "$(mkfile_dir)deps/bin:$(HOME)/go/bin"
endif

prepare_deps:
	@echo "--- Preparing dependencies."
	@bash scripts/prepare-deps.sh

gazelle:
	@bash bazel run --run_under="cd ${mkfile_dir} && " @bazel_gazelle//cmd/gazelle -- update-repos -from_file=go.mod -to_macro=go_deps.bzl%go_repositories
	@bash bazel run //:gazelle -- --mode=fix --exclude=deps --exclude=examples --exclude=test

install:
	@echo "--- Installing 'govalidators' binary to GOBIN."
	go install github.com/lucianoapolo/go-proto-validators/protoc-gen-govalidators

regenerate_test_gogo: prepare_deps install
	@echo "--- Regenerating test .proto files with gogo imports"
	export PATH=$(extra_path):$${PATH}; protoc  \
		--proto_path=deps \
		--proto_path=deps/include \
		--proto_path=test \
		--gogo_out=test/gogo \
		--govalidators_out=gogoimport=true:test/gogo test/*.proto

regenerate_test_golang: prepare_deps install
	@echo "--- Regenerating test .proto files with golang imports"
	export PATH=$(extra_path):$${PATH}; protoc  \
		--proto_path=deps \
		--proto_path=deps/include \
		--proto_path=test \
		--go_out=test/golang \
		--govalidators_out=test/golang test/*.proto

regenerate_example: prepare_deps install
	@echo "--- Regenerating example directory"
	export PATH=$(extra_path):$${PATH}; protoc  \
		--proto_path=deps \
		--proto_path=deps/include \
		--proto_path=. \
		--go_out=. --go_opt paths=source_relative \
		--govalidators_out=lang=pt_br:. examples/*.proto --govalidators_opt paths=source_relative

test: regenerate_test_gogo regenerate_test_golang
	@echo "Running tests"
	go test -v ./...

regenerate: prepare_deps
	@echo "--- Regenerating validator.proto"
	export PATH=$(extra_path):$${PATH}; protoc \
		--proto_path=deps \
		--proto_path=deps/include \
		--proto_path=deps/github.com/gogo/protobuf/protobuf \
		--proto_path=. \
		--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. --gogo_opt paths=source_relative \
		validator.proto

clean:
	rm -rf "deps"