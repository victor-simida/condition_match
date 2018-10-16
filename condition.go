package condition

import (
	"reflect"
	"strings"
)

// condition op define
const (
	OpEqual     = iota //==
	OpNotEqual         //!=
	OpHasPrefix        // xxx*
	OpIn               // is a member of the target slice
	OpNotIn            // not the member of the target slice
)

// Condition single condition struct
type Condition struct {
	Key   string
	Op    byte
	Value interface{}
}

// HasElem ...
func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

// Match return single condition match the input map or not
func (con Condition) Match(input map[string]interface{}) bool {
	v, ok := input[con.Key]
	if !ok {
		return false
	}

	switch con.Op {
	case OpEqual:
		return v == con.Value
	case OpNotEqual:
		return v != con.Value
	case OpHasPrefix:
		target, ok := con.Value.(string)
		if !ok {
			return false
		}
		value, ok := v.(string)
		if !ok {
			return false
		}
		return strings.HasPrefix(value, target)
	case OpIn:
		return HasElem(con.Value, v)
	case OpNotIn:
		return !HasElem(con.Value, v)
	}
	return v == con.Value
}

// NewCondition ...
func NewCondition(key string, op byte, value interface{}) Condition {
	return Condition{
		Key:   key,
		Op:    op,
		Value: value,
	}
}
