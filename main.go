package main

import (
  "io/ioutil"
  "log"
  "os"
  "os/user"
  "path"
  "sort"
  "strings"

  "github.com/huyinghuan/encryption/cbc"
  "github.com/urfave/cli"
  "hcp/action"
)

type Sina struct{

}
var (
  AccessKey string
  SecretKey string
  Bucket string
  CBCKey = "hcp to cloud"

)


func getCommonFlags() []cli.Flag{
  // 关键参数
  keysArgs :=[]cli.Flag{
    cli.StringFlag{
      Name: "ak, A",
      Value: "",
      Usage: "Access Key",
      Destination: &AccessKey,
    },
    cli.StringFlag{
      Name: "sk, S",
      Value: "",
      Usage: "Secret Key",
      Destination: &SecretKey,
    },
    cli.StringFlag{
      Name: "bucket, B",
      Value: "",
      Usage: "bucket name, 默认Bucket",
      Destination: &Bucket,
    },
  }
  return keysArgs
}

func readConfig()(ak string, sk string, b string, e error) {
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
  ak, sk, b =keys[0], keys[1], keys[2]
  return
}


func appAction(c *cli.Context) error{
  ak, sk, bucket, err := readConfig()
  if err!=nil{
    return err
  }
  // 校验密钥
  // log.Println("正在校验密钥...")
  // if err:=sina.Verify(ak, sk, bucket); err!=nil{
  //   return err
  // }
  // log.Println("校验成功")
  log.Println(ak, sk, bucket)
  log.Println(c.Bool("random"))
  c.Args()

  return nil
}

func main() {
  app := cli.NewApp()

  app.Name = "hcp"
  app.Usage = `上传文件到新浪云存储
    hcp sourceFilePath cloudFilePath
    hcp sourceFileArray  cloudDir
    e.g
      hcp a.js a.js
      hcp test/*  test    //test文件夹下到所有文件 上传到 sina云上到test文件夹
      hcp a.js b.js c.js test //上传a.js,b.js, c.js 到 test文件夹下
      //以上都是上传到 init 命令时设置的默认 buckets空间下
      其他参数见help
  `
  app.Version = "1.0.0"

  app.Flags = []cli.Flag{
    cli.BoolFlag{
      Name:"random, r",
      Usage: "上传时使用随机文件名，避免覆盖",
    },
    cli.BoolFlag{
      Name:"upload, u",
      Usage: `命令最后一个参数不作为远程文件夹，远程文件名使用，上传文件到bucket根目录
        如：
          hcp -u a.js a/b.js 
        将上传 a.js 和 b.js到远程根目录 xxx/a.js , xxx/a/b.js
      `,
    },
    cli.BoolFlag{
      Name:"first, f",
      Usage: `所有文件到上传到远程根目录
        如：
          hcp -f a.js a/b.js 
        将上传 a.js 和 b.js到远程根目录 xxx/a.js , xxx/b.js
      `,
    },
  }

  app.Commands = []cli.Command{
    {
      Name:    "init",
      Usage:   "初始化密钥，到本地目录",
      Action:   action.InitAction,
    },
  }
  app.Action = appAction
  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}