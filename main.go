package main

import (
	// "Assignment-3/dbclient"
	"github.com/abiosoft/ishell"
)

func checkCredintials(username string, password string) bool {
	return true
}

func main() {
	shell := ishell.New()

	shell.Println("Welcome to the Store")

	// Login cmd
	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "Authentification for login",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Print("Username: ")
			username := c.ReadLine()

			c.Print("Password: ")
			password := c.ReadPassword()

			// do something with username and password
			if checkCredintials(username, password) {
				c.Println("Authentication Successful.")
				// add commands to shell
			} else {
				c.Println("Authentication Failed.")
			}
		},
	})

	shell.Run()
}
