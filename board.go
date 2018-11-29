package main

import (
	"errors"
	"math/rand"
)

const (
	// StatusMine field contains mine
	StatusMine = -1
)

var (
	// ErrOutOfBoard coordinates are beyond the board boundary
	ErrOutOfBoard = errors.New("point coordinates are beyond the board boundary")
)

// Coords is field coordinate set
type Coords struct {
	x int
	y int
}

// Field is a board field
type Field struct {
	mine bool
}

// Board is a bame board structure
type Board struct {
	fields [][]*Field
	width  int
	height int
}

// NewBoard generates new board
//
// width, height - width and height of board
// x, y - coordinates of starting point
// mineNumber - number of mines
func NewBoard(width, height, x, y, mineNumber uint) (*Board, error) {
	if x > width || y > height {
		return nil, ErrOutOfBoard
	}
	return &Board{
		fields: generateFields(int(width), int(height), int(x), int(y), int(mineNumber)),
		width:  int(width),
		height: int(height),
	}, nil
}

func generateFields(width, height, x, y, mineNumber int) [][]*Field {
	points := map[Coords]struct{}{Coords{x, y}: struct{}{}}

	num := mineNumber

	for num > 0 {
		xNext := rand.Intn(width)
		yNext := rand.Intn(height)
		if _, ok := points[Coords{xNext, yNext}]; !ok {
			points[Coords{xNext, yNext}] = struct{}{}
		}
		num--
	}

	delete(points, Coords{x, y})

	board := make([][]*Field, 0, width)

	for xIdx := 0; xIdx < width; xIdx++ {
		row := make([]*Field, 0, height)

		for yIdx := 0; yIdx < height; yIdx++ {
			_, isMine := points[Coords{xIdx, yIdx}]
			row = append(row, &Field{isMine})
		}

		board = append(board, row)
	}
	return board
}

// GetStatus returns field status - mine (-1) or mined neighbours count (0..9)
func (b *Board) GetStatus(x, y uint) (int, error) {
	if int(x) > b.width || int(y) > b.height {
		return 0, ErrOutOfBoard
	}
	if b.fields[x][y].mine {
		return StatusMine, nil
	}
	return b.getNeighbourMCount(int(x), int(y)), nil
}

func (b *Board) getNeighbourMCount(x, y int) int {
	var count int
	for xIdx := x - 1; xIdx < x+1; xIdx++ {
		if xIdx < 0 || xIdx > b.width {
			continue
		}
		for yIdx := y - 1; yIdx < y+1; yIdx++ {
			if yIdx < 0 || yIdx > b.height || xIdx == x && yIdx == y {
				continue
			}
			if b.fields[x][y].mine {
				count++
			}
		}
	}
	return count
}
