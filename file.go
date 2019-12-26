package utils

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
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
	u, _ := url.Parse(src)
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

// 按行读取文件 通过 chanel 返回
func ReadLine(fileName string, result chan<- string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		result <- line
	}
	return nil
}

// 按行读取文件 通过 chanel 返回
func ReadLineByFn(fileName string, cb func(s string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		cb(line)
	}
	return nil
}
// 通过文件路径复制文件
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}