package main

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// raw url
const rawURL = "http://jmcomic.xyz"

// match key word
const matchKey = "18comic"

// TODO:
// 1. use mem cache
// 2. use http server

// WriteCache WRITE mem cache
func WriteCache() error {
	return nil
}

// CleanCache CLEAN mem cache
func CleanCache() error {
	return nil
}

// ReadCache READ mem cache
func ReadCache() error {
	return nil
}

// MirrorItem single mirror item
type MirrorItem struct {
	Title   string // comic mirror title
	API     string // comic mirror raw api
	InChina bool   // 中国镜像站
}

// GetMirror get 18comic mirror
func GetMirror() ([]MirrorItem, error) {
	dom, err := goquery.NewDocument(rawURL)
	if err != nil {
		return []MirrorItem{}, errors.New("get raw url data is error")
	}
	var strongs = dom.Find("strong")
	var nodes = strongs.Nodes
	var matchNodes = []*html.Node{}
	for _, item := range nodes {
		var ele = goquery.NewDocumentFromNode(item)
		var text = ele.Text()
		if strings.Contains(text, matchKey) {
			matchNodes = append(matchNodes, item)
		}
	}
	var result = []MirrorItem{}
	for _, item := range matchNodes {
		var temp = goquery.NewDocumentFromNode(item)
		var title string
		var data = temp.Text()
		var checkNext = strings.Contains(data, matchKey)
		var eleParent = temp.Parent()
		var ele = eleParent.Prev()
		title = ele.Text()
		if title == "" {
			title = ele.Parent().Text()
		}
		title = strings.ReplaceAll(title, data, "")
		var findIndex = strings.Index(title, "18comic")
		if findIndex >= 0 {
			title = title[0:findIndex]
		}
		if checkNext {
			result = append(result, MirrorItem{
				Title: title,
				API:   data,
			})
		}
	}
	return result, nil
}
