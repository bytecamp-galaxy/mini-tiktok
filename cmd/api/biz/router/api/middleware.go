// Code generated by hertz generator.

package Api

import (
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/jwt"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _comment_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentlistMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{jwt.Middleware.MiddlewareFunc()}
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favorite_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getfeedMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{jwt.Middleware.MiddlewareFunc()}
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publish_ctionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{jwt.Middleware.MiddlewareFunc()}
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userqueryMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{jwt.Middleware.MiddlewareFunc()}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userregisterMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{jwt.Middleware.MiddlewareFunc()}
}
