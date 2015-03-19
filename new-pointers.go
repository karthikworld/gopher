package main
import "fmt"

func one(xPtr *int) {
	*xPtr = 2
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}