package service


//无禁手 No_Firb 、三三禁手 Three_Firb 、四四禁手 Four_Firb 、长连禁手 Long_firb
//无子 None 、黑子 Black 、白子 White

var NoForb = 0
var ThreeForb = 1
var FourForb = 2
var LongForb = 3

var None = 0
var Black = 1

//判断点(x,y)在棋盘上是否构成禁手，并返回结果，chessboard[15][15]为当前棋盘，
//x为点的横坐标，即列数；y为点的纵坐标，即行数

func KeyPointForbCheck(chessboard [16][16]int,x int,y int,adjsame int,direction int)int{
	var i,j int//关键点坐标(i,j)
	adjsame ++
	if direction >= 4 { adjsame = - adjsame}

	switch(direction % 4){
	case 0: i = x;j = y - adjsame; break
	case 1: i = x + adjsame ; j = y - adjsame; break
	case 2: i = x + adjsame ; j = y;break
	default: i = x + adjsame ; j = y + adjsame;break
	}

	if i > 15 { i = 14}else if i < 0 { i = 0 }
	if j > 15 { j = 14}else if j < 0 { j = 0 }

	chessboard[x][y] = Black//向期盼中放入棋子
	chessboard[i][j] = Black

	var result = ForbCheck(chessboard,i,j)

	chessboard[i][j] = None
	chessboard[x][y] = None

	return result
}