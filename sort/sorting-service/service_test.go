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
		testName      string
		initialItems  []*gen.Item
		out           *gen.Item
		expectedError bool
	}{
		{
			testName: "pick with single item in inventory",
			initialItems: []*gen.Item{
				{Code: "123-456", Label: "foo shirt"},
				{Code: "999-888", Label: "foo shirt"},
			},
			out: &gen.Item{Code: "123-456", Label: "foo shirt"},
		},
		{
			testName:      "pick with single item in inventory",
			initialItems:  []*gen.Item{},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			service := newSortingService()
			service.Items = tc.initialItems

			_, err := service.SelectItem(context.Background(), &gen.SelectItemRequest{})

			if err != nil && !tc.expectedError {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.expectedError {
				// Ensure one item is discarded upon SelectItem
				assert.Equal(t, len(tc.initialItems)-1, len(service.Items))
			}
		})
	}
}

func TestMoveItem(t *testing.T) {

	service := newSortingService()

	service.Items = []*gen.Item{
		{Code: "123-456", Label: "foo shirt"},
		{Code: "999-888", Label: "foo shirt"},
	}

	// First select and move
	if _, err := service.SelectItem(context.Background(), &gen.SelectItemRequest{}); err != nil {
		t.Fatalf("cannot select item: %v", err)
	}
	if _, err := service.MoveItem(context.Background(), &gen.MoveItemRequest{Cubby: &gen.Cubby{Id: "A1"}}); err != nil {
		t.Fatalf("cannot move item: %v", err)
	}

	// Second select and move
	if _, err := service.SelectItem(context.Background(), &gen.SelectItemRequest{}); err != nil {
		t.Fatalf("cannot select item: %v", err)
	}

	if _, err := service.MoveItem(context.Background(), &gen.MoveItemRequest{Cubby: &gen.Cubby{Id: "A2"}}); err != nil {
		t.Fatalf("cannot move item: %v", err)
	}

	// Verify that items are placed in the appropriate cubbies
	assert.Equal(t, len(service.Cubbies["A1"]), 1)
	assert.Equal(t, len(service.Cubbies["A2"]), 1)

	// Third select and move - should fail
	_, selectErr := service.SelectItem(context.Background(), &gen.SelectItemRequest{})
	_, moveErr := service.MoveItem(context.Background(), &gen.MoveItemRequest{Cubby: &gen.Cubby{Id: "ZZ"}})

	assert.NotEqual(t, nil, selectErr)
	assert.NotEqual(t, nil, moveErr)

}
