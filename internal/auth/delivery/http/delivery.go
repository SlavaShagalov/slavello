package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/SlavaShagalov/slavello/internal/auth"
	mw "github.com/SlavaShagalov/slavello/internal/middleware"
	"github.com/SlavaShagalov/slavello/internal/pkg/constants"
	pErrors "github.com/SlavaShagalov/slavello/internal/pkg/errors"
	pHTTP "github.com/SlavaShagalov/slavello/internal/pkg/http"
	"github.com/SlavaShagalov/slavello/internal/users"
)

const (
	authPrefix = "/auth"
	signInPath = constants.ApiPrefix + authPrefix + "/signin"
	signUpPath = constants.ApiPrefix + authPrefix + "/signup"
	logoutPath = constants.ApiPrefix + authPrefix + "/logout"
	mePath     = constants.ApiPrefix + authPrefix + "/me"
)

type delivery struct {
	uc      auth.Usecase
	usersUC users.Usecase
	log     *zap.Logger
}

func RegisterHandlers(mux *mux.Router, uc auth.Usecase, usersUC users.Usecase, log *zap.Logger, checkAuth mw.Middleware) {
	del := delivery{
		uc:      uc,
		usersUC: usersUC,
		log:     log,
	}

	mux.HandleFunc(signUpPath, del.signup).Methods(http.MethodPost)
	mux.HandleFunc(signInPath, del.signin).Methods(http.MethodPost)
	mux.HandleFunc(logoutPath, checkAuth(del.logout)).Methods(http.MethodDelete)
	mux.HandleFunc(mePath, checkAuth(del.me)).Methods(http.MethodGet)
}

// signup godoc
//
//	@Summary		Creates new user and returns authentication cookie.
//	@Description	Creates new user and returns authentication cookie.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			signUpParams	body		SignUpRequest	true	"Sign up params."
//	@Success		200				{object}	SignUpResponse	"Successfully created user."
//	@Failure		400				{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/auth/signup [post]
func (d *delivery) signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := pHTTP.ReadBody(r, d.log)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	var request SignUpRequest
	err = request.UnmarshalJSON(body)
	if err != nil {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}

	params := auth.SignUpParams{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	user, authToken, err := d.uc.SignUp(ctx, &params)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	sessionCookie := createSessionCookie(authToken)
	http.SetCookie(w, sessionCookie)

	response := newSignUpResponse(&user)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

// signin godoc
//
//	@Summary		Logs in and returns the authentication cookie
//	@Description	Logs in and returns the authentication cookie
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			signInParams	body		SignInRequest	true	"Successfully authenticated."
//	@Success		200				{object}	SignInResponse	"successfully auth"
//	@Failure		400				{object}	http.JSONError
//	@Failure		404				{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/auth/signin [post]
func (d *delivery) signin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := pHTTP.ReadBody(r, d.log)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	var request SignInRequest
	err = request.UnmarshalJSON(body)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	params := auth.SignInParams{
		Username: request.Username,
		Password: request.Password,
	}

	user, authToken, err := d.uc.SignIn(ctx, &params)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	sessionCookie := createSessionCookie(authToken)
	http.SetCookie(w, sessionCookie)

	response := newSignInResponse(&user)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

// logout godoc
//
//	@Summary		Logs out and deletes the authentication cookie.
//	@Description	Logs out and deletes the authentication cookie.
//	@Tags			auth
//	@Produce		json
//	@Success		204	"Successfully logged out."
//	@Failure		400	{object}	http.JSONError
//	@Failure		401	{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/auth/logout [delete]
//
//	@Security		cookieAuth
func (d *delivery) logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := r.Context().Value(mw.ContextUserID).(int)
	if !ok {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}
	authToken, ok := r.Context().Value(mw.ContextAuthToken).(string)
	if !ok {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}

	err := d.uc.Logout(ctx, userID, authToken)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	newCookie := &http.Cookie{
		Name:     constants.SessionName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Now().Add(-24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, newCookie)

	w.WriteHeader(http.StatusNoContent)
}

func (d *delivery) me(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(mw.ContextUserID).(int)
	if !ok {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}

	user, err := d.usersUC.Get(userID)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newGetResponse(&user)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}
