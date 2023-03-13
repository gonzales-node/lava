//go:build ignore

package v0_x_0_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/app"
	keepertest "github.com/lavanet/lava/testutil/keeper"
	"github.com/stretchr/testify/suite"
)

type UpgradeTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app app.LavaApp
}

func (suite *UpgradeTestSuite) SetupTestApp() {
	suite.app, suite.ctx = app.TestSetup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (suite *UpgradeTestSuite) TestBody() {
	suite.SetupTestApp() // setup test app
	suite.T().Log("test")
}
