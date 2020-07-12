package service

func ForbCheck(chessboard [16][16]int,x int,y int)int{
	//数组下标代表方向
	var adjsame [8]int//记录与(x,y)相邻连续黑色棋子数
	var adjempty [8]int//记录adjsame后相邻连续空位数
	var jumpsame [8]int//记录adjempty后的连续黑色棋子数
	var jumpempty [8]int//记录jumpsame后的空位数
	var jumpjumpsame [8]int//记录jumpempty后的连续黑色棋子数

	for i := 0;i < 8;i++{
		adjsame[i], adjempty[i],jumpsame[i],jumpempty[i],jumpjumpsame[i] = 0,0,0,0,0
	}

	//搜索
	chessboard[x][y] = Black;
	adjsame[0],adjempty[0],jumpsame[0],jumpempty[0],jumpjumpsame[0]  = delay(x,y - 1,chessboard,0,-1)//向上搜索
	adjsame[1],adjempty[1],jumpsame[1],jumpempty[1],jumpjumpsame[1]  = delay(x + 1,y - 1,chessboard,1,-1)//右上搜索
	adjsame[2],adjempty[2],jumpsame[2],jumpempty[2],jumpjumpsame[2]  = delay(x + 1,y,chessboard,1,0)//向右搜索
	adjsame[3],adjempty[3],jumpsame[3],jumpempty[3],jumpjumpsame[3]  = delay(x + 1,y + 1,chessboard,1,1)//右下搜索
	adjsame[4],adjempty[4],jumpsame[4],jumpempty[4],jumpjumpsame[4]  = delay(x,y + 1,chessboard,0,1)//向下搜索
	adjsame[5],adjempty[5],jumpsame[5],jumpempty[5],jumpjumpsame[5]  = delay(x - 1,y + 1,chessboard,-1,1)//左下搜索
	adjsame[6],adjempty[6],jumpsame[6],jumpempty[6],jumpjumpsame[6]  = delay(x - 1,y,chessboard,-1,0)//向左搜索
	adjsame[7],adjempty[7],jumpsame[7],jumpempty[7],jumpjumpsame[7]  = delay(x - 1,y - 1,chessboard,-1,-1)//左上搜索

	//分析
	chessboard[x][y] = None
	for i := 0;i < 4;i++{//先检查是否成连五，若成连五，黑棋获胜，不构成禁手
		if (adjsame[i] + adjsame[i + 4]) == 4{
			return NoForb
		}
	}

	var threecount = 0//禁手分析，棋型统计数
	var fourcount = 0

	for i := 0;i < 4;i++{
		if (adjsame[i] + adjsame[i + 4]) >= 5{//五子以上相连，长连禁手
			return LongForb
		}else if (adjsame[i] + adjsame[i + 4]) == 3{//四子相连
			var isfour = false
			if adjempty[i] > 0{//0000?
				isfour = HuoSiChongSi(chessboard,x,y,adjsame,i)
			}
			if adjempty[i + 4] > 0{//0000?
				isfour = HuoSiChongSi(chessboard,x,y,adjsame,i + 4)
			}
			if isfour { fourcount++ }
		}else if (adjsame[i] + adjsame[i + 4]) == 2{//三子相连
			if adjempty[i] == 1 && jumpsame[i] == 1{
				if HuoSiChongSi(chessboard,x,y,adjsame,i) {fourcount ++}
			}
			if adjempty[i + 4] == 1 && jumpsame[i + 4] == 1{
				if HuoSiChongSi(chessboard,x,y,adjsame,i + 4) {fourcount ++}
			}
			var isthree = false
			if (adjempty[i] > 2 || adjempty[i] == 2 && jumpsame[i] ==0) && (adjempty[i + 4] > 1 || adjempty[i + 4] == 1 && jumpsame[i + 4] == 0){
				isthree = HuoSiChongSi(chessboard,x,y,adjsame,i)
			}
			if (adjempty[i + 4] > 2 || adjempty[i + 4] == 2 && jumpsame[i + 4] ==0) && (adjempty[i] > 1 || adjempty[i] == 1 && jumpsame[i] == 0){
				isthree = HuoSiChongSi(chessboard,x,y,adjsame,i + 4)
			}
			if isthree { threecount++ }
		}else if(adjsame[i] + adjsame[i + 4]) == 1{//两子相连
			if adjempty[i] == 1 && jumpsame[i] == 2{//活四冲实
				if HuoSiChongSi(chessboard,x,y,adjsame,i + 4) { fourcount ++}
			}
			if adjempty[i + 4] == 1 && jumpsame[i + 4] ==2{
				if HuoSiChongSi(chessboard,x,y,adjsame,i + 4) { fourcount ++}
			}
			if adjempty[i] == 1 && jumpsame[i] == 1 && (jumpempty[i] > 1 || jumpempty[i] == 1 &&jumpsame[i] == 0) && (adjempty[i + 4] > 1 || adjempty[i + 4] == 1 && jumpsame[i + 4] == 0){
				if HuoSiChongSi(chessboard,x,y,adjsame,i) { threecount ++}//活三
			}
			if adjempty[i + 4] == 1 && jumpsame[i + 4] == 1 && (jumpempty[i + 4] > 1 || jumpempty[i + 4] == 1 &&jumpsame[i + 4] == 0) && (adjempty[i] > 1 || adjempty[i] == 1 && jumpsame[i] == 0){
				if HuoSiChongSi(chessboard,x,y,adjsame,i + 4) { threecount ++}
			}else if adjsame[i] + adjsame[i + 4] == 0 { //单独一子
				if adjempty[i] == 1 && jumpsame[i] == 3{
					if HuoSiChongSi(chessboard,x,y,adjsame,i) { fourcount ++}
				}
				if adjempty[i + 4] == 1 && jumpsame[i + 4] == 3{
					if HuoSiChongSi(chessboard,x,y,adjsame,i + 4) { fourcount ++}
				}
				if adjempty[i] == 1 && jumpsame[i] == 2 && (jumpempty[i] > 1 || jumpempty[i] == 1 && jumpjumpsame[i] == 0) && (adjempty[i + 4] > 1 || adjempty[i + 4] == 1 && jumpjumpsame[i + 4] == 0){
					if HuoSiChongSi(chessboard,x,y,adjsame,i) { threecount ++}
				}
				if adjempty[i + 4] == 1 && jumpsame[i + 4] == 2 && (jumpempty[i + 4] > 1 || jumpempty[i + 4] == 1 && jumpjumpsame[i + 4] == 0) && (adjempty[i] > 1 || adjempty[i] == 1 && jumpjumpsame[i] == 0){
					if HuoSiChongSi(chessboard,x,y,adjsame,i + 4) { threecount ++}
				}
			}
		}
	}

	if fourcount > 1{ return FourForb
	}
	if threecount > 1 { return ThreeForb
	}
	return NoForb
}

