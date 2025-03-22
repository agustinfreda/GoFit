package main

import "fmt"

func main() {
	pesoPaciente := PedirPeso()
	fmt.Println(pesoPaciente)
}

func PedirPeso() float32 {
	var peso float32
	fmt.Print("Ingrese su peso registrado hoy: ")
	fmt.Scanf("%g", &peso)
	return peso
}
