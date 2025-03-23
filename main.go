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
		fmt.Println("2. Ver ultimos registros.")
		fmt.Println("3. Ver todos los registros.")
		fmt.Println("4. Ver estadisticas totales.")
		fmt.Println("5. Salir.")

		var respuesta int
		fmt.Println("Elegir opcion:")
		fmt.Scanf("%d", &respuesta)

		switch respuesta {
		case 1:
			peso := PedirPeso()
			registrarCSV("registro_de_peso.csv", peso)
		case 2:
			mostrarUltimosRegistros("registro_de_peso.csv", 7)
		case 3:
			mostrarUltimosRegistros("registro_de_peso.csv", -1)
		case 4:
			EstadisticasTotales("registro_de_peso.csv", -1)
		case 5:
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

func mostrarUltimosRegistros(filename string, n int) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo CSV: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	registros, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error al leer el archivo CSV: %v", err)
	}

	// Verificar si hay datos (sin contar encabezado)
	total := len(registros)
	if total <= 1 {
		fmt.Println("No hay suficientes registros para mostrar.")
		return nil
	}
	if n == -1 || n > total-1 {
		n = total - 1
	}

	// Determinar cuántos registros mostrar
	inicio := max(total-n, 1)

	// Imprimir los últimos `n` registros
	fmt.Println("Últimos registros:")
	for i, fila := range registros[inicio:] {
		fmt.Printf("Semana %d: %v\n", i+1, fila)
	}

	return nil
}

func EstadisticasTotales(filename string, n int) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	registros, err := reader.ReadAll()
	if err != nil {
		return
	}

	// Verificar si hay datos (sin contar encabezado)
	total := len(registros)
	if total <= 1 {
		fmt.Println("No hay suficientes registros para mostrar.")
		return
	}
	if n == -1 || n > total-1 {
		n = total - 1
	}

	// Determinar cuántos registros mostrar
	inicio := max(total-n, 1)
	var suma float32
	var cantidad int // Para contar la cantidad de registros
	ultimoRegistro := registros[total-1]
	pesoFinal, err := strconv.ParseFloat(ultimoRegistro[1], 32)

	for i, fila := range registros[inicio:] {
		if i == 0 { // Evita división por cero
			continue
		}

		peso, err := strconv.ParseFloat(fila[1], 32)
		if err != nil {
			fmt.Println("Error al convertir peso:", err)
			continue
		}

		suma += float32(peso)
		cantidad++ // Incrementa la cantidad de registros procesados
	}

	// Si hay registros procesados
	if cantidad > 0 {
		//promedioPeso := suma / float32(cantidad)
		// Porcentaje de pérdida de peso
		porcentajeBajado := ((float32(100) - float32(pesoFinal)) / float32(100)) * 100

		fmt.Printf("Ha bajado %.2f%% de su peso inicial.\n", porcentajeBajado)
	}
}
