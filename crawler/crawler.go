package crawler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type Crawler struct {
	URL     string
	Timeout time.Duration
}

func (c *Crawler) Cities(prefix string) ([]string, error) {
	client := &http.Client{Timeout: c.Timeout}

	req, err := http.NewRequest(http.MethodGet, c.URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("error occurred with closing response body", slog.String("error", err.Error()))
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong satus code: %d", resp.StatusCode)
	}

	var schema Schema
	if err := json.NewDecoder(resp.Body).Decode(&schema); err != nil {
		return nil, err
	}

	var res []string
	for _, g := range schema.Geonames {
		if strings.HasPrefix(g.Name, prefix) {
			res = append(res, g.Name)
		}
	}

	return res, nil
}
