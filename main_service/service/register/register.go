package register

import (
	"context"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/model/out"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
)

type RegisterService interface {
	Handle(ctx context.Context) error
}

type registerServiceImpl struct {
	username string
	password string
	client   net.Conn
	seq      *uint8
}

func GetRegisterServiceImpl(username, password string, client net.Conn, seq *uint8) RegisterService {
	return &registerServiceImpl{
		username: username,
		password: password,
		client:   client,
		seq:      seq,
	}
}

func (r *registerServiceImpl) Handle(ctx context.Context) error {
	ok, err := client.GetUserClient().Register(ctx, r.username, r.password)
	if err != nil {
		out.OnSendMessage(packet.GetNextSeq(r.seq), r.client, constant.MessageDialogBox, constant.CSO_Register_ServerFailed)
		return err
	}
	if !ok {
		out.OnSendMessage(packet.GetNextSeq(r.seq), r.client, constant.MessageDialogBox, constant.CSO_Register_Failed)
		return err
	}

	out.OnSendMessage(packet.GetNextSeq(r.seq), r.client, constant.MessageDialogBox, constant.CSO_Register_Success)

	return nil
}
