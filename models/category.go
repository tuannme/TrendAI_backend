package models

import (
	"cloud.google.com/go/firestore"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
)

var categoryCollection *firestore.CollectionRef

func init() {
	categoryCollection = databases.GetFirestoreCollection("categories")
}

func GetCategoryCollection() *firestore.CollectionRef {
	return categoryCollection
}

type BaseCategory struct {
	Id   string `json:"id" firestore:"-"`
	Slug string `json:"slug" firestore:"slug"`
	Name string `json:"name" firestore:"name"`
}

type SubCategory BaseCategory

type Category struct {
	BaseCategory
	SubCategories []SubCategory `json:"sub_categories" firestore:"sub_categories"`
}

//
//subMusic := []string{"Pop", "Hip-hop/Rap", "Country", "Latino Music", "R&B Soul", "Classic Rock", "Dance/electronic",
//	"Metal", "Rock/Alt", "Indie/Experimental",}
//
//subEntertainment := []string{"Industry News", "Digital Creators", "Movies", "Music", "Television", "Pop Culture",
//	"Style", "Arts", "Books",}
//

//subArts := []string{"Design & Architecture", "Literature", "Photography", "Art", "Interesting Pictures",}
//
//subGovernment := []string{"Gov Officials & Agencies",}
//
//subGame := []string{"Celebrity Gamer", "Games", "Gaming News", "eSport",}
//
//subNonprofits := []string{"Humanitarian",}
//
//subFun := []string{"Trending", "Amazing", "Cute", "Haha", "Weird", "Holidays", "Animals", "Memes", "Humor",}
//
//subScience := []string{"Science News", "Space News",}
//
//subTechnology := []string{"Technology Professionals & Reporters", "Teach News",}
//
//titles := []string{
//	"News",
//	"Lifestyle",
//	"Entertainment",
//	"Fun",
//	"Music",
//	"Technology",
//	"Government & Polytics",
//	"Science",
//	"Arts & Culture",
//	"Nonprofits",
//	"Sports",
//	"Gaming",
//}

var categories = []Category{
	{
		BaseCategory: BaseCategory{
			Name: "News",
			Slug: "news",
		},
		SubCategories: []SubCategory{
			{
				Name: "Weather",
				Slug: "weather",
			},
			{
				Name: "History",
				Slug: "history",
			},
			{
				Name: "Politics",
				Slug: "politics",
			},
			{
				Name: "Health",
				Slug: "health",
			},
			{
				Name: "General News",
				Slug: "general-news",
			},
			{
				Name: "Business & Finance",
				Slug: "business-and-finance",
			},
			{
				Name: "US News",
				Slug: "us-news",
			},
			{
				Name: "World News",
				Slug: "world-news",
			},
			{
				Name: "Technology",
				Slug: "technology",
			},
			{
				Name: "Science",
				Slug: "science",
			},
		},
	},
	{
		BaseCategory: BaseCategory{
			Name: "Lifestyle",
			Slug: "lifestyle",
		},
		SubCategories: []SubCategory{
			{
				Name: "Parenting",
				Slug: "parenting",
			},
			{
				Name: "DIY & Home",
				Slug: "diy-and-home",
			},
			{
				Name: "Travel",
				Slug: "travel",
			},
			{
				Name: "Finess & Wellness",
				Slug: "finess-and-wellness",
			},
			{
				Name: "Carc Culture",
				Slug: "carc-culture",
			},
			{
				Name: "Fashion & Beauty",
				Slug: "fashion-and-beauty",
			},
			{
				Name: "Lifestyle Personalities",
				Slug: "lifestyle-personalities",
			},
			{
				Name: "Food",
				Slug: "food",
			},
		},
	},
	{
		BaseCategory: BaseCategory{
			Name: "Sports",
			Slug: "sports",
		},
		SubCategories: []SubCategory{
			{
				Name: "NEL",
				Slug: "nel",
			},
			{
				Name: "NBA",
				Slug: "nba",
			},
			{
				Name: "MLB",
				Slug: "mlb",
			},
			{
				Name: "Soccer",
				Slug: "soccer",
			},
			{
				Name: "NASCAR",
				Slug: "nascar",
			},
			{
				Name: "WWE",
				Slug: "wwe",
			},
			{
				Name: "MMA",
				Slug: "mma",
			},
			{
				Name: "Golf",
				Slug: "golf",
			},
			{
				Name: "Tennis",
				Slug: "tennis",
			},
			{
				Name: "Basketball",
				Slug: "basketball",
			},
			{
				Name: "Track & Field",
				Slug: "track-and-field",
			},
			{
				Name: "Premeier League",
				Slug: "premeier-league",
			},
			{
				Name: "Olympics",
				Slug: "olympics",
			},
			{
				Name: "UFC",
				Slug: "ufc",
			},
			{
				Name: "MLS",
				Slug: "mls",
			},
			{
				Name: "PGA",
				Slug: "pga",
			},
			{
				Name: "Hockey",
				Slug: "hockey",
			},
			{
				Name: "Wrestling",
				Slug: "wrestling",
			},
		},
	},
}

func GetCategories() []Category {
	return categories
}
