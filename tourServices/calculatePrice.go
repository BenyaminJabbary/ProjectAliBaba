package tourServices

import (
	"fmt"
	"github.com/pkg/errors"
)

func DubaiTourPriceCalculation() error {
	logger.Warn().Str("calculatePrice", "DubaiTourPriceCalculation").Msg("Start of DubaiTourPriceCalculation")

	for _, tour := range dubaiTours {
		fmt.Println("\"Choose the number of passengers\"")

		var numberOfAdultsInDoubleRoom int
		fmt.Print("Number of adults in double room: ")
		fmt.Scanln(&numberOfAdultsInDoubleRoom)
		adultDoubleRoomPrice := numberOfAdultsInDoubleRoom * tour.DubaiTourFee
		extraChargeForAnAdultInDoubleRoom := numberOfAdultsInDoubleRoom * tour.ExtraChargeForAdultsInDoubleRoomDubaiTour

		var numberOfAdultsInSingleRoom int
		fmt.Print("Number of adults in single room: ")
		fmt.Scanln(&numberOfAdultsInSingleRoom)
		adultSingleRoomPrice := numberOfAdultsInSingleRoom * tour.DubaiTourFee
		extraChargeForAnAdultInSingleRoom := numberOfAdultsInSingleRoom * tour.ExtraChargeForAdultsInSingleRoomDubaiTour

		if numberOfAdultsInDoubleRoom == 0 && numberOfAdultsInSingleRoom == 0 {
			fmt.Println("Passengers must include at least one adult!")
			return errors.New("passengers must include at least one adult")
		}

		var numberOfChildren int
		fmt.Print("Number of children: ")
		fmt.Scanln(&numberOfChildren)
		tourPriceForAChild := numberOfChildren * tour.DubaiTourFee
		dubaiTourChildSurcharge := numberOfChildren * tour.DubaiTourChildSurcharge

		var numberOfBaby int
		fmt.Print("Number of baby: ")
		fmt.Scanln(&numberOfBaby)
		dubaiTourBabyFee := numberOfBaby * tour.DubaiTourBabyFee

		dubaiTourFee := adultDoubleRoomPrice + adultSingleRoomPrice + tourPriceForAChild
		extraCharge := extraChargeForAnAdultInDoubleRoom + extraChargeForAnAdultInSingleRoom + dubaiTourChildSurcharge + dubaiTourBabyFee
		fmt.Printf("DubaiTourFee: %d + %d+ Dollar\n", dubaiTourFee, extraCharge)

		ConfirmationOfReservation(numberOfAdultsInDoubleRoom, numberOfAdultsInSingleRoom, numberOfChildren, numberOfBaby)

		break
	}

	logger.Warn().Str("calculatePrice", "DubaiTourPriceCalculation").Msg("End of DubaiTourPriceCalculation")

	return nil
}

