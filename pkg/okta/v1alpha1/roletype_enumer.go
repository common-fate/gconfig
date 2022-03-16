// Code generated by "enumer -type=RoleType -linecomment"; DO NOT EDIT.

//
package gcoktav1alpha1

import (
	"fmt"
)

const _RoleTypeName = "OKTA_GROUP"

var _RoleTypeIndex = [...]uint8{0, 10}

func (i RoleType) String() string {
	i -= 1
	if i < 0 || i >= RoleType(len(_RoleTypeIndex)-1) {
		return fmt.Sprintf("RoleType(%d)", i+1)
	}
	return _RoleTypeName[_RoleTypeIndex[i]:_RoleTypeIndex[i+1]]
}

var _RoleTypeValues = []RoleType{1}

var _RoleTypeNameToValueMap = map[string]RoleType{
	_RoleTypeName[0:10]: 1,
}

// RoleTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RoleTypeString(s string) (RoleType, error) {
	if val, ok := _RoleTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to RoleType values", s)
}

// RoleTypeValues returns all values of the enum
func RoleTypeValues() []RoleType {
	return _RoleTypeValues
}

// IsARoleType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i RoleType) IsARoleType() bool {
	for _, v := range _RoleTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
