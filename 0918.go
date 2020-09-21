package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
)

func main(){
	str := "飞机"
	fmt.Println(str)
	HASH(str,"md5",false);
	//fmt.Println(str1)
}

func HASH(text string, hashtype string,isHex bool)[]byte{
	var hashinstance hash.Hash
	hashinstance = md5.New()
	if isHex{
		arr,_ :=hex.DecodeString(text)
		hashinstance.Write(arr)
	}else{
		hashinstance.Write([]byte(text))
	}
	cipherBytes := hashinstance.Sum(nil)
	fmt.Printf("%x",cipherBytes)
	return  cipherBytes
}
