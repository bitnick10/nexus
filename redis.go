package nexus

import (
	"github.com/garyburd/redigo/redis"
	// "bytes"
	// "database/sql"
	// "encoding/json"
	"fmt"
	// _ "github.com/lib/pq"
	"net/http"
	// "os"
	"io/ioutil"
	"strings"
	// "sync"
	// "strings"
	// "strconv"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

var pool = newPool()

func CloseRedisPool() {
	pool.Close()
}

func Redis(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Query().Get("method")
	key := r.URL.Query().Get("key")
	c := pool.Get()
	defer c.Close()

	if strings.ToUpper(method) == "SET" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		c.Do("SET", key, string(body))
	}
	// c.s
	// fmt.Println("abababababababa")
	// c, _ := redis.Dial("tcp", ":6379")
	// defer c.Close()
	// if err != nil {
	// 	fmt.Println("redis.Dial")
	// 	fmt.Println(err)
	// 	fmt.Println(redisConn.Err())
	// }
	data, _ := redis.String(c.Do(method, key))
	// fmt.Println("dadadadadadadadad")
	fmt.Fprintf(w, "%s", data)
}
