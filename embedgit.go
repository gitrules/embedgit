package embedgit

import (
	"fmt"
	"os"
)

// place the git binary in a temporary location
func init() {
	f, err := os.CreateTemp("", "embedgit")
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating git binary (%v)\n", err)
	}
	name := f.Name()
	_, err = f.Write(GitBinary)
	if err != nil {
		fmt.Fprintf(os.Stderr, "writing git binary (%v)\n", err)
	}
	err = f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "closing git binary (%v)\n", err)
	}
	err = os.Chmod(name, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "chmod git binary (%v)\n", err)
	}
}
