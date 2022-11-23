package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/lukmanlukmin/golib/log"
)

func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}

	var validatorErrs validator.ValidationErrors
	if errors.As(err, &validatorErrs) {
		for _, e := range validatorErrs {
			translatedErr := fmt.Errorf(fmt.Sprintf("%v.", e.Translate(trans)))
			errs = append(errs, translatedErr)
		}
	} else {
		errs = append(errs, err)
	}

	return errs
}

func TranslatorValidator(v *validator.Validate) ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	if err := enTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		log.WithError(err).Warnln("failed register default translation")
	}

	return trans
}
