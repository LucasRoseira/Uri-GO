package main
import "fmt"

func main() {
	var raio float64
	fmt.Scanf("%f\n", &raio)
	raio *= 3.14159 * raio
	fmt.Printf("A=%.4f\n", raio)
}