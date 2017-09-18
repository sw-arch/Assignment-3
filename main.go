package main

import (
	// "Assignment-3/dbclient"
	"github.com/abiosoft/ishell"
)

func main() {
	shell := ishell.New()

	shell.Println("Welcome to the Store")

	addLoginToShell(shell)
	addRegisterToShell(shell)

	shell.Run()
}
