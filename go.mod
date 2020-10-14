module github.com/mmeloni/lcvsimmu.git

go 1.15

require (
	github.com/codenotary/immudb v0.8.0
	github.com/vchain-us/ledger-compliance-go v0.0.0-20201014151406-5654aa5886dc
	google.golang.org/grpc v1.29.1
)

replace github.com/codenotary/immudb v0.8.0 => github.com/codenotary/immudb v0.0.0-20201014151548-74e7b6c86339
