// Code generated by hertz generator. DO NOT EDIT.

package Api

import (
	api "github.com/bytecamp-galaxy/mini-tiktok/api-server/biz/handler/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action := _publish.Group("/action", _actionMw()...)
				_action.POST("/", append(_publish_ctionMw(), api.PublishAction)...)
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_userqueryMw(), api.UserQuery)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_userloginMw(), api.UserLogin)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_userregisterMw(), api.UserRegister)...)
			}
		}
	}
}
