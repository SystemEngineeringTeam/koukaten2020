package apictl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseURL = "https://www.googleapis.com/books/v1/volumes"
)

type book struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []struct {
		Kind       string `json:"kind"`
		ID         string `json:"id"`
		Etag       string `json:"etag"`
		SelfLink   string `json:"selfLink"`
		VolumeInfo struct {
			Title               string   `json:"title"`
			Authors             []string `json:"authors"`
			Publisher           string   `json:"publisher"`
			PublishedDate       string   `json:"publishedDate"`
			Description         string   `json:"description"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
			ReadingModes struct {
				Text  bool `json:"text"`
				Image bool `json:"image"`
			} `json:"readingModes"`
			PageCount           int      `json:"pageCount"`
			PrintType           string   `json:"printType"`
			Categories          []string `json:"categories"`
			MaturityRating      string   `json:"maturityRating"`
			AllowAnonLogging    bool     `json:"allowAnonLogging"`
			ContentVersion      string   `json:"contentVersion"`
			PanelizationSummary struct {
				ContainsEpubBubbles  bool   `json:"containsEpubBubbles"`
				ContainsImageBubbles bool   `json:"containsImageBubbles"`
				EpubBubbleVersion    string `json:"epubBubbleVersion"`
				ImageBubbleVersion   string `json:"imageBubbleVersion"`
			} `json:"panelizationSummary"`
			ImageLinks struct {
				SmallThumbnail string `json:"smallThumbnail"`
				Thumbnail      string `json:"thumbnail"`
			} `json:"imageLinks"`
			Language            string `json:"language"`
			PreviewLink         string `json:"previewLink"`
			InfoLink            string `json:"infoLink"`
			CanonicalVolumeLink string `json:"canonicalVolumeLink"`
			SeriesInfo          struct {
				Kind              string `json:"kind"`
				BookDisplayNumber string `json:"bookDisplayNumber"`
				VolumeSeries      []struct {
					SeriesID       string `json:"seriesId"`
					SeriesBookType string `json:"seriesBookType"`
					OrderNumber    int    `json:"orderNumber"`
				} `json:"volumeSeries"`
			} `json:"seriesInfo"`
		} `json:"volumeInfo,omitempty"`
		SaleInfo struct {
			Country     string `json:"country"`
			Saleability string `json:"saleability"`
			IsEbook     bool   `json:"isEbook"`
			ListPrice   struct {
				Amount       float64 `json:"amount"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"listPrice"`
			RetailPrice struct {
				Amount       float64 `json:"amount"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"retailPrice"`
			BuyLink string `json:"buyLink"`
			Offers  []struct {
				FinskyOfferType int `json:"finskyOfferType"`
				ListPrice       struct {
					AmountInMicros int    `json:"amountInMicros"`
					CurrencyCode   string `json:"currencyCode"`
				} `json:"listPrice"`
				RetailPrice struct {
					AmountInMicros int    `json:"amountInMicros"`
					CurrencyCode   string `json:"currencyCode"`
				} `json:"retailPrice"`
			} `json:"offers"`
		} `json:"saleInfo,omitempty"`
		AccessInfo struct {
			Country                string `json:"country"`
			Viewability            string `json:"viewability"`
			Embeddable             bool   `json:"embeddable"`
			PublicDomain           bool   `json:"publicDomain"`
			TextToSpeechPermission string `json:"textToSpeechPermission"`
			Epub                   struct {
				IsAvailable  bool   `json:"isAvailable"`
				AcsTokenLink string `json:"acsTokenLink"`
			} `json:"epub"`
			Pdf struct {
				IsAvailable  bool   `json:"isAvailable"`
				AcsTokenLink string `json:"acsTokenLink"`
			} `json:"pdf"`
			WebReaderLink       string `json:"webReaderLink"`
			AccessViewStatus    string `json:"accessViewStatus"`
			QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
		} `json:"accessInfo"`
		SearchInfo struct {
			TextSnippet string `json:"textSnippet"`
		} `json:"searchInfo"`
	} `json:"items"`
}

// SearchBooks はAPIから本を検索するための関数です
// 本を登録するときに使います
// なんらかの構造体を返す
func SearchBooks(keyword string) {
	keyword = strings.ReplaceAll(keyword, " ", "+")
	data, err := http.Get(baseURL + "?q=" + keyword)
	fmt.Println(baseURL + "?q=" + keyword)
	if err != nil {
		return
	}
	defer data.Body.Close()

	d, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return
	}

	b := book{}
	json.Unmarshal(d, &b)
	for _, dat := range b.Items {
		fmt.Println(dat.VolumeInfo.Title)
	}
}

// BookDetail は本の詳細ページに情報を渡す関数です
// なんらかの構造体を返す
func BookDetail(ISBN string) {

}
