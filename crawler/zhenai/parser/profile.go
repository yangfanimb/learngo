package parser

import (
	"bytes"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const info0CSS = `#app > div:nth-child(2) > div.CONTAINER.f-cl.f-topIndex.primary > div.CONTAINER.f-fl > div.m-userInfo > div.top.f-cl > div.right.f-fl > div.info`
const info1CSS = `#app > div:nth-child(2) > div.CONTAINER.f-cl.f-topIndex.primary > div.CONTAINER.f-fl > div.CONTAINER > div > div:nth-child(4) > div.purple-btns`
const info2CSS = `#app > div:nth-child(2) > div.CONTAINER.f-cl.f-topIndex.primary > div.CONTAINER.f-fl > div.CONTAINER > div > div:nth-child(4) > div.pink-btns`

var ageRe = regexp.MustCompile(`([\d]+)岁`)
var heightRe = regexp.MustCompile(`([\d]+)cm`)
var weightRe = regexp.MustCompile(`([\d]+)kg`)
var incomeRe = regexp.MustCompile(`月收入:(.+)`)
var workAreaRe = regexp.MustCompile(`工作地:(.+)`)
var hukouRe = regexp.MustCompile(`籍贯:(.+)`)

func extraString(s string, re *regexp.Regexp) string {
	match := re.FindStringSubmatch(s)
	if len(match) < 2 {
		return ""
	}

	return match[1]
}

func extraInt(s string, re *regexp.Regexp) int {
	match := re.FindStringSubmatch(s)
	if len(match) < 2 {
		return -1
	}

	num, err := strconv.Atoi(match[1])
	if err != nil {
		return -1
	}

	return num
}

func ParseProfile(contents []byte, item engine.Item) engine.ParseResult {
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return engine.ParseResult{}
	}

	profile := item.Payload.(model.Profile)
	info0 := dom.Find(info0CSS)
	profile.Name = strings.TrimSpace(info0.Find(".name").Text())
	item.ID = strings.Split(info0.Find(".id").Text(), "：")[1]

	info1 := dom.Find(info1CSS).Children().Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	info2 := dom.Find(info2CSS).Children().Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})

	var info []string
	info = append(info, info1...)
	info = append(info, info2...)

	for _, s := range info {
		if strings.HasSuffix(s, "岁") {
			profile.Age = extraInt(s, ageRe)
		}
		if strings.HasSuffix(s, "cm") {
			profile.Height = extraInt(s, heightRe)
		}
		if strings.HasSuffix(s, "kg") {
			profile.Weight = extraInt(s, weightRe)
		}

		if strings.HasPrefix(s, "月收入:") {
			profile.Income = extraString(s, incomeRe)
		}

		if s == "离异" || s == "未婚" || s == "丧偶" {
			profile.Marriage = s
		}

		if strings.Contains(s, "博士") ||
			strings.Contains(s, "硕士") ||
			strings.Contains(s, "大学") ||
			strings.Contains(s, "高中") ||
			strings.Contains(s, "大专") ||
			strings.Contains(s, "中专") {
			profile.Education = s
		}

		if strings.HasPrefix(s, "工作地:") {
			profile.WorkArea = extraString(s, workAreaRe)
		}

		if strings.HasPrefix(s, "籍贯:") {
			profile.Hukou = extraString(s, hukouRe)
		}

		if strings.Contains(s, "座(") {
			profile.Xinzuo = s
		}

		if strings.HasSuffix(s, "租房") ||
			strings.HasSuffix(s, "购房") {
			profile.House = s
		}

		if strings.HasSuffix(s, "买车") {
			profile.Car = s
		}
	}

	item.Payload = profile
	return engine.ParseResult{
		Items: []engine.Item{item},
	}
}
