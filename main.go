package Nozomi_challenge

import (
	"fmt"
	"strings"
)

func main(){

}

// return a string containing more than twice occurrence of a char
// we can flag to not consider space as char

func AppearingMoreThanTwiceChars(s string,countSpace bool) (string,error){
	if s ==""{
		return "", fmt.Errorf("String is empty")
	}
	s=strings.ToLower(s)
	tempMap:=make(map[int32]int,len(s))
	var resultString string
	for _,ch:=range s{
		if string(ch) ==" " && countSpace !=true{
			continue
		}
		if tempMap[ch]++;tempMap[ch] ==2{
			resultString = resultString+string(ch)
		}
	}
	return resultString,nil
}
