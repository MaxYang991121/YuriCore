package user_service

import (
	"context"
	"errors"
	"sync"

	"github.com/KouKouChan/YuriCore/main_service/constant"
	"github.com/KouKouChan/YuriCore/main_service/gen-go/yuricore/user_service"
	"github.com/KouKouChan/YuriCore/main_service/model/convert"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	. "github.com/KouKouChan/YuriCore/verbose"
	"github.com/apache/thrift/lib/go/thrift"
)

type UserServiceImpl struct {
	client *user_service.UserServiceClient
	lock   sync.Mutex
}

func NewUserService(addr string) *UserServiceImpl {
	var transport thrift.TTransport
	var err error
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		panic(err)
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		panic(err)
	}
	if err := transport.Open(); err != nil {
		panic(err)
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return &UserServiceImpl{
		client: user_service.NewUserServiceClient(thrift.NewTStandardClient(iprot, oprot)),
		lock:   sync.Mutex{},
	}
}

func (u *UserServiceImpl) Login(ctx context.Context, username, password string) (*user.UserInfo, int8) {
	req := user_service.NewLoginRequest()
	req.UserName = username
	req.PassWord = password

	DebugPrintf(2, "Call Login req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.Login(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return nil, constant.Login_RPC_ERROR
	}
	DebugPrintf(2, "Call Login resp=%+v", resp)

	if resp == nil {
		return nil, constant.Login_NULL_Resp
	}

	return convert.ConvertToUserInfo(resp.UserInfo), resp.StatusCode
}

func (u *UserServiceImpl) Register(ctx context.Context, username, password string) (bool, error) {
	req := user_service.NewRegisterRequest()
	req.UserName = username
	req.NickName = ""
	req.PassWord = password

	DebugPrintf(2, "Call Register req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.Register(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return false, err
	}
	DebugPrintf(2, "Call Register resp=%+v", resp)

	if resp == nil {
		return false, errors.New("null response")
	}

	return resp.Success, nil
}

func (u *UserServiceImpl) GetUserInfo(ctx context.Context, userID uint32) (*user.UserInfo, error) {
	req := user_service.NewGetUserInfoRequest()
	req.UserID = int32(userID)

	DebugPrintf(2, "Call GetUserInfo req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.GetUserInfo(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call GetUserInfo resp=%+v", resp)

	if resp == nil {
		return nil, errors.New("got null resp!")
	}

	if resp.UserInfo == nil || !resp.Success || resp.UserInfo.UserID == 0 {
		return nil, errors.New("got null resp data!")
	}

	return convert.ConvertToUserInfo(resp.UserInfo), nil
}

func (u *UserServiceImpl) UserDown(ctx context.Context, userID uint32) (bool, error) {
	req := user_service.NewUserDownRequest()
	req.UserID = int32(userID)

	DebugPrintf(2, "Call UserDown req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.UserDown(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return false, err
	}
	DebugPrintf(2, "Call UserDown resp=%+v", resp)

	if resp == nil || !resp.Success {
		return false, errors.New("got wrong resp!")
	}

	return true, nil
}

func (u *UserServiceImpl) GetUserFriends(ctx context.Context, userID uint32) ([]user.UserInfo, error) {
	return nil, nil
}

func (u *UserServiceImpl) AddUserPoints(ctx context.Context, userID, add uint32) (uint32, error) {
	return 0, nil
}

func (u *UserServiceImpl) AddUserCash(ctx context.Context, userID, add uint32) (uint32, error) {
	return 0, nil
}

func (u *UserServiceImpl) UserPlayedGame(ctx context.Context, userID, IsWin, Kills, Deaths, HeadShots uint32) (*user.UserInfo, error) {
	return nil, nil
}

func (u *UserServiceImpl) UserPayPoints(ctx context.Context, userID, used uint32) (uint32, error) {
	return 0, nil
}

func (u *UserServiceImpl) UserPayCash(ctx context.Context, userID, used uint32) (uint32, error) {
	return 0, nil
}

func (u *UserServiceImpl) UserAddItem(ctx context.Context, userID, item uint32) (*user.UserInfo, error) {
	return nil, nil
}

