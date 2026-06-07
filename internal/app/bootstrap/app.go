// audit:allow-oversize reason=bootstrap-wiring
package bootstrap

import (
	"context"

	"pos-go/internal/config"
	authhttp "pos-go/internal/modules/auth/transport/http"
	authusecase "pos-go/internal/modules/auth/usecase"
	capabilityhttp "pos-go/internal/modules/capability/transport/http"
	capabilityusecase "pos-go/internal/modules/capability/usecase"
	systemhttp "pos-go/internal/modules/system/transport/http"
	googleoidc "pos-go/internal/platform/google"
	"pos-go/internal/platform/postgres"
	"pos-go/internal/platform/state/memory"
	jwtissuer "pos-go/internal/platform/token/jwt"
	httpmw "pos-go/internal/transport/http/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type App struct {
	Echo *echo.Echo
	DB   *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	pool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	e := echo.New()
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
	}

	return &App{
		Echo: e,
		DB:   pool,
	}, nil
}
