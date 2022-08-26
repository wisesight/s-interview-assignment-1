package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Channel struct {
	ChannelId       string `json:"id"`
	Title           string `json:"snippet.title"`
	Description     string `json:"snippet.description"`
	SubscriberCount int    `json:"statistics.subscriberCount"`
}

type TrendingVideoResult struct {
	Items         []TrendingVideo `json:"items"`
	NextPageToken string          `json:"nextPageToken"`
	PageInfo      PageInfo        `json:"pageInfo"`
}

type TrendingVideo struct {
	ChannelId            string `json:"snippet.channelId"`
	Title                string `json:"snippet.title"`
	Description          string `json:"snippet.description"`
	ChannelTitle         string `json:"snippet.channelTitle"`
	DefaultLanguage      string `json:"snippet.defaultLanguage"`
	DefaultAudioLanguage string `json:"snippet.defaultAudioLanguage"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

func checkAccountExist(accountID string) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < 0.5
}

func getNewSource() []Channel {
	return []Channel{}
}

func main() {
	newChannels := getNewSource()
	fmt.Println("=== New Channels ===")
	for _, channel := range newChannels {
		fmt.Printf("Title: %s \nsubscriberCount: %d\n\n", channel.Title, channel.SubscriberCount)
	}
	fmt.Println("=== End ===")
}
