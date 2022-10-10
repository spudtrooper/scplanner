package api

import (
	"bytes"
	"fmt"

	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/scplanner/log"
)

// https://mholt.github.io/json-to-go/
type TradeContractsSearchInfo struct {
	DSP     string                       `json:"dsp"`
	Page    int                          `json:"page"`
	Pages   int                          `json:"pages"`
	Size    int                          `json:"size"`
	Results []TradeContractsSearchResult `json:"results"`
}

type TradeContractsSearchResult struct {
	ID       string    `json:"id"`
	Created  int64     `json:"created"`
	Updated  int64     `json:"updated"`
	UserID   string    `json:"userId"`
	Dsp      string    `json:"dsp"`
	Media    MediaInfo `json:"media"`
	Variants []struct {
		MinimumFollowers int `json:"minimumFollowers"`
		Targets          []struct {
			ID         string `json:"id"`
			Created    int64  `json:"created"`
			Updated    int64  `json:"updated"`
			Dsp        string `json:"dsp"`
			Type       string `json:"type"`
			ExternalID string `json:"externalId"`
			OwnerID    string `json:"ownerId"`
			Metadata   struct {
				ID       string      `json:"id"`
				Name     string      `json:"name"`
				URL      string      `json:"url"`
				ImageURL string      `json:"imageUrl"`
				OwnerID  interface{} `json:"ownerId"`
				Stats    struct {
					Followers  int         `json:"followers"`
					Plays      interface{} `json:"plays"`
					Reposts    interface{} `json:"reposts"`
					Comments   interface{} `json:"comments"`
					Likes      interface{} `json:"likes"`
					Popularity interface{} `json:"popularity"`
				} `json:"stats"`
			} `json:"metadata"`
			JoinDate int64 `json:"joinDate"`
			LastSync int64 `json:"lastSync"`
		} `json:"targets"`
	} `json:"variants"`
	Genre              string `json:"genre"`
	UnrepostAfterHours int    `json:"unrepostAfterHours"`
	ValidatedBids      int    `json:"validatedBids"`
	Expired            bool   `json:"expired"`
	MinimumFollowers   int    `json:"minimumFollowers"`
	MaximumFollowers   int    `json:"maximumFollowers"`
}

