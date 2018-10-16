package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditions_Judge(t *testing.T) {
	c := new(Conditions)
	// means key1=1&&key1=2
	c.And(Condition{
		Key:   "key1",
		Value: "1",
	})
	c.And(Condition{
		Key:   "key1",
		Value: "2",
	})

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.False(t, c.Judge(input), "actually is key1=1&&key1=2, need to be false")
}
func TestConditions_Judge2(t *testing.T) {
	c := new(Conditions)
	// means key1=1||key1=2
	c.And(Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Or(Condition{
		Key:   "key1",
		Value: "2",
	})

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.True(t, c.Judge(input), "actually is key1=1||key1=2, need to be true")
}

func TestConditions_Judge3(t *testing.T) {
	c := new(Conditions)
	c.And(Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Or(Condition{
		Key:   "key1",
		Value: "2",
	})
	c.And(Condition{
		Key:   "key2",
		Value: "3",
	})

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.True(t, c.Judge(input), "actually is key1=1||key1=2&&key2=3, need to be true")
}

func TestConditions_Judge4(t *testing.T) {
	c := new(Conditions)
	c.And(Condition{
		Key:   "key1",
		Value: "1",
	})
	c.And(Condition{
		Key:   "key1",
		Value: "2",
	})
	c.Or(Condition{
		Key:   "key2",
		Value: "3",
	})
	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.True(t, c.Judge(input), "actually is key1=1&&key1=2||key2=3, need to be true")
}

func TestConditions_Judge5(t *testing.T) {
	c := new(Conditions)
	c.And(Condition{
		Key:   "key1",
		Value: "1",
	})
	c.And(Condition{
		Key:   "key1",
		Value: "2",
	})
	c.Or(Condition{
		Key:   "key2",
		Value: "4",
	})
	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.False(t, c.Judge(input), "actually is key1=1&&key1=2||key2=3, need to be false")
}

func TestConditions_Judge6(t *testing.T) {
	c := new(Conditions)
	// means key1=1&&key2=2&&key3=3
	c.And(Condition{
		Key:   "key1",
		Value: "1",
	})
	c.And(Condition{
		Key:   "key2",
		Value: "2",
	})
	c.And(Condition{
		Key:   "key3",
		Value: "3",
	})

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
func TestConditions_Judge7(t *testing.T) {
	c := new(Conditions)
	c.And(Condition{
		Key:   "key1",
		Value: "1",
	})
	c.And(Condition{
		Key:   "key2",
		Value: "2",
	})
	c.And(Condition{
		Key:   "key3",
		Value: "3",
	})
	c.And(Condition{
		Key:   "key4",
		Value: "4",
	})
	c.And(Condition{
		Key:   "key5",
		Value: "5",
	})
	c.And(Condition{
		Key:   "key6",
		Value: "6",
	})
	c.And(Condition{
		Key:   "key7",
		Value: "7",
	})
	c.And(Condition{
		Key:   "key8",
		Value: "8",
	})

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "2"
	input["key3"] = "3"
	input["key4"] = "4"
	input["key5"] = "5"
	input["key6"] = "6"
	input["key7"] = "7"
	assert.False(t, c.Judge(input), "actually is key1=1&&key2=2&&key3=3, need to be false")

	input = make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "2"
	input["key3"] = "3"
	input["key4"] = "4"
	input["key5"] = "5"
	input["key6"] = "6"
	input["key7"] = "7"
	input["key8"] = "8"
	input["key9"] = "9"
	assert.True(t, c.Judge(input), "actually is key1=1&&key2=2&&key3=3, need to be true")
}
func TestConditions_Judge8(t *testing.T) {
	c := new(Conditions)
	c.And(Condition{
		Key:   "key1",
		Value: "1",
	})
	c.And(Condition{
		Key:   "key2",
		Value: "2",
	})
	c.And(Condition{
		Key:   "key3",
		Value: "3",
	})
	c.And(Condition{
		Key:   "key4",
		Value: "4",
	})
	c.And(Condition{
		Key:   "key5",
		Value: "5",
	})
	c.And(Condition{
		Key:   "key6",
		Value: "6",
	})
	c.And(Condition{
		Key:   "key7",
		Value: "7",
	})
	c.And(Condition{
		Key:   "key8",
		Value: "8",
	})
	c.And(Condition{
		Key:   "key9",
		Value: "91",
	})
	c.And(Condition{
		Key:   "key10",
		Value: "10",
	})

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "2"
	input["key3"] = "3"
	input["key4"] = "4"
	input["key5"] = "5"
	input["key6"] = "6"
	input["key7"] = "7"
	assert.False(t, c.Judge(input), "actually is key1=1&&key2=2&&key3=3, need to be false")

	input = make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "2"
	input["key3"] = "3"
	input["key4"] = "4"
	input["key5"] = "5"
	input["key6"] = "6"
	input["key7"] = "7"
	input["key8"] = "8"
	input["key9"] = "91"
	input["key10"] = "10"
	input["key11"] = "11"
	assert.True(t, c.Judge(input), "actually is key1=1&&key2=2&&key3=3, need to be true")
}
