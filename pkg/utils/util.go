package utils

import (
	"encoding/json"
	"fmt"
	"github.com/wpp/PDDComments/types"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	EmptyComment = "此用户未填写文字评论"
	FullScore = 5
	commentUrl = "https://mobile.yangkeduo.com/proxy/api/reviews/%s/list?page=%d&size=10&enable_video=1&enable_group_review=1"
	configFile = "./PDDComments.json"
)

func RemoveEmptyComments(comments []types.Comment) []types.Comment {
	result := make([]types.Comment, 0)
	for _, c := range comments {
		if c.Comment != EmptyComment {
			result = append(result, c)
		}
	}
	return result
}

func RemoveLowScoreComments(comments []types.Comment) []types.Comment {
	result := make([]types.Comment, 0)
	for _, c := range comments {
		if c.DescScore == FullScore &&
			c.LogisticsScore == FullScore && c.ServiceScore == FullScore {
			result = append(result, c)
		}
	}
	return result
}

func GetOnePageComments(key string, itemId string, page int) []types.Comment {
	ak := key
	client := &http.Client{}
	var req *http.Request
	itemUrl := fmt.Sprintf(commentUrl, itemId, page)
	req, _ = http.NewRequest(http.MethodGet, itemUrl, nil)
	req.Header.Add("accesstoken", ak)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		resp.Body.Close()
	} ()
	data, _ := ioutil.ReadAll(resp.Body)
	tmp := &types.CommentData{}
	if err := json.Unmarshal(data, tmp); err != nil {
		fmt.Println(err)
		return nil
	}
	return tmp.Data
}

func FilterKeys(filter string, comments []types.Comment) []types.Comment {
	filters :=strings.Split(filter, ",")
	if len(filter) == 0 || len(filters) == 0{
		return comments
	}
	result := make([]types.Comment, 0)
	for _, c := range comments {
		add := true
		for _, f := range filters {
			if strings.Index(c.Comment, f) > 0 {
				add = false
			}
		}
		if add {
			result = append(result, c)
		}
	}
	return result
}

func InitConfig(filename string) {
	if len(filename) == 0 {
		filename = configFile
	}
	_, err := os.Stat(filename)
	fmt.Println(err)
	if os.IsNotExist(err) {
		os.Create(filename)
	}
}

func SaveConfig(filename string, json string) {
	if len(filename) == 0 {
		filename = configFile
	}
	InitConfig(filename)
	if err := ioutil.WriteFile(filename, []byte(json), 0755); err != nil {
		fmt.Println(err)
	}
}
