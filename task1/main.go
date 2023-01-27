// 1. youtubeからタイトルに「SHOWROOM」が入ってるyoutubeの動画のURLを最新順に100件取得し出力

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("faild to load env : %v", err)
	}
	searchVideo()
}

func searchVideo() {

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))
	if err != nil {
		fmt.Printf("unable create service : %v", err)
	}

	var urls []string
	var nextPageToken string

	for i := 0; i <= 1; i++ {

		call := service.Search.List([]string{"snippet"}).Q("SHOWROOM").Type("video").Order("date").MaxResults(50).PageToken(nextPageToken)

		response, err := call.Do()
		if err != nil {
			fmt.Printf("Error making API call to list channels: %v", err.Error())
		}

		nextPageToken = response.NextPageToken

		for _, searchResult := range response.Items {
			url := fmt.Sprintf("https://www.youtube.com/watch?v=%v", searchResult.Id.VideoId)
			urls = append(urls, url)
		}

	}

	// 出力する
	for _, v := range urls {
		fmt.Println(v)
	}

}
