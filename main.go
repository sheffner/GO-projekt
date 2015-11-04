// Projekt project main.go
package main

import (
	"log"
	"os/exec"
)

// diese Funktion öffnet ein Programm mit einem Argument(z.B. Seite die geöffnet werden soll). Es wird überprüft ob der Pfad vorhanden ist und das Programm gestartet werden kann
func openProgramWithOneArgument(path string, argument string) {
	cmd := exec.Command(path, argument)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

// diese Funktion öffnet ein Programm. Es wird überprüft ob der Pfad vorhanden ist und das Programm gestartet werden kann
func openProgramWithoutArgument(path string) {
	cmd := exec.Command(path)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
func main() {
	openProgramWithOneArgument("C:/Program Files (x86)/Mozilla Firefox/firefox.exe", "www.google.de")
	openProgramWithoutArgument("C:/Program Files/Microsoft Games/Solitaire/solitaire.exe")
}
