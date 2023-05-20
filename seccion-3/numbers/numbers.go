package main

import "fmt"

//sumOfDivisorsOfNum is a function that calculates the sum of the divisors of a number.
func sumOfDivisorsOfNum(number int) int {
	var sum int

	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}
	return sum
}

//IsAPerfectNumber is a function that determines if a number is perfect
func IsAPerfectNumber(number int) int {

	sum := sumOfDivisorsOfNum(number)

	if sum == number {
		fmt.Printf("The sum of the divisors of %d is equal to %d. The number is perfect!\n", number, sum)
		return sum
	}

	fmt.Printf("The sum of the divisors of %d is equal to %d. The number is not perfect, please try again.\n", number, sum)

	return sum
}

//AreTheyFriends is a function that determines if two numbers are friends.
func AreTheyFriends(num1, num2 int) (int, int) {

	sumOfDivisorsOfNum1 := sumOfDivisorsOfNum(num1)
	sumOfDivisorsOfNum2 := sumOfDivisorsOfNum(num2)
	if (sumOfDivisorsOfNum1 == num2) && (sumOfDivisorsOfNum2 == num1) {
		fmt.Printf("Numbers are friends! The sum of the divisors of %d has a value of %d that is equal to the value of the second number and the sum of the divisors of %d is equal to the value of the first number with a value of %d.", num1, sumOfDivisorsOfNum1, sumOfDivisorsOfNum2, num1)
		return sumOfDivisorsOfNum1, sumOfDivisorsOfNum2
	}

	fmt.Printf("Numbers are not friends.The sum of the divisors of %d has a value of %d and is different from the sum of the divisors of %d with a value of %d.\n", num1, sumOfDivisorsOfNum1, num2, sumOfDivisorsOfNum2)

	return sumOfDivisorsOfNum1, sumOfDivisorsOfNum2

}

func main() {
	var option int
	var num int
	var num1, num2 int

	fmt.Println("Welcome!")
	for {
		fmt.Println("\nEnter the operation you want to perform:\n\t1. Define if a number is perfect\n\t2. Define if two numbers are friends\n\t3.Exit")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("\nPlease enter the number ")
			fmt.Scanln(&num)
			IsAPerfectNumber(num)
		case 2:
			fmt.Println("\nPlease enter the first number")
			fmt.Scanln(&num1)
			fmt.Println("\nPlease enter the second number")
			fmt.Scanln(&num2)
			AreTheyFriends(num1, num2)

		case 3:
			fmt.Println("Thanks come back soon")
			return
		default:
			fmt.Println("Wrong option, try again")
		}

	}

}
