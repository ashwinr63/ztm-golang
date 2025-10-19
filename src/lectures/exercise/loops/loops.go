//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

func main() {

	for i := 1; i <= 50; i++ {
		divisibleby3 := i%3 == 0
		divisibleby5 := i%5 == 0

		if divisibleby3 && divisibleby5 {
			fmt.Println("FizzBuzz")
		} else if divisibleby3 {
			fmt.Println("Fizz")
		} else if divisibleby5 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

		//  Implement the classic "FizzBuzz" problem using a `for` loop.
		//
		//--Requirements:
		//* Print integers 1 to 50, except:
		//  - Print "Fizz" if the integer is divisible by 3
		//  - Print "Buzz" if the integer is divisible by 5
		//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
		//
		//--Notes:
		//* The remainder operator (%) can be used to determine divisibility
	}
}

