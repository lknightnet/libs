package genToken

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func convert(chtoto any) bytes.Buffer {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(chtoto)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func (st *ServiceToken) GenerateToken(clientUUID string, timeout time.Duration) (*TokenResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}
	generateToken := &GenerateTokenRequest{
		ClientUUID: clientUUID,
	}
	buf := convert(generateToken)
	request, err := http.NewRequest(http.MethodPost, st.API+"api/tokens", &buf)
	if err != nil {
		return nil, NewError("libs/genToken/GenerateToken: failed to create request", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, NewError("libs/genToken/GenerateToken: failed to create response", err)
	}
	var tokenResponse TokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	if err != nil {
		log.Println(NewError("libs/genToken/GenerateToken: failed to decode response", err))
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(NewError("libs/genToken/GenerateToken: failed to decode response to string", err))
		}
		fmt.Println(string(data))
	}
	return &tokenResponse, nil
}

func (st *ServiceToken) GetAccessToken(accessToken string, timeout time.Duration) (*AccessTokenResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}
	generateToken := &GetAccessTokenRequest{
		AccessToken: accessToken,
	}
	buf := convert(generateToken)
	request, err := http.NewRequest(http.MethodPost, st.API+"api/get/access", &buf)
	if err != nil {
		return nil, NewError("libs/genToken/GetAccessToken: failed to create request", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, NewError("libs/genToken/GetAccessToken: failed to create response", err)
	}
	var tokenResponse AccessTokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	if err != nil {
		log.Println(NewError("libs/genToken/GetAccessToken: failed to decode response", err))
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(NewError("libs/genToken/GetAccessToken: failed to decode response to string", err))
		}
		fmt.Println(string(data))
	}
	return &tokenResponse, nil
}

func (st *ServiceToken) GetRefreshToken(refreshToken string, timeout time.Duration) (*RefreshTokenResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}
	generateToken := &GetRefreshTokenRequest{
		RefreshToken: refreshToken,
	}
	buf := convert(generateToken)
	request, err := http.NewRequest(http.MethodPost, st.API+"api/get/refresh", &buf)
	if err != nil {
		return nil, NewError("libs/genToken/GetRefreshToken: failed to create request", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, NewError("libs/genToken/GetRefreshToken: failed to create response", err)
	}
	var tokenResponse RefreshTokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	if err != nil {
		log.Println(NewError("libs/genToken/GetRefreshToken: failed to decode response", err))
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(NewError("libs/genToken/GetRefreshToken: failed to decode response to string", err))
		}
		fmt.Println(string(data))
	}
	return &tokenResponse, nil
}

func (st *ServiceToken) Refresh(refreshToken string, timeout time.Duration) (*TokenResponse, error) {
	client := http.Client{
		Timeout: timeout,
	}
	generateToken := &RefreshRequest{
		RefreshToken: refreshToken,
	}
	buf := convert(generateToken)
	request, err := http.NewRequest(http.MethodPost, st.API+"api/refresh", &buf)
	if err != nil {
		return nil, NewError("libs/genToken/Refresh: failed to create request", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, NewError("libs/genToken/Refresh: failed to create response", err)
	}
	var tokenResponse TokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	if err != nil {
		log.Println(NewError("libs/genToken/Refresh: failed to decode response", err))
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(NewError("libs/genToken/Refresh: failed to decode response to string", err))
		}
		fmt.Println(string(data))
	}
	return &tokenResponse, nil
}
