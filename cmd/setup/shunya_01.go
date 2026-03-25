package setup

import (
	"context"
	_ "embed"

	"github.com/zitadel/zitadel/internal/database"
	"github.com/zitadel/zitadel/internal/eventstore"
)

//go:embed shunya_01.sql
var addOIDCAppLimitAudience string

type AppsOIDCConfigsLimitAudience struct {
	dbClient *database.DB
}

func (mig *AppsOIDCConfigsLimitAudience) Execute(ctx context.Context, _ eventstore.Event) error {
	_, err := mig.dbClient.ExecContext(ctx, addOIDCAppLimitAudience)
	return err
}

func (mig *AppsOIDCConfigsLimitAudience) String() string {
	return "70_apps7_oidc_configs_limit_audience"
}
