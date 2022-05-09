package errcode

var (
	StatusOK         = NewCode(1001, "OK")
	ErrSave          = NewCode(1002, "保存失败")
	ErrQuery         = NewCode(1003, "查询失败")
	ErrLengthOver    = NewCode(1004, "长度超过限制")
	ErrLengthZero    = NewCode(1005, "长度为零")
	ErrParamNotValid = NewCode(1006, "参数绑定有误")
	NotFound         = NewCode(1007, "没找到")
)
