package deepl_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ravernkoh/deepl"
)

func TestSplitIntoSentences(t *testing.T) {
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
		res, err := cli.SplitIntoSentences(test.Texts, deepl.English)
		if err != nil {
			t.Errorf("Test %d: Unexpected error %s.", i+1, err.Error())
			return
		}
		if !reflect.DeepEqual(test.Expected, res) {
			t.Errorf("Test %d: Expected %v but got %v.", i+1, test.Expected, res)
		}
	}
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		Sentences []string
		Expected  [][]string
	}{
		{
			[]string{
				"Hello world!",
				"This is the unofficial Go client for DeepL. Cheers!",
			},
			[][]string{
				{
					"Hallo Welt!",
					"Hallo, Welt!",
				},
				{
					"Das ist der inoffizielle Go-Client für DeepL. Prost!",
					"Das ist der inoffizielle Go-Kunde für DeepL. Prost!",
					"Dies ist der inoffizielle Go-Client für DeepL. Prost!",
					"Dies ist der inoffizielle Go-Kunde für DeepL. Prost!",
				},
			},
		},
	}

	cli := deepl.NewClient()

	for i, test := range tests {
		res, err := cli.Translate(test.Sentences, deepl.English, deepl.German)
		if err != nil {
			t.Errorf("Test %d: Unexpected error %s.", i+1, err.Error())
			return
		}
		if !reflect.DeepEqual(test.Expected, res) {
			t.Errorf("Test %d: Expected %v but got %v.", i+1, test.Expected, res)
		}
	}
}

func TestTranslateBegin(t *testing.T) {
	tests := []struct {
		Sentences  []string
		Beginnings []string
		Expected   [][]string
	}{
		{
			[]string{
				"Hello world!",
				"This is the unofficial Go client for DeepL. Cheers!",
			},
			[]string{
				"Welt",
				"Dies",
			},
			[][]string{
				{
					"Welt, hallo!",
					"Welt!",
					"Welt! Hallo!",
				},
				{
					"Dies ist der inoffizielle Go-Client für DeepL. Prost!",
					"Dies ist der inoffizielle Go-Kunde für DeepL. Prost!",
					"Dies ist der inoffizielle Go-Client für DeepL. Cheers!",
					"Dies ist der inoffizielle Go-Kunde für DeepL. Cheers!",
				},
			},
		},
	}

	cli := deepl.NewClient()

	for i, test := range tests {
		res, err := cli.TranslateBegin(test.Sentences, test.Beginnings, deepl.English, deepl.German)
		if err != nil {
			t.Errorf("Test %d: Unexpected error %s.", i+1, err.Error())
			return
		}
		if !reflect.DeepEqual(test.Expected, res) {
			for _, i := range res {
				for _, j := range i {
					fmt.Println(j)
				}
			}
			t.Errorf("Test %d: Expected %v but got %v.", i+1, test.Expected, res)
		}
	}
}
