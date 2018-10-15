package condition

import "fmt"

// Condition single condition struct
type Condition struct {
	Key   string
	Value interface{}
}

// Conditions conditions need to match
type Conditions struct {
	Op   byte //now we only allow eight conditions at most so use byte bit field to get the
	Spec []Condition
}

// Match return single condition match the input map or not
func (con Condition) Match(input map[string]interface{}) bool {
	if v, ok := input[con.Key]; !ok {
		return false
	} else {
		return v == con.Value
	}
}

// Judge judge the conditions with input key-value map
func (c Conditions) Judge(input map[string]interface{}) bool {
	l := len(c.Spec)
	if l == 0 {
		return true
	}
	index := make([]int, l-1)
	for i := 0; i < l-1; i++ {
		if c.Op&(0x1<<uint(i)) != 0 {
			index[i] = 1
		} else {
			index[i] = 0
		}
	}
	fmt.Println(index)

	result := c.Spec[0].Match(input)
	for i := 0; i < l-1; i++ {
		if index[i] == 0 {
			if result {
				return true
			}
			result = c.Spec[i+1].Match(input)
		} else {
			result = result && c.Spec[i+1].Match(input)
		}
	}

	return result
}
