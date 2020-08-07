package main
import "fmt"
//Média Simples
func main() {
	var a,b,c float64
	fmt.Scanf("%f\n", &a)
	fmt.Scanf("%f", &b)
	c = ((a*3.5)+(b*7.5))/11
	fmt.Printf("MEDIA = %.5f\n", c)
}