package models

type Domain struct {
	Domain string
	Flag   string // Malicious or Benign
}

type Classifier interface {
	Learn(domains []Domain)
	Predict(domain Domain) string
}

type DomainModel struct {
	Classifier Classifier
}

func (model *DomainModel) Learn(domains []Domain) {
	model.Classifier.Learn(domains)
}

func (model *DomainModel) Predict(domain Domain) string {
	return model.Classifier.Predict(domain)
}

type Model interface {
	Learn(domains []Domain)
	Predict(domain Domain) string
}
