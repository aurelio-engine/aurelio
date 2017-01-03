package task

import "github.com/stretchr/testify/mock"

type runnableMock struct {
	mock.Mock
}

func (rm *runnableMock) Run() {
	rm.Called()
}
