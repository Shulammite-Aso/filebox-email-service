package services

import (
	"context"
	"crypto/tls"
	"log"
	"strconv"

	gomail "gopkg.in/mail.v2"

	"github.com/Shulammite-Aso/filebox-email-service/pkg/config"
	"github.com/Shulammite-Aso/filebox-email-service/pkg/proto"
)

type Server struct {
	proto.UnimplementedEmailServiceServer
}

func (s *Server) SendEmail(ctx context.Context, req *proto.SendEmailRequest) (*proto.SendEmailResponse, error) {
	c, err := config.LoadConfig()
	if err != nil {
		return &proto.SendEmailResponse{}, err
	}

	emailPort, err := strconv.Atoi(c.EmailPort)

	if err != nil {
		return &proto.SendEmailResponse{}, err
	}
	m := gomail.NewMessage()

	m.SetHeader("From", c.Sender)

	m.SetHeader("To", req.User)

	m.SetHeader("Subject", "New file in box")

	m.SetBody("text/plain", "A new file has been sent to your box by "+req.FileSender+". Filename: "+req.FileSent)

	// Settings for SMTP server
	d := gomail.NewDialer(c.EmailHost, emailPort, c.Sender, c.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return &proto.SendEmailResponse{}, err
	}

	log.Printf("email sent to %v \n", req.User)

	return &proto.SendEmailResponse{
		Message: "Email sent",
	}, nil
}
