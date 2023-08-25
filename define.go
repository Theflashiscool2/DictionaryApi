package dictionary

import (
	"github.com/gocolly/colly"
	"strings"
)

func TranslateWord(word string, lang Language) string {
	var newWord string
	c := colly.NewCollector()
	c.OnHTML("body > div.container > div:nth-child(2) > div > div.panel-body > table:nth-child(1) > tbody > tr:nth-child(1) > td:nth-child(2)", func(e *colly.HTMLElement) {
		newWord = e.Text
	})
	url := "http://www.langtolang.com/?txtLang=" + word + "&submitButton=Search&selectFrom=" + string(lang) + "&selectTo=english"
	_ = c.Visit(url)
	return newWord
}

func Define(word string, lang Language) string {
	if lang != LanguageEnglish {
		word = TranslateWord(word, lang)
	}
	var definition string
	c := colly.NewCollector()
	c.OnHTML("span.definition", func(e *colly.HTMLElement) {
		definition = e.Text
	})
	url := "http://www.langtolang.com/?txtLang=" + word + "&submitButton=Search&selectFrom=english&selectTo=spanish"
	_ = c.Visit(url)
	return strings.TrimSpace(definition)
}

func Example(word string, lang Language) string {
	if lang != LanguageEnglish {
		word = TranslateWord(word, lang)
	}
	var definition string
	c := colly.NewCollector()
	c.OnHTML("span.example", func(e *colly.HTMLElement) {
		definition = e.Text
	})
	url := "http://www.langtolang.com/?txtLang=" + word + "&submitButton=Search&selectFrom=english&selectTo=spanish"
	_ = c.Visit(url)
	return strings.TrimSpace(definition)
}

func Synonyms(word string, lang Language) string {
	if lang != LanguageEnglish {
		word = TranslateWord(word, lang)
	}
	var synonyms string
	c := colly.NewCollector()
	c.OnHTML("span.synonym", func(e *colly.HTMLElement) {
		synonyms = e.Text
	})
	url := "http://www.langtolang.com/?txtLang=" + word + "&submitButton=Search&selectFrom=english&selectTo=spanish"
	_ = c.Visit(url)
	return strings.Trim(synonyms[6:], "]")
}
