// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: centerms.proto

package proto

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Event) Validate() error {
	if this.AggregateId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AggregateId", fmt.Errorf(`value '%v' must not be an empty string`, this.AggregateId))
	}
	if this.AggregateType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AggregateType", fmt.Errorf(`value '%v' must not be an empty string`, this.AggregateType))
	}
	if this.EventId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EventId", fmt.Errorf(`value '%v' must not be an empty string`, this.EventId))
	}
	if this.EventType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EventType", fmt.Errorf(`value '%v' must not be an empty string`, this.EventType))
	}
	if this.EventData == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EventData", fmt.Errorf(`value '%v' must not be an empty string`, this.EventData))
	}
	return nil
}
func (this *DevMeta) Validate() error {
	if this.Type == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must not be an empty string`, this.Type))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Mac == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Mac", fmt.Errorf(`value '%v' must not be an empty string`, this.Mac))
	}
	return nil
}
func (this *SetDevInitCfgRequest) Validate() error {
	if nil == this.Meta {
		return github_com_mwitkow_go_proto_validators.FieldError("Meta", fmt.Errorf("message must exist"))
	}
	if this.Meta != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Meta); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Meta", err)
		}
	}
	return nil
}
func (this *SetDevInitCfgResponse) Validate() error {
	return nil
}
func (this *SaveDevDataRequest) Validate() error {
	if nil == this.Meta {
		return github_com_mwitkow_go_proto_validators.FieldError("Meta", fmt.Errorf("message must exist"))
	}
	if this.Meta != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Meta); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Meta", err)
		}
	}
	return nil
}
func (this *SaveDevDataResponse) Validate() error {
	if this.Status == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Status", fmt.Errorf(`value '%v' must not be an empty string`, this.Status))
	}
	return nil
}
