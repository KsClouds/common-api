package weather

import (
	"fmt"
	"testing"
)

func TestWeather(t *testing.T) {
	rsp := getWeather("成都")
	fmt.Println(rsp)
}
