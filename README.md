# merchant-account-management

1. Run test: go test github.com/tikivn/ops-delivery/saga/aftersales/fd-inbound -coverprofile=coverage.out TestMain
2. Viewing the results:
   a) in console: go tool cover -func=coverage.out
   b) in UI: go tool cover -html=coverage.out

mockery --name=Stringer