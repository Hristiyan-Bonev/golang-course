package main

import (
	"context"
	"fmt"
	"github.com/Hristiyan-Bonev/golang-course/sort/gen"
	"github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

type cubbyList map[string][]*gen.Item
type RobotState string

const (
	LoadingItemState    RobotState = "LOADING_ITEMS"
	RobotSortingState              = "SORTING"
	WaitingForItemState            = "WAITING_FOR_ITEMS"
	ItemSelectedState              = "ITEM_SELECTED"
	InitializationState            = "INITIALIZING"
)

type sortingService struct {
	State              RobotState
	Items              []*gen.Item
	Cubbies            cubbyList
	SelectedItem       *gen.Item
	tex                sync.Mutex
	fullfillmentClient gen.FulfillmentClient
}

func generateCubbies(cubbyIDs ...string) cubbyList {
	cubbies := make(cubbyList)

	for _, id := range cubbyIDs {
		cubbies[id] = []*gen.Item{}
	}

	return cubbies
}

func newSortingService() *sortingService {

	return &sortingService{
		Cubbies: generateCubbies("1", "2", "3", "4", "5", "6", "7", "8", "9", "10"),
		State:   InitializationState,
		Items:   []*gen.Item{},
	}
}

func (s *sortingService) LoadItems(context context.Context, loadRequest *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {

	if len(loadRequest.Items) == 0 {
		logrus.Warnf("Request with no items received! Ignoring...")
		return nil, fmt.Errorf("request contains no items")
	}

	for _, item := range loadRequest.Items {
		logrus.Warnf("Added item %s", item.String())
		s.Items = append(s.Items, item)
	}

	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) SelectItem(context.Context, *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {
	s.tex.Lock()
	defer s.tex.Unlock()

	if len(s.Items) < 1 {
		logrus.Errorf("Cannot select item! No items available!")
		return nil, fmt.Errorf("cannot select item because there are no items available")
	}

	s.SelectedItem = s.getRandomItem()

	randItem := &gen.SelectItemResponse{Item: s.SelectedItem}
	logrus.Infof("Selected item: %s", s.SelectedItem.Code)

	return randItem, nil
}

func (s *sortingService) MoveItem(context context.Context, request *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	s.tex.Lock()
	defer s.tex.Unlock()
	if s.SelectedItem == nil {
		logrus.Warnf("Unable to move item as no item is selected. Ignoring...")
		return nil, fmt.Errorf("no item selected")
	}

	for cubbyID, items := range s.Cubbies {
		if cubbyID == request.Cubby.Id {
			s.Cubbies[cubbyID] = append(items, s.SelectedItem)
			logrus.Infof("added item %+v to cubby %v", s.SelectedItem, cubbyID)
			s.SelectedItem = nil
			return &gen.MoveItemResponse{}, nil

		}
	}
	return nil, fmt.Errorf("unknown cubby ID:%v", request.Cubby.Id)
}

func (s *sortingService) getRandomItem() *gen.Item {
	rand.Seed(time.Now().Unix())
	randInt := rand.Intn(len(s.Items))
	randItem := s.Items[randInt]
	s.Items = append(s.Items[:randInt], s.Items[randInt+1:]...)
	return randItem
}
