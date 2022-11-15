package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"

	"github.com/go-resty/resty/v2"
)

type RestClient interface {
	Get(url string, result interface{}) error
	Post(url string, data interface{}, result interface{}) error
	Put(url string, data interface{}, result interface{}) error
	Delete(url string, data interface{}, result interface{}) error
	Patch(url string, data interface{}, result interface{}) error
	Upload(url string, files map[string]string, formData map[string]string, result interface{}) error
	UploadBytes(url string, files []FileParam, formData map[string]string, result interface{}) error
	UploadBytesAndDownload(url string, files []FileParam, formData map[string]string, result *[]byte) error
	PostDownload(url string, data interface{}, result *[]byte) error
	GetDownload(url string, result *[]byte) error
}

type FileParam struct {
	Name     string
	FileName string
	Reader   io.Reader
}

type ClientOption struct {
	Debug    bool
	Insecure bool
	Timeout  time.Duration
}

// restClient @impl client.RestClient
type restClient struct {
	Client *resty.Client
}

func NewRestClient(option ClientOption) RestClient {
	client := resty.New()
	if option.Debug {
		client.SetDebug(true)
	}
	// or One can disable security check (https)
	if option.Insecure {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	if option.Timeout != 0 {
		client.SetTimeout(option.Timeout)
	}

	// Registering Request Middleware
	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		fmt.Println("############## REQ : REST CLIENT ##################")
		fmt.Println("Request URL : " + req.URL)
		fmt.Println("############## REQ : REST CLIENT ##################")
		// Now you have access to Client and current Request object
		// manipulate it as per your need

		return nil // if its success otherwise return error
	})

	// Registering Response Middleware
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		fmt.Println("############## RES : REST CLIENT ##################")
		fmt.Println(resp.StatusCode())
		if resp.StatusCode() != 200 {
			fmt.Println(resp.IsError())
			fmt.Println(string(resp.Body()))
		}
		fmt.Println("############## RES : REST CLIENT ##################")
		return nil // if its success otherwise return error
	})

	return &restClient{Client: client}
}

func (r *restClient) Get(url string, result interface{}) error {
	_, err := r.Client.R().
		SetResult(result).
		Get(url)
	if err != nil {
		return err
	}
	return nil
}

func (r *restClient) Post(url string, data interface{}, result interface{}) error {
	_, err := r.Client.R().
		SetBody(data).
		SetResult(result).
		Post(url)
	if err != nil {
		return err
	}
	return nil
}

func (r *restClient) Put(url string, data interface{}, result interface{}) error {
	_, err := r.Client.R().
		SetBody(data).
		SetResult(result).
		Put(url)
	if err != nil {
		return err
	}
	return nil
}

func (r *restClient) Delete(url string, data interface{}, result interface{}) error {
	_, err := r.Client.R().
		SetBody(data).
		SetResult(result).
		Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (r *restClient) Patch(url string, data interface{}, result interface{}) error {
	_, err := r.Client.R().
		SetBody(data).
		SetResult(result).
		Patch(url)
	if err != nil {
		return err
	}
	return nil
}

func (r *restClient) Upload(url string, files map[string]string, formData map[string]string, result interface{}) error {
	_, err := r.Client.R().
		SetResult(result).
		SetFiles(files).
		SetFormData(formData).
		Post(url)
	if err != nil {
		return err
	}
	return nil
}

func (r *restClient) UploadBytes(url string, files []FileParam, formData map[string]string, result interface{}) error {
	request := r.Client.R()
	for _, file := range files {
		request = request.SetFileReader(file.Name, file.FileName, file.Reader)
	}
	_, err := request.
		SetResult(result).
		SetFormData(formData).
		Post(url)
	if err != nil {
		return err
	}
	return nil
}

func (r *restClient) UploadBytesAndDownload(url string, files []FileParam, formData map[string]string, result *[]byte) error {
	request := r.Client.R()
	for _, file := range files {
		request = request.SetFileReader(file.Name, file.FileName, file.Reader)
	}
	resp, err := request.
		SetResult(result).
		SetFormData(formData).
		Post(url)
	if err != nil {
		return err
	}
	fmt.Println("  Success Code:", resp.StatusCode())
	if resp.StatusCode() != 200 {
		fmt.Println("  Body       :\n", resp)
		return errors.New("Call service error")
	}
	*result = resp.Body()
	return nil
}

func (r *restClient) PostDownload(url string, data interface{}, result *[]byte) error {
	resp, err := r.Client.R().
		SetBody(data).
		Post(url)
	if err != nil {
		return err
	}
	*result = resp.Body()
	return nil
}

func (r *restClient) GetDownload(url string, result *[]byte) error {
	resp, err := r.Client.R().
		Get(url)
	if err != nil {
		return err
	}
	*result = resp.Body()
	return nil
}
