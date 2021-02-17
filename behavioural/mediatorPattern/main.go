package main

import (
	"DesignPatterns/behavioural/mediatorPattern/mediator"
)

func main() {
    stationManager := mediator.NewStationManager()

    passengerTrain := &mediator.PassengerTrain{
    	Mediator: stationManager,
	}

	goodsTrain := &mediator.GoodsTrain{
		Mediator: stationManager,
	}

	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()
}
