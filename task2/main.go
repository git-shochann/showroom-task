// 2. 日本人による直近3日以内のApex Legendsに関しての動画を人気top10を出力

package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/leekchan/timeutil"
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

	// publishedAfterはその日時以降に投稿された動画
	// 現在の時間 = publishAfter = 3日以内
	time := time.Now()
	date := time.AddDate(0, 0, -3)
	newDate := timeutil.Strftime(&date, "%Y-%m-%dT%H:%M:%S.%fZ")

	// 最初に動画のリストを取得
	listCall := service.Search.List([]string{"id"}).Type("video").Q("Apex Legends").RegionCode("jp").PublishedAfter(newDate).MaxResults(10).Order("viewCount")
	response, err := listCall.Do()
	if err != nil {
		fmt.Printf("failed to call api: %v", err)
	}

	var videoIds []string
	for _, video := range response.Items {
		videoId := video.Id.VideoId
		videoIds = append(videoIds, videoId)
	}

	// タイトルを取得する
	videoListCall := service.Videos.List([]string{"snippet"}).Id(strings.Join(videoIds, ","))

	videoListResponse, err := videoListCall.Do()
	if err != nil {
		fmt.Printf("failed to call api: %v", err)
	}

	// 出力する
	for _, video := range videoListResponse.Items {
		fmt.Println(video.Snippet.Title)
	}
}
