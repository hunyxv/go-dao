package mock

import "github.com/steenzout/go-dao"

// MockDAO interface for the mock data access object.
type MockDAO interface {
	// MockSomething does nothing.
	MockSomething()
}

// MockDAOImpl mock implementation of the MockDAO interface.
type MockDAOImpl struct {
	dao.BaseDataAccessObject
}

// MockSomething does nothing.
func (m *MockDAOImpl) MockSomething() {}

var _ MockDAO = (*MockDAOImpl)(nil)