package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

/*
处理用户模块的crud
*/

const secret = "zhyfgzm"

// CheckUserExist  检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条心的用户记录
func InsertUser(user *models.User) (err error) {
	//加密密码
	password := encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user (user_id, username, password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	return
}

// encryptPassword 加密密码
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Login GetUserByUsernameAndPassword 通过username 获取用户 判断密码是不是相等
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id , username , password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	//返回的数据为空ErrorUserNotExist
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		//密码不匹配 ErrorInvalidPassword
		return ErrorInvalidPassword
	}
	return nil
}
// GetUserById
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id , username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return user, err
}
