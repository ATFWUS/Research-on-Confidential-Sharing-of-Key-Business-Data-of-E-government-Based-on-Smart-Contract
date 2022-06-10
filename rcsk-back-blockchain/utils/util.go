package utils

// 定义常用的错误码
const (
	RECODE_OK                  = "0"
	RECODE_UNKOWNERR           = "999"
	RECODE_USER_ALREADY_EXISTS = "100"
	RECODE_USER_NOT_EXISTS     = "101"
	RECODE_PASSWORD_ERR        = "102"
	RECODE_BALANCE_NOT_ENOUGH  = "103"
	RECODE_TASK_NOT_EXISTS     = "104"
	RECODE_TASK_CANNOT_TAKE    = "105"
	RECODE_TASK_CANNOT_COMMIT  = "106"
	RECODE_TASK_CANNOT_CONFIRM = "107"
	RECODE_CC_PUT_ERR          = "108"
	RECODE_CC_GET_ERR          = "109"
	RECODE_CC_DEL_ERR          = "110"
	RECODE_JSON_ERR            = "111"
	RECODE_AUTH_ERR            = "112"
	RECODE_INVOKE_ERR          = "113"
	RECODE_BIND_ERR            = "114"
	RECODE_CC_CALL_ERR         = "115"
	RECODE_SESSION_ERR         = "116"
)

var mapRespMsg map[string]string = map[string]string{
	RECODE_OK:                  "正常",
	RECODE_UNKOWNERR:           "未知错误",
	RECODE_USER_ALREADY_EXISTS: "用户已经存在",
	RECODE_USER_NOT_EXISTS:     "用户不存在",
	RECODE_PASSWORD_ERR:        "密码错误",
	RECODE_BALANCE_NOT_ENOUGH:  "余额不足",
	RECODE_TASK_NOT_EXISTS:     "任务不存在",
	RECODE_TASK_CANNOT_TAKE:    "任务不允许接收",
	RECODE_TASK_CANNOT_COMMIT:  "任务不允许提交",
	RECODE_TASK_CANNOT_CONFIRM: "任务不允许确认",
	RECODE_CC_PUT_ERR:          "链码PUT失败",
	RECODE_CC_GET_ERR:          "链码GET失败",
	RECODE_CC_DEL_ERR:          "链码DEL失败",
	RECODE_JSON_ERR:            "JSON处理错误",
	RECODE_AUTH_ERR:            "权限错误",
	RECODE_INVOKE_ERR:          "跨链码调用错误",
	RECODE_BIND_ERR:            "参数绑定错误",
	RECODE_CC_CALL_ERR:         "链码调用失败",
	RECODE_SESSION_ERR:         "SESSION错误",
}

// 定义链码状态
const (
	STATUS_PUBLISH = iota
	STATUS_TAKE
	STATUS_COMMIT
	STATUS_CONFIRM
)

func GetCodeMsg(code string) string {
	return mapRespMsg[code]
}
