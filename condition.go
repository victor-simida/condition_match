package condition

// Condition single condition struct
type Condition struct {
	Key   string
	Op    byte
	Value interface{}
}

// Match return single condition match the input map or not
func (con Condition) Match(input map[string]interface{}) bool {
	if v, ok := input[con.Key]; !ok {
		return false
	} else {
		return v == con.Value
	}
}

// NewCondition ...
func NewCondition(key string, op byte, value interface{}) Condition {
	return Condition{
		Key:   key,
		Op:    op,
		Value: value,
	}
}
