package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

type AutoGenerated struct {
	A []struct {
		B struct {
			Title                         string `json:"title"`
			INDEX                         string `json:"INDEX"`
			DEADEND                       string `json:"DEAD_END"`
			Q01IschemicSymptoms           string `json:"Q01 - Ischemic Symptoms"`
			Q02AntiIschemicMedicalTherapy string `json:"Q02 - Anti-ischemic Medical Therapy:"`
			Q03NonInvasiveTestResults     string `json:"Q03 - Non-invasive Test Results:"`
			Q04PRIORCABG                  string `json:"Q04 - PRIOR CABG"`
			Q01                           string `json:"Q01"`
			Q02                           string `json:"Q02"`
			Q03                           string `json:"Q03"`
			Q04                           string `json:"Q04"`
			E01PRIMARY                    string `json:"E01 PRIMARY"`
			E01CABG                       string `json:"E01 CABG"`
			E01PCI                        string `json:"E01 PCI"`
			E02PRIMARY                    string `json:"E02 PRIMARY"`
			E02CABG                       string `json:"E02 CABG"`
			E02PCI                        string `json:"E02 PCI"`
			E03PRIMARY                    string `json:"E03 PRIMARY"`
			E03CABG                       string `json:"E03 CABG"`
			E03PCI                        string `json:"E03 PCI"`
			E04PRIMARY                    string `json:"E04 PRIMARY"`
			E04CABG                       string `json:"E04 CABG"`
			E04PCI                        string `json:"E04 PCI"`
			E05PRIMARY                    string `json:"E05 PRIMARY"`
			E05CABG                       string `json:"E05 CABG "`
			E05PCI                        string `json:"E05 PCI"`
			E05APRIMARY                   string `json:"E05a PRIMARY"`
			E05ACABG                      string `json:"E05a CABG"`
			E05APCI                       string `json:"E05a PCI"`
			E05BPRIMARY                   string `json:"E05b PRIMARY"`
			E05BCABG                      string `json:"E05b CABG"`
			E05BPCI                       string `json:"E05b PCI"`
			E05CPRIMARY                   string `json:"E05c PRIMARY"`
			E05CCABG                      string `json:"E05c CABG"`
			E05CPCI                       string `json:"E05c PCI"`
			E06PRIMARY                    string `json:"E06 PRIMARY"`
			E06CABG                       string `json:"E06 CABG"`
			E06PCI                        string `json:"E06 PCI"`
			E06APRIMARY                   string `json:"E06a PRIMARY"`
			E06ACABG                      string `json:"E06a CABG"`
			E06APCI                       string `json:"E06a PCI"`
			E06BPRIMARY                   string `json:"E06b PRIMARY"`
			E06BCABG                      string `json:"E06b CABG"`
			E06BPCI                       string `json:"E06b PCI"`
			E06CPRIMARY                   string `json:"E06c PRIMARY"`
			E06CCABG                      string `json:"E06c CABG"`
			E06CPCI                       string `json:"E06c PCI"`
			E07PRIMARY                    string `json:"E07 PRIMARY"`
			E07CABG                       string `json:"E07 CABG"`
			E07PCI                        string `json:"E07 PCI"`
			E07APRIMARY                   string `json:"E07a PRIMARY"`
			E07ACABG                      string `json:"E07a CABG"`
			E07APCI                       string `json:"E07a PCI"`
			E07BPRIMARY                   string `json:"E07b PRIMARY"`
			E07BCABG                      string `json:"E07b CABG"`
			E07BPCI                       string `json:"E07b PCI"`
			E07CPRIMARY                   string `json:"E07c PRIMARY"`
			E07CCABG                      string `json:"E07c CABG"`
			E07CPCI                       string `json:"E07c PCI"`
			F01PRIMARY                    string `json:"F01 PRIMARY"`
			F02PRIMARY                    string `json:"F02 PRIMARY"`
			F03PRIMARY                    string `json:"F03 PRIMARY"`
			F03CABGPCI                    string `json:"F03 CABG PCI"`
			F03CABG                       string `json:"F03 CABG"`
			F03PCI                        string `json:"F03 PCI"`
			F04PRIMARY                    string `json:"F04 PRIMARY"`
			F04CABGPCI                    string `json:"F04 CABG PCI"`
			F04CABG                       string `json:"F04 CABG"`
			F04PCI                        string `json:"F04 PCI"`
		} `json:"b"`
	} `json:"a"`
}

