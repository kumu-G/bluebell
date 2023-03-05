package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/router"
	"bluebell/settings"
	"fmt"
)

/**
 * @Description
 * @Author 枯木.
 * @Date 2023/3/5 13:44
 **/
func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("need config file.eg: bluebell config.yaml")
	//	return
	//}
	// 加载配置
	if err := settings.Init("./conf/config.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 注册路由
	r := router.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
