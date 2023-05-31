package utils

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	// 建立连接池
	pool = &redis.Pool{
		// Maximum number of connections allocated by the pool at a given time.
		// When zero, there is no limit on the number of connections in the pool.
		//最大活跃连接数，0代表无限
		MaxActive: 888,
		//最大闲置连接数
		// Maximum number of idle connections in the pool.
		MaxIdle: 20,
		//闲置连接的超时时间
		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: time.Second * 100,
		//定义拨号获得连接的函数
		// Dial is an application supplied function for creating and configuring a
		// connection.
		//
		// The connection returned from Dial must not be in a special state
		// (subscribed to pubsub channel, transaction started, ...).
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	// //延迟关闭连接池
	// defer pool.Close()

}

// 根据cookie中的tempid查询对应的uid，假如没有的话就说明没登陆
func AlreadyLogin(tempid string) string {
	conn := pool.Get()
	defer conn.Close()
	// is_email_exit, _ := redis.Bool(conn.Do("EXISTS", tempid))
	uid, _ := redis.String(conn.Do("Get", tempid))

	return uid
}

// 登陆成功之后，给用户生成tempid存入cookie记住登陆状态，维持时间7天
func SetTempId(uid string) (tempid string) {
	//通过连接池获得连接
	conn := pool.Get()
	//延时关闭连接
	defer conn.Close()
	// 生成tempid
	salt := time.Now().String()
	tempid = fmt.Sprintf("%x", sha256.Sum256([]byte(uid+salt)))
	//使用连接操作数据
	if _, err := conn.Do("set", tempid, uid, "EX", "604800"); err != nil {
		return ""
	}
	return tempid
}

// 获得用户app检测的detectid
func GetDetectId(uid string) (detectid int) {
	conn := pool.Get()
	defer conn.Close()
	detectid, _ = redis.Int(conn.Do("Get", uid))
	return
}

// 用户开启了新的app检测
func SetDetectId(uid string, detectid int) {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("set", uid, detectid)
}

// app检测开始
func DetectStart(detectid int) {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("set", detectid, true)
}

// app检测结束
func DetectEnd(detectid int) {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("DEL", detectid)
}

// 获得该次app检测状态：进行中/结束
func IsDetecting(detectid int) bool {
	conn := pool.Get()
	defer conn.Close()
	is_detecting, _ := redis.Bool(conn.Do("EXISTS", detectid))
	return is_detecting
}

// 假如正常获得验证码，说明发送了验证码而且没有过期
// 假如没有获得验证码，说明根本没发送验证码或者验证码过期了
// func GetVariCode(email string) string {
// 	conn := pool.Get()
// 	defer conn.Close()
// 	varicode, _ := redis.String(conn.Do("Get", email))

// 	return varicode
// }

// func DelVariCode(email string) error {
// 	conn := pool.Get()
// 	defer conn.Close()
// 	_, err := conn.Do("DEL", email)
// 	return err
// }
