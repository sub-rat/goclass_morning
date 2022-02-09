package exampledatastructure

import "fmt"

func exampleRangeWithArray() {
	arrayExample2 := [5]int{10, 12, 23, 25, 40}
	arrayExample2[1] = 15

	for i := 0; i < len(arrayExample2); i++ {
		fmt.Printf("Index = %d and Value = %d\n", i, arrayExample2[i])
	}

	for _, value := range arrayExample2 {
		// fmt.Printf("Index = %d and Value = %d\n", index, value)
		fmt.Printf("Value = %d\n", value)
	}
}

func exampleRangeWithMap() {
	m2 := map[string]int{"a": 97, "b": 98, "c": 99}
	m2["z"] = 110
	delete(m2, "a")
	fmt.Println("map2 = ", m2)

	for key, value := range m2 {
		fmt.Printf("key = %s, Value = %d\n", key, value)
	}

	for i, value := range "golang" {
		fmt.Println(i, string(value))
	}
}