type NONACS struct {
	Title                         string `json:"title"`
	INDEX                         string `json:"INDEX"`
	DEADEND                       string `json:"DEAD_END"`
	Q01IschemicSymptoms           string `json:"Q01 - Ischemic Symptoms"`
	Q02AntiIschemicMedicalTherapy string `json:"Q02 - Anti-ischemic Medical Therapy:"`
	Q03NonInvasiveTestResults     string `json:"Q03 - Non-invasive Test Results:"`
	Q04PRIORCABG                  string `json:"Q04 - PRIOR CABG"`
	Q01                           string `json:"Q01"`
	Q02                           string `json:"Q02"`
	Q03                           string `json:"Q03"`
	Q04                           string `json:"Q04"`
	E01PRIMARY                    string `json:"E01 PRIMARY"`
	E01CABG                       string `json:"E01 CABG"`
	E01PCI                        string `json:"E01 PCI"`
	E02PRIMARY                    string `json:"E02 PRIMARY"`
	E02CABG                       string `json:"E02 CABG"`
	E02PCI                        string `json:"E02 PCI"`
	E03PRIMARY                    string `json:"E03 PRIMARY"`
	E03CABG                       string `json:"E03 CABG"`
	E03PCI                        string `json:"E03 PCI"`
	E04PRIMARY                    string `json:"E04 PRIMARY"`
	E04CABG                       string `json:"E04 CABG"`
	E04PCI                        string `json:"E04 PCI"`
	E05PRIMARY                    string `json:"E05 PRIMARY"`
	E05CABG                       string `json:"E05 CABG "`
	E05PCI                        string `json:"E05 PCI"`
	E05APRIMARY                   string `json:"E05a PRIMARY"`
	E05ACABG                      string `json:"E05a CABG"`
	E05APCI                       string `json:"E05a PCI"`
	E05BPRIMARY                   string `json:"E05b PRIMARY"`
	E05BCABG                      string `json:"E05b CABG"`
	E05BPCI                       string `json:"E05b PCI"`
	E05CPRIMARY                   string `json:"E05c PRIMARY"`
	E05CCABG                      string `json:"E05c CABG"`
	E05CPCI                       string `json:"E05c PCI"`
	E06PRIMARY                    string `json:"E06 PRIMARY"`
	E06CABG                       string `json:"E06 CABG"`
	E06PCI                        string `json:"E06 PCI"`
	E06APRIMARY                   string `json:"E06a PRIMARY"`
	E06ACABG                      string `json:"E06a CABG"`
	E06APCI                       string `json:"E06a PCI"`
	E06BPRIMARY                   string `json:"E06b PRIMARY"`
	E06BCABG                      string `json:"E06b CABG"`
	E06BPCI                       string `json:"E06b PCI"`
	E06CPRIMARY                   string `json:"E06c PRIMARY"`
	E06CCABG                      string `json:"E06c CABG"`
	E06CPCI                       string `json:"E06c PCI"`
	E07PRIMARY                    string `json:"E07 PRIMARY"`
	E07CABG                       string `json:"E07 CABG"`
	E07PCI                        string `json:"E07 PCI"`
	E07APRIMARY                   string `json:"E07a PRIMARY"`
	E07ACABG                      string `json:"E07a CABG"`
	E07APCI                       string `json:"E07a PCI"`
	E07BPRIMARY                   string `json:"E07b PRIMARY"`
	E07BCABG                      string `json:"E07b CABG"`
	E07BPCI                       string `json:"E07b PCI"`
	E07CPRIMARY                   string `json:"E07c PRIMARY"`
	E07CCABG                      string `json:"E07c CABG"`
	E07CPCI                       string `json:"E07c PCI"`
	F01PRIMARY                    string `json:"F01 PRIMARY"`
	F02PRIMARY                    string `json:"F02 PRIMARY"`
	F03PRIMARY                    string `json:"F03 PRIMARY"`
	F03CABGPCI                    string `json:"F03 CABG PCI"`
	F03CABG                       string `json:"F03 CABG"`
	F03PCI                        string `json:"F03 PCI"`
	F04PRIMARY                    string `json:"F04 PRIMARY"`
	F04CABGPCI                    string `json:"F04 CABG PCI"`
	F04CABG                       string `json:"F04 CABG"`
	F04PCI                        string `json:"F04 PCI"`
}

