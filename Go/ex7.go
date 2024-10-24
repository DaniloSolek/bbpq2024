package main

import(
	"fmt"
	"encoding/base64"
	"crypto/aes"
	"os"
)

func main(){

	msg, err := os.ReadFile("7.txt")
	if err != nil{
		fmt.Println("Erro:", err)
		return
	}
	stringmsg := string(msg)
	key := []byte("YELLOW SUBMARINE")

	b64decrypted, err := base64.StdEncoding.DecodeString(stringmsg)
	if err != nil{
		fmt.Println("Erro:", err)
		return
	}

	cifra, err := aes.NewCipher(key)
	if err != nil{
		fmt.Println("Erro:", err)
		return
	}

	n := len(b64decrypted)
	decrypted := make([]byte, n)

	for i:=0; i<n; i+= cifra.BlockSize(){
		cifra.Decrypt(decrypted[i:i+cifra.BlockSize()], b64decrypted[i:i+cifra.BlockSize()])
	}
	
	fmt.Println(string(decrypted))
}