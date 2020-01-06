package utils

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"
)

func PostFile(filepath string, filename string, target_url string) (*http.Response, error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	_, err := body_writer.CreateFormFile("Icon", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
	}
	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()
	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
		fmt.Printf("Error Stating file: %s", filepath)
		return nil, err
	}
	req, err := http.NewRequest("POST", target_url, request_reader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	return http.DefaultClient.Do(req)
}

//http get 请求，返回结果
func HttpGetURL(url string) (body string, code int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Msgf("[GET]%s error:%s", url, err)
		return "", 0, err
	}

	defer resp.Body.Close()

	if (resp.StatusCode == 200) {
		log.Debug().Msgf("%s,%d", url, resp.StatusCode)
	} else {
		log.Warn().Msgf("%s,%d", url, resp.StatusCode)
	}
	resp_body, _ := ioutil.ReadAll(resp.Body)
	return string(resp_body), resp.StatusCode, err
}

//下载文件
//todo 可以增加下载进度
func HttpGetFile(url, local string) (err error) {
	return HttpGetFileWithHeader(url, nil, local)
}

// 下载Url到指定目录
func HttpGetFileToDir(url, dir string, headers map[string]string) (error, string) {
	if !IsDirExists(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err, ""
		}
	}
	localPath := dir + "/" + path.Base(url)
	err := HttpGetFileWithHeader(url, headers, localPath)
	return err, localPath
}

// 下载文件，可以追加头
func HttpGetFileWithHeader(url string, headeers map[string]string, local string) error {
	log.Debug().Msgf("%s download start", url)
	// todo 应检查大小
	if IsFileExists(local) {
		log.Debug().Msgf("[%s]貌似本地文件已存在, 无需下载", url)
		return nil
	}
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	for key, value := range headeers {
		request.Header.Add(key, value)
	}

	//处理返回结果
	res, err := client.Do(request)

	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request %s error: %s", url, res.Status)
	}
	defer res.Body.Close()
	file, err := os.Create(local)
	if err != nil {
		return err
	}
	defer file.Close()

	var tick = time.Tick(1000 * time.Millisecond)
	var nb int64 = 0
	downloadSize := make(chan int64)
	remoteSize := res.ContentLength

	go func() {
		var buf = make([]byte, 1024*1024*1024)
		for {
			n, _ := res.Body.Read(buf)
			if n == 0 {
				close(downloadSize)
				break
			}
			file.Write(buf[:n])
			downloadSize <- int64(n)
		}
	}()

loop:
	for {
		select {
		case z, ok := <-downloadSize:
			if !ok {
				// 表示下载完成
				break loop
			}
			nb += z
			if (nb == remoteSize) {
				fmt.Printf(" 已下载:%v/%v, 进度: %.2f%% \r", nb, remoteSize,
					float64(nb)/float64(remoteSize)*100)
			}
		case <-tick:
			fmt.Printf(" 已下载:%v/%v, 进度: %.2f%% \r", nb, remoteSize,
				float64(nb)/float64(remoteSize)*100)
		}
	}

	//io.Copy(f, res.Body)
	log.Debug().Msgf("[%s]已下载完成，保存至[%s]", url, local)
	return nil
}
