package models

import (
	"github.com/astaxie/goredis"
)

const (
	REDIS_HOST = "127.0.0.1:6379"
)

var (
	client goredis.Client
)

func init()  {
	connectRedis(REDIS_HOST)
}

func connectRedis(addr string) {
	client.Addr = addr
	// client.Db = 1
	// client.Password = ""
}

// 将要爬取的 URL 放入队列
func PutinQueue(url, key string){
	client.Lpush(key, []byte(url))
}

// 从队列中取出一个 URL
func PopfromQueue(key string) string {
	res, err := client.Rpop(key)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// 获取队列长度
func GetQueueLength(key string) int {
	length, err := client.Llen(key)
	if err != nil {
		return 0
	}

	return length
}

// 将已访问过的 URL 放入集合当中
func AddToSet(url, key string) {
	client.Sadd(key, []byte(url))
}

func RmFromSet(url, key string) {
	client.Srem(key, []byte(url))
}

func GetSetLenth(key string) int {
	length, err := client.Scard(key)
	if err != nil {
		return 0
	}

	return length
}

func PopFromSet(key string) string{
	item, err := client.Spop(key)
	if nil != err {
		return ""
	}
	return string(item)
}

// 判断 URL 是否已访问过
func IsVisit(url, key string) bool {
	bIsVisit, err := client.Sismember(key, []byte(url))
	if err != nil {
		return false
	}

	return bIsVisit
}
