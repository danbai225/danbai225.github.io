package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	re := regexp.MustCompile(`!\[.*]\(https://danbai.*\)`)
	re2 := regexp.MustCompile(`url: https://p00q.cn/\?p=(.*)\n`)

	_ = filepath.Walk("./md", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		s, _ := os.ReadFile(path)
		urls := re.FindAllString(string(s), -1)
		arr := make([]string, 0)
		for _, url := range urls {
			str := strings.ReplaceAll(strings.Split(url, "](")[1], ")", "")
			arr = append(arr, str)
		}
		if len(arr) > 0 {
			id := strings.ReplaceAll(re2.FindString(string(s)), "url: https://p00q.cn/?p=", "")
			for i, s2 := range arr {
				resp, err := http.Get(s2)
				if err != nil {
					continue
				}
				all, _ := io.ReadAll(resp.Body)
				_ = os.WriteFile(fmt.Sprintf("res/img2/%s-%d.png", id, i+1), all, 0644)
			}
		}

		return nil
	})

}
