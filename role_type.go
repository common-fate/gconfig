package gconfig

import (
	"fmt"
	"strings"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
)

type RoleType int

//go:generate go run github.com/alvaroloes/enumer -type=RoleType -linecomment
const (
	RoleTypeAWS  RoleType = iota + 1 // aws
	RoleTypeOkta                     // okta
)

func (rt RoleType) ToProto() gconfigv1alpha1.RoleType {
	switch rt {
	case RoleTypeAWS:
		return gconfigv1alpha1.RoleType_ROLE_TYPE_AWS
	case RoleTypeOkta:
		return gconfigv1alpha1.RoleType_ROLE_TYPE_OKTA
	}

	return gconfigv1alpha1.RoleType_ROLE_TYPE_UNSPECIFIED
}

func RoleTypeFromProto(in gconfigv1alpha1.RoleType) RoleType {
	switch in {
	case gconfigv1alpha1.RoleType_ROLE_TYPE_AWS:
		return RoleTypeAWS
	case gconfigv1alpha1.RoleType_ROLE_TYPE_OKTA:
		return RoleTypeOkta
	}
	return 0
}

// ParseRoleType displays a more helpful error message than the default enum parser method.
func ParseRoleType(s string) (RoleType, error) {
	rt, err := RoleTypeString(s)
	if err != nil {
		vals := RoleTypeValues()
		var roletypes []string
		for _, v := range vals {
			roletypes = append(roletypes, v.String())
		}
		return 0, fmt.Errorf("%s is not a valid role type, valid types are: %s", s, strings.Join(roletypes, ", "))
	}

	return rt, nil
}
