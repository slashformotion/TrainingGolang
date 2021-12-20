package pdf

import (
	"fmt"
	"os"
	"path"

	gofpdf "github.com/jung-kurt/gofpdf"
	"github.com/slashformotion/TrainingGolang/PROJECTS/certificats/cert"
)

type PdfSaver struct {
	OutputDir string
}

func New(outdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outdir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{OutputDir: outdir}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()
	//bg
	bg(pdf)
	// header
	header(pdf, &cert)
	pdf.Ln(30)
	//body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)
	pdf.SetFont("Helvetica", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")
	pdf.Ln(30)

	footer(pdf)
	//savefile
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

func bg(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	w, h := pdf.GetPageSize()
	pdf.ImageOptions("img/bg.png", 0, 0, w, h, false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, cert *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	w, h := pdf.GetPageSize()
	wu := w / 10
	hu := h / 6
	pdf.ImageOptions("img/gophers.png", wu*2, hu, wu, hu, false, opts, 0, "")
	pdf.ImageOptions("img/gophers.png", w-wu*3, hu, wu, hu, false, opts, 0, "")

	pdf.SetFont("Helvetica", "", 30)
	pdf.WriteAligned(0, 80, cert.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	w, h := pdf.GetPageSize()
	wu := w / 6
	hu := h / 6
	pdf.ImageOptions("img/stamp.png", w-wu*2, h-hu*2, wu, hu, false, opts, 0, "")

}
