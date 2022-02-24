// Code generated by mockery v2.9.4. DO NOT EDIT.

package monitoring

import mock "github.com/stretchr/testify/mock"

// SourceFactoryMock is an autogenerated mock type for the SourceFactory type
type SourceFactoryMock struct {
	mock.Mock
}

// NewSource provides a mock function with given fields: chainConfig, feedConfig
func (_m *SourceFactoryMock) NewSource(chainConfig ChainConfig, feedConfig FeedConfig) (Source, error) {
	ret := _m.Called(chainConfig, feedConfig)

	var r0 Source
	if rf, ok := ret.Get(0).(func(ChainConfig, FeedConfig) Source); ok {
		r0 = rf(chainConfig, feedConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Source)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ChainConfig, FeedConfig) error); ok {
		r1 = rf(chainConfig, feedConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
