package interview

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	// compile error, string empty only is "", can't be nil
	// return nil, false
	return "", false
}

func StringError() {

	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)

}
