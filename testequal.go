package main

type obj struct {
	ID   uint
	Name string
}

func main() {

	o1 := new(obj)
	o1.ID = 1
	o1.Name = "jack"

	o2 := new(obj)
}
