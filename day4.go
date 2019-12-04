package main

import (
    "fmt"
)

func main() {

    counter := 0
    
    for i := 248345; i < 746315; i++ {
        if(isValid(i)) {
            counter++
        }
    }
    
    fmt.Println(counter)
}

func isValid(number int) (bool) {
    var digits [6]int = getDigits(number)
    if(
        digitsIncreasing(digits) && consecutiveDigits(digits) && 
        containsAStrictDouble(digits)) {
        return true
    }
    return false
}

func consecutiveDigits(digits [6]int) (bool) {
    return (
        (digits[0] == digits[1]) ||
        (digits[1] == digits[2]) ||
        (digits[2] == digits[3]) ||
        (digits[3] == digits[4]) ||
        (digits[4] == digits[5]))
}

func containsAStrictDouble(digits [6]int) (bool) {
    return (
        ((digits[0] == digits[1]) && (digits[1] != digits[2])) ||
        ((digits[1] == digits[2]) && (digits[2] != digits[3]) && (digits[1] != digits[0])) ||
        ((digits[2] == digits[3]) && (digits[3] != digits[4]) && (digits[2] != digits[1])) ||
        ((digits[3] == digits[4]) && (digits[4] != digits[5]) && (digits[3] != digits[2])) ||
        ((digits[4] == digits[5]) && (digits[4] != digits[3])))
}

func digitsIncreasing(digits [6]int) (bool) {
    for i, digit := range digits {
        if(!greater(digit, digits[i:6])) {
            return false            
        }
    }
    return true
}

func greater(digit int, slice []int) (bool) {
    for _, compDigit := range slice {
        if (digit > compDigit) {
            return false
        }
    }
    return true
}

func getDigits(number int) ([6]int) {
    var digits [6]int
    digits[5] = number % 10
    digits[4] = (number/10) % 10
    digits[3] = (number/100) % 10
    digits[2] = (number/1000) % 10
    digits[1] = (number/10000) % 10
    digits[0] = (number/100000) % 10
    return digits
}

func getNumber(digits [6]int) (number int) {
    number = digits[5] + digits[4]*10 + digits[3]*100 + digits[2]*1000 + digits[1]*10000 + digits[0]*100000
    return number
}