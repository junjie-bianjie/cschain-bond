module cschain-bond

go 1.14

require (
	github.com/bianjieai/irita-sdk-go v1.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.15
	github.com/aliyun/aliyun-oss-go-sdk v2.1.4+incompatible
	github.com/aws/aws-sdk-go v1.34.4
)

replace github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703
