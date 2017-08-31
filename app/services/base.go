package services

import (
	"errors"
	"go-chat/app/utils"
)

// Base -
type Base struct {
}

func (s *Base) existsMailAdress() error {
	return errors.New(utils.I18n.Translate("validation_error.is_exists_mailadress"))
}

func (s *Base) notFound() error {
	return errors.New(utils.I18n.Translate("user_admin.error.not_found"))
}
