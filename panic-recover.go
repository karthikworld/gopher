package main 
import "fmt"

defer func() {
           str := recover()
           fmt.Println(str)
     }()
     panic("PANIC")
}     