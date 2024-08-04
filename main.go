package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
	"gopkg.in/yaml.v3"
	"html/template"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type site struct {
	Domain string
	Name   string
	Author string
}
type meta struct {
	Title      string   `yaml:"title"`
	Date       string   `yaml:"date"`
	Updated    string   `yaml:"updated"`
	URL        string   `yaml:"url"`
	Categories []string `yaml:"categories"`
	Tags       []string `yaml:"tags"`
	Summary    string   `yaml:"summary"`
	Id         string
}
type data struct {
	site
	meta
	Content template.HTML
}
type index struct {
	site
	Posts []meta
}

var siteData = site{
	Domain: "https://p00q.cn",
	Name:   "淡白的记忆",
	Author: "淡白",
}

func main() {
	var metas []meta
	err := filepath.Walk("./md", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		metas = append(metas, parse(path))
		return nil
	})
	if err != nil {
		fmt.Printf("遍历路径时出错 %q: %v\n", "./md", err)
		return
	}
	//index
	sort.Slice(metas, func(i, j int) bool {
		return metas[i].Date > metas[j].Date
	})
	t, err := template.ParseFiles("./index.tpl.html")
	if err != nil {
		fmt.Printf("解析模板报错: %v", err)
		return
	}
	b := bytes.NewBufferString("")
	err = t.Execute(b, index{siteData, metas})
	if err != nil {
		fmt.Printf("执行模板报错: %v", err)
		return
	}
	_ = os.WriteFile("index.html", []byte(b.String()), 0644)
}

func parse(path string) meta {
	d, _ := os.ReadFile(path)
	lines := strings.Split(string(d), "\n")
	status := 0
	metadata := ""
	dataStr := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "---") {
			if status == 0 {
				status = 1
			} else if status == 1 {
				status = 2
				continue
			}
		}
		if status == 1 {
			metadata += line + "\n"
		}
		if status == 2 {
			dataStr += line + "\n"
		}
	}
	m := meta{}
	err := yaml.Unmarshal([]byte(metadata), &m)
	if err != nil {
		fmt.Printf("解析元数据报错: %v", err)
		return m
	}

	u, _ := url.Parse(m.URL)
	id := u.Query().Get("p")
	m.Id = id
	m.Date = m.Date[:19]
	m.Updated = m.Updated[:19]
	htmlPath := fmt.Sprintf("posts/%s.html", id)
	stat, err := os.Stat(htmlPath)
	if err == nil && stat.Size() > 0 {
		return m
	}
	if m.Summary == "" {
		m.Summary = summary(dataStr)
		if m.Summary == "" {
			return m
		}
		out, _ := yaml.Marshal(m)
		_ = os.WriteFile(path, []byte(fmt.Sprintf("---\n%s---\n%s", string(out), dataStr)), 0644)
	}
	fmt.Println(m.Title, m.Summary)
	s := fillTemplate(string(mdToHTML([]byte(dataStr))), m)
	_ = os.WriteFile(htmlPath, []byte(s), 0644)
	return m
}
func mdToHTML(md []byte) []byte {
	goMd := goldmark.New(
		goldmark.WithExtensions(extension.GFM, &mermaid.Extender{}),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := goMd.Convert(md, &buf); err != nil {
		return []byte(err.Error())
	}
	return buf.Bytes()
}
func fillTemplate(content string, meta meta) string {
	t, err := template.ParseFiles("./post.tpl.html")
	if err != nil {
		fmt.Printf("解析模板报错: %v", err)
		return ""
	}
	b := bytes.NewBufferString("")
	err = t.Execute(b, data{siteData, meta, template.HTML(content)})
	if err != nil {
		fmt.Printf("执行模板报错: %v", err)
		return ""
	}
	return b.String()
}

var openaiClient *openai.Client

func summary(content string) string {
	if openaiClient == nil {
		openaiClient = openai.NewClient(os.Getenv("TOKEN"))
	}
	resp, err := openaiClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "提取文章的摘要，用于博客展示摘要。",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
	})
	if err != nil {
		fmt.Printf("openai 摘要错误: %v", err)
		return ""
	}
	time.Sleep(10 * time.Second)
	return resp.Choices[0].Message.Content
}
