package main

import (
	"testing"
)

// Test zur Ausführung von Firefox mit der Startseite google.de
func TestOpenFirefox(t *testing.T) {
	openProgramWithOneArgument("C:/Program Files (x86)/Mozilla Firefox/firefox.exe", "www.google.de")
}

// Test zur Ausführung von Solitaire
func TestOpenSolitaire(t *testing.T) {
	openProgramWithoutArgument("C:/Program Files/Microsoft Games/Solitaire/solitaire.exe")
}

// Test zum Beenden von Firefox
func TestCloseFirefox(t *testing.T) {
	TestOpenFirefox(t)
	ExitHardWithName("firefox.exe")
}
