package main

import (
    "fmt"
    "bufio"
    "strconv"
    "os"
    "log"
)

func main() {
    masses := readFile("day1-input")
    sumFuel := 0
    for _, mass := range masses {
        massInt, _ := strconv.Atoi(mass)
        fuel := getFuel(massInt)
        sumFuel += fuel
    }
    fmt.Println(sumFuel)
}

func getFuel(mass int) (int) {
    if (mass < 7) {
        return 0
    }
    
    fuel := ((mass / 3) - 2) + getFuel((mass/3) - 2)
    return fuel
}

func readFile (filename string) ([]string) {
    var masses []string
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        masses = append(masses, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return masses
}
