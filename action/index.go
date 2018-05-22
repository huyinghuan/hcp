package action

import (
	"github.com/urfave/cli"
	"log"
	"hcp/utils"
)

func Action(c *cli.Context) error{
	cloud, id, sk, bucket, err := utils.ReadConfig()
	if err!=nil{
		return err
	}
	// 校验密钥
	// log.Println("正在校验密钥...")
	// if err:=sina.Verify(ak, sk, bucket); err!=nil{
	//   return err
	// }
	// log.Println("校验成功")
	log.Println(cloud, id, sk, bucket)
	log.Println(c.Bool("random"))
	c.Args()

	return nil
}


