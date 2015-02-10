package transmissionrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	SESSION_HEADER = "X-Transmission-Session-Id"
)

var logger *log.Logger

type Client struct {
	address   string
	sessionId string

	httpClient *http.Client
}

func New(host, port string) *Client {
	return &Client{
		address:    fmt.Sprintf("%s:%s/transmission/rpc", host, port),
		httpClient: &http.Client{},
	}
}

func NewWithCredentials(username, password, host, port string) *Client {
	return &Client{
		address:    fmt.Sprintf("%s:%s@%s:%s/transmission/rpc", username, password, host, port),
		httpClient: &http.Client{},
	}
}

func setLogging(log *log.Logger) {
	logger = log
}

func dealWithIt(err string) {
	if logger != nil {
		logger.Println(err)
	} else {
		fmt.Println(err)
	}
}

type Request struct {
	Method string      `json:"method"`
	Args   interface{} `json:"arguments"`
}

type Response struct {
	Args   json.RawMessage `json:"arguments,omitempty"`
	Result string          `json:"result"`
}

func (trans *Client) makeRequest(request Request) (Response, error) {
	jsonified, err := json.Marshal(request)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", trans.address, bytes.NewReader(jsonified))
	if err != nil {
		return Response{}, err
	}
	req.Header.Add(SESSION_HEADER, trans.sessionId)

	resp, err := trans.httpClient.Do(req)
	if err != nil {
		dealWithIt("error request")
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 409 {
		dealWithIt("bad session, redoing it with right header")
		trans.sessionId = resp.Header.Get(SESSION_HEADER)
		return trans.makeRequest(request)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dealWithIt("error read all" + string(body))
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		dealWithIt("bad unmarshal")
		return Response{}, err
	}

	return response, nil
}
