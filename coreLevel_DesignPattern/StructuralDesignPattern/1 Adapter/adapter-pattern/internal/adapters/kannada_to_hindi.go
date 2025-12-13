package adapters

import (
	"adapter-pattern/internal/adaptees"
)

// Kannada → Hindi Adapter
type KannadaToHindiAdapter struct {
	Kn adaptees.KannadaSpeaker
}

func (a KannadaToHindiAdapter) SpeakHindi() string {
	return "नमस्ते अर्जुन (Translated from: " + a.Kn.SpeakKannada() + ")"
}