func (c *core) TradeContractsSearch(cOpts ...TradeContractsSearchOption) (*TradeContractsSearchInfo, error) {
	opts := MakeTradeContractsSearchOptions(cOpts...)
	page := or.Int(opts.Page(), 1)
	size := or.Int(opts.Size(), 25)
	dsp := or.String(opts.Dsp(), "SOUNDCLOUD")
	params := []param{
		{"page", page},
		{"size", size},
		{"dsp", dsp},
	}
	if opts.Genre() != "" {
		params = append(params, param{"genre", opts.Genre()})
	}
	if opts.MaximumFollowers() >= 0 {
		params = append(params, param{"maximumFollowers", opts.MaximumFollowers()})
	}
	if opts.MinimumFollowers() >= 0 {
		params = append(params, param{"minimumFollowers", opts.MinimumFollowers()})
	}
	route := createRoute("tradeContractsSearch", params...)
	var payload TradeContractsSearchInfo
	if _, err := c.get(route, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

type LastEvaluatedKeyInfo struct {
	ID struct {
		S string `json:"s"`
		// N    interface{} `json:"n"`
		// B    interface{} `json:"b"`
		// M    interface{} `json:"m"`
		// L    interface{} `json:"l"`
		// Ss   interface{} `json:"ss"`
		// Null interface{} `json:"null"`
		// Bool interface{} `json:"bool"`
		// Bs   interface{} `json:"bs"`
		// Ns   interface{} `json:"ns"`
	} `json:"id"`
	BidderUserID struct {
		S string `json:"s"`
		// N    interface{} `json:"n"`
		// B    interface{} `json:"b"`
		// M    interface{} `json:"m"`
		// L    interface{} `json:"l"`
		// Ss   interface{} `json:"ss"`
		// Null interface{} `json:"null"`
		// Bool interface{} `json:"bool"`
		// Bs   interface{} `json:"bs"`
		// Ns   interface{} `json:"ns"`
	} `json:"bidderUserId"`
	DspAndStatus struct {
		S string `json:"s"`
		// N    interface{} `json:"n"`
		// B    interface{} `json:"b"`
		// M    interface{} `json:"m"`
		// L    interface{} `json:"l"`
		// Ss   interface{} `json:"ss"`
		// Null interface{} `json:"null"`
		// Bool interface{} `json:"bool"`
		// Bs   interface{} `json:"bs"`
		// Ns   interface{} `json:"ns"`
	} `json:"dspAndStatus"`
}

func (i LastEvaluatedKeyInfo) Empty() bool {
	return i.ID.S == ""
}

type BidsInfo struct {
	LastEvaluatedKey LastEvaluatedKeyInfo `json:"lastEvaluatedKey"`
	Results          []BidsResult         `json:"results"`
}

type BidsResult struct {
	ID                   string       `json:"id"`
	Created              int64        `json:"created"`
	Updated              int64        `json:"updated"`
	BidderUserID         string       `json:"bidderUserId"`
	ContractAuthorUserID string       `json:"contractAuthorUserId"`
	Dsp                  string       `json:"dsp"`
	ContractID           string       `json:"contractId"`
	ContractMedia        MediaInfo    `json:"contractMedia"`
	BidMedia             MediaInfo    `json:"bidMedia"`
	ContractTargets      []TargetInfo `json:"contractTargets"`
	BidTargets           []TargetInfo `json:"bidTargets"`
	Status               string       `json:"status"`
	SelectedVariantIndex int          `json:"selectedVariantIndex"`
	RetryCount           int          `json:"retryCount"`
}

//go:generate genopts --function Bids --params "status:string" "dsp:string" "bidderUserId:string" "lastEvaluatedKey:LastEvaluatedKeyInfo"
func (c *core) Bids(cOpts ...BidsOption) (*BidsInfo, error) {
	opts := MakeBidsOptions(cOpts...)
	status := or.String(opts.Status(), "PENDING")
	dsp := or.String(opts.Dsp(), "SOUNDCLOUD")
	bidderUserId := or.String(opts.BidderUserId(), c.userID)
	params := []param{
		{"bidderUserId", bidderUserId},
		{"status", status},
		{"dsp", dsp},
	}
	if !opts.LastEvaluatedKey().Empty() {
		lek := opts.LastEvaluatedKey()
		lastEvaluatedKey, err := jsonMarshal(&lek)
		if err != nil {
			return nil, err
		}
		params = append(params, param{"lastEvaluatedKey", string(lastEvaluatedKey)})
	}
	route := createRoute("bids", params...)
	var payload BidsInfo
	if _, err := c.get(route, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

type AuthInfo struct {
	ID           string   `json:"id"`
	Created      int64    `json:"created"`
	Updated      int64    `json:"updated"`
	EmailAddress string   `json:"emailAddress"`
	User         UserInfo `json:"user"`
	RepostMember bool     `json:"repostMember"`
	ExternalID   string   `json:"externalId"`
	Permalink    string   `json:"permalink"`
	Subscription string   `json:"subscription"`
	UnlimitedPro bool     `json:"unlimitedPro"`
	Admin        bool     `json:"admin"`
}

type UserInfo struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	URL      string      `json:"url"`
	ImageURL string      `json:"imageUrl"`
	OwnerID  interface{} `json:"ownerId"`
	Stats    struct {
		Followers  int         `json:"followers"`
		Plays      interface{} `json:"plays"`
		Reposts    interface{} `json:"reposts"`
		Comments   interface{} `json:"comments"`
		Likes      interface{} `json:"likes"`
		Popularity interface{} `json:"popularity"`
	} `json:"stats"`
	Username       string      `json:"username"`
	Permalink      string      `json:"permalink"`
	Self           interface{} `json:"self"`
	PermalinkURL   string      `json:"permalink_url"`
	AvatarURL      string      `json:"avatar_url"`
	FollowersCount int         `json:"followers_count"`
	Subscriptions  []struct {
		Product struct {
			ID string `json:"id"`
		} `json:"product"`
	} `json:"subscriptions"`
}

//go:generate genopts --function Auth --params
func (c *core) Auth() (*AuthInfo, error) {
	route := createRoute("auth/me")
	var payload AuthInfo
	if _, err := c.get(route, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

type TargetInfo struct {
	ID         string `json:"id"`
	Created    int64  `json:"created"`
	Updated    int64  `json:"updated"`
	Dsp        string `json:"dsp"`
	Type       string `json:"type"`
	ExternalID string `json:"externalId"`
	OwnerID    string `json:"ownerId"`
	Metadata   struct {
		ID       string      `json:"id"`
		Name     string      `json:"name"`
		URL      string      `json:"url"`
		ImageURL string      `json:"imageUrl"`
		OwnerID  interface{} `json:"ownerId"`
		Stats    struct {
			Followers  int         `json:"followers"`
			Plays      interface{} `json:"plays"`
			Reposts    interface{} `json:"reposts"`
			Comments   interface{} `json:"comments"`
			Likes      interface{} `json:"likes"`
			Popularity interface{} `json:"popularity"`
		} `json:"stats"`
	} `json:"metadata"`
	JoinDate int64 `json:"joinDate"`
	LastSync int64 `json:"lastSync"`
}

type TradeContractsInfo struct {
	ID       string    `json:"id"`
	Created  int64     `json:"created"`
	Updated  int64     `json:"updated"`
	UserID   string    `json:"userId"`
	Dsp      string    `json:"dsp"`
	Media    MediaInfo `json:"media"`
	Variants []struct {
		MinimumFollowers int          `json:"minimumFollowers"`
		Targets          []TargetInfo `json:"targets"`
	} `json:"variants"`
	Genre              string `json:"genre"`
	UnrepostAfterHours int    `json:"unrepostAfterHours"`
	ValidatedBids      int    `json:"validatedBids"`
	Expired            bool   `json:"expired"`
	MinimumFollowers   int    `json:"minimumFollowers"`
	MaximumFollowers   int    `json:"maximumFollowers"`
}

//go:generate genopts --function TradeContractsSearch --params "page:int" "size:int" "dsp:string" "genre:string" "minimumFollowers:int" "maximumFollowers:int"
func (c *core) TradeContracts(id string) (*TradeContractsInfo, error) {
	route := createRoute(fmt.Sprintf("tradeContracts/%s", id))
	var payload TradeContractsInfo
	if _, err := c.get(route, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

//go:generate genopts --function Resolve --params --required "url string"
func (c *core) Resolve(url string) (*TargetInfo, error) {
	route := createRoute("resolve", param{"url", url})
	var payload TargetInfo
	if _, err := c.get(route, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}

type MediaInfo struct {
	ID         string `json:"id"`
	Created    int64  `json:"created"`
	Updated    int64  `json:"updated"`
	Dsp        string `json:"dsp"`
	Type       string `json:"type"`
	ExternalID string `json:"externalId"`
	OwnerID    string `json:"ownerId"`
	Metadata   struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		URL      string `json:"url"`
		ImageURL string `json:"imageUrl"`
		OwnerID  string `json:"ownerId"`
		Stats    struct {
			Followers  interface{} `json:"followers"`
			Plays      int         `json:"plays"`
			Reposts    int         `json:"reposts"`
			Comments   int         `json:"comments"`
			Likes      int         `json:"likes"`
			Popularity interface{} `json:"popularity"`
		} `json:"stats"`
		User struct {
			ID           string `json:"id"`
			Username     string `json:"username"`
			PermalinkURL string `json:"permalink_url"`
			AvatarURL    string `json:"avatar_url"`
		} `json:"user"`
		Title            string `json:"title"`
		Genre            string `json:"genre"`
		Duration         int    `json:"duration"`
		ArtworkURL       string `json:"artwork_url"`
		PermalinkURL     string `json:"permalink_url"`
		PlaybackCount    int    `json:"playback_count"`
		RepostsCount     int    `json:"reposts_count"`
		LikesCount       int    `json:"likes_count"`
		CommentCount     int    `json:"comment_count"`
		FavoritingsCount int    `json:"favoritings_count"`
	} `json:"metadata"`
	JoinDate interface{} `json:"joinDate"`
	LastSync interface{} `json:"lastSync"`
}

type BidInfo struct {
	ID                   string       `json:"id"`
	Created              int64        `json:"created"`
	Updated              int64        `json:"updated"`
	BidderUserID         string       `json:"bidderUserId"`
	ContractAuthorUserID string       `json:"contractAuthorUserId"`
	Dsp                  string       `json:"dsp"`
	ContractID           string       `json:"contractId"`
	ContractMedia        MediaInfo    `json:"contractMedia"`
	BidMedia             MediaInfo    `json:"bidMedia"`
	ContractTargets      []TargetInfo `json:"contractTargets"`
	BidTargets           []TargetInfo `json:"bidTargets"`
	Status               string       `json:"status"`
	SelectedVariantIndex int          `json:"selectedVariantIndex"`
	RetryCount           int          `json:"retryCount"`
}

//go:generate genopts --function CreateBid "dsp:string" "bidderUserId:string" "contractID:string" "debugBody"
func (c *core) CreateBid(contractId string, auth AuthInfo, bidMedia TargetInfo, cOpts ...CreateBidOption) (*BidInfo, error) {
	opts := MakeCreateBidOptions(cOpts...)
	dsp := or.String(opts.Dsp(), "SOUNDCLOUD")
	params := []param{
		{"userId", auth.ID},
	}
	route := createRoute("bids", params...)
	type bid struct {
		Dsp        string      `json:"dsp"`
		Type       string      `json:"type"`
		ExternalID string      `json:"externalId"`
		OwnerID    interface{} `json:"ownerId"`
		Metadata   interface{} `json:"metadata"`
		ID         interface{} `json:"id"`
		Created    interface{} `json:"created"`
		Updated    interface{} `json:"updated"`
	}
	type data struct {
		BidderUserID         interface{} `json:"bidderUserId"`
		ContractAuthorUserID interface{} `json:"contractAuthorUserId"`
		Dsp                  string      `json:"dsp"`
		ContractID           string      `json:"contractId"`
		ContractMedia        interface{} `json:"contractMedia"`
		BidMedia             bid         `json:"bidMedia"`
		ContractTargets      interface{} `json:"contractTargets"`
		BidTargets           []bid       `json:"bidTargets"`
		Status               interface{} `json:"status"`
		SelectedVariantIndex int         `json:"selectedVariantIndex"`
		ID                   interface{} `json:"id"`
		Created              interface{} `json:"created"`
		Updated              interface{} `json:"updated"`
	}
	bodyData := data{
		Dsp:        dsp,
		ContractID: contractId,
		BidMedia: bid{
			Dsp:        bidMedia.Dsp,
			Type:       bidMedia.Type,
			ExternalID: bidMedia.ExternalID,
		},
		BidTargets: []bid{{
			Dsp:        dsp,
			Type:       "USER",
			ExternalID: auth.ExternalID,
		}},
	}
	if opts.DebugBody() {
		j, err := formatAsJSON(&bodyData)
		if err != nil {
			log.Printf("ignoring: %v", err)
		} else {
			log.Printf("bodyData: %s", j)
		}
	}
	body, err := jsonMarshal(&bodyData)
	if err != nil {
		return nil, err
	}
	extraHeaders := map[string]string{
		"content-type": "application/json",
	}
	var payload BidInfo
	if _, err := c.post(route, &payload, bytes.NewBuffer(body), RequestExtraHeaders(extraHeaders)); err != nil {
		return nil, err
	}
	return &payload, nil
}

//go:generate genopts --function Bid --params --required "contractId string, auth AuthInfo, bidMedia TargetInfo"  "dsp:string" "bidderUserId:string" "contractID:string" "debugBody"
func (c *core) Bid(contractId string, auth AuthInfo, bidMedia TargetInfo, cOpts ...BidOption) (*BidInfo, error) {
	opts := MakeBidOptions(cOpts...)
	dsp := or.String(opts.Dsp(), "SOUNDCLOUD")
	params := []param{
		{"userId", auth.ID},
	}
	route := createRoute("bids", params...)
	type bid struct {
		Dsp        string      `json:"dsp"`
		Type       string      `json:"type"`
		ExternalID string      `json:"externalId"`
		OwnerID    interface{} `json:"ownerId"`
		Metadata   interface{} `json:"metadata"`
		ID         interface{} `json:"id"`
		Created    interface{} `json:"created"`
		Updated    interface{} `json:"updated"`
	}
	type data struct {
		BidderUserID         interface{} `json:"bidderUserId"`
		ContractAuthorUserID interface{} `json:"contractAuthorUserId"`
		Dsp                  string      `json:"dsp"`
		ContractID           string      `json:"contractId"`
		ContractMedia        interface{} `json:"contractMedia"`
		BidMedia             bid         `json:"bidMedia"`
		ContractTargets      interface{} `json:"contractTargets"`
		BidTargets           []bid       `json:"bidTargets"`
		Status               interface{} `json:"status"`
		SelectedVariantIndex int         `json:"selectedVariantIndex"`
		ID                   interface{} `json:"id"`
		Created              interface{} `json:"created"`
		Updated              interface{} `json:"updated"`
	}
	bodyData := data{
		Dsp:        dsp,
		ContractID: contractId,
		BidMedia: bid{
			Dsp:        bidMedia.Dsp,
			Type:       bidMedia.Type,
			ExternalID: bidMedia.ExternalID,
		},
		BidTargets: []bid{{
			Dsp:        dsp,
			Type:       "USER",
			ExternalID: auth.ExternalID,
		}},
	}
	if opts.DebugBody() {
		j, err := formatAsJSON(&bodyData)
		if err != nil {
			log.Printf("ignoring: %v", err)
		} else {
			log.Printf("bodyData: %s", j)
		}
	}
	body, err := jsonMarshal(&bodyData)
	if err != nil {
		return nil, err
	}
	extraHeaders := map[string]string{
		"content-type": "application/json",
	}
	var payload BidInfo
	if _, err := c.post(route, &payload, bytes.NewBuffer(body), RequestExtraHeaders(extraHeaders)); err != nil {
		return nil, err
	}
	return &payload, nil
}

//go:generate genopts --function DeleteBid --params "status:string" "dsp:string" "bidderUserId:string" "lastEvaluatedKey:LastEvaluatedKeyInfo"
func (c *core) DeleteBid(contractID string) (interface{}, error) {
	route := createRoute(fmt.Sprintf("bids/%s", contractID))
	var payload interface{}
	if _, err := c.delete(route, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
