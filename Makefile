install ::
	@go mod vendor

example1 :: install
	@go run example1/main.go
