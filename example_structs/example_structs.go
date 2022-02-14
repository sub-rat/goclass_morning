package examplestructs

import "fmt"

// Student => name = "ram",id = 1, roll_no = 12, address = home=> kathmandu, phone

type address struct {
	street   string
	houseno  string
	location string
}

type student struct {
	id      int
	name    string
	rollNo  int
	address map[string]address //
	phone   []string
}

func creatStudent(id int, name string) *student {
	s := student{
		id:   id,
		name: name,
	}
	return &s
}

func (s student) Display() {
	fmt.Println("In Student Display func")
	fmt.Println("id = ", s.id)
	fmt.Println("name = ", s.name)
	fmt.Println("rollno = ", s.rollNo)
	fmt.Println("address = ", s.address)
}

func DisplayStudent(s student) {
	fmt.Println("IN Global display func")
	fmt.Println("id = ", s.id)
	fmt.Println("name = ", s.name)
	fmt.Println("rollno = ", s.rollNo)
	fmt.Println("address = ", s.address)
}

func InitExampleStruct() {

	// std1Address := map[string]string{"home": "buddhanagar", "office": "pulchowk"}
	std1Address := map[string]address{
		"home": {
			street:   "xyz",
			houseno:  "1pz",
			location: "new york",
		},
		"office": {
			street:   "xyz1",
			houseno:  "1pz1",
			location: "new york1",
		},
	}
	std1 := student{
		id:      1,
		name:    "Ram",
		rollNo:  1,
		address: std1Address,
		phone:   []string{"9090909090", "89898989"},
	}

	fmt.Printf("%+v\n", std1)

	std2 := student{
		id:   1,
		name: "Ram",
	}

	std2.rollNo = 2
	std2.address = map[string]address{
		"home": {
			street:   "xyz",
			houseno:  "1pz",
			location: "new york",
		},
	}
	fmt.Printf("%+v\n", std2)

	fmt.Println(std2.name)
	fmt.Println(std2.address["home"])

	std3 := creatStudent(3, "shyam")
	fmt.Println(std3)
	DisplayStudent(std1)
	std1.Display()
	// var a = 10
	// var p *int
	// p = &a
	// fmt.Println(*p)

}
