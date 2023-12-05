package main

import (
	"time"
)

type BCP struct {
	ProcessID      string
	State          string
	ProgramCounter int
	CoreID         int
	// Otras propiedades del BCP
}

func simulateProcess(processID string, instructions []string, coreID int, dispatcher chan<- BCP, done chan<- bool) {
	for i, instruction := range instructions {
		time.Sleep(time.Millisecond * 500)

		dispatcher <- BCP{
			ProcessID:      processID,
			State:          "Ejecutando",
			ProgramCounter: i + 1,
			CoreID:         coreID,
			// Otras propiedades del BCP
		}

		if i == len(instructions)-1 {
			done <- true
		}
	}
}

func main() {
	// Leer el número de núcleos desde la línea de comandos

	// ENTRADA NUMERO ENTERO

	// Resto del código...

	// Canales adicionales para la comunicación entre el dispatcher y las goroutines
	processesQueue := make(chan string, 10) // Cola de procesos disponibles
	finishedProcesses := make(chan bool)    // Canal para notificar al dispatcher cuando se hayan terminado todos los procesos

	// Función del Dispatcher
	dispatcher := func() {
		// Implementa la lógica del dispatcher
		// ...
	}

	// Lanzar el dispatcher en una goroutine
	go dispatcher()

	// Goroutine para agregar procesos a la cola de procesos disponibles
	go func() {
		for processID := range processesQueue {
			readyQueue <- BCP{
				ProcessID: processID,
				State:     "Listo",
			}
		}
		close(readyQueue)
	}()

	// Resto del código...
}
