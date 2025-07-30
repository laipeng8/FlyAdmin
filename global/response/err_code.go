package response

import (
	"fmt"
	"strings"
)

var (
	WhiteImageList = []string{

		"jpg", "jpeg", "png", "jpeg", "bmp", "ico", "tiff", "svg", "webp",
	}
	HeadImgWhiteList = []string{
		"jpg", "jpeg", "png", "jpeg", "bmp", "ico", "webp",
	}
	FileSize     = 5
	FileSavePath = "static/"
)

type ErrorCode int

const (
	SettingsError            ErrorCode = 300 // 系统错误
	ArgumentError                      = 301
	DatabaseError                      = 302
	TransactionError                   = 303
	UpdateDaTaError                    = 304
	SelectDataError                    = 305
	CommitDataError                    = 306
	DeleteDataError                    = 307
	AddDataError                       = 308
	AuthorizationLockError             = 309
	ParseJWTError                      = 310
	ScanDataError                      = 311
	WXParamError                       = 312
	FileError                          = 313
	GetUserIDError                     = 314
	JudgeFileError                     = 315
	FileSizeTransborderError           = 316
	FileSaVeError                      = 317
	AffectedDataError                  = 318
	PathArgumentError                  = 319
	DataIsNotNullError                 = 320
	InsertDataBaseError                = 321
	EchoDataBaseIDError                = 322
	LoginError                         = 323
	UserIsNotExistError                = 324
	JwtGenerateError                   = 325
	LastInsertIdError                  = 326
	AlreadyBindShop                    = 327
	AlreadyBindTheme                   = 328
	ParseError                         = 329
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:            "系统错误",
		ArgumentError:            "参数错误",
		DatabaseError:            "数据库错误",
		TransactionError:         "事务开启失败",
		UpdateDaTaError:          "数据更新错误",
		SelectDataError:          "查询数据失败",
		CommitDataError:          "数据提交失败",
		DeleteDataError:          "删除数据错误",
		AddDataError:             "增加数据错误",
		AuthorizationLockError:   "请求头Authorization缺失",
		ParseJWTError:            "无效的jwt令牌",
		ScanDataError:            "数据绑定失败",
		WXParamError:             "微信用户参数错误",
		FileError:                "文件上传错误",
		GetUserIDError:           "解析userID错误",
		JudgeFileError:           fmt.Sprintf("非法文件，文件只允许:%s", strings.Join(HeadImgWhiteList, ",")),
		FileSizeTransborderError: fmt.Sprintf("文件大小越界，文件应该小于，%dMB", FileSize),
		FileSaVeError:            "文件保存错误",
		AffectedDataError:        "数据更改失败，未对数据库产生影响",
		PathArgumentError:        "路径参数为空串出错误",
		DataIsNotNullError:       "数据不能为空",
		InsertDataBaseError:      "数据插入错误",
		EchoDataBaseIDError:      "数据库ID回显错误",
		LoginError:               "登录失败",
		UserIsNotExistError:      "用户不存在错误",
		JwtGenerateError:         "jwt令牌生成错误",
		LastInsertIdError:        "获取回显主键错误",
		AlreadyBindShop:          "已经绑定了商品",
		AlreadyBindTheme:         "已经有社区绑定",
		ParseError:               "解析错误",
	}
)
