package script

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunWithYq(t *testing.T) {
	_, err := exec.LookPath("yq")
	assert.NoError(t, err, "Binary 'yq' must exist")

	output := Var()
	err = Exec("yq", "read", "fixtures/test.yaml", "spec.state").To(output)
	assert.NoError(t, err, "Should not have any error")

	assert.Equal(t, "absent", output.String())
}

func TestEcho(t *testing.T) {
	output := Var()
	err := Echo("hello world").To(output)
	assert.NoError(t, err, "Should not have any error")

	assert.Equal(t, "hello world\n", output.RawString())
	assert.Equal(t, "hello world", output.String())
}


func TestEchoToStdout(t *testing.T) {
	err := Echo("hello world").To(Stdout())
	assert.NoError(t, err, "Should not have any error")
}
