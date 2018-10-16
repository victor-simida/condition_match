package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCondition_Match(t *testing.T) {
	c := Condition{
		Key:   "key1",
		Value: "1",
		Op:    OpNotEqual,
	}

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.False(t, c.Match(input))

	input["key1"] = "12"
	input["key2"] = "3"
	assert.True(t, c.Match(input))
}
func TestCondition_Match2(t *testing.T) {
	c := Condition{
		Key:   "key1",
		Value: "1",
		Op:    OpHasPrefix,
	}

	input := make(map[string]interface{})
	input["key1"] = "23123"
	input["key2"] = "3"
	assert.False(t, c.Match(input))

	input["key1"] = "12"
	input["key2"] = "3"
	assert.True(t, c.Match(input))
}
func TestCondition_Match3(t *testing.T) {
	c := Condition{
		Key:   "key1",
		Value: []string{"1", "2", "3"},
		Op:    OpIn,
	}

	input := make(map[string]interface{})
	input["key1"] = "23123"
	input["key2"] = "3"
	assert.False(t, c.Match(input))

	input["key1"] = "1"
	input["key2"] = "3"
	assert.True(t, c.Match(input))
}
func TestCondition_Match4(t *testing.T) {
	c := Condition{
		Key:   "key1",
		Value: []string{"1", "2", "3"},
		Op:    OpNotIn,
	}

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "3"
	assert.False(t, c.Match(input))

	input["key1"] = "12"
	input["key2"] = "3"
	assert.True(t, c.Match(input))
}
