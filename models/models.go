package models

type Movie struct {
	Label    string `string:"label"`
	Title    string `json:"title"`
	Released int64  `json:"released"`
	Tagline  string `json:"tagline"`
}

type Person struct {
	Name  string   `json:"name"`
	Born  int64    `json:"born,omitempty"`
	Job   string   `json:"job,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type Job struct {
	Name  string   `json:"name"`
	Movie string   `json:"movie"`
	Job   string   `json:"job,omitempty"`
	Roles []string `json:"roles,omitempty"`
}
