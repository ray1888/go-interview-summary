package interview

import "fmt"

func MapEqual() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	/*
		with below code.Can't compile , because map can't be equal with map ,
		needed to check one by one
	*/

	// sm1 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}
	// sm2 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }
}
