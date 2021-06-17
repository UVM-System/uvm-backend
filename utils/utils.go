package utils

import (
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

func GetUUID() string {
	return uuid.NewV4().String()
}

/**
读取指定文件夹下按名称排序的文件名
*/
func ReadDir(dirPath string) (list []os.FileInfo, err error) {
	list, err = ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name() < list[j].Name()
	})
	//for _, item := range list {
	//	log.Println(item.Name())
	//}
	return
}

/**
输入文件名带后缀，获取文件名和后缀
*/
func GetFileNameAndSuffix(fullName string) (fileName, suffix string) {
	suffix = path.Ext(fullName)
	fileName = strings.TrimSuffix(fullName, suffix)
	return
}
