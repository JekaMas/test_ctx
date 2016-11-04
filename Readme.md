Test project shows how we can use context without passing it as a 1st parameter in service/dao layers.

Description: 
* 2 services: product and image call each other
* product service has unit test

Logic of code:
* To get service - call function in package "service": service.YourServiceName(ctx)
* This will create new instance of service
* Services are not global across application, but creating on place where you need them

Unit tests:
* product service contains one
* `ctx := system.SetupServices(context.Background())` - add to context factory of services, same code you can find in main.go
* generator must prepare methods for mocks creation: `mockB := service.MockImage(ctx, mockCtrl)`
