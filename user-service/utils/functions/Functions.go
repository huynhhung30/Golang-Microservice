package functions

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func StringToJson(jsonStr string) interface{} {
	var res interface{}
	json.Unmarshal([]byte(jsonStr), &res)
	return res
}

func StructToJsonStr(modelA interface{}) string {
	json, err := json.Marshal(modelA)
	if err != nil {
		return "{}"
	}
	return string(json)
}

func ShowLog(tag string, msg ...interface{}) {
	Block{
		Try: func() {
			body, _ := json.Marshal(msg)
			fmt.Println("[ProductService] " + tag + " -> " + string(body[1:len(body)-1]))
		},
		Catch: func(e Exception) {
			fmt.Println("[ProductService] " + tag + " -> " + fmt.Sprintln(e))
		},
	}.Do()
}

func MergeTwoStructToJson(a interface{}, b interface{}) interface{} {
	var m map[string]string
	ja, _ := json.Marshal(a)
	json.Unmarshal(ja, &m)
	jb, _ := json.Marshal(b)
	json.Unmarshal(jb, &m)

	body, _ := json.Marshal(m)
	return string(body[1 : len(body)-1])
}

func CurrentTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now().In(loc)
	now = now.Add(time.Hour * time.Duration(7))
	return now
}

func CurrentTimeWithTimeZone() time.Time {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now().In(loc)
	return now
}

func CurrentTimeWithoutTimeZone() time.Time {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	return now
}

func GetStartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func GetEndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

func ConvertIdListToSqlIdList(idList []int) (sqlIdList string) {
	sqlIdList = "(" + strings.Trim(strings.Replace(fmt.Sprint(idList), " ", ",", -1), "[]") + ")"
	return sqlIdList
}

// Find element in array
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	if reflect.TypeOf(array).Kind() == reflect.Slice {
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func GenerateRandomNumer(length int) int {
	const charsetFirst = "123456789"
	const charsetOther = "0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(CurrentTime().UnixNano()))
	number := 0
	number = seededRand.Intn(len(charsetFirst))
	for i := 1; i < length; i++ {
		number = number*10 + seededRand.Intn(len(charsetOther))
	}
	return number
}

func CalculateTotalPage(totalCount int, limit int) int {
	totalPage := int((totalCount) / limit)
	if totalCount-totalPage*int(limit) > 0 {
		totalPage += 1
	}
	return totalPage
}
