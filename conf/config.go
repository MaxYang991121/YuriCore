package conf

import (
	"fmt"

	"github.com/KouKouChan/YuriCore/utils"
)

type CSOConf struct {
	MainPort         uint32
	UserAdress       string
	UserPort         uint32
	RoomAdress       string
	RoomPort         uint32
	MaxUsers         uint32
	UnlockAllWeapons uint32

	EnableDataBase uint32
	DBUserName     string
	DBpassword     string
	DBaddress      string
	DBport         string

	DebugLevel uint32
	LogFile    uint32

	EnableRegister uint32
	EnableMail     uint32
	REGPort        uint32
	REGEmail       string
	REGPassWord    string
	REGSMTPaddr    string

	CodePage string
}

var (
	Config CSOConf
)

func (conf *CSOConf) InitConf(path string) {
	if conf == nil {
		return
	}
	fmt.Printf("Reading configure file ...\n")
	ini_parser := utils.IniParser{}
	file := path
	if err := ini_parser.LoadIni(file); err != nil {
		fmt.Printf("Loading config file error[%s]\n", err.Error())
		fmt.Printf("Using default data ...\n")
		Config.EnableDataBase = 1
		Config.DBUserName = "root"
		Config.DBpassword = "123456"
		Config.DBaddress = "localhost"
		Config.DBport = "3306"

		Config.MaxUsers = 0
		Config.UnlockAllWeapons = 1
		Config.MainPort = 30001
		Config.UserPort = 30002
		Config.RoomPort = 30003
		Config.UserAdress = "127.0.0.1"
		Config.RoomAdress = "127.0.0.1"

		Config.DebugLevel = 2
		Config.LogFile = 1

		Config.EnableRegister = 0
		Config.EnableMail = 0
		Config.CodePage = "gbk"
		return
	}

	Config.EnableDataBase = ini_parser.IniGetUint32("Database", "EnableDataBase")
	Config.DBUserName = ini_parser.IniGetString("Database", "DBUserName")
	Config.DBpassword = ini_parser.IniGetString("Database", "DBpassword")
	Config.DBaddress = ini_parser.IniGetString("Database", "DBaddress")
	Config.DBport = ini_parser.IniGetString("Database", "DBport")

	Config.MaxUsers = ini_parser.IniGetUint32("Server", "MaxUsers")
	if conf.MaxUsers < 0 {
		conf.MaxUsers = 0
	}
	Config.UnlockAllWeapons = ini_parser.IniGetUint32("Server", "UnlockAllWeapons")
	Config.MainPort = ini_parser.IniGetUint32("Server", "MainPort")
	Config.UserAdress = ini_parser.IniGetString("Server", "UserAdress")
	Config.UserPort = ini_parser.IniGetUint32("Server", "UserPort")
	Config.RoomAdress = ini_parser.IniGetString("Server", "RoomAdress")
	Config.RoomPort = ini_parser.IniGetUint32("Server", "RoomPort")

	Config.DebugLevel = ini_parser.IniGetUint32("Debug", "DebugLevel")
	if conf.DebugLevel > 2 || conf.DebugLevel < 0 {
		conf.DebugLevel = 2
	}
	Config.LogFile = ini_parser.IniGetUint32("Debug", "LogFile")

	Config.EnableRegister = ini_parser.IniGetUint32("Register", "EnableRegister")
	Config.EnableMail = ini_parser.IniGetUint32("Register", "EnableMail")
	Config.REGPort = ini_parser.IniGetUint32("Register", "REGPort")
	Config.REGEmail = ini_parser.IniGetString("Register", "REGEmail")
	Config.REGPassWord = ini_parser.IniGetString("Register", "REGPassWord")
	Config.REGSMTPaddr = ini_parser.IniGetString("Register", "REGSMTPaddr")

	Config.CodePage = ini_parser.IniGetString("Encode", "CodePage")
}
