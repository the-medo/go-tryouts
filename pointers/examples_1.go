package pointers

import "fmt"

func RunExamples1() {
	// ================= int
	x := 1

	modifyValueIntNoPointer(x)
	fmt.Println("Modify value no pointer: ", x)

	modifyValueIntPointer(&x)
	fmt.Println("Modify value pointer: ", x)

	// ============== string
	s := "String value"
	modifyValueStringNoPointer(s)
	fmt.Println("Modify value string no pointer: ", s)

	modifyPartOfValueStringNoPointer(s)
	fmt.Println("Modify part of value string no pointer: ", s)

	modifyValueStringPointer(&s)
	fmt.Println("Modify value string pointer: ", s)
}

func modifyValueIntNoPointer(x int) {
	x = 20
}

func modifyValueIntPointer(x *int) {
	*x = 10
}

func modifyValueStringNoPointer(x string) {
	x = "123"
}

func modifyPartOfValueStringNoPointer(x string) {
	tmp := []rune(x)
	tmp[3] = '1'
	x = string(tmp)
}

func modifyValueStringPointer(x *string) {
	*x = "123"
}
