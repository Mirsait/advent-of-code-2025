package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func runRender(initialState State) {
	nextState := initialState
	r := renderWithIndex(1)
	for {
		var count int
		nextState, count = rpuzzle(nextState, r)
		if count == 0 {
			fmt.Println("rendering ended.")
			break
		}
	}
}

func rpuzzle(state State, r func(State)) (State, int) {
	nextState := cloneState(state)
	r(nextState)
	count := 0
	rows := len(state)
	cols := len(state[0])

	for j := range rows {
		for k := range cols {
			if isRoll(state[j][k]) {
				if checkNeighborhood(state, j, k, rows, cols) {
					count++
					nextState[j][k] = 'x'
				}
			}
		}
	}

	r(nextState)                    // render with x
	xReplace(nextState, rows, cols) // clear x
	r(nextState)                    // render
	return nextState, count
}

func xReplace(state State, rows, cols int) {
	for j := range rows {
		for k := range cols {
			if state[j][k] == 'x' {
				state[j][k] = '.'
			}
		}
	}
}

func renderWithIndex(start int) func(state State) {
	index := start
	return func(state State) {
		render(index, state)
		index++
	}
}

func render(num int, state State) {
	w := len(state[0])
	h := len(state)
	var img = image.NewRGBA(image.Rect(0, 0, w*2, h*2))
	for x := range w {
		for y := range h {
			var s = state[x][y]
			xx := x * 3
			yy := y * 3
			var rectColor color.Color
			switch s {
			case '.':
				rectColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
			case 'x':
				rectColor = color.RGBA{G: 255, A: 255}
			default:
				rectColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			}
			img.Set(xx, yy, rectColor)
		}
	}
	filename := fmt.Sprintf("images/%03d.png", num)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
