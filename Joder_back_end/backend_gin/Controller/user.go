package controller

import (
	"fmt"
	"joder/beans"
	"joder/commonapi"
	"joder/dao"
	"joder/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context) {
	insertUserDto := beans.InsertUserDto{}
	c.BindJSON(&insertUserDto)
	//log.Printf("%v", &insertUserDto)
	googleId, err := c.Cookie("googleId")
	if err != nil {
		log.Printf("%v", err)
		googleId = ""
	}
	facebookId, err := c.Cookie("facebookId")
	if err != nil {
		log.Printf("%v", err)

		facebookId = ""

	}
	if facebookId == "" && googleId == "" {
		commonapi.SysError(c, beans.ResponseWrapper{}, "新增會員失敗，請確認FB/GOOGLE帳號登入正確")
		return
	}

	format := "2006-01-02"
	var userId string
	if googleId != "" {
		userId = googleId
	} else {
		userId = facebookId
	}

	birthday, _ := time.Parse(format, insertUserDto.Birthday)
	user := model.PersonalInfo{
		UserId:        userId,
		GoogleId:      googleId,
		FacebookId:    facebookId,
		Name:          insertUserDto.Name,
		Email:         insertUserDto.Email,
		Birthday:      birthday,
		UserType:      insertUserDto.Usertype,
		Gender:        insertUserDto.Gender,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		LastLoginTime: time.Now(),
	}

	result := dao.MysqlDb.Create(user)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		commonapi.SysError(c, beans.ResponseWrapper{}, "新增會員失敗，此帳號已註冊")
		return
	}

	selfIntro := model.UserIntro{
		UserId:    user.UserId,
		SelfIntro: insertUserDto.SelfIntro,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}
	dao.MysqlDb.Create(selfIntro)

	if insertUserDto.Latitude != 0 && insertUserDto.Longitude != 0 {
		userLoc := model.UserLocation{
			UserId:    user.UserId,
			Longitude: insertUserDto.Longitude,
			Latitude:  insertUserDto.Latitude,
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
		}
		dao.MysqlDb.Create(userLoc)
	}

	commonapi.Success(c, user, "新增會員成功，請重新登入")
}

func QueryUser(c *gin.Context, userId string) (model.PersonalInfo, int64) {

	personalInfo := model.PersonalInfo{}
	rows := dao.MysqlDb.Where("user_id = ?", userId).Find(&personalInfo).Select("email", "gender", "name", "user_type", "last_login_time").RowsAffected
	return personalInfo, rows

}

func QueryUserLoc(c *gin.Context, userId string) (model.UserLocation, int64) {
	userLoc := model.UserLocation{}
	rows := dao.MysqlDb.Where("user_id = ?", userId).Find(&userLoc).Select("longitude", "latitude").RowsAffected

	return userLoc, rows
}

func QueryUserIntro(c *gin.Context, userId string) model.UserIntro {
	userIntro := model.UserIntro{}
	dao.MysqlDb.Where("user_id = ?", userId).Find(&userIntro).Select("self_intro")
	return userIntro
}

func QueryUserPic(c *gin.Context, userId string) []model.UserPhoto {
	userPhoto := []model.UserPhoto{}
	dao.MysqlDb.Where("user_id = ?", userId).Find(&userPhoto).Select("photo_url")
	return userPhoto
}

func GetUserInfoWithProcedure(userId string) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	dao.MysqlDb.Raw("call sp_get_user_info_with_id(?)", userId).Scan(&ret)
	fmt.Printf("%+v /n", ret)
	return ret, nil
}

func GetUnmatchedUser(c *gin.Context) {
	userId, ok := c.Get("userId")
	unmatchedUserDto := beans.UnmatchedUserDto{}
	c.BindJSON(&unmatchedUserDto)
	//fmt.Println(userId)
	if !ok || userId == "" {
		commonapi.SysError(c, beans.ResponseWrapper{}, "No valid cookie")
	}
	matches := []model.PersonalInfo{}
	dao.MysqlDb.Raw("call sp_query_not_matched(?,?)", userId, unmatchedUserDto.PaginateDto.PaginateNum).Scan(&matches)

	commonapi.Success(c, matches, "")
}

func GetUserData(c *gin.Context) {
	userId, ok := c.Get("userId")
	fmt.Printf("%v \n", userId)
	if !ok || userId == "" {
		commonapi.SysError(c, beans.ResponseWrapper{}, "No valid cookie")
	}
	userInfo, rows := QueryUser(c, userId.(string))
	if rows == 0 {
		commonapi.SysError(c, beans.ResponseWrapper{}, "No valid user")
	}
	selfIntro := QueryUserIntro(c, userId.(string))
	userPhoto := QueryUserPic(c, userId.(string))
	var urls []string
	for _, obj := range userPhoto {
		urls = append(urls, obj.PhotoUrl)
	}
	userInfoRes := beans.UserInfo{
		Name:      userInfo.Name,
		UserId:    userInfo.UserId,
		Email:     userInfo.Email,
		PicUrl:    urls,
		SelfIntro: selfIntro.SelfIntro,
		UserType:  userInfo.UserType,
	}

	//queryUser(UserId);
	commonapi.Success(c, userInfoRes, "")
}
