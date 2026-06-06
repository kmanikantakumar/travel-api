package intent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Parse(text string) TripIntent {
	today := time.Now().Format("2021-06-06")
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
		"http:localhost:8085/api/intent",
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

	var intent TripIntent
	json.Unmarshal([]byte(r.Response), &intent)
	return intent

}
