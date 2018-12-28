package main

import "github.com/ilyakaznacheev/gosweeper"

func main() {
	l, _ := gosweeper.NewLauncher(20, 20, 300)
	l.Start()
}
