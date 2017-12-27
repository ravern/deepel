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

	return r.Result.SplittedTexts, nil
}
