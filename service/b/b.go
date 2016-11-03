package b

import (
	"test_ctx/service/base"
)

type B struct{
	base.BaseService
}


func (*B) DoB() int {
	return 1
}