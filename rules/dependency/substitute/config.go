package substitute

type Library struct {
	Name               string    `yaml:"name"`
	ChangeTo           string    `yaml:"changeTo"`
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