func (u *UserServiceImpl) UserAddFriend(ctx context.Context, userID, friendID uint32) (*user.UserInfo, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateBag(ctx context.Context, UserID uint32, BagID uint16, Slot uint8, ItemID uint16) (*user.UserInfo, error) {
	req := user_service.NewUpdateBagRequest()
	req.UserID = int32(UserID)
	req.BagID = int8(BagID)
	req.Slot = int8(Slot)
	req.ItemID = int16(ItemID)

	DebugPrintf(2, "Call UpdateBag req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.UpdateBag(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call UpdateBag resp=%+v", resp)

	if resp == nil {
		return nil, errors.New("got null resp!")
	}

	if resp.UserInfo == nil || !resp.Success || resp.UserInfo.UserID == 0 {
		return nil, errors.New("got null resp data!")
	}

	return convert.ConvertToUserInfo(resp.UserInfo), nil
}

func (u *UserServiceImpl) UpdateBuymenu(ctx context.Context, UserID uint32, BuymenuID uint16, Slot uint8, ItemID uint16) (*user.UserInfo, error) {
	req := user_service.NewUpdateBuymenuRequest()
	req.UserID = int32(UserID)
	req.BuymenuID = int8(BuymenuID)
	req.Slot = int8(Slot)
	req.ItemID = int16(ItemID)

	DebugPrintf(2, "Call UpdateBuymenu req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.UpdateBuymenu(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call UpdateBuymenu resp=%+v", resp)

	if resp == nil {
		return nil, errors.New("got null resp!")
	}

	if resp.UserInfo == nil || !resp.Success || resp.UserInfo.UserID == 0 {
		return nil, errors.New("got null resp data!")
	}

	return convert.ConvertToUserInfo(resp.UserInfo), nil
}

func (u *UserServiceImpl) UpdateCosmetics(ctx context.Context, UserID uint32, CosmeticsID uint8, cosmetics *user.UserCosmetics) (*user.UserInfo, error) {
	req := user_service.NewUpdateCosmeticsRequest()
	req.UserID = int32(UserID)
	req.CosmeticsID = int8(CosmeticsID)
	req.CosmeticsName = cosmetics.CosmeticsName
	req.MainWeapon = int16(cosmetics.MainWeapon)
	req.MainBullet = int16(cosmetics.MainBullet)
	req.SecondWeapon = int16(cosmetics.SecondWeapon)
	req.SecondBullet = int16(cosmetics.SecondBullet)
	req.FlashbangNum = int16(cosmetics.FlashbangNum)
	req.GrenadeID = int16(cosmetics.GrenadeID)
	req.SmokeNum = int16(cosmetics.SmokeNum)
	req.DefuserNum = int16(cosmetics.DefuserNum)
	req.TelescopeNum = int16(cosmetics.TelescopeNum)
	req.BulletproofNum = int16(cosmetics.BulletproofNum)
	req.KnifeID = int16(cosmetics.KnifeID)

	DebugPrintf(2, "Call UpdateCosmetics req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.UpdateCosmetics(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call UpdateCosmetics resp=%+v", resp)

	if resp == nil {
		return nil, errors.New("got null resp!")
	}

	if resp.UserInfo == nil || !resp.Success || resp.UserInfo.UserID == 0 {
		return nil, errors.New("got null resp data!")
	}

	return convert.ConvertToUserInfo(resp.UserInfo), nil
}

func (u *UserServiceImpl) UpdateCampaign(ctx context.Context, UserID uint32, CampaignID uint8) (*user.UserInfo, error) {
	req := user_service.NewUpdateCampaignRequest()
	req.UserID = int32(UserID)
	req.Campaign = int8(CampaignID)

	DebugPrintf(2, "Call UpdateCampaign req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.UpdateCampaign(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call UpdateCampaign resp=%+v", resp)

	if resp == nil {
		return nil, errors.New("got null resp!")
	}

	if resp.UserInfo == nil || !resp.Success || resp.UserInfo.UserID == 0 {
		return nil, errors.New("got null resp data!")
	}

	return convert.ConvertToUserInfo(resp.UserInfo), nil
}
func (u *UserServiceImpl) UpdateOptions(ctx context.Context, UserID uint32, options []byte) (*user.UserInfo, error) {
	req := user_service.NewUpdateOptionsRequest()
	req.UserID = int32(UserID)
	req.Options = options

	DebugPrintf(2, "Call UpdateOptions req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.UpdateOptions(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call UpdateOptions resp=%+v", resp)

	if resp == nil {
		return nil, errors.New("got null resp!")
	}

	if resp.UserInfo == nil || !resp.Success || resp.UserInfo.UserID == 0 {
		return nil, errors.New("got null resp data!")
	}

	return convert.ConvertToUserInfo(resp.UserInfo), nil
}

func (u *UserServiceImpl) UpdateNickname(ctx context.Context, UserID uint32, nickname string) (*user.UserInfo, error) {
	req := user_service.NewUpdateNickNameRequest()
	req.UserID = int32(UserID)
	req.NickName = nickname

	DebugPrintf(2, "Call UpdateNickname req=%+v", req)
	u.lock.Lock()
	resp, err := u.client.UpdateNickName(ctx, req)
	u.lock.Unlock()
	if err != nil {
		return nil, err
	}
	DebugPrintf(2, "Call UpdateNickname resp=%+v", resp)

	if resp == nil {
		return nil, errors.New("got null resp!")
	}

	if resp.UserInfo == nil || !resp.Success || resp.UserInfo.UserID == 0 {
		return nil, errors.New("got null resp data!")
	}

	return convert.ConvertToUserInfo(resp.UserInfo), nil
}
