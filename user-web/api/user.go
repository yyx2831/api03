package api

import (
	"api03/user-web/global/reponse"
	"api03/user-web/proto"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

func HandlerGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc错误码转换为http错误码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func GetUserList(c *gin.Context) {
	//ip := "127.0.0.1"
	//port := 50051
	//拨号连接grpc服务器
	userConn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		zap.L().Error("[GetUserList]连接【用户服务失败】", zap.Error(err))
	}
	//生成grpc的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 10,
	})
	if err != nil {
		zap.L().Error("[GetUserList]查询【用户列表】失败", zap.Error(err))
		HandlerGrpcErrorToHttp(err, c)
		return
	}

	result := make([]interface{}, 0)
	for _, v := range rsp.Data {

		user := reponse.UserResponse{
			Id:       v.Id,
			NickName: v.NickName,
			Birthday: reponse.JsonTime(time.Time(time.Unix(int64(v.BirthDay), 0))),
		}
		result = append(result, user)
		//result = append(result, gin.H{
		//	"id": v.Id,
		//	"name": v.NickName,
		//	"birthDay": v.BirthDay,
		//	"gender":v.Gender,
		//	"mobile":v.Mobile,
		//})
		//data := make(map[string]interface{})
		//data["id"] = v.Id
		//data["name"] = v.NickName
		//data["birthDay"] = v.BirthDay
		//data["gender"] = v.Gender
		//data["mobile"] = v.Mobile

	}
	c.JSON(http.StatusOK, result)
}
