package dbgorm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"gorm.io/gorm/clause"

	"go.autokitteh.dev/autokitteh/backend/internal/db/dbgorm/scheme"
	"go.autokitteh.dev/autokitteh/internal/kittehs"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

func (db *gormdb) CreateIntegration(ctx context.Context, i sdktypes.Integration) error {
	err := db.db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(convertTypeToRecord(i)).Error
	if err != nil {
		return translateError(err)
	}
	return nil
}

func (db *gormdb) UpdateIntegration(ctx context.Context, i sdktypes.Integration) error {
	integ := convertTypeToRecord(i)
	err := db.db.WithContext(ctx).
		Where("integration_id = ?", integ.IntegrationID).
		Updates(integ).Error
	if err != nil {
		return translateError(err)
	}
	return nil
}

func convertTypeToRecord(i sdktypes.Integration) *scheme.Integration {
	l := sdktypes.GetIntegrationLogoURL(i)
	if l == nil {
		l = &url.URL{}
	}

	uls, err := json.Marshal(sdktypes.GetIntegrationUserLinks(i))
	if err != nil {
		uls = []byte{}
	}

	c := sdktypes.GetIntegrationConnectionURL(i)
	if c == nil {
		c = &url.URL{}
	}

	return &scheme.Integration{
		IntegrationID: sdktypes.GetIntegrationID(i).String(),
		UniqueName:    sdktypes.GetIntegrationUniqueName(i).String(),
		DisplayName:   sdktypes.GetIntegrationDisplayName(i),
		Description:   sdktypes.GetIntegrationDescription(i),
		LogoURL:       l.String(),
		UserLinks:     uls,
		// TODO: Tags
		// TODO(ENG-346): Connection UI specification instead of a URL.
		ConnectionURL: c.String(),
		// TODO: Functions
		// TODO: Events
		// TODO: APIKey
		// TODO: SigningKey
	}
}

func (db *gormdb) DeleteIntegration(ctx context.Context, id sdktypes.IntegrationID) error {
	// Desired product behavior: if user tries to delete an integration which
	// already has associated connections, AK should confirm with the user
	// what they want to do - abort, or cascade the deletion.
	// Note that a similar decision exists when deleting connections that
	// have active project mappings.
	var i scheme.Integration
	err := db.db.WithContext(ctx).
		Where("integration_id = ?", id.String()).
		Delete(&i).Error
	if err != nil {
		return translateError(err)
	}
	return nil
}

func (db *gormdb) GetIntegration(ctx context.Context, id sdktypes.IntegrationID) (sdktypes.Integration, error) {
	return get(db.db, ctx, scheme.ParseIntegration, "integration_id = ?", id.String())
}

func (db *gormdb) ListIntegrations(ctx context.Context, filter sdkservices.ListIntegrationsFilter) ([]sdktypes.Integration, error) {
	q := db.db.WithContext(ctx)

	if filter.NameSubstring != "" {
		s := fmt.Sprintf("%%%s%%", filter.NameSubstring)
		q = q.Where("uniqe_name LIKE ? OR display_name LIKE ?", s, s)
	}

	// TODO: Tags

	var is []scheme.Integration
	if err := q.Find(&is).Error; err != nil {
		return nil, translateError(err)
	}
	return kittehs.TransformError(is, scheme.ParseIntegration)
}
