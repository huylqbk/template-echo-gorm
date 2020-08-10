package console

import (
	"github.com/Jeffail/gabs"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"red-coins/config"
)

func Schedule() {
	Job := cron.New()
	Job.AddFunc("30 * * * *", getBitcoinPrice)
	Job.Start()

	log.Printf("Scheduled jobs loaded")
}

func getBitcoinPrice() {
	var value float64

	coinSecret := os.Getenv("COIN_SECRET")

	client := &http.Client{}
	req, err := http.NewRequest("GET","https://pro-api.coinmarketcap.com/v1/tools/price-conversion", nil)
	
	if err != nil {
		log.Panic(err)
	}

	q := url.Values{}
	q.Add("id", "1")
	q.Add("amount", "1")
	q.Add("convert", "BRL")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", coinSecret)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		log.Panic("Error sending request to server")
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	jsonParsed, err := gabs.ParseJSON([]byte(respBody))
	
	if err != nil {
		panic(err)
	}
	
	value = jsonParsed.Path("data.quote.BRL.price").Data().(float64)

	err = config.RC.Set("price", value, 0).Err()
	
	if err != nil {
		log.Panic(err)
	}

	val, err := config.RC.Get("price").Result()

	if err != nil {
		log.Panic(err)
	}

	log.Println("Geting Bitcoin Price", val)
}