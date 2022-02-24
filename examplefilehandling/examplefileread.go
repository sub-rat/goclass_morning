package examplefilehandling

import (
	"bufio"
	"fmt"
	"os"
)

func p(err error) {
	if err != nil {
		panic(err)
	}
}

// File Read and Write operation
func InitFileHandling(fileName string) {
	output, err := os.ReadFile(fileName)
	p(err)
	fmt.Println(string(output))

	file, err := os.Open("test.txt")
	p(err)
	defer file.Close()

	b1 := make([]byte, 3)
	n1, err := file.Read(b1)
	p(err)
	fmt.Printf("output = %s and n = %d\n", string(b1), n1)

	b2 := make([]byte, 2)
	n2, err := file.Read(b2)
	p(err)
	fmt.Printf("output = %s and n = %d\n", string(b2[:n2]), n2)

	o1, err := file.Seek(2, 0)
	p(err)
	fmt.Println("o1 = ", o1)

	b3 := make([]byte, 4)
	n3, err := file.Read(b3)
	p(err)
	fmt.Printf("output = %s and n = %d\n", string(b3[:n3]), n3)

	file.Seek(0, 0)
	r1 := bufio.NewReader(file)
	b4, err := r1.Peek(7)
	p(err)
	fmt.Println(string(b4))
}
