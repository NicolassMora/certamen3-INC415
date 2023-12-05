package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

// Definición de una estructura para el BCP1 (Bloque de Control de Proceso)
type BCP1 struct {
	ProcesoID      string
	Estado         string
	ProgramCounter int
}

// Función que simula la ejecución de un proceso
func Proceso(procesoID string, instrucciones []string, dispatcher chan<- BCP1, done chan<- bool) {
	// Implementa la lógica de ejecución del proceso y notifica al dispatcher cuando sea necesario
	for i, instruction := range instrucciones {
		// Simula la ejecución de una instrucción
		time.Sleep(time.Millisecond * 500)

		// Notifica al dispatcher sobre el cambio de estado o eventos relevantes
		dispatcher <- BCP1{
			ProcesoID:      procesoID,
			Estado:         "Ejecutando",
			ProgramCounter: i + 1,
			// Otras propiedades del BCP1
		}

		// Simula eventos de E/S u otros cambios de estado
		// ...

		// Finaliza el proceso si es la última instrucción
		if i == len(instrucciones)-1 {
			done <- true
		}
	}
}

func main() {

	var err error

	// Configuración del simulador
	maxInstruccion := 100 // Máximo de instrucciones a ejecutar por la CPU antes de cambiar de proceso
	Probabilidad := 4     // Probabilidad P de que una instrucción finalice antes de tiempo (1/P)

	// Configuración de las colas de procesos (Listo y Bloqueado)
	readyQueue := make(chan BCP1, 10)
	blockedQueue := make(chan BCP1, 10)

	// Canal para notificar al dispatcher sobre eventos
	dispatcherChan := make(chan BCP1)

	// Canal para notificar al dispatcher sobre la finalización de un proceso
	procesoHecho := make(chan bool)

	// Espera grupal para esperar que todas las goroutines finalicen
	var wg sync.WaitGroup

	// Función del Dispatcher
	dispatcher := func() {
		// Implementa la lógica del dispatcher
		// ...
	}

	// Lanzar el dispatcher en una goroutine
	go dispatcher()

	// Archivo de entrada simulado
	fileProceso, err := os.Open("proceso.txt")

	scanner := bufio.NewScanner(fileProceso)

	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}

	procesos := map[string][]string{
		"Proceso1": {"I", "I", "ES 7", "I", "I", "F"},
		"Proceso2": {"I", "I", "ES 5", "I", "F"},
		// Otros procesos...
	}

	// Inicialización de los procesos
	for procesoID, instrucciones := range procesos {
		wg.Add(1)
		go func(id string, instr []string) {
			defer wg.Done()
			Proceso(id, instr, dispatcherChan, procesoHecho)
		}(procesoID, instrucciones)
	}

	// Esperar a que todas las goroutines de procesos finalicen
	go func() {
		wg.Wait()
		close(dispatcherChan)
		close(procesoHecho)
	}()

	// Esperar la señal de finalización de un proceso
	<-procesoHecho

	// Esperar que el dispatcher también finalice
	<-dispatcherChan
}