type COMPLEX_NONACS struct {
	Title                         string    `json:"title"`
	INDEX                         string    `json:"INDEX"`
	DEADEND                       string    `json:"DEAD_END"`
	Q01IschemicSymptoms           string    `json:"Q01 - Ischemic Symptoms"`
	Q02AntiIschemicMedicalTherapy string    `json:"Q02 - Anti-ischemic Medical Therapy:"`
	Q03NonInvasiveTestResults     string    `json:"Q03 - Non-invasive Test Results:"`
	Q04PRIORCABG                  string    `json:"Q04 - PRIOR CABG"`
	Q01                           string    `json:"Q01"`
	Q02                           string    `json:"Q02"`
	Q03                           string    `json:"Q03"`
	Q04                           string    `json:"Q04"`
	E01PRIMARY                    string    `json:"E01 PRIMARY"`
	DE01PRIMARY                   NACSTable `json:"de01_primary"`
	E01CABG                       string    `json:"E01 CABG"`
	DE01CABG                      NACSTable `json:"de01_cabg"`
	E01PCI                        string    `json:"E01 PCI"`
	DE01PCI                       NACSTable `json:"de01_pci"`
	E02PRIMARY                    string    `json:"E02 PRIMARY"`
	DE02PRIMARY                   NACSTable `json:"de02_primary"`
	E02CABG                       string    `json:"E02 CABG"`
	DE02CABG                      NACSTable `json:"de02_cabg"`
	E02PCI                        string    `json:"E02 PCI"`
	DE02PCI                       NACSTable `json:"de02_pci"`
	E03PRIMARY                    string    `json:"E03 PRIMARY"`
	DE03PRIMARY                   NACSTable `json:"de03_primary"`
	E03CABG                       string    `json:"E03 CABG"`
	DE03CABG                      NACSTable `json:"de03_cabg"`
	E03PCI                        string    `json:"E03 PCI"`
	DE03PCI                       NACSTable `json:"de03_pci"`
	E04PRIMARY                    string    `json:"E04 PRIMARY"`
	DE04PRIMARY                   NACSTable `json:"de04_primary"`
	E04CABG                       string    `json:"E04 CABG"`
	DE04CABG                      NACSTable `json:"de04_cabg"`
	E04PCI                        string    `json:"E04 PCI"`
	DE04PCI                       NACSTable `json:"de04_pci"`
	E05PRIMARY                    string    `json:"E05 PRIMARY"`
	DE05PRIMARY                   NACSTable `json:"de05_primary"`
	E05CABG                       string    `json:"E05 CABG"`
	DE05CABG                      NACSTable `json:"de05_cabg"`
	E05PCI                        string    `json:"E05 PCI"`
	DE05PCI                       NACSTable `json:"de05_pci"`
	E05APRIMARY                   string    `json:"E05a PRIMARY"`
	DE05APRIMARY                  NACSTable `json:"de05a_primary"`
	E05ACABG                      string    `json:"E05a CABG"`
	DE05ACABG                     NACSTable `json:"de05a_cabg"`
	E05APCI                       string    `json:"E05a PCI"`
	DE05APCI                      NACSTable `json:"de05a_pci"`
	E05BPRIMARY                   string    `json:"E05b PRIMARY"`
	DE05BPRIMARY                  NACSTable `json:"de05b_primary"`
	E05BCABG                      string    `json:"E05b CABG"`
	DE05BCABG                     NACSTable `json:"de05b_cabg"`
	E05BPCI                       string    `json:"E05b PCI"`
	DE05BPCI                      NACSTable `json:"de05b_pci"`
	E05CPRIMARY                   string    `json:"E05c PRIMARY"`
	DE05CPRIMARY                  NACSTable `json:"de05c_primary"`
	E05CCABG                      string    `json:"E05c CABG"`
	DE05CCABG                     NACSTable `json:"de05c_cabg"`
	E05CPCI                       string    `json:"E05c PCI"`
	DE05CPCI                      NACSTable `json:"de05c_pci"`
	E06PRIMARY                    string    `json:"E06 PRIMARY"`
	DE06PRIMARY                   NACSTable `json:"de06_primary"`
	E06CABG                       string    `json:"E06 CABG"`
	DE06CABG                      NACSTable `json:"de06_cabg"`
	E06PCI                        string    `json:"E06 PCI"`
	DE06PCI                       NACSTable `json:"de06_pci"`
	E06APRIMARY                   string    `json:"E06a PRIMARY"`
	DE06APRIMARY                  NACSTable `json:"de06a_primary"`
	E06ACABG                      string    `json:"E06a CABG"`
	DE06ACABG                     NACSTable `json:"de06a_cabg"`
	E06APCI                       string    `json:"E06a PCI"`
	DE06APCI                      NACSTable `json:"de06a_pci"`
	E06BPRIMARY                   string    `json:"E06b PRIMARY"`
	DE06BPRIMARY                  NACSTable `json:"de06b_primary"`
	E06BCABG                      string    `json:"E06b CABG"`
	DE06BCABG                     NACSTable `json:"de06b_cabg"`
	E06BPCI                       string    `json:"E06b PCI"`
	DE06BPCI                      NACSTable `json:"de06b_pci"`
	E06CPRIMARY                   string    `json:"E06c PRIMARY"`
	DE06CPRIMARY                  NACSTable `json:"de06c_primary"`
	E06CCABG                      string    `json:"E06c CABG"`
	DE06CCABG                     NACSTable `json:"de06c_cabg"`
	E06CPCI                       string    `json:"E06c PCI"`
	DE06CPCI                      NACSTable `json:"de06c_pci"`
	E07PRIMARY                    string    `json:"E07 PRIMARY"`
	DE07PRIMARY                   NACSTable `json:"de07_primary"`
	E07CABG                       string    `json:"E07 CABG"`
	DE07CABG                      NACSTable `json:"de07_cabg"`
	E07PCI                        string    `json:"E07 PCI"`
	DE07PCI                       NACSTable `json:"de07_pci"`
	E07APRIMARY                   string    `json:"E07a PRIMARY"`
	DE07APRIMARY                  NACSTable `json:"de07a_primary"`
	E07ACABG                      string    `json:"E07a CABG"`
	DE07ACABG                     NACSTable `json:"de07a_cabg"`
	E07APCI                       string    `json:"E07a PCI"`
	DE07APCI                      NACSTable `json:"de07a_pci"`
	E07BPRIMARY                   string    `json:"E07b PRIMARY"`
	DE07BPRIMARY                  NACSTable `json:"de07b_primary"`
	E07BCABG                      string    `json:"E07b CABG"`
	DE07BCABG                     NACSTable `json:"de07b_cabg"`
	E07BPCI                       string    `json:"E07b PCI"`
	DE07BPCI                      NACSTable `json:"de07b_pci"`
	E07CPRIMARY                   string    `json:"E07c PRIMARY"`
	DE07CPRIMARY                  NACSTable `json:"de07c_primary"`
	E07CCABG                      string    `json:"E07c CABG"`
	DE07CCABG                     NACSTable `json:"de07c_cabg"`
	E07CPCI                       string    `json:"E07c PCI"`
	DE07CPCI                      NACSTable `json:"de07c_pci"`
	F01PRIMARY                    string    `json:"F01 PRIMARY"`
	DF01PRIMARY                   NACSTable `json:"df01_primary"`
	F02PRIMARY                    string    `json:"F02 PRIMARY"`
	DF02PRIMARY                   NACSTable `json:"df02_primary"`
	F03PRIMARY                    string    `json:"F03 PRIMARY"`
	DF03PRIMARY                   NACSTable `json:"df03_primary"`
	F03CABGPCI                    string    `json:"F03 CABG PCI"`
	DF03CABGPCI                   NACSTable `json:"df03_cabg_pci"`
	F03CABG                       string    `json:"F03 CABG"`
	DF03CABG                      NACSTable `json:"df03_cabg"`
	F03PCI                        string    `json:"F03 PCI"`
	DF03PCI                       NACSTable `json:"df03_pci"`
	F04PRIMARY                    string    `json:"F04 PRIMARY"`
	DF04PRIMARY                   NACSTable `json:"df04_primary"`
	F04CABGPCI                    string    `json:"F04 CABG PCI"`
	DF04CABGPCI                   string    `json:"df04_cabg_pci"`
	F04CABG                       string    `json:"F04 CABG"`
	DF04CABG                      NACSTable `json:"df04_cabg"`
	F04PCI                        string    `json:"F04 PCI"`
	DF04PCI                       NACSTable `json:"df04_pci"`
}

