package http

import (
	"autodom/services/advisor"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"autodom/services/advisor/transport"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func NewService(
	svcEndpoints transport.Endpoints, options []kithttp.ServerOption, logger log.Logger,
) http.Handler {

	var (
		r            = mux.NewRouter()
		errorLogger  = kithttp.ServerErrorLogger(logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)
	options = append(options, errorLogger, errorEncoder)

	r.Methods("POST").Path("/solutions").Handler(kithttp.NewServer(
		svcEndpoints.GetByTitle,
		decodeGetByTitleRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeGetByTitleRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	text, count := decodeIncomingRequest(r)

	return transport.GetByTitleRequest{Title: text, Number: count}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(codeFrom(err))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case advisor.ErrSolutionNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

type incomingRequest struct {
	SearchText  string `json:"searchText"`
	AnswerCount int    `json:"answerCount"`
}

func decodeIncomingRequest(r *http.Request) (string, int) {
	defer r.Body.Close()
	var ir incomingRequest

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &ir)

	if err != nil {
		panic(err)
	}

	text := string(ir.SearchText)
	count := ir.AnswerCount

	return text, count
}
