package main

import (
	"bufio"
	"fmt"
	"github.com/ajph/nbclassifier-go"
	"github.com/yanyiwu/gojieba"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"syscall"
)

const SPAM_CHECK_SOCKET_FILE = "/tmp/spamcheck.sock"

// 使用go 实现简单的贝叶斯分类
func getWords(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var ret []string
	for {
		line, err := reader.ReadString('n')
		if err != nil || io.EOF == err {
			if line == "" {
				break
			}
		}
		line = strings.Trim(line, "n")
		fmt.Println("处理单词：" + line)
		ret = append(ret, line)
	}
	return ret
}

func learn() {
	m := nbclassifier.New()

	m.NewClass("normal")
	normalwords := getWords("normalwords.txt")
	//fmt.Println(normalwords)
	m.Learn("normal", normalwords...)
	//m.Learn("normal", "a", "need")

	m.NewClass("forbidden")
	forbiddenwords := getWords("forbiddenwords.txt")
	//fmt.Println(forbiddenwords)
	m.Learn("forbidden", forbiddenwords...)
	//m.Learn("forbidden", " design ", "banner", " picture", " logo ", "clip art", " ad ", "clipart", "hairstyles", " drawing", " rendering", " diagram ", " poster", "изображение")

	m.NewClass("terror")
	terrorwords := getWords("terrorwords.txt")
	//fmt.Println(terrorwords)
	m.Learn("terror", terrorwords...)
	//m.Learn("terror", "...", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "…", "image", "pinterest", ".c", "ltd.", "vector", "quote", "video", "search", "?", "click", "psd", "ai", "print", "file", "related", "download", "submit", "view", "buy", "how", "maker", "online", " on", "by")

	m.SaveToFile("materiel.json")

}

func reloadModel() *nbclassifier.Model {
	model, _ := nbclassifier.LoadFromFile("materiel.json")
	//fmt.Println(model.Classes[0].Items[0])
	//fmt.Println(model.Classes[1])
	//fmt.Println(model.Classes[2])
	return model
}

func match(model *nbclassifier.Model, content string) string {
	// 分词
	jieba := gojieba.NewJieba()
	defer jieba.Free()
	words := jieba.Cut(content, true)
	cls, unsure, _ := model.Classify(words...)
	fmt.Println("检测到分类为：" + cls.Id)

	result := "normal"
	if unsure == false {
		result = cls.Id
		fmt.Println(cls, unsure)
	}
	return result
}

func run() {
	socket, _ := net.Listen("unix", SPAM_CHECK_SOCKET_FILE)
	defer syscall.Unlink(SPAM_CHECK_SOCKET_FILE)
	learn()
	// 训练物料
	model := reloadModel()

	for {
		client, _ := socket.Accept()

		buf := make([]byte, 1024)
		datalength, _ := client.Read(buf)
		data := buf[:datalength]
		fmt.Println("client msg:" + string(data))

		checkret := match(model, string(data))
		fmt.Println("check result: " + checkret)
		response := []byte("")
		if len(checkret) > 0 {
			response = []byte(checkret)
		}
		_, _ = client.Write(response)
	}
}

func main() {
	// 开启sock，检测服务
	run()
	//fmt.Println(reloadModel())
}
