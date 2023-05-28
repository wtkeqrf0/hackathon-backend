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
	IndustryBranch    string
	OrganizationType  string
	FullTimeEmployers int
	LandArea          float64
	WageFund          float64
	InsurancePayment  float64
	IncomeTax         float64
	LandValueMin      float64
	LandValueMax      float64
	LandValue         float64
	Equipment         float64
	Taxes             float64
	Other             *string
	Total             float64
	///////
	Staff               float64
	RentalProperty      float64
	Services            float64
	StaffNum            int
	MinStaffMaintenance float64
	MaxStaffMaintenance float64
	MinPensionInsurance float64
	MaxPensionInsurance float64
	MinHealthInsurance  float64
	MaxHealthInsurance  float64
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
