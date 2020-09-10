package superclog

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
)

const (
	// TemplateExternalRelease template for external releases
	TemplateExternalRelease = "ExternalRelease"

	// TemplateInternalQARelease template for internal QA releases
	TemplateInternalQARelease = "InternalQARelease"
)

func TemplateReleaseNotes(w io.Writer, t *template.Template, commits []Commit, categories Categories) error {
	//t := template.Must(template.New("markdown-internal-qa-release-notes.tmpl").ParseFiles("markdown-internal-qa-release-notes.tmpl"))
	a := struct {
		Commits    []Commit
		Categories Categories
	}{
		Commits:    commits,
		Categories: categories,
	}

	err := t.Execute(w, a)
	if err != nil {
		return err
	}
	return nil
}

func DefaultTemplate(name string) (*template.Template, error) {
	var file string
	switch name {
	case TemplateExternalRelease:
		file = "markdown-external-release-notes.tmpl"
	case TemplateInternalQARelease:
		file = "markdown-internal-qa-release-notes.tmpl"
	default:
		return nil, fmt.Errorf("template not known: %s", name)
	}

	f, err := Assets.Open(file)
	if err != nil {
		return nil, fmt.Errorf("built in template open: %s: %w", file, err)
	}

	bts, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	t, err := template.New(file).Parse(string(bts))
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}

	return t, nil

	//return template.New("markdown-internal-qa-release-notes.tmpl").ParseFiles("markdown-internal-qa-release-notes.tmpl")
}
