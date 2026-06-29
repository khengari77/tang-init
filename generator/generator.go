package generator

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/khengari77/tang-init/profiles"
)

//go:embed templates/*
var templateFS embed.FS

func Generate(projName string, projType string, board profiles.BoardProfile) error {
	dirs := []string{
		filepath.Join(projName, "constraints"),
		filepath.Join(projName, "src"),
		filepath.Join(projName, "tb"),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	topTmpl := "templates/top.v.tmpl"
	if projType == "hdmi" {
		topTmpl = "templates/hdmi_top.v.tmpl"
	}

	files := []struct {
		tmpl string
		dest string
	}{
		{tmpl: "templates/Makefile.tmpl", dest: filepath.Join(projName, "Makefile")},
		{tmpl: "templates/board.cst.tmpl", dest: filepath.Join(projName, "constraints", "board.cst")},
		{tmpl: topTmpl, dest: filepath.Join(projName, "src", "top.v")},
		{tmpl: "templates/top_tb.v.tmpl", dest: filepath.Join(projName, "tb", "top_tb.v")},
		{tmpl: "templates/gitignore.tmpl", dest: filepath.Join(projName, ".gitignore")},
	}

	base := filepath.Base(projName)

	data := struct {
		profiles.BoardProfile
		ProjectName string
		ProjectType string
	}{
		BoardProfile: board,
		ProjectName:  base,
		ProjectType:  projType,
	}

	for _, f := range files {
		tmplBytes, err := fs.ReadFile(templateFS, f.tmpl)
		if err != nil {
			return err
		}

		tmpl, err := template.New(f.tmpl).Funcs(funcMap).Parse(string(tmplBytes))
		if err != nil {
			return err
		}

		outFile, err := os.Create(f.dest)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(outFile, data); err != nil {
			outFile.Close()
			return err
		}
		outFile.Close()
	}

	return nil
}
