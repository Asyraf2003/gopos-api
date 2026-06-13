// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

// audit:allow-oversize reason=bootstrap-wiring
package bootstrap

import (
	"context"
	"time"

	"pos-go/internal/config"
	authhttp "pos-go/internal/modules/auth/transport/http"
	authusecase "pos-go/internal/modules/auth/usecase"
	capabilityhttp "pos-go/internal/modules/capability/transport/http"
	capabilityusecase "pos-go/internal/modules/capability/usecase"
	productcataloghttp "pos-go/internal/modules/productcatalog/transport/http"
	productcatalogusecase "pos-go/internal/modules/productcatalog/usecase"
	servicecatalogdomain "pos-go/internal/modules/servicecatalog/domain"
	servicecataloghttp "pos-go/internal/modules/servicecatalog/transport/http"
	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"
	systemhttp "pos-go/internal/modules/system/transport/http"
	googleoidc "pos-go/internal/platform/google"
	"pos-go/internal/platform/postgres"
	"pos-go/internal/platform/state/memory"
	jwtissuer "pos-go/internal/platform/token/jwt"
	httpmw "pos-go/internal/transport/http/middleware"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type App struct {
	Echo *echo.Echo
	DB   *pgxpool.Pool
}

func newServiceCatalogItemID() (servicecatalogdomain.ServiceCatalogItemID, error) {
	return servicecatalogdomain.ServiceCatalogItemID(uuid.NewString()), nil
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	pool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.HTTPErrorHandler = httpresponse.HTTPErrorHandler
	e.HideBanner = true
	e.HidePort = true

	e.Use(httpmw.RequestID)
	e.Use(httpmw.Recover)

	api := e.Group("/api")

	healthChecker := postgres.NewHealthChecker(pool)
	healthHandler := systemhttp.NewHealthHandler(healthChecker)
	healthHandler.Register(api)

	if cfg.Auth.Google.IsConfigured() || cfg.Auth.Debug.Enabled {
		tokenIssuer, err := jwtissuer.NewHMACIssuer(
			cfg.Auth.JWT.Issuer,
			cfg.Auth.JWT.Aud,
			cfg.Auth.JWT.Kid,
			cfg.Auth.JWT.Secret,
			cfg.Auth.JWT.TTL,
		)
		if err != nil {
			pool.Close()
			return nil, err
		}

		tokenVerifier, err := jwtissuer.NewHMACVerifier(
			cfg.Auth.JWT.Issuer,
			cfg.Auth.JWT.Aud,
			cfg.Auth.JWT.Secret,
		)
		if err != nil {
			pool.Close()
			return nil, err
		}

		stateStore := memory.NewAuthStateStore()
		accountRepo := postgres.NewAccountIdentityRepository(pool)
		sessionStore := postgres.NewSessionStore(pool)
		sessionStatusChecker := postgres.NewSessionStatusChecker(pool)
		sessionRevoker := postgres.NewSessionRevoker(pool)
		refreshRepo := postgres.NewRefreshSessionRepository(pool)
		transactor := postgres.NewTransactor(pool)
		roleAssigner := postgres.NewAccountRoleAssigner(pool)
		roleRemover := postgres.NewAccountRoleRemover(pool)
		principalResolver := postgres.NewPrincipalResolver(pool)
		capabilityRepository := postgres.NewCapabilityRepository(pool)
		checkCapabilityUsecase := capabilityusecase.NewCheckCapability(capabilityRepository)

		serviceCatalogRepository := postgres.NewServiceCatalogRepository(pool)
		listServiceCatalogItemsUsecase := servicecatalogusecase.NewListServiceCatalogItems(serviceCatalogRepository)
		lookupServiceCatalogItemsUsecase := servicecatalogusecase.NewLookupServiceCatalogItems(serviceCatalogRepository)
		showServiceCatalogItemUsecase := servicecatalogusecase.NewShowServiceCatalogItem(serviceCatalogRepository)
		createServiceCatalogItemUsecase := servicecatalogusecase.NewCreateServiceCatalogItem(
			serviceCatalogRepository,
			newServiceCatalogItemID,
			time.Now,
		)
		updateServiceCatalogItemUsecase := servicecatalogusecase.NewUpdateServiceCatalogItem(
			serviceCatalogRepository,
			time.Now,
		)
		activateServiceCatalogItemUsecase := servicecatalogusecase.NewActivateServiceCatalogItem(serviceCatalogRepository)
		deactivateServiceCatalogItemUsecase := servicecatalogusecase.NewDeactivateServiceCatalogItem(serviceCatalogRepository)
		serviceCatalogHandler := servicecataloghttp.NewServiceCatalogHandler(
			listServiceCatalogItemsUsecase,
			lookupServiceCatalogItemsUsecase,
			showServiceCatalogItemUsecase,
			createServiceCatalogItemUsecase,
			updateServiceCatalogItemUsecase,
			activateServiceCatalogItemUsecase,
			deactivateServiceCatalogItemUsecase,
		)

		productRepository := postgres.NewProductRepository(pool)
		productAuditRecorder := productCatalogNoopAuditRecorder{}
		productIDGenerator := productCatalogUUIDGenerator{}
		listProductsUsecase := productcatalogusecase.NewListProducts(productRepository)
		lookupProductsUsecase := productcatalogusecase.NewLookupProducts(productRepository)
		getProductDetailUsecase := productcatalogusecase.NewGetProductDetail(productRepository)
		createProductUsecase := productcatalogusecase.NewCreateProduct(
			productRepository,
			productRepository,
			productRepository,
			productAuditRecorder,
			productIDGenerator,
			time.Now,
		)
		updateProductUsecase := productcatalogusecase.NewUpdateProduct(
			productRepository,
			productRepository,
			productRepository,
			productAuditRecorder,
			time.Now,
		)
		softDeleteProductUsecase := productcatalogusecase.NewSoftDeleteProduct(
			productRepository,
			productRepository,
			productAuditRecorder,
			time.Now,
		)
		restoreProductUsecase := productcatalogusecase.NewRestoreProduct(
			productRepository,
			productRepository,
			productAuditRecorder,
			time.Now,
		)
		listProductVersionsUsecase := productcatalogusecase.NewListProductVersions(productRepository)
		productCatalogHandler := productcataloghttp.NewProductCatalogHandler(
			listProductsUsecase,
			lookupProductsUsecase,
			getProductDetailUsecase,
			createProductUsecase,
			updateProductUsecase,
			softDeleteProductUsecase,
			restoreProductUsecase,
			listProductVersionsUsecase,
		)

		authGroup := api.Group("/auth")

		if cfg.Auth.Google.IsConfigured() {
			oidcProvider, err := googleoidc.NewOIDC(ctx, googleoidc.OIDCConfig{
				Issuer:       cfg.Auth.Google.Issuer,
				ClientID:     cfg.Auth.Google.ClientID,
				ClientSecret: cfg.Auth.Google.ClientSecret,
				RedirectURL:  cfg.Auth.Google.RedirectURL,
			})
			if err != nil {
				pool.Close()
				return nil, err
			}

			googleFlow := authusecase.NewGoogleFlow(
				oidcProvider,
				stateStore,
				accountRepo,
				sessionStore,
				tokenIssuer,
				transactor,
				cfg.Auth.StateTTL,
				cfg.Auth.SessionTTL,
			).WithRoleAssigner(roleAssigner)

			googleHandler := authhttp.NewGoogleHandler(
				authhttp.NewGoogleFlowAdapter(googleFlow),
				cfg.Auth.Google.RedirectURL,
			)
			googleHandler.Register(authGroup)
		}

		if cfg.Auth.Debug.Enabled {
			manualLoginUsecase := authusecase.NewManualLogin(
				accountRepo,
				roleAssigner,
				sessionStore,
				tokenIssuer,
				transactor,
				cfg.Auth.SessionTTL,
			)
			manualLoginHandler := authhttp.NewManualLoginHandler(manualLoginUsecase)
			manualLoginHandler.Register(authGroup)
		}

		refreshUsecase := authusecase.NewRefreshToken(
			refreshRepo,
			tokenIssuer,
			cfg.Auth.SessionTTL,
		)
		refreshHandler := authhttp.NewRefreshHandler(refreshUsecase)
		refreshHandler.Register(authGroup)

		meHandler := systemhttp.NewMeHandler()

		meGroup := api.Group("")
		meGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		meGroup.Use(httpmw.RequirePermission("profile.self.read"))
		meGroup.Use(httpmw.RequireCapability("profile.self.show", checkCapabilityUsecase))
		meHandler.Register(meGroup)

		authzGroup := api.Group("/authz")
		authzGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		authzGroup.Use(httpmw.RequirePermission("profile.self.read"))
		authzGroup.Use(httpmw.RequireCapability("authz.profile.self.show", checkCapabilityUsecase))
		meHandler.Register(authzGroup)

		logoutGroup := api.Group("/auth")
		logoutGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		logoutGroup.Use(httpmw.RequirePermission("auth.session.logout"))
		logoutGroup.Use(httpmw.RequireCapability("auth.session.logout", checkCapabilityUsecase))

		logoutUsecase := authusecase.NewLogoutCurrentSession(sessionRevoker)
		logoutHandler := authhttp.NewLogoutHandler(logoutUsecase)
		logoutHandler.Register(logoutGroup)

		assignAccountRoleUsecase := authusecase.NewAssignAccountRole(roleAssigner)
		removeAccountRoleUsecase := authusecase.NewRemoveAccountRole(roleRemover)
		accountRoleHandler := authhttp.NewAccountRoleHandler(
			assignAccountRoleUsecase,
			removeAccountRoleUsecase,
		)

		adminAssignGroup := api.Group("/admin")
		adminAssignGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		adminAssignGroup.Use(httpmw.RequirePermission("account.role.assign"))
		adminAssignGroup.Use(httpmw.RequireCapability("account.role.assign", checkCapabilityUsecase))
		accountRoleHandler.RegisterAssign(adminAssignGroup)

		adminRemoveGroup := api.Group("/admin")
		adminRemoveGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		adminRemoveGroup.Use(httpmw.RequirePermission("account.role.assign"))
		adminRemoveGroup.Use(httpmw.RequireCapability("account.role.remove", checkCapabilityUsecase))
		accountRoleHandler.RegisterRemove(adminRemoveGroup)

		adminCapabilityGroup := api.Group("/admin")
		adminCapabilityGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		adminCapabilityGroup.Use(httpmw.RequirePermission("capability.manage"))
		adminCapabilityGroup.Use(httpmw.RequireCapability("capability.manage", checkCapabilityUsecase))

		listCapabilitiesUsecase := capabilityusecase.NewListCapabilities(capabilityRepository)
		showCapabilityUsecase := capabilityusecase.NewShowCapability(capabilityRepository)
		enableCapabilityUsecase := capabilityusecase.NewEnableCapability(capabilityRepository)
		disableCapabilityUsecase := capabilityusecase.NewDisableCapability(capabilityRepository)
		capabilityHandler := capabilityhttp.NewCapabilityHandler(
			listCapabilitiesUsecase,
			showCapabilityUsecase,
			enableCapabilityUsecase,
			disableCapabilityUsecase,
		)
		capabilityHandler.Register(adminCapabilityGroup)

		serviceCatalogListGroup := api.Group("/service-catalog")
		serviceCatalogListGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		serviceCatalogListGroup.Use(httpmw.RequirePermission("service_catalog.read"))
		serviceCatalogListGroup.Use(httpmw.RequireCapability("service_catalog.list", checkCapabilityUsecase))
		serviceCatalogHandler.RegisterList(serviceCatalogListGroup)

		serviceCatalogCreateGroup := api.Group("/service-catalog")
		serviceCatalogCreateGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		serviceCatalogCreateGroup.Use(httpmw.RequirePermission("service_catalog.manage"))
		serviceCatalogCreateGroup.Use(httpmw.RequireCapability("service_catalog.create", checkCapabilityUsecase))
		serviceCatalogHandler.RegisterCreate(serviceCatalogCreateGroup)

		serviceCatalogLookupGroup := api.Group("/service-catalog")
		serviceCatalogLookupGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		serviceCatalogLookupGroup.Use(httpmw.RequirePermission("service_catalog.read"))
		serviceCatalogLookupGroup.Use(httpmw.RequireCapability("service_catalog.lookup", checkCapabilityUsecase))
		serviceCatalogHandler.RegisterLookup(serviceCatalogLookupGroup)

		serviceCatalogShowGroup := api.Group("/service-catalog")
		serviceCatalogShowGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		serviceCatalogShowGroup.Use(httpmw.RequirePermission("service_catalog.read"))
		serviceCatalogShowGroup.Use(httpmw.RequireCapability("service_catalog.show", checkCapabilityUsecase))
		serviceCatalogHandler.RegisterShow(serviceCatalogShowGroup)

		serviceCatalogUpdateGroup := api.Group("/service-catalog")
		serviceCatalogUpdateGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		serviceCatalogUpdateGroup.Use(httpmw.RequirePermission("service_catalog.manage"))
		serviceCatalogUpdateGroup.Use(httpmw.RequireCapability("service_catalog.update", checkCapabilityUsecase))
		serviceCatalogHandler.RegisterUpdate(serviceCatalogUpdateGroup)

		serviceCatalogActivateGroup := api.Group("/service-catalog")
		serviceCatalogActivateGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		serviceCatalogActivateGroup.Use(httpmw.RequirePermission("service_catalog.manage"))
		serviceCatalogActivateGroup.Use(httpmw.RequireCapability("service_catalog.activate", checkCapabilityUsecase))
		serviceCatalogHandler.RegisterActivate(serviceCatalogActivateGroup)

		serviceCatalogDeactivateGroup := api.Group("/service-catalog")
		serviceCatalogDeactivateGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		serviceCatalogDeactivateGroup.Use(httpmw.RequirePermission("service_catalog.manage"))
		serviceCatalogDeactivateGroup.Use(httpmw.RequireCapability("service_catalog.deactivate", checkCapabilityUsecase))
		serviceCatalogHandler.RegisterDeactivate(serviceCatalogDeactivateGroup)

		newProductCatalogGroup := func(permissionKey, capabilityKey string) *echo.Group {
			group := api.Group("")
			group.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
			group.Use(httpmw.RequirePermission(permissionKey))
			group.Use(httpmw.RequireCapability(capabilityKey, checkCapabilityUsecase))

			return group
		}

		productCatalogHandler.RegisterList(newProductCatalogGroup(
			"product_catalog.read",
			"product_catalog.list",
		))
		productCatalogHandler.RegisterCreate(newProductCatalogGroup(
			"product_catalog.manage",
			"product_catalog.create",
		))
		productCatalogHandler.RegisterLookup(newProductCatalogGroup(
			"product_catalog.read",
			"product_catalog.lookup",
		))
		productCatalogHandler.RegisterVersions(newProductCatalogGroup(
			"product_catalog.read",
			"product_catalog.versions",
		))
		productCatalogHandler.RegisterRestore(newProductCatalogGroup(
			"product_catalog.manage",
			"product_catalog.restore",
		))
		productCatalogHandler.RegisterShow(newProductCatalogGroup(
			"product_catalog.read",
			"product_catalog.show",
		))
		productCatalogHandler.RegisterUpdate(newProductCatalogGroup(
			"product_catalog.manage",
			"product_catalog.update",
		))
		productCatalogHandler.RegisterDelete(newProductCatalogGroup(
			"product_catalog.manage",
			"product_catalog.delete",
		))
	}

	return &App{
		Echo: e,
		DB:   pool,
	}, nil
}
