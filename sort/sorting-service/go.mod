module github.com/Hristiyan-Bonev/golang-course/sort/sorting-service

go 1.16

replace github.com/Hristiyan-Bonev/golang-course/sort/gen => ../gen

require (
	github.com/Hristiyan-Bonev/golang-course/sort/gen v0.0.0-00010101000000-000000000000
	github.com/fullstorydev/grpcurl v1.8.1 // indirect
	github.com/sirupsen/logrus v1.2.0
	github.com/stretchr/testify v1.5.1
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
	google.golang.org/genproto v0.0.0-20210608205507-b6d2f5bf0d7d // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0 // indirect
)
