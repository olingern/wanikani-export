package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

func getAPIResponse(level string) ApiResponse {
	token := "a3020bbe-c9a3-4b90-89f2-7c0fdc4441c5"
	path := "https://api.wanikani.com/v2/subjects?levels=" + level

	req, err := http.NewRequest("GET", path, nil)
	client := &http.Client{}

	if err != nil {
		log.Fatal("Error creating request")
	}

	req.Header.Add("Wanikani-Revision", "20170710")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Request error")
		log.Fatal(err)
	}

	var apiResp ApiResponse

	json.NewDecoder(resp.Body).Decode(&apiResp)

	return apiResp
}

func getReadingType(readingType string) string {
	if readingType == "onyomi" {
		return "on"
	}

	return "kun"
}

func createFile(filename string, output []string) {
	f, err := os.Create("./" + filename + ".csv")

	if err != nil {
		log.Fatal("Error creating CSV")
	}

	w := bufio.NewWriter(f)
	w.WriteString(strings.Join(output[:], "\n"))
	w.Flush()
}

func createCSV(objectType string, resp ApiResponse) {
	var d KanjiData
	var r string = ""
	var m string = ""
	var out = make([]string, 0)

	for i := 0; i < len(resp.Data); i++ {
		if resp.Data[i].Object == objectType {
			d = resp.Data[i].Data

			for j := 0; j < len(d.Readings); j++ {
				if d.Readings[j].AcceptedAnswer == true {
					r += d.Readings[j].Reading
					if objectType == "kanji" {
						r += " [" + getReadingType(d.Readings[j].Type) + "] "
					}
				}
			}

			for k := 0; k < len(d.Meanings); k++ {
				if d.Meanings[k].AcceptedAnswer == true {
					m += d.Meanings[k].Meaning
					if k != len(d.Meanings)-1 {
						m += " / "
					}
				}
			}

			out = append(out, d.Characters+","+m+","+r)
			r = ""
			m = ""
		}
	}
	createFile(objectType, out)
}

func main() {

	lPtr := flag.String("level", "", "Desired levels. ex. --level=9,10")
	kPtr := flag.Bool("kanji", false, "Create a kanji CSV")
	vPtr := flag.Bool("vocab", false, "Create a vocabulary CSV")
	sPtr := flag.Bool("sentence", false, "Create a sentence CSV")
	flag.Parse()

	if len(*lPtr) < 1 {
		log.Fatal("You must provide a level option as an argument")
	}

	resp := getAPIResponse(*lPtr)

	if *kPtr {
		createCSV("kanji", resp)
	}

	if *vPtr {
		createCSV("vocabulary", resp)
	}

	if *sPtr {
		createCSV("context_sentences", resp)
	}
}
