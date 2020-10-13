module github.com/mmeloni/lcvsimmu.git

go 1.15

require (
	github.com/codenotary/immudb v0.8.0
	github.com/vchain-us/ledger-compliance-go v0.0.0-20201012160855-de9cc58c77da
	google.golang.org/grpc v1.29.1
)

replace github.com/codenotary/immudb v0.8.0 => github.com/codenotary/immudb v0.0.0-20201009135111-4ecdf1bf01bd
