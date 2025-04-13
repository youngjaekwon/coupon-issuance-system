package couponcode

import "github.com/stretchr/testify/mock"

type MockCodeGenerator struct {
	mock.Mock
}

func (m *MockCodeGenerator) Generate() string {
	args := m.Called()
	return args.String(0)
}
