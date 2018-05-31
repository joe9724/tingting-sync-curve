package main

import (
	"time"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"tingting-sync/var"
)

func main() {
	// RunTimer()
	GetDashboardData()
}
func RunTimer() {
	GetDashboardTimer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-GetDashboardTimer.C:
			GetDashboardData()
		}
	}
}
func GetDashboardData() {
	//go func() {
		fmt.Println("timer...")
         //今日新增用户
		db,err := utils.OpenConnection()
		if err!=nil{
			fmt.Println(err.Error())
		}
		defer db.Close()

		var number_newuser int64
		db.Raw("select count(*) from members where to_days(ts) = to_days(now())").Count(&number_newuser)
        fmt.Println("number_newuser is",number_newuser)

         //今日购买专辑
         var number_today_buy_albums int64
         db.Raw("select count(*) from orders  where to_days(createtime) = to_days(now())").Count(&number_today_buy_albums)

		 fmt.Println("number_today_buy_albums is",number_today_buy_albums)
         //今日成交额
         var money_today int64
         db.Raw("select sum(value) from orders where to_days(createtime) = to_days(now())").Count(&money_today)
		 fmt.Println("money_today is",money_today)
		 db.Exec("insert into everyday_deal_album(album_count,money) values (?,?)",number_today_buy_albums,money_today )
         //今日上传录音
         /*var number_today_record int64
         db.Raw("select count(*) from records where to_days(createtime) = to_days(now())").Count(&number_today_record)
		 fmt.Println("number_today_record is",number_today_record)
         //月内购买专辑数
         var number_month_buy_albums int64
         db.Raw("select count(*) from orders where status=0 and DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(createtime)").Count(&number_month_buy_albums)
		 fmt.Println("number_month_buy_albums is",number_month_buy_albums)
         //月内成交额
         var money_month_buy_albums int64
         db.Raw("select sum(value) from orders where status=0 and DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(createtime)").Count(&money_month_buy_albums)
		 fmt.Println("money_month_buy_albums is",money_month_buy_albums)
         //近期热门


         //今日新增分类数
         var number_today_add_category int64
         db.Raw("select count(*) from sub_category_items where to_days(time) = to_days(now())").Count(&number_today_add_category)
		 fmt.Println("number_today_add_category is",number_today_add_category)
         //今日新增专辑数
         var number_today_add_albums int64
         db.Raw("select count(*) from albums where to_days(date_format(from_UNIXTIME(`time`),'%Y-%m-%d')) = to_days(now())").Count(&number_today_add_albums)
		 fmt.Println("number_today_add_albums is",number_today_add_albums)
         //今日新增书本数
         var number_today_add_book int64
         db.Raw("select count(*) from books where to_days(date_format(from_UNIXTIME(`time`),'%Y-%m-%d')) = to_days(now())").Count(&number_today_add_book)
		 fmt.Println("number_today_add_book is",number_today_add_book)
         //今日新增章节数
         var number_today_add_chapter int64
         db.Raw("select count(*) from chapters where to_days(date_format(from_UNIXTIME(`time`),'%Y-%m-%d')) = to_days(now())").Count(&number_today_add_chapter)
		 fmt.Println("number_today_add_chapter is",number_today_add_chapter)*/
         //
         // db.Exec("update dashboard set number_newuser=?,number_today_buy_albums=?,money_today=?,number_today_record=?,number_month_buy_albums=?,number_today_add_category=?,number_today_add_albums=?,number_today_add_book=?,number_today_add_chapter=?", number_newuser,number_today_buy_albums,money_today,number_today_record,number_month_buy_albums,number_today_add_category,number_today_add_albums,number_today_add_book,number_today_add_chapter)

	//}()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}