package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-test/client"
	"github.com/hashicorp/go-hclog"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:    "test",
		Version: "v0.0.11",
		Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, diag.Diagnostics) {
			return &client.TestClient{L: logger}, nil
		},
		ResourceMap: map[string]*schema.Table{
			"slow_resource": {
				Name: "slow_resource",
				Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
					meta.Logger().Info("fetching")
					select {
					case <-ctx.Done():
						return nil
					case <-time.After(time.Second * 5):
						return nil
					}
				},
				Columns: []schema.Column{
					{
						Name: "some_bool",
						Type: schema.TypeBool,
					},
				},
			},
			"very_slow_resource": {
				Name: "very_slow_resource",
				Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
					meta.Logger().Info("fetching very slow")
					select {
					case <-ctx.Done():
						return nil
					case <-time.After(time.Second * 8):
						return nil
					}
				},
			},
			"error_resource": {
				Name: "error_resource",
				Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
					return diag.WrapError(fmt.Errorf("error from provider"))
				},
			},
			"panic_resource": {
				Name: "panic_resource",
				Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
					panic("resource with panic")
				},
			},
			"migrate_resource": {
				Name: "migrate_resource",
				Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
					return nil
				},
				Columns: []schema.Column{
					{
						Name: "upgrade_column",
						Type: schema.TypeInt,
					},
					{
						Name: "upgrade_column_3",
						Type: schema.TypeInt,
					},
				},
			},
		},
		Config: func() provider.Config {
			return &client.Configuration{}
		},
		Logger: hclog.NewNullLogger(),
	}
}
