/*
map is a key-value data structure,it's unordered
map[key]value
 */
package main

import "fmt"

func main() {
	// declare a map
	ages := make(map[string]int)
	weight := map[string]int{"lihua": 50, "ming": 65}
	// declare empty map
	height := map[string]int{}

	// assign age
	ages["lihua"] = 21
	ages["ming"] = 20

	// if key is not exist, return 0
	fmt.Println(height["lihua"])
	delete(weight, "lihua")
}
