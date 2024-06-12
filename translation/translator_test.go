package translation_test

import (
	"testing"

	"github.com/luzeduardo/shipping-go/translation"
)

func TestTranslate(t *testing.T) {
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		}, {
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		}, {
			Word:        "Hello",
			Language:    "finnish",
			Translation: "hei",
		}, {
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		}, {
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
	}

	for _, ts := range tt {

		res := translation.Translate(ts.Word, ts.Language)
		if res != ts.Translation {
			t.Errorf(`expected "%s" but received "%s" from "%s"`, ts.Translation, res, ts.Language)
		}
	}
}
