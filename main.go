package main

import (
	"github.com/pkg/errors"
	"fmt"
	"go-aliBaba/core"
	"go-aliBaba/tourServices"
)

var logger = core.NewFileLogger()

func main() {
	logger.Info().Msg("Start of main")

	var choose = ""
	for choose != "exit" {
	    fmt.Println("Place(Dubai, Qatar, Maldives,...) : List of tours of the place")
	    fmt.Println("Reserve : Tour reservation")
	    fmt.Print("Choose: ")
	    fmt.Scanln(&choose)
	    switch choose {
	    case "Dubai" :
		    tourServices.ListDubaiTours()
		case "Qatar" :
			tourServices.ListQatarTours()
		case "Maldives" :
			tourServices.ListMaldivesTours()
	    case "Reserve" :
	        tourServices.TourReservation()
	    case "exit" :
	    	fmt.Println("Exiting...")
	    default:
			err := errors.New("Invalid found")
		    logger.Error().Err(err).Msg("Error in the main")
		}
	}

	logger.Info().Msg("End of main")
}