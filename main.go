package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func parseProblem(lines [][]string) []problem {

	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r

}

func problemPuller(fileName string) ([]problem, error) {

	//read al the problem
	if fObj, err := os.Open(fileName); err == nil {

		csvR := csv.NewReader(fObj)
		if cLines, err := csvR.ReadAll(); err == nil {
			return parseProblem(cLines), nil
		} else {
			return nil, fmt.Errorf("error in reading the data  from the %s file %s ", fileName, err.Error())
		}

	} else {
		return nil, fmt.Errorf("error in reading", fileName, err.Error())
	}

	//1 open the file
	//create new reader & read the file

}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}

func main() {

	//input the name of the file

	fName := flag.String("f", "quiz.csv", "path of csv file ")
	timer := flag.Int("t", 20, "timer for the quiz")
	flag.Parse()
	//duration of the timer

	problems, err := problemPuller(*fName)

	if err != nil {
		exit(fmt.Sprintf("some is wrong %s", err.Error()))

	}

	//pull the problem
	//handle the error
	//create
	correctAns := 0
	//timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)

problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("problem %d : %s", i+1, p.q)

		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}

	}
	fmt.Printf("your result is %d out of %d\n", correctAns, len(problems))
	fmt.Printf("press enter to exit")
	<-ansC

}
