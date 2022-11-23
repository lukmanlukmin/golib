package query

import (
	"fmt"
	"time"
)

func FilterAndQuery(filter map[string]interface{}) string {
	query := ""
	for key, element := range filter {
		switch val := element.(type) {
		case int:
			query = fmt.Sprintf("%s %s=%v AND", query, key, element)
		case uint:
			query = fmt.Sprintf("%s %s=%v AND", query, key, element)
		case int32:
			query = fmt.Sprintf("%s %s=%v AND", query, key, element)
		case uint32:
			query = fmt.Sprintf("%s %s=%v AND", query, key, element)
		case int64:
			query = fmt.Sprintf("%s %s=%v AND", query, key, element)
		case uint64:
			query = fmt.Sprintf("%s %s=%v AND", query, key, element)
		case float64:
			query = fmt.Sprintf("%s %s=%v AND", query, key, element)
		case string:
			query = fmt.Sprintf("%s %s='%v' AND", query, key, element)
		case bool:
			query = fmt.Sprintf("%s %s=%v AND", query, key, element)
		default:
			fmt.Printf("cant identify %v", val)
		}
	}
	lenFilter := len(query)
	if lenFilter > 1 {
		query = query[:lenFilter-4]
	}
	return query
}

func FilterORQuery(filter map[string]interface{}) string {
	query := ""
	for key, element := range filter {
		switch val := element.(type) {
		case int:
			query = fmt.Sprintf("%s %s=%v OR", query, key, element)
		case int64:
			query = fmt.Sprintf("%s %s=%v OR", query, key, element)
		case float64:
			query = fmt.Sprintf("%s %s=%v OR", query, key, element)
		case string:
			query = fmt.Sprintf("%s %s='%v' OR", query, key, element)
		case bool:
			query = fmt.Sprintf("%s %s=%v OR", query, key, element)
		default:
			fmt.Printf("cant identify %v", val)
		}
	}
	lenFilter := len(query)
	if lenFilter > 1 {
		query = query[:lenFilter-3]
	}
	return query
}

func TranslatePagination(page int, perPage int) (int, int) {
	limit := perPage
	offset := (page * perPage) - perPage
	if limit < 1 {
		limit = 1
	}
	if page <= 1 {
		offset = 0
	}
	return limit, offset
}

func BuildQueryIN(list []int64) string {
	query := ""
	for _, element := range list {
		query = fmt.Sprintf("%s%v,", query, element)
	}
	lenFilter := len(query)
	if lenFilter > 1 {
		query = query[:lenFilter-1]
	}
	return query
}

func BuildQueryINString(list []string) string {
	query := ""
	for _, element := range list {
		query = fmt.Sprintf("%s'%v',", query, element)
	}
	lenFilter := len(query)
	if lenFilter > 1 {
		query = query[:lenFilter-1]
	}
	return query
}

func NormalizePerPage(perPage *int) {
	if perPage == nil {
		*perPage = 25
	}
	if *perPage > 25 || *perPage <= 0 {
		*perPage = 25
	}
}

func NormalizePage(page *int) {
	if page == nil {
		*page = 1
	}
	if *page <= 0 {
		*page = 1
	}
}

const DateSimpleFormat = "2006-01-02"

func ValidateTimeRangeString(startDate string, endDate string) (time.Time, time.Time) {
	start, err := time.Parse(DateSimpleFormat, startDate)
	if err != nil {
		return genereteDefaultTimeRange()
	}
	end, err := time.Parse(DateSimpleFormat, endDate)
	if err != nil {
		return genereteDefaultTimeRange()
	}
	resStart, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", start.Format(DateSimpleFormat)))
	resEnd, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 23:59:59", end.Format(DateSimpleFormat)))
	return resStart, resEnd
}

func genereteDefaultTimeRange() (time.Time, time.Time) {
	yesterday := time.Time{}
	todayString := time.Now().Format(DateSimpleFormat)
	end, _ := time.Parse(DateSimpleFormat, todayString)
	today, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 23:59:59", end.Format(DateSimpleFormat)))
	return yesterday, today
}
