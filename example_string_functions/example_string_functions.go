package example_string_functions

import (
	"fmt"
	"strings"
)

var print = fmt.Println

func InitExampleStringFunction() {
	print("Contains:  	", strings.Contains("golang", "lang"))
	print("Has Prefix: 	", strings.HasPrefix("golang", "gol"))
	print("Has Suffix: 	", strings.HasSuffix("golang!", "g"))
	print("Count: 		", strings.Count("golang", "g"))
	print("Index: 		", strings.Index("golang", "la"))
	print("Repeat:		", strings.Repeat("go ", 10))
	print("Replace: 	", strings.Replace("golang!", "g", "$", -1))
	print("Split:		", strings.Split("g-o-l-a-n-g", "-"))
	print("Split:		", strings.Split("go lang", " "))
	print("Join:		", strings.Join([]string{"g", "o", "l", "a", "n", "g"}, "-"))
	print("Join: 		", strings.Join([]string{"go", "lang"}, " "))
	print("ToUpper:		", strings.ToUpper("GoLang"))
	print("ToLower:		", strings.ToLower("GoLang"))

	fmt.Printf("booean\t\tvalue = %t,\nInteger\t\tvalue = %d\n", true, 1)
	// fmt.Printf("Formatted string", ...args)

	fmt.Printf("int = %d\n", 2)    // 00 01 10
	fmt.Printf("binary = %b\n", 4) // 00 01 10 11 100
	fmt.Printf("char = %c\n", 97)
	fmt.Printf("float value = %f\n", 1.1)
	fmt.Printf("String value = %s\n", "golang")

	l := location{
		lat: 21.0123,
		lng: 22.2032,
	}
	fmt.Printf("Struct 1 = %v\n", l)
	fmt.Printf("Struct 2 = %+v\n", l)
	fmt.Printf("Struct 3 = %#v\n", l)

	fmt.Printf("|%10s|%10s|%10s|\n", "==========", "==========", "==========")
	fmt.Printf("|%10s|%10s|%10s|\n", "", "MIN PRICE", "MAX PRICE")
	fmt.Printf("|%10s|%10s|%10s|\n", "==========", "==========", "==========")
	fmt.Printf("|%10s|%10d|%10d|\n", "a", 10, 12)
	fmt.Printf("|%10s|%10d|%10d|\n", "b", 1000, 1200)
	fmt.Printf("|%10s|%10d|%10d|\n", "c", 1, 15000)
	fmt.Printf("|%10s|%10s|%10s|\n", "==========", "==========", "==========")
	fmt.Printf("|%-10d|%-10d|\n", 10, 12)
	fmt.Printf("|%-10d|%-10d|\n", 1000, 1200)
	fmt.Printf("|%-10d|%-10d|\n", 1, 15000)
	fmt.Printf("|%10s|%10s|\n", "==========", "==========")

	s := fmt.Sprintf("Value of a = %d\n", 10)
	fmt.Println(s)
}

type location struct {
	lat, lng float64
}