type NACSTable struct {
	Name             string   `json:"name"`
	INDICATION       string   `json:"indication"`
	TermDescription  []string `json:"termdescription"`
	Q01CAT           string   `json:"q01"`
	SCORE            string   `json:"score"`
	SCOREDEFINITIONS string   `json:"score_def"`
	SCOREGRAPHIC     string   `json:"score_graphic"`
}

func (p AutoGenerated) toString() string {
	return toJson(p)
}

//===================================================================
func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

//===================================================================

//===================================================================
func main() {
	fmt.Println("===============================================")

	var newNONACS NONACS
	//ctx = context.Background()
	sa := option.WithCredentialsFile("./SQADMINSDK.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatal(err)
	}
	client2, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	pages := getPages()
	var tNACSStruct []NONACS
	for _, p := range pages.A {
		//fmt.Println(p.B.INDEX)
		fmt.Println("===============================================")
		fmt.Println(p.B)
		newNONACS = p.B
		fmt.Println("============aaaaaaaaaaaaa=====================")
		fmt.Println(newNONACS.INDEX)
		fmt.Println("===============================================")
		fmt.Println("=========  /////////////////////// ==========")

		_, err := client2.Collection("NONACSv03").Doc(newNONACS.INDEX).Set(context.Background(), newNONACS)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("=========  /////////////////////// ==========")

		_, err = client2.Collection("NONACSv01_byTitleV03").Doc(newNONACS.Title).Set(context.Background(), newNONACS)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("=========  /////////////////////// ==========")

		//Add to TNACSStruct
		tNACSStruct = append(tNACSStruct, newNONACS)

	}
	defer client2.Close()

	// create local json files

	fmt.Println("=========  /////////////////////// ==========")
	fmt.Println("=========  writing JSON file   ==========")
	jsonFile, err := os.Create("./NACSv03_OUT.json")

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	//b, err := json.Marshal(pages.A)
	b, err := json.Marshal(tNACSStruct)
	jsonFile.Write(b)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())

}

//===================================================================

//===================================================================
func getPages() AutoGenerated {
	raw, err := ioutil.ReadFile("./nonACS.json")
	if err != nil {
		fmt.Print("=========  error in Collection ==========")
		fmt.Println(err.Error())
		log.Fatal(err)
		os.Exit(1)
	}

	var c AutoGenerated
	json.Unmarshal(raw, &c)
	return c
}
