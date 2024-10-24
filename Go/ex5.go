package main

import(
	"fmt"
	"encoding/hex"
)

func main(){
	
	stanza := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	ice := "ICE"
	resposta := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	n := len(stanza)
	xored := make([]byte, n)
	j := 0
	
	for i:=0; i<n; i++{
		xored[i] = stanza[i] ^ ice[j]
		if j != len(ice) - 1 {
			j++
		}else{
			j = 0
		}
	}

	saida := hex.EncodeToString(xored)
	if saida == resposta{
		fmt.Println("Sucesso: ", saida)
	} else {
		fmt.Println("Erro: ", saida)
	}
}