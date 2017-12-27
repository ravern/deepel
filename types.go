package deepl

// body of the POST request
type body struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Jobs []struct {
			DeSentenceBeginning string `json:"de_setence_beginning"`
			RawEnSetence        string `json:"raw_en_sentence"`
			Kind                string `json:"kind"`
		} `json:"jobs,omitempty"`
		Texts string `json:"texts,omitempty"`
		Lang  struct {
			UserPreferredLangs     []string `json:"user_preferred_langs,omitempty"`
			SourceLangUserSelected string   `json:"source_lang_user_selected,omitempty"`
			SourceLangComputed     string   `json:"source_lang_computed,omitempty"`
			LangUserSelected       string   `json:"lang_user_selected,omitempty"`
			TargetLang             string   `json:"target_lang,omitempty"`
		} `json:"lang"`
	} `json:"params"`
}
