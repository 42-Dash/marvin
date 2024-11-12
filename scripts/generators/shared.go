package utils

import (
	"log"
	"strconv"
	"strings"
)

type Point struct {
	Row, Col int
}

func ParseArr(arguments, separator string, arr ...*int) {
	args := strings.Split(arguments, separator)

	if len(args) != len(arr) {
		log.Fatalf("Expected %d arguments, got %d", len(arr), len(args))
	}

	for i, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal(err)
		}
		*(arr[i]) = num
	}
}
