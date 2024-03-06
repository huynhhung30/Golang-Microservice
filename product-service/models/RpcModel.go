package models

import (
	"encoding/json"
	"net/rpc"
	"os"
	"product_service/utils/functions"
)

func RpcGetUserInfoById(user_id int) (userInfo *UserModel) {
	host := os.Getenv("DB_HOST")
	useRpcPort := os.Getenv("USER_RPC_PORT")
	client, err := rpc.Dial("tcp", host+":"+useRpcPort)
	if err != nil {
		functions.ShowLog("RpcGetUserInfoByIdError1", err)
		return
	}
	reply := ""
	err = client.Call("Listener.RpcFindUserProfileById", user_id, &reply)
	if err != nil {
		functions.ShowLog("RpcGetUserInfoByIdError2", err.Error())
		return
	}
	json.Unmarshal([]byte(reply), &userInfo)
	return userInfo
}
