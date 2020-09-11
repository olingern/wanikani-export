package main

// Reading represents the "on" or "kun" reading for a character
type Reading struct {
	AcceptedAnswer bool   `json:"accepted_answer"`
	Primary        bool   `json:"primary"`
	Reading        string `json:"reading"`
	Type           string `json:"type"`
}

// ContextSentences represents a sentence associated with
// a character
type ContextSentences struct {
	En string `json:"en"`
	Ja string `json:"ja"`
}

// Meaning represents a character meaning
type Meaning struct {
	AcceptedAnswer bool   `json:"accepted_answer"`
	Meaning        string `json:"meaning"`
	Primary        bool   `json:"primary"`
}

// KanjiData is the nested data structure that's most interesting
// to this application
type KanjiData struct {
	AmalgamationSubjectIds    []int64              `json:"amalgamation_subject_ids"`
	AuxiliaryMeanings         []AuxiliaryMeanings  `json:"auxiliary_meanings"`
	CharacterImages           []CharacterImages    `json:"character_images"`
	Characters                string               `json:"characters"`
	ComponentSubjectIds       []int64              `json:"component_subject_ids"`
	ContextSentences          []ContextSentences   `json:"context_sentences"`
	CreatedAt                 string               `json:"created_at"`
	DocumentURL               string               `json:"document_url"`
	HiddenAt                  interface{}          `json:"hidden_at"`
	LessonPosition            int64                `json:"lesson_position"`
	Level                     int64                `json:"level"`
	MeaningHint               string               `json:"meaning_hint"`
	MeaningMnemonic           string               `json:"meaning_mnemonic"`
	Meanings                  []Meaning            `json:"meanings"`
	PartsOfSpeech             []string             `json:"parts_of_speech"`
	PronunciationAudios       []PronunciationAudio `json:"pronunciation_audios"`
	ReadingHint               string               `json:"reading_hint"`
	ReadingMnemonic           string               `json:"reading_mnemonic"`
	Readings                  []Reading            `json:"readings"`
	Slug                      string               `json:"slug"`
	SpacedRepetitionSystemID  int64                `json:"spaced_repetition_system_id"`
	VisuallySimilarSubjectIds []int64              `json:"visually_similar_subject_ids"`
}

// APIResponse is the root of wanikani's subject response
type APIResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  []struct {
		Data          KanjiData `json:"data"`
		DataUpdatedAt string    `json:"data_updated_at"`
		ID            int64     `json:"id"`
		Object        string    `json:"object"`
		URL           string    `json:"url"`
	} `json:"data"`
	DataUpdatedAt string `json:"data_updated_at"`
	Object        string `json:"object"`
	Pages         struct {
		NextURL     interface{} `json:"next_url"`
		PerPage     int64       `json:"per_page"`
		PreviousURL interface{} `json:"previous_url"`
	} `json:"pages"`
	TotalCount int64  `json:"total_count"`
	URL        string `json:"url"`
}

/**
UNUSED API DATASTRUCTURES
*/

// CharacterImages is unused
type CharacterImages struct {
	ContentType string `json:"content_type"`
	Metadata    struct {
		Color        string `json:"color"`
		Dimensions   string `json:"dimensions"`
		InlineStyles bool   `json:"inline_styles"`
		StyleName    string `json:"style_name"`
	} `json:"metadata"`
	URL string `json:"url"`
}

// AuxiliaryMeanings is unused
type AuxiliaryMeanings []struct {
	Meaning string `json:"meaning"`
	Type    string `json:"type"`
}

// PronunciationAudio is unused
type PronunciationAudio struct {
	ContentType string `json:"content_type"`
	Metadata    struct {
		Gender           string `json:"gender"`
		Pronunciation    string `json:"pronunciation"`
		SourceID         int64  `json:"source_id"`
		VoiceActorID     int64  `json:"voice_actor_id"`
		VoiceActorName   string `json:"voice_actor_name"`
		VoiceDescription string `json:"voice_description"`
	} `json:"metadata"`
	URL string `json:"url"`
}
