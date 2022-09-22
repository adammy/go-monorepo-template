package create

// TmplCfg defines how to construct a file from a template.
type TmplCfg struct {
	// Name is the name of the template. Needs to be the exact name of the template file.
	Name string

	// Tmpl is the path template file to use.
	Tmpl string

	// Dir is the output directory.
	Dir string

	// File is the output filename.
	File string

	// Data is data for the template.
	Data interface{}
}

// Cfg defines some global configuration for the create package.
type Cfg struct {
	SuiteName string `yaml:"suiteName"`
	RootURL   string `yaml:"rootUrl"`
}
