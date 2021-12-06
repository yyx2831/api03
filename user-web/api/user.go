package api

import (
	"api03/user-web/forms"
	"api03/user-web/global/reponse"
	"api03/user-web/proto"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

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
	//获取前端get请求
	//page := c.DefaultQuery("pn","0")
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	pageInt, err := strconv.Atoi(page)
	pageSizeInt, err := strconv.Atoi(pageSize)
	userSrvClient := proto.NewUserClient(userConn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pageInt),
		PSize: uint32(pageSizeInt),
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
	}
	c.JSON(http.StatusOK, result)
}

func PassWordLogin(c *gin.Context) {
	//表单验证
	PassWordLoginForm := forms.PassWordLoginForm{}
	err := c.ShouldBind(&PassWordLoginForm)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(validate)),
		})
	}
}
