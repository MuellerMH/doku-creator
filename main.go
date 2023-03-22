package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
    fmt.Println("Hello, World!")
    // gehe durch alle Datein im Ordner 'dir' mit der Endung 'extension'
    // Step 1: Receive the parameter values from the CLI.
		args := os.Args[1:]
		extension := args[0]
		repoURL := args[1]

		// Step 2: Print the parameter values to the console.
		fmt.Println("Extension:", extension)
		fmt.Println("Repo URL:", repoURL)
		splitURL := strings.Split(repoURL, "/")
		dir := strings.Replace(splitURL[4],".git","",-1)
		fmt.Println("Directory: ", dir)
		// wenn order nicht existiert
		if _, err := os.Stat(dir); os.IsNotExist(err) {
      gitClone(repoURL)
    }

    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			fmt.Println(path,info)
        if err != nil {
            return err
        }
        if filepath.Ext(path) == "."+extension {
            fmt.Println(path)
            // Lese den Inhalt der Datei und sende in an die Open.AI API um eine Dokumentation zum code zu erhalten. Der API Token steht in den ENV als OPEN_AI_API_TOKE
            summary := getDokumentation(path)
						fmt.Println(summary)
						// summary string to Struct
						task := getTask(summary)
						// replace .go to .md in path
						mdFile := strings.Replace(path, "."+extension, ".md", -1)
						mdFile = strings.Replace(mdFile, dir, dir+"/doku", -1)
						// make dir ist missing
						os.MkdirAll(filepath.Dir(mdFile), os.ModePerm)
						// write summary to file
						for _,choise := range task.Choices {
							err = ioutil.WriteFile(mdFile, []byte(choise.Text), 0644)
							if err!= nil {
									return err
							}
						}
        }
				return nil
		})

    if err != nil {
        fmt.Println(err)
    }

}

type Task struct {
	Choices []Choise `json:"choices"`
	ID   int    `json:"id"`
	Usage Usage `json:"usage"`
	Object string `json:"object"`
	CreatedAt string `json:"created"`
	Model string `json:"model"`
}
type Choise struct {
	Text string `json:"text"`
	Index int `json:"index"`
	Logprops string `json:"logprobs"`
	Finished bool `json:"finish_reason"`
}
type Usage struct {
	PromtTokens int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens int `json:"total_tokens"`
}

func getTask(content string) Task {
	task := Task{}
	json.Unmarshal([]byte(content), &task)
	return task
}

func gitClone(repoURL string) error {
	cmd := exec.Command("git", "clone", repoURL)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func getDokumentation(filename string) string {
	// Read the content of the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Setup the API request
	return askGPT(string(content))
}

type RequestPayload struct {
	Model string `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
	MaxTokens   int     `json:"max_tokens"`
}


func askGPT(content string) string {
	apiKey := os.Getenv("OPEN_AI_API_TOKEN")
	payload := RequestPayload{
		Model: "text-davinci-003",
		Prompt:      fmt.Sprintf("Analysiere den folgenden Code und fasse es für die Dokumentation zusammen und formatiere den Text als Markdown: %s", content),
		Temperature: 0.7,
		MaxTokens:   2048,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	url := "https://api.openai.com/v1/completions"

	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Read the API response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Output the response
	return string(body)
}