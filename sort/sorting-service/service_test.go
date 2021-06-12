package main

import (
	"context"
	"github.com/Hristiyan-Bonev/golang-course/sort/gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadItems(t *testing.T) {

	ctx := context.Background()

	testCases := []struct {
		name        string
		in          []*gen.Item
		expected    []*gen.Item
		errorString string
	}{
		{
			name: "happy path case",
			in: []*gen.Item{
				{Code: "123-456-789", Label: "foo shirt"},
				{Code: "234-123-123", Label: "foo pants"},
				{Code: "334-123-123", Label: "foo jacket"},
				{Code: "123-410-213", Label: "foo hat"},
			},
			errorString: "",
		},
		{
			name:        "bad case",
			in:          []*gen.Item{},
			errorString: "request contains no items",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			service := newSortingService()
			_, err := service.LoadItems(ctx, &gen.LoadItemsRequest{Items: tc.in})

			if err != nil {
				if tc.errorString != "" {
					assert.Equal(t, tc.errorString, err.Error())
				} else {
					t.Fatalf("unknown error received: %v", err)
				}
			}
			assert.EqualValues(t, service.Items, tc.in)
		})
	}
}

func TestSelectItem(t *testing.T) {

	testCases := []struct {
		testName     string
		initialItems []*gen.Item
		out          *gen.Item
	}{
		{
			testName: "pick with single item in inventory",
			initialItems: []*gen.Item{
				{Code: "123-456", Label: "foo shirt"},
				{Code: "999-888", Label: "foo shirt"},
			},
			out: &gen.Item{Code: "123-456", Label: "foo shirt"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			service := newSortingService()
			service.Items = tc.initialItems

			_, err := service.SelectItem(context.Background(), &gen.SelectItemRequest{})

			if err != nil {
				t.Fatalf("unable to select item: %v", err)
			}

			// Assert one item is removed
			assert.Equal(t, len(tc.initialItems)-1, len(service.Items))
		})
	}
}
