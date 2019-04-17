package models

import (
	"github.com/astaxie/goredis"
)

const (
	REDIS_HOST = "127.0.0.1:6379"
	KEY_URL_QUEUE = "movie_url_queue"
	KEY_URL_VISIT_SET = "movie_url_visited_set"
)

var (
	client goredis.Client
)

// 连接至 Redis
func ConnectRedis(addr string) {
	client.Addr = addr
	// client.Db = 1
	// client.Password = ""
}

// 将要爬取的 URL 放入队列
func PutinQueue(url string){
	client.Lpush(KEY_URL_QUEUE, []byte(url))
}

// 从队列中取出一个 URL
func PopfromQueue() string {
	res, err := client.Rpop(KEY_URL_QUEUE)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// 获取队列长度
func GetQueueLength() int {
	length, err := client.Llen(KEY_URL_QUEUE)
	if err != nil {
		return 0
	}

	return length
}

// 将已访问过的 URL 放入集合当中
func AddToSet(url string){
	client.Sadd(KEY_URL_VISIT_SET, []byte(url))
}

// 判断 URL 是否已访问过
func IsVisit(url string) bool {
	bIsVisit, err := client.Sismember(KEY_URL_VISIT_SET, []byte(url))
	if err != nil {
		return false
	}

	return bIsVisit
}
