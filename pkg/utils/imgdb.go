package utils

import (
	"AiServer/configs"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

var ImgDbClient *memcache.Client

func InitImgDb(c configs.ImageDb) {
	ImgDbClient = memcache.New(fmt.Sprintf("%s:%d", c.Host, c.Port))
}
