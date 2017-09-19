package main

import (
	"Assignment-3/dao"
	"Assignment-3/dbclient"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/satori/go.uuid"

	"github.com/abiosoft/ishell"
)

func addListItemsToShell(shell *ishell.Shell) {
	listItemsCmd := &ishell.Cmd{
		Name:     "list",
		Help:     "list available items",
		LongHelp: `List available items by category.`,
	}

	for _, category := range dbclient.GetInventoryDBClient().GetCategoryDescriptions() {
		listItemsCmd.AddCmd(&ishell.Cmd{
			Name: category[0],
			Help: category[1],
			Func: func(c *ishell.Context) {
				items := dbclient.GetInventoryDBClient().GetItemsByCategory(c.Cmd.Name)
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
					dbclient.GetInventoryDBClient().Reserve(item, uint64(quantityDesired))
					GetUserManager().user.PersonalCart.AddItem(item, uint64(quantityDesired))
					dbclient.GetUserDBClient().SetCart(GetUserManager().user)
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

func addRemoveItemFromCartToShell(shell *ishell.Shell) {
	removeItemCmd := &ishell.Cmd{
		Name: "remove",
		Help: "remove items from the cart",
		Func: func(c *ishell.Context) {
			var itemsToRemove []dao.CartItem

			items := GetUserManager().user.PersonalCart.Items
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
					c.Printf("You selected %s\nDescription:\n\t%s\n\nYou have %d in your cart. How many would you like to remove?\n",
						item.Item.Name, item.Item.Description, item.Quantity)

					quantityToRemove, err := strconv.Atoi(c.ReadLine())
					for {
						if err == nil && quantityToRemove > 0 && uint64(quantityToRemove) <= item.Quantity {
							dbclient.GetInventoryDBClient().Release(item.Item, uint64(quantityToRemove))
							GetUserManager().user.PersonalCart.RemoveItem(item.Item, uint64(quantityToRemove))
							dbclient.GetUserDBClient().SetCart(GetUserManager().user)
							break
						} else {
							c.Printf("Invalid quantity. Choose a number between 1 and %d\n", item.Quantity)
							quantityToRemove, err = strconv.Atoi(c.ReadLine())
						}
					}
				}
			}
		},
	}

	shell.AddCmd(removeItemCmd)
}

func addDisplayCartToShell(shell *ishell.Shell) {
	displayCartCmd := &ishell.Cmd{
		Name: "show",
		Help: "show items in the cart",
		Func: func(c *ishell.Context) {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)
			fmt.Fprint(w, displayCart(GetUserManager().user.PersonalCart))
			w.Flush()
		},
	}

	shell.AddCmd(displayCartCmd)
}

func addCheckoutToShell(shell *ishell.Shell) {
	checkoutCmd := &ishell.Cmd{
		Name: "checkout",
		Help: "Proceed to checkout your stuffs",
		Func: func(c *ishell.Context) {
			user := GetUserManager().user
			if len(user.PersonalCart.Items) > 0 {
				// preview cart
				buf := bytes.NewBufferString("")
				w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', tabwriter.AlignRight)
				fmt.Fprint(w, "Cart:\t\t\t\n")
				fmt.Fprint(w, displayCart(user.PersonalCart))
				fmt.Fprint(w, "\t\t\t\n")
				// show address and osc number
				fmt.Fprintf(w, "OSC Card:\t\t%d\t\n", user.OscCardNumber)
				fmt.Fprintf(w, "Address:\t\t%s\t\n", user.Address)
				fmt.Fprint(w, "\t\t\t\n")
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
					dbclient.GetUserDBClient().ChangeAddress(user, newAddress)
				}
				c.ReadLine()
				// confirm order
				buf = bytes.NewBufferString("")
				w = tabwriter.NewWriter(buf, 0, 0, 3, ' ', tabwriter.AlignRight)
				fmt.Fprint(w, "Cart:\t\t\t\n")
				fmt.Fprint(w, displayCart(user.PersonalCart))
				fmt.Fprint(w, "\t\t\t\n")
				// show address and osc number
				fmt.Fprintf(w, "OSC Card:\t\t%d\t\n", user.OscCardNumber)
				fmt.Fprintf(w, "Address:\t\t%s\t\n", user.Address)
				fmt.Fprint(w, "\t\t\t\n")
				fmt.Fprint(w, "Confirm Purchase?")
				w.Flush()
				confirmPurchase := c.MultiChoice([]string{"No", "Yes"}, buf.String())
				if confirmPurchase == 1 {
					// remove items from inventory and add Purchase to purchase history
					purchase := dao.Purchase{uuid.NewV4(), time.Now(), user.Username, user.Address, user.OscCardNumber, user.PersonalCart.GetTotalCost(), *user.PersonalCart}

					for _, cartItem := range user.PersonalCart.Items {
						dbclient.GetInventoryDBClient().Remove(cartItem.Item, cartItem.Quantity)
					}
					dbclient.GetPurchaseDBClient().AddPurchase(purchase)
					user.PersonalCart.EmptyCart()
				}
			} else {
				c.Println("Cart is currently empty!")
			}
		},
	}

	shell.AddCmd(checkoutCmd)
}

func addPurchaseHistoryToShell(shell *ishell.Shell) {
	purchaseHistoryCmd := &ishell.Cmd{
		Name: "purchases",
		Help: "View purchase history",
		Func: func(c *ishell.Context) {
			purchases := dbclient.GetPurchaseDBClient().GetPurchasesByUsername(GetUserManager().user.Username)
			buf := bytes.NewBufferString("")
			w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', tabwriter.AlignRight)

			for p := 0; p < len(purchases); p++ {
				fmt.Fprintf(w, "Purchase:\t%s\n", purchases[p].Id.String())
				fmt.Fprintf(w, "Address:\t%s\n", purchases[p].Address)
				fmt.Fprintf(w, "Date:\t%s\n", purchases[p].CheckoutDate.String())
				fmt.Fprintf(w, "OSC Card:\t%d\n", purchases[p].OscCardNumber)
				fmt.Fprint(w, "Cart:\n")
				fmt.Fprintf(w, displayCart(&purchases[p].Cart))
			}
			c.ShowPaged(buf.String())
		},
	}

	shell.AddCmd(purchaseHistoryCmd)
}

func addWhoamiToShell(shell *ishell.Shell) {
	whoamiCmd := &ishell.Cmd{
		Name: "whoami",
		Help: "show user information",
		Func: func(c *ishell.Context) {
			c.Printf("Username: %s\n", GetUserManager().user.Username)
			c.Printf("Card Number: %d\n", GetUserManager().user.OscCardNumber)
			c.Printf("Address: %s\n", GetUserManager().user.Address)
			c.Printf("Cart: %s\n", GetUserManager().user.PersonalCart)
		},
	}

	shell.AddCmd(whoamiCmd)
}

func displayCart(cart *dao.Cart) string {
	display := fmt.Sprintln("Item\tQuantity\tCost Each\t")
	for _, cartItem := range cart.Items {
		display += fmt.Sprintf(" %s\t%d\t%.2f\t\n", cartItem.Item.Name, cartItem.Quantity, cartItem.Item.Price)
	}
	display += fmt.Sprintln("\t\t\t")
	display += fmt.Sprintf(" \tTotal Cost:\t%.2f\t\n", cart.GetTotalCost())
	return display
}
