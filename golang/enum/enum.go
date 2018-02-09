// 枚举，一组预定义好的值
// 需要遍历，获取和校验，容易扩展
// 一般用于特定范围的变量
package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	for _, v := range weekdays {
		fmt.Println(v)
	}
	fmt.Println(weekdays.One("sunday"))

	for _, v := range months {
		fmt.Println(v)
	}
	// fmt.Println(months.One("january"))

	for _, v := range hours {
		fmt.Println(v)
	}
	// fmt.Println(hours.One("1"))

	for k, v := range ColorMap {
		fmt.Println(k, ":", v)
	}
}

func timezone() {
	// GMT，UTC，DST，CST
	now := time.Now()
	fmt.Println(now.Zone()) // DST 28800
	// GMT: 格林威治标准时间 Greenwich Mean Time
	// UTC: 世界协调时间 Coordinated Universal Time
	// DST: 夏日节约时间(夏令时 Summer Time) Daylight Saving Time
	// CST: 可以同时表示美国，澳大利亚，中国，古巴四个国家的标准时间
}

// type Weekday int

// const (
// 	Sunday Weekday = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// )

// var weekdays = []Weekday{
// 	Sunday,
// 	Monday,
// 	Tuesday,
// 	Wednesday,
// 	Thursday,
// 	Friday,
// 	Saturday,
// }

// Enum 枚举
type Enum struct {
	Value string // 值
	EN    string // 英文注释
	CN    string // 中文注释
	// 如果要添加更多种语言表示呢？
}

// String 打印
func (e Enum) String() string {
	return fmt.Sprintf("value is %s, en is %s, cn is %s", e.Value, e.EN, e.CN)
}

// WeekdayCNT 周数
const WeekdayCNT = 7

// Weekday 周列表
type Weekday [WeekdayCNT]Enum

// One 取一个
func (w Weekday) One(value string) (Enum, error) {
	for _, v := range w {
		if v.Value == value {
			return v, nil
		}
	}
	return Enum{}, errors.New("Not Exist")
}

// weekdays 周列表变量
var weekdays = Weekday{
	{"sunday", "Sunday", "周日"},
	{"monday", "Monday", "周一"},
	{"tuesday", "Tuesday", "周二"},
	{"wednesday", "Wednesday", "周三"},
	{"thursday", "Thursday", "周四"},
	{"friday", "Friday", "周五"},
	{"saturday", "Saturday", "周六"},
}

// MonthCNT 月数
const MonthCNT = 12

// Month 月列表
type Month [MonthCNT]Enum

// months 月列表变量
var months = Month{
	{"january", "January", "一月"},
	{"february", "February", "二月"},
	{"march", "March", "三月"},
	{"april", "April", "四月"},
	{"may", "May", "五月"},
	{"june", "June", "六月"},
	{"july", "July", "七月"},
	{"august", "August", "八月"},
	{"september", "September", "九月"},
	{"october", "October", "十月"},
	{"november", "November", "十一月"},
	{"december", "December", "十二月"},
}

// HourCNT 时数
const HourCNT = 24

// Hour 时列表
type Hour [HourCNT]Enum

// hours 时列表变量
var hours = Hour{
	{"0", "00:00:00", "零点"},
	{"1", "01:00:00", "一点"},
	{"2", "02:00:00", "二点"},
	{"3", "03:00:00", "三点"},
	{"4", "04:00:00", "四点"},
	{"5", "05:00:00", "五点"},
	{"6", "06:00:00", "六点"},
	{"7", "07:00:00", "七点"},
	{"8", "08:00:00", "八点"},
	{"9", "09:00:00", "九点"},
	{"10", "10:00:00", "十点"},
	{"11", "11:00:00", "十一点"},
	{"12", "12:00:00", "十二点"},
	{"13", "13:00:00", "十三点"},
	{"14", "14:00:00", "十四点"},
	{"15", "15:00:00", "十五点"},
	{"16", "16:00:00", "十六点"},
	{"17", "17:00:00", "十七点"},
	{"18", "18:00:00", "十八点"},
	{"19", "19:00:00", "十九点"},
	{"20", "20:00:00", "二十点"},
	{"21", "21:00:00", "二十一点"},
	{"22", "22:00:00", "二十二点"},
	{"23", "23:00:00", "二十三点"},
}

