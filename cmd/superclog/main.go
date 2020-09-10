package main

import (
	"flag"
	"fmt"
	"github.com/lwaddicor/superclog"
	"log"
	"os"
)

func main() {
	var pathFlag = flag.String("path", "", "Git Path (File only)")
	var fromFlag = flag.String("from", "", "From refspec")
	var toFlag = flag.String("to", "", "To refspec")
	var tmplFlag = flag.String("tmpl", superclog.TemplateInternalQARelease, "Inbuilt template (InternalQARelease, ExternalRelease)")
	var outFlag = flag.String("out", "", "File to output to. Default when empty is stdout")
	flag.Parse()

	if *pathFlag == "" {
		*pathFlag = "."
	}

	if *fromFlag == "" {
		log.Fatal("no 'from' flag set")
	}

	if *fromFlag == "" {
		log.Fatal("no 'to' flag set")
	}

	w := os.Stdout
	if *outFlag != "" {
		var err error
		w, err = os.Create(*outFlag)
		if err != nil {
			log.Fatal(fmt.Errorf("open output file: %w", err))
		}
		defer w.Close()
	}

	commits, err := superclog.GetCommits(*pathFlag, *fromFlag, *toFlag)
	if err != nil {
		log.Fatal(err)
	}

	categories := superclog.CalculateCategories(commits)

	template, err := superclog.DefaultTemplate(*tmplFlag)
	if err != nil {
		log.Fatal(err)
	}

	err = superclog.TemplateReleaseNotes(w, template, commits, categories)
	if err != nil {
		log.Fatal(err)
	}
}
