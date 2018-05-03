package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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

type AG struct {
	A AutoGenerated
}

type AutoGenerated02 struct {
	Node struct {
		Name            string `json:"name"`
		INDICATION      string `json:"INDICATION"`
		TermDescription []struct {
			L1 string `json:"l1"`
			L2 string `json:"l2"`
			L3 string `json:"l3"`
		} `json:"Term description"`
		Q01CAT           string `json:"Q01_CAT"`
		SCORE            string `json:"SCORE"`
		SCOREDEFINITIONS string `json:"SCORE DEFINITIONS"`
		SCOREGRAPHIC     string `json:"SCORE_GRAPHIC"`
	} `json:"node"`
}

func main() {

	//var newNONACS AutoGenerated
	//ctx = context.Background()
	//sa := option.WithCredentialsFile("./SQADMINSDK.json")
	// app, err := firebase.NewApp(context.Background(), nil, sa)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client2, err := app.Firestore(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	getPages()
	//pages := getPages()

	// for x := 0; x < len(pages)-1; x++ {
	// 	fmt.Println(pages[x].toJson)
	// }

	// for _, p := range pages {
	// 	//fmt.Println(p.B.INDEX)
	// 	fmt.Println("===============================================")
	// 	fmt.Println(p.B)
	// 	newNONACS = p
	// 	fmt.Println("============aaaaaaaaaaaaa=====================")
	// 	fmt.Println(newNONACS.toString)
	// 	fmt.Println("===============================================")
	// 	fmt.Println("=========  /////////////////////// ==========")

	// _, err := client2.Collection("NONACSv01").Doc(newNONACS.INDEX).Set(context.Background(), newNONACS)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("=========  /////////////////////// ==========")

	// _, err = client2.Collection("NONACSv01_byTitle").Doc(newNONACS.Title).Set(context.Background(), newNONACS)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("=========  /////////////////////// ==========")
	// }

}

//===================================================================
func getPages() AutoGenerated {
	//read from local file
	raw, err := ioutil.ReadFile("./NACS_TABLE01.json")

	//read from URL json feed

	if err != nil {
		fmt.Print("=========  error in Collection ==========")
		fmt.Println(err.Error())
		log.Fatal(err)
		os.Exit(1)
	}

	var c AutoGenerated
	json.Unmarshal(raw, &c)

	fmt.Println(c.toString)
	//var d AG

	return c
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
