package tweetentities

// HashtagEntity represents a hashtag which has been parsed from text.
type HashtagEntity struct {
	Indices Indices `json:"indices" bson:"indices"`
	Text    string  `json:"text" bson:"text"`
}

// URLEntity represents a URL which has been parsed from text.
type URLEntity struct {
	Indices     Indices `json:"indices" bson:"indices"`
	DisplayURL  string  `json:"display_url" bson:"display_url"`
	ExpandedURL string  `json:"expanded_url" bson:"expanded_url"`
	URL         string  `json:"url" bson:"url"`
}

// MediaEntity represents media elements associated with a Tweet.
type MediaEntity struct {
	URLEntity
	TwitterId      int64      `json:"twitter_id" bson:"twitter_id"`
	MediaURL       string     `json:"media_url" bson:"media_url"`
	MediaURLHttps  string     `json:"media_url_https" bson:"media_url_https"`
	SourceStatusID int64      `json:"source_status_id" bson:"source_status_id"`
	Type           string     `json:"type" bson:"type"`
	Sizes          MediaSizes `json:"sizes" bson:"sizes"`
	VideoInfo      VideoInfo  `json:"video_info,omitempty" bson:"video_info,omitempty"`
}

// MentionEntity represents Twitter user mentions parsed from text.
type MentionEntity struct {
	Indices    Indices `json:"indices" bson:"indices"`
	TwitterId  int64   `json:"twitter_id" bson:"twitter_id"`
	Name       string  `json:"name" bson:"name"`
	ScreenName string  `json:"screen_name" bson:"screen_name"`
}

// Indices represent the start and end offsets within text.
type Indices [2]int

// Start returns the index at which an entity starts, inclusive.
func (i Indices) Start() int {
	return i[0]
}

// End returns the index at which an entity ends, exclusive.
func (i Indices) End() int {
	return i[1]
}

// MediaSizes contain the different size media that are available.
// https://dev.twitter.com/overview/api/entities#obj-sizes
type MediaSizes struct {
	Thumb  MediaSize `json:"thumb" bson:"thumb"`
	Large  MediaSize `json:"large" bson:"large"`
	Medium MediaSize `json:"medium" bson:"medium"`
	Small  MediaSize `json:"small" bson:"small"`
}

// MediaSize describes the height, width, and resizing method used.
type MediaSize struct {
	Width  int    `json:"w" bson:"w"`
	Height int    `json:"h" bson:"h"`
	Resize string `json:"resize" bson:"resize"`
}

// VideoInfo is available on video media objects.
type VideoInfo struct {
	AspectRatio    [2]int         `json:"aspect_ratio" bson:"aspect_ratio"`
	DurationMillis int            `json:"duration_millis" bson:"duration_millis"`
	Variants       []VideoVariant `json:"variants" bson:"variants"`
}

// VideoVariant describes one of the available video formats.
type VideoVariant struct {
	ContentType string `json:"content_type" bson:"content_type"`
	Bitrate     int    `json:"bitrate" bson:"bitrate"`
	URL         string `json:"url" bson:"url"`
}
