package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//Titulo obtem o titulo de uma paǵina HTML
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) { //funcao anonima
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url) //aplica a funcao anonima (de forma goroutine) a todos urls
	}
	return c //é retornada antes da execução da requisição http da linha 17 devido a concorrencia
}
