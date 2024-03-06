package models

import (
	"net"
	"net/rpc"
	"os"
	"user_service/utils/functions"
)

type Listener int

func StartRpcServer() {
	rpcPort := os.Getenv("RPC_PORT")
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+rpcPort)
	if err != nil {
		functions.ShowLog("StartRpcServerError1", err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		functions.ShowLog("StartRpcServerError2", err)
	}
	listener := new(Listener)
	err = rpc.Register(listener)
	if err != nil {
		functions.ShowLog("StartRpcServerError3", err)
	}
	go rpc.Accept(inbound)
}

func (l *Listener) RpcFindUserProfileById(userId int, reply *string) error {
	userInfo := FindUserInfoById(userId)
	userInfoJsonStr := functions.StructToJsonStr(userInfo)
	*reply = userInfoJsonStr
	return nil
}
