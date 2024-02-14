package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer -type=Pill

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
)

func main() {
	fmt.Println(Placebo.String())
}
