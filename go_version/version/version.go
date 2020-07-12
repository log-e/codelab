package version

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type Version struct {
	Platform string `json:"platform"` // 构建平台
	GoVersion string `json:"goVersion"` // go 编译器版本
	BuildDate string `json:"buildDate"` // 构建日期
	HashSelf string `json:"hash"` // 程序摘要
}

func (version Version) String() string {
	marshaled, err := json.MarshalIndent(&version, "", "    ")
	if err != nil {
		return err.Error()
	}
	return string(marshaled)
}

func selfMd5sum() string {
	path, e := exec.LookPath(os.Args[0])
	if e != nil {
		log.Println(`Self md5sum error: look path fail:`, e)
		return ""
	}
	bs, e := ioutil.ReadFile(path)
	if e != nil {
		log.Println(`SelfMd5sum error: read file fail:`, e)
		return ""
	}
	sum := md5.Sum(bs)
	return hex.EncodeToString(sum[:])
}

func Get() Version {
	return Version{
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		GoVersion: runtime.Version(),
		BuildDate: buildDate,
		HashSelf:  selfMd5sum(),
	}
}

