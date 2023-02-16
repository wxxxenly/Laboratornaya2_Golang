package main
import "fmt"

func main() {
    var n int
    fmt.Println("Введите число: ")
    fmt.Scanf("%d\n", &n)	
    var n1, rez float32 = 1, 1
    for i := 1; i <= n; i++{
        n1 = n1 * float32(i)
        rez = rez + 1 / n1
    }
    fmt.Println("Результат: ",  rez)
}
