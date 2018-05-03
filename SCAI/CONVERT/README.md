The original TableDB for the SCIA project was complicated
with many remnants of experiments and not used fields.
This project mines the db. and pumps it into Firestore
using Firestore admin Go SDK
I can also convert a lot of the fields on the fly.

convert the table data from NACS to JSON to be imported into FIRESTORE

http://nacs.scai-qit.org/nacs-table01-dump



type AutoGenerated struct {
	Nodes []struct {
		Node struct {
			Name             string `json:"name"`
			INDICATION       string `json:"INDICATION"`
			TermDescription  string `json:"Term description"`
			Q01CAT           string `json:"Q01_CAT"`
			SCORE            string `json:"SCORE"`
			SCOREDEFINITIONS string `json:"SCORE DEFINITIONS"`
			SCOREGRAPHIC     string `json:"SCORE_GRAPHIC"`
		} `json:"node"`
	} `json:"nodes"`
}


//with description chopped up
type AutoGenerated struct {
	Node struct {
		Name            string `json:"name"`
		INDICATION      string `json:"INDICATION"`
		Name            string `json:"Name"`
		TermDescription []string `json:"Term description"`
		Q01CAT           string `json:"Q01_CAT"`
		SCORE            string `json:"SCORE"`
		SCOREDEFINITIONS string `json:"SCORE DEFINITIONS"`
		SCOREGRAPHIC     string `json:"SCORE_GRAPHIC"`
	} `json:"node"`
}
