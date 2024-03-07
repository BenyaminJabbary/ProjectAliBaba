package entities

type TravelTour struct {
	Id int
	PlaceName string
	NumberTravelDays string
	StartAndEndTime string
	TypeTravelVehicle string
	TourDate string
	NameHotel string
	TypeHotel string
}

type DubaiTour struct {
	TravelTour
	DubaiTourFee int
	ExtraChargeForAdultsInDoubleRoomDubaiTour int
	ExtraChargeForAdultsInSingleRoomDubaiTour int
	DubaiTourChildSurcharge int
	DubaiTourBabyFee int
}

type QatarTour struct {
	TravelTour
	QatarTourFee int
	ExtraChargeForAdultsInDoubleRoomQatarTour int
	ExtraChargeForAdultsInSingleRoomQatarTour int
	QatarTourChildSurcharge int
	QatarTourBabyFee int
}

type MaldivesTour struct {
	TravelTour
	MaldivesTourFee int
	ExtraChargeForAdultsInDoubleRoomMaldivesTour int
	ExtraChargeForAdultsInSingleRoomMaldivesTour int
	MaldivesTourChildSurcharge int
	MaldivesTourBabyFee int
}