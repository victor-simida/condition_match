package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditions_Judge(t *testing.T) {
	var c Conditions
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "2",
	})

	c.Op = 0 | (0x1 << 0) // means key1=1&&key1=2
	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.False(t, c.Judge(input), "actually is key1=1&&key1=2, need to be false")
}
func TestConditions_Judge2(t *testing.T) {
	var c Conditions
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "2",
	})

	c.Op = 0 // means key1=1||key1=2
	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.True(t, c.Judge(input), "actually is key1=1||key1=2, need to be true")
}

func TestConditions_Judge3(t *testing.T) {
	var c Conditions
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "2",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key2",
		Value: "3",
	})

	c.Op = 0 | 0x1<<1 // means key1=1||key1=2&&key2=3
	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.True(t, c.Judge(input), "actually is key1=1||key1=2&&key2=3, need to be true")
}

func TestConditions_Judge4(t *testing.T) {
	var c Conditions
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "2",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key2",
		Value: "3",
	})

	c.Op = 0 | 0x1<<0 // means key1=1&&key1=2||key2=3
	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.True(t, c.Judge(input), "actually is key1=1&&key1=2||key2=3, need to be true")
}

func TestConditions_Judge5(t *testing.T) {
	var c Conditions
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "2",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key2",
		Value: "4",
	})

	c.Op = 0 | 0x1<<0 // means key1=1&&key1=2||key2=3
	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.False(t, c.Judge(input), "actually is key1=1&&key1=2||key2=3, need to be false")
}

func TestConditions_Judge6(t *testing.T) {
	var c Conditions
	c.Spec = append(c.Spec, Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key2",
		Value: "2",
	})
	c.Spec = append(c.Spec, Condition{
		Key:   "key3",
		Value: "3",
	})

	c.Op = c.Op | 0x1<<0
	c.Op = c.Op | 0x1<<1 // means key1=1&&key2=2&&key3=3
	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "2"
	assert.False(t, c.Judge(input), "actually is key1=1&&key2=2&&key3=3, need to be false")

	input = make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "2"
	input["key3"] = "3"
	assert.True(t, c.Judge(input), "actually is key1=1&&key2=2&&key3=3, need to be true")
}
