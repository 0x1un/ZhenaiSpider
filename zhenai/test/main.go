package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const url = "http://album.zhenai.com/u/1750987033"

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		r, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		profile := regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)

		//profile := regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^<]+)</div>`)
		body := profile.FindAllSubmatch(r, -1)
		print(len(body))
		fmt.Printf("%s\n", body[2][1])
		//profiles := model.Profile{}
		//profiles.Car = string(body[6][1])
		//fmt.Println(profiles.Car)
		//for _, matches := range body {
		//	fmt.Printf("%s\n", matches)
		//}

	}
}


