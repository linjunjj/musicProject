package db
var tables =[]string{
	/*
	TBL_USER 用户信息表
	ID:用户id
	USER_NAME:用户名
	AGE:年龄
	ACCOUNT:账户
	PASSWORD:密码
	*/
  `
   CREATE TABLE TBL_USER (
    ID bigint not null AUTO_INCREMENT,
    USER_NAME VARCHAR (32),
    AGE VARCHAR (32),
    ACCOUNT VARCHAR(32),
    PASSWORD VARCHAR(32),
    PRIMARY KEY (ID)
  );`,
  /*
   TBL_MUSIC音乐
   ID：音乐序号id
   SRC:音乐路径
  */

   `
  CREATE TABLE TBL_MUSIC (
    ID bigint not null AUTO_INCREMENT,
    SRC VARCHAR (32),
    PRIMARY KEY (ID)
  );`,

  /*
    TBL_RELATEMUSIC 音乐
     ID：序号
     USER_ID 用户id
    MUSIC_ID 音乐id
  */

   `CREATE TABLE TBL_RELATEMUSIC (
    ID bigint not null AUTO_INCREMENT,
    USER_ID bigint,
    MUSIC_ID bigint,
    PRIMARY KEY (ID)
  );`,

}