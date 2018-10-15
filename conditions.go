package condition

// Conditions conditions need to match
type Conditions struct {
	Op   byte //now we only allow eight conditions at most so use byte bit field to get the
	Spec []Condition
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

// Or append an logical or condition to conditions, Or has lower precedence than And, note that conditions a.Or(b).And(c) means a || (b&&c),
func (c *Conditions) Or(input Condition) {
	c.Spec = append(c.Spec, input)
}

// And append an logical and condition to conditions, And has higher precedence than Or, note that conditions a.Or(b).And(c) means a || (b&&c),
func (c *Conditions) And(input Condition) {
	i := len(c.Spec)
	c.Spec = append(c.Spec, input)
	if i != 0 {
		c.Op = c.Op | (0x1 << (uint(i) - 1))
	}
}
