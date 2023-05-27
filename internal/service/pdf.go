package service

import (
	"bytes"
	pdff "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/sirupsen/logrus"
	"html/template"
	"io"
	"log"
)

type PDF struct {
	t *template.Template
}

type Params struct {
}

func NewPDF(templatePath string) *PDF {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalln(err)
	}

	pdff.SetPath("pkg/wkhtmltopdf/wkhtmltopdf.exe")

	_, err = pdff.NewPDFGenerator()
	if err != nil {
		logrus.WithError(err).Fatal("can't find wkhtmltopdf.exe")
	}

	return &PDF{t: t}
}

func (r *PDF) GeneratePDF(out io.Writer, data Params) error {

	pdf, err := pdff.NewPDFGenerator()
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = r.t.Execute(buf, data); err != nil {
		return err
	}

	pdf.NoPdfCompression.Set(true)
	pdf.SetOutput(out)
	pdf.PageSize.Set(pdff.PageSizeA4)
	pdf.AddPage(pdff.NewPageReader(buf))

	return pdf.Create()
}
