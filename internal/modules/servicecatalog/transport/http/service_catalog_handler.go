package http

import "github.com/labstack/echo/v4"

type ServiceCatalogHandler struct {
	list       listServiceCatalogItems
	lookup     lookupServiceCatalogItems
	show       showServiceCatalogItem
	create     createServiceCatalogItem
	update     updateServiceCatalogItem
	activate   activateServiceCatalogItem
	deactivate deactivateServiceCatalogItem
}

func NewServiceCatalogHandler(
	list listServiceCatalogItems,
	lookup lookupServiceCatalogItems,
	show showServiceCatalogItem,
	create createServiceCatalogItem,
	update updateServiceCatalogItem,
	activate activateServiceCatalogItem,
	deactivate deactivateServiceCatalogItem,
) ServiceCatalogHandler {
	return ServiceCatalogHandler{
		list:       list,
		lookup:     lookup,
		show:       show,
		create:     create,
		update:     update,
		activate:   activate,
		deactivate: deactivate,
	}
}

func (h ServiceCatalogHandler) RegisterList(group *echo.Group) {
	group.GET("/items", h.List)
}

func (h ServiceCatalogHandler) RegisterCreate(group *echo.Group) {
	group.POST("/items", h.Create)
}

func (h ServiceCatalogHandler) RegisterLookup(group *echo.Group) {
	group.GET("/items/lookup", h.Lookup)
}

func (h ServiceCatalogHandler) RegisterShow(group *echo.Group) {
	group.GET("/items/:id", h.Show)
}

func (h ServiceCatalogHandler) RegisterUpdate(group *echo.Group) {
	group.PUT("/items/:id", h.Update)
}

func (h ServiceCatalogHandler) RegisterActivate(group *echo.Group) {
	group.POST("/items/:id/activate", h.Activate)
}

func (h ServiceCatalogHandler) RegisterDeactivate(group *echo.Group) {
	group.POST("/items/:id/deactivate", h.Deactivate)
}
