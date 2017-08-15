package meta

import "github.com/keratin/authn-server/api"

func Routes(app *api.App) []*api.HandledRoute {
	return []*api.HandledRoute{
		api.Get("/health").
			SecuredWith(api.Unsecured()).
			Handle(getHealth(app)),
		api.Get("/jwks").
			SecuredWith(api.Unsecured()).
			Handle(getJWKs(app)),
		api.Get("/configuration").
			SecuredWith(api.Unsecured()).
			Handle(getConfiguration(app)),
	}
}