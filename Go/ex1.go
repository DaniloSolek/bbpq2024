package main

import(
	"fmt"
	"encoding/base64"
	"encoding/hex"
)

func main(){
	
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	resposta := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	
	decoded, err := hex.DecodeString(hexString)
	if err != nil{
		fmt.Println("Erro:", err)
		return
	}
	b64String := base64.StdEncoding.EncodeToString(decoded)

	if b64String == resposta{
		fmt.Println("Sucesso: ", b64String)
	}
}