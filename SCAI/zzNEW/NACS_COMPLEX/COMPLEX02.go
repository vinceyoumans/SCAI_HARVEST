//  Will combine NACS with NACS_table
// Create a File.
// Then Submit to FIRESTORE

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var TableMap = map[string]NACSTable{}

func main() {

	var ComplexNACS []COMPLEX_NONACS
	NACS01json, err := os.Open("./NACS_TABLE_EDIT.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened NACS_TABLE_EDIT.json")

	fmt.Println(NACS01json)
	defer NACS01json.Close()
	var NACSTableC []NACSTable

	NACS01jsonA, _ := ioutil.ReadAll(NACS01json)
	json.Unmarshal(NACS01jsonA, &NACSTableC)
	for i := 0; i < len(NACSTableC); i++ {
		fmt.Println("==========================================")
		fmt.Println("INdication: " + NACSTableC[i].INDICATION)
		fmt.Println("Name: " + NACSTableC[i].Name)
		fmt.Println("Q01CAT: " + NACSTableC[i].Q01CAT)
		fmt.Println("SCORE: " + NACSTableC[i].SCORE)
		fmt.Println("SCOREDEF: " + NACSTableC[i].SCOREDEFINITIONS)
		fmt.Println("SCOREGRAPHIC: " + NACSTableC[i].SCOREGRAPHIC)
		fmt.Println("TERM DESCRIP: " + NACSTableC[i].TermDescription[0])
		fmt.Println("==========================================")
	}

	fmt.Println("============Stating Table MaP=================")
	var t string

	for i := 0; i < len(NACSTableC); i++ {
		t = NACSTableC[i].Name
		TableMap[t] = NACSTableC[i]
	}
	fmt.Println(TableMap["69 - PCI"].Name)

	fmt.Println("============   ended making table map   =================")

	//=====================================================================
	fmt.Println("============  starting Map of NACS  =================")

	NACS01json, err = os.Open("./NACSv03_OUT.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened NACSv03_OUT.json")
	// defer the closing of our jsonFile so that we can parse it later on

	fmt.Println(NACS01json)
	defer NACS01json.Close()

	// we initialize our Users array
	//var users Users
	var NACS []NONACS

	// read our opened xmlFile as a byte array.
	NACS02jsonA, _ := ioutil.ReadAll(NACS01json)
	json.Unmarshal(NACS02jsonA, &NACS)
	for i := 0; i < len(NACS); i++ {
		fmt.Println("==========================================")
		fmt.Println("Index: " + NACS[i].INDEX)
		fmt.Println("DeadEnd" + NACS[i].DEADEND)
		fmt.Println("E01CABG: " + NACS[i].E01CABG)
		fmt.Println("E01PCI: " + NACS[i].E01PCI)
		fmt.Println("==========================================")
	}

	fmt.Println("============Building ComplexNACS=================")

	var TempComplexNACS COMPLEX_NONACS

	for i := 0; i < len(NACS); i++ {
		fmt.Println("==========================================")
		fmt.Println("Index: " + NACS[i].INDEX)

		//pppppppppppppppppppppppppppppppppp

		TempComplexNACS.Title = NACS[i].Title                                                 //    string    `json:"title"`
		TempComplexNACS.INDEX = NACS[i].INDEX                                                 //                     string    `json:"INDEX"`
		TempComplexNACS.DEADEND = NACS[i].DEADEND                                             //                     string    `json:"DEAD_END"`
		TempComplexNACS.Q01IschemicSymptoms = NACS[i].Q01IschemicSymptoms                     //          string    `json:"Q01 - Ischemic Symptoms"`
		TempComplexNACS.Q02AntiIschemicMedicalTherapy = NACS[i].Q02AntiIschemicMedicalTherapy // string    `json:"Q02 - Anti-ischemic Medical Therapy:"`
		TempComplexNACS.Q03NonInvasiveTestResults = NACS[i].Q03NonInvasiveTestResults         //  string    `json:"Q03 - Non-invasive Test Results:"`
		TempComplexNACS.Q04PRIORCABG = NACS[i].Q04PRIORCABG                                   //                  string    `json:"Q04 - PRIOR CABG"`
		TempComplexNACS.Q01 = NACS[i].Q01                                                     //                         string    `json:"Q01"`
		TempComplexNACS.Q02 = NACS[i].Q02                                                     //                         string    `json:"Q02"`
		TempComplexNACS.Q03 = NACS[i].Q03                                                     //                         string    `json:"Q03"`
		TempComplexNACS.Q04 = NACS[i].Q04                                                     //                        string    `json:"Q04"`
		TempComplexNACS.E01PRIMARY = NACS[i].E01PRIMARY                                       //                 string    `json:"E01 PRIMARY"`

		TempComplexNACS.DE01PRIMARY = getTableMap(NACS[i].E01PRIMARY) //	NACSTable `json:"de01_primary"`
		//TempComplexNACS.DE01PRIMARY = getTableMap("22-Asymptomatic") 		//	NACSTable `json:"de01_primary"`

		TempComplexNACS.E01CABG = NACS[i].E01CABG                       //	string    `json:"E01 CABG"`
		TempComplexNACS.DE01CABG = getTableMap(NACS[i].E01CABG)         //	NACSTable `json:"de01_cabg"`
		TempComplexNACS.E01PCI = NACS[i].E01PCI                         //	string    `json:"E01 PCI"`
		TempComplexNACS.DE01PCI = getTableMap(NACS[i].E01PCI)           // 	NACSTable `json:"de01_pci"`
		TempComplexNACS.E02PRIMARY = NACS[i].E02PRIMARY                 //	string    `json:"E02 PRIMARY"`
		TempComplexNACS.DE02PRIMARY = getTableMap(NACS[i].E02PRIMARY)   //	NACSTable `json:"de02_primary"`
		TempComplexNACS.E02CABG = NACS[i].E02CABG                       //	string    `json:"E02 CABG"`
		TempComplexNACS.DE02CABG = getTableMap(NACS[i].E02CABG)         //	NACSTable `json:"de02_cabg"`
		TempComplexNACS.E02PCI = NACS[i].E02PCI                         //	string    `json:"E02 PCI"`
		TempComplexNACS.DE02PCI = getTableMap(NACS[i].E02PCI)           //	NACSTable `json:"de02_pci"`
		TempComplexNACS.E03PRIMARY = NACS[i].E03PRIMARY                 //	string    `json:"E03 PRIMARY"`
		TempComplexNACS.DE03PRIMARY = getTableMap(NACS[i].E03PRIMARY)   //	NACSTable `json:"de03_primary"`
		TempComplexNACS.E03CABG = NACS[i].E03CABG                       //	string    `json:"E03 CABG"`
		TempComplexNACS.DE03CABG = getTableMap(NACS[i].E03CABG)         //	NACSTable `json:"de03_cabg"`
		TempComplexNACS.E03PCI = NACS[i].E03PCI                         //	string    `json:"E03 PCI"`
		TempComplexNACS.DE03PCI = getTableMap(NACS[i].E03PCI)           //	NACSTable `json:"de03_pci"`
		TempComplexNACS.E04PRIMARY = NACS[i].E04PRIMARY                 //	string    `json:"E04 PRIMARY"`
		TempComplexNACS.DE04PRIMARY = getTableMap(NACS[i].E04PRIMARY)   //	NACSTable `json:"de04_primary"`
		TempComplexNACS.E04CABG = NACS[i].E04CABG                       // 	string    `json:"E04 CABG"`
		TempComplexNACS.DE04CABG = getTableMap(NACS[i].E04CABG)         //	NACSTable `json:"de04_cabg"`
		TempComplexNACS.E04PCI = NACS[i].E04PCI                         //	string    `json:"E04 PCI"`
		TempComplexNACS.DE04PCI = getTableMap(NACS[i].E04PCI)           //	NACSTable `json:"de04_pci"`
		TempComplexNACS.E05PRIMARY = NACS[i].E05PRIMARY                 //	string    `json:"E05 PRIMARY"`
		TempComplexNACS.DE05PRIMARY = getTableMap(NACS[i].E05PRIMARY)   //	NACSTable `json:"de05_primary"`
		TempComplexNACS.E05CABG = NACS[i].E05CABG                       //	string    `json:"E05 CABG"`
		TempComplexNACS.DE05CABG = getTableMap(NACS[i].E05CABG)         //	NACSTable `json:"de05_cabg"`
		TempComplexNACS.E05PCI = NACS[i].E05PCI                         //	string    `json:"E05 PCI"`
		TempComplexNACS.DE05PCI = getTableMap(NACS[i].E05PCI)           //	NACSTable `json:"de05_pci"`
		TempComplexNACS.E05APRIMARY = NACS[i].E05APRIMARY               //	string    `json:"E05a PRIMARY"`
		TempComplexNACS.DE05APRIMARY = getTableMap(NACS[i].E05APRIMARY) //	NACSTable `json:"de05a_primary"`
		TempComplexNACS.E05ACABG = NACS[i].E05ACABG                     //	string    `json:"E05a CABG"`
		TempComplexNACS.DE05ACABG = getTableMap(NACS[i].E05ACABG)       //	NACSTable `json:"de05a_cabg"`
		TempComplexNACS.E05APCI = NACS[i].E05APCI                       //	string    `json:"E05a PCI"`
		TempComplexNACS.DE05APCI = getTableMap(NACS[i].E05APCI)         //	NACSTable `json:"de05a_pci"`
		TempComplexNACS.E05BPRIMARY = NACS[i].E05BPRIMARY               //	string    `json:"E05b PRIMARY"`
		TempComplexNACS.DE05BPRIMARY = getTableMap(NACS[i].E05BPRIMARY) //	NACSTable `json:"de05b_primary"`
		TempComplexNACS.E05BCABG = NACS[i].E05BCABG                     //	string    `json:"E05b CABG"`
		TempComplexNACS.DE05BCABG = getTableMap(NACS[i].E05BCABG)       //	NACSTable `json:"de05b_cabg"`
		TempComplexNACS.E05BPCI = NACS[i].E05BPCI                       //	string    `json:"E05b PCI"`
		TempComplexNACS.DE05BPCI = getTableMap(NACS[i].E05BPCI)         //	NACSTable `json:"de05b_pci"`
		TempComplexNACS.E05CPRIMARY = NACS[i].E05CPRIMARY               //	string    `json:"E05c PRIMARY"`
		TempComplexNACS.DE05CPRIMARY = getTableMap(NACS[i].E05CPRIMARY) //	NACSTable `json:"de05c_primary"`
		TempComplexNACS.E05CCABG = NACS[i].E05CCABG                     //	string    `json:"E05c CABG"`
		TempComplexNACS.DE05CCABG = getTableMap(NACS[i].E05CCABG)       //	NACSTable `json:"de05c_cabg"`
		TempComplexNACS.E05CPCI = NACS[i].E05CPCI                       //	string    `json:"E05c PCI"`
		TempComplexNACS.DE05CPCI = getTableMap(NACS[i].E05CPCI)         //	NACSTable `json:"de05c_pci"`
		TempComplexNACS.E06PRIMARY = NACS[i].E06PRIMARY                 //	string    `json:"E06 PRIMARY"`
		TempComplexNACS.DE06PRIMARY = getTableMap(NACS[i].E06PRIMARY)   //	NACSTable `json:"de06_primary"`
		TempComplexNACS.E06CABG = NACS[i].E06CABG                       //	string    `json:"E06 CABG"`
		TempComplexNACS.DE06CABG = getTableMap(NACS[i].E06CABG)         //	NACSTable `json:"de06_cabg"`
		TempComplexNACS.E06PCI = NACS[i].E06PCI                         //	string    `json:"E06 PCI"`
		TempComplexNACS.DE06PCI = getTableMap(NACS[i].E06PCI)           //	NACSTable `json:"de06_pci"`
		TempComplexNACS.E06APRIMARY = NACS[i].E06APRIMARY               //	string    `json:"E06a PRIMARY"`
		TempComplexNACS.DE06APRIMARY = getTableMap(NACS[i].E06APRIMARY) //	NACSTable `json:"de06a_primary"`
		TempComplexNACS.E06ACABG = NACS[i].E06ACABG                     //	string    `json:"E06a CABG"`
		TempComplexNACS.DE06ACABG = getTableMap(NACS[i].E06ACABG)       //	NACSTable `json:"de06a_cabg"`
		TempComplexNACS.E06APCI = NACS[i].E06APCI                       //	string    `json:"E06a PCI"`
		TempComplexNACS.DE06APCI = getTableMap(NACS[i].E06APCI)         //	NACSTable `json:"de06a_pci"`
		TempComplexNACS.E06BPRIMARY = NACS[i].E06BPRIMARY               //	string    `json:"E06b PRIMARY"`
		TempComplexNACS.DE06BPRIMARY = getTableMap(NACS[i].E06BPRIMARY) //	NACSTable `json:"de06b_primary"`
		TempComplexNACS.E06BCABG = NACS[i].E06BCABG                     //	string    `json:"E06b CABG"`
		TempComplexNACS.DE06BCABG = getTableMap(NACS[i].E06BCABG)       //	NACSTable `json:"de06b_cabg"`
		TempComplexNACS.E06BPCI = NACS[i].E06BPCI                       //	string    `json:"E06b PCI"`
		TempComplexNACS.DE06BPCI = getTableMap(NACS[i].E06BPCI)         //	NACSTable `json:"de06b_pci"`
		TempComplexNACS.E06CPRIMARY = NACS[i].E06CPRIMARY               //	string    `json:"E06c PRIMARY"`
		TempComplexNACS.DE06CPRIMARY = getTableMap(NACS[i].E06CPRIMARY) //	NACSTable `json:"de06c_primary"`
		TempComplexNACS.E06CCABG = NACS[i].E06CCABG                     //	string    `json:"E06c CABG"`
		TempComplexNACS.DE06CCABG = getTableMap(NACS[i].E06CCABG)       //	NACSTable `json:"de06c_cabg"`
		TempComplexNACS.E06CPCI = NACS[i].E06CPCI                       //	string    `json:"E06c PCI"`
		TempComplexNACS.DE06CPCI = getTableMap(NACS[i].E06CPCI)         //	NACSTable `json:"de06c_pci"`
		TempComplexNACS.E07PRIMARY = NACS[i].E07PRIMARY                 //	string    `json:"E07 PRIMARY"`
		TempComplexNACS.DE07PRIMARY = getTableMap(NACS[i].E07PRIMARY)   //	NACSTable `json:"de07_primary"`
		TempComplexNACS.E07CABG = NACS[i].E07CABG                       //	string    `json:"E07 CABG"`
		TempComplexNACS.DE07CABG = getTableMap(NACS[i].E07CABG)         //	NACSTable `json:"de07_cabg"`
		TempComplexNACS.E07PCI = NACS[i].E07PCI                         //	string    `json:"E07 PCI"`
		TempComplexNACS.DE07PCI = getTableMap(NACS[i].E07PCI)           //	NACSTable `json:"de07_pci"`
		TempComplexNACS.E07APRIMARY = NACS[i].E07APRIMARY               //	string    `json:"E07a PRIMARY"`
		TempComplexNACS.DE07APRIMARY = getTableMap(NACS[i].E07APRIMARY) //	NACSTable `json:"de07a_primary"`
		TempComplexNACS.E07ACABG = NACS[i].E07ACABG                     //	string    `json:"E07a CABG"`
		TempComplexNACS.DE07ACABG = getTableMap(NACS[i].E07ACABG)       //	NACSTable `json:"de07a_cabg"`
		TempComplexNACS.E07APCI = NACS[i].E07APCI                       //	string    `json:"E07a PCI"`
		TempComplexNACS.DE07APCI = getTableMap(NACS[i].E07APCI)         //	NACSTable `json:"de07a_pci"`
		TempComplexNACS.E07BPRIMARY = NACS[i].E07BPRIMARY               //	string    `json:"E07b PRIMARY"`
		TempComplexNACS.DE07BPRIMARY = getTableMap(NACS[i].E07BPRIMARY) //	NACSTable `json:"de07b_primary"`
		TempComplexNACS.E07BCABG = NACS[i].E07BCABG                     //	string    `json:"E07b CABG"`
		TempComplexNACS.DE07BCABG = getTableMap(NACS[i].E07BCABG)       //	NACSTable `json:"de07b_cabg"`
		TempComplexNACS.E07BPCI = NACS[i].E07BPCI                       //	string    `json:"E07b PCI"`
		TempComplexNACS.DE07BPCI = getTableMap(NACS[i].E07BPCI)         //	NACSTable `json:"de07b_pci"`
		TempComplexNACS.E07CPRIMARY = NACS[i].E07CPRIMARY               //	string    `json:"E07c PRIMARY"`
		TempComplexNACS.DE07CPRIMARY = getTableMap(NACS[i].E07CPRIMARY) //	NACSTable `json:"de07c_primary"`
		TempComplexNACS.E07CCABG = NACS[i].E07CCABG                     //	string    `json:"E07c CABG"`
		TempComplexNACS.DE07CCABG = getTableMap(NACS[i].E07CCABG)       //	NACSTable `json:"de07c_cabg"`
		TempComplexNACS.E07CPCI = NACS[i].E07CPCI                       //	string    `json:"E07c PCI"`
		TempComplexNACS.DE07CPCI = getTableMap(NACS[i].E07CPCI)         //	NACSTable `json:"de07c_pci"`
		TempComplexNACS.F01PRIMARY = NACS[i].F01PRIMARY                 //	string    `json:"F01 PRIMARY"`
		TempComplexNACS.DF01PRIMARY = getTableMap(NACS[i].F01PRIMARY)   //	NACSTable `json:"df01_primary"`
		TempComplexNACS.F02PRIMARY = NACS[i].F02PRIMARY                 //	string    `json:"F02 PRIMARY"`
		TempComplexNACS.DF02PRIMARY = getTableMap(NACS[i].F02PRIMARY)   //	NACSTable `json:"df02_primary"`
		TempComplexNACS.F03PRIMARY = NACS[i].F03PRIMARY                 //	string    `json:"F03 PRIMARY"`
		TempComplexNACS.DF03PRIMARY = getTableMap(NACS[i].F03PRIMARY)   //	NACSTable `json:"df03_primary"`
		TempComplexNACS.F03CABGPCI = NACS[i].F03CABGPCI                 //	string    `json:"F03 CABG PCI"`
		TempComplexNACS.DF03CABGPCI = getTableMap(NACS[i].F03CABGPCI)   //	NACSTable `json:"df03_cabg_pci"`
		TempComplexNACS.F03CABG = NACS[i].F03CABG                       //	string    `json:"F03 CABG"`
		TempComplexNACS.DF03CABG = getTableMap(NACS[i].F03CABG)         //	NACSTable `json:"df03_cabg"`
		TempComplexNACS.F03PCI = NACS[i].F03PCI                         //	string    `json:"F03 PCI"`
		TempComplexNACS.DF03PCI = getTableMap(NACS[i].F03PCI)           //	NACSTable `json:"df03_pci"`
		TempComplexNACS.F04PRIMARY = NACS[i].F04PRIMARY                 //	string    `json:"F04 PRIMARY"`
		TempComplexNACS.DF04PRIMARY = getTableMap(NACS[i].F04PRIMARY)   //	NACSTable `json:"df04_primary"`
		TempComplexNACS.F04CABGPCI = NACS[i].F04CABGPCI                 //	string    `json:"F04 CABG PCI"`
		TempComplexNACS.DF04CABGPCI = getTableMap(NACS[i].F04CABGPCI)   //	NACSTable    `json:"df04_cabg_pci"`
		TempComplexNACS.F04CABG = NACS[i].F04CABG                       //	string    `json:"F04 CABG"`
		TempComplexNACS.DF04CABG = getTableMap(NACS[i].F04CABG)         //	NACSTable `json:"df04_cabg"`
		TempComplexNACS.F04PCI = NACS[i].F04PCI                         //	string    `json:"F04 PCI"`
		TempComplexNACS.DF04PCI = getTableMap(NACS[i].F04PCI)           //	NACSTable `json:"df04_pci"`

		fmt.Println("TempComplexNACS.INDEXIndex: " + TempComplexNACS.INDEX)
		fmt.Println("==========================================")
		ComplexNACS = append(ComplexNACS, TempComplexNACS)

	}

	// //   Save TempComplex to a ComplexNACS Slice
	// ComplexJson, _ := json.Marshal(ComplexNACS)
	// err = ioutil.WriteFile("COMPLEX_NACSv01.json", ComplexJson, 0644)
	// fmt.Printf("%+v", ComplexJson)

	// create local json files

	fmt.Println("=========  /////////////////////// ==========")
	fmt.Println("=========  writing JSON file   ==========")
	jsonFile, err := os.Create("./ComplexNACSv03_OUT.json")

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	//b, err := json.Marshal(pages.A)
	b, err := json.Marshal(ComplexNACS)
	jsonFile.Write(b)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())

	sa := option.WithCredentialsFile("./scai-qit-firebase-adminsdk.json")

	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatal(err)
	}
	client2, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(ComplexNACS); i++ {
		_, err := client2.Collection("COMPLEX_NONACS").Doc(ComplexNACS[i].INDEX).Set(context.Background(), ComplexNACS[i])
		if err != nil {
			log.Fatal(err)
		}
	}

}

//=============================================================
func getTableMap(lookup string) NACSTable {

	var lookUptable NACSTable = TableMap[lookup]
	fmt.Println("======================")
	fmt.Println(lookUptable.INDICATION)
	fmt.Println(lookUptable)
	return lookUptable
}

//=============================================================
type NACSTable struct {
	Name             string   `json:"name"`
	INDICATION       string   `json:"indication"`
	TermDescription  []string `json:"termdescription"`
	Q01CAT           string   `json:"q01"`
	SCORE            string   `json:"score"`
	SCOREDEFINITIONS string   `json:"score_def"`
	SCOREGRAPHIC     string   `json:"score_graphic"`
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
	DF04CABGPCI                   NACSTable `json:"df04_cabg_pci"`
	F04CABG                       string    `json:"F04 CABG"`
	DF04CABG                      NACSTable `json:"df04_cabg"`
	F04PCI                        string    `json:"F04 PCI"`
	DF04PCI                       NACSTable `json:"df04_pci"`
}
