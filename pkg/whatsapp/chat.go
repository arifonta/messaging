package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Chat struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type ApiResponse struct {
	Message string `json:"message"`
	Data    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

func CallAPIChat(req Chat) (err error) {
	payload := []byte(`{
		"to": req.To,
		"recipientType": "individual",
		"type": "template",
		"template": {
		  "namespace": "5c922676_64f2_42ec_a56d_965a23a6df9b",
		  "name": "sedna_dev_sent_file_patient",
		  "language": {
			"code": "id",
			"policy": "deterministic"
		  },
		  "components": [
			{
			  "type": "header",
			  "parameters": [
				{
				  "type": "document",
				  "document": {
					"link": "https://storebox.jec.co.id/v1/?file=/storage_umum/doc/Swagger UI.pdf",
					"filename": "Sample Data Source"
				  }
				}
			  ]
			},
			{
			  "type": "body",
			  "parameters": [
				{
				  "type": "text",
				  "text": "Sample Data Source"
				},
				{
				  "type": "text",
				  "text": "Tn. Patient Demo (001-001-00-01)"
				},
				{
				  "type": "text",
				  "text": "Last Visit: JEC@DEVELOPMENT, 02 Mar 2023"
				}
			  ]
			}
		  ]
		}
	  }`)

	// Make a POST request
	url := "https://api-whatsapp.kata.ai/v1/messages"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the response body
	var apiResponse ApiResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&apiResponse)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// Print the parsed data
	fmt.Println("Message:", apiResponse.Message)
	fmt.Println("ID:", apiResponse.Data.ID)
	fmt.Println("Name:", apiResponse.Data.Name)

	return
}
