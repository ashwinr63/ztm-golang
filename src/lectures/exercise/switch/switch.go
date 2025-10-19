//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

func main() {
	switch age := 26; {
	case age == 0:
		// newborn age is 0
		fmt.Println("newborn")
	case age >= 1 && age <= 3:
		// toddler age is 1, 2, or 3
		fmt.Println("toddler")
	case age >= 4 && age <= 12:
		// child age is 4 through 12
		fmt.Println("child")
	case age >= 13 && age <= 17:
		// teenager age is 13-17
		fmt.Println("teenager")
	default:
		// adult age is 18+
		fmt.Println("adult")
	}

	//* Use a `switch` statement to print the following:
	//  - "newborn" when age is 0
	//  - "toddler" when age is 1, 2, or 3
	//  - "child" when age is 4 through 12
	//  - "teenager" when age is 13 through 17
	//  - "adult" when age is 18+
}

