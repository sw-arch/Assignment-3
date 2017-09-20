package ui

import (
	"Assignment-3/dao"
	"Assignment-3/manager"
	"bytes"
	"fmt"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/abiosoft/ishell"
)

func AddListItemsToShell(shell *ishell.Shell) {
	listItemsCmd := &ishell.Cmd{
		Name:     "list",
		Help:     "list available items",
		LongHelp: `List available items by category.`,
	}

	for _, category := range manager.GetStore().GetCategories() {
		listItemsCmd.AddCmd(&ishell.Cmd{
			Name: category.Identifier,
			Help: category.Description,
			Func: func(c *ishell.Context) {
				for _, item := range manager.GetStore().GetItemsInCategory(c.Cmd.Name) {
					c.Printf("%.2f\t%s\n", item.Price, item.Name)
				}
			},
		})
	}

	shell.AddCmd(listItemsCmd)
}

func AddAddItemToCartToShell(shell *ishell.Shell) {
	addItemCmd := &ishell.Cmd{
		Name: "add",
		Help: "add item to cart",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			categories := manager.GetStore().GetCategories()
			var categoryDescriptions []string
			for _, category := range categories {
				categoryDescriptions = append(categoryDescriptions, category.Description)
			}
			categoryIdx := c.MultiChoice(categoryDescriptions, "Which category do you want to add an item from?")
			category := categories[categoryIdx]

			items := manager.GetStore().GetItemsInCategory(categories[categoryIdx].Identifier)
			var itemTexts []string
			for _, item := range items {
				itemTexts = append(itemTexts, fmt.Sprintf("%s - $%.2f", item.Name, item.Price))
			}
			itemIdx := c.MultiChoice(itemTexts, fmt.Sprintf("Which item do you want to add from the %s category?", category))
			item := items[itemIdx]

			attributeID1, attributeID2 := manager.GetStore().GetAttributesByCategory(category.Identifier)
			attribute1 := fmt.Sprintf("%s: %s", attributeID1, item.AttributeOne)
			attribute2 := fmt.Sprintf("%s: %s", attributeID2, item.AttributeTwo)

			c.Printf("You selected %s\nDescription:\n\t%s\n\t%s\n\t%s\nThere are %d available. How many would you like to add to your cart?\n",
				item.Name, item.Description, attribute1, attribute2, item.QuantityAvailable)

			quantityDesired, err := strconv.Atoi(c.ReadLine())
			for {
				if err == nil && quantityDesired > 0 && int64(quantityDesired) <= item.QuantityAvailable {
					manager.GetStore().AddToCart(item, quantityDesired)
					break
				} else if quantityDesired == 0 {
					break
				} else {
					c.Printf("Invalid quantity. Choose a number between 0 and %d\n", item.QuantityAvailable)
					quantityDesired, err = strconv.Atoi(c.ReadLine())
				}
			}

			c.Printf("%d %s added to cart\n", quantityDesired, item.Name)
		},
	}

	shell.AddCmd(addItemCmd)
}

func AddRemoveItemFromCartToShell(shell *ishell.Shell) {
	// TODO: Remove isn't working
	removeItemCmd := &ishell.Cmd{
		Name: "remove",
		Help: "remove items from the cart",
		Func: func(c *ishell.Context) {
			user := manager.GetUserManager().User
			if len(user.PersonalCart.Items) == 0 {
				c.Println("Cart is currently empty!")
				return
			}
			var itemsToRemove []dao.CartItem

			items := user.PersonalCart.Items
			var itemTexts []string
			for _, item := range items {
				itemTexts = append(itemTexts, item.Item.Name)
			}

			itemIdxs := c.Checklist(itemTexts,
				"Which items do you want to remove?", nil)
			for itemIdx := range itemIdxs {
				item := items[itemIdx]
				if item.Quantity == 1 {
					itemsToRemove = append(itemsToRemove, item)
				} else {
					attributeID1, attributeID2 := manager.GetStore().GetAttributesByCategory(item.Item.Category)
					attribute1 := fmt.Sprintf("%s: %s", attributeID1, item.Item.AttributeOne)
					attribute2 := fmt.Sprintf("%s: %s", attributeID2, item.Item.AttributeTwo)

					c.Printf("You selected %s\nDescription:\n\t%s\n\t%s\n\t%s\nThere are %d in your cart. How many would you like to remove?\n",
						item.Item.Name, item.Item.Description, attribute1, attribute2, item.Quantity)

					quantityToRemove, err := strconv.Atoi(c.ReadLine())
					for {
						if err == nil && quantityToRemove > 0 && int64(quantityToRemove) <= item.Quantity {
							manager.GetStore().RemoveFromCart(item.Item, quantityToRemove)
							break
						} else if quantityToRemove == 0 {
							break
						} else {
							c.Printf("Invalid quantity. Choose a number between 0 and %d\n", item.Quantity)
							quantityToRemove, err = strconv.Atoi(c.ReadLine())
						}
					}
				}
			}
		},
	}

	shell.AddCmd(removeItemCmd)
}

func AddDisplayCartToShell(shell *ishell.Shell) {
	displayCartCmd := &ishell.Cmd{
		Name: "show",
		Help: "show items in the cart",
		Func: func(c *ishell.Context) {
			c.Print(displayCart(manager.GetUserManager().User.PersonalCart))
		},
	}

	shell.AddCmd(displayCartCmd)
}

