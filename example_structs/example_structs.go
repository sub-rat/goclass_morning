package examplestructs

import "fmt"

// Student => name = "ram",id = 1, roll_no = 12, address = home=> kathmandu, phone

type address struct {
	street   string
	houseno  string
	location string
}

type person struct {
	id    int
	name  string
	phone []string
}

//embedding
type student struct {
	person  // embedding person in strudent
	rollNo  int
	address map[string]address //
}

type teacher struct {
	person          // embedding person in teacher
	id              int
	faculty         string
	teachingSubject string
}

type admin struct {
	person   // embedding person in admin
	position string
}

func creatStudent(id int, name string) *student {
	s := student{
		person: person{
			id:   id,
			name: name,
		},
	}
	return &s
}

// methods that is only accessable by instance of person struct
func (s person) Display() {
	fmt.Println("In Person Display func")
	fmt.Println("id = ", s.id)
	fmt.Println("name = ", s.name)
	// fmt.Println("rollno = ", s.rollNo)
	// fmt.Println("address = ", s.address)
}

// methods that is only accessable by instance of teacher struct
func (s teacher) Display() {
	fmt.Println("In Teacher Display func")
	fmt.Println("id = ", s.id)
	fmt.Println("name = ", s.name)
	// fmt.Println("rollno = ", s.rollNo)
	// fmt.Println("address = ", s.address)
}

// function
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
		person: person{
			id:    1,
			name:  "Ram",
			phone: []string{"9090909090", "89898989"},
		},
		rollNo:  1,
		address: std1Address,
	}

	fmt.Printf("%+v\n", std1)

	std2 := student{
		person: person{
			id:   1,
			name: "Ram",
		},
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

	teacher := teacher{
		person: person{
			id:   10,
			name: "hari",
		},
		id:              2,
		faculty:         "science",
		teachingSubject: "science",
	}

	fmt.Println("=========example embeddings==========")
	fmt.Println(teacher.id)
	fmt.Println(teacher.person.id)
	teacher.Display()
	teacher.person.Display()

	a1 := admin{
		person: person{
			id:   4,
			name: "gopal",
		},
		position: "woker1",
	}

	a1.Display()

}
