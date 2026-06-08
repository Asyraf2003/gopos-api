package http

import (
	"context"

	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"
)

type listServiceCatalogItems interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.ListServiceCatalogItemsCommand,
	) ([]servicecatalogusecase.ServiceCatalogItemResult, error)
}

type lookupServiceCatalogItems interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.LookupServiceCatalogItemsCommand,
	) ([]servicecatalogusecase.ServiceCatalogLookupResult, error)
}

type showServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.ShowServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}

type createServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.CreateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}

type updateServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.UpdateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}

type activateServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.ActivateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}

type deactivateServiceCatalogItem interface {
	Execute(
		ctx context.Context,
		cmd servicecatalogusecase.DeactivateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error)
}
