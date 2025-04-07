package easypage

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func Paginate(c *gin.Context, table string, db *gorm.DB, conditions string, join string, globalSearch ...string) func(db *gorm.DB) *gorm.DB {

	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	//Check if optional parameter globalSearchString is passed
	globalSearchString := ""
	if len(globalSearch) > 0 {
		globalSearchString = globalSearch[0]
	}

	if pageSize != 0 {
		totalRows := countElement(db, table, "", conditions, join, globalSearchString)
		c.Header("X-Total-Count", strconv.FormatInt(totalRows, 10))

		page, _ := strconv.Atoi(c.Query("page"))
		offset := page * pageSize

		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(offset).Limit(pageSize)
		}
	}

	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func PaginateCustomQuery(c *gin.Context, db *gorm.DB, customQuery string) string {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	totalRows := countElement(db, "", customQuery, "", "", "")

	c.Header("X-Total-Count", strconv.FormatInt(totalRows, 10))

	page, _ := strconv.Atoi(c.Query("page"))
	if pageSize == 0 {
		return ""
	}
	offset := page * pageSize
	pageString := strconv.Itoa(pageSize)
	offsetString := strconv.Itoa(offset)
	return " LIMIT " + pageString + " OFFSET " + offsetString

}

func countElement(db *gorm.DB, table string, customQuery string, conditions string, join string, globalSearchString string) int64 {
	var totalRows int64
	if customQuery != "" {
		if globalSearchString != "" {
			customQuery = customQuery + " WHERE " + globalSearchString
		}
		countQuery := "SELECT COUNT(*) FROM (" + customQuery + ") AS TB"
		db.Raw(countQuery).Scan(&totalRows)
	} else {
		db.Table(table).
			Joins(join).
			Where(conditions).
			Where(globalSearchString).
			Count(&totalRows)
	}
	return totalRows
}