func QatarTourPriceCalculation() error {
	logger.Warn().Str("calculatePrice", "QatarTourPriceCalculation").Msg("Start of QatarTourPriceCalculation")

	for _, tour := range qatarTours {
		fmt.Println("\"Choose the number of passengers\"")

		var numberOfAdultsInDoubleRoom int
		fmt.Print("Number of adults in double room: ")
		fmt.Scanln(&numberOfAdultsInDoubleRoom)
		adultDoubleRoomPrice := numberOfAdultsInDoubleRoom * tour.QatarTourFee
		extraChargeForAnAdultInDoubleRoom := numberOfAdultsInDoubleRoom * tour.ExtraChargeForAdultsInDoubleRoomQatarTour

		var numberOfAdultsInSingleRoom int
		fmt.Print("Number of adults in single room: ")
		fmt.Scanln(&numberOfAdultsInSingleRoom)
		adultSingleRoomPrice := numberOfAdultsInSingleRoom * tour.QatarTourFee
		extraChargeForAnAdultInSingleRoom := numberOfAdultsInSingleRoom * tour.ExtraChargeForAdultsInSingleRoomQatarTour

		if numberOfAdultsInDoubleRoom == 0 && numberOfAdultsInSingleRoom == 0 {
			fmt.Println("Passengers must include at least one adult!")
			return errors.New("passengers must include at least one adult")
		}

		var numberOfChildren int
		fmt.Print("Number of children: ")
		fmt.Scanln(&numberOfChildren)
		tourPriceForAChild := numberOfChildren * tour.QatarTourFee
		qatarTourChildSurcharge := numberOfChildren * tour.QatarTourChildSurcharge

		var numberOfBaby int
		fmt.Print("Number of baby: ")
		fmt.Scanln(&numberOfBaby)
		qatarTourBabyFee := numberOfBaby * tour.QatarTourBabyFee

		qatarTourFee := adultDoubleRoomPrice + adultSingleRoomPrice + tourPriceForAChild
		extraCharge := extraChargeForAnAdultInDoubleRoom + extraChargeForAnAdultInSingleRoom + qatarTourChildSurcharge + qatarTourBabyFee
		fmt.Printf("QatarTourFee: %d + %d+ Dollar\n", qatarTourFee, extraCharge)

		ConfirmationOfReservation(numberOfAdultsInDoubleRoom, numberOfAdultsInSingleRoom, numberOfChildren, numberOfBaby)

		break
	}

	logger.Warn().Str("calculatePrice", "QatarTourPriceCalculation").Msg("End of QatarTourPriceCalculation")

	return nil
}

func MaldivesTourPriceCalculation() error {
	logger.Warn().Str("calculatePrice", "MaldivesTourPriceCalculation").Msg("Start of MaldivesTourPriceCalculation")

	for _, tour := range maldivesTours {
		fmt.Println("\"Choose the number of passengers\"")

		var numberOfAdultsInDoubleRoom int
		fmt.Print("Number of adults in double room: ")
		fmt.Scanln(&numberOfAdultsInDoubleRoom)
		adultDoubleRoomPrice := numberOfAdultsInDoubleRoom * tour.MaldivesTourFee
		extraChargeForAnAdultInDoubleRoom := numberOfAdultsInDoubleRoom * tour.ExtraChargeForAdultsInDoubleRoomMaldivesTour

		var numberOfAdultsInSingleRoom int
		fmt.Print("Number of adults in single room: ")
		fmt.Scanln(&numberOfAdultsInSingleRoom)
		adultSingleRoomPrice := numberOfAdultsInSingleRoom * tour.MaldivesTourFee
		extraChargeForAnAdultInSingleRoom := numberOfAdultsInSingleRoom * tour.ExtraChargeForAdultsInSingleRoomMaldivesTour

		if numberOfAdultsInDoubleRoom == 0 && numberOfAdultsInSingleRoom == 0 {
			fmt.Println("Passengers must include at least one adult!")
			return errors.New("passengers must include at least one adult")
		}

		var numberOfChildren int
		fmt.Print("Number of children: ")
		fmt.Scanln(&numberOfChildren)
		tourPriceForChild := numberOfChildren * tour.MaldivesTourFee
		maldivesTourChildSurcharge := numberOfChildren * tour.MaldivesTourChildSurcharge

		var numberOfBaby int
		fmt.Print("Number of baby: ")
		fmt.Scanln(&numberOfBaby)
		maldivesTourBabyFee := numberOfBaby * tour.MaldivesTourBabyFee

		maldivesTourFee := adultDoubleRoomPrice + adultSingleRoomPrice + tourPriceForChild
		extraCharge := extraChargeForAnAdultInDoubleRoom + extraChargeForAnAdultInSingleRoom + maldivesTourChildSurcharge + maldivesTourBabyFee
		fmt.Printf("MaldivesTourFee: %d + %d+ Dollar\n", maldivesTourFee, extraCharge)

		ConfirmationOfReservation(numberOfAdultsInDoubleRoom, numberOfAdultsInSingleRoom, numberOfChildren, numberOfBaby)

		break
	}

	logger.Warn().Str("calculatePrice", "MaldivesTourPriceCalculation").Msg("End of MaldivesTourPriceCalculation")

	return nil
}
