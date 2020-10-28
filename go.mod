module gitee.com/grandeep/org-svc

go 1.13

require (
	gitee.com/grandeep/device-plugin v1.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-kit/kit v0.10.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/jinzhu/now v1.1.1
	github.com/lib/pq v1.8.0
	github.com/tealeg/xlsx v1.0.5
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20201006153459-a7d1128ccaa0
	google.golang.org/grpc v1.33.1
	gopkg.in/yaml.v2 v2.3.0
	gorm.io/driver/mysql v1.0.2
	gorm.io/gorm v1.20.2
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
