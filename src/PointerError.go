package interview

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

/*
if the method is to pointer type rather than actual Student type
this can be error for this
```
cannot use Student literal (type Student) as type People in assignment:
	Student does not implement People (Speak method has pointer receiver)
```
*/

// func (stu *Student) Speak(think string) (talk string) {
func (stu Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func PointerError() {
	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
