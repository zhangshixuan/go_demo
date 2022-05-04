package model

type student struct {
	Name string
	Age  float64
}

func GetStudent(n string, a float64) *student {
	return &student{
		Name: n,
		Age:  a,
	}
}
