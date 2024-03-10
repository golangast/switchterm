package tokens

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	Name   string `json:"name"`
	Admin  bool   `json:"admin"`
	Claims jwt.StandardClaims
}

func (j MyCustomClaims) Valid() error {
	var err error
	return err
}

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Timername() string {
	ticker := time.NewTicker(time.Hour)
	for ; true; <-ticker.C {
		return CreateHash("lannarr")
	}
	return ""
}
func Timerkey() []byte {
	ticker := time.NewTicker(time.Hour)
	for ; true; <-ticker.C {
		return Keyen()
	}
	b := []byte("")
	return b
}

func Sevenkey() []byte {
	ticker := time.NewTicker((time.Minute))
	for ; true; <-ticker.C {
		return Keyen()
	}
	b := []byte("")
	return b
}

type Keys struct {
	Name   string
	Key    []byte
	Seven  []byte
	Claims MyCustomClaims
}

var claimer = MyCustomClaims{
	Timername(),
	true,
	jwt.StandardClaims{
		Audience:  Timername(),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "moviecast",
	},
}

var K = Keys{Name: Timername(), Key: []byte(Timerkey()), Seven: []byte(Sevenkey()), Claims: claimer}

func Keyen() []byte {
	bytes := make([]byte, 64) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	return bytes
}

func ContextToken() string {
	mySigningKey := K.Key

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		K.Name,
		jwt.StandardClaims{
			Audience:  Timername(),
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "lannarr",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss

}
func Checktoken(t string) bool {

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return K.Key, nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			return false

		} else {
			fmt.Println("Couldn't handle this token:", err)
			return false

		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return false

	}

}
func Checktokencontext(t string) bool {

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return K.Seven, nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			return false

		} else {
			fmt.Println("Couldn't handle this token:", err)
			return false

		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return false

	}

}
func Overrideclaims(t string) {
	// sample token is expired.  override time so it parses as valid
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// sample token is expired.  override time so it parses as valid
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(t, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return K.Key, nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
		} else {
			fmt.Println(err)
		}
	})

}

// Override time value for tests.  Restore default value after.
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
func Tokerner() string {
	mySigningKey := K.Key

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		K.Name,
		jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "lannarr",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss

}
func HeadTokerner(k interface{}) string {

	mySigningKey := k

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"header" + K.Name,
		jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "lannarr",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss

}
