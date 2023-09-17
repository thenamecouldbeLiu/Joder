package controller

import (
	"fmt"
	"io"
	"joder/beans"
	"joder/commonapi"
	"joder/config"
	"joder/jwt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

const redirectURL = "http://localhost:9090/api/ouath/google/login"
const scope = "https://www.googleapis.com/auth/userinfo.profile+https://www.googleapis.com/auth/userinfo.email"

func oauthUrl() string {
	link := "https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&response_type=code&scope=%s&redirect_uri=%s"

	return fmt.Sprintf(link, config.Val.GoogleClientId, scope, redirectURL)
}

func GetOauthUrl(c *gin.Context) {

	url := oauthUrl()

	c.String(200, url)
}

func accessToken(code string) (token string, err error) {
	link := "https://www.googleapis.com/oauth2/v4/token"
	data := url.Values{"code": {code}, "client_id": {config.Val.GoogleClientId},
		"client_secret": {config.Val.GoogleSecretKey},
		"grant_type":    {"authorization_code"}, "redirect_uri": {redirectURL}}

	body := strings.NewReader(data.Encode())

	res, err := http.Post(link, "application/x-www-form-urlencoded", body)
	if err != nil {
		return token, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return token, err
	}

	token = gjson.GetBytes(resBody, "access_token").String()

	return token, nil
}

func getGoogleUserInfo(token string) (id string, name string, email string, err error) {
	link := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", token)

	res, err := http.Get(link)

	if err != nil {
		return id, name, email, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return id, name, email, err
	}
	email = gjson.GetBytes(body, "email").String()
	name = gjson.GetBytes(body, "name").String()
	id = gjson.GetBytes(body, "id").String()
	return id, name, email, nil
}

func GoogleAccess(c *gin.Context) {
	commonapi.Success(c, gin.H{
		"url": oauthUrl(),
	})
}

func GoogleLogin(c *gin.Context) {
	code := c.Query("code")
	token, err := accessToken(code)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Info("token err")
		c.Redirect(http.StatusFound, "/")
		return
	}

	id, name, email, err := getGoogleUserInfo(token)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Info("user info err")
		c.Redirect(http.StatusFound, "/")
		return
	}
	fmt.Printf("id: %v, name: %v, email: %v \n", id, name, email)

	jwtToken, err := jwt.GenerateJWT(id)

	if err != nil {
		fmt.Printf("err happend %v\n", err)
		log.WithFields(log.Fields{
			"err": err,
		}).Info("GenerateToken error \n")
		return
	}

	c.SetCookie(config.Val.COOKIE_KEY, jwtToken, int(config.Val.JWT_TOKEN_LIFE), "/api/ouath/google/oauthUrl", "localhost", false, true)

	model := make(map[string]any)
	res := beans.ResponseWrapper{ResultCode: "0000", ResultMessage: "", Model: model}
	c.JSON(200, res)

	c.Redirect(http.StatusFound, "http://localhost:5173/")

}
