package requests

import (
	"bytes"
	"io"
	"net/http"
)

type Response struct {
	URL  string
	Text []byte
}

func Get(url string) Response {
	resp, err := http.Get(url)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	err = resp.Body.Close()
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	return Response{url, data}
}

func Post(url string, bytesBody []byte) Response {
	body := bytes.NewReader(bytesBody)

	resp, err := http.Post(url, "application/octet-stream", body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	err = resp.Body.Close()
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	return Response{url, data}
}

func Put(url string, bytesBody []byte) Response {
	body := bytes.NewReader(bytesBody)

	resp, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	err = resp.Body.Close()
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	return Response{url, data}
}

func Delete(url string, bytesBody []byte) Response {
	body := bytes.NewReader(bytesBody)

	resp, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	err = resp.Body.Close()
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	return Response{url, data}
}

func Head(url string) Response {
	resp, err := http.Head(url)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	err = resp.Body.Close()
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	return Response{url, data}
}

func Options(url string, bytesBody []byte) Response {
	body := bytes.NewReader(bytesBody)

	resp, err := http.NewRequest(http.MethodOptions, url, body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	err = resp.Body.Close()
	if err != nil {
		return Response{url, []byte(err.Error())}
	}

	return Response{url, data}
}
