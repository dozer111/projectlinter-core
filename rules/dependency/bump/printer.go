package bump

import (
	"fmt"

	"github.com/dozer111/projectlinter-core/util/painter"
)

const lineDelimiter = "================================================================================================="

type bumpDependencyPrinter struct {
	libraries []*bumpDependencyConfig
}

type bumpDependencyConfig struct {
	Library
	CurrentVersion string
}

func (p *bumpDependencyPrinter) Print() []string {
	if len(p.libraries) == 0 {
		return nil
	}

	paint := painter.NewPainter()
	var output []string

	for idx, library := range p.libraries {
		if idx > 0 {
			output = append(output, paint.Yellow(lineDelimiter))
		}

		bumpName := library.Name
		bumpVersion := library.Version
		output = append(output,
			paint.Warning(fmt.Sprintf("%s", bumpName)),
			paint.Yellow(fmt.Sprintf("\tCurrent: %s", library.CurrentVersion)),
			paint.Green(fmt.Sprintf("\tBump to: %s", bumpVersion)),
		)

		if len(library.Description) > 0 {
			output = append(output, "")
			for _, descriptionMessage := range library.Description {
				output = append(output, paint.Yellow(descriptionMessage))
			}
			output = append(output, "")
		}

		if len(library.ResponsiblePersons) > 0 {
			output = append(output, paint.Yellow("\tPeople, who can help u:"))
			for _, person := range library.ResponsiblePersons {
				output = append(output, paint.Yellow(fmt.Sprintf("\t  - %s", person)))
			}
		}

		if len(library.Examples) > 0 {
			output = append(output, paint.Yellow("\tExamples:"))
			for _, example := range library.Examples {
				output = append(output, paint.Yellow(fmt.Sprintf("\t  %s: (%s)", example.ProjectName, example.Programmer)))

				if len(example.Description) > 0 {
					output = append(output, "")
					for _, descriptionMessage := range example.Description {
						output = append(output, paint.Yellow(fmt.Sprintf("\t    %s", descriptionMessage)))
					}
					output = append(output, "")
				}

				for _, exampleLink := range example.Links {
					output = append(output, paint.Yellow(fmt.Sprintf("\t    - %s", exampleLink)))
				}
			}
		}
	}

	return output
}
