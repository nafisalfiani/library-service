package entity

import (
	"bytes"
	"html/template"
)

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

func (m *Mail) ParseHtml(req XenditPaymentResponse) (string, error) {
	// Define the HTML template
	const htmlTemplate = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Payment Details</title>
	</head>
	<body>
	
	    <h1>Payment Details</h1>
	
	    <ul>
	        <li><strong>External Payment ID:</strong> {{.XenditPaymentId}}</li>
	        <li><strong>Payment ID:</strong> {{.PaymentId}}</li>
	        <li><strong>Invoice Expiry Date:</strong> {{.InvoiceExpiryDate}}</li>
	        <li><strong>Invoice Status:</strong> {{.InvoiceStatus}}</li>
	        <li><strong>Invoice Amount:</strong> {{.InvoiceAmount}}</li>
	        <li><strong>Invoice URL:</strong> <a href="{{.InvoiceUrl}}" target="_blank">{{.InvoiceUrl}}</a></li>
	        <li><strong>Payment Method:</strong> {{.PaymentMethod}}</li>
	    </ul>
	
	</body>
	</html>
	`

	// Create a template from the HTML string
	tmpl, err := template.New("emailTemplate").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}

	// Create a buffer to write the template output
	var bodyContentBuffer = new(bytes.Buffer)

	// Execute the template with the data and write to the buffer
	if err := tmpl.Execute(bodyContentBuffer, req); err != nil {
		return "", err
	}

	return bodyContentBuffer.String(), nil
}
