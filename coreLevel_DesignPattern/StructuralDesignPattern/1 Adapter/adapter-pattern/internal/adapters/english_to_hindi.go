package adapters

import (
	"adapter-pattern/internal/adaptees"
)

// English → Hindi Adapter
type EnglishToHindiAdapter struct {
	Eng adaptees.EnglishSpeaker
}

func (a EnglishToHindiAdapter) SpeakHindi() string {
	return "नमस्ते अर्जुन (Translated from: " + a.Eng.SpeakEnglish() + ")"
}
