package assetAnalyser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func AnalyseAssetByCsv(csvFilename string) string {
	//1. Open the file
	csvFile, err := os.Open("./src/csv/" + csvFilename)
	//2. Check for errors
	if err != nil {
		fmt.Println(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	//3. Create the array with struct that represents the daily results
	var results = (make([]dailyResult, 0))

	//4. Loop through the lines and add them to the array
	for _, line := range csvLines {
		//1. Create the struct object
		result := dailyResult{}

		//2. Parse the date and fill the in struct
		time, err := time.Parse("2006-04-05", line[0])
		if err == nil {
			result.date = time
		}

		//3. Parse the open price and fill the in struct
		low, err := strconv.ParseFloat(line[1], 64)

		//4. Parse the high price and fill the in struct
		high, err := strconv.ParseFloat(line[2], 64)

		if err == nil {
			result.averagePrice = (low + high) / 2
		}

		if err == nil {
			//5. Append the struct to the array
			results = append(results, result)
		}

	}

	//5. Close the file
	csvFile.Close()

	mostProfitable := execAnalysis(results)

	str := "The best investiment for " + csvFilename + " was buy in " + mostProfitable.buyDate.Format("2006-04-05") + " with the value of " + strconv.FormatFloat(mostProfitable.buyValue, 'f', 2, 64) + " and sell at " + mostProfitable.sellDate.Format("2006-04-05") + " with the value of " + strconv.FormatFloat(mostProfitable.sellValue, 'f', 2, 64) + " with a profit of " + strconv.FormatFloat(mostProfitable.profit, 'f', 2, 64)
	return str
}

func execAnalysis(results []dailyResult) *MostProfitableInvestiment {
	//1. Create the struct object
	var mostProfitable MostProfitableInvestiment

	//2. Loop through the array and find the most profitable day
	for i := 0; i < (len(results) - 1); i++ {
		//Current day
		currentDay := results[i]
		for j := i + 1; j < len(results); j++ {
			currentProfit := results[j].averagePrice - currentDay.averagePrice
			if currentProfit > mostProfitable.profit {
				mostProfitable.profit = currentProfit
				mostProfitable.buyDate = currentDay.date
				mostProfitable.sellDate = results[j].date
				mostProfitable.buyValue = currentDay.averagePrice
				mostProfitable.sellValue = results[j].averagePrice
			}
		}
	}

	return &mostProfitable
}

type MostProfitableInvestiment struct {
	buyDate   time.Time
	sellDate  time.Time
	buyValue  float64
	sellValue float64
	profit    float64
}

type dailyResult struct {
	date         time.Time
	averagePrice float64
}
