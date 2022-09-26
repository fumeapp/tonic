package search

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/setting"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type SearchQuery struct {
	Index    string
	Sorting  string
	Order    string
	Limiting int
	Offset   int
}

type SearchResult struct {
	SearchQuery *SearchQuery
	Took        int  `json:"took"`
	TimedOut    bool `json:"timed_out"`
	Shards      struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string      `json:"_index"`
			ID     string      `json:"_id"`
			Score  float64     `json:"_score"`
			Source interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

/**
	TODO: add terms
  "sort" : [
    { "post_date" : {"order" : "asc", "format": "strict_date_optional_time_nanos"}},
    "user",
    { "name" : "desc" },
    { "age" : "desc" },
    "_score"
  ],
  "query" : {
    "term" : { "user" : "kimchy" }
  }
*/

func Search(index string) *SearchQuery {
	return &SearchQuery{
		Index:    index,
		Order:    "asc",
		Limiting: 10000,
		Offset:   0,
	}
}

func (sq *SearchQuery) Sort(sort string) *SearchQuery {
	sq.Sorting = sort
	return sq
}

func (sq *SearchQuery) Asc() *SearchQuery {
	sq.Order = "asc"
	return sq
}

func (sq *SearchQuery) Desc() *SearchQuery {
	sq.Order = "desc"
	return sq
}

func (sq *SearchQuery) Limit(limit int) *SearchQuery {
	sq.Limiting = limit
	return sq
}

func (sq *SearchQuery) toString() string {

	query := "{"

	query += fmt.Sprintf(`
		"sort": [
			{
				"%s": {"order": "%s"}
			}
		],
		"size": %d
		`, sq.Sorting, sq.Order, sq.Limiting)

	query += "}"
	return query
}

func (sq *SearchQuery) Get() (*SearchResult, error) {
	req := opensearchapi.SearchRequest{
		Index: []string{sq.Index},
		Body:  strings.NewReader(sq.toString()),
	}

	res, err := req.Do(context.Background(), database.Os)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result SearchResult
	if err := json.Unmarshal(bs, &result); err != nil {
		return nil, err
	}
	if setting.IsDebug() {
		result.SearchQuery = sq
	}

	return &result, nil
}
