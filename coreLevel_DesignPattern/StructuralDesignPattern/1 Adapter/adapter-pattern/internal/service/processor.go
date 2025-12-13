package service

import (
	"fmt"
	"adapter-pattern/internal/domain/hindi"
)

// ProcessMessage is the main service that accepts only Hindi
func ProcessMessage(h hindi.HindiSpeaker) {
	fmt.Println("Main Service Received:", h.SpeakHindi())
}
