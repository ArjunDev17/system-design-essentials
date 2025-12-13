package adapters

import (
	"adapter-pattern/internal/adaptees"
)

// Punjabi → Hindi Adapter
type PunjabiToHindiAdapter struct {
	Pb adaptees.PunjabiSpeaker
}

func (a PunjabiToHindiAdapter) SpeakHindi() string {
	return "नमस्ते अर्जुन (Translated from: " + a.Pb.SpeakPunjabi() + ")"
}
