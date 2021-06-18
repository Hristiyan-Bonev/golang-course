package main

import (
	"fmt"
	"github.com/Hristiyan-Bonev/golang-course/sort/gen"
	"github.com/preslavmihaylov/ordertocubby"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"sync"
)

func NewFulfillmentService() gen.FulfillmentServer {
	return &fullfillmentService{}
}

type fullfillmentService struct {
	sortingService      gen.SortingRobotClient
	cubbyToOrderMapping map[string]*gen.Order
	mutex               sync.Mutex
}

func (f *fullfillmentService) prepareOrders(orders []*gen.Order) (map[string]*gen.Order, error) {

	f.mutex.Lock()
	defer f.mutex.Unlock()

	cubbyToOrderMap := make(map[string]*gen.Order)

	for _, order := range orders {
		retries := 1
		allocatedCubbyID := ""
		for retries < 15 {
			cubbyID := ordertocubby.Map(order.Id, uint32(retries), 10)
			if _, ok := cubbyToOrderMap[cubbyID]; !ok {
				allocatedCubbyID = cubbyID
				break
			}
			retries++
		}

		if allocatedCubbyID == "" {
			return nil, fmt.Errorf("unable to allocate cubby")
		}

		cubbyToOrderMap[allocatedCubbyID] = order

		logrus.Infof("Order %s allocated to Cubby %s", order.Id, allocatedCubbyID)
	}
	return cubbyToOrderMap, nil
}

func (f *fullfillmentService) LoadOrders(_ context.Context, req *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {

	orderMapping, err := f.prepareOrders(req.Orders)

	f.cubbyToOrderMapping = orderMapping

	if err != nil {
		return &gen.CompleteResponse{Status: "Error"}, fmt.Errorf("cannot map orders to cubbies. Reason: %s", err)
	}

	for {
		selectedItem, err := f.sortingService.SelectItem(context.Background(), &gen.SelectItemRequest{})
		if err != nil {
			logrus.Warnf("foo")
		}
			if selectedItem == nil {
				logrus.Warnf("nil item found. Skipping...")
				break
			}

			cubbyID := f.allocateToCubby(selectedItem.Item)

			if cubbyID == "" {
				return nil, fmt.Errorf("foo")
			}
			
			_, err := f.sortingService.MoveItem(context.Background(), &gen.MoveItemRequest{Cubby: &gen.Cubby{Id: cubbyID}})
			if err != nil {
				return nil, fmt.Errorf("cannot move item: %s", err)
			}
		}
	}


	preparedOrders := make([]*gen.PreparedOrder, len(req.Orders))

	for cubbyID, order := range orderMapping {
		preparedOrder := &gen.PreparedOrder{
			Order: order,
			Cubby: &gen.Cubby{Id: cubbyID},
		}
		preparedOrders = append(preparedOrders, preparedOrder)
	}

	return &gen.CompleteResponse{Status: "Done", Orders: preparedOrders}, nil
}

func (f *fullfillmentService) allocateToCubby(unallocatedItem *gen.Item) string {
	for cubby, order := range f.cubbyToOrderMapping {
		for _, item := range order.Items {
			if unallocatedItem.Code == item.Code {
				cubbyItems := f.cubbyToOrderMapping[cubby]
				itemIndex := find(cubbyItems, item.Code)
				order.Items = append(order.Items[:itemIndex], order.Items[itemIndex+1:]...)
				return cubby
			}
		}
	}
	return ""
}

func find(a *gen.Order, itemCode string) int {
	for i, n := range a.Items {
		if itemCode == n.Code {
			return i
		}
	}
	return len(a.Items)
}
