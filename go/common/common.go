package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

//Log is a decorator middleware log request if DEBUG is set
func Log(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if IsDebugEnabled() {
			body, _ := ioutil.ReadAll(r.Body)
			log.Println("Recived a request")
			header := r.Header
			log.Println("The headers are : ")
			for k, s := range header {
				fmt.Println(k, "=", s)
			}
			log.Println("The request is : ")
			log.Println(string(body))
			//Check for non empty body and close it
			//Do not defer the call as we need to set back the body
			if body != nil {
				r.Body.Close()
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		f(w, r)
	}
}

//Auth is a basic authorization decorator middeware
func Auth(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorized := false
		for _, v := range r.Header.Values(authHeaderKey) {
			if v == authValue {
				authorized = true
				break
			}
		}
		if authorized {
			f(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(struct{ Error string }{"Unauthorized"})
		}
	}
}
