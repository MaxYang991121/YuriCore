package main

import (
	"context"
	"fmt"

	"github.com/KouKouChan/YuriCore/conf"
	"github.com/KouKouChan/YuriCore/main_service/gen-go/yuricore/user_service"
	"github.com/KouKouChan/YuriCore/main_service/model/convert"
	"github.com/KouKouChan/YuriCore/main_service/model/user"
	"github.com/KouKouChan/YuriCore/user_service/controller"
	. "github.com/KouKouChan/YuriCore/verbose"
	"github.com/apache/thrift/lib/go/thrift"
)

type UserServiceImpl struct{}

func (l *UserServiceImpl) Login(ctx context.Context, req *user_service.LoginRequest) (*user_service.LoginResponse, error) {
	DebugPrintf(2, "Login called req=%+v", req)
	resp := user_service.NewLoginResponse()
	defer DebugPrintf(2, "Get Login resp=%+v", resp)

	info, code := controller.NewLoginController(req.GetUserName(), req.GetPassWord()).Handle(ctx)

	resp.StatusCode = code
	resp.UserInfo = convert.ConvertToServiceUser(info)

	return resp, nil
}

func (l *UserServiceImpl) Register(ctx context.Context, req *user_service.RegisterRequest) (*user_service.RegisterResponse, error) {
	DebugPrintf(2, "Register called req=%+v", req)
	resp := user_service.NewRegisterResponse()
	defer DebugPrintf(2, "Get Register resp=%+v", resp)

	err := controller.NewRegisterController(req.UserName, req.NickName, req.PassWord).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "Register err=%+v", err)
		resp.Success = false

		return resp, nil
	}

	resp.Success = true

	return resp, nil
}

func (l *UserServiceImpl) GetUserInfo(ctx context.Context, req *user_service.GetUserInfoRequest) (*user_service.GetUserInfoResponse, error) {
	DebugPrintf(2, "GetUserInfo called req=%+v", req)
	resp := user_service.NewGetUserInfoResponse()
	defer DebugPrintf(2, "Get GetUserInfo resp=%+v", resp)

	info, err := controller.NewUserInfoController(uint32(req.GetUserID())).Handle(ctx)
	if err != nil || info == nil {
		DebugPrintf(2, "GetUserInfo err=%+v", err)
		resp.Success = false
		resp.UserInfo = nil

		return resp, nil
	}

	resp.Success = true
	resp.UserInfo = convert.ConvertToServiceUser(info)

	return resp, nil
}

func (l *UserServiceImpl) UserDown(ctx context.Context, req *user_service.UserDownRequest) (*user_service.UserDownResponse, error) {
	DebugPrintf(2, "UserDown called req=%+v", req)
	resp := user_service.NewUserDownResponse()
	defer DebugPrintf(2, "Get UserDown resp=%+v", resp)

	err := controller.NewDownController(uint32(req.GetUserID())).Handle(ctx)
	if err != nil {
		DebugPrintf(2, "UserDown err=%+v", err)
		resp.Success = false

		return resp, nil
	}

	resp.Success = true

	return resp, nil
}

func (l *UserServiceImpl) GetUserFriends(ctx context.Context, req *user_service.GetUserFriendsRequest) (*user_service.GetUserFriendsResponse, error) {
	DebugPrintf(2, "GetUserFriends called req=%+v", req)
	resp := user_service.NewGetUserFriendsResponse()
	defer DebugPrintf(2, "Get GetUserFriends resp=%+v", resp)

	infos, err := controller.NewFriendsController(uint32(req.GetUserID())).Handle(ctx)
	if err != nil || infos == nil {
		DebugPrintf(2, "GetUserFriends err=%+v", err)
		resp.Success = false
		resp.Friends = nil

		return resp, nil
	}

	resp.Success = true
	resp.Friends = []*user_service.UserInfo{}
	for i := range infos {
		resp.Friends = append(resp.Friends, convert.ConvertToServiceUser(&infos[i]))
	}

	return resp, nil
}

