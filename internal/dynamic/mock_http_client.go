// Code generated by mockery. DO NOT EDIT.

package dynamic

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// mockHTTPDoer is an autogenerated mock type for the HTTPDoer type
type mockHTTPDoer struct {
	mock.Mock
}

// Do provides a mock function with given fields: _a0
func (_m *mockHTTPDoer) Do(_a0 *http.Request) (*http.Response, error) {
	ret := _m.Called(_a0)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
