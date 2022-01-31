// Code generated by "enumer -type=Policy -linecomment"; DO NOT EDIT.

//
package gconfig

import (
	"fmt"
)

const _PolicyName = "allowrequireApprovalrequireReason"

var _PolicyIndex = [...]uint8{0, 5, 20, 33}

func (i Policy) String() string {
	i -= 1
	if i < 0 || i >= Policy(len(_PolicyIndex)-1) {
		return fmt.Sprintf("Policy(%d)", i+1)
	}
	return _PolicyName[_PolicyIndex[i]:_PolicyIndex[i+1]]
}

var _PolicyValues = []Policy{1, 2, 3}

var _PolicyNameToValueMap = map[string]Policy{
	_PolicyName[0:5]:   1,
	_PolicyName[5:20]:  2,
	_PolicyName[20:33]: 3,
}

// PolicyString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PolicyString(s string) (Policy, error) {
	if val, ok := _PolicyNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Policy values", s)
}

// PolicyValues returns all values of the enum
func PolicyValues() []Policy {
	return _PolicyValues
}

// IsAPolicy returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Policy) IsAPolicy() bool {
	for _, v := range _PolicyValues {
		if i == v {
			return true
		}
	}
	return false
}
