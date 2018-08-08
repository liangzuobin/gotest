package main

import (
	"fmt"
	"reflect"
)

// Value ...
type Value struct {
	gallonsPerMin int
	enabled       bool
}

// Control ...
type Control struct {
	gallonsPerMin int
	enabled       bool
	Value
}

func main() {
	c := Control{
		gallonsPerMin: 700,
		enabled:       false,
		Value: Value{
			gallonsPerMin: 800,
			enabled:       true,
		},
	}
	fmt.Printf("c = %#v \n", c)

	{
		v := reflect.ValueOf(&(c.Value.enabled)).Elem()
		// f := v.FieldByName("Value")
		// fmt.Printf("field = %+v, canset = %t \n", f, f.CanSet())
		fmt.Printf("canset = %t \n", v.CanSet())
		v.SetBool(false)
		fmt.Printf("c = %#v \n", c)
	}

	{
		t := reflect.TypeOf(c)
		f, _ := t.FieldByName("enabled")
		fmt.Printf("field = %+v \n", f)
	}
}
