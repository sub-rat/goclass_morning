package json_example

import (
	"encoding/json"
	"fmt"
)

type address struct {
	Country string `json:"country"`
	Street  string `json:"street"`
}

type person struct {
	FirstName string              `json:"first_name"`
	LastName  string              `json:"last_name"`
	Age       uint8               `json:"age"`
	Address   address             `json:"address"`
	Phone     []map[string]string `json:"phone"`
}

type response struct {
	Page int      `json:"page"`
	Data []person `json:"data"`
}

var p = fmt.Println

func InitJsonExample() {
	jsonExample1, _ := json.Marshal(true)
	p(string(jsonExample1))

	jsonExample2, _ := json.Marshal(10)
	p(string(jsonExample2))

	jsonExample3, _ := json.Marshal("golang")
	p(string(jsonExample3))

	stringInput := []string{"go", "language", "hello"}
	p(stringInput)
	jsonExample4, _ := json.Marshal(stringInput)
	p(string(jsonExample4))

	mapInput := map[string]string{"first_name": "Ram", "last_name": "Berma"}
	p(mapInput)
	jsonExample5, _ := json.Marshal(mapInput)
	p(string(jsonExample5))

	mapInput2 := map[string]interface{}{
		"first_name": "Ram", "last_name": "Berma", "age": 10,
		"phone": []string{"899089", "8989"},
	}
	p(mapInput2)
	jsonExample6, _ := json.Marshal(mapInput2)
	p(string(jsonExample6))

	person1 := person{
		FirstName: "Ram",
		LastName:  "Berma",
		Address: address{
			Country: "Nepal",
			Street:  "Buddhanager",
		},
		Phone: []map[string]string{
			{
				"home": "9090909",
			},
			{
				"office": "1111111",
			},
		},
	}
	person2 := person{
		FirstName: "Ram1",
		LastName:  "Berma1",
		Age:       10,
		Address: address{
			Country: "Nepal1",
			Street:  "Buddhanager1",
		},
		Phone: []map[string]string{
			{
				"home": "9090909",
			},
			{
				"office": "1111111",
			},
		},
	}

	response := response{
		Page: 1,
		Data: []person{
			person1,
			person2,
		},
	}

	fmt.Printf("%+v\n", response)
	jsonExample8, _ := json.Marshal(response)
	p(string(jsonExample8))
}
