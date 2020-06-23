package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	logger "github.com/sirupsen/logrus"
)

func IsDebugEnabled() bool {
	//Could have stored it in a localvariable but want a dynamic behaviour without restarting tha app
	return strings.ToLower(os.Getenv("DEBUG")) == "true"
}

var authValue string

const (
	authKey          = "AUTH_HEADER"
	authHeaderKey    = "Authorization"
	authDefaultValue = "CaputDraconis"
)

func init() {
	authValue = os.Getenv(authKey)
	if authValue == "" {
		authValue = authDefaultValue
	}
}

//Log is a middleware log request if DEBUG is set
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsDebugEnabled() {
			body, _ := ioutil.ReadAll(r.Body)
			header := r.Header
			logger.WithField("header", header).WithField("body", string(body)).Info("Received a request")
			defer logger.Info("Request processed")
			//Check for non empty body and close it
			//Do not defer the call as we need to set back the body
			if body != nil {
				r.Body.Close()
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		next.ServeHTTP(w, r)
	})
}

//Auth is a basic authorization middeware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorized := false
		for _, v := range r.Header.Values(authHeaderKey) {
			if v == authValue {
				authorized = true
				break
			}
		}
		if authorized {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(struct{ Error string }{"Unauthorized"})
		}
	})
}
