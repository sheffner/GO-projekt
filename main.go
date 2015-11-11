// Projekt project main.go
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Line struct {
	name         string
	PID          int64
	session_name string
	SID          int64
	mem          int64
}

func main() {
	//openProgramWithOneArgument("C:/Program Files (x86)/Mozilla Firefox/firefox.exe", "www.google.de")
	//openProgramWithoutArgument("C:/Program Files/Microsoft Games/Solitaire/solitaire.exe")
	ExitHardWithName("firefox.exe")
}

// diese Funktion öffnet ein Programm mit einem Argument(z.B. Seite die geöffnet werden soll). Es wird überprüft ob der Pfad vorhanden ist und das Programm gestartet werden kann
func openProgramWithOneArgument(path string, argument string) {
	cmd := exec.Command(path, argument)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
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
	log.Printf("Command finished with error: %v", err)
}

// diese Funktion listet alle geöffneten Prozesse mit ihrer PID
func listProcess() {
	cmd := exec.Command("tasklist")
	out, _ := cmd.Output()
	processes := strings.Split(string(out[:len(out)-1]), "\n")[3:]

	lines := []Line{}
	for i := 0; i < len(processes); i++ {
		temp := strings.Fields(processes[i])
		name := strings.Join(temp[:len(temp)-5], " ")
		pid, _ := strconv.ParseInt(temp[len(temp)-5], 10, 32)
		s_name := temp[len(temp)-4]
		sid, _ := strconv.ParseInt(temp[len(temp)-3], 10, 32)
		mem, _ := strconv.ParseInt(strings.Replace(temp[len(temp)-2], ".", "", -1), 10, 32)
		lines = append(lines, Line{name, pid, s_name, sid, mem})
	}
	fmt.Println(lines)
}

// diese Funktion lässt Prozesse anhand ihrer ID mit kill beenden
func ExitHardWithID(pid int) {
	killCmd := exec.Command("taskkill", "/pid", strconv.Itoa(pid))
	//killCmd := exec.Command("taskkill", "/pid", strconv.Itoa(pid))
	err := killCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Command finished with error: %v", err)
}

// diese Funktion lässt Prozesse anahnd ihres Namens mit kill beenden
func ExitHardWithName(name string) {
	killCmd := exec.Command("taskkill", "/IM", name, "/F")
	err := killCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	log.Printf("Command finished with error: %v", err)
}
