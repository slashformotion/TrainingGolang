package html

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/slashformotion/TrainingGolang/PROJECTS/certificats/cert"
)

type HtmlSaver struct {
	OutputDir string
}

func New(outdir string) (*HtmlSaver, error) {
	var h *HtmlSaver
	err := os.MkdirAll(outdir, os.ModePerm)
	if err != nil {
		return h, err
	}
	h = &HtmlSaver{OutputDir: outdir}
	return h, nil
}

func (h *HtmlSaver) Save(cert cert.Cert) error {
	t, err := template.New("certificate").Parse(tpl)
	if err != nil {
		return err
	}

	fn := fmt.Sprintf("%v.html", cert.LabelTitle)
	path := path.Join(h.OutputDir, fn)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	t.Execute(f, cert)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

var tpl string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.LabelTitle}}</title>
</head>
<body>
    <h1>{{.LabelCompletion}}</h1>
    <h2>{{.LabelPresented}}</h2>
    <h2>{{.Name}}</h2>
    <h2>{{.LabelPartipation}}</h2>
    <h2>{{.LabelDate}}</h2>

</body>
</html>
`
