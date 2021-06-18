module github.com/Hristiyan-Bonev/golang-course/sort/sorting-service

go 1.16

replace github.com/Hristiyan-Bonev/golang-course/sort/gen => ../gen

require (
	github.com/Hristiyan-Bonev/golang-course/sort/gen v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.38.0
)
