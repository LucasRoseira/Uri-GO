package main
import "fmt"
//Média Simples
func main() {
	var a,b,c,d int
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	fmt.Scanf("%d\n", &c)
	fmt.Scanf("%d", &d)
	d = a*b-c*d
	fmt.Printf("DIFERENCA = %d\n", d)
}