// Package run is a runtime library for microservices
package run

// Runtime is the top level interface representing a runtime
// capable managing the lifecycle of a microservice
type Runtime interface {
	// Fetch source from url
	Fetch(url string, opts ...FetchOption) (*Source, error)
	// Build the binary from source
	Build(*Source) (*Binary, error)
	// Execute a binary
	Exec(*Binary) (*Process, error)
	// Kill a process
	Kill(*Process) error
	// Wait for a process to exit
	Wait(*Process) error
}

// Source represents source code fetched from a URL
type Source struct {
	// URL fetched from
	URL string
	// Directory on disk
	Dir string
}

// Binary represents source which has been built
type Binary struct {
	// Path to binary
	Path string
	// Source from which it's build
	Source *Source
}

// Process represents a binary which has been executed
type Process struct {
	// The process ID
	ID string
	// Binary executed
	Binary *Binary
}
