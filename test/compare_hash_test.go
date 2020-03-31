package test

import (
	"bufio"
	"fmt"
	"murmur3"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestHashMurmur( t *testing.T)  {
	lines, err := readLines("murmur3.txt")
	if err != nil {
		fmt.Println(err)
	}

	check := true

	for _ , line := range lines {
		input := strings.Split(line, " ")
		if len(input) == 3 {
			//fmt.Println(input[0], input[1], input[2])
			seed,_ := strconv.Atoi(input[0])
			result,_ := strconv.Atoi(input[2])
			postID := input[1]
			output := murmur3.HashString(int64(seed), postID).AsInt()
			if output != result {
				 fmt.Println(line)
				 check = false
			}
		} else {
			fmt.Println("Len(line) != 3")
		}
	}

	if !check {
		fmt.Println("False!!!")
	} else {
		fmt.Println("OK!!!")
	}

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


