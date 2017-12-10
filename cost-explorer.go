package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/bradfitz/slice"
	"github.com/olekukonko/tablewriter"
)

var colors = map[string]int{
	"red":    31,
	"green":  32,
	"yellow": 33,
}

func addColor(s string, i int) string {
	return fmt.Sprintf("\u001b[%dm%v\u001b[0m", i, s)
}

// getDates returns a DateInterval for the last week
func getDates() *costexplorer.DateInterval {
	now := time.Now()
	then := now.AddDate(0, 0, -7)
	dateRange := costexplorer.DateInterval{}
	dateRange.SetEnd(now.Format("2006-01-02"))
	dateRange.SetStart(then.Format("2006-01-02"))
	return &dateRange
}

// covert string to float to string for formatting
func formatNumber(s string) string {
	f, _ := strconv.ParseFloat(s, 64)
	return fmt.Sprintf("%.2f", f)
}

// generate the date headers for the table
func dateHeaders() []string {
	now := time.Now()
	dates := []string{"AWS Service"}
	for i := 7; i > 0; i-- {
		n := now.AddDate(0, 0, -i)
		dates = append(dates, n.Format("01-02"))
	}
	return dates
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	svc := costexplorer.New(sess)

	// Make the API call to cost explorer
	resp, err := svc.GetCostAndUsage((&costexplorer.GetCostAndUsageInput{
		Metrics:     []*string{aws.String("BlendedCost")},
		TimePeriod:  getDates(),
		Granularity: aws.String("DAILY"),
		GroupBy: []*costexplorer.GroupDefinition{
			&costexplorer.GroupDefinition{
				Key:  aws.String("SERVICE"),
				Type: aws.String("DIMENSION"),
			},
		},
	}))

	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println(resp)
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(dateHeaders())
	// append(table)

	// sorted := sort.Sort(resp.ResultsByTime[0].Groups)
	var data [][]string
	for i, group := range resp.ResultsByTime {

		// Sort Map by dollar cost
		slice.Sort(group.Groups[:], func(i, j int) bool {
			a, _ := strconv.ParseFloat(*group.Groups[i].Metrics["BlendedCost"].Amount, 64)
			b, _ := strconv.ParseFloat(*group.Groups[j].Metrics["BlendedCost"].Amount, 64)
			return a > b
		})

		//
		for j, key := range group.Groups {
			// fmt.Println(j, key)
			//dollas := formatNumber(aws.StringValue(key.Metrics["BlendedCost"].Amount))
			// if dollas != "0.00" {
			if i == 0 {
				data = append(data, []string{aws.StringValue(key.Keys[0]), dollas})
			} else {
				if j < len(data) {
					f, _ := strconv.ParseFloat(*key.Metrics["BlendedCost"].Amount, 64)
					data[j] = append(data[j], f)
				}
			}
			// }
		}
	}
	for i := range data {
		for j := range data[i] {
			fmt.Println(j)
			//data[0][i] = addColor(data[0][i], colors["yellow"])
			if j != 0 {
				a, _ := strconv.ParseFloat(data[i][j], 64)
				b, err := strconv.ParseFloat(data[i][j-1], 64)
				if err != nil {
					fmt.Println(err)
				}
				if a > b {
					fmt.Printf("%v is > %v\n", a, b)
					data[i][j] = addColor(data[i][j], colors["red"])
				} else {
					data[i][j] = addColor(data[i][j], colors["green"])
				}
			}
		}
	}

	table.AppendBulk(data)
	table.Render()
}
