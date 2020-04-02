package murmur3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestHashMurmur3_128(t *testing.T) {
	lines, err := readLines("/storage/Learn/Work/MyRepositories/murmur3/murmur3.txt")
	if err != nil {
		fmt.Println(err)
	}

	total := 0
	cntCorrect := 0

	for _, line := range lines {
		input := strings.Split(line, " ")
		if len(input) == 3 {
			total++
			seed, _ := strconv.Atoi(input[0])
			result, _ := strconv.Atoi(input[2])
			postID := input[1]
			output := HashString(int64(seed), postID).AsInt()
			if output != result {
				fmt.Println(line)
			} else {
				cntCorrect++
			}
		} else {
			fmt.Println("Len(line) != 3", line)
		}
	}

	if cntCorrect != total {
		fmt.Printf("Exactly %.2f", float32(cntCorrect*100/total))
	} else {
		fmt.Println("Exactly 100% ^v^")
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
