package main
import "fmt"
//Média Simples
func main() {
	var a,b,c,d float64
	fmt.Scanf("%f\n", &a)
	fmt.Scanf("%f\n", &b)
	fmt.Scanf("%f", &c)
	d = ((a*2)+(b*3)+(c*5))/10
	fmt.Printf("MEDIA = %.5f\n", d)
}