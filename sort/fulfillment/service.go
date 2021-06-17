package main

import (
	"errors"
	"fmt"
	"github.com/Hristiyan-Bonev/golang-course/sort/gen"
)

func NewFulfillmentService() gen.FulfillmentServer {
	return fullfillmentService{}
}

type fullfillmentService struct{}

func (f fullfillmentService) LoadOrders(context.Context, *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	return nil, errors.New("Not implemented")
}
