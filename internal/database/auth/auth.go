package auth

import (
	"net/http"
	"errors"
	"strings"
)


/*
* Authorization: ApiKey {actual api key here}
*/
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == ""{
		return "", errors.New("No authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed auth header")
	}
	if vals[0] != "ApiKey"{
		return "", errors.New("Malformed first part of the auth header")
	}

	return vals[1], nil

}
