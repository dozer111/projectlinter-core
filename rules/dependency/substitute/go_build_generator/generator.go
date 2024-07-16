//go:build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/dozer111/projectlinter-core/rules/dependency/substitute"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/1set/gut/yos"
)

func main() {
	flagCalledFrom := flag.String("c", "", "called from")
	flag.Parse()

	if *flagCalledFrom == "" {
		log.Fatalf("calledFrom flag is required.")
	}

	path := strings.Split(*flagCalledFrom, "/")
	packageName := path[len(path)-1]

	path[len(path)-1] = "shared_source"
	path = append(path, "substitute_library")

	pathToMain := *flagCalledFrom
	pathToConfigs := strings.Join(path, "/")

	if !yos.ExistDir(pathToConfigs) {
		log.Fatalf("directory %s does not exists", pathToConfigs)
	}

	parser := substitute.NewParser(pathToConfigs)
	configs, err := parser.Parse()

	if err != nil {
		log.Fatalf("cannot parse substitute configs: %v", err)
	}

	pathToFile := fmt.Sprintf("%s/generated_substitute.go", pathToMain)
	f, err := os.Create(pathToFile)
	if err != nil {
		log.Fatalf("cannot create file %s: %v", pathToFile, err)
	}
	defer f.Close()

	_, err = fmt.Fprintf(
		f,
		headTpl,
		packageName,
		importTpl,
	)
	if err != nil {
		panic(err)
	}

	for _, bumpConfig := range configs {
		if _, err := fmt.Fprintf(f, "%s", printConfig(bumpConfig)); err != nil {
			panic(err)
		}
	}

	if _, err := fmt.Fprintln(f, "}"); err != nil {
		panic(err)
	}

	runCommand("go", "fmt", pathToFile)
}

func printConfig(cfg substitute.Library) string {
	tpl := make([]string, 0, 100)

	tpl = []string{"{"}

	tpl = append(
		tpl,
		fmt.Sprintf(`Name: "%s",`, cfg.Name),
		fmt.Sprintf(`ChangeTo: "%s",`, cfg.ChangeTo),
	)

	if len(cfg.Description) > 0 {
		tpl = append(tpl, "Description: []string{")
		for _, rp := range cfg.Description {
			tpl = append(tpl, fmt.Sprintf(`%q,`, rp))
		}
		tpl = append(tpl, "		},")
	}

	if len(cfg.ResponsiblePersons) != 0 {
		tpl = append(tpl, "ResponsiblePersons: []string{")
		for _, rp := range cfg.ResponsiblePersons {
			tpl = append(tpl, fmt.Sprintf(`"%s",`, rp))
		}
		tpl = append(tpl, "},")
	}

	if len(cfg.Examples) != 0 {
		tpl = append(tpl, "Examples: []substitute.Example{")

		for _, ex := range cfg.Examples {
			tpl = append(tpl, "			{")
			tpl = append(
				tpl,
				fmt.Sprintf(`ProjectName: "%s",`, ex.ProjectName),
				fmt.Sprintf(`Programmer: "%s",`, ex.Programmer),
			)

			tpl = append(tpl, "Links: []string{")
			for _, exLink := range ex.Links {
				tpl = append(tpl, fmt.Sprintf(`"%s",`, exLink))
			}
			tpl = append(tpl, "},")

			tpl = append(tpl, "},")
		}

		tpl = append(tpl, "},")
	}

	tpl = append(tpl, "},\n")

	return strings.Join(tpl, "\n")
}

var headTpl = `// Code generated by go generate; DO NOT EDIT.
// This file was generated
// 

package %s
%s
var substituteLibraryConfigs = []substitute.Library{
`
var importTpl = `
import (
	"github.com/dozer111/projectlinter-core/rules/dependency/substitute"
)
`

func runCommand(cmd string, args ...string) ([]byte, error) {
	c := exec.Command(cmd, args...)

	b := &bytes.Buffer{}
	errb := &bytes.Buffer{}
	c.Stdout = b
	c.Stderr = errb
	err := c.Run()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, errb)
	}

	return b.Bytes(), nil
}
