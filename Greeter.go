package main

import (
	"github.com/abiosoft/ishell"
)

func addLoginToShell(shell *ishell.Shell) {
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

			success := GetUserManager().logIn(username, password)
			if success {
				c.Println("Authentication Successful.")
			} else {
				c.Println("Authentication Failed.")
			}
		},
	})
}

func addRegisterToShell(shell *ishell.Shell) {
	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Register a new user",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Print("Username: ")
			username := c.ReadLine()

			var password1, password2 string
			password1 = "Not equal to password2"

			for password1 != password2 {
				c.Print("Password: ")
				password1 = c.ReadPassword()

				c.Print("Repeat Password: ")
				password2 = c.ReadPassword()
			}

			c.Print("Address: ")
			address := c.ReadLine()

			success := GetUserManager().register(username, password1, address)
			if success {
				c.Println("New User created.")
			} else {
				c.Println("New User failed to create.")
			}
		},
	})
}
