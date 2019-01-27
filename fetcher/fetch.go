package fetcher

import (
	"net/http"
	"fmt"
	"bufio"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
)

func Fetch(url string )([]byte,error){


	resp,err:= http.Get("https://book.douban.com/")

	if err!=nil{
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Error status code:%d",resp.StatusCode)
	}

	bodyReader:= bufio.NewReader(resp.Body)
	e:= DeterminEncoding(bodyReader)

	utf8Reader:= transform.NewReader(bodyReader,e.NewDecoder())



	return ioutil.ReadAll(utf8Reader)


}


func DeterminEncoding(r * bufio.Reader) encoding.Encoding{

	bytes,err:= r.Peek(1024)

	if err!=nil{
		log.Printf("fetch error:%v",err)
		return unicode.UTF8
	}

	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}

