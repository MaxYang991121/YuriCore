package option

import (
	"context"
	"errors"
	"net"

	"github.com/KouKouChan/YuriCore/main_service/client"
	"github.com/KouKouChan/YuriCore/main_service/model/packet"
	"github.com/KouKouChan/YuriCore/utils"
)

type OptionService interface {
	Handle(ctx context.Context) error
}

type optionServiceImpl struct {
	data   []byte
	client net.Conn
}

func NewOptionService(p *packet.PacketData, client net.Conn) OptionService {
	len := utils.ReadUint16(p.Data, &p.CurOffset)
	return &optionServiceImpl{
		data:   utils.ReadString(p.Data, &p.CurOffset, int(len)),
		client: client,
	}
}

func (s *optionServiceImpl) Handle(ctx context.Context) error {
	u := client.GetUserCacheClient().GetUserByConnection(ctx, s.client)
	if u == nil {
		return errors.New("can't find user")
	}

	_, err := client.GetUserClient().UpdateOptions(ctx, u.UserID, s.data)

	return err
}
