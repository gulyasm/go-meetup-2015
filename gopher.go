package main

import "io"

// Gopher represents a cute little gopher
type Gopher struct {
	Name string
}

// Sometimes we want to write out our little gopher.
// easy to test
func (g *Gopher) WriteTo(w *io.Writer) (int64, err) {

}
