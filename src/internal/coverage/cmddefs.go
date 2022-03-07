// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package coverage

// CoverPkgConfig is a bundle of information passed from the Go
// command to the cover command during "go build -cover" runs. The
// Go command creates and fills in a struct as below, then passes
// file containing the encoded JSON for the struct to the "cover"
// tool when instrumenting the source files in a Go package.
type CoverPkgConfig struct {
	// File into which cmd/cover should emit summary info
	// when instrumentation is complete.
	OutConfig string

	// Import path for the package being instrumented.
	PkgPath string

	// Package name.
	PkgName string

	// Instrumentation granularity: one of "perfunc" or "perblock" (default)
	Granularity string

	// Module path for this package (empty if no go.mod in use)
	ModulePath string
}

// CoverFixupConfig contains annotations/notes generated by the
// cmd/cover tool (during instrumentation) to be passed on to the
// compiler when the instrumented code is compiled. The cmd/cover tool
// creates a struct of this type, JSON-encodes it, and emits the
// result to a file, which the Go command then passes to the compiler
// when the instrumented package is built.
type CoverFixupConfig struct {
	// Name of the variable (created by cmd/cover) containing the
	// encoded meta-data for the package.
	MetaVar string

	// Length of the meta-data.
	MetaLen int

	// Hash computed by cmd/cover of the meta-data.
	MetaHash string

	// Instrumentation strategy. For now this is always set to
	// "normal", but in the future we may add new values (for example,
	// if panic paths are instrumented, or if the instrumenter
	// eliminates redundant counters).
	Strategy string

	// Prefix assigned to the names of counter variables generated
	// during instrumentation by cmd/cover.
	CounterPrefix string

	// Name chosen for the package ID variable generated during
	// instrumentation.
	PkgIdVar string

	// Counter mode (e.g. set/count/atomic)
	CounterMode string

	// Counter granularity (perblock or perfunc).
	CounterGranularity string
}
