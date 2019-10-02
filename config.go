package mymod

import "rsc.io/quote"

func Config() string {
	return "mymod config"
}

func Hello() string {
	return quote.Hello()
}
