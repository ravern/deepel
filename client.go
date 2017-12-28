package deepl

import (
	"time"
)

// Client is a configurable client for the DeepL API.
type Client struct {
	Config
}

// Config for the Client.
type Config struct {
	PreferredLangs []string      // Preferred languages of the client.
	Timeout        time.Duration // Timeout for requests.
}

// The default configuration used for clients.
var DefaultConfig = Config{
	PreferredLangs: []string{English},
	Timeout:        10 * time.Second,
}

// NewClient returns a new client with the DefaultConfig.
func NewClient() *Client {
	return &Client{
		Config: DefaultConfig,
	}
}

// Construct the default body from the client configuration
func (cli *Client) newCall(method string) call {
	return call{
		JSONRPC: "2.0",
		Method:  method,
		Params: params{
			Lang: lang{
				UserPreferredLangs: cli.PreferredLangs,
			},
		},
	}
}

// SplitIntoSentences splits the texts into sentences via the DeepL API. Each text given
// will be split into its own slice of sentences.
func (cli *Client) SplitIntoSentences(texts []string, lang string) ([][]string, error) {
	c := cli.newCall("LMT_split_into_sentences")
	c.Params.Lang.LangUserSelected = lang
	c.Params.Texts = texts

	r, err := request(c, cli.Timeout)
	if err != nil {
		return nil, err
	}

	sens := r.Result.SplittedTexts

	return sens, nil
}

// Translate translates the given sentences to the target language via the DeepL API,
// returning all the possible translations. They are returned in descending order of
// confidence (i.e. Most to least confident).
func (cli *Client) Translate(stcs []string, source, target string) ([][]string, error) {
	begin := []string{}
	for _ = range stcs {
		begin = append(begin, "")
	}
	return cli.TranslateBegin(stcs, begin, source, target)
}

func (c *call) addJobs(stcs, begin []string, kind string) {
	for i, stc := range stcs {
		j := job{
			DeSentenceBeginning: begin[i],
			RawEnSentence:       stc,
			Kind:                kind,
		}
		c.Params.Jobs = append(c.Params.Jobs, j)
	}
}

func (r *reply) mapTranslations() [][]string {
	trans := [][]string{}
	for _, t := range r.Result.Translations {
		trans = append(trans, []string{})
		for _, b := range t.Beams {
			cur := len(trans) - 1
			trans[cur] = append(trans[cur], b.PostprocessedSentence)
		}
	}
	return trans
}

// TranslateBegin is the same as Translate, with the added constraint that the result
// must start with the given beginnings.
func (cli *Client) TranslateBegin(stcs, begin []string, source, target string) ([][]string, error) {
	c := cli.newCall("LMT_handle_jobs")
	c.Params.Lang.SourceLangUserSelected = source
	c.Params.Lang.TargetLang = target
	c.addJobs(stcs, begin, "default")

	r, err := request(c, cli.Timeout)
	if err != nil {
		return nil, err
	}

	return r.mapTranslations(), nil
}

// Alternatives returns the alternative beginnings to the given beginning in the context
// of the sentence.
func (cli *Client) Alternatives(stcs, begin []string, source, target string) ([][]string, error) {
	c := cli.newCall("LMT_handle_jobs")
	c.Params.Lang.SourceLangComputed = source
	c.Params.Lang.TargetLang = target
	c.addJobs(stcs, begin, "alternatives_at_position")

	r, err := request(c, cli.Timeout)
	if err != nil {
		return nil, err
	}

	return r.mapTranslations(), nil
}
