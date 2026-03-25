package setup

import (
	"context"
	_ "embed"

	"github.com/zitadel/zitadel/internal/database"
	"github.com/zitadel/zitadel/internal/eventstore"
)

var (
	//go:embed 70.sql
	addOIDCAppLimitAudience string
)

type Apps7OIDCConfigsLimitAudience struct {
	dbClient *database.DB
}

func (mig *Apps7OIDCConfigsLimitAudience) Execute(ctx context.Context, _ eventstore.Event) error {
	_, err := mig.dbClient.ExecContext(ctx, addOIDCAppLimitAudience)
	return err
}

func (mig *Apps7OIDCConfigsLimitAudience) String() string {
	return "70_apps7_oidc_configs_limit_audience"
}
