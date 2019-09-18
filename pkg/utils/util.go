package utils

import (
	"encoding/json"
	"fmt"
	"github.com/wpp/PDDComments/types"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	EmptyComment = "此用户未填写文字评论"
	FullScore = 5
	commentUrl = "https://mobile.yangkeduo.com/proxy/api/reviews/%s/list?page=%d&size=10&enable_video=1&enable_group_review=1"
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

func GenerateComment(minLength string, comments []types.Comment) string {
	result := []byte{}
	lastComment := ""
	length, _ := strconv.Atoi(minLength)
	size := len(comments)
	for _, c := range comments {
		fmt.Println(c)
		i := rand.Intn(size -1) + 1
		if comments[i].Comment == lastComment {
			continue
		}
		result = append(result, comments[i].Comment...)
		lastComment = comments[i].Comment
		fmt.Println(utf8.RuneCountInString(string(result)))
		if utf8.RuneCountInString(string(result)) > length {
			return string(result)
		}
		result = append(result, "，"...)
	}
	return string(result)
}
