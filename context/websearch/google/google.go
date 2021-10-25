package google

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/betNevS/code-examples/context/websearch/userip"
)

type Results []Result

type Result struct {
	Title, URL string
}

func Search(ctx context.Context, query string) (Results, error) {
	req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", "query")
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()
	var results Results
	err = httpDO(ctx, req, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer response.Body.Close()
		var data struct {
			ResponseData struct {
				Results []struct {
					TitleNoFormatting string
					URL               string
				}
			}
		}
		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			return err
		}
		for _, res := range data.ResponseData.Results {
			results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
		}
		return nil
	})
	return results, nil
}

func httpDO(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() {
		c <- f(http.DefaultClient.Do(req))
	}()
	select {
	case <-ctx.Done():
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}
