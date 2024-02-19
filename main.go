package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer -type=Pill
//go:generate go run go.uber.org/mock/mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=mock_$GOPACKAGE/mock_$GOFILE

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
)

type Tester interface {
	Test()
}

func main() {
	fmt.Println(Placebo.String())
}
