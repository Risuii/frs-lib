package frsProvider

import (
	"github.com/ttacon/libphonenumber"
)

type PhoneProvider interface {
	ConvertToE164(phone string, countryCode string) (string, error)
}

type GPhone struct{}

func (g *GPhone) ConvertToE164(phone string, countryCode string) (string, error) {
	var formattedPhoneNumber string

	phoneNumber, err := libphonenumber.Parse(phone, countryCode)
	if err != nil {
		return formattedPhoneNumber, err
	}

	formattedPhoneNumber = libphonenumber.Format(phoneNumber, libphonenumber.E164)

	return formattedPhoneNumber, nil
}
