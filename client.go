package tauth

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type Status int

var (
	SUCCESS               Status = 100
	INVALID_REQUEST       Status = 101
	INTERNAL_SERVER_ERROR Status = 102
	AUTHORIZATION_FAILED  Status = 103
	ALREADY_EXISTS        Status = 104
	NOT_FOUND             Status = 105
)

var client fasthttp.HostClient

func init() {
	client = fasthttp.HostClient{
		Addr:                "127.0.0.1:8001",
		MaxIdleConnDuration: time.Minute,
		ReadTimeout:         30 * time.Second,
		WriteTimeout:        30 * time.Second,
	}
}

func get(function string) (Status, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://127.0.0.1:8001/auth/" + function)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		return 0, err
	}
	response := resp.Body()
	ReleaseResponse(resp)
	i, err := strconv.Atoi(string(response))
	if err != nil {
		return 0, err
	}
	return Status(i), nil
}
func post(function string, json string) (Status, error) {
	req := fasthttp.AcquireRequest()
	req.SetBody([]byte(json))
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetRequestURI("http://127.0.0.1:8001/auth/" + function)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	if err != nil {
		return 0, err
	}
	response := resp.Body()
	ReleaseResponse(resp)
	i, err := strconv.Atoi(string(response))
	if err != nil {
		return 0, err
	}
	return Status(i), nil
}
func put(function string, json string) (Status, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://127.0.0.1:8001/auth/" + function)
	req.Header.SetContentType("application/json")
	req.Header.SetMethodBytes([]byte("PUT"))
	req.SetBody([]byte(json))
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		return 0, err
	}
	response := resp.Body()
	ReleaseResponse(resp)
	i, err := strconv.Atoi(string(response))
	if err != nil {
		return 0, err
	}
	return Status(i), nil
}
func delete(function string, json string) (Status, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://127.0.0.1:8001/auth/" + function)
	req.Header.SetContentType("application/json")
	req.Header.SetMethodBytes([]byte("DELETE"))
	req.SetBody([]byte(json))
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	if err != nil {
		return 0, err
	}
	response := resp.Body()
	ReleaseResponse(resp)
	i, err := strconv.Atoi(string(response))
	if err != nil {
		return 0, err
	}
	return Status(i), nil
}
func ReleaseResponse(response *fasthttp.Response) {
	fasthttp.ReleaseResponse(response)
}

func SignIn(username string, password string) (Status, error) {
	return put("signIn", fmt.Sprintf(`{"username":"%s", "password":"%s"}`, username, password))
}
func CheckPassword(username string, password string) (Status, error) {
	return get(fmt.Sprintf("checkPassword?u=%s&p=%s", username, password))
}
func ResetPassword(username string, oldPassword string, newPassword string) (Status, error) {
	return put("resetPassword", fmt.Sprintf(`{"username":"%s", "oldPassword":"%s", "newPassword":"%s"}`, username, oldPassword, newPassword))
}
func RequestSignUpCode(username string, email string) (Status, error) {
	return post("requestSignUpCode", fmt.Sprintf(`{"username":"%s", "email":"%s"}`, username, email))
}
func RequestPasswordRecovery(username string) (Status, error) {
	return put("requestPasswordRecovery", fmt.Sprintf(`{"username":"%s"}`, username))
}
func RecoverPassword(username string, newPassword string, code string) (Status, error) {
	return put("recoverPassword", fmt.Sprintf(`{"username":"%s", "newPassword":"%s", "code":"%s"}`, username, newPassword, code))
}
func SignUp(username string, password string, code string) (Status, error) {
	return post("signUp", fmt.Sprintf(`{"username":"%s", "password":"%s", "code":"%s"}`, username, password, code))
}
