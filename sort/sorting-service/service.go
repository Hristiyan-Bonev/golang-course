package main

import (
	"context"
	"fmt"
	"github.com/Hristiyan-Bonev/golang-course/sort/gen"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func newSortingService() *sortingService {
	cubbies := make(map[string][]*gen.Item)

	return &sortingService{
		Cubbies: cubbies,
		State:   InitializationState,
		Items:   []*gen.Item{},
	}
}

type RobotState string

const (
	LoadingItemState    RobotState = "LOADING_ITEMS"
	RobotSortingState              = "SORTING"
	WaitingForItemState            = "WAITING_FOR_ITEMS"
	ItemSelectedState              = "ITEM_SELECTED"
	InitializationState            = "INITIALIZING"
)

type sortingService struct {
	State        RobotState
	Items        []*gen.Item
	Cubbies      map[string][]*gen.Item
	SelectedItem *gen.Item
}

func (s *sortingService) LoadItems(context context.Context, loadRequest *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {

	if len(loadRequest.Items) == 0 {
		return nil, fmt.Errorf("request contains no items")
	}

	for _, item := range loadRequest.Items {
		s.Items = append(s.Items, item)
	}
	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) MoveItem(context context.Context, request *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	if s.SelectedItem == nil {
		return nil, fmt.Errorf("no item selected")
	}

	for cubbyID, items := range s.Cubbies {
		s.Cubbies[cubbyID] = append(items, s.SelectedItem)
	}

	s.SelectedItem = nil

	return &gen.MoveItemResponse{}, nil
}

func (s *sortingService) SelectItem(context.Context, *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {

	if len(s.Items) < 1 {
		logrus.Errorf("Cannot select item! No items available!")
		return nil, fmt.Errorf("cannot select item because there are no items available")
	}

	//if s.SelectedItem {
	//	return nil, fmt.Errorf("item al")
	//}

	s.SelectedItem = s.getRandomItem()

	randItem := &gen.SelectItemResponse{Item: s.SelectedItem}

	logrus.Infof("Selected item: %s", randItem.Item.Code)

	return randItem, nil
}

func (s *sortingService) getRandomItem() *gen.Item {
	rand.Seed(time.Now().Unix())
	randInt := rand.Intn(len(s.Items))
	randItem := s.Items[randInt]

	s.Items = append(s.Items[:randInt], s.Items[randInt+1:]...)

	return randItem
}
