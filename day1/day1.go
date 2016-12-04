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
		Facing: N,
		XPos:   0,
		YPos:   0,
	}
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		position.Turn(string(input[0]))
		toInt, _ := strconv.Atoi(string(input[1:]))
		position.Move(toInt)
	}
	fmt.Printf("Total Distance = %+v\n", position.TaxiDistance())
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

type Position struct {
	Facing Direction
	XPos   int
	YPos   int
}

func (p *Position) Move(spaces int) {
	switch p.Facing {
	case N:
		p.YPos += spaces
	case E:
		p.XPos += spaces
	case S:
		p.YPos -= spaces
	case W:
		p.XPos -= spaces
	}
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
