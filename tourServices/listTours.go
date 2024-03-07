package tourServices

import (
	"fmt"
	"go-aliBaba/entities"
)

var dubaiTours []entities.DubaiTour = ConstructionDubaiTours()

func ConstructionDubaiTours() []entities.DubaiTour {
	dubaiTours := []entities.DubaiTour{}
	dubaiTours = append(dubaiTours, entities.DubaiTour{TravelTour: entities.TravelTour{Id: 1, PlaceName: "Dubai",
	NumberTravelDays: "7 night And 8 days", StartAndEndTime: "9 night And 10 days", TypeTravelVehicle: "Airplane",
	TourDate: "From 4 to 13 Bahman", NameHotel: "Hotel1", TypeHotel: "5 Star"}, DubaiTourFee: 38e6, ExtraChargeForAdultsInDoubleRoomDubaiTour: 1800,
    ExtraChargeForAdultsInSingleRoomDubaiTour: 2300, DubaiTourChildSurcharge: 1600, DubaiTourBabyFee: 700 })

	dubaiTours = append(dubaiTours, entities.DubaiTour{TravelTour: entities.TravelTour{Id: 2, PlaceName: "Dubai",
	NumberTravelDays: "10 night And 11 days", StartAndEndTime: "12 night And 13 days", TypeTravelVehicle: "Airplane",
	TourDate: "From 5 to 17 Tir", NameHotel: "Hotel2", TypeHotel: "4 Star"}, DubaiTourFee: 35e6, ExtraChargeForAdultsInDoubleRoomDubaiTour: 1700,
    ExtraChargeForAdultsInSingleRoomDubaiTour: 2250, DubaiTourChildSurcharge: 1500, DubaiTourBabyFee: 600 })
    
	return dubaiTours
}

func ListDubaiTours() {
	for _, tour := range dubaiTours {
		fmt.Printf("DubaiTours: %+v\n", tour)
	}
}

var qatarTours []entities.QatarTour = ConstructionQatarTours()

func ConstructionQatarTours() []entities.QatarTour {
	qatarTours := []entities.QatarTour{}
	qatarTours = append(qatarTours, entities.QatarTour{TravelTour: entities.TravelTour{Id: 1, PlaceName: "Qatar",
	NumberTravelDays: "7 night And 8 days", StartAndEndTime: "9 night And 10 days", TypeTravelVehicle: "Airplane",
	TourDate: "From 4 to 13 Sfand", NameHotel: "Hotel1", TypeHotel: "5 Star"}, QatarTourFee: 43e6, ExtraChargeForAdultsInDoubleRoomQatarTour: 2000,
    ExtraChargeForAdultsInSingleRoomQatarTour: 2400, QatarTourChildSurcharge: 1800, QatarTourBabyFee: 800 })

	qatarTours = append(qatarTours, entities.QatarTour{TravelTour: entities.TravelTour{Id: 2, PlaceName: "Qatar",
	NumberTravelDays: "10 night And 11 days", StartAndEndTime: "12 night And 13 days", TypeTravelVehicle: "Airplane",
	TourDate: "From 8 to 20 Mordad", NameHotel: "Hotel2", TypeHotel: "5 Star"}, QatarTourFee: 45e6, ExtraChargeForAdultsInDoubleRoomQatarTour: 2400,
    ExtraChargeForAdultsInSingleRoomQatarTour: 2600, QatarTourChildSurcharge: 1800, QatarTourBabyFee: 1000 })
     
	return qatarTours
}

func ListQatarTours() {
	for _, tour := range qatarTours {
		fmt.Printf("QatarTours: %+v\n", tour)
	}
}

var maldivesTours []entities.MaldivesTour = ConstructionMaldivesTours()

func ConstructionMaldivesTours() []entities.MaldivesTour {
	maldivesTours := []entities.MaldivesTour{}
	maldivesTours = append(maldivesTours, entities.MaldivesTour{TravelTour: entities.TravelTour{Id: 1, PlaceName: "Maldives",
	NumberTravelDays: "9 night And 10 days", StartAndEndTime: "11 night And 12 days", TypeTravelVehicle: "Airplane",
	TourDate: "From 1 to 12 Khordad", NameHotel: "Hotel1", TypeHotel: "4 Star"}, MaldivesTourFee: 65e6, ExtraChargeForAdultsInDoubleRoomMaldivesTour: 2500,
    ExtraChargeForAdultsInSingleRoomMaldivesTour: 2700, MaldivesTourChildSurcharge: 2200, MaldivesTourBabyFee: 1300 })

	maldivesTours = append(maldivesTours, entities.MaldivesTour{TravelTour: entities.TravelTour{Id: 2, PlaceName: "Maldives",
	NumberTravelDays: "14 night And 15 days", StartAndEndTime: "16 night And 17 days", TypeTravelVehicle: "Airplane",
	TourDate: "From 1 to 17 Azar", NameHotel: "Hotel2", TypeHotel: "5 Star"}, MaldivesTourFee: 78e6, ExtraChargeForAdultsInDoubleRoomMaldivesTour: 2800,
    ExtraChargeForAdultsInSingleRoomMaldivesTour: 3100, MaldivesTourChildSurcharge: 2400, MaldivesTourBabyFee: 1500 })

	return maldivesTours
}

func ListMaldivesTours() {
	for _, tour := range maldivesTours {
		fmt.Printf("MaldivesTours: %+v\n", tour)
	}
}