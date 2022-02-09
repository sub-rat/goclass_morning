package exampledatastructure

import "fmt"

func initExampleMap() {
	m := make(map[string]int) // map[key-data-type]value-data-type
	//  "a" => 97
	//  "b" => 98
	//  "c" => 99
	m["a"] = 97
	m["b"] = 98
	m["c"] = 99
	fmt.Println("lenght of map = ", len(m))
	fmt.Println("map = ", m)

	if v1, ok := m["a"]; ok {
		fmt.Println("value of a = ", v1)
	}
	m["d"] = 100
	fmt.Println("map = ", m)
	fmt.Println("lenght of map = ", len(m))

	delete(m, "c")
	fmt.Println("map = ", m)

	m2 := map[string]int{"a": 97, "b": 98, "c": 99}
	m2["z"] = 110
	delete(m2, "a")
	fmt.Println("map2 = ", m2)

	// exampleRange()
}
