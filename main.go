package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	bc := NewBlockchain()
	size := 1000 //생성할 블럭 수
	//생성할 블럭만큼 난수를 생산
	num := RandData(size, 1500) //블럭수, [0,1000)의 범위로 난수 생산
	for i := 1; i < size; i++ {
		data := strconv.Itoa(num[i])
		bc.AddBlock(data)
	}
	for _, v := range bc.Blocks {
		v.Bprint()
	}

	reqnum, _ := rand.Int(rand.Reader, big.NewInt(1500))
	request := []byte(strconv.Itoa(int(reqnum.Int64())))
	fmt.Printf("찾기 : %s\n", reqnum)
	inHash := bc.Blocks[len(bc.Blocks)-1].Hash

	for _, v := range bc.Blocks {
		block := bc.FindBlock(inHash)
		if block != nil {
			if v.EqualData(request) {
				fmt.Println("found")
				v.Bprint()
				break
			}

		}
		inHash = block.PrevBlockHash
		if block.IsGenBlock() {
			fmt.Println("Not found")
		}

	}

}
