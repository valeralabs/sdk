//go:build darwin

package build

//go:generate go run golang.org/x/mobile/cmd/gomobile bind --target=macos,ios -o ../VDK.xcframework ..
