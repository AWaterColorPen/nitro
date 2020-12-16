package memory

import (
	"github.com/awatercolorpen/nitro/config/loader"
	"github.com/awatercolorpen/nitro/config/reader"
	"github.com/awatercolorpen/nitro/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
