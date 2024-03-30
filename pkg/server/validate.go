package server

import (
	"net/http"

	"github.com/cvancleave/go-auth-server/pkg/utils"
	"github.com/julienschmidt/httprouter"
)

// example of jwt validation
func (s *server) handleValidateRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SetCorsHeaders(w, r)

	// get token from auth header
	token, err := utils.GetTokenFromRequest(r)
	if err != nil {
		utils.RespondError(w, 400, "error getting token from request")
		return
	}

	// validate jwt
	if err := utils.ValidateJwt(token, s.secretKey, s.issuer); err != nil {
		utils.RespondError(w, 500, "error validating token")
		return
	}

	// return whatever hidden data is needed
	// ---

	// send data response
	response := map[string]string{
		"data": "token is valid",
	}

	utils.RespondJson(w, 200, response)
}
