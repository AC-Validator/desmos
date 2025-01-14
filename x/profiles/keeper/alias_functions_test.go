package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_IterateProfile() {
	date, err := time.Parse(time.RFC3339, "2010-10-02T12:10:00.000Z")
	suite.Require().NoError(err)

	addr1, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	addr2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.Require().NoError(err)

	addr3, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.Require().NoError(err)

	addr4, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	suite.Require().NoError(err)

	profiles := []*types.Profile{
		suite.CheckProfileNoError(types.NewProfile(
			"first",
			"",
			"",
			types.NewPictures("", ""),
			date,
			authtypes.NewBaseAccountWithAddress(addr1),
		)),
		suite.CheckProfileNoError(types.NewProfile(
			"second",
			"",
			"",
			types.NewPictures("", ""),
			date,
			authtypes.NewBaseAccountWithAddress(addr2),
		)),
		suite.CheckProfileNoError(types.NewProfile(
			"not",
			"",
			"",
			types.NewPictures("", ""),
			date,
			authtypes.NewBaseAccountWithAddress(addr3),
		)),
		suite.CheckProfileNoError(types.NewProfile(
			"third",
			"",
			"",
			types.NewPictures("", ""),
			date,
			authtypes.NewBaseAccountWithAddress(addr4),
		)),
	}

	expProfiles := []*types.Profile{
		profiles[0],
		profiles[1],
		profiles[3],
	}

	for _, profile := range profiles {
		err := suite.k.StoreProfile(suite.ctx, profile)
		suite.Require().NoError(err)
	}

	var validProfiles []*types.Profile
	suite.k.IterateProfiles(suite.ctx, func(_ int64, profile *types.Profile) (stop bool) {
		if profile.Dtag == "not" {
			return false
		}
		validProfiles = append(validProfiles, profile)
		return false
	})

	suite.Len(expProfiles, len(validProfiles))
	for _, profile := range validProfiles {
		suite.Contains(expProfiles, profile)
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetProfiles() {
	tests := []struct {
		name     string
		accounts []*types.Profile
	}{
		{
			name:     "Non empty Profiles list returned",
			accounts: []*types.Profile{suite.testData.profile},
		},
		{
			name:     "Profile not found",
			accounts: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, profile := range test.accounts {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			res := suite.k.GetProfiles(suite.ctx)
			suite.Require().Equal(test.accounts, res)
		})
	}
}
