package main

import (
	"Assignment-3/dbclient"
	"fmt"
	"strconv"

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
					c.Println(item.Name)
				}
			},
		})
	}

	shell.AddCmd(listItemsCmd)
}

func addAddItemToCartToShell(shell *ishell.Shell) {
	addItemCmd := &ishell.Cmd{
		Name:     "add",
		Help:     "add item to cart",
		LongHelp: `Interactive cart selection.`,
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

			c.Println(fmt.Sprintf("You selected %s\nDescription:\n\t%s\n\nThere are %d available. How many would you like to add to your cart?",
				item.Name, item.Description, item.QuantityAvailable))

			quantityDesired, err := strconv.Atoi(c.ReadLine())
			for {
				if err == nil && quantityDesired > 0 && uint64(quantityDesired) <= item.QuantityAvailable {
					userManagerInstance.user.Cart.AddItem(item, uint64(quantityDesired))
					break
				} else {
					c.Println(fmt.Sprintf("Invalid quantity. Choose a number between 1 and %d", item.QuantityAvailable))
					quantityDesired, err = strconv.Atoi(c.ReadLine())
				}
			}

			c.Println(fmt.Sprintf("%d %s added to cart", quantityDesired, item.Name))
		},
	}

	shell.AddCmd(addItemCmd)
}
