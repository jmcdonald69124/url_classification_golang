package classifiers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"

	"./models"
)

var Successful = 0
var Failed = 0
var BadPercentage = .20

var Bad = []models.Domain{}
var Good = []models.Domain{}

func Shuffle(slice []models.Domain) {
	// Fisher–Yates algorithm:
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func CreateTrainingDomains() []models.Domain {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.LazyQuotes = true
	reader.Comma = '≈'
	var domains []models.Domain
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if line[1] == "bad" {
			Bad = append(Bad, models.Domain{
				Domain: line[0],
				Flag:   line[1],
			})
		} else if line[1] == "good" {
			Good = append(Good, models.Domain{
				Domain: line[0],
				Flag:   line[1],
			})
		}
	}

	Shuffle(Bad)
	Shuffle(Good)

	badTotal := len(Bad)
	goodTotal := badTotal

	fmt.Print("\n Number of bad domains:  ", badTotal)
	fmt.Print("\n Number of good domains: ", goodTotal)

	// Randomly select bad domains for training data
	// removing them from the Bad slice of domains and
	// inserting them into the domains slice that is passed
	// back by this function

	trainingDataSize := float64(badTotal) * BadPercentage

	fmt.Print("\n Training data size: ", trainingDataSize)
	var element = 0
	var elementG = 0
	var deleteTo = 0
	for i := 0; i < int(trainingDataSize); i++ {
		// Random number between 0 and length of bad data slice
		element = rand.Intn(len(Bad))
		elementG = rand.Intn(len(Good))
		// fmt.Printf("\n Adding element [%v]: ", element)

		domains = append(domains, Bad[element])
		domains = append(domains, Good[elementG])
		// Remove the element from the bad slice and then resize it
		deleteTo = deleteTo + 1
		copy(Bad[element:], Bad[deleteTo:])
		copy(Good[element:], Good[deleteTo:])
		Bad = Bad[:len(Bad)-deleteTo+i]
		Good = Good[:len(Good)-deleteTo+i]
	}
	fmt.Print("\n Domain array data size: ", len(domains))
	fmt.Print("\n Bad array data size: ", len(Bad))
	fmt.Print("\n Good array data size: ", len(Good))
	Shuffle(domains)
	return domains
}

func CreateValidationDomains() []models.Domain {
	Bad := append(Bad, Good...)
	Shuffle(Bad)
	return Bad
}

func TestLearn(t *testing.T) {
	nbModel := models.DomainModel{Classifier: &NaiveBayesClassifier{}}
	trainingSet := CreateTrainingDomains()
	validationSet := CreateValidationDomains()

	nbModel.Learn(trainingSet)

	for i := 0; i < len(validationSet); i++ {
		input := validationSet[i].Domain
		expected := validationSet[i].Flag
		actual := nbModel.Predict(validationSet[i])
		Assert(t, expected, actual, input)
	}

	fmt.Print(
		"\n Successful        | ", Successful,
		"\n Failed            | ", Failed,
		"\n Total             | ", Successful+Failed,
		"\n Success Rate      | ", float32(Successful)/float32((Successful+Failed)),
		"\n ********************************",
	)

}

func Assert(t *testing.T, expected string, actual string, input string) {
	if actual != expected {
		t.Error(
			"\nFOR:       ", input,
			"\nEXPECTED:  ", expected,
			"\nACTUAL:    ", actual,
		)
		Failed = Failed + 1
	} else {
		Successful = Successful + 1
	}
}

func AssertList(t *testing.T, expected string, actual []string, input string) {
	for i := 0; i < len(actual); i++ {
		if actual[i] == expected {
			Successful = Successful + 1
			return
		}
	}
	Failed = Failed + 1
	t.Error(
		"\nFOR:       ", input,
		"\nEXPECTED:  ", expected,
		"\nACTUAL:    ", strings.Join(actual, ","),
	)
}
