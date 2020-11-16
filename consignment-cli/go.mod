module github.com/AndreyShep2012/EVService/consignment-cli

go 1.15

replace github.com/AndreyShep2012/EVService/shippingService => ../ShippingService

require (
	github.com/AndreyShep2012/EVService/shippingService v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.33.2
)
