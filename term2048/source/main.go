/*
	Author(s):	drburnett
	Update(s):	2019/03/06
	Version(s):	0.1.8
	Note(s):	Will change Game.move??() to recursion, but for now enjoy the forception.
*/


package main

import (
    "fmt"
    "math/rand"
    "time"
	"github.com/eiannone/keyboard"
    "os"
    "os/exec"
    "runtime"
	"strconv"
)

func main(){
	game := Game{}
	game.new()
	game.print()
	
	quit := false
	for quit == false {
		char, _, err := keyboard.GetSingleKey()
		if (err != nil) {}

		switch char {
			case 'w': 
				game.moveUp()
				if(game.addTile() == false){
					quit = true
				}
				break
			case 'a':
				game.moveLeft()
				if(game.addTile() == false){
					quit = true
				}
				break
			case 's':
				game.moveDown()
				if(game.addTile() == false){
					quit = true
				}
				break
			case 'd':
				game.moveRight()
				if(game.addTile() == false){
					quit = true
				}
				break
			case 'q':
				quit = true
			default: 
		}
		game.print()
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("  Author(s):	DRBurnett")
	fmt.Println("  Update(s):	2019/03/06")
	fmt.Println("  Version(s):	0.1.8");
	fmt.Println("")
	fmt.Println("")
}




type Game struct {
	board [4][4]int
	score int
}
func (g *Game) addTile() bool{
	random_numbers := [3]int{0,0,0}
	for i := 0; i < len(random_numbers); i++ {
		s := rand.NewSource(time.Now().UnixNano())
		n := rand.New(s)

		if(i < 2){// positions: 0,1,2,3
			random_numbers[i] = n.Intn(3)
		}else{// tile: 2,4
			random_numbers[i] = (n.Intn(2) + 1)*2
		}
	}
	if(g.board[random_numbers[0]][random_numbers[1]] == 0){
		g.board[random_numbers[0]][random_numbers[1]] = random_numbers[2]
		return true
	}else{
		for i := 0; i < len(g.board); i++ {
			for j := 0; j < len(g.board[i]); j++ {
				if(g.board[i][j] == 0){
					g.board[i][j] = random_numbers[2]
					return true
				}
			}
		}
	}
	return false
}
func (g *Game) new() {
	random_numbers := [6]int{0,0,0,0,0,0}
	for i := 0; i < len(random_numbers); i++ {
		s := rand.NewSource(time.Now().UnixNano())
		n := rand.New(s)

		if(i < 4){// positions: 0,1,2,3
			random_numbers[i] = n.Intn(3)
		}else{// tile: 2,4
			random_numbers[i] = (n.Intn(2) + 1)*2
		}
	}
	// Makes sure they are not duplicate locations
	if(random_numbers[0] == random_numbers[2] && random_numbers[1] == random_numbers[3]){
		random_numbers[0]++;
		if(random_numbers[0] > 3){random_numbers[0] = 2}
	}
	
	g.board = [4][4]int{  
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	g.board[random_numbers[0]][random_numbers[1]] = random_numbers[4]
	g.board[random_numbers[2]][random_numbers[3]] = random_numbers[5]
	g.score = 0
}
func (g Game) print(){
	// clear screen
	switch runtime.GOOS {
		case "windows": 
			cmd := exec.Command("cmd", "/c", "cls")
    	    cmd.Stdout = os.Stdout
	        cmd.Run()
	        break;
		default:
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
	}
	fmt.Println("  term2048             Scord:",g.score)
	fmt.Println("")
	fmt.Println("  w: Slide number up.")
	fmt.Println("  a: Slide number left.")
	fmt.Println("  s: Slide number right.")
	fmt.Println("  d: Slide number down.")
	fmt.Println("  q: Quit")
	fmt.Println("")
	fmt.Println("  ┏━━━━━━━┳━━━━━━━┳━━━━━━━┳━━━━━━━┓")
	for i := 0; i < len(g.board); i++ {
		fmt.Println("  ┃       ┃       ┃       ┃       ┃")
		fmt.Printf("  ┃%s┃%s┃%s┃%s┃\r\n",g.outputFormat(i, 0),g.outputFormat(i, 1),g.outputFormat(i, 2),g.outputFormat(i, 3))
		fmt.Println("  ┃       ┃       ┃       ┃       ┃")
		if(i < len(g.board) - 1){
			fmt.Println("  ┣━━━━━━━╋━━━━━━━╋━━━━━━━╋━━━━━━━┫")
		}
	}
	fmt.Print("  ┗━━━━━━━┻━━━━━━━┻━━━━━━━┻━━━━━━━┛ ")
}
func (g Game) outputFormat(i int, j int) string{
	s := strconv.Itoa(g.board[i][j]);

	if(g.board[i][j] > 9999){
		return " " + s + " "
	}else if (g.board[i][j] > 999){
		return "  " + s + " "
	}else if (g.board[i][j] > 99){
		return "  " + s + "  "
	}else if (g.board[i][j] > 9){
		return "   " + s + "  "
	}else if (g.board[i][j] > 0){
		return "   " + s + "   "
	}else{
		return "       "
	}
}
func (g *Game) moveAdjustment(j int, jAdjustment int, i int, iAdjustment int){
	if(g.board[j + jAdjustment][i + iAdjustment] == 0){
		g.board[j + jAdjustment][i + iAdjustment] = g.board[j][i];
		g.board[j][i] = 0;
	}else if(g.board[j + jAdjustment][i + iAdjustment] == g.board[j][i]){
		g.board[j + jAdjustment][i + iAdjustment] = g.board[j + jAdjustment][i + iAdjustment] + g.board[j][i];
		g.score = g.score + g.board[j + jAdjustment][i + iAdjustment];
		g.board[j][i] = 0;
	}
}
func (g *Game) moveDown(){
	for loop := 0; loop < 3; loop++ {// loop 3 times
		for i := 0; i < 4; i++ {// start on the left to right
			length := len(g.board) -1
			for j := length; j > -1; j-- {// bottom up
				if(j < length){
					g.moveAdjustment(j, 1, i, 0)
				}
			}
		}
	}
}
func (g *Game) moveUp(){
	for loop := 0; loop < 4; loop++ {// loop 3 times
		for i := 0; i < 4; i++ {// start on the left to right
			length := len(g.board)
			for j := 0; j < length; j++ {// bottom up
				if(j > 0){
					g.moveAdjustment(j, -1, i, 0)
				}
			}
		}
	}
}
func (g *Game) moveLeft(){
	for loop := 0; loop < 3; loop++ {// loop 3 times
		for i := 0; i < 4; i++ {// start on the left to right
			length := len(g.board)
			for j := 0; j < length; j++ {// bottom up
				if(j > 0){
					g.moveAdjustment(i, 0, j, -1)
				}
			}
		}
	}
}

func (g *Game) moveRight(){
	for loop := 0; loop < 3; loop++ {// loop 3 times
		for i := 0; i < 4; i++ {// start on the left to right
			length := len(g.board) - 1
			for j := length; j > -1; j-- {// bottom up
				if(j < length){
					g.moveAdjustment(i, 0, j, 1)
				}
			}
		}
	}
}
