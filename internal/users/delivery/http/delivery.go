package http

import (
	mw "github.com/SlavaShagalov/slavello/internal/middleware"
	"github.com/SlavaShagalov/slavello/internal/pkg/constants"
	pErrors "github.com/SlavaShagalov/slavello/internal/pkg/errors"
	pHTTP "github.com/SlavaShagalov/slavello/internal/pkg/http"
	pUsers "github.com/SlavaShagalov/slavello/internal/users"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type delivery struct {
	uc  pUsers.Usecase
	log *zap.Logger
}

func RegisterHandlers(mux *mux.Router, uc pUsers.Usecase, log *zap.Logger, checkAuth mw.Middleware) {
	del := delivery{
		uc:  uc,
		log: log,
	}

	const (
		usersPrefix = "/users"
		usersPath   = constants.ApiPrefix + usersPrefix
		userPath    = usersPath + "/{id}"
	)

	mux.HandleFunc(userPath, checkAuth(del.get)).Methods(http.MethodGet)
	mux.HandleFunc(userPath, checkAuth(del.partialUpdate)).Methods(http.MethodPatch)
}

// get godoc
//
//	@Summary		Returns user with specified id
//	@Description	Returns user with specified id
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int			true	"User ID"
//	@Success		200	{object}	getResponse	"User data"
//	@Failure		400	{object}	http.JSONError
//	@Failure		401	{object}	http.JSONError
//	@Failure		404	{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/users/{id} [get]
//
//	@Security		cookieAuth
func (del *delivery) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	user, err := del.uc.Get(userID)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newGetResponse(&user)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

// partialUpdate godoc
//
//	@Summary		Partial update of user
//	@Description	Partial update of user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int						true	"User ID"
//	@Param			UserUpdateData	body		partialUpdateRequest	true	"User data to update"
//	@Success		200				{object}	getResponse				"Updated user data"
//	@Failure		400				{object}	http.JSONError
//	@Failure		401				{object}	http.JSONError
//	@Failure		403				{object}	http.JSONError
//	@Failure		404				{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/users/{id} [patch]
//
//	@Security		cookieAuth
func (del *delivery) partialUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	body, err := pHTTP.ReadBody(r, del.log)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	var request partialUpdateRequest
	err = request.UnmarshalJSON(body)
	if err != nil {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}

	params := pUsers.PartialUpdateParams{ID: userID}
	params.UpdateUsername = request.Username != nil
	if params.UpdateUsername {
		params.Username = *request.Username
	}
	params.UpdateEmail = request.Email != nil
	if params.UpdateEmail {
		params.Email = *request.Email
	}
	params.UpdateName = request.Name != nil
	if params.UpdateName {
		params.Name = *request.Name
	}

	workspace, err := del.uc.PartialUpdate(&params)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newGetResponse(&workspace)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}
