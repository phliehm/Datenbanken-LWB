package main

import (
		"fmt"
		//"strconv"
		)
	
func interfaceToString(i interface{}) string {
	switch v := i.(type) {
		case string:
			return v
		case []rune:
			return string(v)
		case []uint8:
			return string(v)
		case float64,int64:
			return fmt.Sprint(v)
		default:
			return ""
	}
}
	
func main() {
	var z float64
	z = 1.3
	fmt.Println(fmt.Sprint(z))
	
	var i interface{}
	var x rune
	x = '.'
	y := []rune{'1','.','3'}
	i = []uint8{'1','.','3'}
	fmt.Println(string(x))
	fmt.Println(string(fmt.Sprint(y)))
	fmt.Println(y)
	fmt.Println(interfaceToString(i))

}

