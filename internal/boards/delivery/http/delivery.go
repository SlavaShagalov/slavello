package http

import (
	"context"
	pBoards "github.com/SlavaShagalov/slavello/internal/boards"
	mw "github.com/SlavaShagalov/slavello/internal/middleware"
	"github.com/SlavaShagalov/slavello/internal/pkg/constants"
	pErrors "github.com/SlavaShagalov/slavello/internal/pkg/errors"
	pHTTP "github.com/SlavaShagalov/slavello/internal/pkg/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type delivery struct {
	uc  pBoards.Usecase
	log *zap.Logger
}

func RegisterHandlers(mux *mux.Router, uc pBoards.Usecase, log *zap.Logger, checkAuth mw.Middleware, metrics mw.Middleware) {
	del := delivery{
		uc:  uc,
		log: log,
	}

	const (
		workspaceBoardsPrefix = "/workspaces/{id}/boards"
		workspaceBoardsPath   = constants.ApiPrefix + workspaceBoardsPrefix

		boardsPrefix = "/boards"
		boardsPath   = constants.ApiPrefix + boardsPrefix
		boardPath    = boardsPath + "/{id}"
	)

	mux.HandleFunc(workspaceBoardsPath, checkAuth(del.create)).Methods(http.MethodPost)
	mux.HandleFunc(workspaceBoardsPath, checkAuth(del.listByWorkspace)).Methods(http.MethodGet)

	mux.HandleFunc(boardsPath, checkAuth(del.list)).Methods(http.MethodGet).
		Queries("title", "{title}")

	mux.HandleFunc(boardPath, checkAuth(del.get)).Methods(http.MethodGet)
	mux.HandleFunc(boardPath, checkAuth(del.partialUpdate)).Methods(http.MethodPatch)
	mux.HandleFunc(boardPath, checkAuth(del.delete)).Methods(http.MethodDelete)
}

// create godoc
//
//	@Summary		Create a new board
//	@Description	Create a new board
//	@Tags			workspaces
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int				true	"Workspace ID"
//	@Param			BoardCreateData	body		createRequest	true	"Board create data"
//	@Success		200				{object}	createResponse	"Created board data."
//	@Failure		400				{object}	http.JSONError
//	@Failure		401				{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/workspaces/{id}/boards [post]
//
//	@Security		cookieAuth
func (del *delivery) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workspaceID, err := strconv.Atoi(vars["id"])
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	body, err := pHTTP.ReadBody(r, del.log)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	var request createRequest
	err = request.UnmarshalJSON(body)
	if err != nil {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}

	params := pBoards.CreateParams{
		Title:       request.Title,
		Description: request.Description,
		WorkspaceID: workspaceID,
	}

	board, err := del.uc.Create(ctx, &params)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newCreateResponse(&board)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

// listByWorkspace godoc
//
//	@Summary		Returns boards by workspace id
//	@Description	Returns boards by workspace id
//	@Tags			workspaces
//	@Produce		json
//	@Param			id	path		int				true	"Workspace ID"
//	@Success		200	{object}	listResponse	"Boards data"
//	@Failure		400	{object}	http.JSONError
//	@Failure		401	{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/workspaces/{id}/boards [get]
//
//	@Security		cookieAuth
func (del *delivery) listByWorkspace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workspaceID, err := strconv.Atoi(vars["id"])
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	boards, err := del.uc.ListByWorkspace(context.Background(), workspaceID)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newListResponse(boards)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

// list godoc
//
//	@Summary		Returns boards by workspace id
//	@Description	Returns boards by workspace id
//	@Tags			boards
//	@Produce		json
//	@Param			title	query		string			true	"Title filter"
//	@Success		200		{object}	listResponse	"Boards data"
//	@Failure		400		{object}	http.JSONError
//	@Failure		401		{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/boards [get]
//
//	@Security		cookieAuth
func (del *delivery) list(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(mw.ContextUserID).(int)
	if !ok {
		pHTTP.HandleError(w, r, pErrors.ErrReadBody)
		return
	}

	title := r.FormValue("title")

	boards, err := del.uc.ListByTitle(context.Background(), title, userID)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newListResponse(boards)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

// get godoc
//
//	@Summary		Returns board by id
//	@Description	Returns board by id
//	@Tags			boards
//	@Produce		json
//	@Param			id	path		int			true	"Board ID"
//	@Success		200	{object}	getResponse	"Board data"
//	@Failure		400	{object}	http.JSONError
//	@Failure		401	{object}	http.JSONError
//	@Failure		404	{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/boards/{id} [get]
//
//	@Security		cookieAuth
func (del *delivery) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardID, err := strconv.Atoi(vars["id"])
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	board, err := del.uc.Get(context.Background(), boardID)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newGetResponse(&board)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

// partialUpdate godoc
//
//	@Summary		Partial update of board
//	@Description	Partial update of board
//	@Tags			boards
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int						true	"Board ID"
//	@Param			BoardUpdateData	body		partialUpdateRequest	true	"Board data to update"
//	@Success		200				{object}	getResponse				"Updated board data."
//	@Failure		400				{object}	http.JSONError
//	@Failure		401				{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/boards/{id}  [patch]
//
//	@Security		cookieAuth
func (del *delivery) partialUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardID, err := strconv.Atoi(vars["id"])
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

	params := pBoards.PartialUpdateParams{ID: boardID}
	params.UpdateTitle = request.Title != nil
	if params.UpdateTitle {
		params.Title = *request.Title
	}
	params.UpdateDescription = request.Description != nil
	if params.UpdateDescription {
		params.Description = *request.Description
	}

	board, err := del.uc.PartialUpdate(context.Background(), &params)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	response := newGetResponse(&board)
	pHTTP.SendJSON(w, r, http.StatusOK, response)
}

// delete godoc
//
//	@Summary		Delete board by id
//	@Description	Delete board by id
//	@Tags			boards
//	@Produce		json
//	@Param			id	path	int	true	"Board ID"
//	@Success		204	"Board deleted successfully"
//	@Failure		400	{object}	http.JSONError
//	@Failure		401	{object}	http.JSONError
//	@Failure		404	{object}	http.JSONError
//	@Failure		405
//	@Failure		500
//	@Router			/boards/{id} [delete]
//
//	@Security		cookieAuth
func (del *delivery) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardID, err := strconv.Atoi(vars["id"])
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	err = del.uc.Delete(context.Background(), boardID)
	if err != nil {
		pHTTP.HandleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
