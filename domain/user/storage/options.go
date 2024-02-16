package storage

type mode string

const (
	Information mode = "Information"
)

type Options struct {
	modes        []mode
	AllDevice    bool
	ActiveDevice bool
}
