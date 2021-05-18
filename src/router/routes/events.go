package routes

import (
	"moneycount-api/src/controllers"
	"net/http"
)

var eventsRoutes = []Route{
	{
		URI: "/event-maps",
		Method: http.MethodPost,
		Function: controllers.CreateEventMap,
		RequiresAuthentication: false,
	},
}