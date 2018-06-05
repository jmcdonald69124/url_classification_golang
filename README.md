# Malicious Domain Classification Naive Bayes / Go

Classification of malicious domains using go

## This code operates in the following way:

This code uses the follwoing package [github.com/jbrukh/bayesian](https://github.com/jbrukh/bayesian)

The inspiration behind this can be found in the following repository, sadly the site that contained the related blog post has been taken down. The training data set is the data.csv file that is found in the repository. [github.com/faizann24/Using-machine-learning-to-detect-malicious-URLs](https://github.com/faizann24/Using-machine-learning-to-detect-malicious-URLs)

The domains are read in and split on the `/`, they labelled either bad or good. The remaining domains are used to test the prediction by running them through the model.

## To run the go tests :

In a terminal navigate to the classifiers folder and run :

```
go test 
```

The test will read in the contents of the domains found in the data.csv file into an array of structures that contain the domain and a flag indicating that the domain is labelled bad or good.

The test will take a percentage of the bad domains as set on line 19 of the test, these domains are removed from the array of bad domains so that no known bad domains will be passed later in the test. An equal number of good domains will also be used to train the model.

Later in the function TestLearn() we iterate over the remaining domains after the good and bad domain arrays have been combined and shuffeled and print out the expected and actual flag value for each entry. 