func (l *UserServiceImpl) AddUserPoints(ctx context.Context, req *user_service.AddUserPointsRequest) (*user_service.AddUserPointsResponse, error) {
	DebugPrintf(2, "AddUserPoints called req=%+v", req)
	resp := user_service.NewAddUserPointsResponse()
	defer DebugPrintf(2, "Get AddUserPoints resp=%+v", resp)

	return resp, nil
}

func (l *UserServiceImpl) AddUserCash(ctx context.Context, req *user_service.AddUserCashRequest) (*user_service.AddUserCashResponse, error) {
	DebugPrintf(2, "AddUserCash called req=%+v", req)
	return nil, nil
}

func (l *UserServiceImpl) UserPlayedGame(ctx context.Context, req *user_service.UserPlayedGameRequest) (*user_service.UserPlayedGameResponse, error) {
	DebugPrintf(2, "UserPlayedGame called req=%+v", req)
	return nil, nil
}

func (l *UserServiceImpl) UserPayPoints(ctx context.Context, req *user_service.UserPayPointsRequest) (*user_service.UserPayPointsResponse, error) {
	DebugPrintf(2, "UserPayPoints called req=%+v", req)
	return nil, nil
}

func (l *UserServiceImpl) UserPayCash(ctx context.Context, req *user_service.UserPayCashRequest) (*user_service.UserPayCashResponse, error) {
	DebugPrintf(2, "UserPayCash called req=%+v", req)
	return nil, nil
}

func (l *UserServiceImpl) UserAddItem(ctx context.Context, req *user_service.UserAddItemRequest) (*user_service.UserAddItemResponse, error) {
	DebugPrintf(2, "UserAddItem called req=%+v", req)
	return nil, nil
}

func (l *UserServiceImpl) UserAddFriend(ctx context.Context, req *user_service.UserAddFriendRequest) (*user_service.UserAddFriendResponse, error) {
	DebugPrintf(2, "UserAddFriend called req=%+v", req)
	return nil, nil
}

func (l *UserServiceImpl) UpdateBag(ctx context.Context, req *user_service.UpdateBagRequest) (*user_service.UpdateBagResponse, error) {
	DebugPrintf(2, "UpdateBag called req=%+v", req)
	resp := user_service.NewUpdateBagResponse()
	defer DebugPrintf(2, "Get UpdateBag resp=%+v", resp)

	info, err := controller.NewUpdateBagController(uint32(req.UserID), uint16(req.BagID), uint8(req.Slot), uint16(req.ItemID)).Handle(ctx)
	if err != nil || info == nil {
		DebugPrintf(2, "UpdateBag err=%+v", err)
		resp.Success = false
		resp.UserInfo = nil

		return resp, nil
	}

	resp.Success = true
	resp.UserInfo = convert.ConvertToServiceUser(info)

	return resp, nil
}

func (l *UserServiceImpl) UpdateBuymenu(ctx context.Context, req *user_service.UpdateBuymenuRequest) (*user_service.UpdateBuymenuResponse, error) {
	DebugPrintf(2, "UpdateBuymenu called req=%+v", req)
	resp := user_service.NewUpdateBuymenuResponse()
	defer DebugPrintf(2, "Get UpdateBuymenu resp=%+v", resp)

	info, err := controller.NewUpdateBuymenuController(uint32(req.UserID), uint16(req.BuymenuID), uint8(req.Slot), uint16(req.ItemID)).Handle(ctx)
	if err != nil || info == nil {
		DebugPrintf(2, "UpdateBuymenu err=%+v", err)
		resp.Success = false
		resp.UserInfo = nil

		return resp, nil
	}

	resp.Success = true
	resp.UserInfo = convert.ConvertToServiceUser(info)

	return resp, nil
}

