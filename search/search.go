package search

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fumeapp/tonic/database"
	"github.com/octoper/go-ray"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func DeleteIndex(index string) (*opensearchapi.Response, error) {

	req := opensearchapi.IndicesDeleteRequest{
		Index: []string{index},
	}

	res, err := req.Do(context.Background(), database.Os)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func MapIndex(index string, stringses []string, dates []string) error {
	mappings := `{
		"mappings": {
			"properties": {
`
	for _, str := range stringses {
		mappings += fmt.Sprintf(
			`"%s": {
					"type": "text",
					"fields": {
						"keyword": {
							"type": "keyword",
							"ignore_above": 256
						}
					}
				},
`, str)
	}
	for i, str := range dates {
		mappings += fmt.Sprintf(
			`"%s": {
					"type": "date"
				}`, str)
		if (i + 1) < len(dates) {
			mappings += ","
		}
	}

	mappings += `
		}
	}
}`
	ray.Ray(mappings)
	req := opensearchapi.IndicesCreateRequest{
		Index:  index,
		Body:   strings.NewReader(mappings),
		Pretty: true,
	}

	if res, err := req.Do(context.Background(), database.Os); err != nil {
		return err
	} else {
		defer res.Body.Close()
		if res.IsError() {
			return fmt.Errorf("Error: %s", res.String())
		}
	}
	return nil
}

func CreateDocument(index string, Id string, body interface{}) (*opensearchapi.Response, error) {

	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	document := strings.NewReader(string(buf))
	req := opensearchapi.IndexRequest{
		DocumentID: Id,
		Index:      index,
		Body:       document,
	}

	res, err := req.Do(context.Background(), database.Os)
	if err != nil {
		return nil, err
	}

	return res, nil
}
