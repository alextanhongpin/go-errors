package errors

import (
	"fmt"

	"golang.org/x/text/language"
)

type Code string
type Kind string

type Bundle struct {
	errorByCode map[Code]*Error
	kinds       map[Kind]bool
	langs       map[language.Tag]bool
	defaultLang language.Tag
}

func NewBundle(defaultLang language.Tag, langs []language.Tag, kinds []Kind) *Bundle {
	langs = append(langs, defaultLang)

	langMap := make(map[language.Tag]bool)
	for _, lang := range langs {
		langMap[lang] = true
	}

	kindMap := make(map[Kind]bool)
	for _, kind := range kinds {
		kindMap[kind] = true
	}

	return &Bundle{
		defaultLang: defaultLang,
		langs:       langMap,
		kinds:       kindMap,
		errorByCode: make(map[Code]*Error),
	}
}

type unmarshallFn = func(data []byte, v any) error

func (b *Bundle) Load(errorBytes []byte, fn unmarshallFn) error {
	var data map[Kind]map[Code]map[string]string
	if err := fn(errorBytes, &data); err != nil {
		return err
	}

	for kind, translationsByCode := range data {
		if !b.kinds[kind] {
			return fmt.Errorf("%w: %s", ErrInvalidKind, kind)
		}

		for code, messageByLanguage := range translationsByCode {
			if _, ok := b.errorByCode[code]; ok {
				return fmt.Errorf("%w: %s", ErrDuplicateCode, code)
			}

			translations := make(map[language.Tag]string)
			for lang, message := range messageByLanguage {
				translations[language.MustParse(lang)] = message
			}

			for lang := range b.langs {
				if _, ok := translations[lang]; !ok {
					return fmt.Errorf("%w: %q.%q", ErrTranslationUndefined, code, lang)
				}
			}

			b.errorByCode[code] = &Error{
				Code:         string(code),
				Kind:         string(kind),
				Message:      translations[b.defaultLang],
				Params:       nil,
				lang:         b.defaultLang,
				translations: translations,
			}
		}
	}

	return nil
}

func (b *Bundle) MustLoad(errorBytes []byte, fn unmarshallFn) bool {
	if err := b.Load(errorBytes, fn); err != nil {
		panic(err)
	}

	return true
}

func (b *Bundle) Code(code Code) *Error {
	err, ok := b.errorByCode[code]
	if !ok {
		panic(fmt.Errorf("%w: %s", ErrCodeNotFound, code))
	}

	return err.Clone()
}
