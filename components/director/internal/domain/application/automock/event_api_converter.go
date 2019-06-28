// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import "github.com/kyma-incubator/compass/components/director/pkg/graphql"
import "github.com/stretchr/testify/mock"
import "github.com/kyma-incubator/compass/components/director/internal/model"

// EventAPIConverter is an autogenerated mock type for the EventAPIConverter type
type EventAPIConverter struct {
	mock.Mock
}

// MultipleInputFromGraphQL provides a mock function with given fields: in
func (_m *EventAPIConverter) MultipleInputFromGraphQL(in []*graphql.EventAPIDefinitionInput) []*model.EventAPIDefinitionInput {
	ret := _m.Called(in)

	var r0 []*model.EventAPIDefinitionInput
	if rf, ok := ret.Get(0).(func([]*graphql.EventAPIDefinitionInput) []*model.EventAPIDefinitionInput); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.EventAPIDefinitionInput)
		}
	}

	return r0
}
