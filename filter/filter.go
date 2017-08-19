package filter

import "github.com/kwoodhouse93/audio-playground/source"

// Pass  returns a filter that doesn't modify the source
func Pass(source source.Source) source.Source {
	return source
}
