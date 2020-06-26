/*
 * Maintained by jemo from 2020.6.24 to now
 * Created by jemo on 2020.6.24 16:44:09
 * Counting item by created time
 */

package handle

import (
  "encoding/json"
  "log"
  "math"
  "time"
  "net/http"
  "github.com/jiangtaozy/pdd-management-api/database"
)

func CountingItemByCreatedTime(w http.ResponseWriter, r *http.Request) {
  db := database.DB
  rows, err := db.Query("SELECT createdAt FROM pddItem ORDER BY createdAt ASC")
  if err != nil {
    log.Println("counting-item-by-created-time-query-error: ", err)
  }
  defer rows.Close()
  var timeList []interface{}
  for rows.Next() {
    var createdAt string
    if err := rows.Scan(&createdAt); err != nil {
      log.Println("counting-item-by-created-time-scan-error: ", err)
    }
    timeList = append(timeList, createdAt)
  }
  length := len(timeList)
  var countList []interface{}
  if length > 0 {
    start := timeList[0].(string)
    startTime, _ := time.Parse("2006-01-02 15:04:05", start)
    location, _ := time.LoadLocation("Local")
    startDate := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, location)
    now := time.Now()
    duration := now.Sub(startDate)
    days := int(math.Floor(duration.Hours() / 24))
    for i := 0; i <= days; i++ {
      date := startDate.AddDate(0, 0, i)
      count := 0
      for j := 0; j < len(timeList); j++ {
        timeStr := timeList[j].(string)
        time, _ := time.Parse("2006-01-02 15:04:05", timeStr)
        subTime := time.Sub(date)
        if subTime.Hours() > 0 && subTime.Hours() < 24 {
          count++
        }
      }
      timeCount := map[string]interface{}{
        "date": date,
        "count": count,
      }
      countList = append(countList, timeCount)
    }
  }
  json.NewEncoder(w).Encode(countList)
}
