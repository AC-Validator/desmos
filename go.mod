module github.com/desmos-labs/desmos

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/desmos-labs/Go-Emoji-Utils v1.1.1-0.20200515063516-9c493b11de3e
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => github.com/RiccardoM/cosmos-sdk v0.40.2-0.20210413061315-e274652f0ed8
