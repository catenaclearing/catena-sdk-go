module github.com/catenaclearing/catena-sdk-go

go 1.24.0

toolchain go1.24.11

require (
	github.com/catenaclearing/catena-sdk-go/gen/authentication v0.0.0-00010101000000-000000000000
	github.com/catenaclearing/catena-sdk-go/gen/integrations v0.0.0-00010101000000-000000000000
	github.com/catenaclearing/catena-sdk-go/gen/notifications v0.0.0-00010101000000-000000000000
	github.com/catenaclearing/catena-sdk-go/gen/orgs v0.0.0-00010101000000-000000000000
	github.com/catenaclearing/catena-sdk-go/gen/telematics v0.0.0-00010101000000-000000000000
	golang.org/x/sync v0.19.0
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace (
	github.com/catenaclearing/catena-sdk-go/gen/authentication => ./gen/authentication
	github.com/catenaclearing/catena-sdk-go/gen/integrations => ./gen/integrations
	github.com/catenaclearing/catena-sdk-go/gen/notifications => ./gen/notifications
	github.com/catenaclearing/catena-sdk-go/gen/orgs => ./gen/orgs
	github.com/catenaclearing/catena-sdk-go/gen/telematics => ./gen/telematics
)
