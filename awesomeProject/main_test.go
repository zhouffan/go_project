package main_test

import (
	"awesomeProject/simplemath"
	"fmt"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Println("...")
	fmt.Println("...")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 3{
		Usage()
		return
	}
	param1 := args[1]
	param2 := args[2]
	//param3 := args[3]
	switch param1 {
	case "add":
		if len(args) != 4 {
			fmt.Println(",,")
		}
		atoi, err := strconv.Atoi(param2)
		atoi1, err1 := strconv.Atoi(args[3])
		//有错误 return
		if err != nil || err1 != nil{
			fmt.Println("x")
			return
		}
		ret := simplemath.Add(atoi, atoi1)
		fmt.Println("", ret)
	case "sqrt":
		fmt.Println("")
	default:

	}
}