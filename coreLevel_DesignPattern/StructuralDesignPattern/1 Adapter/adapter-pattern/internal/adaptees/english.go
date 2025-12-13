package adaptees

type EnglishSpeaker struct{}

func (EnglishSpeaker) SpeakEnglish() string {
	return "Hello Arjun"
}
