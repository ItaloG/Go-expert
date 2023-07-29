package main

import "fmt"

func main() {
	var minhaVar interface{} = "Italo Gabriel"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res %v e o resultado de ok é %v", res, ok)
	res2 := minhaVar.(int)
	fmt.Fprintf("O valor de res2 é %v", res2)

}
