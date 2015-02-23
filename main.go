package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

const tpl = `package {{.Package}}

import "github.com/gchaincl/gotic/fs"

func init() {
{{range $file, $data := .Files}}
	fs.Add("{{$file}}", {{$data}})
{{end}}
}
`

type gotic struct {
	Package string
	Files   map[string]string
}

func main() {
	g := &gotic{}

	flag.StringVar(&g.Package, "package", "main", "package name")
	flag.Parse()

	t := template.Must(template.New(g.Package + "_gotic").Parse(tpl))
	g.Files = make(map[string]string)

	for _, pattern := range flag.Args() {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			panic(err)
		}

		for _, file := range matches {
			abs, err := filepath.Abs(file)
			if err != nil {
				panic(err)
			}

			data, err := ioutil.ReadFile(abs)
			if err != nil {
				panic(err)
			}

			g.Files[abs] = fmt.Sprintf("%q", data)
		}
	}

	err := t.Execute(os.Stdout, g)
	if err != nil {
		println("Error", err)
	}

}