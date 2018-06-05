package classifiers

import (
	"sort"
	"strings"

	"./models"

	"github.com/jbrukh/bayesian"
)

type NaiveBayesClassifier struct {
	classifier *bayesian.Classifier
	output     []bayesian.Class
}

func (c *NaiveBayesClassifier) Learn(domains []models.Domain) {
	// Split the tld as this will be too common

	c.output = distinctFlags(domains)
	c.classifier = bayesian.NewClassifierTfIdf(c.output...)

	for i := 0; i < len(domains); i++ {
		c.classifier.Learn(strings.Split(domains[i].Domain, "/"), bayesian.Class(domains[i].Flag))
	}
	c.classifier.ConvertTermsFreqToTfIdf()
}

func (c *NaiveBayesClassifier) Predict(domains models.Domain) string {
	// Split the tld as this will be too common

	scores, _, _ := c.classifier.LogScores(strings.Split(domains.Domain, "/"))
	results := models.Results{}
	for i := 0; i < len(scores); i++ {
		results = append(results, models.Result{ID: i, Score: scores[i]})
	}

	sort.Sort(sort.Reverse(results))

	flags := []string{}
	for i := 0; i < len(results); i++ {
		flags = append(flags, string(c.output[results[i].ID]))
	}
	return flags[0]
}

func distinctFlags(domains []models.Domain) []bayesian.Class {
	result := []bayesian.Class{}
	j := 0
	for i := 0; i < len(domains); i++ {
		for j = 0; j < len(result); j++ {
			if domains[i].Flag == string(result[j]) {
				break
			}
		}
		if j == len(result) {
			result = append(result, bayesian.Class(domains[i].Flag))
		}
	}
	return result
}
