package middleware

import (
	"errors"
	"net/http"

	"github.com/diyabhandari/Go-API/api"
	"github.com/diyabhandari/Go-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

// custom unauthorized error
var UnAuthorizedError = errors.New("Invalid username or taken ")

// cuz its middleware, it must take in AND return an HTTP handler interface
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.RepsonseWriter, r *http.Request) { //handlerFunc is from http package, it takes in a func, and that func also takes in response writer(to construct a response to the caller, set response body header etc) and  a pointer to the request(contains info on incoming req like headers payload)
		var username string = r.URL.Query().Get("username") //grab username from request pointer
		var token = r.Header.Get("Authorization")           //auth token grabbed from header
		//if either of the above are empty, return an error
		var err error
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError) //called new req handler
			return
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil { //now in this case if we get an error, we return the internal error handler instead
			api.InternalErrorHandler(w)
		}
		//now, query the database
		var loginDetails = *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		} //we didnt find a client with a username or we didnt get the token from the database
		next.ServeHTTP(w, r) //calls the next middleware in line or the handler func for the endpoint when no middleware left, we'll call getcoinbalance in our case
	}) //) on a new line gave error
}
