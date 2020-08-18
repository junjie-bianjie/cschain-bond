module cschain-bond

go 1.14

require (
	github.com/aws/aws-sdk-go v1.34.4
	github.com/bianjieai/irita-sdk-go v1.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.15
	gitlab.bianjie.ai/csrb-bond/umbral-go v0.0.0-20200628091106-cbd01115d206
	go.uber.org/zap v1.13.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	google.golang.org/grpc v1.30.0
	gitlab.bianjie.ai/csrb-bond/pre-server v0.0.0-20200630110848-a89f05313740
)

replace github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703
