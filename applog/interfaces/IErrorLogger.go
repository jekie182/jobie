package interfaces

import "appdomain/result/interfaces"

type IErrorLogger interface {
	func WriteLog(result *IResult)
}
