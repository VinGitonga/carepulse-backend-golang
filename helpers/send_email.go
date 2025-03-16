package helpers

import (
	"CarepluseBackend/config"
	"bytes"
	"fmt"
	"github.com/resend/resend-go/v2"
	"os"
	"path/filepath"
	"text/template"
)

type EmailTemplate string

const (
	WelcomeEmail EmailTemplate = "welcome"
)

type ISendEmailOptions struct {
	To       []string
	Subject  string
	Template EmailTemplate
	Data     *map[string]interface{}
}

func SendEmail(options ISendEmailOptions) (string, error) {
	htmlTemplatesPath := "assets/templates/emails"
	templateName := options.Template + ".html"

	templatePath := filepath.Join(htmlTemplatesPath, string(templateName))

	templateContent, err := os.ReadFile(templatePath)

	if err != nil {
		return "", fmt.Errorf("error reading template email path: %s: %v", templatePath, err)
	}

	// Parse html content
	tmpl, err := template.New("emailTemplate").Parse(string(templateContent))

	if err != nil {
		return "", fmt.Errorf("error parsing template content: %s: %v", templatePath, err)
	}

	// Create a buffer that holder the email content
	var buffer bytes.Buffer

	if err = tmpl.Execute(&buffer, &options.Data); err != nil {
		return "", fmt.Errorf("error executing template: %s: %v", templatePath, err)
	}

	emittedHtml := buffer.String()

	client := resend.NewClient(config.Config("RESEND_API_KEY"))

	params := &resend.SendEmailRequest{
		To:      options.To,
		Subject: options.Subject,
		From:    "Carepulse <support-carepulse@vingitonga.xyz>",
		Html:    emittedHtml,
	}

	sent, err := client.Emails.Send(params)

	if err != nil {
		return "", fmt.Errorf("error sending email: %v", err)
	}

	return fmt.Sprintf("email sent with id: %s", sent.Id), nil
}
