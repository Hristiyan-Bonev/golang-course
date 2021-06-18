module github.com/Hristiyan-Bonev/golang-course/sort/fulfillment

replace github.com/Hristiyan-Bonev/golang-course/sort/gen => ../gen

go 1.16

require (
	github.com/Hristiyan-Bonev/golang-course/sort/gen v0.0.0-00010101000000-000000000000
	github.com/preslavmihaylov/ordertocubby v0.0.0-20210617074346-1704d311e402
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.38.0
)
