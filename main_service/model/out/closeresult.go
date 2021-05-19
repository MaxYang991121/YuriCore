package out

import "github.com/KouKouChan/YuriCore/main_service/constant"

func BuildCloseResultWindow() []byte {
	buf := make([]byte, 1)
	buf[0] = constant.LeaveResultWindow
	return buf
}
