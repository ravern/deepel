package deepl

/* Request */
type call struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  params `json:"params"`
}

type params struct {
	Jobs  []job    `json:"jobs,omitempty"`
	Texts []string `json:"texts,omitempty"`
	Lang  lang     `json:"lang"`
}

type job struct {
	DeSentenceBeginning string `json:"de_sentence_beginning"`
	RawEnSentence       string `json:"raw_en_sentence"`
	Kind                string `json:"kind"`
}

type lang struct {
	UserPreferredLangs     []string `json:"user_preferred_langs,omitempty"`
	SourceLangUserSelected string   `json:"source_lang_user_selected,omitempty"`
	SourceLangComputed     string   `json:"source_lang_computed,omitempty"`
	LangUserSelected       string   `json:"lang_user_selected,omitempty"`
	TargetLang             string   `json:"target_lang,omitempty"`
}

/* Response */
type reply struct {
	JSONRPC string `json:"jsonrpc"`
	Result  result `json:"result"`
}

type result struct {
	Lang            string        `json:"lang"`
	SourceLang      string        `json:"source_lang"`
	TargetLang      string        `json:"target_lang"`
	LangIsConfident int           `json:"lang_is_confident"`
	SplittedTexts   [][]string    `json:"splitted_texts"`
	Translations    []translation `json:"translations"`
}

type translation struct {
	Beams []beam `json:"beams"`
}

type beam struct {
	PostprocessedSentence string `json:"postprocessed_sentence"`
}
