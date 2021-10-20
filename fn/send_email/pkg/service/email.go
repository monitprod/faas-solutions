package service

import m "github.com/monitprod/core/pkg/models"

type EmailOptions struct {
	subject string
	body    string
}

type EmailService interface {
	SendToMany(recipients []m.User, opts EmailOptions) error
}

type EmailServiceImp struct {
}

func newEmailService() EmailService {
	return &EmailServiceImp{}
}

func (e *EmailServiceImp) SendToMany(recipients []m.User, opts EmailOptions) error {
	return nil
}
