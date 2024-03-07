package tourServices

import (
	"github.com/pkg/errors"
	"fmt"
	"go-aliBaba/core"
)

var logger = core.NewFileLogger()

func TourReservation() {
	logger.Warn().Str("tourReservation", "TourReservation").Msg("Start of ReservationTour")

	var chooseLocation string
	fmt.Print("Choose a place: ")
	fmt.Scanln(&chooseLocation)

	var id int
	fmt.Print("ID: ")
	fmt.Scanln(&id)
	switch chooseLocation {
	case "Dubai":
		err := GetDubaiTours(id)
		if err != nil {
			logger.Error().Stack().Err(err).Str("tourReservation", "TourReservation").Msg("There is an error in the tour reservation")
		}
	case "Qatar":
		err := GetQatarTours(id)
		if err != nil {
			logger.Error().Stack().Err(err).Str("tourReservation", "TourReservation").Msg("There is an error in the tour reservation")
		}
	case "Maldives":
		err := GetMaldivesTours(id)
		if err != nil {
			logger.Error().Stack().Err(err).Str("tourReservation", "TourReservation").Msg("There is an error in the tour reservation")
		}
	default:
		err := errors.New("Tour is not available at this place")
		logger.Error().Err(err).Str("tourReservation", "TourReservation").Msg("Error in the (func TourReservation)")
	}

	logger.Warn().Str("tourReservation", "TourReservation").Msg("End of ReservationTour")
}

func GetDubaiTours(id int) error {
	logger.Warn().Str("tourReservation", "GetDubaiTours").Msg("Start of GetDubaiTours")

	for _, tour := range dubaiTours {
		if tour.Id == id {
			fmt.Printf("DubaiTour: %+v\n", tour)
			err := DubaiTourPriceCalculation()
			if err != nil {
				return err
			}
		}
	}

	logger.Warn().Str("tourReservation", "GetDubaiTours").Msg("End of GetDubaiTours")

	return nil
}

func GetQatarTours(id int) error {
	logger.Warn().Str("tourReservation", "GetQatarTours").Msg("Start of GetQatarTours")

	for _, tour := range qatarTours {
		if tour.Id == id {
			fmt.Printf("QatarTour: %+v\n", tour)
			err := QatarTourPriceCalculation()
			if err != nil {
				return err
			}
		}
	}

	logger.Warn().Str("tourReservation", "GetQatarTours").Msg("Start of GetQatarTours")

	return nil
}

func GetMaldivesTours(id int) error {
	logger.Warn().Str("tourReservation", "GetMaldivesTours").Msg("Start of GetMaldivesTours")

	for _, tour := range maldivesTours {
		if tour.Id == id {
			fmt.Printf("MaldivesTour: %+v\n", tour)
			err := MaldivesTourPriceCalculation()
			if err != nil {
				return err
			}
		}
	}

	logger.Warn().Str("tourReservation", "GetMaldivesTours").Msg("Start of GetMaldivesTours")

	return nil
}