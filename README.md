# Golang ProtoBuf Validator Compiler

[![Travis Build](https://travis-ci.org/lucianoapolo/go-proto-validators.svg)](https://travis-ci.org/lucianoapolo/go-proto-validators)
[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

A `protoc` plugin that generates `Validate() error` functions on Go proto `struct`s based on field options inside `.proto`
files. The validation functions are code-generated and thus don't suffer on performance from tag-based reflection on
deeply-nested messages.

## Requirements

Using Protobuf validators is currently verified to work with:

- Go 1.11, 1.12, 1.13
- [Protobuf](https://github.com/protocolbuffers/protobuf) @ `v3.8.0`
- [Go Protobuf](https://github.com/golang/protobuf) @ `v1.5.2`
- [Gogo Protobuf](https://github.com/gogo/protobuf) @ `v1.3.2`

It _should_ still be possible to use it in project using earlier Go versions. However if you want to contribute to this
repository you'll need at least 1.11 for Go module support.

## Paint me a code picture

Let's take the following `proto3` snippet:

```proto
syntax = "proto3";
package validator.examples;
import "github.com/lucianoapolo/go-proto-validators/validator.proto";

message InnerMessage {
  // some_integer can only be in range (0, 100).
  int32 some_integer = 1 [(validator.field) = {int_gt: 0, int_lt: 100}];
  // some_float can only be in range (0;1).
  double some_float = 2 [(validator.field) = {float_gte: 0, float_lte: 1}];
}

message OuterMessage {
  // important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
  string important_string = 1 [(validator.field) = {regex: "^[a-z0-9]{5,30}$"}];
  // proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
  InnerMessage inner = 2 [(validator.field) = {msg_exists : true}];
}
```

First, the **`required` keyword is back** for `proto3`, under the guise of `msg_exists`. The painful `if-nil` checks are taken care of!

Second, the expected values in fields are now part of the contract `.proto` file. No more hunting down conditions in code!

Third, the generated code is understandable and has clear understandable error messages. Take a look:

```go
func (this *InnerMessage) Validate() []*google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation {
	fieldsViolations := []*google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{}
	if !(this.SomeInteger > 0) {
		fieldViolation := &google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{Field: "InnerMessage.SomeInteger", Description: "InnerMessage.SomeInteger must be greater than '0'"}
		fieldsViolations = append(fieldsViolations, fieldViolation)
	}
	if !(this.SomeInteger < 100) {
		fieldViolation := &google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{Field: "InnerMessage.SomeInteger", Description: "InnerMessage.SomeInteger must be less than '100'"}
		fieldsViolations = append(fieldsViolations, fieldViolation)
	}
	if !(this.SomeFloat >= 0) {
		fieldViolation := &google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{Field: "InnerMessage.SomeFloat", Description: "InnerMessage.SomeFloat must be greater than or equal to '0'"}
		fieldsViolations = append(fieldsViolations, fieldViolation)
	}
	if !(this.SomeFloat <= 1) {
		fieldViolation := &google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{Field: "InnerMessage.SomeFloat", Description: "InnerMessage.SomeFloat must be less than or equal to '1'"}
		fieldsViolations = append(fieldsViolations, fieldViolation)
	}
	if len(fieldsViolations) > 0 {
		return fieldsViolations
	} else {
		return nil
	}
}

var _regex_OuterMessage_ImportantString = regexp.MustCompile("^[a-z0-9]{5,30}$")

func (this *OuterMessage) Validate() []*google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation {
	if !_regex_OuterMessage_ImportantString.MatchString(this.ImportantString) {
		fieldViolation := &google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{Field: "OuterMessage.ImportantString", Description: "OuterMessage.ImportantString must conform to regex '^[a-z0-9]{5,30}$'"}
		fieldsViolations = append(fieldsViolations, fieldViolation)
	}
	if nil == this.Inner {
		fieldViolation := &google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{Field: "OuterMessage.Inner", Description: "OuterMessage.Inner message must exist"}
		fieldsViolations = append(fieldsViolations, fieldViolation)
	}
	if this.Inner != nil {
		if fieldsViolationsChild := github_com_lucianoapolo_go_proto_validators.CallValidatorIfExists(this.Inner); fieldsViolationsChild != nil {
			fieldsViolations = append(fieldsViolations, fieldsViolationsChild...)
		}
	}
	if len(fieldsViolations) > 0 {
		return fieldsViolations
	} else {
		return nil
	}
}
```

## Installing and using

The `protoc` compiler expects to find plugins named `proto-gen-XYZ` on the execution `$PATH`. So first:

```sh
export PATH=${PATH}:${GOPATH}/bin
```

Then, do the usual

```sh
go get github.com/lucianoapolo/go-proto-validators/protoc-gen-govalidators
```

Your `protoc` builds probably look very simple like:

```sh
protoc  \
  --proto_path=. \
  --go_out=. \
  *.proto
```

That's fine, until you encounter `.proto` includes. Because `go-proto-validators` uses field options inside the `.proto` 
files themselves, it's `.proto` definition (and the Google `descriptor.proto` itself) need to on the `protoc` include
path. Hence the above becomes:

```sh
protoc  \
  --proto_path=${GOPATH}/src \
  --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
  --proto_path=. \
  --go_out=. \
  --govalidators_out=. \
  *.proto
```

Or with gogo protobufs:

```sh
protoc  \
  --proto_path=${GOPATH}/src \
  --proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
  --proto_path=. \
  --gogo_out=. \
  --govalidators_out=gogoimport=true:. \
  *.proto
```

Basically the magical incantation (apart from includes) is the `--govalidators_out`. That triggers the 
`protoc-gen-govalidators` plugin to generate `mymessage.validator.pb.go`. That's it :)

## License

`go-proto-validators` is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
