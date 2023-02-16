package main
import "fmt"

func main() {
    var y,z float64
    fmt.Println("Введите Y: ")
    fmt.Scanf("%g\n", &y)
    
    for x := 0.1; x <= 0.5; x+=0.1{
        z = ((10*x) - (y/x))
        fmt.Println("X = ", float32(x),";", "Z = ", float32(z))
    } 
}
