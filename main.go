package main

import (
  "log"
  "os"
  "sort"
  "github.com/urfave/cli"
  "hcp/action"
)



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
  app.Action = action.Action
  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}