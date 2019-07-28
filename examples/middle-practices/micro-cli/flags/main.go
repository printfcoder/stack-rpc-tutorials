package main

import (
	"fmt"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		// 我们可以通过下面的方式给我们的服务增加flag参数
		micro.Flags(
			cli.StringFlag{
				Name:  "string_flag",
				Usage: "string_flag是字符串flag",
				Value: "我要被覆盖",
			},
			cli.IntFlag{
				Name:  "int_flag",
				Usage: "int_flag是整形flag",
				// 可以使用EnvVar来声明支持使用环境变量，但是如果flag和env同时使用时，以flag为准
				EnvVar: "INT_FLAG",
			},
			cli.BoolFlag{
				Name:  "bool_flag",
				Usage: "bool_flag是布尔值flag",
			},
			cli.StringFlag{
				Name:  "string_flag_default",
				Usage: "我是缺省值",
				Value: "我是缺省值",
			},
		),
	)

	// Init 初始化方法会解析flag。在命令行中传入的参数会覆盖上面定义的默认值
	// 而在下面的选项Option中定义的值会覆盖命令行中传入的值。
	service.Init(
		micro.Action(func(c *cli.Context) {
			fmt.Printf("字符串flag值： %s\n", c.String("string_flag"))
			fmt.Printf("字符串缺省值： %s\n", c.String("string_flag_default"))
			fmt.Printf("整形flag值：%d\n", c.Int("int_flag"))
			fmt.Printf("布尔值flag值：%t\n", c.Bool("bool_flag"))

			// 打印完后退出
			os.Exit(0)
		}),
	)

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
