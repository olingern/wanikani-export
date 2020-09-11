package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// SuccessColor green
const SuccessColor = "\033[1;32m%s\033[0m"

//ErrorColor red
const ErrorColor = "\033[1;31m%s\033[0m"

// InfoColor yellow
const InfoColor = "\033[1;33m%s\033[0m"

func getAPIResponse(level string) (apiResp APIResponse, err error) {
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
		return apiResp, err
	}

	json.NewDecoder(resp.Body).Decode(&apiResp)

	if apiResp.Code > 400 {
		fmt.Printf(ErrorColor, "API Request error: "+apiResp.Error+"\n")
		fmt.Printf(InfoColor, "Is your API key correct?\n")
		fmt.Printf(InfoColor, "https://www.wanikani.com/settings/personal_access_tokens\n")
		return apiResp, errors.New(apiResp.Error)
	}

	return apiResp, nil
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

func createSentenceCSV(resp APIResponse) {
	var d KanjiData
	var out = make([]string, 0)

	for i := 0; i < len(resp.Data); i++ {

		d = resp.Data[i].Data

		for j := 0; j < len(d.ContextSentences); j++ {
			s := strconv.Quote(d.ContextSentences[j].Ja) + "," + strconv.Quote(d.ContextSentences[j].En)

			out = append(out, s)
		}
	}

	createFile("sentences", out)
}

func createCSV(objectType string, resp APIResponse) {
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

	if len(strings.Split(*lPtr, ",")) > 5 {
		fmt.Printf(InfoColor, "WARNING: More than 4 levels not suppported.\n")
		os.Exit(0)
	}

	resp, err := getAPIResponse(*lPtr)

	if err != nil {
		os.Exit(1)
	}

	if *kPtr {
		createCSV("kanji", resp)
		fmt.Printf(SuccessColor, "✔️ Kanji CSV written\n")
	}

	if *vPtr {
		createCSV("vocabulary", resp)
		fmt.Printf(SuccessColor, "✔️ Vocabulary CSV written\n")
	}

	if *sPtr {
		createSentenceCSV(resp)
		fmt.Printf(SuccessColor, "✔️ Sentence CSV written\n")
	}
}
