package dujitang

import (
	"context"
	"io/ioutil"
	"net/http"
)

type Args struct {
}

type Reply struct {
	Dujitang string
}

type Dujitang int

func (t *Dujitang) Get(ctx context.Context, args *Args, reply *Reply) error {
	response, err := http.Get("https://du.shadiao.app/api.php")
	if err == nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			reply.Dujitang = string(body)
		}
	}
	return nil
}
