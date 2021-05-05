// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/enum.proto

package validatorexamples

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lucianoapolo/go-proto-validators"
	google_golang_org_genproto_googleapis_rpc_errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SomeMsg) Validate() []*google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation {
	fieldsViolations := []*google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{}
	if _, ok := Action_name[int32(this.Do)]; !ok {
		fieldViolation := &google_golang_org_genproto_googleapis_rpc_errdetails.BadRequest_FieldViolation{Field: "Do", Description: fmt.Sprintf(`valor '%v' deve ser um válido Action enumerador`, this.Do)}
		fieldsViolations = append(fieldsViolations, fieldViolation)
	}
	if len(fieldsViolations) > 0 {
		return fieldsViolations
	} else {
		return nil
	}
}
