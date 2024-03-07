package tourServices

import (
	"fmt"
	"go-aliBaba/entities"
)

 var passengers []entities.ProfileOfForeignTravelers

func GetPassengerInfo() {
	getInfo := entities.ProfileOfForeignTravelers{}
	fmt.Print("FirstName: ")
	fmt.Scanln(&getInfo.FirstName)

	fmt.Print("LastName: ")
	fmt.Scanln(&getInfo.LastName)

	fmt.Print("Gender: ")
	fmt.Scanln(&getInfo.Gender)

	fmt.Print("NationalCode: ")
	fmt.Scanln(&getInfo.NationalCode)

	fmt.Print("DateOfBirth: ")
	fmt.Scanln(&getInfo.DateOfBirth)
	
	fmt.Print("NumberPhone: ")
	fmt.Scanln(&getInfo.NumberPhone)

	fmt.Print("Nationality: ")
	fmt.Scanln(&getInfo.Nationality)

	fmt.Print("PassportNumber: ")
	fmt.Scanln(&getInfo.PassportNumber)

	fmt.Print("ExpiryDateOfPassport: ")
	fmt.Scanln(&getInfo.ExpiryDateOfPassport)

	passengers = append(passengers, getInfo)
}

func NumberOfPassengers(numberOfAdultsInDoubleRoom, numberOfAdultsInSingleRoom, numberOfChildren, numberOfBaby int) {
	if numberOfAdultsInDoubleRoom != 0 {
		for i := 0; i < numberOfAdultsInDoubleRoom; i++ {
			fmt.Println("Adults in double room: ")
			GetPassengerInfo()
			
		}
	}

	if numberOfAdultsInSingleRoom != 0 {
		for i := 0; i < numberOfAdultsInSingleRoom; i++ {
			fmt.Println("Adults in single room: ")
			GetPassengerInfo()
		}
	}

	if numberOfChildren != 0 {
		for i := 0; i < numberOfChildren; i++ {
			fmt.Println("Children: ")
			GetPassengerInfo()
		}
	}

	if numberOfBaby != 0 {
		for i := 0; i < numberOfBaby; i++ {
			fmt.Println("Baby: ")
			GetPassengerInfo()
		}
	}
	
}