package webhook_test

import (
	"testing"

	"github.com/kyma-incubator/compass/components/director/internal/domain/webhook"
	"github.com/kyma-incubator/compass/components/director/internal/domain/webhook/automock"
	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/pkg/graphql"
	"github.com/stretchr/testify/assert"
)

func TestConverter_ToGraphQL(t *testing.T) {
	// given
	testCases := []struct {
		Name     string
		Input    *model.ApplicationWebhook
		Expected *graphql.ApplicationWebhook
	}{
		{
			Name:     "All properties given",
			Input:    fixModelWebhook("1", "foo", "bar"),
			Expected: fixGQLWebhook("foo", "bar"),
		},
		{
			Name:     "Empty",
			Input:    &model.ApplicationWebhook{},
			Expected: &graphql.ApplicationWebhook{},
		},
		{
			Name:     "Nil",
			Input:    nil,
			Expected: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			authConv := &automock.AuthConverter{}
			if testCase.Input != nil {
				authConv.On("ToGraphQL", testCase.Input.Auth).Return(testCase.Expected.Auth)
			}
			converter := webhook.NewConverter(authConv)

			// when
			res := converter.ToGraphQL(testCase.Input)

			// then
			assert.Equal(t, testCase.Expected, res)
			authConv.AssertExpectations(t)
		})
	}
}

func TestConverter_MultipleToGraphQL(t *testing.T) {
	// given
	input := []*model.ApplicationWebhook{
		fixModelWebhook("1", "foo", "baz"),
		fixModelWebhook("2", "bar", "bez"),
		{},
		nil,
	}
	expected := []*graphql.ApplicationWebhook{
		fixGQLWebhook("foo", "baz"),
		fixGQLWebhook("bar", "bez"),
		{},
	}
	authConv := &automock.AuthConverter{}
	authConv.On("ToGraphQL", input[0].Auth).Return(expected[0].Auth)
	authConv.On("ToGraphQL", (*model.Auth)(nil)).Return(nil)
	converter := webhook.NewConverter(authConv)

	// when
	res := converter.MultipleToGraphQL(input)

	// then
	assert.Equal(t, expected, res)
	authConv.AssertExpectations(t)
}

func TestConverter_InputFromGraphQL(t *testing.T) {
	// given
	testCases := []struct {
		Name     string
		Input    *graphql.ApplicationWebhookInput
		Expected *model.ApplicationWebhookInput
	}{
		{
			Name:     "All properties given",
			Input:    fixGQLWebhookInput("foo"),
			Expected: fixModelWebhookInput("foo"),
		},
		{
			Name:     "Empty",
			Input:    &graphql.ApplicationWebhookInput{},
			Expected: &model.ApplicationWebhookInput{},
		},
		{
			Name:     "Nil",
			Input:    nil,
			Expected: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			authConv := &automock.AuthConverter{}
			if testCase.Input != nil {
				authConv.On("InputFromGraphQL", testCase.Input.Auth).Return(testCase.Expected.Auth)
			}
			converter := webhook.NewConverter(authConv)

			// when
			res := converter.InputFromGraphQL(testCase.Input)

			// then
			assert.Equal(t, testCase.Expected, res)
			authConv.AssertExpectations(t)
		})
	}
}

func TestConverter_MultipleInputFromGraphQL(t *testing.T) {
	// given
	input := []*graphql.ApplicationWebhookInput{
		fixGQLWebhookInput("foo"),
		fixGQLWebhookInput("bar"),
		{},
		nil,
	}
	expected := []*model.ApplicationWebhookInput{
		fixModelWebhookInput("foo"),
		fixModelWebhookInput("bar"),
		{},
	}
	authConv := &automock.AuthConverter{}
	authConv.On("InputFromGraphQL", input[0].Auth).Return(expected[0].Auth)
	authConv.On("InputFromGraphQL", (*graphql.AuthInput)(nil)).Return(nil)
	converter := webhook.NewConverter(authConv)

	// when
	res := converter.MultipleInputFromGraphQL(input)

	// then
	assert.Equal(t, expected, res)
	authConv.AssertExpectations(t)
}
