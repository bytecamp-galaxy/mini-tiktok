package jwt

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, _, err := JwtMiddleware.TokenGenerator("123")
	fmt.Println(token)
	fmt.Println(err)
}

func TestParseToken(t *testing.T) {

}
