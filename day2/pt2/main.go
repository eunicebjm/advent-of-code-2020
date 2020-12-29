package main

import (
"bufio"
"fmt"
"log"
"os"
"regexp"
"strconv"
"strings"
)

type record struct {
	rule string
	password string
}

type policy struct {
	index1 int
	index2 int
	char string
}

// NewRecord formats string into record struct
func NewRecord(line string) *record {
	s := strings.Split(line, ":")

	return &record{
		rule: s[0],
		password: strings.TrimSpace(s[1]),
	}
}

// NewPolicy generates policy struct given the record
func NewPolicy(r record) (*policy, error) {
	re := regexp.MustCompile("[0-9a-z]+")
	limits := re.FindAllString(r.rule, -1)

	min, err := strconv.Atoi(limits[0])
	if err != nil{
		return &policy{0,0,""}, err
	}

	max, err := strconv.Atoi(limits[1])
	if err != nil{
		return &policy{0,0,""}, err
	}

	return &policy{min, max, limits[2]}, nil
}

// IsValid checks if the password has exactly one character
// in the right position
func (r record) IsValid() bool {
	p, err:= NewPolicy(r)
	if err != nil {
		log.Fatal(err)
	}

	var matchCount = 0
	if string(r.password[p.index1 - 1]) == p.char{
		matchCount++
	}
	if string(r.password[p.index2 - 1]) == p.char {
		matchCount++
	}

	if matchCount == 1{
		return true
	}
	return false
}

// ReadLines returns lines from a text file as a slice
func ReadLines(filename string) ([]string, error){
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

// GetValidPasswordCount finds the count of valid passwords from
// rows in a text file
func GetValidPasswordCount(lines []string) int{
	validCount := 0
	for _, line := range lines{
		record := NewRecord(line)
		if record.IsValid(){
			validCount++
		}
	}
	return validCount
}

func main() {
	lines, err := ReadLines("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	count := GetValidPasswordCount(lines)
	fmt.Println(count)
}