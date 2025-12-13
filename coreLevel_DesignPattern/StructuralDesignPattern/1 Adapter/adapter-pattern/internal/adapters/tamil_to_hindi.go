package adapters

import (
	"adapter-pattern/internal/adaptees"
)

// Tamil → Hindi Adapter
type TamilToHindiAdapter struct {
	Tm adaptees.TamilSpeaker
}

func (a TamilToHindiAdapter) SpeakHindi() string {
	return "नमस्ते अर्जुन (Translated from: " + a.Tm.SpeakTamil() + ")"
}
