package loopexample

import "fmt"

func InitLoopExample() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	/*
			i = 0  ;  sum = 0 + 0  ; i < 10 = true
		 	i = 1  ;  sum = 0 + 1  ; i < 10 = true
			i = 2  ; sum = 1 + 2   ; i < 10 = true
		 	..
		 	i = 10;  		; i < 10 = false
	*/
	fmt.Println("sum = ", sum)

	sum = 1
	for sum < 45 {
		sum += sum
		fmt.Println("Sum In loop = ", sum)
		fmt.Println(sum < 45)
	}
	fmt.Println("sum2=", sum)

	a := 0
	for a < 10 {
		if a == 5 {
			/* skip the iteration */ a = a + 1
			continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}

	fmt.Println("===================")
	fmt.Println("Using Break")
	a = 0
	for a < 10 {
		a++
		if a > 5 {
			/* terminate the loop using break statement */
			break
		}
		fmt.Printf("value of a: %d\n", a)
	}

}
