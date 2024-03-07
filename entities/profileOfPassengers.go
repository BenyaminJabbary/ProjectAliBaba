package entities

type DetailsOfDomesticPassengers struct {
	FirstName    string
	LastName     string
	Gender       string
	NationalCode string
	DateOfBirth  string
	NumberPhone string
}

type ProfileOfForeignTravelers struct {
	DetailsOfDomesticPassengers
	Nationality string
	PassportNumber string
	ExpiryDateOfPassport string
}