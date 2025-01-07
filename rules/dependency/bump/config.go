package bump

type Library struct {
	Name               string    `yaml:"name"`
	Version            string    `yaml:"version"`
	Description        []string  `yaml:"description,omitempty"`
	ResponsiblePersons []string  `yaml:"responsiblePersons,omitempty"`
	Examples           []Example `yaml:"examples,omitempty"`
}

type Example struct {
	ProjectName string   `yaml:"projectName"`
	Programmer  string   `yaml:"committee"`
	Description []string `yaml:"description,omitempty"`
	Links       []string `yaml:"links"`
}
