package main

import (
	"Assignment-3/dbclient"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/abiosoft/ishell"
)

func addListItemsToShell(shell *ishell.Shell) {
	listItemsCmd := &ishell.Cmd{
		Name:     "list",
		Help:     "list available items",
		LongHelp: `List available items by category.`,
	}

	for _, category := range dbclient.GetInventoryDBClient().GetAvailableCategories() {
		listItemsCmd.AddCmd(&ishell.Cmd{
			Name: category,
			Help: category,
			Func: func(c *ishell.Context) {
				items := dbclient.GetInventoryDBClient().GetItemsByCategory(c.Cmd.Help)
				for _, item := range items {
					c.Printf("%.2f\t%s\n", item.Price, item.Name)
				}
			},
		})
	}

	shell.AddCmd(listItemsCmd)
}

func addAddItemToCartToShell(shell *ishell.Shell) {
	addItemCmd := &ishell.Cmd{
		Name: "add",
		Help: "add item to cart",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			categories := dbclient.GetInventoryDBClient().GetAvailableCategories()
			categoryIdx := c.MultiChoice(categories, "Which category do you want to add an item from?")
			category := categories[categoryIdx]

			items := dbclient.GetInventoryDBClient().GetItemsByCategory(category)
			var itemTexts []string
			for _, item := range items {
				itemTexts = append(itemTexts, item.Name)
			}
			itemIdx := c.MultiChoice(itemTexts, fmt.Sprintf("Which item do you want to add from the %s category?", category))
			item := items[itemIdx]

			c.Printf("You selected %s\nDescription:\n\t%s\n\nThere are %d available. How many would you like to add to your cart?\n",
				item.Name, item.Description, item.QuantityAvailable)

			quantityDesired, err := strconv.Atoi(c.ReadLine())
			for {
				if err == nil && quantityDesired > 0 && uint64(quantityDesired) <= item.QuantityAvailable {
					GetUserManager().user.PersonalCart.AddItem(item, uint64(quantityDesired))
					break
				} else {
					c.Printf("Invalid quantity. Choose a number between 1 and %d\n", item.QuantityAvailable)
					quantityDesired, err = strconv.Atoi(c.ReadLine())
				}
			}

			c.Printf("%d %s added to cart\n", quantityDesired, item.Name)
		},
	}

	shell.AddCmd(addItemCmd)
}

func addDisplayCartToShell(shell *ishell.Shell) {
	displayCartCmd := &ishell.Cmd{
		Name: "show",
		Help: "show items in the cart",
		Func: func(c *ishell.Context) {

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)

			fmt.Fprintln(w, "Item\tQuantity\tCost Each\t")
			for _, cartItem := range GetUserManager().user.PersonalCart.Items {
				fmt.Fprintf(w, " %s\t%d\t%.2f\t\n", cartItem.Item.Name, cartItem.Quantity, cartItem.Item.Price)
			}
			fmt.Fprintln(w, "\t\t\t")
			fmt.Fprintf(w, " \tTotal Cost:\t%.2f\t\n", GetUserManager().user.PersonalCart.GetTotalCost())
			w.Flush()
		},
	}

	shell.AddCmd(displayCartCmd)
}

func addWhoamiToShell(shell *ishell.Shell) {
	whoamiCmd := &ishell.Cmd{
		Name: "whoami",
		Help: "show user information",
		Func: func(c *ishell.Context) {
			c.Printf("Username: %s\n", GetUserManager().user.Username)
			c.Printf("Card Number: %s\n", GetUserManager().user.OscCardNumber)
			c.Printf("Address: %s\n", GetUserManager().user.Address)
			c.Printf("Cart: %s\n", GetUserManager().user.PersonalCart)
		},
	}

	shell.AddCmd(whoamiCmd)
}