// TimeZone 时区
type TimeZone Enum

// const ( // const initializer Enum literal is not a constant
var (
	IDL  = TimeZone{"idl", "IDL", "国际换日线"}
	MIT  = TimeZone{"mit", "MIT", "中途岛标准时间"}
	HST  = TimeZone{"hst", "HST", "夏威夷 阿留申标准时间"}
	MSIT = TimeZone{"msit", "MSIT", "马克萨斯群岛标准时间"}
	AKST = TimeZone{"akst", "AKST", "阿拉斯加标准时间"}
	PST  = TimeZone{"pst", "PST", "太平洋标准时间"}
	MST  = TimeZone{"mst", "MST", "北美山区标准时间"}
	CST  = TimeZone{"cst", "CST", "北美中部标准时间"}
	EST  = TimeZone{"est", "EST", "北美东部标准时间"}
	AST  = TimeZone{"ast", "AST", "大西洋标准时间"}
	NST  = TimeZone{"nst", "NST", "纽芬兰岛标准时间"}
	SAT  = TimeZone{"sat", "SAT", "南美标准时间"}
	BRT  = TimeZone{"brt", "BRT", "巴西时间"}
	CVT  = TimeZone{"cvt", "CVT", "佛得角标准时间"}
	WET  = TimeZone{"wet", "WET", "欧洲西部时区，GMT   格林威治标准时间"}
	CET  = TimeZone{"cet", "CET", "欧洲中部时区"}
	EET  = TimeZone{"eet", "EET", "欧洲东部时区"}
	MSK  = TimeZone{"msk", "MSK", "莫斯科时区"}
	IRT  = TimeZone{"irt", "IRT", "伊朗标准时间"}
	META = TimeZone{"meta", "META", "中东时区A"}
	AFT  = TimeZone{"aft", "AFT", "阿富汗标准时间"}
	METB = TimeZone{"metb", "METB", "中东时区B"}
	IDT  = TimeZone{"idt", "IDT", "印度标准时间"}
	NPT  = TimeZone{"npt", "NPT", "尼泊尔标准时间"}
	BHT  = TimeZone{"bht", "BHT", "孟加拉标准时间"}
	MRT  = TimeZone{"mrt", "MRT", "缅甸标准时间"}
	IST  = TimeZone{"ist", "IST", "中南半岛标准时间"}
	EAT  = TimeZone{"eat", "EAT", "东亚标准时间/中国标准时间(BJT)"}
	KRT  = TimeZone{"krt", "KRT", "朝鲜标准时间"}
	FET  = TimeZone{"fet", "FET", "远东标准时间"}
	ACST = TimeZone{"acst", "ACST", "澳大利亚中部标准时间"}
	AEST = TimeZone{"aest", "AEST", "澳大利亚东部标准时间"}
	FAST = TimeZone{"fast", "FAST", "澳大利亚远东标准时间"}
	VTT  = TimeZone{"vtt", "VTT", "瓦努阿图标准时间"}
	NFT  = TimeZone{"nft", "NFT", "诺福克岛标准时间"}
	PSTB = TimeZone{"pstb", "PSTB", "太平洋标准时间B"}
	CIT  = TimeZone{"cit", "CIT", "查塔姆群岛标准时间"}
	PSTC = TimeZone{"pstc", "PSTC", "太平洋标准时间C"}
	PSTD = TimeZone{"pstd", "PSTD", "太平洋标准时间D"}
)

// In 是否在数组里面
func In(enums []Enum, value string) bool {
	for _, v := range enums {
		if v.Value == value {
			return true
		}
	}
	return false
}

// Color 颜色
type Color string

const (
	// Black 黑色
	Black Color = "black"
	// Red 红色
	Red Color = "red"
	// Blue 蓝色
	Blue Color = "blue"
)

var (
	// ColorMap 遍历是随机顺序的
	ColorMap = map[Color]Enum{
		Black: Enum{string(Black), "Black", "黑色"},
		Red:   Enum{string(Red), "Red", "红色"},
		Blue:  Enum{string(Blue), "Blue", "蓝色"},
	}
)
