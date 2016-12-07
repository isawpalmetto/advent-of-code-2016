package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileLocation := flag.String("f", "", "File to direction input")
	flag.Parse()
	//read stream of comma separated inputs
	f, err := os.Open(*fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	onComma := func(data []byte, atEOF bool) (int, []byte, error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)

	position := Position{
		Facing:         N,
		XPos:           0,
		YPos:           0,
		AllThePosition: make(map[string]bool),
	}
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		position.Turn(string(input[0]))
		toInt, _ := strconv.Atoi(string(input[1:]))
		if position.Move(toInt) {
			fmt.Printf("Reached HQ: Total Distance = %+v\n", position.TaxiDistance())
			return
		}
	}
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

type Position struct {
	Facing         Direction
	XPos           int
	YPos           int
	AllThePosition map[string]bool
}

// Move moves the position a certain amount of spaces, returns true if a position has been seen
func (p *Position) Move(spaces int) bool {
	switch p.Facing {
	case N:
		for i := 0; i < spaces; i++ {
			p.YPos += 1
			if p.Seen() {
				return true
			}
		}
	case E:
		for i := 0; i < spaces; i++ {
			p.XPos += 1
			if p.Seen() {
				return true
			}
		}
	case S:
		for i := 0; i < spaces; i++ {
			p.YPos -= 1
			if p.Seen() {
				return true
			}
		}
	case W:
		for i := 0; i < spaces; i++ {
			p.XPos -= 1
			if p.Seen() {
				return true
			}
		}
	}
	return false
}

func (p *Position) Seen() bool {
	x := strconv.Itoa(p.XPos)
	y := strconv.Itoa(p.YPos)
	if _, ok := p.AllThePosition[x+","+y]; !ok {
		p.AllThePosition[x+","+y] = true
		return false
	}
	return true
}

func (p *Position) Turn(t string) {
	if t == "R" {
		switch p.Facing {
		case N:
			p.Facing = E
		case E:
			p.Facing = S
		case S:
			p.Facing = W
		case W:
			p.Facing = N
		}
	} else if t == "L" {
		switch p.Facing {
		case N:
			p.Facing = W
		case E:
			p.Facing = N
		case S:
			p.Facing = E
		case W:
			p.Facing = S
		}
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *Position) TaxiDistance() int {
	return abs(p.XPos) + abs(p.YPos)
}
