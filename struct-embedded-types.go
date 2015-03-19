package main
import ("fmt"; "math")

type Rectangle struct {

	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

type Circle struct {

	x float64
	y float64
	r float64
}

type Person struct{

	Name string
}

type Android struct {

	Person 
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi My name is", p.Name)
}
func (c *Circle) area() float64 {
	
	return math.Pi * c.r*c.r
}



func main() {
	
	//var rx1, ry1 float64 = 0, 0
	//var rx2, ry2 float64 = 10, 10
	//var cx, cy, cr float64 = 0, 0, 5

	r := Rectangle{0, 0, 10, 10}

	//c := new(Circle)
	c := Circle{0, 0, 5}
	//p := Person{"Karthik"}

	//a := new(Android)
	//a.Person.Talk()
	p := Person{"Karthik"}
	//
	p.Talk()

	a := new(Android)
	a.Talk()

	fmt.Println((r.area())) 
	fmt.Println((c.area()))
}