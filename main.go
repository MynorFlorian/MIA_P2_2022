package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Gola")
}

func terminal() {
	for true {
		fmt.Print("~$ ")
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')
		eleccion := strings.TrimRight(entrada, "\r\n")

		if eleccion == "" {
			return
		} else if strings.Contains(eleccion, "\\*") {
			eleccionEnvio := strings.Replace(eleccion, "\\*", " ", -1)
			terminalSalto(eleccionEnvio)
		} else if eleccion[0] == 35 {
			//#Comentario
		} else if eleccion == "x" || eleccion == "X" {
			return
		} else {
			//Comandos en una sola linea
			analizador(eleccion)
		}
	}
}

func terminalSalto(eleccionAnterior string) {
	fmt.Print(">>")
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
	eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario

	eleccionAnterior += eleccion

	if strings.Contains(eleccion, "/*") {
		eleccionEnvio := strings.Replace(eleccionAnterior, "/*", " ", -1)
		terminalSalto(eleccionEnvio)
	} else if eleccion[0] == 35 {

	} else {
		fmt.Println(eleccionAnterior)
		analizador(eleccionAnterior)
	}

}

func salida() {
	return
}

func analizador(comando string) {
	if comando == "x" || comando == "X" {
		salida()
	} else if comando == "" {

	} else {
		if comando != "" {

			arregloComandos := strings.Split(comando, " ")
			ejecucionComando(arregloComandos)
		}

	}
}

func ejecucionComando(comandos []string) {

}
