package service

import (
	"bytes"
	gen "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/sirupsen/logrus"
	"html/template"
	"io"
)

type PDF struct {
	t *template.Template
}

func NewPDF(templatePath string) *PDF {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		logrus.WithError(err).Fatal("can't parse template file")
	}
	return &PDF{t: t}
}

func (r *PDF) GeneratePDF(out io.Writer, data any) error {
	buf := new(bytes.Buffer)

	if err := r.t.Execute(buf, data); err != nil {
		return err
	}

	pdf, err := gen.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdf.NoPdfCompression.Set(true)
	pdf.PageSize.Set(gen.PageSizeA4)
	pdf.Dpi.Set(300)
	pdf.SetOutput(out)
	pdf.AddPage(gen.NewPageReader(buf))

	return pdf.Create()
}
