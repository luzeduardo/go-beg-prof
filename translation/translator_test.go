package translation_test

import (
	"testing"

	"github.com/luzeduardo/shipping-go/translation"
)

func TestTranslate(t *testing.T) {
	word := "hello"
	language := "english"

	res := translation.Translate(word, language)
	if res != "hello" {
		t.Errorf(`expected "hello" but received "%s"`, res)
	}

	res = translation.Translate("hello", "german")
	if res != "hallo" {
		t.Errorf(`expected "hei" but received "%s"`, res)
	}

	res = translation.Translate("hello", "finnish")
	if res != "hei" {
		t.Errorf(`expected "hei" but received "%s"`, res)
	}

	res = translation.Translate("hello", "dutch")
	if res != "" {
		t.Errorf(`expected "hei" but received "%s"`, res)
	}
}
