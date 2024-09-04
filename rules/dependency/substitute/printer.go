package substitute

import (
	"fmt"

	"github.com/dozer111/projectlinter-core/util/painter"
)

const lineDelimiter = "================================================================================================="

type substituteLibraryPrinter struct {
	libraries []Library
}

func (p *substituteLibraryPrinter) Print() []string {
	if len(p.libraries) == 0 {
		return nil
	}

	paint := painter.NewPainter()
	var output []string

	for idx, library := range p.libraries {
		if idx > 0 {
			output = append(output, paint.Yellow(lineDelimiter))
		}

		changeLibraryName := library.Name
		output = append(
			output,
			paint.Warning(
				fmt.Sprintf(
					"Dependency \"%s\" is deprecated.", changeLibraryName,
				),
			),
			paint.Yellow(
				fmt.Sprintf(
					"  Change it to \"%s\"", library.ChangeTo,
				),
			),
		)

		if len(library.Description) > 0 {
			output = append(output, "")
			for _, descriptionMessage := range library.Description {
				output = append(output, paint.Yellow(descriptionMessage))
			}
			output = append(output, "")
		}

		if len(library.ResponsiblePersons) > 0 {
			output = append(output, paint.Yellow("People, who can help u:"))
			for _, person := range library.ResponsiblePersons {
				output = append(output, paint.Yellow(fmt.Sprintf("\t- %s", person)))
			}
		}

		if len(library.Examples) > 0 {
			output = append(output, paint.Yellow("Examples:"))
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
