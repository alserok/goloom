package v1

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client interface {
	NotifyStart() error
	NotifyClosure() error
}

func NewClient(addr, appPort string) Client {
	cl := &http.Client{Timeout: time.Second}

	return &client{
		addr:    addr,
		appPort: appPort,
		cl:      cl,
	}
}

type client struct {
	addr string

	appPort string

	cl *http.Client
}

func (c *client) NotifyStart() error {
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/v1/service/add", c.addr),
		nil,
	)

	q := req.URL.Query()
	q.Add("port", c.appPort)
	req.URL.RawQuery = q.Encode()

	res, err := c.cl.Do(req)
	if err != nil {
		return fmt.Errorf("failed to request server: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("server status code is not 200")
	}

	return nil
}

func (c *client) NotifyClosure() error {
	req, _ := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/v1/service/remove", c.addr),
		nil,
	)

	q := req.URL.Query()
	q.Add("port", c.appPort)
	req.URL.RawQuery = q.Encode()

	res, err := c.cl.Do(req)
	if err != nil {
		return fmt.Errorf("failed to request server: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("server status code is not 200")
	}

	return nil
}
