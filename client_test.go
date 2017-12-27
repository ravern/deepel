package deepl_test

import (
	"reflect"
	"testing"

	"github.com/ravernkoh/deepl"
)

func TestSplitSentencesRequest(t *testing.T) {
	tests := []struct {
		Texts    []string
		Expected [][]string
	}{
		{
			[]string{
				"Hello world!",
				"This is the unofficial Go client for DeepL. Cheers!",
				"These, are some tests.",
			},
			[][]string{
				{"Hello world!"},
				{"This is the unofficial Go client for DeepL.", "Cheers!"},
				{"These, are some tests."},
			},
		},
	}

	cli := deepl.NewClient()

	for i, test := range tests {
		res, err := cli.SplitSentences(test.Texts, deepl.English)
		if err != nil {
			t.Errorf("Test %d: Unexpected error %s.", i+1, err.Error())
			return
		}
		if !reflect.DeepEqual(test.Expected, res) {
			t.Errorf("Test %d: Expected %v but got %v.", i+1, test.Expected, res)
		}
	}
}
