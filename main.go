package main

import (
	"fmt"
	_ "strconv"
	"sync"
	"time"

	"github.com/correavitor4/concorrentes2/src/assetAnalyser"
)

func main() {
	//1. Define csv files names
	csvFilesNames := []string{"BBAS3.csv", "VALE3.SA.csv", "GGBR4.SA.csv", "BBDC4.SA.csv", "JBSS3.SA.csv"}
	startTime1 := time.Now()
	startAnalysisSingleThread(csvFilesNames)
	elapsed1 := time.Since(startTime1)

	startTime2 := time.Now()
	startAnalysisMultiThread(csvFilesNames)
	elapsed2 := time.Since(startTime2)

	fmt.Println("Single thread elapsed time: ", elapsed1)
	fmt.Println("Multi thread elapsed time: ", elapsed2)

}

func startAnalysisSingleThread(csvFilesNames []string) {
	fmt.Println("Starting single thread analysis")
	for _, csvFilename := range csvFilesNames {
		result := assetAnalyser.AnalyseAssetByCsv(csvFilename)
		// fmt.Println()
		fmt.Println(result)
	}
	fmt.Println("Ending single thread analysis")
}

func startAnalysisMultiThread(csvFilesNames []string) {
	println("Starting multi thread analysis")
	var wg = sync.WaitGroup{}
	wg.Add(len(csvFilesNames))
	for _, csvFilename := range csvFilesNames {
		go anl(csvFilename, &wg)
	}
	wg.Wait()
	println("Ending multi thread analysis")
}

func anl(csvFilename string, wg *sync.WaitGroup) {
	result := assetAnalyser.AnalyseAssetByCsv(csvFilename)
	println(result)
	wg.Done()
}
