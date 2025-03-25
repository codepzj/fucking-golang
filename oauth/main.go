package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	ClientId     string
	ClientSecret string
	httpClient   = &http.Client{}
)

type AccessTokenReq struct {
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Code         string `json:"code,omitempty"`
}

type UserInfo struct {
	Login     string `json:"login,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}
	ClientId = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{"client_id": ClientId})
	})
	r.GET("/oauth/redirect", Login)
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func Login(ctx *gin.Context) {
	code := ctx.Query("code")
	accessToken, err := GetAccessToken(code)
	if err != nil {
		log.Printf("Failed to get access token: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get access token"})
		return
	}

	info, err := GetUserInfo(accessToken)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"name":       info.Login,
		"avatar_url": info.AvatarUrl,
	})
}

func GetAccessToken(code string) (string, error) {
	accessTokenReq := AccessTokenReq{
		ClientId:     ClientId,
		ClientSecret: ClientSecret,
		Code:         code,
	}
	jsonData, err := json.Marshal(accessTokenReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal access token request: %w", err)
	}

	resp, err := httpClient.Post("https://github.com/login/oauth/access_token", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to request access token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned non-200 status: %d", resp.StatusCode)
	}

	respJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	val, err := url.ParseQuery(string(respJson))
	if err != nil {
		return "", fmt.Errorf("failed to parse access token response: %w", err)
	}

	return val.Get("access_token"), nil
}

func GetUserInfo(accessToken string) (UserInfo, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to create user info request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to send user info request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("GitHub API returned non-200 status: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, fmt.Errorf("failed to read user info response: %w", err)
	}

	var userInfo UserInfo
	if err := json.Unmarshal(respBody, &userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	return userInfo, nil
}
