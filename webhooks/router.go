package webhooks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type Router struct {
	handlers map[string]func(context.Context, *Envelope, json.RawMessage) error
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]func(context.Context, *Envelope, json.RawMessage) error),
	}
}

// Handle registers a handler for the given event name.
// h must be a function with signature: func(context.Context, *Envelope, []T) error
func (r *Router) Handle(eventName string, h interface{}) {
	val := reflect.ValueOf(h)
	typ := val.Type()

	if typ.Kind() != reflect.Func {
		panic("handler must be a function")
	}
	if typ.NumIn() != 3 {
		panic("handler must have 3 arguments")
	}
	if typ.NumOut() != 1 {
		panic("handler must have 1 return value")
	}
	if !typ.Out(0).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		panic("handler must return error")
	}

	// Check argument types
	// arg0: context.Context
	// arg1: *Envelope
	// arg2: []T

	sliceType := typ.In(2)
	if sliceType.Kind() != reflect.Slice {
		panic("handler 3rd argument must be a slice")
	}

	r.handlers[eventName] = func(ctx context.Context, env *Envelope, rawData json.RawMessage) error {
		// Create pointer to slice for Unmarshal
		slicePtr := reflect.New(sliceType)

		if err := json.Unmarshal(rawData, slicePtr.Interface()); err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidJSON, err)
		}

		in := []reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(env),
			slicePtr.Elem(),
		}

		out := val.Call(in)
		if !out[0].IsNil() {
			return out[0].Interface().(error)
		}
		return nil
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Limit body size (e.g. 10MB)
	req.Body = http.MaxBytesReader(w, req.Body, 10*1024*1024)
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Request body too large", http.StatusRequestEntityTooLarge)
		return
	}

	var env Envelope
	if err := json.Unmarshal(body, &env); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	handler, ok := r.handlers[env.EventName]
	if !ok {
		http.Error(w, "Unregistered event", http.StatusNotFound)
		return
	}

	if err := handler(req.Context(), &env, env.Data); err != nil {
		if errors.Is(err, ErrInvalidJSON) {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
