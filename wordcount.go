package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	if str, err := readInput(); err!=nil {
		fmt.Println(err)
	} else {
		num, err2 := getWordsCount(str)
		if err2!=nil {
			fmt.Println(err2)
		} else {
			fmt.Println(num)
		}
	}
}

func readInput() (string, error){
	args := os.Args[1:]
	if len(args) != 1 {
		return "", errors.New("wrong amount of arguments, needed one string")
	}
	str := string(args[0])
	return str, nil
}

func getWordsCount(str string) (int, error) {
	if str == "" {
		return 0, errors.New("empty string")
	}
	split := strings.Split(str, " ")
	return len(split), nil
}