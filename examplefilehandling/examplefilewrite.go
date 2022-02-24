package examplefilehandling

import (
	"bufio"
	"fmt"
	"os"
)

func InitFileWrite() {
	data := []byte("golang!\nhello there\n")
	err := os.WriteFile("example.txt", data, 0644)
	p(err)

	file, err := os.Create("example.txt")
	p(err)
	defer file.Close()

	n1, err := file.Write(data)
	p(err)
	fmt.Printf("data written = %d\n", n1)

	n2, err := file.Write(data)
	p(err)
	fmt.Printf("data written = %d\n", n2)

	n3, err := file.WriteString("I am writing String!\n")
	p(err)
	fmt.Printf("data written = %d\n", n3)

	file.Sync()

	writter := bufio.NewWriter(file)
	n4, err := writter.WriteString("buffered String\n")
	p(err)
	fmt.Println("byte written = ", n4)

	writter.Flush()

}
