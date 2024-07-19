package main
import ("fmt"
        "strconv")

func main(){
    var n int;
    var array []string;
    fmt.Scan(&n)
    for i:=0;i<n;i++{
        var word string;
        fmt.Scan(&word);
        array = append(array,word)
        }
        for _, word:= range array{
        if (len(word)>10){
            firstChar := word[0]
            lastChar := word[len(word)-1]
            countstr := strconv.Itoa(len(word)-2)
            ans:= string(firstChar)+countstr+string(lastChar)
            fmt.Println(ans)
        }else{
            fmt.Println(word)
        }
        }
    } 


//https://codeforces.com/problemset/problem/71/A
//input
// 4
// word
// localization
// internationalization
// pneumonoultramicroscopicsilicovolcanoconiosis

// output
// word
// l10n
// i18n
// p43s


