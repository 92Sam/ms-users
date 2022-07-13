# go-ddd

Golang Domain Driven Design code - Using REST APP with DynamoDB


## Thoughts

It would be great if we can separate golang domains (user, comment) into individual folders, but there are some disadvantages:
- circular dependency. Say if we have a `/user` and `/comment` folder, and we want to add a method to get user comments for both folder. This would not work due to circular dependency. The solution is to put all the interfaces/entity into a single folder called `/domain`

## Should be separated?

It is probably better to put the controller in a separate folder, outside of the implementation. This is because the controller is dependent on the framework used, and we want to keep that layer separate from the implementation folder. Also, controller belongs to the `transport` category. There can be other ways of transport such as `grpc`, `graphql`, `http` etc that just calls the `service` layer.

## cmd folder

The applications (cli or server) are placed in the `cmd` folder. Each app is a composition of different use cases, and it depends on the usage. We can run an api server to allow clients to call the api, as well as running scripts that allows admins to interact with the application directly. N


## Expected workflow

When designing services, we do not want to concern ourselves with cross-cutting concerns such as database, logging, etc. The infrastructure should be pluggable, and the backend engineer should only focus on using the existing tools to add a new service.

Example, when creating a new service:
- do we need an entity? creates the structs from the database schema into models /folder
- do we need a repository? creates the repo and call the instances inside domain app.go
- do we need a service? first define the service interface in the domain. Then add the service layer with business logic
- do we need a controller? create a controller that calls the service and call him in routers inside trasnport/http
