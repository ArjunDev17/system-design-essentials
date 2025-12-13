package hindi

// HindiSpeaker is the target interface.
// Our main system ONLY understands Hindi.
type HindiSpeaker interface {
	SpeakHindi() string
}
