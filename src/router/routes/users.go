package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeUser = []Route{
	{
		URI:               "/users",
		Method:            http.MethodPost,
		Function:          controllers.GetAllUsers,
		NeedAutentication: false,
	},
	{
		URI:               "/users/{id}",
		Method:            http.MethodGet,
		Function:          controllers.GetUser,
		NeedAutentication: false,
	},
	{
		URI:               "/users",
		Method:            http.MethodPost,
		Function:          controllers.PostUser,
		NeedAutentication: false,
	},
	{
		URI:               "/users/{id}",
		Method:            http.MethodPatch,
		Function:          controllers.UpdateUser,
		NeedAutentication: false,
	},
	{
		URI:               "/users/{id}",
		Method:            http.MethodDelete,
		Function:          controllers.DeleteUser,
		NeedAutentication: false,
	},
}
