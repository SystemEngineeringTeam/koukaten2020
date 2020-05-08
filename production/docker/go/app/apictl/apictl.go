package apictl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../dbctl"
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

// BookPage は本の詳細ページに渡す構造体です
type BookPage struct {
	BookImgURL      string
	BookName        string
	BookAuthor      string
	BookPublication string
	Status          string
	BookData        string
	BookDescription string
}

// SearchedBook は検索した本を表示するための
type SearchedBook struct {
	BookImgURL    string
	BookName      string
	BookAuthor    string
	Identify      string
	PublishedDate string
}

// SearchBooks はAPIから本を検索するための関数です
// 本を登録するときに使います
func SearchBooks(keyword string) []SearchedBook {
	// スペースを+に直す
	keyword = strings.ReplaceAll(keyword, " ", "+")
	keyword = strings.ReplaceAll(keyword, "　", "+")
	// apiから値の取得
	data, err := http.Get(baseURL + "?q=" + keyword)
	fmt.Println(baseURL + "?q=" + keyword)

	result := make([]SearchedBook, 0)
	if err != nil {
		return nil
	}
	defer data.Body.Close()

	d, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return nil
	}

	b := book{}
	json.Unmarshal(d, &b)

	fmt.Println("len(b.items)", len(b.Items))

	for _, dat := range b.Items {
		tmp := SearchedBook{
			BookImgURL: dat.VolumeInfo.ImageLinks.Thumbnail,
			BookName:   dat.VolumeInfo.Title,
			// PublishedDate: dat.VolumeInfo.PublishedDate,
		}
		if len(dat.VolumeInfo.Authors) > 0 {
			tmp.BookAuthor = dat.VolumeInfo.Authors[0]
		}
		// if len(dat.VolumeInfo.IndustryIdentifiers) > 0 {
		// 	tmp.Identify = dat.VolumeInfo.IndustryIdentifiers[0].Identifier
		// }
		result = append(result, tmp)
	}

	return result
}

// BookDetail は本の詳細ページに情報を渡す関数です
func BookDetail(id string) BookPage {
	data, err := http.Get(baseURL + "?q=id:" + id)
	if err != nil {
		log.Println(err)
	}
	defer data.Body.Close()
	d, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
	}
	b := book{}

	json.Unmarshal(d, &b)

	detail := BookPage{
		BookImgURL:      b.Items[0].VolumeInfo.ImageLinks.Thumbnail,
		BookName:        b.Items[0].VolumeInfo.Title,
		BookAuthor:      b.Items[0].VolumeInfo.Authors[0],
		BookPublication: b.Items[0].VolumeInfo.PublishedDate,
		Status:          "hoge",
		BookData:        "hogehoge",
		BookDescription: b.Items[0].VolumeInfo.Description,
	}

	return detail
}

// BookRegister は
func BookRegister(id string) dbctl.Book {

	// apiから値の取得
	data, err := http.Get(baseURL + "?q=id:" + id)
	if err != nil {
		log.Println(err)
	}
	defer data.Body.Close()
	d, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
	}
	b := book{}

	json.Unmarshal(d, &b)

	// // Book は本の登録、詳細な情報の表示に使用する構造体
	// type Book struct {
	// 	RFID          string
	// 	Status        string
	// 	PlaceID       int
	// 	BookName      string
	// 	Author        string
	// 	Publisher     string
	// 	PublishedDate string
	// 	Description   string
	// 	ISBN          string
	// }

	book := dbctl.Book{
		RFID:          "hoge",
		Status:        "hogehoge",
		PlaceID:       0,
		BookName:      b.Items[0].VolumeInfo.Title,
		Author:        b.Items[0].VolumeInfo.Authors[0],
		Publisher:     b.Items[0].VolumeInfo.Publisher,
		PublishedDate: b.Items[0].VolumeInfo.PublishedDate,
		Description:   b.Items[0].VolumeInfo.Description,
		APIID:         b.Items[0].ID,
	}

	return book
}
