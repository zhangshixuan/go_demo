package model

type person struct {
	Name   string
	salary float64
}

//构造方法
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//set方法
func (p *person) SetSalary(salary float64) {
	p.salary = salary
}

//get方法
func (p *person) GetSalary() float64 {
	return p.salary
}
