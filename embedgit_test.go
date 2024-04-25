package embedgit

import (
	"context"
	"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	gout, gerr, err := RunGit(context.Background(), "version")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("out:", gout)
	fmt.Println("err:", gerr)
}
