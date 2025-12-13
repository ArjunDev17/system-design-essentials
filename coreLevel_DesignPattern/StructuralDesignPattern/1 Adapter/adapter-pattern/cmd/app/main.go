package main

import (
	"fmt"
	"adapter-pattern/internal/adapters"
	"adapter-pattern/internal/adaptees"
	"adapter-pattern/internal/service"
)

func main() {

	fmt.Println("---- Adapter Pattern Demo ----")

	// English incoming
	engAdapter := adapters.EnglishToHindiAdapter{Eng: adaptees.EnglishSpeaker{}}
	service.ProcessMessage(engAdapter)

	// Kannada incoming
	knAdapter := adapters.KannadaToHindiAdapter{Kn: adaptees.KannadaSpeaker{}}
	service.ProcessMessage(knAdapter)

	// Tamil incoming
	tmAdapter := adapters.TamilToHindiAdapter{Tm: adaptees.TamilSpeaker{}}
	service.ProcessMessage(tmAdapter)

	// Punjabi incoming
	
}