module github.com/Hristiyan-Bonev/golang-course/sort/fulfillment

replace github.com/Hristiyan-Bonev/golang-course/sort/gen => ../gen

go 1.16

require (
	github.com/Hristiyan-Bonev/golang-course/sort/gen v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.38.0
)
