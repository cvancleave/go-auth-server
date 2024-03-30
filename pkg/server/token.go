package server

import (
	"net/http"

	"github.com/cvancleave/go-auth-server/pkg/utils"
	"github.com/julienschmidt/httprouter"
)

// oauth standard endpoint
func (s *server) handleTokenRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SetCorsHeaders(w, r)

	// get form values from request
	if err := r.ParseForm(); err != nil {
		utils.RespondError(w, 500, "error parsing form")
		return
	}

	clientId, idOk := r.Form["client_id"]
	clientSecret, secretOk := r.Form["client_secret"]
	_ = clientSecret

	if !idOk || !secretOk {
		utils.RespondError(w, 400, "error finding form values")
		return
	}

	// check clientid and secret against values in a database to verify
	// --

	// create new jwt
	accessToken, err := utils.NewJwt(clientId[0], s.audience, s.secretKey, s.issuer, 60)
	if err != nil {
		utils.RespondError(w, 500, "error validating token")
		return
	}

	// send token response
	response := TokenResponse{
		AccessToken: accessToken,
	}

	utils.RespondJson(w, 200, response)
}
