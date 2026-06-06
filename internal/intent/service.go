package intent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Parse(text string) TripIntent {
	today := time.Now().Format("2006-01-02")

	prompt := fmt.Sprintf(

		`Return ONLY this JSON (no other text): 

        {"destination":"City, Country","dateFrom":"YYYY-MM-DD", 

         "dateTo":"YYYY-MM-DD","purpose":"Meeting type"} 


        Today is %s. Resolve relative dates like 'next Thursday'. 

        Trip description: "%s"`, today, text)

	body, _ := json.Marshal(map[string]interface{}{
		"model":  "llama3.1:8b",
		"prompt": prompt,
		"stream": false,
	})

	resp, err := http.Post(
		"http://localhost:8080/api/intent",
		"application/json",
		bytes.NewReader(body),
	)

	if err != nil {

		return TripIntent{Destination: "unknown"}
	}

	defer resp.Body.Close()

	var r struct {
		Response string `json:"response"`
	}
	json.NewDecoder(resp.Body).Decode(&r)

	dummyResponse := `{
        "destination":"London, UK",
        "dateFrom":"2026-06-11",
        "dateTo":"2026-06-14",
        "purpose":"Sales Meeting"
    }`

	var intent TripIntent
	json.Unmarshal([]byte(dummyResponse), &intent)
	return intent

}
