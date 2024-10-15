package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// PageContent 用于存储网页内容
type PageContent struct {
	Title    string   `json:"title"`
	Headings []string `json:"headings"`
	Links    []string `json:"links"`
	Videos   []string `json:"videos"`
}

func CrawlerTester() {
	// 替换为目标URL
	url := "exampleurl"

	// 发起请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("请求失败，状态码:", resp.StatusCode)
		return
	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("解析HTML失败:", err)
		return
	}

	// 创建PageContent结构体实例
	content := PageContent{}

	// 获取标题
	content.Title = doc.Find("title").Text()

	// 获取所有的标题（h1, h2, h3等）
	doc.Find("h1, h2, h3").Each(func(i int, s *goquery.Selection) {
		content.Headings = append(content.Headings, s.Text())
	})

	// 获取所有链接
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		content.Links = append(content.Links, link)
	})

	// 获取视频源
	doc.Find("video").Each(func(i int, s *goquery.Selection) {
		videoSrc, exists := s.Attr("src")
		if exists {
			content.Videos = append(content.Videos, videoSrc)
		} else {
			// 检查<source>标签
			s.Find("source").Each(func(i int, source *goquery.Selection) {
				sourceSrc, exists := source.Attr("src")
				if exists {
					content.Videos = append(content.Videos, sourceSrc)
				}
			})
		}
	})

	// 查找所有<iframe>标签
	doc.Find("iframe").Each(func(i int, s *goquery.Selection) {
		iframeSrc, exists := s.Attr("src")
		if exists {
			content.Videos = append(content.Videos, iframeSrc)
		}
	})

	// 将内容转换为JSON格式
	jsonData, err := json.MarshalIndent(content, "", "    ")
	if err != nil {
		fmt.Println("转换为JSON失败:", err)
		return
	}

	// 输出JSON数据
	fmt.Println(string(jsonData))
}
