package parser

import (
	"bytes"
	"learngo/crawler/engine"
	"learngo/crawler/helper/log"
	"learngo/crawler/model"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var logger = log.DLogger()

const userListCSS = `#app > div.g-container > div.main.clearfix > div.m-left > div.g-list`

var (
	profileRe   = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	_genderRe   = regexp.MustCompile(`性别：</span>([^<]+)</td>`)
	_workAreaRe = regexp.MustCompile(`居住地：</span>([^<]+)</td>`)
	_ageRe      = regexp.MustCompile(`年龄：</span>([\d]+)</td>`)
	_incomeRe   = regexp.MustCompile(`薪：</span>([^<]+)</td>`)
	_marriageRe = regexp.MustCompile(`婚况：</span>([^<]+)</td>`)
	_heightArea = regexp.MustCompile(`高：</span>([\d]+)</td>`)
	_cityURLRe  = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func extractInfoString(re *regexp.Regexp, info string) string {
	m := re.FindStringSubmatch(info)
	if len(m) != 2 {
		return ""
	}

	return m[1]
}

func extractInfoInt(re *regexp.Regexp, info string) int {
	m := re.FindStringSubmatch(info)
	if len(m) != 2 {
		return 0
	}
	num, _ := strconv.Atoi(m[1])
	return num
}

func ParseCity(contents []byte) engine.ParseResult {
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return engine.ParseResult{}
	}

	result := engine.ParseResult{}
	dom.Find(userListCSS).Children().EachWithBreak(func(i int, s *goquery.Selection) bool {
		className, _ := s.Attr("class")
		if className != "list-item" {
			return false
		}
		info, _ := s.Html()

		profile := model.Profile{}
		profile.Gender = extractInfoString(_genderRe, info)
		profile.WorkArea = extractInfoString(_workAreaRe, info)
		profile.Age = extractInfoInt(_ageRe, info)
		profile.Income = extractInfoString(_incomeRe, info)
		profile.Marriage = extractInfoString(_marriageRe, info)
		profile.Height = extractInfoInt(_heightArea, info)

		item := engine.Item{}
		m := profileRe.FindStringSubmatch(info)
		if len(m) >= 2 {
			profile.Name = m[2]
			item.URL = m[1]
			item.Type = "zhenai"
			item.ID = ""
			item.Payload = profile
		}

		result.Requests = append(result.Requests, engine.Request{
			Url:    item.URL,
			Parser: NewProfileParser(item),
		})

		return true
	})

	// 增加更多城市
	matches := _cityURLRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}

type ProfileParser struct {
	item engine.Item
}

func NewProfileParser(item engine.Item) *ProfileParser {
	return &ProfileParser{item:item}
}

func (p *ProfileParser) Parse(contents []byte) engine.ParseResult {
	return ParseProfile(contents, p.item)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ParseProfile", p.item
}
