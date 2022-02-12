package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:          "/publications",
		Method:       http.MethodPost,
		Handler:      controllers.CreatePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications",
		Method:       http.MethodGet,
		Handler:      controllers.SearchPublications,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodGet,
		Handler:      controllers.SearchPublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdatePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeletePublication,
		RequiresAuth: true,
	},
}
