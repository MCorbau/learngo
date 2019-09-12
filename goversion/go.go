package goversion

import "runtime"

//Version returns the version of go
func Version() string {
	return runtime.Version()
}
