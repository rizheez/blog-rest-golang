package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate *validator.Validate
	Trans    ut.Translator
)

func InitValidator() {
	Validate = validator.New()

	// Setup locale & translator
	locale := en.New()
	uni := ut.New(locale, locale)
	Trans, _ = uni.GetTranslator("en")

	// Register default translation
	_ = en_translations.RegisterDefaultTranslations(Validate, Trans)
}
