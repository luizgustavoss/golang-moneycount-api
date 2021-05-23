package routes

import (
	"moneycount-api/src/interfaces/rest/controllers"
	"net/http"
)

var eventEntriesRoutes = []Route{
	{
		URI:                    "/event-entries-maps",
		Method:                 http.MethodPost,
		Function:               controllers.CreateEventEntryMap,
		RequiresAuthentication: false,
	},
}