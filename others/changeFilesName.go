package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	path := "G:\\BaiduNetdiskDownload\\50本Golang电子书\\50本Golang电子书"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
		re := regexp.MustCompile("^\\d{3}[\u2014]{2}(.*)")
		//re := regexp.MustCompile(`^\d*(.*)`)
		//re := regexp.MustCompile(`[\\u2014]{2}`)
		//re := regexp.MustCompile(`(\d*)`)
		//re := regexp.MustCompile(`\\—{2}`)
		//re := regexp.MustCompile("[\u2014]{2}")
		matches := re.FindStringSubmatch(f.Name())
		//fmt.Println(matches[0])
		fmt.Println(matches[1])
		//for _,s :=range matches{
		//	fmt.Println(s)
		//}
		newName := matches[1]
		//newName := re.Find([]byte(f.Name()))
		//fmt.Println(string(newName[:]))

		_ = os.Rename(path+"\\"+f.Name(), path+"\\"+newName)
	}
}

/**
目前golang只支持捕捉分组，但是拿捕获的分组进行匹配还不支持；此外也不支持向前的零宽断言。
断言有4种：向前的肯定/否定，向后的肯定/否定。
目前绝大多数正则实现都不支持向前的。

(?<!\s)\d 前面不是空bai格的du数zhi字，不包含dao空格 不支持
(?<=\s)\d 前面一zhuan位是空格的数字，不包含空格 不支持
\d(?!\s) 后面一位不是shu空格的数字，不包含空格
\d(?=\s) 后面一位是空格的数字，不包含空格

 */
