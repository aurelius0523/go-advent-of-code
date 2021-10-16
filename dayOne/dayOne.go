package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	absPath, _ := filepath.Abs("./dayOne/input.txt")
	dat, err := os.ReadFile(absPath)
	check(err)

	masses := strings.Split(string(dat), "\r\n")
	totalMass := 0

	for _, mass := range masses {
		j, _ := strconv.Atoi(mass)
		totalMass += counterUpper(j)
	}

	fmt.Print(totalMass)
}

func counterUpper(mass int) int {
	return (mass / 3) - 2
}
