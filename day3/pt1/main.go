package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

// ReadLines returns lines from a text file as a slice
func ReadLines(filename string) ([]string, error) {
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
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

func GetLength(lines []string) int {
	return len(lines)
}

func GetWidth(lines []string) int {
	return len(lines[0])
}

func (p Point) GetX(width int) int {
	xCoord := p.x + 3
	if xCoord+1 > width {
		xCoord = xCoord - width
	}
	return xCoord
}

func (p Point) IdentifyObject(lines []string) string {
	return string(lines[p.y][p.x])
}

func IsTree(obj string) bool {
	switch obj {
	case "#":
		return true
	case ".":
		return false
	default:
		fmt.Printf("Unkown object")
		return false
	}
}

func IterateGrid(lines []string) int {
	length := GetLength(lines)
	width := GetWidth(lines)

	TreeCount := 0
	p := Point{}
	for i := 1; i < length; i++ {
		p.x = p.GetX(width) // 3, 6, 9..
		p.y = i             // 1, 2, 3..
		obj := p.IdentifyObject(lines)

		if IsTree(obj) {
			TreeCount++
		}
	}
	return TreeCount
}

func main() {
	lines, _ := ReadLines("../input.txt")
	treeCount := IterateGrid(lines)
	fmt.Println(treeCount)
}
