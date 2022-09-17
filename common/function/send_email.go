package function

import (
	"github.com/jau1jz/cornus"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var emailConfig struct {
	Email struct {
		Key  string `yaml:"key"`
		From string `yaml:"from"`
	} `yaml:"email"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&emailConfig)
}

type EmailUser struct {
	Name    string
	Address string
}

func NewEmailUser(name string, address string) EmailUser {
	return EmailUser{
		Name:    name,
		Address: address,
	}
}

func SendEmail(toUser EmailUser, subject string, plainTextContent string, htmlContent string) (string, error) {
	from := mail.NewEmail(emailConfig.Email.From, emailConfig.Email.From)
	to := mail.NewEmail(toUser.Name, toUser.Address)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(emailConfig.Email.Key)
	response, err := client.Send(message)
	if err != nil {
		return response.Body, err
	} else {
		return response.Body, nil
	}
}
