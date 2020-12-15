package e

var MsgFlags = map[int]string {
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorExistTag:              "已存在该标签名称",
	ErrorNotExistTag:           "该标签不存在",
	ErrorNotExistArticle:      "该文章不存在",
	ErrorAuthCheckTokenFail:    "token鉴权失败",
	ErrorAuthCheckTokenTimeout: "token已超时",
	ErrorAuth:                  "token错误",
	ErrorAuthToken:             "token生成失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
