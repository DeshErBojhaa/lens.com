package views

import (
	"errors"
	"html/template"
	"log"

	"lens.com/models"
)

const (
	AlertLevError   = "danger"
	AlertLevWarning = "warning"
	AlertLevInfo    = "info"
	AlertLevSuccess = "success"
	AlertMsgGeneric = "something went wrong"
)

type Alert struct {
	Level   string
	Message string
}

type Data struct {
	Alert *Alert
	User  *models.User
	CSRF  template.HTML
	Yield interface{}
}

func (d *Data) SetAlert(err error) {
	var pErr PublicError
	if errors.As(err, &pErr) {
		d.Alert = &Alert{
			Level:   AlertLevError,
			Message: pErr.Public(),
		}
	} else {
		log.Println(err)
		d.Alert = &Alert{
			Level:   AlertLevError,
			Message: AlertMsgGeneric,
		}
	}
}

func (d *Data) AlertError(msg string) {
	d.Alert = &Alert{
		Level:   AlertLevError,
		Message: msg,
	}
}

type PublicError interface {
	error
	Public() string
}
