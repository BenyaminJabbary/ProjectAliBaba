package tourServices

import "fmt"

func ConfirmationOfReservation(numberOfAdultsInDoubleRoom, numberOfAdultsInSingleRoom, numberOfChildren, numberOfBaby int) {
	var confirmationOfReservation string
	fmt.Print("Confirmation of reservation(yes or no): ")
	fmt.Scanln(&confirmationOfReservation)
	if confirmationOfReservation == "yes" {
		fmt.Println("Confirmed:")
		NumberOfPassengers(numberOfAdultsInDoubleRoom, numberOfAdultsInSingleRoom, numberOfChildren, numberOfBaby)
		PurchaseConfirmation()
	} else if confirmationOfReservation == "no" {
		fmt.Println("Not confirmed!")
	}
}

func PurchaseConfirmation() {
	var purchaseConfirmation = ""
	fmt.Print("Purchase confirmation(yes or no): ")
	fmt.Scanln(&purchaseConfirmation)
	if purchaseConfirmation == "yes" {
		fmt.Println("The purchase was made!")
	} else if purchaseConfirmation == "no" {
		fmt.Println("The purchase was not made!")
	}
}
