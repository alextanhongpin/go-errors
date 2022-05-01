package errors

import (
	"bytes"
	"fmt"
	"html/template"

	"golang.org/x/text/language"
)

var (
	errorLanguages = map[language.Tag]bool{
		language.English: true,
	}
	errorByCode = map[string]*Error{}
)

// AddLanguage adds new language support.
// Also checks if existing errors has the translations for that language.
func AddLanguage(lang language.Tag) bool {
	if errorLanguages[lang] {
		return false
	}
	for _, err := range errorByCode {
		_, ok := err.translations[lang]
		if !ok {
			panic(fmt.Errorf("missing language %q for error code %q", lang, err.Code))
		}
	}

	errorLanguages[lang] = true
	return true
}

type unmarshallFn = func(data []byte, v any) error

func MustRegister(lang language.Tag, raw []byte, fn unmarshallFn) bool {
	if err := Register(lang, raw, fn); err != nil {
		panic(err)
	}
	return true
}

func Register(lang language.Tag, raw []byte, fn unmarshallFn) error {
	var data map[string]map[string]map[string]string
	if err := fn(raw, &data); err != nil {
		return err
	}

	for kind, translationsByCode := range data {
		for code, messageByLanguage := range translationsByCode {
			if _, ok := errorByCode[code]; ok {
				return fmt.Errorf("error code exists: %s", code)
			}
			translations := make(map[language.Tag]string)
			for lang, message := range messageByLanguage {
				translations[language.MustParse(lang)] = message
			}
			for supportedLang := range errorLanguages {
				if _, ok := translations[supportedLang]; !ok {
					return fmt.Errorf("missing language %q for error code %q", supportedLang, code)
				}
			}

			errorByCode[code] = &Error{
				Code:         code,
				Kind:         kind,
				Message:      translations[lang],
				lang:         lang,
				translations: translations,
			}
		}
	}

	return nil
}

func New(code string) *Error {
	err, ok := errorByCode[code]
	if !ok {
		panic(fmt.Errorf("error code not found: %s", code))
	}

	cerr := *err
	cerr.translations = make(map[language.Tag]string)
	for k, v := range err.translations {
		cerr.translations[k] = v
	}
	return &cerr
}

func NewParams[T any](code string) *ErrorParams[T] {
	return NewErrorParams[T](New(code))
}

func makeTemplate(msg string, data any) (string, error) {
	t := template.Must(template.New("").Parse(msg))

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return msg, err
	}

	return buf.String(), nil
}
