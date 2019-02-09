package main

import "fmt"

func main() {
	request := map[string]interface{}{"abc": []int{1, 2}}
	i, _ := request["abc"].([]string)
	fmt.Println(i)
}
