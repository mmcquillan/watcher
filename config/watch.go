package config

type Watch struct {
	Server    string
	Type      string
	Input     string
	Filters   []string
	Output    []string
	Frequency int
}
