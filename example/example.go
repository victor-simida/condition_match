package main

import (
	"fmt"

	condition "github.com/victor-simida/condition_match"
)

func main() {
	var c condition.Conditions
	c.Spec = append(c.Spec, condition.Condition{
		Key:   "key1",
		Value: "1",
	})
	c.Spec = append(c.Spec, condition.Condition{
		Key:   "key2",
		Value: "2",
	})
	c.Spec = append(c.Spec, condition.Condition{
		Key:   "key3",
		Value: "3",
	})

	c.Op = c.Op | 0x1<<0
	c.Op = c.Op | 0x1<<1 // means key1=1&&key2=2&&key3=3

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "2"
	input["key3"] = "3"

	fmt.Println(c.Judge(input))

}
