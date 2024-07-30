package model

type Domain struct {
	Name          string `csv:"name"`
	SourceVersion string `csv:"source_version"`
	TargetVersion string `csv:"target_version"`
}
