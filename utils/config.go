package utils

import (
	"os/user"
	"io/ioutil"
	"path"
	"strings"
	"github.com/huyinghuan/encryption/cbc"
)

var (
	CBCKey = "hcp to cloud"
)

func ReadConfig()(cloud string, id string, secretKey string, bucket string, e error) {
	current, err:=user.Current()
	if err!=nil{
		e = err
		return
	}
	content, err := ioutil.ReadFile(path.Join(current.HomeDir, ".hcp"))
	if err!=nil{
		e = err
		return
	}
	encrypt := cbc.New(CBCKey)
	keyContent, err := encrypt.DecryptString(string(content))
	if err!=nil{
		e = err
		return
	}
	keys := strings.Split(keyContent, ",")
	cloud, id, secretKey, bucket = keys[0], keys[1], keys[2], keys[3]
	return
}

func WriteConfig(cloud string, id string, secretKey string, bucket string ) error{
	// 将key加密写入用户隐藏文件
	current, err:=user.Current()
	if err!=nil{
		return err
	}
	configFilePath:=path.Join(current.HomeDir, ".hcp")
	content := strings.Join([]string{cloud, id, secretKey, bucket}, ",")
	encrypt := cbc.New(CBCKey)
	encryptStr, err:=encrypt.EncryptString(content)
	if err!=nil{
		return err
	}
	if err:=ioutil.WriteFile(configFilePath, []byte(encryptStr), 0644); err!=nil{
		return err
	}
	return nil
}