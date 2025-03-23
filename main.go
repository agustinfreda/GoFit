package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	Menu()
}

func obtenerFecha() string {
	fecha := time.Now().Format("02-01-2006")
	return fecha
}

func PedirPeso() float32 {
	var peso float32
	fmt.Print("Ingrese su peso registrado hoy: ")
	fmt.Scanf("%g", &peso)
	return peso
}

func Menu() {
	option := true

	for option {
		fmt.Println("|---| REGISTRO DE PESO |---|")
		fmt.Println("1. Registrar peso de hoy.")
		fmt.Println("2. Salir.")

		var respuesta int
		fmt.Println("Elegir opcion:")
		fmt.Scanf("%d", &respuesta)

		switch respuesta {
		case 1:
			peso := PedirPeso()
			registrarCSV("registro_de_peso.csv", peso)
		case 2:
			fmt.Println("Saliendo...")
			option = false
		}
	}
}

func registrarCSV(filename string, peso float32) error {
	// Verificar si el archivo existe
	existe := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existe = false
	}

	// Abrir archivo en modo append
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo CSV: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir encabezado solo si el archivo no existía
	if !existe {
		headers := []string{"fecha", "peso"}
		if err := writer.Write(headers); err != nil {
			return fmt.Errorf("error al escribir encabezado en CSV: %v", err)
		}
	}

	// Agregar nueva línea con el peso
	fecha := obtenerFecha()
	peso_recibido := strconv.FormatFloat(float64(peso), 'f', -1, 32)
	if err := writer.Write([]string{fecha, peso_recibido}); err != nil {
		return fmt.Errorf("error al escribir en CSV: %v", err)
	}

	return nil
}
