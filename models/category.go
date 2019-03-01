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

type Category struct {
	Id     string                   `json:"id" firestore:"-"`
	Slug   string                   `json:"slug" firestore:"slug"`
	Name   string                   `json:"name" firestore:"name"`
	Parent *firestore.DocumentRef   `json:"parent" firestore:"parent,omitempty"`
	Child  []*firestore.DocumentRef `json:"child" firestore:"child,omitempty"`
}

type RawCategory struct {
	Name  string        `json:"name"`
	Slug  string        `json:"slug"`
	Child []RawCategory `json:"child,omitempty"`
}

func (c *RawCategory) ToCategory(parent *firestore.DocumentRef, child []*firestore.DocumentRef) Category {
	return Category{
		Name:   c.Name,
		Slug:   c.Slug,
		Parent: parent,
		Child:  child,
	}
}

var categories = []RawCategory{
	{
		Name: "News",
		Slug: "news",
		Child: []RawCategory{
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
		Name: "Lifestyle",
		Slug: "lifestyle",
		Child: []RawCategory{
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
		Name: "Entertainment",
		Slug: "entertainment",
		Child: []RawCategory{
			{
				Name: "Industry News",
				Slug: "industry-news",
			},
			{
				Name: "Digital Creators",
				Slug: "digital-creators",
			},
			{
				Name: "Movies",
				Slug: "movies",
			},
			{
				Name: "Music",
				Slug: "music",
			},
			{
				Name: "Television",
				Slug: "television",
			},
			{
				Name: "Pop Culture",
				Slug: "pop-culture",
			},
			{
				Name: "Style",
				Slug: "style",
			},
			{
				Name: "Arts",
				Slug: "arts",
			},
			{
				Name: "Books",
				Slug: "books",
			},
		},
	},
	{
		Name: "Fun",
		Slug: "fun",
		Child: []RawCategory{
			{
				Name: "Trending",
				Slug: "trending",
			},
			{
				Name: "Amazing",
				Slug: "amazing",
			},
			{
				Name: "Cute",
				Slug: "cute",
			},
			{
				Name: "Haha",
				Slug: "haha",
			},
			{
				Name: "Weird",
				Slug: "weird",
			},
			{
				Name: "Holidays",
				Slug: "holidays",
			},
			{
				Name: "Animals",
				Slug: "animals",
			},
			{
				Name: "Memes",
				Slug: "memes",
			},
			{
				Name: "Humor",
				Slug: "humor",
			},
		},
	},
	{
		Name: "Music",
		Slug: "music",
		Child: []RawCategory{
			{
				Name: "Pop",
				Slug: "pop",
			},
			{
				Name: "Hip-hop/Rap",
				Slug: "hip-hop-rap",
			},
			{
				Name: "Country",
				Slug: "country",
			},
			{
				Name: "Latino Music",
				Slug: "latino-music",
			},
			{
				Name: "R&B Soul",
				Slug: "r-and-b-soul",
			},
			{
				Name: "Classic Rock",
				Slug: "classic-rock",
			},
			{
				Name: "Dance/electronic",
				Slug: "dance-electronic",
			},
			{
				Name: "Metal",
				Slug: "metal",
			},
			{
				Name: "Rock/Alt",
				Slug: "rock-alt",
			},
			{
				Name: "Indie/Experimental",
				Slug: "indie-experimental",
			},
		},
	},
	{
		Name: "Technology",
		Slug: "technology",
		Child: []RawCategory{
			{
				Name: "Technology Professionals & Reporters",
				Slug: "technology-professionals-and-reporters",
			},
			{
				Name: "Teach News",
				Slug: "teach-news",
			},
		},
	},
	{
		Name: "Government & Polytics",
		Slug: "government-and-polytics",
		Child: []RawCategory{
			{
				Name: "Gov Officials & Agencies",
				Slug: "gov-officials-and-agencies",
			},
		},
	},
	{
		Name: "Science",
		Slug: "science",
		Child: []RawCategory{
			{
				Name: "Science News",
				Slug: "science-sews",
			},
			{
				Name: "Space News",
				Slug: "space-sews",
			},
		},
	},
	{
		Name: "Arts & Culture",
		Slug: "arts-and-culture",
		Child: []RawCategory{
			{
				Name: "Design & Architecture",
				Slug: "design-and-architecture",
			},
			{
				Name: "Literature",
				Slug: "literature",
			},
			{
				Name: "Photography",
				Slug: "photography",
			},
			{
				Name: "Art",
				Slug: "art",
			},
			{
				Name: "Interesting Pictures",
				Slug: "interesting-pictures",
			},
		},
	},
	{
		Name: "Nonprofits",
		Slug: "nonprofits",
		Child: []RawCategory{
			{
				Name: "Humanitarian",
				Slug: "humanitarian",
			},
		},
	},
	{
		Name: "Gaming",
		Slug: "gaming",
		Child: []RawCategory{
			{
				Name: "Celebrity Gamer",
				Slug: "celebrity-gamer",
			},
			{
				Name: "Games",
				Slug: "games",
			},
			{
				Name: "Gaming News",
				Slug: "gaming-news",
			},
			{
				Name: "eSport",
				Slug: "esport",
			},
		},
	},
	{
		Name: "Sports",
		Slug: "sports",
		Child: []RawCategory{
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

func GetCategories() []RawCategory {
	return categories
}
