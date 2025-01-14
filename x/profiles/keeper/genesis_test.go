package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_ExportGenesis() {
	usecases := []struct {
		name  string
		state struct {
			DTagRequests []types.DTagTransferRequest
			Params       types.Params
		}
		expGenesis *types.GenesisState
	}{
		{
			name: "empty state",
			state: struct {
				DTagRequests []types.DTagTransferRequest
				Params       types.Params
			}{
				DTagRequests: nil,
				Params:       types.DefaultParams(),
			},
			expGenesis: types.NewGenesisState(nil, types.DefaultParams()),
		},
		{
			name: "non-empty state",
			state: struct {
				DTagRequests []types.DTagTransferRequest
				Params       types.Params
			}{
				DTagRequests: []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				Params: types.NewParams(
					types.NewMonikerParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDtagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
			},
			expGenesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				types.NewParams(
					types.NewMonikerParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDtagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
			),
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, req := range uc.state.DTagRequests {
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(suite.ctx, req))
			}
			suite.k.SetParams(suite.ctx, uc.state.Params)

			exported := suite.k.ExportGenesis(suite.ctx)
			suite.Require().Equal(uc.expGenesis, exported)
		})
	}
}

func (suite *KeeperTestSuite) Test_InitGenesis() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.Require().NoError(err)

	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	addr3, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.Require().NoError(err)

	profile1, err := types.NewProfile(
		"dtag-1",
		"moniker-1",
		"bio-1",
		types.NewPictures("profile-1", "cover-1"),
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr1),
	)
	suite.Require().NoError(err)

	profile2, err := types.NewProfile(
		"dtag-2",
		"moniker-2",
		"bio-2",
		types.NewPictures("profile-2", "cover-2"),
		time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr2),
	)
	suite.Require().NoError(err)

	usecases := []struct {
		name         string
		authAccounts []authtypes.AccountI
		genesis      *types.GenesisState
		expState     struct {
			profiles             []*types.Profile
			dTagTransferRequests []types.DTagTransferRequest
			params               types.Params
		}
	}{
		{
			name:    "empty genesis",
			genesis: types.NewGenesisState(nil, types.DefaultParams()),
			expState: struct {
				profiles             []*types.Profile
				dTagTransferRequests []types.DTagTransferRequest
				params               types.Params
			}{
				profiles:             nil,
				dTagTransferRequests: nil,
				params:               types.DefaultParams(),
			},
		},
		{
			name: "non-empty genesis",
			authAccounts: []authtypes.AccountI{
				profile1,
				profile2,
				authtypes.NewBaseAccountWithAddress(addr3),
			},
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				types.NewParams(
					types.NewMonikerParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDtagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
			),
			expState: struct {
				profiles             []*types.Profile
				dTagTransferRequests []types.DTagTransferRequest
				params               types.Params
			}{
				profiles: []*types.Profile{
					profile1,
					profile2,
				},
				dTagTransferRequests: []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				params: types.NewParams(
					types.NewMonikerParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDtagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, acc := range uc.authAccounts {
				suite.ak.SetAccount(suite.ctx, acc)
			}
			suite.k.InitGenesis(suite.ctx, *uc.genesis)

			suite.Require().Equal(uc.expState.profiles, suite.k.GetProfiles(suite.ctx))
			suite.Require().Equal(uc.expState.dTagTransferRequests, suite.k.GetDTagTransferRequests(suite.ctx))
			suite.Require().Equal(uc.expState.params, suite.k.GetParams(suite.ctx))
		})
	}
}
