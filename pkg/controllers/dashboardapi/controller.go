package dashboardapi

import (
	"context"

	"github.com/rancher/rancher/pkg/controllers/dashboardapi/feature"
	"github.com/rancher/rancher/pkg/controllers/dashboardapi/mcmstart"
	"github.com/rancher/rancher/pkg/controllers/dashboardapi/settings"
	"github.com/rancher/rancher/pkg/wrangler"
)

func Register(ctx context.Context, wrangler *wrangler.Context) error {
	mcmstart.Register(ctx, wrangler.Mgmt.Feature(), wrangler.MultiClusterManager)
	feature.Register(ctx, wrangler.Mgmt.Feature())
	return settings.Register(wrangler.Mgmt.Setting())
}
