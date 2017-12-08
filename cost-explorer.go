package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/olekukonko/tablewriter"
)

func getDates() costexplorer.DateInterval {
	now := time.Now()
	then := now.AddDate(0, 0, -7)
	dateRange := costexplorer.DateInterval{}
	dateRange.SetEnd(now.Format("2006-01-02"))
	dateRange.SetStart(then.Format("2006-01-02"))
	return dateRange
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	svc := costexplorer.New(sess)
	input := costexplorer.GetCostAndUsageInput{}

	metrics := aws.String("BlendedCost")
	metricsSlice := []*string{metrics}

	group := costexplorer.GroupDefinition{}
	group.SetType("DIMENSION")
	group.SetKey("SERVICE")
	groups := []*costexplorer.GroupDefinition{&group}

	input.SetGroupBy(groups)
	input.SetMetrics(metricsSlice)

	dates := getDates()
	input.SetGranularity("DAILY")
	input.SetTimePeriod(&dates)
	// input.SetGroupBy("SERVICE")
	// fmt.Println(dateRange.Validate())

	resp, err := svc.GetCostAndUsage(&input)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%#v", resp.ResultsByTime[0])
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Service", "Cost"})
	// sorted := sort.Sort(resp.ResultsByTime[0].Groups)
	var data [][]string
	for _, key := range resp.ResultsByTime[0].Groups {
		data = append(data, []string{aws.StringValue(key.Keys[0]), aws.StringValue(key.Metrics["BlendedCost"].Amount)})
		//		table.Append(data)
	}

	sort.Slice(data[:], func(i, j int) bool {
		fmt.Println(i, j)
		// return i[1] < j[1]
		return false
	})

	table.Render()
	// fmt.Println(err)

	// fmt.Printf("%d-%d-%d\n", now.Year(), now.Month(), now.Day())
	//
	// now = now.AddDate(0, 0, 1)
	//
	// fmt.Printf("%d-%d-%d\n", now.Year(), now.Month(), now.Day())

}
