package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
)

var page = 1
var amountOfCoins = 0;

func main() {
	StartScrape(page)
}

func StartScrape(pageNumber int) {
	url := fmt.Sprintf("https://coinmarketcap.com/%d", pageNumber)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	ParseTableRows(doc)
	nextpage := HasNextPage(doc)
	if nextpage {
		page++
		StartScrape(page)
	} else {
		timeTrack(time.Now(), "scraping")
		fmt.Printf("Done scraping %d pages, found %d coins", page, amountOfCoins)
	}
}

func ParseTableRows(doc *goquery.Document)  {
	doc.Find("#currencies > tbody > tr").Each(func(i int, s *goquery.Selection) {
		amountOfCoins++
		// For each item found, get the band and title
		name := s.Find("td.currency-name > a.currency-name-container").Text()
		abbreviation := s.Find("td.currency-name > span.currency-symbol").Text()
		marketCap := strings.TrimSpace(s.Find("td.market-cap").Text())
		marketCapBTC := s.Find("td.market-cap").AttrOr("data-btc", "N/A")
		price := s.Find("td > a.price").Text()
		priceBTC := s.Find("td > a.price").AttrOr("data-btc", "N/A")
		volume := s.Find("td > a.volume").Text()
		volumeBTC := s.Find("td > a.volume").AttrOr("data-btc", "N/A")
		circulatingSupply := s.Find("td.circulating-supply > a").AttrOr("data-supply", "N/A")
		percentCHange24h := s.Find("td.percent-24h").Text()
		fmt.Println(name, abbreviation, marketCap, marketCapBTC, price, priceBTC, volume, volumeBTC, circulatingSupply, percentCHange24h)
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