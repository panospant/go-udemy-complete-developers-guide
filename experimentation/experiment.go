package main

import "fmt"

type Saiyan struct {
	Name string
	Power int
}

func main() {

	// goku := Saiyan{}
	// or
	goku := Saiyan{Name: "Goku"}
	goku.Power = 9000
	goku.Super()
	fmt.Println(goku.Power)

	//if len(os.Args) != 2 {
	//	os.Exit(1)
	//}
	//fmt.Println("It's over", os.Args[0])
}

func (s *Saiyan) Super() {
	s.Power += 10000
}
