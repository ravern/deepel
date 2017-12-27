package deepl

type SplitSentencesRequest struct {
	Texts []string
	Lang  string
}

func NewSplitSentencesRequest(texts []string, lang string) *SplitSentencesRequest {
	return nil
}
