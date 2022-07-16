package random

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strings"
	"sync/atomic"
)

const imageNum = 350

var index int32 = 0
var images [imageNum][]byte

func init() {
	fmt.Println("Starting image.go init")
	var files []fs.FileInfo

	var err error
	// 画像ファイル群の読み込み
	files, err = ioutil.ReadDir(imageFolderPath)
	if err != nil {
		log.Fatalf("%+v", fmt.Errorf("%w", err))
	}

	for i := 0; i < imageNum; i++ {
		fileInfo := files[rand.Intn(len(files))]
		//default.jpg以外の、.jpgで終わるファイルに限定する
		for fileInfo.Name() == "default.jpg" || !strings.HasSuffix(fileInfo.Name(), ".jpg") {
			fileInfo = files[rand.Intn(len(files))]
		}
		imagePath := filepath.Join(imageFolderPath, fileInfo.Name())
		byteArray, err := ioutil.ReadFile(imagePath)
		if err != nil {
			panic("failed to read the image")
		}
		images[i] = byteArray
	}
}

func Image() ([]byte, error) {
	// MEMO: 現状 error は返してないがメモリがやばければファイル読み込みに変える
	return images[atomic.AddInt32(&index, 1)%imageNum], nil
}
