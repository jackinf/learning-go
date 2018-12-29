package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/harlow/authtoken"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	privKeyPath = "/Users/jrumjantsev/dev/tmp/app.rsa"
	pubKeyPath  = "/Users/jrumjantsev/dev/tmp/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s", signKey)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s", signKey)
		panic(err)
	}
}

func GenerateJWT(name, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "admin",
		"UserInfo": struct {
			Name string
			Role string
		}{name, role},
		"exp": time.Now().Add(time.Minute * 20).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authToken, err := authtoken.FromRequest(r)
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return signKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				DisplayAppError(w, err, "Access Token expired, get a new Token", 400)
				return
			default:
				DisplayAppError(w, err, "Error while parsing the access token", 400)
				return
			}
		default:
			DisplayAppError(w, err, "Error while parsing the access token", 400)
			return
		}
	}

	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(w, err, "Invalid Access Token", 401)
	}
}
