module github.com/jace-ys/super-smash-heroes/services/battle

go 1.12

replace (
	github.com/jace-ys/super-smash-heroes/api/proto/generated/go v0.0.0 => ../../api/proto/generated/go
	github.com/jace-ys/super-smash-heroes/libraries/go v0.0.0 => ../../libraries/go
)

require (
	github.com/jace-ys/super-smash-heroes/api/proto/generated/go v0.0.0
	github.com/jace-ys/super-smash-heroes/libraries/go v0.0.0
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980 // indirect
	golang.org/x/sys v0.0.0-20190616124812-15dcb6c0061f // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20190611190212-a7e196e89fd3 // indirect
	google.golang.org/grpc v1.21.1
)
