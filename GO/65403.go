package main

import	"fmt"
import "time"


type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	reportID := int(time.Now().UnixNano()%1e9) + user.ID
	return Report{
		User:     user,
		ReportID: reportID,
		Date:     reportDate,
	}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	reports := make([]Report, 0, len(users))
	for _, user := range users {
		report := CreateReport(user, reportDate)
		reports = append(reports, report)
	}
	return reports
}

func main() {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 25},
	}
	reportDate := "2025-06-01"
	reports := GenerateUserReports(users, reportDate)
	for _, r := range reports {
		fmt.Printf("ReportID: %d, Date: %s, UserID: %d, Name: %s, Email: %s, Age: %d\n",
			r.ReportID, r.Date, r.ID, r.Name, r.Email, r.Age)
	}
}
