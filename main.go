package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
	"github.com/jasonlvhit/gocron"
	"strconv"
	"github.com/steffen25/go-coinmarketcap-scraper/models"
)

var amountOfCoins = 0;

func main() {
	gocron.Every(10).Seconds().Do(StartScrape)
	<- gocron.Start()
}

func StartScrape() {
	doc, err := goquery.NewDocument("https://coinmarketcap.com/all/views/all/")
	if err != nil {
		log.Fatal(err)
	}

	ParseTableRows(doc)
	nextpage := HasNextPage(doc)
	if nextpage {
		StartScrape()
	} else {
		defer timeTrack(time.Now(), "scraping")
		fmt.Printf("Done scraping found %d coins", amountOfCoins)
		amountOfCoins = 0
	}
}

func ParseTableRows(doc *goquery.Document)  {
	doc.Find("#currencies-all > tbody > tr").Each(func(i int, s *goquery.Selection) {
		amountOfCoins++
		// For each item found, get the band and title
		name := s.Find("td.currency-name > a.currency-name-container").Text()
		symbol := s.Find("td.currency-name > span.currency-symbol").Text()
		marketCapusd := s.Find("td.market-cap").AttrOr("data-usd", "N/A")
		marketCapbtc := s.Find("td.market-cap").AttrOr("data-btc", "N/A")
		marketCapBTC, _ := strconv.ParseFloat(marketCapbtc, 64)
		marketCapUSD, _ := strconv.ParseInt(marketCapusd, 10, 64)
		priceusd := s.Find("td > a.price").AttrOr("data-usd", "N/A")
		priceUSD, _ := strconv.ParseFloat(priceusd, 64)
		pricebtc := s.Find("td > a.price").AttrOr("data-btc", "N/A")
		priceBTC, _ := strconv.ParseFloat(pricebtc, 64)
		circulatingSupplyAmount := s.Find("td.circulating-supply > a").AttrOr("data-supply", "N/A")
		circulatingSupply, _ := strconv.ParseInt(circulatingSupplyAmount, 10, 64)
		volumeusd := s.Find("td > a.volume").AttrOr("data-usd", "N/A")
		volumeUSD, _ := strconv.ParseInt(volumeusd, 10, 64)
		volumebtc := s.Find("td > a.volume").AttrOr("data-btc", "N/A")
		volumeBTC, _ := strconv.ParseFloat(volumebtc, 64)
		percentCHange1Husd := s.Find("td.percent-1h").AttrOr("data-usd", "N/A")
		percentCHange1HUSD, _ := strconv.ParseFloat(percentCHange1Husd, 64)
		percentCHange1Hbtc := s.Find("td.percent-1h").AttrOr("data-btc", "N/A")
		percentCHange1HBTC, _ := strconv.ParseFloat(percentCHange1Hbtc, 64)
		percentCHange24Husd := s.Find("td.percent-24h").AttrOr("data-usd", "N/A")
		percentCHange24HUSD, _ := strconv.ParseFloat(percentCHange24Husd, 64)
		percentCHange24Hbtc := s.Find("td.percent-24h").AttrOr("data-btc", "N/A")
		percentCHange24HBTC, _ := strconv.ParseFloat(percentCHange24Hbtc, 64)
		percentCHange7Dusd := s.Find("td.percent-7d").AttrOr("data-usd", "N/A")
		percentCHange7DUSD, _ := strconv.ParseFloat(percentCHange7Dusd, 64)
		percentCHange7Dbtc := s.Find("td.percent-7d").AttrOr("data-btc", "N/A")
		percentCHange7DBTC, _ := strconv.ParseFloat(percentCHange7Dbtc, 64)

		coin := &models.Coin{
			Name: name,
			Symbol: symbol,
			MarketCapUSD: marketCapUSD,
			MarketCapBTC: marketCapBTC,
			PriceUSD: priceUSD,
			PriceBTC: priceBTC,
			CirculatingSupply: circulatingSupply,
			Volume24HUSD: volumeUSD,
			Volume24HBTC: volumeBTC,
			PriceChange1HUSD: percentCHange1HUSD,
			PriceChange1HBTC: percentCHange1HBTC,
			PriceChange24HUSD: percentCHange24HUSD,
			PriceChange24HBTC: percentCHange24HBTC,
			PriceChange7DHUSD: percentCHange7DUSD,
			PriceChange7DHBTC: percentCHange7DBTC,
		}

		log.Println(coin)
	})
}


func HasNextPage(doc *goquery.Document) bool {
	var nextPage = false
	doc.Find("ul.top-paginator > li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		name := s.Find("a").Text()
		name = strings.ToLower(name)
		if strings.Contains(name, "next") {
			nextPage = true
		}
	})

	return nextPage
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}