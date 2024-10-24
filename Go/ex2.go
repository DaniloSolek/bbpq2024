package main

import(
	"fmt"
	"encoding/hex"
)

func main(){

	resposta := "746865206b696420646f6e277420706c6179"
	str1 := "1c0111001f010100061a024b53535009181c"
	decodedstr1, err := hex.DecodeString(str1)
	if err != nil{
		fmt.Println("Erro:", err)
		return
	}

	str2 := "686974207468652062756c6c277320657965"
	decodedstr2, err := hex.DecodeString(str2)
	if err != nil{
		fmt.Println("Erro:", err)
		return
	}

	n := len(decodedstr1)
	xored := make([]byte, n)

	for i:=0; i<n; i++{
		xored[i] = decodedstr1[i] ^ decodedstr2[i]
	}

	str3 := hex.EncodeToString(xored)

	if str3 == resposta{
		fmt.Println("Sucesso: ", str3)
	} else {
		fmt.Println("Erro: ", str3)
	}
}