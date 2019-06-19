module github.com/jace-ys/super-smash-heroes/services/rest-api

go 1.12

replace (
	github.com/jace-ys/super-smash-heroes/api/proto/generated/go v0.0.0 => ../../api/proto/generated/go
	github.com/jace-ys/super-smash-heroes/libraries/go v0.0.0 => ../../libraries/go
)

require (
	github.com/gorilla/mux v1.7.2
	github.com/jace-ys/super-smash-heroes/api/proto/generated/go v0.0.0
	github.com/jace-ys/super-smash-heroes/libraries/go v0.0.0
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/net v0.0.0-20190619014844-b5b0513f8c1b // indirect
	golang.org/x/sys v0.0.0-20190619223125-e40ef342dc56 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20190611190212-a7e196e89fd3 // indirect
	google.golang.org/grpc v1.21.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
