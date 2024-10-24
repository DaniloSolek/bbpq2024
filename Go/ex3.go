package main

import(
	"fmt"
	"encoding/hex"
)

func main(){

	hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decrypted, err := hex.DecodeString(hexStr)
	if err != nil{
		fmt.Println("Erro:", err)
		return
	}
	best := make([]byte, len(decrypted))
	scoremax := 0
	for i:=0; i<256; i++{
		xored := make([]byte, len(decrypted))
		for j:=0; j<len(decrypted); j++{
			xored[j] = decrypted[j] ^ byte(i)
		}
		score := 0
		for k:=0; k<len(xored); k++{
			if !(xored[k] < 'A' || xored[k] > 'Z' && xored[k] < 'a' || xored[k] > 'z'){
				score = score + 1
			}
			if(score > scoremax){
				scoremax = score
				best = xored
			}
		}

	}

	fmt.Println(string(best))

}