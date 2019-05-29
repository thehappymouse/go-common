package utils

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
)
var (
	//记录已下载的图片，减少IO请次
	mu               sync.RWMutex
	downloadedImages map[string]bool
)

func init() {
	downloadedImages = map[string]bool{}
}

//检查文件是否存在
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//查检目录是否存在
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("not reached")
}


// 下载文件到指定的目录
func DownLoadImgToDir(src string, localDir string) (filename string, err error) {
	mu.RLock()
	if downloadedImages[src] {
		mu.RUnlock()
		return
	}
	mu.RUnlock()
	mu.Lock()
	downloadedImages[src] = true
	mu.Unlock()
	u, _:=url.Parse(src)
	filename = path.Base(u.Path)
	localPath := fmt.Sprintf("%s/%s", localDir, filename)

	if IsFileExists(localPath) {
		return
	}
	log.Debug().Msgf("[download][%s]==>[%s]", src, localPath)
	resp, err := http.Get(src)

	if err != nil {
		log.Warn().Msgf("请求【%s】内容失败，请参考:%s", src, err)
		return
	}
	defer resp.Body.Close()

	//小于10K的图片，不下载
	//if resp.ContentLength < 1024*20 {
	//	log.Println("可能是图片太小，忽略", resp.ContentLength)
	//	return
	//}

	if !IsDirExists(localDir) {
		os.MkdirAll(localDir, 0755)
	}
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn().Msgf("接收【%s】内容失败，请参考:%s", src, err)
		return
	}

	err = ioutil.WriteFile(localPath, pix, 0755)
	if err != nil {
		log.Fatal().Msgf("写入磁盘【%s】，请参考:%s", localPath, err)
	}
	return filename, nil
}