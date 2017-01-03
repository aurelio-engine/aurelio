package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunner_Run(t *testing.T) {
	t1 := new(runnableMock)
	t1.On("Run").Return()

	t2 := new(runnableMock)
	t2.On("Run").Return()

	t3 := new(runnableMock)
	t3.On("Run").Return()

	runner := NewRunner(1)
	runner.tasks = []Runnable{t1, t2, t3}
	runner.Run()

	t1.AssertCalled(t, "Run")
	t2.AssertCalled(t, "Run")
	t3.AssertCalled(t, "Run")

	assert.Equal(t, 0, len(runner.tasks))
}

func TestRunner_AddTask(t *testing.T) {
	runner := NewRunner(1)

	assert.Equal(t, 0, len(runner.tasks))

	t1 := new(runnableMock)
	runner.AddTask(t1)

	assert.Equal(t, 1, len(runner.tasks))
	assert.Equal(t, t1, runner.tasks[0])
}
