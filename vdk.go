// Provides a gomobile-safe wrapper.
package VDK

// wrap enumerables with stringer
//go:generate go generate ./address
//go:generate go generate ./constant
//go:generate go generate ./transaction
//go:generate go generate ./encoding/c32
//go:generate go generate ./encoding/clarity

// package VDK with gomobile
//go:generate go run golang.org/x/mobile/cmd/gomobile init .
//go:generate go generate ./build
