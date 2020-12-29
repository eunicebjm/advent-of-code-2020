package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type Grid struct {
	lines         []string
	width, length int
	point         Point
}

type Point struct {
	x, y int
}

type Move struct {
	right, down int
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

func (g Grid) GetX(traverseRight int) int {
	x := g.point.x + traverseRight
	if x+1 > g.width {
		x %= g.width
	}
	return x
}

func (g Grid) GetY(traverseDown int) (int, error) {
	y := g.point.y + traverseDown
	if y > g.length {
		return 0, errors.New("index error")
	}
	return y, nil
}

func (g Grid) GetPoint(traverseRight int, traverseDown int) Point {
	x := g.GetX(traverseRight)
	y, err := g.GetY(traverseDown)
	if err != nil {
		return Point{}
	}
	return Point{x: x, y: y}
}

func (g Grid) IdentifyObject() string {
	return string(g.lines[g.point.y][g.point.x])
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

func (g Grid) IterateGrid(traverseRight int, traverseDown int) int {
	TreeCount := 0
	for i := 1; i < g.length; i++ {
		g.point = g.GetPoint(traverseRight, traverseDown)
		if g.point == (Point{}) {
			break
		}
		obj := g.IdentifyObject()

		if IsTree(obj) {
			TreeCount++
		}
	}
	return TreeCount
}

func main() {
	lines, _ := ReadLines("../input.txt")
	g := Grid{
		lines: lines,
		width: len(lines[0]),
		length: len(lines),
		point: Point{},
	}

	slopes := []Move{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	answer := 1
	for _, slope := range slopes {
		treeCount := g.IterateGrid(slope.right, slope.down)
		answer *= treeCount
	}

	fmt.Println(answer)
}
