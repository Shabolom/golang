package main

import (
	converter "example.com/m/convert_string_to_camel_case"
	"fmt"
)

func main() {
	result := converter.ConvertString("hui_nahui-hui")
	fmt.Println(result)
}
