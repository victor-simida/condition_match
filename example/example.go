package main

import (
	"encoding/json"
	"fmt"

	condition "github.com/victor-simida/condition_match"
)

func main() {
	var c condition.Conditions

	// means key1=1&&key2=2&&key3=3
	c.And(condition.Condition{
		Key:   "key1",
		Value: "1",
	})
	c.And(condition.Condition{
		Key:   "key2",
		Value: "2",
	})
	c.And(condition.Condition{
		Key:   "key3",
		Value: "3",
	})

	o, _ := json.Marshal(&c)
	fmt.Println(c.Operators)
	fmt.Println(o)
	fmt.Println(string(o))
	fmt.Println(len(o))
	fmt.Println(len(string(o)))

	input := make(map[string]interface{})
	input["key1"] = "1"
	input["key2"] = "2"
	input["key3"] = "3"

	fmt.Println(c.Judge(input))

}
