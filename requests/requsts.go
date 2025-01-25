package requests

import (
	"bytes"
	"io"
	"net/http"
)

type Response struct {
	URL  string
	Text []byte
	resp http.Response
}

func Get(url string) (Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	return Response{url, data, *resp}, nil
}

func Post(url, contentType string, bytesBody []byte) (Response, error) {
	body := bytes.NewReader(bytesBody)

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return Response{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	return Response{url, data, *resp}, nil
}

func Put(url string, bytesBody []byte) (Response, error) {
	body := bytes.NewReader(bytesBody)

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return Response{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	return Response{url, data, *resp}, nil
}

func Delete(url string, bytesBody []byte) (Response, error) {
	body := bytes.NewReader(bytesBody)

	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return Response{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	return Response{url, data, *resp}, nil
}

func Head(url string) (Response, error) {
	resp, err := http.Head(url)
	if err != nil {
		return Response{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	return Response{url, data, *resp}, nil
}

func Options(url string, bytesBody []byte) (Response, error) {
	body := bytes.NewReader(bytesBody)

	req, err := http.NewRequest(http.MethodOptions, url, body)
	if err != nil {
		return Response{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Response{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	return Response{url, data, *resp}, nil
}
