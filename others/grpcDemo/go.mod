module goProjects/daylearning/others/grpcDemo

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.36.0