func (l *UserServiceImpl) UpdateCosmetics(ctx context.Context, req *user_service.UpdateCosmeticsRequest) (*user_service.UpdateCosmeticsResponse, error) {
	DebugPrintf(2, "UpdateCosmetics called req=%+v", req)
	resp := user_service.NewUpdateCosmeticsResponse()
	defer DebugPrintf(2, "Get UpdateCosmetics resp=%+v", resp)

	info, err := controller.NewUpdateCosmeticsController(
		uint32(req.UserID),
		uint8(req.CosmeticsID),
		&user.UserCosmetics{
			CosmeticsName:  req.CosmeticsName,
			MainWeapon:     uint16(req.MainWeapon),
			MainBullet:     uint16(req.MainBullet),
			SecondWeapon:   uint16(req.SecondWeapon),
			SecondBullet:   uint16(req.SecondBullet),
			FlashbangNum:   uint16(req.FlashbangNum),
			GrenadeID:      uint16(req.GrenadeID),
			SmokeNum:       uint16(req.SmokeNum),
			DefuserNum:     uint16(req.DefuserNum),
			TelescopeNum:   uint16(req.TelescopeNum),
			BulletproofNum: uint16(req.BulletproofNum),
			KnifeID:        uint16(req.KnifeID),
		},
	).Handle(ctx)
	if err != nil || info == nil {
		DebugPrintf(2, "UpdateCosmetics err=%+v", err)
		resp.Success = false
		resp.UserInfo = nil

		return resp, nil
	}

	resp.Success = true
	resp.UserInfo = convert.ConvertToServiceUser(info)

	return resp, nil
}

func (l *UserServiceImpl) UpdateCampaign(ctx context.Context, req *user_service.UpdateCampaignRequest) (*user_service.UpdateCampaignResponse, error) {
	DebugPrintf(2, "UpdateCampaign called req=%+v", req)
	resp := user_service.NewUpdateCampaignResponse()
	defer DebugPrintf(2, "Get UpdateCampaign resp=%+v", resp)

	info, err := controller.NewUpdateCampaignController(
		uint32(req.UserID),
		uint8(req.Campaign)).Handle(ctx)

	if err != nil || info == nil {
		DebugPrintf(2, "UpdateCampaign err=%+v", err)
		resp.Success = false
		resp.UserInfo = nil

		return resp, nil
	}

	resp.Success = true
	resp.UserInfo = convert.ConvertToServiceUser(info)

	return resp, nil
}

func (l *UserServiceImpl) UpdateOptions(ctx context.Context, req *user_service.UpdateOptionsRequest) (*user_service.UpdateOptionsResponse, error) {
	DebugPrintf(2, "UpdateOptions called req=%+v", req)
	resp := user_service.NewUpdateOptionsResponse()
	defer DebugPrintf(2, "Get UpdateOptions resp=%+v", resp)

	info, err := controller.NewUpdateOptionController(
		uint32(req.UserID),
		req.Options).Handle(ctx)

	if err != nil || info == nil {
		DebugPrintf(2, "UpdateOptions err=%+v", err)
		resp.Success = false
		resp.UserInfo = nil

		return resp, nil
	}

	resp.Success = true
	resp.UserInfo = convert.ConvertToServiceUser(info)

	return resp, nil
}

func (l *UserServiceImpl) UpdateNickName(ctx context.Context, req *user_service.UpdateNickNameRequest) (*user_service.UpdateNickNameResponse, error) {
	DebugPrintf(2, "UpdateNickName called req=%+v", req)
	resp := user_service.NewUpdateNickNameResponse()
	defer DebugPrintf(2, "Get UpdateNickName resp=%+v", resp)

	info, err := controller.NewNickNameController(
		uint32(req.UserID),
		req.NickName).Handle(ctx)

	if err != nil || info == nil {
		DebugPrintf(2, "UpdateNickName err=%+v", err)
		resp.Success = false
		resp.UserInfo = nil

		return resp, nil
	}

	resp.Success = true
	resp.UserInfo = convert.ConvertToServiceUser(info)

	return resp, nil
}

func initServer(path string, addr string) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		panic(err)
	}

	handler := &UserServiceImpl{}
	processor := user_service.NewUserServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("user service is running at", fmt.Sprintf("%s:%d", conf.Config.UserAdress, conf.Config.UserPort))
	server.Serve()
}
