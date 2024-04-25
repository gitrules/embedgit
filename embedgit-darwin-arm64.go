//go:build darwin && arm64

package embedgit

import (
	_ "embed"
)

var (
	//go:embed dist/bingit-2.44.0-darwin-arm64
	gitBinaryData []byte
)