func delay(x1 int,y1 int,chessboard [16][16]int,i int,j int)(as int,ae int,js int,je int,jjs int){
	for ;x1 < 15 && x1 >= 0 && y1 < 15 && y1 >= 0 && chessboard[x1][y1] == Black; { x1 += i;y1 += j;as ++}
	for ;x1 < 15 && x1 >= 0 && y1 < 15 && y1 >= 0 && chessboard[x1][y1] == None;  { x1 += i;y1 += j;ae ++}
	for ;x1 < 15 && x1 >= 0 && y1 < 15 && y1 >= 0 && chessboard[x1][y1] == Black; { x1 += i;y1 += j;js ++}
	for ;x1 < 15 && x1 >= 0 && y1 < 15 && y1 >= 0 && chessboard[x1][y1] == None;  { x1 += i;y1 += j;je ++}
	for ;x1 < 15 && x1 >= 0 && y1 < 15 && y1 >= 0 && chessboard[x1][y1] == Black; { x1 += i;y1 += j;jjs++}
	return as , ae , js , je , jjs
}

func HuoSiChongSi(chessboard [16][16]int,x int,y int,adjsame [8]int,i int)bool{ ////活四冲四判断
	if KeyPointForbCheck(chessboard,x,y,adjsame[i],i) == NoForb {               //通过递归关键点是否可下
		return true
	}
	return false
}