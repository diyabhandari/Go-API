package api
import(
	"encoding/json"
	"net/http"
)
type CoinBalanceParams struct{
	Username string
}//parameters that our api endpoint will take (postman here)

type CoinBalanceResponse struct{
	Code int //success code is usually 200
	Balance int64 //account balance
}//reps sucessful response from server

type Error struct{
	Code int //error code
	Message string //error message
}
//write an error response to http response writer, returns error response to person who called the endpoint
//we have it here cuz we want to reuse it whenever theres an error somewhere else
func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resp)
}
var {
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w,err.Error(),http.StatusBadRequest)
	}//to return specific error in response, may guide the user to correct themselves
	InternalErrorHandler = func(w http.ResponseWriter){
		writeerror(w,"An unexpected Error Occurred",http.StatusInternalServerError)//generic error message cuz if the error is due to a bug in the code a long error message isnt needed as it wont help the user
	}
}