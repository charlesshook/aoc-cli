package helpers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetDataFromUrlWithCookie(url string, cookie string) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("setting up request: %s", err)
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: cookie,
	}

	request.AddCookie(&sessionCookie)

	respose, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("could not get response: %s", err)
	}

	data, err = io.ReadAll(respose.Body)

	if err != nil {
		return nil, fmt.Errorf("could not read response body: %s", err)
	}

	return data, nil
}
