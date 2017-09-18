package redistest

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"log"
)

var RedisPort = "6379"
var RedisIP = "140.115.153.185"


func Redis_Get(KEY_NAME string) (string, int){
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
    CheckError(err)
    defer c.Close()
	v, err := redis.Strings(c.Do("GET", KEY_NAME))
	CheckError(err)
	// fmt.Println(v)
	return v, 710
}

func Redis_Set(KEY_NAME string, VAL string) int{
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
    CheckError(err)
    defer c.Close()
	_, err2 := c.Do("SET", KEY_NAME, VAL)
	CheckError(err2)
	return 720
}

func CheckError(err error) {
    if err  != nil {
        log.Println("Error: " , err)
        // os.Exit(0)
    }
}
