package embedgit

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
)

var (
	GitBinaryName string
)

// place the git binary in a temporary location
func init() {
	f, err := os.CreateTemp("", "embedgit")
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating git binary (%v)\n", err)
	}
	GitBinaryName = f.Name()
	_, err = f.Write(gitBinaryData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "writing git binary (%v)\n", err)
	}
	err = f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "closing git binary (%v)\n", err)
	}
	err = os.Chmod(GitBinaryName, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "chmod git binary (%v)\n", err)
	}
}

func RunGit(
	ctx context.Context,
	args ...string,
) (stdout string, stderr string, err error) {

	cmd := exec.Command(GitBinaryName, args...)
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err = cmd.Run()
	return stdoutBuf.String(), stderrBuf.String(), err
}
