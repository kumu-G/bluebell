package main

import "fmt"

/**
 * @Description
 * @Author 枯木.
 * @Date 2023/3/5 13:44
 **/
func main() {
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
}
