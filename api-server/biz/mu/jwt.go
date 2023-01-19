package mu

import (
	"github.com/hertz-contrib/jwt"
	"log"
	"mini-tiktok-v2/pkg/dal/model"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func Init() {
	var err error
	if JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:   "mini-tiktok-v2",
		Key:     []byte("114514"),
		Timeout: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		TokenLookup: "query:token",
	}); err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}
