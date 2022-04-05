package github

import (
	"fmt"
	transport "github.com/misteeka/fasthttp"
)

type Response []byte

var (
	SUCCESS               = Response{100}
	INVALID_REQUEST       = Response{101}
	INTERNAL_SERVER_ERROR = Response{102}
	AUTHORIZATION_FAILED  = Response{103}
	ALREADY_EXISTS        = Response{104}
	NOT_FOUND             = Response{105}
)

func get(function string) ([]byte, error) {
	resp, err := transport.Get("http://127.0.0.1:8001/auth/" + function)
	if err != nil {
		return nil, err
	}
	response := resp.Body()
	transport.ReleaseResponse(resp)
	return response, nil
}
func post(function string, json string) ([]byte, error) {
	resp, err := transport.Post("http://127.0.0.1:8001/auth/"+function, []byte(json))
	if err != nil {
		return nil, err
	}
	response := resp.Body()
	transport.ReleaseResponse(resp)
	return response, nil
}
func put(function string, json string) ([]byte, error) {
	resp, err := transport.Put("http://127.0.0.1:8001/auth/"+function, []byte(json))
	if err != nil {
		return nil, err
	}
	response := resp.Body()
	transport.ReleaseResponse(resp)
	return response, nil
}

func SignIn(username string, password string) (Response, error) {
	return put("signIn", fmt.Sprintf(`{"username":"%s", "password":"%s"}`, username, password))
}
func CheckPassword(username string, password string) (Response, error) {
	return get(fmt.Sprintf("checkPassword?u=%s&p=%s", username, password))
}
func ResetPassword(username string, oldPassword string, newPassword string) (Response, error) {
	return put("resetPassword", fmt.Sprintf(`{"username":"%s", "oldPassword":"%s", "newPassword":"%s"}`, username, oldPassword, newPassword))
}
func RequestSignUpCode(username string, email string) (Response, error) {
	return post("requestSignUpCode", fmt.Sprintf(`{"username":"%s", "email":"%s"}`, username, email))
}
func RequestPasswordRecovery(username string) (Response, error) {
	return put("requestPasswordRecovery", fmt.Sprintf(`{"username":"%s"}`, username))
}
func RecoverPassword(username string, newPassword string, code string) (Response, error) {
	return put("recoverPassword", fmt.Sprintf(`{"username":"%s", "newPassword":"%s", "code":"%s"}`, username, newPassword, code))
}
func SignUp(username string, password string, code string) (Response, error) {
	return post("signUp", fmt.Sprintf(`{"username":"%s", "password":"%s", "code":"%s"}`, username, password, code))
}