func AddCheckoutToShell(shell *ishell.Shell) {
	checkoutCmd := &ishell.Cmd{
		Name: "checkout",
		Help: "Proceed to checkout your stuffs",
		Func: func(c *ishell.Context) {
			user := manager.GetUserManager().User
			if len(user.PersonalCart.Items) == 0 {
				c.Println("Cart is currently empty!")
				return
			}
			// preview cart
			buf := bytes.NewBufferString("")
			w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', tabwriter.AlignRight)
			fmt.Fprintln(w, "Cart:\t\t")
			w.Flush()
			fmt.Fprint(buf, displayCart(user.PersonalCart))
			fmt.Fprintln(w, "\t\t")
			// show address and osc number
			fmt.Fprintf(w, "OSC Card:\t\t%d\n", user.OscCardNumber)
			fmt.Fprintf(w, "Address:\t\t%s\n", user.Address)
			fmt.Fprintln(w, "\t\t")
			fmt.Fprint(w, "Would you like to change your address?")
			w.Flush()
			// offer to change address
			changeAddress := c.MultiChoice([]string{"No", "Yes"}, buf.String())
			if changeAddress == 1 {
			enterAddress:
				c.Print("Enter New Address: ")
				newAddress := c.ReadLine()
				c.Printf("Address:\t%s", newAddress)
				correctAddress := c.MultiChoice([]string{"Yes", "No"}, "Is this address correct?")
				if correctAddress == 1 {
					goto enterAddress
				}
				user.Address = newAddress
				manager.GetUserManager().ChangeAddress(newAddress)
			}
			// confirm order
			buf = bytes.NewBufferString("")
			w = tabwriter.NewWriter(buf, 0, 0, 3, ' ', tabwriter.AlignRight)
			fmt.Fprint(w, "Order:\t\t\n")
			w.Flush()
			fmt.Fprint(buf, displayCart(user.PersonalCart))
			fmt.Fprintln(w, "\t\t")
			// show address and osc number
			fmt.Fprintf(w, "OSC Card:\t\t%d\n", user.OscCardNumber)
			fmt.Fprintf(w, "Address:\t\t%s\n", user.Address)
			fmt.Fprintln(w, "\t\t")
			fmt.Fprint(w, "Confirm Purchase?")
			w.Flush()
			confirmPurchase := c.MultiChoice([]string{"No", "Yes"}, buf.String())
			if confirmPurchase == 1 {
				// remove items from inventory and add Purchase to purchase history
				purchase := manager.GetCashier().CreatePurchase(user)
				manager.GetCashier().ConfirmPurchase(user, &purchase)
			}

			// This print ensures the command completes. Something isn't flushing right.
			c.Println()
		},
	}

	shell.AddCmd(checkoutCmd)
}

func AddPurchaseHistoryToShell(shell *ishell.Shell) {
	purchaseHistoryCmd := &ishell.Cmd{
		Name: "purchases",
		Help: "View purchase history",
		Func: func(c *ishell.Context) {
			purchases := manager.GetUserManager().GetHistory()

			if len(purchases) == 0 {
				c.Println("You have not made any purchases")
				return
			}

			buf := bytes.NewBufferString("")

			for _, p := range purchases {
				w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', tabwriter.AlignRight)
				fmt.Fprintf(w, "Purchase:\t\t%s\n", p.Id.String())
				fmt.Fprintf(w, "Address:\t\t%s\n", p.Address)
				fmt.Fprintf(w, "Date:\t\t%s\n", p.CheckoutDate.Format(time.RFC1123))
				fmt.Fprintf(w, "OSC Card:\t\t%d\n", p.OscCardNumber)
				w.Flush()
				fmt.Fprintln(buf)
				fmt.Fprint(buf, displayCart(p.Cart))
				fmt.Fprintln(buf)
			}

			c.ShowPaged(buf.String())
		},
	}

	shell.AddCmd(purchaseHistoryCmd)
}

func AddWhoamiToShell(shell *ishell.Shell) {
	whoamiCmd := &ishell.Cmd{
		Name: "whoami",
		Help: "show user information",
		Func: func(c *ishell.Context) {
			user := manager.GetUserManager().User
			c.Printf("Username: %s\n", user.Username)
			c.Printf("Card Number: %d\n", user.OscCardNumber)
			c.Printf("Address: %s\n", user.Address)
			c.Printf("Cart: %s\n", user.PersonalCart)
		},
	}

	shell.AddCmd(whoamiCmd)
}

func AddLogoutToShell(shell *ishell.Shell) {
	logoutCmd := &ishell.Cmd{
		Name: "logout",
		Help: "logout of user session",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			manager.GetUserManager().LogOut()
			c.Stop()
		},
	}

	shell.AddCmd(logoutCmd)
}

func displayCart(cart *dao.Cart) string {
	buf := bytes.NewBufferString("")
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "Item\tQuantity\tCost Each\t")
	for _, cartItem := range cart.Items {
		attributeID1, attributeID2 := manager.GetStore().GetAttributesByCategory(cartItem.Item.Category)
		attribute1 := fmt.Sprintf("%s: %s", attributeID1, cartItem.Item.AttributeOne)
		attribute2 := fmt.Sprintf("%s: %s", attributeID2, cartItem.Item.AttributeTwo)

		fmt.Fprintf(w, "%s\t%d\t%.2f\t\n", cartItem.Item.Name, cartItem.Quantity, cartItem.Item.Price)
		fmt.Fprintf(w, "%s   %s\t\t\t\n", attribute1, attribute2)
	}
	fmt.Fprintln(w, "\t\t\t")
	fmt.Fprintf(w, "\tTotal Cost:\t%.2f\t\n", cart.GetTotalCost())
	w.Flush()
	return buf.String()
}
