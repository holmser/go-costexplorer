package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/olekukonko/tablewriter"
)

// getDates returns a DateInterval for the last week
func getDates() *costexplorer.DateInterval {
	now := time.Now()
	then := now.AddDate(0, 0, -7)
	dateRange := costexplorer.DateInterval{}
	dateRange.SetEnd(now.Format("2006-01-02"))
	dateRange.SetStart(then.Format("2006-01-02"))
	return &dateRange
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	svc := costexplorer.New(sess)

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
		fmt.Println(resp)
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Service", "Cost"})
	// sorted := sort.Sort(resp.ResultsByTime[0].Groups)
	var data [][]string
	for _, key := range resp.ResultsByTime[0].Groups {
		data = append(data, []string{aws.StringValue(key.Keys[0]), aws.StringValue(key.Metrics["BlendedCost"].Amount)})
	}
	table.AppendBulk(data)

	table.Render()
}
