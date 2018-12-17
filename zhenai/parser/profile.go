package parser

import (
	"learning/crawler/engine"
	"learning/crawler/model"
	"regexp"
)

var profileList = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)
var profileList2 = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	matches := profileList.FindAllSubmatch(contents, -1)
	if matches != nil {
		profile.Name = name
		profile.Age = string(matches[1][1])

		profile.Marriage = string(matches[0][1])
		profile.Xingzuo = string(matches[2][1])
		profile.Height = string(matches[3][1])
		profile.Weight = string(matches[4][1])
		if len(matches) > 5 {
			profile.WorkPlace = string(matches[5][1])
		}
		if len(matches) > 6 {
			profile.Income = string(matches[6][1])
		}
		if len(matches) > 7 {
			profile.Occupation = string(matches[7][1])
		}
		if len(matches) > 8 {
			profile.Education = string(matches[8][1])
		}
	}
	matchess := profileList2.FindAllSubmatch(contents, -1)
	if matchess != nil {
		if len(matchess) > 6 {
			profile.Car = string(matchess[6][1])
		}

		profile.Race = string(matchess[0][1])
		if len(matchess) > 5 {
			profile.House = string(matchess[5][1])
		}
	}

	result := engine.ParseResult{Items: []interface{}{profile}}
	return result

}
