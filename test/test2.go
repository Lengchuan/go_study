package test2

//全局变量声明
var (
	name  = "lengchuan"
	emial = "lishuijun1992@gmail.com"
)

const (
	PI   = 3.14
	PI_1 = 3.1415926
)

//普通变量声明
type (
	inttype int     //根据平台32位或者64位
	inttype8 int8   //	-128~127
	uinttype8 uint8 //	0~255
	inttype32 int32
	stringtype string
	bytetype byte
	booltype bool //1字节  true false
)
