package ebay

import (
	"encoding/xml"
	"time"
)

type AddDisputeRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DisputeExplanation   string `xml:",omitempty"`
	DisputeReason        string `xml:",omitempty"`
	ItemID               string `xml:",omitempty"`
	OrderLineItemID      string `xml:",omitempty"`
	TransactionID        string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type AddDisputeResponse struct {
	XMLName               xml.Name
	DisputeID             string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type AddDisputeResponseRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DisputeActivity      string    `xml:",omitempty"`
	DisputeID            string    `xml:",omitempty"`
	MessageText          string    `xml:",omitempty"`
	ShipmentTrackNumber  string    `xml:",omitempty"`
	ShippingCarrierUsed  string    `xml:",omitempty"`
	ShippingTime         time.Time `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type AddDisputeResponseResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type AddFixedPriceItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Item                 *ItemType `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type AddFixedPriceItemResponse struct {
	XMLName                    xml.Name
	Category2ID                string                          `xml:",omitempty"`
	CategoryID                 string                          `xml:",omitempty"`
	DiscountReason             string                          `xml:",omitempty"`
	EndTime                    time.Time                       `xml:",omitempty"`
	Fees                       *FeesType                       `xml:",omitempty"`
	ItemID                     string                          `xml:",omitempty"`
	ListingRecommendations     *ListingRecommendationsType     `xml:",omitempty"`
	ProductSuggestions         *ProductSuggestionsType         `xml:",omitempty"`
	SKU                        string                          `xml:",omitempty"`
	StartTime                  time.Time                       `xml:",omitempty"`
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Message                    string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type AddItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Item                 *ItemType `xml:",omitempty"`
	ErrorHandling        string    `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}

type AddItemResponse struct {
	XMLName                    xml.Name
	Category2ID                string                          `xml:",omitempty"`
	CategoryID                 string                          `xml:",omitempty"`
	DiscountReason             string                          `xml:",omitempty"`
	EndTime                    *time.Time                      `xml:",omitempty"`
	Fees                       *FeesType                       `xml:",omitempty"`
	ItemID                     string                          `xml:",omitempty"`
	ListingRecommendations     *ListingRecommendationsType     `xml:",omitempty"`
	ProductSuggestions         *ProductSuggestionsType         `xml:",omitempty"`
	StartTime                  *time.Time                      `xml:",omitempty"`
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     struct {
		Error []*struct {
			Code int
		}
	}
	HardExpirationWarning string     `xml:",omitempty"`
	Message               string     `xml:",omitempty"`
	Timestamp             *time.Time `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type AddItemFromSellingManagerTemplateRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Item                 *ItemType `xml:",omitempty"`
	SaleTemplateID       float64   `xml:",omitempty"`
	ScheduleTime         time.Time `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type AddItemFromSellingManagerTemplateResponse struct {
	XMLName               xml.Name
	Category2ID           string     `xml:",omitempty"`
	CategoryID            string     `xml:",omitempty"`
	EndTime               time.Time  `xml:",omitempty"`
	Fees                  *FeesType  `xml:",omitempty"`
	ItemID                string     `xml:",omitempty"`
	StartTime             time.Time  `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type AddItemsRequest struct {
	XMLName                 xml.Name
	RequesterCredentials    *RequesterCredentialsType
	AddItemRequestContainer *AddItemRequestContainerType `xml:",omitempty"`
	ErrorHandling           string                       `xml:",omitempty"`
	ErrorLanguage           string                       `xml:",omitempty"`
	MessageID               string                       `xml:",omitempty"`
	Version                 string                       `xml:",omitempty"`
	WarningLevel            string                       `xml:",omitempty"`
}
type AddItemsResponse struct {
	XMLName                    xml.Name
	AddItemResponseContainer   *AddItemResponseContainerType   `xml:",omitempty"`
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Message                    string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type AddMemberMessageAAQToPartnerRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ItemID               string             `xml:",omitempty"`
	MemberMessage        *MemberMessageType `xml:",omitempty"`
	ErrorLanguage        string             `xml:",omitempty"`
	MessageID            string             `xml:",omitempty"`
	Version              string             `xml:",omitempty"`
	WarningLevel         string             `xml:",omitempty"`
}
type AddMemberMessageAAQToPartnerResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type AddMemberMessageRTQRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ItemID               string             `xml:",omitempty"`
	MemberMessage        *MemberMessageType `xml:",omitempty"`
	ErrorLanguage        string             `xml:",omitempty"`
	MessageID            string             `xml:",omitempty"`
	Version              string             `xml:",omitempty"`
	WarningLevel         string             `xml:",omitempty"`
}
type AddMemberMessageRTQResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type AddMemberMessagesAAQToBidderRequest struct {
	XMLName                                      xml.Name
	RequesterCredentials                         *RequesterCredentialsType
	AddMemberMessagesAAQToBidderRequestContainer *AddMemberMessagesAAQToBidderRequestContainerType `xml:",omitempty"`
	ErrorLanguage                                string                                            `xml:",omitempty"`
	MessageID                                    string                                            `xml:",omitempty"`
	Version                                      string                                            `xml:",omitempty"`
	WarningLevel                                 string                                            `xml:",omitempty"`
}
type AddMemberMessagesAAQToBidderResponse struct {
	XMLName                                       xml.Name
	AddMemberMessagesAAQToBidderResponseContainer *AddMemberMessagesAAQToBidderResponseContainerType `xml:",omitempty"`
	Ack                                           string                                             `xml:",omitempty"`
	Build                                         string                                             `xml:",omitempty"`
	CorrelationID                                 string                                             `xml:",omitempty"`
	Errors                                        *ErrorType                                         `xml:",omitempty"`
	HardExpirationWarning                         string                                             `xml:",omitempty"`
	Timestamp                                     time.Time                                          `xml:",omitempty"`
	Version                                       string                                             `xml:",omitempty"`
}
type AddOrderRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Order                *OrderType `xml:",omitempty"`
	ErrorLanguage        string     `xml:",omitempty"`
	InvocationID         string     `xml:",omitempty"`
	MessageID            string     `xml:",omitempty"`
	Version              string     `xml:",omitempty"`
	WarningLevel         string     `xml:",omitempty"`
}
type AddOrderResponse struct {
	XMLName                    xml.Name
	CreatedTime                time.Time                       `xml:",omitempty"`
	OrderID                    string                          `xml:",omitempty"`
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type AddSecondChanceItemRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	BuyItNowPrice         float64 `xml:",omitempty"`
	Duration              string  `xml:",omitempty"`
	ItemID                string  `xml:",omitempty"`
	RecipientBidderUserID string  `xml:",omitempty"`
	SellerMessage         string  `xml:",omitempty"`
	ErrorLanguage         string  `xml:",omitempty"`
	MessageID             string  `xml:",omitempty"`
	Version               string  `xml:",omitempty"`
	WarningLevel          string  `xml:",omitempty"`
}
type AddSecondChanceItemResponse struct {
	XMLName                    xml.Name
	EndTime                    time.Time                       `xml:",omitempty"`
	ItemID                     string                          `xml:",omitempty"`
	StartTime                  time.Time                       `xml:",omitempty"`
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type AddSellingManagerInventoryFolderRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Comment              string  `xml:",omitempty"`
	FolderName           string  `xml:",omitempty"`
	ParentFolderID       float64 `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type AddSellingManagerInventoryFolderResponse struct {
	XMLName               xml.Name
	FolderID              float64    `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type AddSellingManagerProductRequest struct {
	XMLName                        xml.Name
	RequesterCredentials           *RequesterCredentialsType
	FolderID                       float64                             `xml:",omitempty"`
	SellingManagerProductDetails   *SellingManagerProductDetailsType   `xml:",omitempty"`
	SellingManagerProductSpecifics *SellingManagerProductSpecificsType `xml:",omitempty"`
	ErrorLanguage                  string                              `xml:",omitempty"`
	MessageID                      string                              `xml:",omitempty"`
	Version                        string                              `xml:",omitempty"`
	WarningLevel                   string                              `xml:",omitempty"`
}
type AddSellingManagerProductResponse struct {
	XMLName                      xml.Name
	SellingManagerProductDetails *SellingManagerProductDetailsType `xml:",omitempty"`
	Ack                          string                            `xml:",omitempty"`
	Build                        string                            `xml:",omitempty"`
	CorrelationID                string                            `xml:",omitempty"`
	Errors                       *ErrorType                        `xml:",omitempty"`
	HardExpirationWarning        string                            `xml:",omitempty"`
	Timestamp                    time.Time                         `xml:",omitempty"`
	Version                      string                            `xml:",omitempty"`
}
type AddSellingManagerTemplateRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Item                 *ItemType `xml:",omitempty"`
	ProductID            float64   `xml:",omitempty"`
	SaleTemplateName     string    `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type AddSellingManagerTemplateResponse struct {
	XMLName                      xml.Name
	Category2ID                  float64                           `xml:",omitempty"`
	CategoryID                   float64                           `xml:",omitempty"`
	Fees                         *FeesType                         `xml:",omitempty"`
	SaleTemplateID               float64                           `xml:",omitempty"`
	SaleTemplateName             string                            `xml:",omitempty"`
	SellingManagerProductDetails *SellingManagerProductDetailsType `xml:",omitempty"`
	Ack                          string                            `xml:",omitempty"`
	Build                        string                            `xml:",omitempty"`
	CorrelationID                string                            `xml:",omitempty"`
	Errors                       *ErrorType                        `xml:",omitempty"`
	HardExpirationWarning        string                            `xml:",omitempty"`
	Timestamp                    time.Time                         `xml:",omitempty"`
	Version                      string                            `xml:",omitempty"`
}
type AddToItemDescriptionRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Description          string `xml:",omitempty"`
	ItemID               string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	InvocationID         string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type AddToItemDescriptionResponse struct {
	XMLName                    xml.Name
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Message                    string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type AddToWatchListRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ItemID               string            `xml:",omitempty"`
	VariationKey         *VariationKeyType `xml:",omitempty"`
	ErrorLanguage        string            `xml:",omitempty"`
	MessageID            string            `xml:",omitempty"`
	Version              string            `xml:",omitempty"`
	WarningLevel         string            `xml:",omitempty"`
}
type AddToWatchListResponse struct {
	XMLName               xml.Name
	WatchListCount        int        `xml:",omitempty"`
	WatchListMaximum      int        `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type AddTransactionConfirmationItemRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	Comments              string  `xml:",omitempty"`
	ItemID                string  `xml:",omitempty"`
	ListingDuration       string  `xml:",omitempty"`
	NegotiatedPrice       float64 `xml:",omitempty"`
	RecipientPostalCode   string  `xml:",omitempty"`
	RecipientRelationType string  `xml:",omitempty"`
	RecipientUserID       string  `xml:",omitempty"`
	VerifyEligibilityOnly string  `xml:",omitempty"`
	ErrorLanguage         string  `xml:",omitempty"`
	MessageID             string  `xml:",omitempty"`
	Version               string  `xml:",omitempty"`
	WarningLevel          string  `xml:",omitempty"`
}
type AddTransactionConfirmationItemResponse struct {
	XMLName               xml.Name
	EndTime               time.Time  `xml:",omitempty"`
	ItemID                string     `xml:",omitempty"`
	StartTime             time.Time  `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type CompleteSaleRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	FeedbackInfo         *FeedbackInfoType `xml:",omitempty"`
	ItemID               string            `xml:",omitempty"`
	ListingType          string            `xml:",omitempty"`
	OrderID              string            `xml:",omitempty"`
	OrderLineItemID      string            `xml:",omitempty"`
	Paid                 bool              `xml:",omitempty"`
	Shipment             *ShipmentType     `xml:",omitempty"`
	Shipped              bool              `xml:",omitempty"`
	TransactionID        string            `xml:",omitempty"`
	ErrorHandling        string            `xml:",omitempty"`
	ErrorLanguage        string            `xml:",omitempty"`
	MessageID            string            `xml:",omitempty"`
	Version              string            `xml:",omitempty"`
	WarningLevel         string            `xml:",omitempty"`
}
type CompleteSaleResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type ConfirmIdentityRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	SessionID            string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type ConfirmIdentityResponse struct {
	XMLName               xml.Name
	UserID                string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type DeleteMyMessagesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	MessageIDs           []MyMessagesMessageIDArrayType `xml:",omitempty"`
	ErrorLanguage        string                         `xml:",omitempty"`
	MessageID            string                         `xml:",omitempty"`
	Version              string                         `xml:",omitempty"`
	WarningLevel         string                         `xml:",omitempty"`
}
type DeleteMyMessagesResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type DeleteSellingManagerInventoryFolderRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	FolderID             float64 `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type DeleteSellingManagerInventoryFolderResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type DeleteSellingManagerItemAutomationRuleRequest struct {
	XMLName                              xml.Name
	RequesterCredentials                 *RequesterCredentialsType
	DeleteAutomatedRelistingRule         bool   `xml:",omitempty"`
	DeleteAutomatedSecondChanceOfferRule bool   `xml:",omitempty"`
	ItemID                               string `xml:",omitempty"`
	ErrorLanguage                        string `xml:",omitempty"`
	MessageID                            string `xml:",omitempty"`
	Version                              string `xml:",omitempty"`
	WarningLevel                         string `xml:",omitempty"`
}
type DeleteSellingManagerItemAutomationRuleResponse struct {
	XMLName                        xml.Name
	AutomatedListingRule           *SellingManagerAutoListType              `xml:",omitempty"`
	AutomatedRelistingRule         *SellingManagerAutoRelistType            `xml:",omitempty"`
	AutomatedSecondChanceOfferRule *SellingManagerAutoSecondChanceOfferType `xml:",omitempty"`
	Fees                           *FeesType                                `xml:",omitempty"`
	Ack                            string                                   `xml:",omitempty"`
	Build                          string                                   `xml:",omitempty"`
	CorrelationID                  string                                   `xml:",omitempty"`
	Errors                         *ErrorType                               `xml:",omitempty"`
	HardExpirationWarning          string                                   `xml:",omitempty"`
	Timestamp                      time.Time                                `xml:",omitempty"`
	Version                        string                                   `xml:",omitempty"`
}
type DeleteSellingManagerProductRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ProductID            float64 `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type DeleteSellingManagerProductResponse struct {
	XMLName                             xml.Name
	DeletedSellingManagerProductDetails *SellingManagerProductDetailsType `xml:",omitempty"`
	Ack                                 string                            `xml:",omitempty"`
	Build                               string                            `xml:",omitempty"`
	CorrelationID                       string                            `xml:",omitempty"`
	Errors                              *ErrorType                        `xml:",omitempty"`
	HardExpirationWarning               string                            `xml:",omitempty"`
	Timestamp                           time.Time                         `xml:",omitempty"`
	Version                             string                            `xml:",omitempty"`
}
type DeleteSellingManagerTemplateRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	SaleTemplateID       float64 `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type DeleteSellingManagerTemplateResponse struct {
	XMLName                 xml.Name
	DeletedSaleTemplateID   string     `xml:",omitempty"`
	DeletedSaleTemplateName string     `xml:",omitempty"`
	Ack                     string     `xml:",omitempty"`
	Build                   string     `xml:",omitempty"`
	CorrelationID           string     `xml:",omitempty"`
	Errors                  *ErrorType `xml:",omitempty"`
	HardExpirationWarning   string     `xml:",omitempty"`
	Timestamp               time.Time  `xml:",omitempty"`
	Version                 string     `xml:",omitempty"`
}
type DeleteSellingManagerTemplateAutomationRuleRequest struct {
	XMLName                              xml.Name
	RequesterCredentials                 *RequesterCredentialsType
	DeleteAutomatedListingRule           bool    `xml:",omitempty"`
	DeleteAutomatedRelistingRule         bool    `xml:",omitempty"`
	DeleteAutomatedSecondChanceOfferRule bool    `xml:",omitempty"`
	SaleTemplateID                       float64 `xml:",omitempty"`
	ErrorLanguage                        string  `xml:",omitempty"`
	MessageID                            string  `xml:",omitempty"`
	Version                              string  `xml:",omitempty"`
	WarningLevel                         string  `xml:",omitempty"`
}
type DeleteSellingManagerTemplateAutomationRuleResponse struct {
	XMLName                        xml.Name
	AutomatedListingRule           *SellingManagerAutoListType              `xml:",omitempty"`
	AutomatedRelistingRule         *SellingManagerAutoRelistType            `xml:",omitempty"`
	AutomatedSecondChanceOfferRule *SellingManagerAutoSecondChanceOfferType `xml:",omitempty"`
	Fees                           *FeesType                                `xml:",omitempty"`
	Ack                            string                                   `xml:",omitempty"`
	Build                          string                                   `xml:",omitempty"`
	CorrelationID                  string                                   `xml:",omitempty"`
	Errors                         *ErrorType                               `xml:",omitempty"`
	HardExpirationWarning          string                                   `xml:",omitempty"`
	Timestamp                      time.Time                                `xml:",omitempty"`
	Version                        string                                   `xml:",omitempty"`
}
type DisableUnpaidItemAssistanceRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DisputeID            string `xml:",omitempty"`
	ItemID               string `xml:",omitempty"`
	OrderLineItemID      string `xml:",omitempty"`
	TransactionID        string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type DisableUnpaidItemAssistanceResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type EndFixedPriceItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	EndingReason         string `xml:",omitempty"`
	ItemID               string `xml:",omitempty"`
	SKU                  string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type EndFixedPriceItemResponse struct {
	XMLName               xml.Name
	SKU                   string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type EndItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	EndingReason         string `xml:",omitempty"`
	ItemID               string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type EndItemResponse struct {
	XMLName               xml.Name
	EndTime               time.Time  `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Message               string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type EndItemsRequest struct {
	XMLName                 xml.Name
	RequesterCredentials    *RequesterCredentialsType
	EndItemRequestContainer *EndItemRequestContainerType `xml:",omitempty"`
	ErrorLanguage           string                       `xml:",omitempty"`
	MessageID               string                       `xml:",omitempty"`
	Version                 string                       `xml:",omitempty"`
	WarningLevel            string                       `xml:",omitempty"`
}
type EndItemsResponse struct {
	XMLName                  xml.Name
	EndItemResponseContainer *EndItemResponseContainerType `xml:",omitempty"`
	Ack                      string                        `xml:",omitempty"`
	Build                    string                        `xml:",omitempty"`
	CorrelationID            string                        `xml:",omitempty"`
	Errors                   *ErrorType                    `xml:",omitempty"`
	HardExpirationWarning    string                        `xml:",omitempty"`
	Timestamp                time.Time                     `xml:",omitempty"`
	Version                  string                        `xml:",omitempty"`
}
type ExtendSiteHostedPicturesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ExtensionInDays      int    `xml:",omitempty"`
	PictureURL           string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type ExtendSiteHostedPicturesResponse struct {
	XMLName               xml.Name
	PictureURL            string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type FetchTokenRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	SecretID             string `xml:",omitempty"`
	SessionID            string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type FetchTokenResponse struct {
	XMLName            xml.Name
	eBayAuthToken      string     `xml:",omitempty"`
	HardExpirationTime time.Time  `xml:",omitempty"`
	RESTToken          string     `xml:",omitempty"`
	Ack                string     `xml:",omitempty"`
	Build              string     `xml:",omitempty"`
	CorrelationID      string     `xml:",omitempty"`
	Errors             *ErrorType `xml:",omitempty"`
	Timestamp          time.Time  `xml:",omitempty"`
	Version            string     `xml:",omitempty"`
}
type GetAccountRequest struct {
	XMLName                 xml.Name
	RequesterCredentials    *RequesterCredentialsType
	AccountEntrySortType    string          `xml:",omitempty"`
	AccountHistorySelection string          `xml:",omitempty"`
	BeginDate               time.Time       `xml:",omitempty"`
	Currency                string          `xml:",omitempty"`
	EndDate                 time.Time       `xml:",omitempty"`
	ExcludeBalance          bool            `xml:",omitempty"`
	ExcludeSummary          bool            `xml:",omitempty"`
	IncludeConversionRate   bool            `xml:",omitempty"`
	InvoiceDate             time.Time       `xml:",omitempty"`
	ItemID                  string          `xml:",omitempty"`
	Pagination              *PaginationType `xml:",omitempty"`
	ErrorLanguage           string          `xml:",omitempty"`
	MessageID               string          `xml:",omitempty"`
	OutputSelector          string          `xml:",omitempty"`
	Version                 string          `xml:",omitempty"`
	WarningLevel            string          `xml:",omitempty"`
}
type GetAccountResponse struct {
	XMLName               xml.Name
	AccountEntries        *AccountEntriesType   `xml:",omitempty"`
	AccountID             string                `xml:",omitempty"`
	AccountSummary        *AccountSummaryType   `xml:",omitempty"`
	Currency              string                `xml:",omitempty"`
	EntriesPerPage        int                   `xml:",omitempty"`
	HasMoreEntries        bool                  `xml:",omitempty"`
	PageNumber            int                   `xml:",omitempty"`
	PaginationResult      *PaginationResultType `xml:",omitempty"`
	Ack                   string                `xml:",omitempty"`
	Build                 string                `xml:",omitempty"`
	CorrelationID         string                `xml:",omitempty"`
	Errors                *ErrorType            `xml:",omitempty"`
	HardExpirationWarning string                `xml:",omitempty"`
	Timestamp             time.Time             `xml:",omitempty"`
	Version               string                `xml:",omitempty"`
}
type GetAdFormatLeadsRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	EndCreationTime       time.Time `xml:",omitempty"`
	IncludeMemberMessages bool      `xml:",omitempty"`
	ItemID                string    `xml:",omitempty"`
	StartCreationTime     time.Time `xml:",omitempty"`
	Status                string    `xml:",omitempty"`
	DetailLevel           string    `xml:",omitempty"`
	ErrorLanguage         string    `xml:",omitempty"`
	MessageID             string    `xml:",omitempty"`
	OutputSelector        string    `xml:",omitempty"`
	Version               string    `xml:",omitempty"`
	WarningLevel          string    `xml:",omitempty"`
}
type GetAdFormatLeadsResponse struct {
	XMLName               xml.Name
	AdFormatLead          *AdFormatLeadType `xml:",omitempty"`
	AdFormatLeadCount     int               `xml:",omitempty"`
	Ack                   string            `xml:",omitempty"`
	Build                 string            `xml:",omitempty"`
	CorrelationID         string            `xml:",omitempty"`
	Errors                *ErrorType        `xml:",omitempty"`
	HardExpirationWarning string            `xml:",omitempty"`
	Timestamp             time.Time         `xml:",omitempty"`
	Version               string            `xml:",omitempty"`
}
type GetAllBiddersRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	CallMode              string `xml:",omitempty"`
	IncludeBiddingSummary bool   `xml:",omitempty"`
	ItemID                string `xml:",omitempty"`
	ErrorLanguage         string `xml:",omitempty"`
	MessageID             string `xml:",omitempty"`
	OutputSelector        string `xml:",omitempty"`
	Version               string `xml:",omitempty"`
	WarningLevel          string `xml:",omitempty"`
}
type GetAllBiddersResponse struct {
	XMLName               xml.Name
	BidArray              []OfferArrayType `xml:",omitempty"`
	HighBidder            string           `xml:",omitempty"`
	HighestBid            float64          `xml:",omitempty"`
	ListingStatus         string           `xml:",omitempty"`
	Ack                   string           `xml:",omitempty"`
	Build                 string           `xml:",omitempty"`
	CorrelationID         string           `xml:",omitempty"`
	Errors                *ErrorType       `xml:",omitempty"`
	HardExpirationWarning string           `xml:",omitempty"`
	Timestamp             time.Time        `xml:",omitempty"`
	Version               string           `xml:",omitempty"`
}
type GetApiAccessRulesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetApiAccessRulesResponse struct {
	XMLName               xml.Name
	ApiAccessRule         *ApiAccessRuleType `xml:",omitempty"`
	Ack                   string             `xml:",omitempty"`
	Build                 string             `xml:",omitempty"`
	CorrelationID         string             `xml:",omitempty"`
	Errors                *ErrorType         `xml:",omitempty"`
	HardExpirationWarning string             `xml:",omitempty"`
	Timestamp             time.Time          `xml:",omitempty"`
	Version               string             `xml:",omitempty"`
}
type GetBestOffersRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	BestOfferID          string          `xml:",omitempty"`
	BestOfferStatus      string          `xml:",omitempty"`
	ItemID               string          `xml:",omitempty"`
	Pagination           *PaginationType `xml:",omitempty"`
	DetailLevel          string          `xml:",omitempty"`
	ErrorLanguage        string          `xml:",omitempty"`
	MessageID            string          `xml:",omitempty"`
	OutputSelector       string          `xml:",omitempty"`
	Version              string          `xml:",omitempty"`
	WarningLevel         string          `xml:",omitempty"`
}
type GetBestOffersResponse struct {
	XMLName               xml.Name
	BestOfferArray        []BestOfferArrayType      `xml:",omitempty"`
	Item                  *ItemType                 `xml:",omitempty"`
	ItemBestOffersArray   []ItemBestOffersArrayType `xml:",omitempty"`
	PageNumber            int                       `xml:",omitempty"`
	PaginationResult      *PaginationResultType     `xml:",omitempty"`
	Ack                   string                    `xml:",omitempty"`
	Build                 string                    `xml:",omitempty"`
	CorrelationID         string                    `xml:",omitempty"`
	Errors                *ErrorType                `xml:",omitempty"`
	HardExpirationWarning string                    `xml:",omitempty"`
	Timestamp             time.Time                 `xml:",omitempty"`
	Version               string                    `xml:",omitempty"`
}
type GetBidderListRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ActiveItemsOnly      bool      `xml:",omitempty"`
	EndTimeFrom          time.Time `xml:",omitempty"`
	EndTimeTo            time.Time `xml:",omitempty"`
	GranularityLevel     string    `xml:",omitempty"`
	UserID               string    `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	OutputSelector       string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type GetBidderListResponse struct {
	XMLName               xml.Name
	Bidder                *UserType       `xml:",omitempty"`
	BidItemArray          []ItemArrayType `xml:",omitempty"`
	Ack                   string          `xml:",omitempty"`
	Build                 string          `xml:",omitempty"`
	CorrelationID         string          `xml:",omitempty"`
	Errors                *ErrorType      `xml:",omitempty"`
	HardExpirationWarning string          `xml:",omitempty"`
	Timestamp             time.Time       `xml:",omitempty"`
	Version               string          `xml:",omitempty"`
}
type GetCategoriesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	CategoryParent       string `xml:",omitempty"`
	CategorySiteID       string `xml:",omitempty"`
	LevelLimit           int    `xml:",omitempty"`
	ViewAllNodes         bool   `xml:",omitempty"`
	DetailLevel          string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	OutputSelector       string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetCategoriesResponse struct {
	XMLName               xml.Name
	CategoryArray         []CategoryArrayType `xml:",omitempty"`
	CategoryCount         int                 `xml:",omitempty"`
	CategoryVersion       string              `xml:",omitempty"`
	MinimumReservePrice   float64             `xml:",omitempty"`
	ReduceReserveAllowed  bool                `xml:",omitempty"`
	ReservePriceAllowed   bool                `xml:",omitempty"`
	UpdateTime            time.Time           `xml:",omitempty"`
	Ack                   string              `xml:",omitempty"`
	Build                 string              `xml:",omitempty"`
	CorrelationID         string              `xml:",omitempty"`
	Errors                *ErrorType          `xml:",omitempty"`
	HardExpirationWarning string              `xml:",omitempty"`
	Timestamp             time.Time           `xml:",omitempty"`
	Version               string              `xml:",omitempty"`
}
type GetCategoryFeaturesRequest struct {
	XMLName                xml.Name
	RequesterCredentials   *RequesterCredentialsType
	AllFeaturesForCategory bool   `xml:",omitempty"`
	CategoryID             string `xml:",omitempty"`
	FeatureID              string `xml:",omitempty"`
	LevelLimit             int    `xml:",omitempty"`
	ViewAllNodes           bool   `xml:",omitempty"`
	DetailLevel            string `xml:",omitempty"`
	ErrorLanguage          string `xml:",omitempty"`
	MessageID              string `xml:",omitempty"`
	OutputSelector         string `xml:",omitempty"`
	Version                string `xml:",omitempty"`
	WarningLevel           string `xml:",omitempty"`
}
type GetCategoryFeaturesResponse struct {
	XMLName               xml.Name
	Category              *CategoryFeatureType    `xml:",omitempty"`
	CategoryVersion       string                  `xml:",omitempty"`
	FeatureDefinitions    *FeatureDefinitionsType `xml:",omitempty"`
	SiteDefaults          *SiteDefaultsType       `xml:",omitempty"`
	UpdateTime            time.Time               `xml:",omitempty"`
	Ack                   string                  `xml:",omitempty"`
	Build                 string                  `xml:",omitempty"`
	CorrelationID         string                  `xml:",omitempty"`
	Errors                *ErrorType              `xml:",omitempty"`
	HardExpirationWarning string                  `xml:",omitempty"`
	Timestamp             time.Time               `xml:",omitempty"`
	Version               string                  `xml:",omitempty"`
}
type GetCategoryMappingsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	CategoryVersion      string `xml:",omitempty"`
	DetailLevel          string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetCategoryMappingsResponse struct {
	XMLName               xml.Name
	CategoryMapping       string     `xml:",omitempty"`
	CategoryVersion       string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type GetCategorySpecificsRequest struct {
	XMLName                   xml.Name
	RequesterCredentials      *RequesterCredentialsType
	CategoryID                string                     `xml:",omitempty"`
	CategorySpecific          *CategoryItemSpecificsType `xml:",omitempty"`
	CategorySpecificsFileInfo bool                       `xml:",omitempty"`
	ExcludeRelationships      bool                       `xml:",omitempty"`
	IncludeConfidence         bool                       `xml:",omitempty"`
	LastUpdateTime            time.Time                  `xml:",omitempty"`
	MaxNames                  int                        `xml:",omitempty"`
	MaxValuesPerName          int                        `xml:",omitempty"`
	Name                      string                     `xml:",omitempty"`
	ErrorLanguage             string                     `xml:",omitempty"`
	MessageID                 string                     `xml:",omitempty"`
	Version                   string                     `xml:",omitempty"`
	WarningLevel              string                     `xml:",omitempty"`
}
type GetCategorySpecificsResponse struct {
	XMLName               xml.Name
	FileReferenceID       string               `xml:",omitempty"`
	Recommendations       *RecommendationsType `xml:",omitempty"`
	TaskReferenceID       string               `xml:",omitempty"`
	Ack                   string               `xml:",omitempty"`
	Build                 string               `xml:",omitempty"`
	CorrelationID         string               `xml:",omitempty"`
	Errors                *ErrorType           `xml:",omitempty"`
	HardExpirationWarning string               `xml:",omitempty"`
	Timestamp             time.Time            `xml:",omitempty"`
	Version               string               `xml:",omitempty"`
}
type GetChallengeTokenRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetChallengeTokenResponse struct {
	XMLName               xml.Name
	AudioChallengeURL     string     `xml:",omitempty"`
	ChallengeToken        string     `xml:",omitempty"`
	ImageChallengeURL     string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type GetCharitiesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	CharityDomain        int    `xml:",omitempty"`
	CharityID            string `xml:",omitempty"`
	CharityName          string `xml:",omitempty"`
	CharityRegion        int    `xml:",omitempty"`
	Featured             bool   `xml:",omitempty"`
	IncludeDescription   bool   `xml:",omitempty"`
	MatchType            string `xml:",omitempty"`
	Query                string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetCharitiesResponse struct {
	XMLName               xml.Name
	Charity               *CharityInfoType `xml:",omitempty"`
	Ack                   string           `xml:",omitempty"`
	Build                 string           `xml:",omitempty"`
	CorrelationID         string           `xml:",omitempty"`
	Errors                *ErrorType       `xml:",omitempty"`
	HardExpirationWarning string           `xml:",omitempty"`
	Timestamp             time.Time        `xml:",omitempty"`
	Version               string           `xml:",omitempty"`
}
type GetClientAlertsAuthTokenRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetClientAlertsAuthTokenResponse struct {
	XMLName               xml.Name
	ClientAlertsAuthToken string     `xml:",omitempty"`
	HardExpirationTime    time.Time  `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type GetContextualKeywordsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	CategoryID           string `xml:",omitempty"`
	Encoding             string `xml:",omitempty"`
	URL                  string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetContextualKeywordsResponse struct {
	XMLName               xml.Name
	ContextSearchAsset    *ContextSearchAssetType `xml:",omitempty"`
	Ack                   string                  `xml:",omitempty"`
	Build                 string                  `xml:",omitempty"`
	CorrelationID         string                  `xml:",omitempty"`
	Errors                *ErrorType              `xml:",omitempty"`
	HardExpirationWarning string                  `xml:",omitempty"`
	Timestamp             time.Time               `xml:",omitempty"`
	Version               string                  `xml:",omitempty"`
}
type GetDescriptionTemplatesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	CategoryID           string    `xml:",omitempty"`
	LastModifiedTime     time.Time `xml:",omitempty"`
	MotorVehicles        bool      `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type GetDescriptionTemplatesResponse struct {
	XMLName               xml.Name
	DescriptionTemplate   *DescriptionTemplateType `xml:",omitempty"`
	LayoutTotal           int                      `xml:",omitempty"`
	ObsoleteLayoutID      int                      `xml:",omitempty"`
	ObsoleteThemeID       int                      `xml:",omitempty"`
	ThemeGroup            *ThemeGroupType          `xml:",omitempty"`
	ThemeTotal            int                      `xml:",omitempty"`
	Ack                   string                   `xml:",omitempty"`
	Build                 string                   `xml:",omitempty"`
	CorrelationID         string                   `xml:",omitempty"`
	Errors                *ErrorType               `xml:",omitempty"`
	HardExpirationWarning string                   `xml:",omitempty"`
	Timestamp             time.Time                `xml:",omitempty"`
	Version               string                   `xml:",omitempty"`
}
type GetDisputeRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DisputeID            string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetDisputeResponse struct {
	XMLName               xml.Name
	Dispute               *DisputeType `xml:",omitempty"`
	Ack                   string       `xml:",omitempty"`
	Build                 string       `xml:",omitempty"`
	CorrelationID         string       `xml:",omitempty"`
	Errors                *ErrorType   `xml:",omitempty"`
	HardExpirationWarning string       `xml:",omitempty"`
	Timestamp             time.Time    `xml:",omitempty"`
	Version               string       `xml:",omitempty"`
}
type GeteBayDetailsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DetailName           string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GeteBayDetailsResponse struct {
	XMLName                        xml.Name
	BuyerRequirementDetails        *SiteBuyerRequirementDetailsType    `xml:",omitempty"`
	CountryDetails                 *CountryDetailsType                 `xml:",omitempty"`
	CurrencyDetails                *CurrencyDetailsType                `xml:",omitempty"`
	DispatchTimeMaxDetails         *DispatchTimeMaxDetailsType         `xml:",omitempty"`
	ExcludeShippingLocationDetails *ExcludeShippingLocationDetailsType `xml:",omitempty"`
	ItemSpecificDetails            *ItemSpecificDetailsType            `xml:",omitempty"`
	ListingFeatureDetails          *ListingFeatureDetailsType          `xml:",omitempty"`
	ListingStartPriceDetails       *ListingStartPriceDetailsType       `xml:",omitempty"`
	ProductDetails                 *ProductDetailsType                 `xml:",omitempty"`
	RecoupmentPolicyDetails        *RecoupmentPolicyDetailsType        `xml:",omitempty"`
	RegionDetails                  string                              `xml:",omitempty"`
	RegionOfOriginDetails          string                              `xml:",omitempty"`
	ReturnPolicyDetails            *ReturnPolicyDetailsType            `xml:",omitempty"`
	ShippingCarrierDetails         *ShippingCarrierDetailsType         `xml:",omitempty"`
	ShippingCategoryDetails        *ShippingCategoryDetailsType        `xml:",omitempty"`
	ShippingLocationDetails        *ShippingLocationDetailsType        `xml:",omitempty"`
	ShippingPackageDetails         *ShippingPackageDetailsType         `xml:",omitempty"`
	ShippingServiceDetails         *ShippingServiceDetailsType         `xml:",omitempty"`
	SiteDetails                    *SiteDetailsType                    `xml:",omitempty"`
	TaxJurisdiction                *TaxJurisdictionType                `xml:",omitempty"`
	TimeZoneDetails                *TimeZoneDetailsType                `xml:",omitempty"`
	UpdateTime                     time.Time                           `xml:",omitempty"`
	URLDetails                     *URLDetailsType                     `xml:",omitempty"`
	VariationDetails               *VariationDetailsType               `xml:",omitempty"`
	Ack                            string                              `xml:",omitempty"`
	Build                          string                              `xml:",omitempty"`
	CorrelationID                  string                              `xml:",omitempty"`
	Errors                         *ErrorType                          `xml:",omitempty"`
	HardExpirationWarning          string                              `xml:",omitempty"`
	Timestamp                      time.Time                           `xml:",omitempty"`
	Version                        string                              `xml:",omitempty"`
}
type GeteBayOfficialTimeRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GeteBayOfficialTimeResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type GetFeedbackRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	CommentType          string          `xml:",omitempty"`
	FeedbackID           string          `xml:",omitempty"`
	FeedbackType         string          `xml:",omitempty"`
	ItemID               string          `xml:",omitempty"`
	OrderLineItemID      string          `xml:",omitempty"`
	Pagination           *PaginationType `xml:",omitempty"`
	TransactionID        string          `xml:",omitempty"`
	UserID               string          `xml:",omitempty"`
	DetailLevel          string          `xml:",omitempty"`
	ErrorLanguage        string          `xml:",omitempty"`
	MessageID            string          `xml:",omitempty"`
	OutputSelector       string          `xml:",omitempty"`
	Version              string          `xml:",omitempty"`
	WarningLevel         string          `xml:",omitempty"`
}
type GetFeedbackResponse struct {
	XMLName                 xml.Name
	EntriesPerPage          int                       `xml:",omitempty"`
	FeedbackDetailArray     []FeedbackDetailArrayType `xml:",omitempty"`
	FeedbackDetailItemTotal int                       `xml:",omitempty"`
	FeedbackScore           int                       `xml:",omitempty"`
	FeedbackSummary         *FeedbackSummaryType      `xml:",omitempty"`
	PageNumber              int                       `xml:",omitempty"`
	PaginationResult        *PaginationResultType     `xml:",omitempty"`
	Ack                     string                    `xml:",omitempty"`
	Build                   string                    `xml:",omitempty"`
	CorrelationID           string                    `xml:",omitempty"`
	Errors                  *ErrorType                `xml:",omitempty"`
	HardExpirationWarning   string                    `xml:",omitempty"`
	Timestamp               time.Time                 `xml:",omitempty"`
	Version                 string                    `xml:",omitempty"`
}
type GetItemRequest struct {
	XMLName                      xml.Name
	RequesterCredentials         *RequesterCredentialsType
	IncludeItemCompatibilityList bool                     `xml:",omitempty"`
	IncludeItemSpecifics         bool                     `xml:",omitempty"`
	IncludeTaxTable              bool                     `xml:",omitempty"`
	IncludeWatchCount            bool                     `xml:",omitempty"`
	ItemID                       string                   `xml:",omitempty"`
	SKU                          string                   `xml:",omitempty"`
	TransactionID                string                   `xml:",omitempty"`
	VariationSKU                 string                   `xml:",omitempty"`
	VariationSpecifics           []NameValueListArrayType `xml:",omitempty"`
	DetailLevel                  string                   `xml:",omitempty"`
	ErrorLanguage                string                   `xml:",omitempty"`
	MessageID                    string                   `xml:",omitempty"`
	OutputSelector               string                   `xml:",omitempty"`
	Version                      string                   `xml:",omitempty"`
	WarningLevel                 string                   `xml:",omitempty"`
}
type GetItemResponse struct {
	XMLName               xml.Name
	Item                  *ItemType  `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type GetItemsAwaitingFeedbackRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Pagination           *PaginationType `xml:",omitempty"`
	Sort                 string          `xml:",omitempty"`
	ErrorLanguage        string          `xml:",omitempty"`
	MessageID            string          `xml:",omitempty"`
	OutputSelector       string          `xml:",omitempty"`
	Version              string          `xml:",omitempty"`
	WarningLevel         string          `xml:",omitempty"`
}
type GetItemsAwaitingFeedbackResponse struct {
	XMLName               xml.Name
	ItemsAwaitingFeedback []PaginatedTransactionArrayType `xml:",omitempty"`
	Ack                   string                          `xml:",omitempty"`
	Build                 string                          `xml:",omitempty"`
	CorrelationID         string                          `xml:",omitempty"`
	Errors                *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning string                          `xml:",omitempty"`
	Timestamp             time.Time                       `xml:",omitempty"`
	Version               string                          `xml:",omitempty"`
}
type GetItemShippingRequest struct {
	XMLName                xml.Name
	RequesterCredentials   *RequesterCredentialsType
	DestinationCountryCode string `xml:",omitempty"`
	DestinationPostalCode  string `xml:",omitempty"`
	ItemID                 string `xml:",omitempty"`
	QuantitySold           int    `xml:",omitempty"`
	ErrorLanguage          string `xml:",omitempty"`
	MessageID              string `xml:",omitempty"`
	OutputSelector         string `xml:",omitempty"`
	Version                string `xml:",omitempty"`
	WarningLevel           string `xml:",omitempty"`
}
type GetItemShippingResponse struct {
	XMLName               xml.Name
	ShippingDetails       *ShippingDetailsType `xml:",omitempty"`
	Ack                   string               `xml:",omitempty"`
	Build                 string               `xml:",omitempty"`
	CorrelationID         string               `xml:",omitempty"`
	Errors                *ErrorType           `xml:",omitempty"`
	HardExpirationWarning string               `xml:",omitempty"`
	Timestamp             time.Time            `xml:",omitempty"`
	Version               string               `xml:",omitempty"`
}
type GetItemTransactionsRequest struct {
	XMLName                xml.Name
	RequesterCredentials   *RequesterCredentialsType
	IncludeContainingOrder bool            `xml:",omitempty"`
	IncludeFinalValueFee   bool            `xml:",omitempty"`
	IncludeVariations      bool            `xml:",omitempty"`
	ItemID                 string          `xml:",omitempty"`
	ModTimeFrom            time.Time       `xml:",omitempty"`
	ModTimeTo              time.Time       `xml:",omitempty"`
	NumberOfDays           int             `xml:",omitempty"`
	OrderLineItemID        string          `xml:",omitempty"`
	Pagination             *PaginationType `xml:",omitempty"`
	Platform               string          `xml:",omitempty"`
	TransactionID          string          `xml:",omitempty"`
	DetailLevel            string          `xml:",omitempty"`
	ErrorLanguage          string          `xml:",omitempty"`
	MessageID              string          `xml:",omitempty"`
	OutputSelector         string          `xml:",omitempty"`
	Version                string          `xml:",omitempty"`
	WarningLevel           string          `xml:",omitempty"`
}
type GetItemTransactionsResponse struct {
	XMLName                        xml.Name
	HasMoreTransactions            bool                   `xml:",omitempty"`
	Item                           *ItemType              `xml:",omitempty"`
	PageNumber                     int                    `xml:",omitempty"`
	PaginationResult               *PaginationResultType  `xml:",omitempty"`
	PayPalPreferred                bool                   `xml:",omitempty"`
	ReturnedTransactionCountActual int                    `xml:",omitempty"`
	TransactionArray               []TransactionArrayType `xml:",omitempty"`
	TransactionsPerPage            int                    `xml:",omitempty"`
	Ack                            string                 `xml:",omitempty"`
	Build                          string                 `xml:",omitempty"`
	CorrelationID                  string                 `xml:",omitempty"`
	Errors                         *ErrorType             `xml:",omitempty"`
	HardExpirationWarning          string                 `xml:",omitempty"`
	Timestamp                      time.Time              `xml:",omitempty"`
	Version                        string                 `xml:",omitempty"`
}
type GetMemberMessagesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DisplayToPublic      bool            `xml:",omitempty"`
	EndCreationTime      time.Time       `xml:",omitempty"`
	ItemID               string          `xml:",omitempty"`
	MailMessageType      string          `xml:",omitempty"`
	MemberMessageID      string          `xml:",omitempty"`
	MessageStatus        string          `xml:",omitempty"`
	Pagination           *PaginationType `xml:",omitempty"`
	SenderID             string          `xml:",omitempty"`
	StartCreationTime    time.Time       `xml:",omitempty"`
	ErrorLanguage        string          `xml:",omitempty"`
	MessageID            string          `xml:",omitempty"`
	OutputSelector       string          `xml:",omitempty"`
	Version              string          `xml:",omitempty"`
	WarningLevel         string          `xml:",omitempty"`
}
type GetMemberMessagesResponse struct {
	XMLName               xml.Name
	HasMoreItems          bool                             `xml:",omitempty"`
	MemberMessage         []MemberMessageExchangeArrayType `xml:",omitempty"`
	PaginationResult      *PaginationResultType            `xml:",omitempty"`
	Ack                   string                           `xml:",omitempty"`
	Build                 string                           `xml:",omitempty"`
	CorrelationID         string                           `xml:",omitempty"`
	Errors                *ErrorType                       `xml:",omitempty"`
	HardExpirationWarning string                           `xml:",omitempty"`
	Timestamp             time.Time                        `xml:",omitempty"`
	Version               string                           `xml:",omitempty"`
}
type GetMessagePreferencesRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	IncludeASQPreferences bool   `xml:",omitempty"`
	SellerID              string `xml:",omitempty"`
	ErrorLanguage         string `xml:",omitempty"`
	MessageID             string `xml:",omitempty"`
	Version               string `xml:",omitempty"`
	WarningLevel          string `xml:",omitempty"`
}
type GetMessagePreferencesResponse struct {
	XMLName               xml.Name
	ASQPreferences        *ASQPreferencesType `xml:",omitempty"`
	Ack                   string              `xml:",omitempty"`
	Build                 string              `xml:",omitempty"`
	CorrelationID         string              `xml:",omitempty"`
	Errors                *ErrorType          `xml:",omitempty"`
	HardExpirationWarning string              `xml:",omitempty"`
	Timestamp             time.Time           `xml:",omitempty"`
	Version               string              `xml:",omitempty"`
}
type GetMyeBayBuyingRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	BestOfferList        *ItemListCustomizationType `xml:",omitempty"`
	BidList              *ItemListCustomizationType `xml:",omitempty"`
	BuyingSummary        *ItemListCustomizationType `xml:",omitempty"`
	DeletedFromLostList  *ItemListCustomizationType `xml:",omitempty"`
	DeletedFromWonList   *ItemListCustomizationType `xml:",omitempty"`
	FavoriteSearches     *MyeBaySelectionType       `xml:",omitempty"`
	FavoriteSellers      *MyeBaySelectionType       `xml:",omitempty"`
	HideVariations       bool                       `xml:",omitempty"`
	LostList             *ItemListCustomizationType `xml:",omitempty"`
	SecondChanceOffer    *MyeBaySelectionType       `xml:",omitempty"`
	UserDefinedLists     *MyeBaySelectionType       `xml:",omitempty"`
	WatchList            *ItemListCustomizationType `xml:",omitempty"`
	WonList              *ItemListCustomizationType `xml:",omitempty"`
	DetailLevel          string                     `xml:",omitempty"`
	ErrorLanguage        string                     `xml:",omitempty"`
	MessageID            string                     `xml:",omitempty"`
	OutputSelector       string                     `xml:",omitempty"`
	Version              string                     `xml:",omitempty"`
	WarningLevel         string                     `xml:",omitempty"`
}
type GetMyeBayBuyingResponse struct {
	XMLName               xml.Name
	BestOfferList         []PaginatedItemArrayType             `xml:",omitempty"`
	BidList               []PaginatedItemArrayType             `xml:",omitempty"`
	BuyingSummary         *BuyingSummaryType                   `xml:",omitempty"`
	DeletedFromLostList   []PaginatedItemArrayType             `xml:",omitempty"`
	DeletedFromWonList    []PaginatedOrderTransactionArrayType `xml:",omitempty"`
	FavoriteSearches      *MyeBayFavoriteSearchListType        `xml:",omitempty"`
	FavoriteSellers       *MyeBayFavoriteSellerListType        `xml:",omitempty"`
	LostList              []PaginatedItemArrayType             `xml:",omitempty"`
	SecondChanceOffer     *ItemType                            `xml:",omitempty"`
	UserDefinedList       *UserDefinedListType                 `xml:",omitempty"`
	WatchList             []PaginatedItemArrayType             `xml:",omitempty"`
	WonList               []PaginatedOrderTransactionArrayType `xml:",omitempty"`
	Ack                   string                               `xml:",omitempty"`
	Build                 string                               `xml:",omitempty"`
	CorrelationID         string                               `xml:",omitempty"`
	Errors                *ErrorType                           `xml:",omitempty"`
	HardExpirationWarning string                               `xml:",omitempty"`
	Timestamp             time.Time                            `xml:",omitempty"`
	Version               string                               `xml:",omitempty"`
}
type GetMyeBayRemindersRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	BuyingReminders      *ReminderCustomizationType `xml:",omitempty"`
	SellingReminders     *ReminderCustomizationType `xml:",omitempty"`
	ErrorLanguage        string                     `xml:",omitempty"`
	MessageID            string                     `xml:",omitempty"`
	Version              string                     `xml:",omitempty"`
	WarningLevel         string                     `xml:",omitempty"`
}
type GetMyeBayRemindersResponse struct {
	XMLName               xml.Name
	BuyingReminders       *RemindersType `xml:",omitempty"`
	SellingReminders      *RemindersType `xml:",omitempty"`
	Ack                   string         `xml:",omitempty"`
	Build                 string         `xml:",omitempty"`
	CorrelationID         string         `xml:",omitempty"`
	Errors                *ErrorType     `xml:",omitempty"`
	HardExpirationWarning string         `xml:",omitempty"`
	Timestamp             time.Time      `xml:",omitempty"`
	Version               string         `xml:",omitempty"`
}
type GetMyeBaySellingRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	ActiveList            *ItemListCustomizationType `xml:",omitempty"`
	BidList               *ItemListCustomizationType `xml:",omitempty"`
	DeletedFromSoldList   *ItemListCustomizationType `xml:",omitempty"`
	DeletedFromUnsoldList *ItemListCustomizationType `xml:",omitempty"`
	HideVariations        bool                       `xml:",omitempty"`
	ScheduledList         *ItemListCustomizationType `xml:",omitempty"`
	SellingSummary        *ItemListCustomizationType `xml:",omitempty"`
	SoldList              *ItemListCustomizationType `xml:",omitempty"`
	UnsoldList            *ItemListCustomizationType `xml:",omitempty"`
	DetailLevel           string                     `xml:",omitempty"`
	ErrorLanguage         string                     `xml:",omitempty"`
	MessageID             string                     `xml:",omitempty"`
	OutputSelector        string                     `xml:",omitempty"`
	Version               string                     `xml:",omitempty"`
	WarningLevel          string                     `xml:",omitempty"`
}
type GetMyeBaySellingResponse struct {
	XMLName               xml.Name
	ActiveList            []PaginatedItemArrayType             `xml:",omitempty"`
	DeletedFromSoldList   []PaginatedOrderTransactionArrayType `xml:",omitempty"`
	DeletedFromUnsoldList []PaginatedItemArrayType             `xml:",omitempty"`
	ScheduledList         []PaginatedItemArrayType             `xml:",omitempty"`
	SellingSummary        *SellingSummaryType                  `xml:",omitempty"`
	SoldList              []PaginatedOrderTransactionArrayType `xml:",omitempty"`
	Summary               *MyeBaySellingSummaryType            `xml:",omitempty"`
	UnsoldList            []PaginatedItemArrayType             `xml:",omitempty"`
	Ack                   string                               `xml:",omitempty"`
	Build                 string                               `xml:",omitempty"`
	CorrelationID         string                               `xml:",omitempty"`
	Errors                *ErrorType                           `xml:",omitempty"`
	HardExpirationWarning string                               `xml:",omitempty"`
	Timestamp             time.Time                            `xml:",omitempty"`
	Version               string                               `xml:",omitempty"`
}
type GetMyMessagesRequest struct {
	XMLName                        xml.Name
	RequesterCredentials           *RequesterCredentialsType
	EndTime                        time.Time                              `xml:",omitempty"`
	ExternalMessageIDs             []MyMessagesExternalMessageIDArrayType `xml:",omitempty"`
	FolderID                       float64                                `xml:",omitempty"`
	IncludeHighPriorityMessageOnly bool                                   `xml:",omitempty"`
	MessageIDs                     []MyMessagesMessageIDArrayType         `xml:",omitempty"`
	Pagination                     *PaginationType                        `xml:",omitempty"`
	StartTime                      time.Time                              `xml:",omitempty"`
	DetailLevel                    string                                 `xml:",omitempty"`
	ErrorLanguage                  string                                 `xml:",omitempty"`
	MessageID                      string                                 `xml:",omitempty"`
	OutputSelector                 string                                 `xml:",omitempty"`
	Version                        string                                 `xml:",omitempty"`
	WarningLevel                   string                                 `xml:",omitempty"`
}
type GetMyMessagesResponse struct {
	XMLName               xml.Name
	Messages              []MyMessagesMessageArrayType `xml:",omitempty"`
	Summary               *MyMessagesSummaryType       `xml:",omitempty"`
	Ack                   string                       `xml:",omitempty"`
	Build                 string                       `xml:",omitempty"`
	CorrelationID         string                       `xml:",omitempty"`
	Errors                *ErrorType                   `xml:",omitempty"`
	HardExpirationWarning string                       `xml:",omitempty"`
	Timestamp             time.Time                    `xml:",omitempty"`
	Version               string                       `xml:",omitempty"`
}
type GetNotificationPreferencesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	PreferenceLevel      string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	OutputSelector       string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetNotificationPreferencesResponse struct {
	XMLName                        xml.Name
	ApplicationDeliveryPreferences *ApplicationDeliveryPreferencesType `xml:",omitempty"`
	DeliveryURLName                string                              `xml:",omitempty"`
	EventProperty                  *NotificationEventPropertyType      `xml:",omitempty"`
	UserData                       *NotificationUserDataType           `xml:",omitempty"`
	UserDeliveryPreferenceArray    []NotificationEnableArrayType       `xml:",omitempty"`
	Ack                            string                              `xml:",omitempty"`
	Build                          string                              `xml:",omitempty"`
	CorrelationID                  string                              `xml:",omitempty"`
	Errors                         *ErrorType                          `xml:",omitempty"`
	ExternalUserData               string                              `xml:",omitempty"`
	HardExpirationWarning          string                              `xml:",omitempty"`
	Timestamp                      time.Time                           `xml:",omitempty"`
	Version                        string                              `xml:",omitempty"`
}
type GetNotificationsUsageRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	EndTime              time.Time `xml:",omitempty"`
	ItemID               string    `xml:",omitempty"`
	StartTime            time.Time `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type GetNotificationsUsageResponse struct {
	XMLName                  xml.Name
	EndTime                  time.Time                      `xml:",omitempty"`
	MarkUpMarkDownHistory    *MarkUpMarkDownHistoryType     `xml:",omitempty"`
	NotificationDetailsArray []NotificationDetailsArrayType `xml:",omitempty"`
	NotificationStatistics   *NotificationStatisticsType    `xml:",omitempty"`
	StartTime                time.Time                      `xml:",omitempty"`
	Ack                      string                         `xml:",omitempty"`
	Build                    string                         `xml:",omitempty"`
	CorrelationID            string                         `xml:",omitempty"`
	Errors                   *ErrorType                     `xml:",omitempty"`
	HardExpirationWarning    string                         `xml:",omitempty"`
	Timestamp                time.Time                      `xml:",omitempty"`
	Version                  string                         `xml:",omitempty"`
}
type GetOrdersRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	CreateTimeFrom       time.Time          `xml:",omitempty"`
	CreateTimeTo         time.Time          `xml:",omitempty"`
	IncludeFinalValueFee bool               `xml:",omitempty"`
	ListingType          string             `xml:",omitempty"`
	ModTimeFrom          time.Time          `xml:",omitempty"`
	ModTimeTo            time.Time          `xml:",omitempty"`
	NumberOfDays         int                `xml:",omitempty"`
	OrderIDArray         []OrderIDArrayType `xml:",omitempty"`
	OrderRole            string             `xml:",omitempty"`
	OrderStatus          string             `xml:",omitempty"`
	Pagination           *PaginationType    `xml:",omitempty"`
	SortingOrder         string             `xml:",omitempty"`
	DetailLevel          string             `xml:",omitempty"`
	ErrorLanguage        string             `xml:",omitempty"`
	MessageID            string             `xml:",omitempty"`
	OutputSelector       string             `xml:",omitempty"`
	Version              string             `xml:",omitempty"`
	WarningLevel         string             `xml:",omitempty"`
}
type GetOrdersResponse struct {
	XMLName                  xml.Name
	HasMoreOrders            bool                  `xml:",omitempty"`
	OrderArray               []OrderArrayType      `xml:",omitempty"`
	OrdersPerPage            int                   `xml:",omitempty"`
	PageNumber               int                   `xml:",omitempty"`
	PaginationResult         *PaginationResultType `xml:",omitempty"`
	ReturnedOrderCountActual int                   `xml:",omitempty"`
	Ack                      string                `xml:",omitempty"`
	Build                    string                `xml:",omitempty"`
	CorrelationID            string                `xml:",omitempty"`
	Errors                   *ErrorType            `xml:",omitempty"`
	HardExpirationWarning    string                `xml:",omitempty"`
	Timestamp                time.Time             `xml:",omitempty"`
	Version                  string                `xml:",omitempty"`
}
type GetOrderTransactionsRequest struct {
	XMLName                xml.Name
	RequesterCredentials   *RequesterCredentialsType
	IncludeFinalValueFees  bool                         `xml:",omitempty"`
	ItemTransactionIDArray []ItemTransactionIDArrayType `xml:",omitempty"`
	OrderIDArray           []OrderIDArrayType           `xml:",omitempty"`
	Platform               string                       `xml:",omitempty"`
	DetailLevel            string                       `xml:",omitempty"`
	ErrorLanguage          string                       `xml:",omitempty"`
	MessageID              string                       `xml:",omitempty"`
	OutputSelector         string                       `xml:",omitempty"`
	Version                string                       `xml:",omitempty"`
	WarningLevel           string                       `xml:",omitempty"`
}
type GetOrderTransactionsResponse struct {
	XMLName               xml.Name
	OrderArray            []OrderArrayType `xml:",omitempty"`
	Ack                   string           `xml:",omitempty"`
	Build                 string           `xml:",omitempty"`
	CorrelationID         string           `xml:",omitempty"`
	Errors                *ErrorType       `xml:",omitempty"`
	HardExpirationWarning string           `xml:",omitempty"`
	Timestamp             time.Time        `xml:",omitempty"`
	Version               string           `xml:",omitempty"`
}
type GetPromotionalSaleDetailsRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	PromotionalSaleID     float64 `xml:",omitempty"`
	PromotionalSaleStatus string  `xml:",omitempty"`
	ErrorLanguage         string  `xml:",omitempty"`
	MessageID             string  `xml:",omitempty"`
	Version               string  `xml:",omitempty"`
	WarningLevel          string  `xml:",omitempty"`
}
type GetPromotionalSaleDetailsResponse struct {
	XMLName                xml.Name
	PromotionalSaleDetails []PromotionalSaleArrayType `xml:",omitempty"`
	Ack                    string                     `xml:",omitempty"`
	Build                  string                     `xml:",omitempty"`
	CorrelationID          string                     `xml:",omitempty"`
	Errors                 *ErrorType                 `xml:",omitempty"`
	HardExpirationWarning  string                     `xml:",omitempty"`
	Timestamp              time.Time                  `xml:",omitempty"`
	Version                string                     `xml:",omitempty"`
}
type GetSellerDashboardRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetSellerDashboardResponse struct {
	XMLName               xml.Name
	BuyerSatisfaction     *BuyerSatisfactionDashboardType `xml:",omitempty"`
	Performance           *PerformanceDashboardType       `xml:",omitempty"`
	PolicyCompliance      string                          `xml:",omitempty"`
	PowerSellerStatus     *PowerSellerDashboardType       `xml:",omitempty"`
	SearchStanding        *SearchStandingDashboardType    `xml:",omitempty"`
	SellerAccount         *SellerAccountDashboardType     `xml:",omitempty"`
	SellerFeeDiscount     *SellerFeeDiscountDashboardType `xml:",omitempty"`
	Ack                   string                          `xml:",omitempty"`
	Build                 string                          `xml:",omitempty"`
	CorrelationID         string                          `xml:",omitempty"`
	Errors                *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning string                          `xml:",omitempty"`
	Timestamp             time.Time                       `xml:",omitempty"`
	Version               string                          `xml:",omitempty"`
}
type GetSellerEventsRequest struct {
	XMLName                   xml.Name
	RequesterCredentials      *RequesterCredentialsType
	EndTimeFrom               time.Time `xml:",omitempty"`
	EndTimeTo                 time.Time `xml:",omitempty"`
	HideVariations            bool      `xml:",omitempty"`
	IncludeVariationSpecifics bool      `xml:",omitempty"`
	IncludeWatchCount         bool      `xml:",omitempty"`
	ModTimeFrom               time.Time `xml:",omitempty"`
	ModTimeTo                 time.Time `xml:",omitempty"`
	NewItemFilter             bool      `xml:",omitempty"`
	StartTimeFrom             time.Time `xml:",omitempty"`
	StartTimeTo               time.Time `xml:",omitempty"`
	UserID                    string    `xml:",omitempty"`
	DetailLevel               string    `xml:",omitempty"`
	ErrorLanguage             string    `xml:",omitempty"`
	MessageID                 string    `xml:",omitempty"`
	OutputSelector            string    `xml:",omitempty"`
	Version                   string    `xml:",omitempty"`
	WarningLevel              string    `xml:",omitempty"`
}
type GetSellerEventsResponse struct {
	XMLName               xml.Name
	ItemArray             []ItemArrayType `xml:",omitempty"`
	TimeTo                time.Time       `xml:",omitempty"`
	Ack                   string          `xml:",omitempty"`
	Build                 string          `xml:",omitempty"`
	CorrelationID         string          `xml:",omitempty"`
	Errors                *ErrorType      `xml:",omitempty"`
	HardExpirationWarning string          `xml:",omitempty"`
	Timestamp             time.Time       `xml:",omitempty"`
	Version               string          `xml:",omitempty"`
}
type GetSellerListRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	AdminEndedItemsOnly  bool              `xml:",omitempty"`
	CategoryID           int               `xml:",omitempty"`
	EndTimeFrom          time.Time         `xml:",omitempty"`
	EndTimeTo            time.Time         `xml:",omitempty"`
	GranularityLevel     string            `xml:",omitempty"`
	IncludeVariations    bool              `xml:",omitempty"`
	IncludeWatchCount    bool              `xml:",omitempty"`
	MotorsDealerUsers    []UserIDArrayType `xml:",omitempty"`
	Pagination           *PaginationType   `xml:",omitempty"`
	SKUArray             []SKUArrayType    `xml:",omitempty"`
	Sort                 int               `xml:",omitempty"`
	StartTimeFrom        time.Time         `xml:",omitempty"`
	StartTimeTo          time.Time         `xml:",omitempty"`
	UserID               string            `xml:",omitempty"`
	DetailLevel          string            `xml:",omitempty"`
	ErrorLanguage        string            `xml:",omitempty"`
	MessageID            string            `xml:",omitempty"`
	OutputSelector       string            `xml:",omitempty"`
	Version              string            `xml:",omitempty"`
	WarningLevel         string            `xml:",omitempty"`
}
type GetSellerListResponse struct {
	XMLName                 xml.Name
	HasMoreItems            bool                  `xml:",omitempty"`
	ItemArray               []ItemArrayType       `xml:",omitempty"`
	ItemsPerPage            int                   `xml:",omitempty"`
	PageNumber              int                   `xml:",omitempty"`
	PaginationResult        *PaginationResultType `xml:",omitempty"`
	ReturnedItemCountActual int                   `xml:",omitempty"`
	Seller                  *UserType             `xml:",omitempty"`
	Ack                     string                `xml:",omitempty"`
	Build                   string                `xml:",omitempty"`
	CorrelationID           string                `xml:",omitempty"`
	Errors                  *ErrorType            `xml:",omitempty"`
	HardExpirationWarning   string                `xml:",omitempty"`
	Timestamp               time.Time             `xml:",omitempty"`
	Version                 string                `xml:",omitempty"`
}
type GetSellerTransactionsRequest struct {
	XMLName                 xml.Name
	RequesterCredentials    *RequesterCredentialsType
	IncludeCodiceFiscale    bool            `xml:",omitempty"`
	IncludeContainingOrder  bool            `xml:",omitempty"`
	IncludeFinalValueFee    bool            `xml:",omitempty"`
	InventoryTrackingMethod string          `xml:",omitempty"`
	ModTimeFrom             time.Time       `xml:",omitempty"`
	ModTimeTo               time.Time       `xml:",omitempty"`
	NumberOfDays            int             `xml:",omitempty"`
	Pagination              *PaginationType `xml:",omitempty"`
	Platform                string          `xml:",omitempty"`
	SKUArray                []SKUArrayType  `xml:",omitempty"`
	DetailLevel             string          `xml:",omitempty"`
	ErrorLanguage           string          `xml:",omitempty"`
	MessageID               string          `xml:",omitempty"`
	OutputSelector          string          `xml:",omitempty"`
	Version                 string          `xml:",omitempty"`
	WarningLevel            string          `xml:",omitempty"`
}
type GetSellerTransactionsResponse struct {
	XMLName                        xml.Name
	HasMoreTransactions            bool                   `xml:",omitempty"`
	PageNumber                     int                    `xml:",omitempty"`
	PaginationResult               *PaginationResultType  `xml:",omitempty"`
	PayPalPreferred                bool                   `xml:",omitempty"`
	ReturnedTransactionCountActual int                    `xml:",omitempty"`
	Seller                         *UserType              `xml:",omitempty"`
	TransactionArray               []TransactionArrayType `xml:",omitempty"`
	TransactionsPerPage            int                    `xml:",omitempty"`
	Ack                            string                 `xml:",omitempty"`
	Build                          string                 `xml:",omitempty"`
	CorrelationID                  string                 `xml:",omitempty"`
	Errors                         *ErrorType             `xml:",omitempty"`
	HardExpirationWarning          string                 `xml:",omitempty"`
	Timestamp                      time.Time              `xml:",omitempty"`
	Version                        string                 `xml:",omitempty"`
}
type GetSellingManagerAlertsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetSellingManagerAlertsResponse struct {
	XMLName               xml.Name
	Alert                 *SellingManagerAlertType `xml:",omitempty"`
	Ack                   string                   `xml:",omitempty"`
	Build                 string                   `xml:",omitempty"`
	CorrelationID         string                   `xml:",omitempty"`
	Errors                *ErrorType               `xml:",omitempty"`
	HardExpirationWarning string                   `xml:",omitempty"`
	Timestamp             time.Time                `xml:",omitempty"`
	Version               string                   `xml:",omitempty"`
}
type GetSellingManagerEmailLogRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	EmailDateRange       *TimeRangeType `xml:",omitempty"`
	ItemID               string         `xml:",omitempty"`
	OrderID              string         `xml:",omitempty"`
	OrderLineItemID      string         `xml:",omitempty"`
	TransactionID        float64        `xml:",omitempty"`
	ErrorLanguage        string         `xml:",omitempty"`
	MessageID            string         `xml:",omitempty"`
	Version              string         `xml:",omitempty"`
	WarningLevel         string         `xml:",omitempty"`
}
type GetSellingManagerEmailLogResponse struct {
	XMLName               xml.Name
	EmailLog              *SellingManagerEmailLogType `xml:",omitempty"`
	Ack                   string                      `xml:",omitempty"`
	Build                 string                      `xml:",omitempty"`
	CorrelationID         string                      `xml:",omitempty"`
	Errors                *ErrorType                  `xml:",omitempty"`
	HardExpirationWarning string                      `xml:",omitempty"`
	Timestamp             time.Time                   `xml:",omitempty"`
	Version               string                      `xml:",omitempty"`
}
type GetSellingManagerInventoryRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Filter               string                    `xml:",omitempty"`
	FolderID             float64                   `xml:",omitempty"`
	Pagination           *PaginationType           `xml:",omitempty"`
	Search               *SellingManagerSearchType `xml:",omitempty"`
	Sort                 string                    `xml:",omitempty"`
	SortOrder            string                    `xml:",omitempty"`
	StoreCategoryID      float64                   `xml:",omitempty"`
	ErrorLanguage        string                    `xml:",omitempty"`
	MessageID            string                    `xml:",omitempty"`
	Version              string                    `xml:",omitempty"`
	WarningLevel         string                    `xml:",omitempty"`
}
type GetSellingManagerInventoryResponse struct {
	XMLName                          xml.Name
	InventoryCountLastCalculatedDate time.Time                  `xml:",omitempty"`
	PaginationResult                 *PaginationResultType      `xml:",omitempty"`
	SellingManagerProduct            *SellingManagerProductType `xml:",omitempty"`
	Ack                              string                     `xml:",omitempty"`
	Build                            string                     `xml:",omitempty"`
	CorrelationID                    string                     `xml:",omitempty"`
	Errors                           *ErrorType                 `xml:",omitempty"`
	HardExpirationWarning            string                     `xml:",omitempty"`
	Timestamp                        time.Time                  `xml:",omitempty"`
	Version                          string                     `xml:",omitempty"`
}
type GetSellingManagerInventoryFolderRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	FolderID             float64 `xml:",omitempty"`
	FullRecursion        bool    `xml:",omitempty"`
	MaxDepth             int     `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type GetSellingManagerInventoryFolderResponse struct {
	XMLName               xml.Name
	Folder                *SellingManagerFolderDetailsType `xml:",omitempty"`
	Ack                   string                           `xml:",omitempty"`
	Build                 string                           `xml:",omitempty"`
	CorrelationID         string                           `xml:",omitempty"`
	Errors                *ErrorType                       `xml:",omitempty"`
	HardExpirationWarning string                           `xml:",omitempty"`
	Timestamp             time.Time                        `xml:",omitempty"`
	Version               string                           `xml:",omitempty"`
}
type GetSellingManagerItemAutomationRuleRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ItemID               string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetSellingManagerItemAutomationRuleResponse struct {
	XMLName                        xml.Name
	AutomatedListingRule           *SellingManagerAutoListType              `xml:",omitempty"`
	AutomatedRelistingRule         *SellingManagerAutoRelistType            `xml:",omitempty"`
	AutomatedSecondChanceOfferRule *SellingManagerAutoSecondChanceOfferType `xml:",omitempty"`
	Fees                           *FeesType                                `xml:",omitempty"`
	Ack                            string                                   `xml:",omitempty"`
	Build                          string                                   `xml:",omitempty"`
	CorrelationID                  string                                   `xml:",omitempty"`
	Errors                         *ErrorType                               `xml:",omitempty"`
	HardExpirationWarning          string                                   `xml:",omitempty"`
	Timestamp                      time.Time                                `xml:",omitempty"`
	Version                        string                                   `xml:",omitempty"`
}
type GetSellingManagerSaleRecordRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ItemID               string `xml:",omitempty"`
	OrderID              string `xml:",omitempty"`
	OrderLineItemID      string `xml:",omitempty"`
	TransactionID        string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetSellingManagerSaleRecordResponse struct {
	XMLName                 xml.Name
	SellingManagerSoldOrder *SellingManagerSoldOrderType `xml:",omitempty"`
	Ack                     string                       `xml:",omitempty"`
	Build                   string                       `xml:",omitempty"`
	CorrelationID           string                       `xml:",omitempty"`
	Errors                  *ErrorType                   `xml:",omitempty"`
	HardExpirationWarning   string                       `xml:",omitempty"`
	Timestamp               time.Time                    `xml:",omitempty"`
	Version                 string                       `xml:",omitempty"`
}
type GetSellingManagerSoldListingsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Archived             bool                      `xml:",omitempty"`
	Filter               string                    `xml:",omitempty"`
	Pagination           *PaginationType           `xml:",omitempty"`
	SaleDateRange        *TimeRangeType            `xml:",omitempty"`
	Search               *SellingManagerSearchType `xml:",omitempty"`
	Sort                 string                    `xml:",omitempty"`
	SortOrder            string                    `xml:",omitempty"`
	StoreCategoryID      float64                   `xml:",omitempty"`
	ErrorLanguage        string                    `xml:",omitempty"`
	MessageID            string                    `xml:",omitempty"`
	Version              string                    `xml:",omitempty"`
	WarningLevel         string                    `xml:",omitempty"`
}
type GetSellingManagerSoldListingsResponse struct {
	XMLName               xml.Name
	PaginationResult      *PaginationResultType        `xml:",omitempty"`
	SaleRecord            *SellingManagerSoldOrderType `xml:",omitempty"`
	Ack                   string                       `xml:",omitempty"`
	Build                 string                       `xml:",omitempty"`
	CorrelationID         string                       `xml:",omitempty"`
	Errors                *ErrorType                   `xml:",omitempty"`
	HardExpirationWarning string                       `xml:",omitempty"`
	Timestamp             time.Time                    `xml:",omitempty"`
	Version               string                       `xml:",omitempty"`
}
type GetSellingManagerTemplateAutomationRuleRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	SaleTemplateID       float64 `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type GetSellingManagerTemplateAutomationRuleResponse struct {
	XMLName                        xml.Name
	AutomatedListingRule           *SellingManagerAutoListType              `xml:",omitempty"`
	AutomatedRelistingRule         *SellingManagerAutoRelistType            `xml:",omitempty"`
	AutomatedSecondChanceOfferRule *SellingManagerAutoSecondChanceOfferType `xml:",omitempty"`
	Fees                           *FeesType                                `xml:",omitempty"`
	Ack                            string                                   `xml:",omitempty"`
	Build                          string                                   `xml:",omitempty"`
	CorrelationID                  string                                   `xml:",omitempty"`
	Errors                         *ErrorType                               `xml:",omitempty"`
	HardExpirationWarning          string                                   `xml:",omitempty"`
	Timestamp                      time.Time                                `xml:",omitempty"`
	Version                        string                                   `xml:",omitempty"`
}
type GetSellingManagerTemplatesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	SaleTemplateID       float64 `xml:",omitempty"`
	DetailLevel          string  `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type GetSellingManagerTemplatesResponse struct {
	XMLName                            xml.Name
	SellingManagerTemplateDetailsArray []SellingManagerTemplateDetailsArrayType `xml:",omitempty"`
	Ack                                string                                   `xml:",omitempty"`
	Build                              string                                   `xml:",omitempty"`
	CorrelationID                      string                                   `xml:",omitempty"`
	Errors                             *ErrorType                               `xml:",omitempty"`
	HardExpirationWarning              string                                   `xml:",omitempty"`
	Timestamp                          time.Time                                `xml:",omitempty"`
	Version                            string                                   `xml:",omitempty"`
}
type GetSessionIDRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	RuName               string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetSessionIDResponse struct {
	XMLName       xml.Name
	SessionID     string     `xml:",omitempty"`
	Ack           string     `xml:",omitempty"`
	Build         string     `xml:",omitempty"`
	CorrelationID string     `xml:",omitempty"`
	Errors        *ErrorType `xml:",omitempty"`
	Timestamp     time.Time  `xml:",omitempty"`
	Version       string     `xml:",omitempty"`
}
type GetShippingDiscountProfilesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetShippingDiscountProfilesResponse struct {
	XMLName                            xml.Name
	CalculatedHandlingDiscount         *CalculatedHandlingDiscountType         `xml:",omitempty"`
	CalculatedShippingDiscount         *CalculatedShippingDiscountType         `xml:",omitempty"`
	CombinedDuration                   string                                  `xml:",omitempty"`
	CurrencyID                         string                                  `xml:",omitempty"`
	FlatShippingDiscount               *FlatShippingDiscountType               `xml:",omitempty"`
	InternationalShippingInsurance     *ShippingInsuranceType                  `xml:",omitempty"`
	PromotionalShippingDiscount        bool                                    `xml:",omitempty"`
	PromotionalShippingDiscountDetails *PromotionalShippingDiscountDetailsType `xml:",omitempty"`
	ShippingInsurance                  *ShippingInsuranceType                  `xml:",omitempty"`
	Ack                                string                                  `xml:",omitempty"`
	Build                              string                                  `xml:",omitempty"`
	CorrelationID                      string                                  `xml:",omitempty"`
	Errors                             *ErrorType                              `xml:",omitempty"`
	HardExpirationWarning              string                                  `xml:",omitempty"`
	Timestamp                          time.Time                               `xml:",omitempty"`
	Version                            string                                  `xml:",omitempty"`
}
type GetStoreRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	CategoryStructureOnly bool    `xml:",omitempty"`
	LevelLimit            int     `xml:",omitempty"`
	RootCategoryID        float64 `xml:",omitempty"`
	UserID                string  `xml:",omitempty"`
	ErrorLanguage         string  `xml:",omitempty"`
	MessageID             string  `xml:",omitempty"`
	Version               string  `xml:",omitempty"`
	WarningLevel          string  `xml:",omitempty"`
}
type GetStoreResponse struct {
	XMLName               xml.Name
	Store                 *StoreType `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type GetStoreCategoryUpdateStatusRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	TaskID               float64 `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type GetStoreCategoryUpdateStatusResponse struct {
	XMLName               xml.Name
	Status                string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type GetStoreCustomPageRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	PageID               float64 `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type GetStoreCustomPageResponse struct {
	XMLName               xml.Name
	CustomPageArray       []StoreCustomPageArrayType `xml:",omitempty"`
	Ack                   string                     `xml:",omitempty"`
	Build                 string                     `xml:",omitempty"`
	CorrelationID         string                     `xml:",omitempty"`
	Errors                *ErrorType                 `xml:",omitempty"`
	HardExpirationWarning string                     `xml:",omitempty"`
	Timestamp             time.Time                  `xml:",omitempty"`
	Version               string                     `xml:",omitempty"`
}
type GetStoreOptionsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetStoreOptionsResponse struct {
	XMLName               xml.Name
	AdvancedThemeArray    []StoreThemeArrayType        `xml:",omitempty"`
	BasicThemeArray       []StoreThemeArrayType        `xml:",omitempty"`
	LogoArray             []StoreLogoArrayType         `xml:",omitempty"`
	MaxCategories         int                          `xml:",omitempty"`
	MaxCategoryLevels     int                          `xml:",omitempty"`
	SubscriptionArray     []StoreSubscriptionArrayType `xml:",omitempty"`
	Ack                   string                       `xml:",omitempty"`
	Build                 string                       `xml:",omitempty"`
	CorrelationID         string                       `xml:",omitempty"`
	Errors                *ErrorType                   `xml:",omitempty"`
	HardExpirationWarning string                       `xml:",omitempty"`
	Timestamp             time.Time                    `xml:",omitempty"`
	Version               string                       `xml:",omitempty"`
}
type GetStorePreferencesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetStorePreferencesResponse struct {
	XMLName               xml.Name
	StorePreferences      *StorePreferencesType `xml:",omitempty"`
	Ack                   string                `xml:",omitempty"`
	Build                 string                `xml:",omitempty"`
	CorrelationID         string                `xml:",omitempty"`
	Errors                *ErrorType            `xml:",omitempty"`
	HardExpirationWarning string                `xml:",omitempty"`
	Timestamp             time.Time             `xml:",omitempty"`
	Version               string                `xml:",omitempty"`
}
type GetSuggestedCategoriesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Query                string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetSuggestedCategoriesResponse struct {
	XMLName                xml.Name
	CategoryCount          int                          `xml:",omitempty"`
	SuggestedCategoryArray []SuggestedCategoryArrayType `xml:",omitempty"`
	Ack                    string                       `xml:",omitempty"`
	Build                  string                       `xml:",omitempty"`
	CorrelationID          string                       `xml:",omitempty"`
	Errors                 *ErrorType                   `xml:",omitempty"`
	HardExpirationWarning  string                       `xml:",omitempty"`
	Timestamp              time.Time                    `xml:",omitempty"`
	Version                string                       `xml:",omitempty"`
}
type GetTaxTableRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DetailLevel          string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetTaxTableResponse struct {
	XMLName               xml.Name
	LastUpdateTime        time.Time     `xml:",omitempty"`
	TaxTable              *TaxTableType `xml:",omitempty"`
	Ack                   string        `xml:",omitempty"`
	Build                 string        `xml:",omitempty"`
	CorrelationID         string        `xml:",omitempty"`
	Errors                *ErrorType    `xml:",omitempty"`
	HardExpirationWarning string        `xml:",omitempty"`
	Timestamp             time.Time     `xml:",omitempty"`
	Version               string        `xml:",omitempty"`
}
type GetTokenStatusRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetTokenStatusResponse struct {
	XMLName               xml.Name
	TokenStatus           *TokenStatusType `xml:",omitempty"`
	Ack                   string           `xml:",omitempty"`
	Build                 string           `xml:",omitempty"`
	CorrelationID         string           `xml:",omitempty"`
	Errors                *ErrorType       `xml:",omitempty"`
	HardExpirationWarning string           `xml:",omitempty"`
	Timestamp             time.Time        `xml:",omitempty"`
	Version               string           `xml:",omitempty"`
}
type GetUserRequest struct {
	XMLName                   xml.Name
	RequesterCredentials      *RequesterCredentialsType
	IncludeFeatureEligibility bool   `xml:",omitempty"`
	ItemID                    string `xml:",omitempty"`
	UserID                    string `xml:",omitempty"`
	DetailLevel               string `xml:",omitempty"`
	ErrorLanguage             string `xml:",omitempty"`
	MessageID                 string `xml:",omitempty"`
	Version                   string `xml:",omitempty"`
	WarningLevel              string `xml:",omitempty"`
}
type GetUserResponse struct {
	XMLName               xml.Name
	User                  *UserType  `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type GetUserContactDetailsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ContactID            string `xml:",omitempty"`
	ItemID               string `xml:",omitempty"`
	RequesterID          string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type GetUserContactDetailsResponse struct {
	XMLName               xml.Name
	ContactAddress        *AddressType `xml:",omitempty"`
	RegistrationDate      time.Time    `xml:",omitempty"`
	UserID                string       `xml:",omitempty"`
	Ack                   string       `xml:",omitempty"`
	Build                 string       `xml:",omitempty"`
	CorrelationID         string       `xml:",omitempty"`
	Errors                *ErrorType   `xml:",omitempty"`
	HardExpirationWarning string       `xml:",omitempty"`
	Timestamp             time.Time    `xml:",omitempty"`
	Version               string       `xml:",omitempty"`
}
type GetUserDisputesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DisputeFilterType    string          `xml:",omitempty"`
	DisputeSortType      string          `xml:",omitempty"`
	ModTimeFrom          time.Time       `xml:",omitempty"`
	ModTimeTo            time.Time       `xml:",omitempty"`
	Pagination           *PaginationType `xml:",omitempty"`
	DetailLevel          string          `xml:",omitempty"`
	ErrorLanguage        string          `xml:",omitempty"`
	MessageID            string          `xml:",omitempty"`
	Version              string          `xml:",omitempty"`
	WarningLevel         string          `xml:",omitempty"`
}
type GetUserDisputesResponse struct {
	XMLName               xml.Name
	DisputeArray          []DisputeArrayType      `xml:",omitempty"`
	DisputeFilterCount    *DisputeFilterCountType `xml:",omitempty"`
	EndingDisputeID       string                  `xml:",omitempty"`
	ItemsPerPage          int                     `xml:",omitempty"`
	PageNumber            int                     `xml:",omitempty"`
	PaginationResult      *PaginationResultType   `xml:",omitempty"`
	StartingDisputeID     string                  `xml:",omitempty"`
	Ack                   string                  `xml:",omitempty"`
	Build                 string                  `xml:",omitempty"`
	CorrelationID         string                  `xml:",omitempty"`
	Errors                *ErrorType              `xml:",omitempty"`
	HardExpirationWarning string                  `xml:",omitempty"`
	Timestamp             time.Time               `xml:",omitempty"`
	Version               string                  `xml:",omitempty"`
}
type GetUserPreferencesRequest struct {
	XMLName                                         xml.Name
	RequesterCredentials                            *RequesterCredentialsType
	ShowBidderNoticePreferences                     bool   `xml:",omitempty"`
	ShowCombinedPaymentPreferences                  bool   `xml:",omitempty"`
	ShowDispatchCutoffTimePreferences               bool   `xml:",omitempty"`
	ShoweBayPLUSPreference                          bool   `xml:",omitempty"`
	ShowEmailShipmentTrackingNumberPreference       bool   `xml:",omitempty"`
	ShowEndOfAuctionEmailPreferences                bool   `xml:",omitempty"`
	ShowGlobalShippingProgramListingPreference      bool   `xml:",omitempty"`
	ShowGlobalShippingProgramPreference             bool   `xml:",omitempty"`
	ShowOutOfStockControlPreference                 bool   `xml:",omitempty"`
	ShowOverrideGSPServiceWithIntlServicePreference bool   `xml:",omitempty"`
	ShowPickupDropoffPreferences                    bool   `xml:",omitempty"`
	ShowPurchaseReminderEmailPreferences            bool   `xml:",omitempty"`
	ShowRequiredShipPhoneNumberPreference           bool   `xml:",omitempty"`
	ShowSellerExcludeShipToLocationPreference       bool   `xml:",omitempty"`
	ShowSellerFavoriteItemPreferences               bool   `xml:",omitempty"`
	ShowSellerPaymentPreferences                    bool   `xml:",omitempty"`
	ShowSellerProfilePreferences                    bool   `xml:",omitempty"`
	ShowSellerReturnPreferences                     bool   `xml:",omitempty"`
	ShowUnpaidItemAssistanceExclusionList           bool   `xml:",omitempty"`
	ShowUnpaidItemAssistancePreference              bool   `xml:",omitempty"`
	ErrorLanguage                                   string `xml:",omitempty"`
	MessageID                                       string `xml:",omitempty"`
	Version                                         string `xml:",omitempty"`
	WarningLevel                                    string `xml:",omitempty"`
}
type GetUserPreferencesResponse struct {
	XMLName                                     xml.Name
	BidderNoticePreferences                     *BidderNoticePreferencesType                `xml:",omitempty"`
	CombinedPaymentPreferences                  *CombinedPaymentPreferencesType             `xml:",omitempty"`
	DispatchCutoffTimePreference                *DispatchCutoffTimePreferencesType          `xml:",omitempty"`
	eBayPLUSPreference                          *eBayPLUSPreferenceType                     `xml:",omitempty"`
	EmailShipmentTrackingNumberPreference       bool                                        `xml:",omitempty"`
	EndOfAuctionEmailPreferences                *EndOfAuctionEmailPreferencesType           `xml:",omitempty"`
	GlobalShippingProgramListingPreference      bool                                        `xml:",omitempty"`
	OfferGlobalShippingProgramPreference        bool                                        `xml:",omitempty"`
	OutOfStockControlPreference                 bool                                        `xml:",omitempty"`
	OverrideGSPServiceWithIntlServicePreference bool                                        `xml:",omitempty"`
	PickupDropoffSellerPreference               bool                                        `xml:",omitempty"`
	PurchaseReminderEmailPreferences            *PurchaseReminderEmailPreferencesType       `xml:",omitempty"`
	RequiredShipPhoneNumberPreference           bool                                        `xml:",omitempty"`
	SellerExcludeShipToLocationPreferences      *SellerExcludeShipToLocationPreferencesType `xml:",omitempty"`
	SellerFavoriteItemPreferences               *SellerFavoriteItemPreferencesType          `xml:",omitempty"`
	SellerPaymentPreferences                    *SellerPaymentPreferencesType               `xml:",omitempty"`
	SellerProfilePreferences                    *SellerProfilePreferencesType               `xml:",omitempty"`
	SellerReturnPreferences                     *SellerReturnPreferencesType                `xml:",omitempty"`
	SellerThirdPartyCheckoutDisabled            bool                                        `xml:",omitempty"`
	UnpaidItemAssistancePreferences             *UnpaidItemAssistancePreferencesType        `xml:",omitempty"`
	Ack                                         string                                      `xml:",omitempty"`
	Build                                       string                                      `xml:",omitempty"`
	CorrelationID                               string                                      `xml:",omitempty"`
	Errors                                      *ErrorType                                  `xml:",omitempty"`
	HardExpirationWarning                       string                                      `xml:",omitempty"`
	Timestamp                                   time.Time                                   `xml:",omitempty"`
	Version                                     string                                      `xml:",omitempty"`
}
type GetVeROReasonCodeDetailsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ReasonCodeID         float64 `xml:",omitempty"`
	ReturnAllSites       bool    `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type GetVeROReasonCodeDetailsResponse struct {
	XMLName               xml.Name
	VeROReasonCodeDetails *VeROReasonCodeDetailsType `xml:",omitempty"`
	Ack                   string                     `xml:",omitempty"`
	Build                 string                     `xml:",omitempty"`
	CorrelationID         string                     `xml:",omitempty"`
	Errors                *ErrorType                 `xml:",omitempty"`
	HardExpirationWarning string                     `xml:",omitempty"`
	Timestamp             time.Time                  `xml:",omitempty"`
	Version               string                     `xml:",omitempty"`
}
type GetVeROReportStatusRequest struct {
	XMLName                    xml.Name
	RequesterCredentials       *RequesterCredentialsType
	IncludeReportedItemDetails bool            `xml:",omitempty"`
	ItemID                     string          `xml:",omitempty"`
	Pagination                 *PaginationType `xml:",omitempty"`
	TimeFrom                   time.Time       `xml:",omitempty"`
	TimeTo                     time.Time       `xml:",omitempty"`
	VeROReportPacketID         float64         `xml:",omitempty"`
	ErrorLanguage              string          `xml:",omitempty"`
	MessageID                  string          `xml:",omitempty"`
	Version                    string          `xml:",omitempty"`
	WarningLevel               string          `xml:",omitempty"`
}
type GetVeROReportStatusResponse struct {
	XMLName                xml.Name
	HasMoreItems           bool                         `xml:",omitempty"`
	ItemsPerPage           int                          `xml:",omitempty"`
	PageNumber             int                          `xml:",omitempty"`
	PaginationResult       *PaginationResultType        `xml:",omitempty"`
	ReportedItemDetails    *VeROReportedItemDetailsType `xml:",omitempty"`
	VeROReportPacketID     float64                      `xml:",omitempty"`
	VeROReportPacketStatus string                       `xml:",omitempty"`
	Ack                    string                       `xml:",omitempty"`
	Build                  string                       `xml:",omitempty"`
	CorrelationID          string                       `xml:",omitempty"`
	Errors                 *ErrorType                   `xml:",omitempty"`
	HardExpirationWarning  string                       `xml:",omitempty"`
	Timestamp              time.Time                    `xml:",omitempty"`
	Version                string                       `xml:",omitempty"`
}
type LeaveFeedbackRequest struct {
	XMLName                     xml.Name
	RequesterCredentials        *RequesterCredentialsType
	CommentText                 string                      `xml:",omitempty"`
	CommentType                 string                      `xml:",omitempty"`
	ItemArrivedWithinEDDType    string                      `xml:",omitempty"`
	ItemDeliveredWithinEDD      bool                        `xml:",omitempty"`
	ItemID                      string                      `xml:",omitempty"`
	OrderLineItemID             string                      `xml:",omitempty"`
	SellerItemRatingDetailArray []ItemRatingDetailArrayType `xml:",omitempty"`
	TargetUser                  string                      `xml:",omitempty"`
	TransactionID               string                      `xml:",omitempty"`
	ErrorLanguage               string                      `xml:",omitempty"`
	MessageID                   string                      `xml:",omitempty"`
	Version                     string                      `xml:",omitempty"`
	WarningLevel                string                      `xml:",omitempty"`
}
type LeaveFeedbackResponse struct {
	XMLName               xml.Name
	FeedbackID            string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type MoveSellingManagerInventoryFolderRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	FolderID             float64 `xml:",omitempty"`
	NewParentFolderID    float64 `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type MoveSellingManagerInventoryFolderResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type PlaceOfferRequest struct {
	XMLName                  xml.Name
	RequesterCredentials     *RequesterCredentialsType
	AffiliateTrackingDetails *AffiliateTrackingDetailsType `xml:",omitempty"`
	BlockOnWarning           bool                          `xml:",omitempty"`
	ItemID                   string                        `xml:",omitempty"`
	Offer                    *OfferType                    `xml:",omitempty"`
	VariationSpecifics       []NameValueListArrayType      `xml:",omitempty"`
	BotBlock                 *BotBlockRequestType          `xml:",omitempty"`
	EndUserIP                string                        `xml:",omitempty"`
	ErrorLanguage            string                        `xml:",omitempty"`
	InvocationID             string                        `xml:",omitempty"`
	MessageID                string                        `xml:",omitempty"`
	Version                  string                        `xml:",omitempty"`
	WarningLevel             string                        `xml:",omitempty"`
}
type PlaceOfferResponse struct {
	XMLName                    xml.Name
	BestOffer                  *BestOfferType                  `xml:",omitempty"`
	OrderLineItemID            string                          `xml:",omitempty"`
	SellingStatus              *SellingStatusType              `xml:",omitempty"`
	TransactionID              string                          `xml:",omitempty"`
	Ack                        string                          `xml:",omitempty"`
	BotBlock                   *BotBlockResponseType           `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Message                    string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type RelistFixedPriceItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DeletedField         string    `xml:",omitempty"`
	Item                 *ItemType `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type RelistFixedPriceItemResponse struct {
	XMLName                    xml.Name
	Category2ID                string                      `xml:",omitempty"`
	CategoryID                 string                      `xml:",omitempty"`
	DiscountReason             string                      `xml:",omitempty"`
	EndTime                    time.Time                   `xml:",omitempty"`
	Fees                       *FeesType                   `xml:",omitempty"`
	ItemID                     string                      `xml:",omitempty"`
	ListingRecommendations     *ListingRecommendationsType `xml:",omitempty"`
	ProductSuggestions         *ProductSuggestionsType     `xml:",omitempty"`
	SKU                        string                      `xml:",omitempty"`
	StartTime                  time.Time                   `xml:",omitempty"`
	Ack                        string                      `xml:",omitempty"`
	Build                      string                      `xml:",omitempty"`
	CorrelationID              string                      `xml:",omitempty"`
	DuplicateInvocationDetails string                      `xml:",omitempty"`
	Errors                     *ErrorType                  `xml:",omitempty"`
	HardExpirationWarning      string                      `xml:",omitempty"`
	Message                    string                      `xml:",omitempty"`
	Timestamp                  time.Time                   `xml:",omitempty"`
	Version                    string                      `xml:",omitempty"`
}
type RelistItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DeletedField         string    `xml:",omitempty"`
	Item                 *ItemType `xml:",omitempty"`
	ErrorHandling        string    `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type RelistItemResponse struct {
	XMLName                    xml.Name
	Category2ID                string                          `xml:",omitempty"`
	CategoryID                 string                          `xml:",omitempty"`
	DiscountReason             string                          `xml:",omitempty"`
	EndTime                    time.Time                       `xml:",omitempty"`
	Fees                       *FeesType                       `xml:",omitempty"`
	ItemID                     string                          `xml:",omitempty"`
	ListingRecommendations     *ListingRecommendationsType     `xml:",omitempty"`
	ProductSuggestions         *ProductSuggestionsType         `xml:",omitempty"`
	StartTime                  time.Time                       `xml:",omitempty"`
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Message                    string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type RemoveFromWatchListRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ItemID               string            `xml:",omitempty"`
	RemoveAllItems       bool              `xml:",omitempty"`
	VariationKey         *VariationKeyType `xml:",omitempty"`
	ErrorLanguage        string            `xml:",omitempty"`
	MessageID            string            `xml:",omitempty"`
	Version              string            `xml:",omitempty"`
	WarningLevel         string            `xml:",omitempty"`
}
type RemoveFromWatchListResponse struct {
	XMLName               xml.Name
	WatchListCount        int        `xml:",omitempty"`
	WatchListMaximum      int        `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type RespondToBestOfferRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Action               string  `xml:",omitempty"`
	BestOfferID          string  `xml:",omitempty"`
	CounterOfferPrice    float64 `xml:",omitempty"`
	CounterOfferQuantity int     `xml:",omitempty"`
	ItemID               string  `xml:",omitempty"`
	SellerResponse       string  `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type RespondToBestOfferResponse struct {
	XMLName               xml.Name
	RespondToBestOffer    []BestOfferArrayType `xml:",omitempty"`
	Ack                   string               `xml:",omitempty"`
	Build                 string               `xml:",omitempty"`
	CorrelationID         string               `xml:",omitempty"`
	Errors                *ErrorType           `xml:",omitempty"`
	HardExpirationWarning string               `xml:",omitempty"`
	Timestamp             time.Time            `xml:",omitempty"`
	Version               string               `xml:",omitempty"`
}
type RespondToFeedbackRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	FeedbackID           string `xml:",omitempty"`
	ItemID               string `xml:",omitempty"`
	OrderLineItemID      string `xml:",omitempty"`
	ResponseText         string `xml:",omitempty"`
	ResponseType         string `xml:",omitempty"`
	TargetUserID         string `xml:",omitempty"`
	TransactionID        string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type RespondToFeedbackResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type ReviseCheckoutStatusRequest struct {
	XMLName                 xml.Name
	RequesterCredentials    *RequesterCredentialsType
	AdjustmentAmount        float64                  `xml:",omitempty"`
	AmountPaid              float64                  `xml:",omitempty"`
	BuyerID                 string                   `xml:",omitempty"`
	CheckoutStatus          string                   `xml:",omitempty"`
	CODCost                 float64                  `xml:",omitempty"`
	EncryptedID             string                   `xml:",omitempty"`
	ExternalTransaction     *ExternalTransactionType `xml:",omitempty"`
	InsuranceType           string                   `xml:",omitempty"`
	ItemID                  string                   `xml:",omitempty"`
	MultipleSellerPaymentID string                   `xml:",omitempty"`
	OrderID                 string                   `xml:",omitempty"`
	OrderLineItemID         string                   `xml:",omitempty"`
	PaymentMethodUsed       string                   `xml:",omitempty"`
	PaymentStatus           string                   `xml:",omitempty"`
	SalesTax                float64                  `xml:",omitempty"`
	ShippingCost            float64                  `xml:",omitempty"`
	ShippingIncludedInTax   bool                     `xml:",omitempty"`
	ShippingInsuranceCost   float64                  `xml:",omitempty"`
	ShippingService         string                   `xml:",omitempty"`
	TransactionID           string                   `xml:",omitempty"`
	ErrorLanguage           string                   `xml:",omitempty"`
	InvocationID            string                   `xml:",omitempty"`
	MessageID               string                   `xml:",omitempty"`
	Version                 string                   `xml:",omitempty"`
	WarningLevel            string                   `xml:",omitempty"`
}
type ReviseCheckoutStatusResponse struct {
	XMLName                    xml.Name
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type ReviseFixedPriceItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DeletedField         string    `xml:",omitempty"`
	Item                 *ItemType `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type ReviseFixedPriceItemResponse struct {
	XMLName                    xml.Name
	Category2ID                string                      `xml:",omitempty"`
	CategoryID                 string                      `xml:",omitempty"`
	DiscountReason             string                      `xml:",omitempty"`
	EndTime                    time.Time                   `xml:",omitempty"`
	Fees                       *FeesType                   `xml:",omitempty"`
	ItemID                     string                      `xml:",omitempty"`
	ListingRecommendations     *ListingRecommendationsType `xml:",omitempty"`
	ProductSuggestions         *ProductSuggestionsType     `xml:",omitempty"`
	SKU                        string                      `xml:",omitempty"`
	StartTime                  time.Time                   `xml:",omitempty"`
	Ack                        string                      `xml:",omitempty"`
	Build                      string                      `xml:",omitempty"`
	CorrelationID              string                      `xml:",omitempty"`
	DuplicateInvocationDetails string                      `xml:",omitempty"`
	Errors                     *ErrorType                  `xml:",omitempty"`
	HardExpirationWarning      string                      `xml:",omitempty"`
	Message                    string                      `xml:",omitempty"`
	Timestamp                  time.Time                   `xml:",omitempty"`
	Version                    string                      `xml:",omitempty"`
}
type ReviseInventoryStatusRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	InventoryStatus      *InventoryStatusType `xml:",omitempty"`
	ErrorLanguage        string               `xml:",omitempty"`
	MessageID            string               `xml:",omitempty"`
	Version              string               `xml:",omitempty"`
	WarningLevel         string               `xml:",omitempty"`
}
type ReviseInventoryStatusResponse struct {
	XMLName               xml.Name
	Fees                  *InventoryFeesType   `xml:",omitempty"`
	InventoryStatus       *InventoryStatusType `xml:",omitempty"`
	Ack                   string               `xml:",omitempty"`
	Build                 string               `xml:",omitempty"`
	CorrelationID         string               `xml:",omitempty"`
	Errors                *ErrorType           `xml:",omitempty"`
	HardExpirationWarning string               `xml:",omitempty"`
	Timestamp             time.Time            `xml:",omitempty"`
	Version               string               `xml:",omitempty"`
}
type ReviseItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DeletedField         string    `xml:",omitempty"`
	Item                 *ItemType `xml:",omitempty"`
	VerifyOnly           bool      `xml:",omitempty"`
	ErrorHandling        string    `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	InvocationID         string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type ReviseItemResponse struct {
	XMLName                    xml.Name
	Category2ID                string                          `xml:",omitempty"`
	CategoryID                 string                          `xml:",omitempty"`
	DiscountReason             string                          `xml:",omitempty"`
	EndTime                    time.Time                       `xml:",omitempty"`
	Fees                       *FeesType                       `xml:",omitempty"`
	ItemID                     string                          `xml:",omitempty"`
	ListingRecommendations     *ListingRecommendationsType     `xml:",omitempty"`
	ProductSuggestions         *ProductSuggestionsType         `xml:",omitempty"`
	StartTime                  time.Time                       `xml:",omitempty"`
	Ack                        string                          `xml:",omitempty"`
	Build                      string                          `xml:",omitempty"`
	CorrelationID              string                          `xml:",omitempty"`
	DuplicateInvocationDetails *DuplicateInvocationDetailsType `xml:",omitempty"`
	Errors                     *ErrorType                      `xml:",omitempty"`
	HardExpirationWarning      string                          `xml:",omitempty"`
	Message                    string                          `xml:",omitempty"`
	Timestamp                  time.Time                       `xml:",omitempty"`
	Version                    string                          `xml:",omitempty"`
}
type ReviseMyMessagesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Flagged              bool                           `xml:",omitempty"`
	FolderID             float64                        `xml:",omitempty"`
	MessageIDs           []MyMessagesMessageIDArrayType `xml:",omitempty"`
	Read                 bool                           `xml:",omitempty"`
	ErrorLanguage        string                         `xml:",omitempty"`
	MessageID            string                         `xml:",omitempty"`
	Version              string                         `xml:",omitempty"`
	WarningLevel         string                         `xml:",omitempty"`
}
type ReviseMyMessagesResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type ReviseMyMessagesFoldersRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	FolderID             float64 `xml:",omitempty"`
	FolderName           string  `xml:",omitempty"`
	Operation            string  `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type ReviseMyMessagesFoldersResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type ReviseSellingManagerInventoryFolderRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Folder               *SellingManagerFolderDetailsType `xml:",omitempty"`
	ErrorLanguage        string                           `xml:",omitempty"`
	MessageID            string                           `xml:",omitempty"`
	Version              string                           `xml:",omitempty"`
	WarningLevel         string                           `xml:",omitempty"`
}
type ReviseSellingManagerInventoryFolderResponse struct {
	XMLName               xml.Name
	Folder                *SellingManagerFolderDetailsType `xml:",omitempty"`
	Ack                   string                           `xml:",omitempty"`
	Build                 string                           `xml:",omitempty"`
	CorrelationID         string                           `xml:",omitempty"`
	Errors                *ErrorType                       `xml:",omitempty"`
	HardExpirationWarning string                           `xml:",omitempty"`
	Timestamp             time.Time                        `xml:",omitempty"`
	Version               string                           `xml:",omitempty"`
}
type ReviseSellingManagerProductRequest struct {
	XMLName                        xml.Name
	RequesterCredentials           *RequesterCredentialsType
	DeletedField                   string                              `xml:",omitempty"`
	SellingManagerFolderDetails    *SellingManagerFolderDetailsType    `xml:",omitempty"`
	SellingManagerProductDetails   *SellingManagerProductDetailsType   `xml:",omitempty"`
	SellingManagerProductSpecifics *SellingManagerProductSpecificsType `xml:",omitempty"`
	ErrorLanguage                  string                              `xml:",omitempty"`
	MessageID                      string                              `xml:",omitempty"`
	Version                        string                              `xml:",omitempty"`
	WarningLevel                   string                              `xml:",omitempty"`
}
type ReviseSellingManagerProductResponse struct {
	XMLName                      xml.Name
	SellingManagerProductDetails *SellingManagerProductDetailsType `xml:",omitempty"`
	Ack                          string                            `xml:",omitempty"`
	Build                        string                            `xml:",omitempty"`
	CorrelationID                string                            `xml:",omitempty"`
	Errors                       *ErrorType                        `xml:",omitempty"`
	HardExpirationWarning        string                            `xml:",omitempty"`
	Timestamp                    time.Time                         `xml:",omitempty"`
	Version                      string                            `xml:",omitempty"`
}
type ReviseSellingManagerSaleRecordRequest struct {
	XMLName                 xml.Name
	RequesterCredentials    *RequesterCredentialsType
	ItemID                  string                       `xml:",omitempty"`
	OrderID                 string                       `xml:",omitempty"`
	OrderLineItemID         string                       `xml:",omitempty"`
	SellingManagerSoldOrder *SellingManagerSoldOrderType `xml:",omitempty"`
	TransactionID           string                       `xml:",omitempty"`
	ErrorLanguage           string                       `xml:",omitempty"`
	MessageID               string                       `xml:",omitempty"`
	Version                 string                       `xml:",omitempty"`
	WarningLevel            string                       `xml:",omitempty"`
}
type ReviseSellingManagerSaleRecordResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type ReviseSellingManagerTemplateRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DeletedField         string    `xml:",omitempty"`
	Item                 *ItemType `xml:",omitempty"`
	SaleTemplateID       float64   `xml:",omitempty"`
	SaleTemplateName     string    `xml:",omitempty"`
	VerifyOnly           bool      `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type ReviseSellingManagerTemplateResponse struct {
	XMLName                      xml.Name
	Category2ID                  string                            `xml:",omitempty"`
	CategoryID                   string                            `xml:",omitempty"`
	Fees                         *FeesType                         `xml:",omitempty"`
	SaleTemplateID               float64                           `xml:",omitempty"`
	SaleTemplateName             string                            `xml:",omitempty"`
	SellingManagerProductDetails *SellingManagerProductDetailsType `xml:",omitempty"`
	VerifyOnly                   bool                              `xml:",omitempty"`
	Ack                          string                            `xml:",omitempty"`
	Build                        string                            `xml:",omitempty"`
	CorrelationID                string                            `xml:",omitempty"`
	Errors                       *ErrorType                        `xml:",omitempty"`
	HardExpirationWarning        string                            `xml:",omitempty"`
	Timestamp                    time.Time                         `xml:",omitempty"`
	Version                      string                            `xml:",omitempty"`
}
type RevokeTokenRequest struct {
	XMLName                 xml.Name
	RequesterCredentials    *RequesterCredentialsType
	UnsubscribeNotification bool   `xml:",omitempty"`
	ErrorLanguage           string `xml:",omitempty"`
	MessageID               string `xml:",omitempty"`
	Version                 string `xml:",omitempty"`
	WarningLevel            string `xml:",omitempty"`
}
type RevokeTokenResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SaveItemToSellingManagerTemplateRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ItemID               string  `xml:",omitempty"`
	ProductID            float64 `xml:",omitempty"`
	TemplateName         string  `xml:",omitempty"`
	ErrorLanguage        string  `xml:",omitempty"`
	MessageID            string  `xml:",omitempty"`
	Version              string  `xml:",omitempty"`
	WarningLevel         string  `xml:",omitempty"`
}
type SaveItemToSellingManagerTemplateResponse struct {
	XMLName               xml.Name
	TemplateID            float64    `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SellerReverseDisputeRequest struct {
	XMLName                 xml.Name
	RequesterCredentials    *RequesterCredentialsType
	DisputeID               string `xml:",omitempty"`
	DisputeResolutionReason string `xml:",omitempty"`
	ErrorLanguage           string `xml:",omitempty"`
	MessageID               string `xml:",omitempty"`
	Version                 string `xml:",omitempty"`
	WarningLevel            string `xml:",omitempty"`
}
type SellerReverseDisputeResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SendInvoiceRequest struct {
	XMLName                             xml.Name
	RequesterCredentials                *RequesterCredentialsType
	AdjustmentAmount                    float64                                  `xml:",omitempty"`
	CheckoutInstructions                string                                   `xml:",omitempty"`
	CODCost                             float64                                  `xml:",omitempty"`
	EmailCopyToSeller                   bool                                     `xml:",omitempty"`
	InsuranceFee                        float64                                  `xml:",omitempty"`
	InsuranceOption                     string                                   `xml:",omitempty"`
	InternationalShippingServiceOptions *InternationalShippingServiceOptionsType `xml:",omitempty"`
	ItemID                              string                                   `xml:",omitempty"`
	OrderID                             string                                   `xml:",omitempty"`
	OrderLineItemID                     string                                   `xml:",omitempty"`
	PaymentMethods                      string                                   `xml:",omitempty"`
	PayPalEmailAddress                  string                                   `xml:",omitempty"`
	SalesTax                            *SalesTaxType                            `xml:",omitempty"`
	ShippingServiceOptions              *ShippingServiceOptionsType              `xml:",omitempty"`
	SKU                                 string                                   `xml:",omitempty"`
	TransactionID                       string                                   `xml:",omitempty"`
	ErrorLanguage                       string                                   `xml:",omitempty"`
	MessageID                           string                                   `xml:",omitempty"`
	Version                             string                                   `xml:",omitempty"`
	WarningLevel                        string                                   `xml:",omitempty"`
}
type SendInvoiceResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetMessagePreferencesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ASQPreferences       *ASQPreferencesType `xml:",omitempty"`
	ErrorLanguage        string              `xml:",omitempty"`
	MessageID            string              `xml:",omitempty"`
	Version              string              `xml:",omitempty"`
	WarningLevel         string              `xml:",omitempty"`
}
type SetMessagePreferencesResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetNotificationPreferencesRequest struct {
	XMLName                        xml.Name
	RequesterCredentials           *RequesterCredentialsType
	ApplicationDeliveryPreferences *ApplicationDeliveryPreferencesType `xml:",omitempty"`
	DeliveryURLName                string                              `xml:",omitempty"`
	EventProperty                  *NotificationEventPropertyType      `xml:",omitempty"`
	UserData                       *NotificationUserDataType           `xml:",omitempty"`
	UserDeliveryPreferenceArray    []NotificationEnableArrayType       `xml:",omitempty"`
	ErrorLanguage                  string                              `xml:",omitempty"`
	MessageID                      string                              `xml:",omitempty"`
	Version                        string                              `xml:",omitempty"`
	WarningLevel                   string                              `xml:",omitempty"`
}
type SetNotificationPreferencesResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetPromotionalSaleRequest struct {
	XMLName                xml.Name
	RequesterCredentials   *RequesterCredentialsType
	Action                 string               `xml:",omitempty"`
	PromotionalSaleDetails *PromotionalSaleType `xml:",omitempty"`
	ErrorLanguage          string               `xml:",omitempty"`
	MessageID              string               `xml:",omitempty"`
	Version                string               `xml:",omitempty"`
	WarningLevel           string               `xml:",omitempty"`
}
type SetPromotionalSaleResponse struct {
	XMLName               xml.Name
	PromotionalSaleID     float64    `xml:",omitempty"`
	Status                string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetPromotionalSaleListingsRequest struct {
	XMLName                    xml.Name
	RequesterCredentials       *RequesterCredentialsType
	Action                     string            `xml:",omitempty"`
	AllAuctionItems            bool              `xml:",omitempty"`
	AllFixedPriceItems         bool              `xml:",omitempty"`
	CategoryID                 float64           `xml:",omitempty"`
	PromotionalSaleID          float64           `xml:",omitempty"`
	PromotionalSaleItemIDArray []ItemIDArrayType `xml:",omitempty"`
	StoreCategoryID            float64           `xml:",omitempty"`
	ErrorLanguage              string            `xml:",omitempty"`
	MessageID                  string            `xml:",omitempty"`
	Version                    string            `xml:",omitempty"`
	WarningLevel               string            `xml:",omitempty"`
}
type SetPromotionalSaleListingsResponse struct {
	XMLName               xml.Name
	Status                string     `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetSellingManagerFeedbackOptionsRequest struct {
	XMLName                     xml.Name
	RequesterCredentials        *RequesterCredentialsType
	AutomatedLeaveFeedbackEvent string                     `xml:",omitempty"`
	StoredComments              []FeedbackCommentArrayType `xml:",omitempty"`
	ErrorLanguage               string                     `xml:",omitempty"`
	MessageID                   string                     `xml:",omitempty"`
	Version                     string                     `xml:",omitempty"`
	WarningLevel                string                     `xml:",omitempty"`
}
type SetSellingManagerFeedbackOptionsResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetSellingManagerItemAutomationRuleRequest struct {
	XMLName                        xml.Name
	RequesterCredentials           *RequesterCredentialsType
	AutomatedRelistingRule         *SellingManagerAutoRelistType            `xml:",omitempty"`
	AutomatedSecondChanceOfferRule *SellingManagerAutoSecondChanceOfferType `xml:",omitempty"`
	ItemID                         string                                   `xml:",omitempty"`
	ErrorLanguage                  string                                   `xml:",omitempty"`
	MessageID                      string                                   `xml:",omitempty"`
	Version                        string                                   `xml:",omitempty"`
	WarningLevel                   string                                   `xml:",omitempty"`
}
type SetSellingManagerItemAutomationRuleResponse struct {
	XMLName                        xml.Name
	AutomatedListingRule           *SellingManagerAutoListType              `xml:",omitempty"`
	AutomatedRelistingRule         *SellingManagerAutoRelistType            `xml:",omitempty"`
	AutomatedSecondChanceOfferRule *SellingManagerAutoSecondChanceOfferType `xml:",omitempty"`
	Fees                           *FeesType                                `xml:",omitempty"`
	Ack                            string                                   `xml:",omitempty"`
	Build                          string                                   `xml:",omitempty"`
	CorrelationID                  string                                   `xml:",omitempty"`
	Errors                         *ErrorType                               `xml:",omitempty"`
	HardExpirationWarning          string                                   `xml:",omitempty"`
	Timestamp                      time.Time                                `xml:",omitempty"`
	Version                        string                                   `xml:",omitempty"`
}
type SetSellingManagerTemplateAutomationRuleRequest struct {
	XMLName                        xml.Name
	RequesterCredentials           *RequesterCredentialsType
	AutomatedListingRule           *SellingManagerAutoListType              `xml:",omitempty"`
	AutomatedRelistingRule         *SellingManagerAutoRelistType            `xml:",omitempty"`
	AutomatedSecondChanceOfferRule *SellingManagerAutoSecondChanceOfferType `xml:",omitempty"`
	SaleTemplateID                 float64                                  `xml:",omitempty"`
	ErrorLanguage                  string                                   `xml:",omitempty"`
	MessageID                      string                                   `xml:",omitempty"`
	Version                        string                                   `xml:",omitempty"`
	WarningLevel                   string                                   `xml:",omitempty"`
}
type SetSellingManagerTemplateAutomationRuleResponse struct {
	XMLName                        xml.Name
	AutomatedListingRule           *SellingManagerAutoListType              `xml:",omitempty"`
	AutomatedRelistingRule         *SellingManagerAutoRelistType            `xml:",omitempty"`
	AutomatedSecondChanceOfferRule *SellingManagerAutoSecondChanceOfferType `xml:",omitempty"`
	Fees                           *FeesType                                `xml:",omitempty"`
	Ack                            string                                   `xml:",omitempty"`
	Build                          string                                   `xml:",omitempty"`
	CorrelationID                  string                                   `xml:",omitempty"`
	Errors                         *ErrorType                               `xml:",omitempty"`
	HardExpirationWarning          string                                   `xml:",omitempty"`
	Timestamp                      time.Time                                `xml:",omitempty"`
	Version                        string                                   `xml:",omitempty"`
}
type SetShippingDiscountProfilesRequest struct {
	XMLName                            xml.Name
	RequesterCredentials               *RequesterCredentialsType
	CalculatedHandlingDiscount         *CalculatedHandlingDiscountType         `xml:",omitempty"`
	CalculatedShippingDiscount         *CalculatedShippingDiscountType         `xml:",omitempty"`
	CombinedDuration                   string                                  `xml:",omitempty"`
	CurrencyID                         string                                  `xml:",omitempty"`
	FlatShippingDiscount               *FlatShippingDiscountType               `xml:",omitempty"`
	InternationalShippingInsurance     *ShippingInsuranceType                  `xml:",omitempty"`
	ModifyActionCode                   string                                  `xml:",omitempty"`
	PromotionalShippingDiscountDetails *PromotionalShippingDiscountDetailsType `xml:",omitempty"`
	ShippingInsurance                  *ShippingInsuranceType                  `xml:",omitempty"`
	ErrorLanguage                      string                                  `xml:",omitempty"`
	MessageID                          string                                  `xml:",omitempty"`
	Version                            string                                  `xml:",omitempty"`
	WarningLevel                       string                                  `xml:",omitempty"`
}
type SetShippingDiscountProfilesResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetStoreRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Store                *StoreType `xml:",omitempty"`
	ErrorLanguage        string     `xml:",omitempty"`
	MessageID            string     `xml:",omitempty"`
	Version              string     `xml:",omitempty"`
	WarningLevel         string     `xml:",omitempty"`
}
type SetStoreResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetStoreCategoriesRequest struct {
	XMLName                     xml.Name
	RequesterCredentials        *RequesterCredentialsType
	Action                      string                         `xml:",omitempty"`
	DestinationParentCategoryID float64                        `xml:",omitempty"`
	ItemDestinationCategoryID   float64                        `xml:",omitempty"`
	StoreCategories             []StoreCustomCategoryArrayType `xml:",omitempty"`
	ErrorLanguage               string                         `xml:",omitempty"`
	MessageID                   string                         `xml:",omitempty"`
	Version                     string                         `xml:",omitempty"`
	WarningLevel                string                         `xml:",omitempty"`
}
type SetStoreCategoriesResponse struct {
	XMLName               xml.Name
	CustomCategory        string     `xml:",omitempty"`
	Status                string     `xml:",omitempty"`
	TaskID                float64    `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetStoreCustomPageRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	CustomPage           *StoreCustomPageType `xml:",omitempty"`
	ErrorLanguage        string               `xml:",omitempty"`
	MessageID            string               `xml:",omitempty"`
	Version              string               `xml:",omitempty"`
	WarningLevel         string               `xml:",omitempty"`
}
type SetStoreCustomPageResponse struct {
	XMLName               xml.Name
	CustomPage            *StoreCustomPageType `xml:",omitempty"`
	Ack                   string               `xml:",omitempty"`
	Build                 string               `xml:",omitempty"`
	CorrelationID         string               `xml:",omitempty"`
	Errors                *ErrorType           `xml:",omitempty"`
	HardExpirationWarning string               `xml:",omitempty"`
	Timestamp             time.Time            `xml:",omitempty"`
	Version               string               `xml:",omitempty"`
}
type SetStorePreferencesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	StorePreferences     *StorePreferencesType `xml:",omitempty"`
	ErrorLanguage        string                `xml:",omitempty"`
	MessageID            string                `xml:",omitempty"`
	Version              string                `xml:",omitempty"`
	WarningLevel         string                `xml:",omitempty"`
}
type SetStorePreferencesResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetTaxTableRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	TaxTable             *TaxTableType `xml:",omitempty"`
	ErrorLanguage        string        `xml:",omitempty"`
	MessageID            string        `xml:",omitempty"`
	Version              string        `xml:",omitempty"`
	WarningLevel         string        `xml:",omitempty"`
}
type SetTaxTableResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetUserNotesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Action               string                   `xml:",omitempty"`
	ItemID               string                   `xml:",omitempty"`
	NoteText             string                   `xml:",omitempty"`
	OrderLineItemID      string                   `xml:",omitempty"`
	SKU                  string                   `xml:",omitempty"`
	TransactionID        string                   `xml:",omitempty"`
	VariationSpecifics   []NameValueListArrayType `xml:",omitempty"`
	ErrorLanguage        string                   `xml:",omitempty"`
	MessageID            string                   `xml:",omitempty"`
	Version              string                   `xml:",omitempty"`
	WarningLevel         string                   `xml:",omitempty"`
}
type SetUserNotesResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type SetUserPreferencesRequest struct {
	XMLName                                xml.Name
	RequesterCredentials                   *RequesterCredentialsType
	BidderNoticePreferences                *BidderNoticePreferencesType          `xml:",omitempty"`
	CombinedPaymentPreferences             *CombinedPaymentPreferencesType       `xml:",omitempty"`
	DispatchCutoffTimePreference           *DispatchCutoffTimePreferencesType    `xml:",omitempty"`
	EmailShipmentTrackingNumberPreference  bool                                  `xml:",omitempty"`
	EndOfAuctionEmailPreferences           *EndOfAuctionEmailPreferencesType     `xml:",omitempty"`
	GlobalShippingProgramListingPreference bool                                  `xml:",omitempty"`
	OutOfStockControlPreference            bool                                  `xml:",omitempty"`
	OverrideGSPserviceWithIntlService      bool                                  `xml:",omitempty"`
	PurchaseReminderEmailPreferences       *PurchaseReminderEmailPreferencesType `xml:",omitempty"`
	RequiredShipPhoneNumberPreference      bool                                  `xml:",omitempty"`
	SellerFavoriteItemPreferences          *SellerFavoriteItemPreferencesType    `xml:",omitempty"`
	SellerPaymentPreferences               *SellerPaymentPreferencesType         `xml:",omitempty"`
	SellerThirdPartyCheckoutDisabled       bool                                  `xml:",omitempty"`
	UnpaidItemAssistancePreferences        *UnpaidItemAssistancePreferencesType  `xml:",omitempty"`
	ErrorLanguage                          string                                `xml:",omitempty"`
	MessageID                              string                                `xml:",omitempty"`
	Version                                string                                `xml:",omitempty"`
	WarningLevel                           string                                `xml:",omitempty"`
}
type SetUserPreferencesResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type UploadSiteHostedPicturesRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ExtensionInDays      int    `xml:",omitempty"`
	ExternalPictureURL   string `xml:",omitempty"`
	PictureData          string `xml:",omitempty"`
	PictureName          string `xml:",omitempty"`
	PictureSet           string `xml:",omitempty"`
	PictureSystemVersion int    `xml:",omitempty"`
	PictureUploadPolicy  string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type UploadSiteHostedPicturesResponse struct {
	XMLName                  xml.Name
	PictureSystemVersion     int                           `xml:",omitempty"`
	SiteHostedPictureDetails *SiteHostedPictureDetailsType `xml:",omitempty"`
	Ack                      string                        `xml:",omitempty"`
	Build                    string                        `xml:",omitempty"`
	CorrelationID            string                        `xml:",omitempty"`
	Errors                   *ErrorType                    `xml:",omitempty"`
	HardExpirationWarning    string                        `xml:",omitempty"`
	Timestamp                time.Time                     `xml:",omitempty"`
	Version                  string                        `xml:",omitempty"`
}
type ValidateChallengeInputRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ChallengeToken       string `xml:",omitempty"`
	KeepTokenValid       bool   `xml:",omitempty"`
	UserInput            string `xml:",omitempty"`
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}
type ValidateChallengeInputResponse struct {
	XMLName               xml.Name
	ValidToken            bool       `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type ValidateTestUserRegistrationRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	FeedbackScore        int       `xml:",omitempty"`
	RegistrationDate     time.Time `xml:",omitempty"`
	SubscribeSA          bool      `xml:",omitempty"`
	SubscribeSAPro       bool      `xml:",omitempty"`
	SubscribeSM          bool      `xml:",omitempty"`
	SubscribeSMPro       bool      `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type ValidateTestUserRegistrationResponse struct {
	XMLName               xml.Name
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type VerifyAddFixedPriceItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Item                 *ItemType `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type VerifyAddFixedPriceItemResponse struct {
	XMLName                xml.Name
	Category2ID            string                      `xml:",omitempty"`
	CategoryID             string                      `xml:",omitempty"`
	DiscountReason         string                      `xml:",omitempty"`
	ExpressListing         bool                        `xml:",omitempty"`
	Fees                   *FeesType                   `xml:",omitempty"`
	ItemID                 string                      `xml:",omitempty"`
	ListingRecommendations *ListingRecommendationsType `xml:",omitempty"`
	Ack                    string                      `xml:",omitempty"`
	Build                  string                      `xml:",omitempty"`
	CorrelationID          string                      `xml:",omitempty"`
	Errors                 *ErrorType                  `xml:",omitempty"`
	HardExpirationWarning  string                      `xml:",omitempty"`
	Message                string                      `xml:",omitempty"`
	Timestamp              time.Time                   `xml:",omitempty"`
	Version                string                      `xml:",omitempty"`
}
type VerifyAddItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	Item                 *ItemType `xml:",omitempty"`
	ErrorHandling        string    `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type VerifyAddItemResponse struct {
	XMLName                xml.Name
	Category2ID            string                      `xml:",omitempty"`
	CategoryID             string                      `xml:",omitempty"`
	DiscountReason         string                      `xml:",omitempty"`
	Fees                   *FeesType                   `xml:",omitempty"`
	ItemID                 string                      `xml:",omitempty"`
	ListingRecommendations *ListingRecommendationsType `xml:",omitempty"`
	ProductSuggestions     *ProductSuggestionsType     `xml:",omitempty"`
	Ack                    string                      `xml:",omitempty"`
	Build                  string                      `xml:",omitempty"`
	CorrelationID          string                      `xml:",omitempty"`
	Errors                 *ErrorType                  `xml:",omitempty"`
	HardExpirationWarning  string                      `xml:",omitempty"`
	Message                string                      `xml:",omitempty"`
	Timestamp              time.Time                   `xml:",omitempty"`
	Version                string                      `xml:",omitempty"`
}
type VerifyAddSecondChanceItemRequest struct {
	XMLName               xml.Name
	RequesterCredentials  *RequesterCredentialsType
	BuyItNowPrice         float64 `xml:",omitempty"`
	Duration              string  `xml:",omitempty"`
	ItemID                string  `xml:",omitempty"`
	RecipientBidderUserID string  `xml:",omitempty"`
	SellerMessage         string  `xml:",omitempty"`
	ErrorLanguage         string  `xml:",omitempty"`
	MessageID             string  `xml:",omitempty"`
	Version               string  `xml:",omitempty"`
	WarningLevel          string  `xml:",omitempty"`
}
type VerifyAddSecondChanceItemResponse struct {
	XMLName               xml.Name
	EndTime               time.Time  `xml:",omitempty"`
	StartTime             time.Time  `xml:",omitempty"`
	Ack                   string     `xml:",omitempty"`
	Build                 string     `xml:",omitempty"`
	CorrelationID         string     `xml:",omitempty"`
	Errors                *ErrorType `xml:",omitempty"`
	HardExpirationWarning string     `xml:",omitempty"`
	Timestamp             time.Time  `xml:",omitempty"`
	Version               string     `xml:",omitempty"`
}
type VerifyRelistItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	DeletedField         string    `xml:",omitempty"`
	Item                 *ItemType `xml:",omitempty"`
	ErrorLanguage        string    `xml:",omitempty"`
	MessageID            string    `xml:",omitempty"`
	Version              string    `xml:",omitempty"`
	WarningLevel         string    `xml:",omitempty"`
}
type VerifyRelistItemResponse struct {
	XMLName                xml.Name
	DiscountReason         string                      `xml:",omitempty"`
	EndTime                time.Time                   `xml:",omitempty"`
	Fees                   *FeesType                   `xml:",omitempty"`
	ItemID                 string                      `xml:",omitempty"`
	ListingRecommendations *ListingRecommendationsType `xml:",omitempty"`
	ProductSuggestions     *ProductSuggestionsType     `xml:",omitempty"`
	StartTime              time.Time                   `xml:",omitempty"`
	Ack                    string                      `xml:",omitempty"`
	Build                  string                      `xml:",omitempty"`
	CorrelationID          string                      `xml:",omitempty"`
	Errors                 *ErrorType                  `xml:",omitempty"`
	HardExpirationWarning  string                      `xml:",omitempty"`
	Timestamp              time.Time                   `xml:",omitempty"`
	Version                string                      `xml:",omitempty"`
}
type VeROReportItemsRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentialsType
	ReportItems          *VeROReportItemsType `xml:",omitempty"`
	RightsOwnerID        string               `xml:",omitempty"`
	ErrorLanguage        string               `xml:",omitempty"`
	MessageID            string               `xml:",omitempty"`
	Version              string               `xml:",omitempty"`
	WarningLevel         string               `xml:",omitempty"`
}
type VeROReportItemsResponse struct {
	XMLName                xml.Name
	VeROReportPacketID     float64    `xml:",omitempty"`
	VeROReportPacketStatus string     `xml:",omitempty"`
	Ack                    string     `xml:",omitempty"`
	Build                  string     `xml:",omitempty"`
	CorrelationID          string     `xml:",omitempty"`
	Errors                 *ErrorType `xml:",omitempty"`
	HardExpirationWarning  string     `xml:",omitempty"`
	Timestamp              time.Time  `xml:",omitempty"`
	Version                string     `xml:",omitempty"`
}

type ErrorType struct {
	XMLName             xml.Name
	ErrorClassification string              `xml:",omitempty"`
	ErrorCode           string              `xml:",omitempty"`
	ErrorParameters     *ErrorParameterType `xml:",omitempty"`
	LongMessage         string              `xml:",omitempty"`
	SeverityCode        string              `xml:",omitempty"`
	ShortMessage        string              `xml:",omitempty"`
}

type ErrorParameterType struct {
	XMLName xml.Name `xml:"ErrorParameter"`
	Value   string   `xml:",omitempty"`
}

type ItemType struct {
	XMLName                        xml.Name                             `xml:"Item"`
	ApplicationData                string                               `xml:",omitempty"`
	AutoPay                        bool                                 `xml:",omitempty"`
	BestOfferDetails               *BestOfferDetailsType                `xml:",omitempty"`
	BuyerRequirementDetails        *BuyerRequirementDetailsType         `xml:",omitempty"`
	BuyerResponsibleForShipping    bool                                 `xml:",omitempty"`
	BuyItNowPrice                  float64                              `xml:",omitempty"`
	CategoryBasedAttributesPrefill bool                                 `xml:",omitempty"`
	CategoryMappingAllowed         bool                                 `xml:",omitempty"`
	Charity                        *CharityType                         `xml:",omitempty"`
	ConditionDescription           string                               `xml:",omitempty"`
	ConditionID                    int                                  `xml:",omitempty"`
	Country                        string                               `xml:",omitempty"`
	CrossBorderTrade               string                               `xml:",omitempty"`
	Description                    string                               `xml:",omitempty"`
	DigitalGoodInfo                *DigitalGoodInfoType                 `xml:",omitempty"`
	DisableBuyerRequirements       bool                                 `xml:",omitempty"`
	DiscountPriceInfo              string                               `xml:",omitempty"`
	DispatchTimeMax                int                                  `xml:",omitempty"`
	eBayNowEligible                bool                                 `xml:",omitempty"`
	eBayPlus                       bool                                 `xml:",omitempty"`
	ExtendedSellerContactDetails   *ExtendedContactDetailsType          `xml:",omitempty"`
	GiftIcon                       int                                  `xml:",omitempty"`
	GiftServices                   string                               `xml:",omitempty"`
	HitCounter                     string                               `xml:",omitempty"`
	IncludeRecommendations         bool                                 `xml:",omitempty"`
	ItemCompatibilityList          *ItemCompatibilityListType           `xml:",omitempty"`
	ItemID                         string                               `xml:",omitempty"`
	ItemSpecifics                  []NameValueListArrayType             `xml:",omitempty"`
	ListingDesigner                *ListingDesignerType                 `xml:",omitempty"`
	ListingDetails                 *ListingDetailsType                  `xml:",omitempty"`
	ListingDuration                string                               `xml:",omitempty"`
	ListingEnhancement             string                               `xml:",omitempty"`
	ListingSubtype2                string                               `xml:",omitempty"`
	ListingType                    string                               `xml:",omitempty"`
	LiveAuction                    bool                                 `xml:",omitempty"`
	Location                       string                               `xml:",omitempty"`
	LotSize                        int                                  `xml:",omitempty"`
	MotorsGermanySearchable        bool                                 `xml:",omitempty"`
	PaymentDetails                 *PaymentDetailsType                  `xml:",omitempty"`
	PaymentMethods                 string                               `xml:",omitempty"`
	PayPalEmailAddress             string                               `xml:",omitempty"`
	PickupInStoreDetails           *PickupInStoreDetailsType            `xml:",omitempty"`
	PictureDetails                 *PictureDetailsType                  `xml:",omitempty"`
	PostalCode                     string                               `xml:",omitempty"`
	PrimaryCategory                *CategoryType                        `xml:",omitempty"`
	PrivateListing                 bool                                 `xml:",omitempty"`
	PrivateNotes                   string                               `xml:",omitempty"`
	ProductListingDetails          *ProductListingDetailsType           `xml:",omitempty"`
	Quantity                       int                                  `xml:",omitempty"`
	QuantityInfo                   string                               `xml:",omitempty"`
	QuantityRestrictionPerBuyer    *QuantityRestrictionPerBuyerInfoType `xml:",omitempty"`
	ReservePrice                   float64                              `xml:",omitempty"`
	ReturnPolicy                   *ReturnPolicyType                    `xml:",omitempty"`
	ScheduleTime                   *time.Time                           `xml:",omitempty"`
	// SecondaryCategory               *CategoryType                        `xml:",omitempty"`
	SellerContactDetails            *AddressType                         `xml:",omitempty"`
	ShippingDetails                 *ShippingDetailsType                 `xml:",omitempty"`
	ShippingPackageDetails          *ShipPackageDetailsType              `xml:",omitempty"`
	ShippingServiceCostOverrideList *ShippingServiceCostOverrideListType `xml:",omitempty"`
	ShippingTermsInDescription      bool                                 `xml:",omitempty"`
	ShipToLocations                 string                               `xml:",omitempty"`
	Site                            string                               `xml:",omitempty"`
	SKU                             string                               `xml:",omitempty"`
	StartPrice                      float64                              `xml:",omitempty"`
	Storefront                      *StorefrontType                      `xml:",omitempty"`
	SubTitle                        string                               `xml:",omitempty"`
	Title                           string                               `xml:",omitempty"`
	UseTaxTable                     bool                                 `xml:",omitempty"`
	UUID                            string                               `xml:",omitempty"`
	VATDetails                      *VATDetailsType                      `xml:",omitempty"`
}

type BestOfferDetailsType struct {
	XMLName          xml.Name `xml:"BestOfferDetails"`
	BestOfferEnabled bool     `xml:",omitempty"`
}

type BuyerRequirementDetailsType struct {
	XMLName                      xml.Name                          `xml:"BuyerRequirementDetails"`
	LinkedPayPalAccount          bool                              `xml:",omitempty"`
	MaximumBuyerPolicyViolations *MaximumBuyerPolicyViolationsType `xml:",omitempty"`
	MaximumItemRequirements      *MaximumItemRequirementsType      `xml:",omitempty"`
	MaximumUnpaidItemStrikesInfo *MaximumUnpaidItemStrikesInfoType `xml:",omitempty"`
	MinimumFeedbackScore         int                               `xml:",omitempty"`
	ShipToRegistrationCountry    bool                              `xml:",omitempty"`
	VerifiedUserRequirements     *VerifiedUserRequirementsType     `xml:",omitempty"`
	ZeroFeedbackScore            bool                              `xml:",omitempty"`
}

type MaximumBuyerPolicyViolationsType struct {
	XMLName xml.Name `xml:"MaximumBuyerPolicyViolations"`
	Count   int      `xml:",omitempty"`
	Period  string   `xml:",omitempty"`
}

type MaximumItemRequirementsType struct {
	XMLName              xml.Name `xml:"MaximumItemRequirements"`
	MaximumItemCount     int      `xml:",omitempty"`
	MinimumFeedbackScore int      `xml:",omitempty"`
}

type MaximumUnpaidItemStrikesInfoType struct {
	XMLName xml.Name `xml:"MaximumUnpaidItemStrikesInfo"`
	Count   int      `xml:",omitempty"`
	Period  string   `xml:",omitempty"`
}

type VerifiedUserRequirementsType struct {
	XMLName              xml.Name `xml:"VerifiedUserRequirements"`
	MinimumFeedbackScore int      `xml:",omitempty"`
	VerifiedUser         bool     `xml:",omitempty"`
}

type CharityType struct {
	XMLName         xml.Name `xml:"Charity"`
	CharityID       string   `xml:",omitempty"`
	CharityNumber   int      `xml:",omitempty"`
	DonationPercent float64  `xml:",omitempty"`
}

type DigitalGoodInfoType struct {
	XMLName         xml.Name `xml:"DigitalGoodInfo"`
	DigitalDelivery bool     `xml:",omitempty"`
}

type DiscountPriceInfoType struct {
	XMLName                        xml.Name `xml:"DiscountPriceInfo"`
	MadeForOutletComparisonPrice   float64  `xml:",omitempty"`
	MinimumAdvertisedPrice         float64  `xml:",omitempty"`
	MinimumAdvertisedPriceExposure string   `xml:",omitempty"`
	OriginalRetailPrice            float64  `xml:",omitempty"`
	SoldOffeBay                    bool     `xml:",omitempty"`
	SoldOneBay                     bool     `xml:",omitempty"`
}

type ItemCompatibilityListType struct {
	XMLName       xml.Name               `xml:"ItemCompatibilityList"`
	Compatibility *ItemCompatibilityType `xml:",omitempty"`
}

type ItemCompatibilityType struct {
	XMLName            xml.Name           `xml:"ItemCompatibility"`
	CompatibilityNotes string             `xml:",omitempty"`
	NameValueList      *NameValueListType `xml:",omitempty"`
}

type NameValueListType struct {
	XMLName xml.Name `xml:"NameValueList"`
	Name    string   `xml:",omitempty"`
	Value   string   `xml:",omitempty"`
}

type NameValueListArrayType struct {
	XMLName       xml.Name           `xml:"NameValueListArray"`
	NameValueList *NameValueListType `xml:",omitempty"`
}

type ListingDesignerType struct {
	XMLName            xml.Name `xml:"ListingDesigner"`
	LayoutID           int      `xml:",omitempty"`
	OptimalPictureSize bool     `xml:",omitempty"`
	ThemeID            int      `xml:",omitempty"`
}

type ListingDetailsType struct {
	XMLName                  xml.Name `xml:"ListingDetails"`
	BestOfferAutoAcceptPrice float64  `xml:",omitempty"`
	MinimumBestOfferPrice    float64  `xml:",omitempty"`
}

type PickupInStoreDetailsType struct {
	XMLName                  xml.Name `xml:"PickupInStoreDetails"`
	EligibleForPickupDropOff bool     `xml:",omitempty"`
	EligibleForPickupInStore bool     `xml:",omitempty"`
}

type PictureDetailsType struct {
	XMLName         xml.Name `xml:"PictureDetails"`
	GalleryDuration string   `xml:",omitempty"`
	GalleryType     string   `xml:",omitempty"`
	PhotoDisplay    string   `xml:",omitempty"`
	PictureURL      string   `xml:",omitempty"`
}

type CategoryType struct {
	XMLName    xml.Name `xml:"Category"`
	CategoryID string   `xml:",omitempty"`
}

type ProductListingDetailsType struct {
	XMLName                        xml.Name                  `xml:"ProductListingDetails"`
	BrandMPN                       *BrandMPNType             `xml:",omitempty"`
	EAN                            string                    `xml:",omitempty"`
	IncludeeBayProductDetails      bool                      `xml:",omitempty"`
	IncludeStockPhotoURL           bool                      `xml:",omitempty"`
	ISBN                           string                    `xml:",omitempty"`
	NameValueList                  *NameValueListType        `xml:",omitempty"`
	ProductReferenceID             string                    `xml:",omitempty"`
	ReturnSearchResultOnDuplicates bool                      `xml:",omitempty"`
	TicketListingDetails           *TicketListingDetailsType `xml:",omitempty"`
	UPC                            string                    `xml:",omitempty"`
	UseFirstProduct                bool                      `xml:",omitempty"`
	UseStockPhotoURLAsGallery      bool                      `xml:",omitempty"`
}

type BrandMPNType struct {
	XMLName xml.Name `xml:"BrandMPN"`
	Brand   string   `xml:",omitempty"`
	MPN     string   `xml:",omitempty"`
}

type TicketListingDetailsType struct {
	XMLName     xml.Name `xml:"TicketListingDetails"`
	EventTitle  string   `xml:",omitempty"`
	PrintedDate string   `xml:",omitempty"`
	PrintedTime string   `xml:",omitempty"`
	Venue       string   `xml:",omitempty"`
}

type QuantityInfoType struct {
	XMLName           xml.Name `xml:"QuantityInfo"`
	MinimumRemnantSet int      `xml:",omitempty"`
}

type QuantityRestrictionPerBuyerInfoType struct {
	XMLName         xml.Name `xml:"QuantityRestrictionPerBuyerInfo"`
	MaximumQuantity int      `xml:",omitempty"`
}

type ReturnPolicyType struct {
	XMLName                  xml.Name `xml:"ReturnPolicy"`
	Description              string   `xml:",omitempty"`
	ExtendedHolidayReturns   bool     `xml:",omitempty"`
	RefundOption             string   `xml:",omitempty"`
	RestockingFeeValueOption string   `xml:",omitempty"`
	ReturnsAcceptedOption    string   `xml:",omitempty"`
	ReturnsWithinOption      string   `xml:",omitempty"`
	ShippingCostPaidByOption string   `xml:",omitempty"`
	WarrantyDurationOption   string   `xml:",omitempty"`
	WarrantyOfferedOption    string   `xml:",omitempty"`
	WarrantyTypeOption       string   `xml:",omitempty"`
}

type SellerProfilesType struct {
	XMLName               xml.Name                   `xml:"SellerProfiles"`
	SellerPaymentProfile  *SellerPaymentProfileType  `xml:",omitempty"`
	SellerReturnProfile   *SellerReturnProfileType   `xml:",omitempty"`
	SellerShippingProfile *SellerShippingProfileType `xml:",omitempty"`
}

type SellerPaymentProfileType struct {
	XMLName            xml.Name `xml:"SellerPaymentProfile"`
	PaymentProfileID   float64  `xml:",omitempty"`
	PaymentProfileName string   `xml:",omitempty"`
}

type SellerReturnProfileType struct {
	XMLName           xml.Name `xml:"SellerReturnProfile"`
	ReturnProfileID   float64  `xml:",omitempty"`
	ReturnProfileName string   `xml:",omitempty"`
}

type SellerShippingProfileType struct {
	XMLName             xml.Name `xml:"SellerShippingProfile"`
	ShippingProfileID   float64  `xml:",omitempty"`
	ShippingProfileName string   `xml:",omitempty"`
}

type ShippingDetailsType struct {
	XMLName                                  xml.Name                                 `xml:"ShippingDetails"`
	CalculatedShippingRate                   *CalculatedShippingRateType              `xml:",omitempty"`
	CODCost                                  float64                                  `xml:",omitempty"`
	ExcludeShipToLocation                    string                                   `xml:",omitempty"`
	GlobalShipping                           bool                                     `xml:",omitempty"`
	InsuranceDetails                         *InsuranceDetailsType                    `xml:",omitempty"`
	InternationalInsuranceDetails            *InsuranceDetailsType                    `xml:",omitempty"`
	InternationalPromotionalShippingDiscount bool                                     `xml:",omitempty"`
	InternationalShippingDiscountProfileID   string                                   `xml:",omitempty"`
	InternationalShippingServiceOption       *InternationalShippingServiceOptionsType `xml:",omitempty"`
	PaymentInstructions                      string                                   `xml:",omitempty"`
	PromotionalShippingDiscount              bool                                     `xml:",omitempty"`
	RateTableDetails                         *RateTableDetailsType                    `xml:",omitempty"`
	SalesTax                                 *SalesTaxType                            `xml:",omitempty"`
	ShippingDiscountProfileID                string                                   `xml:",omitempty"`
	ShippingServiceOptions                   *ShippingServiceOptionsType              `xml:",omitempty"`
	ShippingType                             string                                   `xml:",omitempty"`
}

type CalculatedShippingRateType struct {
	XMLName                             xml.Name `xml:"CalculatedShippingRate"`
	InternationalPackagingHandlingCosts float64  `xml:",omitempty"`
	MeasurementUnit                     string   `xml:",omitempty"`
	OriginatingPostalCode               string   `xml:",omitempty"`
	PackagingHandlingCosts              float64  `xml:",omitempty"`
	ShippingIrregular                   bool     `xml:",omitempty"`
}

type InsuranceDetailsType struct {
	XMLName         xml.Name `xml:"InsuranceDetails"`
	InsuranceFee    float64  `xml:",omitempty"`
	InsuranceOption string   `xml:",omitempty"`
}

type InternationalShippingServiceOptionsType struct {
	XMLName                       xml.Name `xml:"InternationalShippingServiceOptions"`
	ShippingService               string   `xml:",omitempty"`
	ShippingServiceAdditionalCost float64  `xml:",omitempty"`
	ShippingServiceCost           float64  `xml:",omitempty"`
	ShippingServicePriority       int      `xml:",omitempty"`
	ShipToLocation                string   `xml:",omitempty"`
}

type RateTableDetailsType struct {
	XMLName                  xml.Name `xml:"RateTableDetails"`
	DomesticRateTable        string   `xml:",omitempty"`
	DomesticRateTableId      string   `xml:",omitempty"`
	InternationalRateTable   string   `xml:",omitempty"`
	InternationalRateTableId string   `xml:",omitempty"`
}

type SalesTaxType struct {
	XMLName               xml.Name `xml:"SalesTax"`
	SalesTaxPercent       float64  `xml:",omitempty"`
	SalesTaxState         string   `xml:",omitempty"`
	ShippingIncludedInTax bool     `xml:",omitempty"`
}

type ShippingServiceOptionsType struct {
	XMLName                       xml.Name `xml:"ShippingServiceOptions"`
	FreeShipping                  bool     `xml:",omitempty"`
	ShippingService               string   `xml:",omitempty"`
	ShippingServiceAdditionalCost float64  `xml:",omitempty"`
	ShippingServiceCost           float64  `xml:",omitempty"`
	ShippingServicePriority       int      `xml:",omitempty"`
	ShippingSurcharge             float64  `xml:",omitempty"`
}

type ShipPackageDetailsType struct {
	XMLName           xml.Name `xml:"ShipPackageDetails"`
	MeasurementUnit   string   `xml:",omitempty"`
	PackageDepth      float64  `xml:",omitempty"`
	PackageLength     float64  `xml:",omitempty"`
	PackageWidth      float64  `xml:",omitempty"`
	ShippingIrregular bool     `xml:",omitempty"`
	ShippingPackage   string   `xml:",omitempty"`
	WeightMajor       float64  `xml:",omitempty"`
	WeightMinor       float64  `xml:",omitempty"`
}

type ShippingServiceCostOverrideListType struct {
	XMLName                     xml.Name                         `xml:"ShippingServiceCostOverrideList"`
	ShippingServiceCostOverride *ShippingServiceCostOverrideType `xml:",omitempty"`
}

type ShippingServiceCostOverrideType struct {
	XMLName                       xml.Name `xml:"ShippingServiceCostOverride"`
	ShippingServiceAdditionalCost float64  `xml:",omitempty"`
	ShippingServiceCost           float64  `xml:",omitempty"`
	ShippingServicePriority       int      `xml:",omitempty"`
	ShippingServiceType           string   `xml:",omitempty"`
	ShippingSurcharge             float64  `xml:",omitempty"`
}

type StorefrontType struct {
	XMLName            xml.Name `xml:"Storefront"`
	StoreCategory2ID   float64  `xml:",omitempty"`
	StoreCategory2Name string   `xml:",omitempty"`
	StoreCategoryID    float64  `xml:",omitempty"`
	StoreCategoryName  string   `xml:",omitempty"`
}

type VariationsType struct {
	XMLName               xml.Name                 `xml:"Variations"`
	Pictures              *PicturesType            `xml:",omitempty"`
	Variation             *VariationType           `xml:",omitempty"`
	VariationSpecificsSet []NameValueListArrayType `xml:",omitempty"`
}

type PicturesType struct {
	XMLName                     xml.Name                         `xml:"Pictures"`
	VariationSpecificName       string                           `xml:",omitempty"`
	VariationSpecificPictureSet *VariationSpecificPictureSetType `xml:",omitempty"`
}

type VariationSpecificPictureSetType struct {
	XMLName                xml.Name `xml:"VariationSpecificPictureSet"`
	PictureURL             string   `xml:",omitempty"`
	VariationSpecificValue string   `xml:",omitempty"`
}

type VariationType struct {
	XMLName                        xml.Name                 `xml:"Variation"`
	DiscountPriceInfo              *DiscountPriceInfoType   `xml:",omitempty"`
	Quantity                       int                      `xml:",omitempty"`
	SKU                            string                   `xml:",omitempty"`
	StartPrice                     float64                  `xml:",omitempty"`
	VariationProductListingDetails string                   `xml:",omitempty"`
	VariationSpecifics             []NameValueListArrayType `xml:",omitempty"`
}

type VariationProductListingDetailsType struct {
	XMLName       xml.Name           `xml:"VariationProductListingDetails"`
	EAN           string             `xml:",omitempty"`
	ISBN          string             `xml:",omitempty"`
	NameValueList *NameValueListType `xml:",omitempty"`
	UPC           string             `xml:",omitempty"`
}

type VATDetailsType struct {
	XMLName              xml.Name `xml:"VATDetails"`
	BusinessSeller       bool     `xml:",omitempty"`
	RestrictedToBusiness bool     `xml:",omitempty"`
	VATPercent           float64  `xml:",omitempty"`
}

type FeesType struct {
	XMLName xml.Name `xml:"Fees"`
	Fee     *FeeType `xml:",omitempty"`
}

type FeeType struct {
	XMLName             xml.Name `xml:"Fee"`
	Fee                 float64  `xml:",omitempty"`
	Name                string   `xml:",omitempty"`
	PromotionalDiscount float64  `xml:",omitempty"`
}

type ListingRecommendationsType struct {
	XMLName        xml.Name                   `xml:"ListingRecommendations"`
	Recommendation *ListingRecommendationType `xml:",omitempty"`
}

type ListingRecommendationType struct {
	XMLName   xml.Name      `xml:"ListingRecommendation"`
	Code      string        `xml:",omitempty"`
	FieldName string        `xml:",omitempty"`
	Group     string        `xml:",omitempty"`
	Message   string        `xml:",omitempty"`
	Metadata  *MetadataType `xml:",omitempty"`
	Type      string        `xml:",omitempty"`
	Value     string        `xml:",omitempty"`
}

type MetadataType struct {
	XMLName xml.Name `xml:"Metadata"`
	Name    string   `xml:",omitempty"`
	Value   string   `xml:",omitempty"`
}

type ProductSuggestionsType struct {
	XMLName           xml.Name               `xml:"ProductSuggestions"`
	ProductSuggestion *ProductSuggestionType `xml:",omitempty"`
}

type ProductSuggestionType struct {
	XMLName     xml.Name `xml:"ProductSuggestion"`
	EPID        string   `xml:",omitempty"`
	Recommended bool     `xml:",omitempty"`
	StockPhoto  string   `xml:",omitempty"`
	Title       string   `xml:",omitempty"`
}

type DuplicateInvocationDetailsType struct {
	XMLName               xml.Name `xml:"DuplicateInvocationDetails"`
	DuplicateInvocationID string   `xml:",omitempty"`
	InvocationTrackingID  string   `xml:",omitempty"`
	Status                string   `xml:",omitempty"`
}

type ExtendedContactDetailsType struct {
	XMLName                           xml.Name                 `xml:"ExtendedContactDetails"`
	ClassifiedAdContactByEmailEnabled bool                     `xml:",omitempty"`
	ContactHoursDetails               *ContactHoursDetailsType `xml:",omitempty"`
}

type ContactHoursDetailsType struct {
	XMLName       xml.Name  `xml:"ContactHoursDetails"`
	Hours1AnyTime bool      `xml:",omitempty"`
	Hours1Days    string    `xml:",omitempty"`
	Hours1From    time.Time `xml:",omitempty"`
	Hours1To      time.Time `xml:",omitempty"`
	Hours2AnyTime bool      `xml:",omitempty"`
	Hours2Days    string    `xml:",omitempty"`
	Hours2From    time.Time `xml:",omitempty"`
	Hours2To      time.Time `xml:",omitempty"`
	TimeZoneID    string    `xml:",omitempty"`
}

type PaymentDetailsType struct {
	XMLName           xml.Name `xml:"PaymentDetails"`
	DaysToFullPayment int      `xml:",omitempty"`
	DepositAmount     float64  `xml:",omitempty"`
	DepositType       string   `xml:",omitempty"`
	HoursToDeposit    int      `xml:",omitempty"`
}

type UserType struct {
	XMLName      xml.Name `xml:"User"`
	MotorsDealer bool     `xml:",omitempty"`
}

type AddressType struct {
	XMLName              xml.Name `xml:"Address"`
	CompanyName          string   `xml:",omitempty"`
	County               string   `xml:",omitempty"`
	Phone2AreaOrCityCode string   `xml:",omitempty"`
	Phone2CountryCode    string   `xml:",omitempty"`
	Phone2LocalNumber    string   `xml:",omitempty"`
	PhoneAreaOrCityCode  string   `xml:",omitempty"`
	PhoneCountryCode     string   `xml:",omitempty"`
	PhoneLocalNumber     string   `xml:",omitempty"`
	Street               string   `xml:",omitempty"`
	Street2              string   `xml:",omitempty"`
}

type AddItemRequestContainerType struct {
	XMLName   xml.Name  `xml:"AddItemRequestContainer"`
	Item      *ItemType `xml:",omitempty"`
	MessageID string    `xml:",omitempty"`
}

type AddItemResponseContainerType struct {
	XMLName                xml.Name                    `xml:"AddItemResponseContainer"`
	Category2ID            string                      `xml:",omitempty"`
	CategoryID             string                      `xml:",omitempty"`
	CorrelationID          string                      `xml:",omitempty"`
	DiscountReason         string                      `xml:",omitempty"`
	EndTime                time.Time                   `xml:",omitempty"`
	Errors                 *ErrorType                  `xml:",omitempty"`
	Fees                   *FeesType                   `xml:",omitempty"`
	ItemID                 string                      `xml:",omitempty"`
	ListingRecommendations *ListingRecommendationsType `xml:",omitempty"`
	Message                string                      `xml:",omitempty"`
	StartTime              time.Time                   `xml:",omitempty"`
}

type MemberMessageType struct {
	XMLName         xml.Name          `xml:"MemberMessage"`
	Body            string            `xml:",omitempty"`
	DisplayToPublic bool              `xml:",omitempty"`
	MessageID       string            `xml:",omitempty"`
	MessageMedia    *MessageMediaType `xml:",omitempty"`
	MessageType     string            `xml:",omitempty"`
	QuestionType    string            `xml:",omitempty"`
	RecipientID     string            `xml:",omitempty"`
	SenderEmail     string            `xml:",omitempty"`
	SenderID        string            `xml:",omitempty"`
	Subject         string            `xml:",omitempty"`
}

type MessageMediaType struct {
	XMLName   xml.Name `xml:"MessageMedia"`
	MediaName string   `xml:",omitempty"`
	MediaURL  string   `xml:",omitempty"`
}

type AddMemberMessagesAAQToBidderRequestContainerType struct {
	XMLName       xml.Name           `xml:"AddMemberMessagesAAQToBidderRequestContainer"`
	CorrelationID string             `xml:",omitempty"`
	ItemID        string             `xml:",omitempty"`
	MemberMessage *MemberMessageType `xml:",omitempty"`
}

type AddMemberMessagesAAQToBidderResponseContainerType struct {
	XMLName       xml.Name `xml:"AddMemberMessagesAAQToBidderResponseContainer"`
	Ack           string   `xml:",omitempty"`
	CorrelationID string   `xml:",omitempty"`
}

type OrderType struct {
	XMLName                             xml.Name                     `xml:"Order"`
	BuyerPackageEnclosures              *BuyerPackageEnclosuresType  `xml:",omitempty"`
	BuyerTaxIdentifier                  *TaxIdentifierType           `xml:",omitempty"`
	CancelDetail                        *CancelDetailType            `xml:",omitempty"`
	CancelReason                        string                       `xml:",omitempty"`
	CancelReasonDetails                 string                       `xml:",omitempty"`
	CancelStatus                        string                       `xml:",omitempty"`
	ContainseBayPlusTransaction         bool                         `xml:",omitempty"`
	CreatingUserRole                    string                       `xml:",omitempty"`
	ExtendedOrderID                     string                       `xml:",omitempty"`
	IntegratedMerchantCreditCardEnabled bool                         `xml:",omitempty"`
	IsMultiLegShipping                  bool                         `xml:",omitempty"`
	LogisticsPlanType                   string                       `xml:",omitempty"`
	MonetaryDetails                     *PaymentsInformationType     `xml:",omitempty"`
	MultiLegShippingDetails             *MultiLegShippingDetailsType `xml:",omitempty"`
	OrderID                             string                       `xml:",omitempty"`
	OrderStatus                         string                       `xml:",omitempty"`
	PaymentHoldDetails                  *PaymentHoldDetailType       `xml:",omitempty"`
	PaymentHoldStatus                   string                       `xml:",omitempty"`
	ShippingConvenienceCharge           float64                      `xml:",omitempty"`
}

type TransactionArrayType struct {
	XMLName     xml.Name         `xml:"TransactionArray"`
	Transaction *TransactionType `xml:",omitempty"`
}

type TransactionType struct {
	XMLName                           xml.Name                          `xml:"Transaction"`
	ActualHandlingCost                float64                           `xml:",omitempty"`
	ActualShippingCost                float64                           `xml:",omitempty"`
	AdjustmentAmount                  float64                           `xml:",omitempty"`
	AmountPaid                        float64                           `xml:",omitempty"`
	BestOfferSale                     bool                              `xml:",omitempty"`
	Buyer                             *UserType                         `xml:",omitempty"`
	BuyerCheckoutMessage              string                            `xml:",omitempty"`
	BuyerGuaranteePrice               float64                           `xml:",omitempty"`
	BuyerPackageEnclosures            *BuyerPackageEnclosuresType       `xml:",omitempty"`
	CartID                            string                            `xml:",omitempty"`
	CodiceFiscale                     string                            `xml:",omitempty"`
	ContainingOrder                   *OrderType                        `xml:",omitempty"`
	ConvertedAdjustmentAmount         float64                           `xml:",omitempty"`
	ConvertedAmountPaid               float64                           `xml:",omitempty"`
	ConvertedTransactionPrice         float64                           `xml:",omitempty"`
	CreatedDate                       time.Time                         `xml:",omitempty"`
	DepositType                       string                            `xml:",omitempty"`
	DigitalDeliverySelected           *DigitalDeliverySelectedType      `xml:",omitempty"`
	eBayPaymentID                     string                            `xml:",omitempty"`
	eBayPlusTransaction               bool                              `xml:",omitempty"`
	ExtendedOrderID                   string                            `xml:",omitempty"`
	ExternalTransaction               *ExternalTransactionType          `xml:",omitempty"`
	FinalValueFee                     float64                           `xml:",omitempty"`
	Gift                              bool                              `xml:",omitempty"`
	GiftSummary                       *GiftSummaryType                  `xml:",omitempty"`
	GuaranteedDelivery                bool                              `xml:",omitempty"`
	GuaranteedShipping                bool                              `xml:",omitempty"`
	IntangibleItem                    bool                              `xml:",omitempty"`
	InventoryReservationID            string                            `xml:",omitempty"`
	InvoiceSentTime                   time.Time                         `xml:",omitempty"`
	IsMultiLegShipping                bool                              `xml:",omitempty"`
	Item                              *ItemType                         `xml:",omitempty"`
	ListingCheckoutRedirectPreference string                            `xml:",omitempty"`
	LogisticsPlanType                 string                            `xml:",omitempty"`
	MonetaryDetails                   *PaymentsInformationType          `xml:",omitempty"`
	MultiLegShippingDetails           *MultiLegShippingDetailsType      `xml:",omitempty"`
	OrderLineItemID                   string                            `xml:",omitempty"`
	PaidTime                          time.Time                         `xml:",omitempty"`
	PaymentHoldDetails                *PaymentHoldDetailType            `xml:",omitempty"`
	PayPalEmailAddress                string                            `xml:",omitempty"`
	PickupDetails                     *PickupDetailsType                `xml:",omitempty"`
	PickupMethodSelected              *PickupMethodSelectedType         `xml:",omitempty"`
	Platform                          string                            `xml:",omitempty"`
	QuantityPurchased                 int                               `xml:",omitempty"`
	SellerContactBuyerByEmail         bool                              `xml:",omitempty"`
	SellerDiscounts                   *SellerDiscountsType              `xml:",omitempty"`
	SellingManagerProductDetails      *SellingManagerProductDetailsType `xml:",omitempty"`
	ShippedTime                       time.Time                         `xml:",omitempty"`
	ShippingConvenienceCharge         float64                           `xml:",omitempty"`
	ShippingDetails                   *ShippingDetailsType              `xml:",omitempty"`
	ShippingServiceSelected           *ShippingServiceOptionsType       `xml:",omitempty"`
	Status                            *TransactionStatusType            `xml:",omitempty"`
	Taxes                             *TaxesType                        `xml:",omitempty"`
	TransactionID                     string                            `xml:",omitempty"`
	TransactionPrice                  float64                           `xml:",omitempty"`
	TransactionSiteID                 string                            `xml:",omitempty"`
	UnpaidItem                        *UnpaidItemType                   `xml:",omitempty"`
	Variation                         *VariationType                    `xml:",omitempty"`
}

type SellingManagerProductDetailsType struct {
	XMLName           xml.Name                         `xml:"SellingManagerProductDetails"`
	CustomLabel       string                           `xml:",omitempty"`
	FolderID          float64                          `xml:",omitempty"`
	Note              string                           `xml:",omitempty"`
	ProductID         float64                          `xml:",omitempty"`
	ProductName       string                           `xml:",omitempty"`
	QuantityAvailable int                              `xml:",omitempty"`
	RestockAlert      bool                             `xml:",omitempty"`
	RestockThreshold  int                              `xml:",omitempty"`
	UnitCost          float64                          `xml:",omitempty"`
	VendorInfo        *SellingManagerVendorDetailsType `xml:",omitempty"`
}

type SellingManagerProductSpecificsType struct {
	XMLName           xml.Name                 `xml:"SellingManagerProductSpecifics"`
	ItemSpecifics     []NameValueListArrayType `xml:",omitempty"`
	PrimaryCategoryID string                   `xml:",omitempty"`
	Variations        *VariationsType          `xml:",omitempty"`
}

type VariationKeyType struct {
	XMLName            xml.Name                 `xml:"VariationKey"`
	ItemID             string                   `xml:",omitempty"`
	VariationSpecifics []NameValueListArrayType `xml:",omitempty"`
}

type FeedbackInfoType struct {
	XMLName     xml.Name `xml:"FeedbackInfo"`
	CommentType string   `xml:",omitempty"`
}

type ShipmentType struct {
	XMLName                 xml.Name                     `xml:"Shipment"`
	ShipmentTrackingDetails *ShipmentTrackingDetailsType `xml:",omitempty"`
}

type ShipmentTrackingDetailsType struct {
	XMLName                xml.Name              `xml:"ShipmentTrackingDetails"`
	ShipmentLineItem       *ShipmentLineItemType `xml:",omitempty"`
	ShipmentTrackingNumber string                `xml:",omitempty"`
	ShippingCarrierUsed    string                `xml:",omitempty"`
}

type ShipmentLineItemType struct {
	XMLName  xml.Name      `xml:"ShipmentLineItem"`
	LineItem *LineItemType `xml:",omitempty"`
}

type LineItemType struct {
	XMLName         xml.Name `xml:"LineItem"`
	CountryOfOrigin string   `xml:",omitempty"`
	Description     string   `xml:",omitempty"`
	ItemID          string   `xml:",omitempty"`
	Quantity        int      `xml:",omitempty"`
	TransactionID   string   `xml:",omitempty"`
}

type MyMessagesMessageIDArrayType struct {
	XMLName   xml.Name `xml:"MyMessagesMessageIDArray"`
	MessageID string   `xml:",omitempty"`
}

type SellingManagerAutoListType struct {
	XMLName                 xml.Name                                       `xml:"SellingManagerAutoList"`
	KeepMinActive           *SellingManagerAutoListMinActiveItemsType      `xml:",omitempty"`
	ListAccordingToSchedule *SellingManagerAutoListAccordingToScheduleType `xml:",omitempty"`
	SourceSaleTemplateID    float64                                        `xml:",omitempty"`
}

type SellingManagerAutoListMinActiveItemsType struct {
	XMLName                   xml.Name  `xml:"SellingManagerAutoListMinActiveItems"`
	ListingHoldInventoryLevel int       `xml:",omitempty"`
	ListTimeFrom              time.Time `xml:",omitempty"`
	ListTimeTo                time.Time `xml:",omitempty"`
	MinActiveItemCount        int       `xml:",omitempty"`
	SpacingIntervalInMinutes  int       `xml:",omitempty"`
}

type SellingManagerAutoListAccordingToScheduleType struct {
	XMLName                   xml.Name  `xml:"SellingManagerAutoListAccordingToSchedule"`
	DayOfWeek                 string    `xml:",omitempty"`
	EndTime                   time.Time `xml:",omitempty"`
	ListAtSpecificTimeOfDay   time.Time `xml:",omitempty"`
	ListingHoldInventoryLevel int       `xml:",omitempty"`
	ListingPeriodInWeeks      int       `xml:",omitempty"`
	MaxActiveItemCount        int       `xml:",omitempty"`
	StartTime                 time.Time `xml:",omitempty"`
}

type SellingManagerAutoRelistType struct {
	XMLName                   xml.Name              `xml:"SellingManagerAutoRelist"`
	BestOfferDetails          *BestOfferDetailsType `xml:",omitempty"`
	ListingHoldInventoryLevel int                   `xml:",omitempty"`
	RelistAfterDays           int                   `xml:",omitempty"`
	RelistAfterHours          int                   `xml:",omitempty"`
	RelistAtSpecificTimeOfDay time.Time             `xml:",omitempty"`
	RelistCondition           string                `xml:",omitempty"`
	Type                      string                `xml:",omitempty"`
}

type SellingManagerAutoSecondChanceOfferType struct {
	XMLName                    xml.Name `xml:"SellingManagerAutoSecondChanceOffer"`
	Amount                     float64  `xml:",omitempty"`
	Duration                   string   `xml:",omitempty"`
	ListingHoldInventoryLevel  int      `xml:",omitempty"`
	ProfitPercent              float64  `xml:",omitempty"`
	SecondChanceOfferCondition string   `xml:",omitempty"`
}

type EndItemRequestContainerType struct {
	XMLName      xml.Name `xml:"EndItemRequestContainer"`
	EndingReason string   `xml:",omitempty"`
	ItemID       string   `xml:",omitempty"`
	MessageID    string   `xml:",omitempty"`
}

type EndItemResponseContainerType struct {
	XMLName       xml.Name   `xml:"EndItemResponseContainer"`
	CorrelationID string     `xml:",omitempty"`
	EndTime       time.Time  `xml:",omitempty"`
	Errors        *ErrorType `xml:",omitempty"`
}

type PaginationType struct {
	XMLName        xml.Name `xml:"Pagination"`
	EntriesPerPage int      `xml:",omitempty"`
	PageNumber     int      `xml:",omitempty"`
}

type AccountEntriesType struct {
	XMLName      xml.Name          `xml:"AccountEntries"`
	AccountEntry *AccountEntryType `xml:",omitempty"`
}

type AccountEntryType struct {
	XMLName                  xml.Name  `xml:"AccountEntry"`
	AccountDetailsEntryType  string    `xml:",omitempty"`
	Balance                  float64   `xml:",omitempty"`
	ConversionRate           float64   `xml:",omitempty"`
	Date                     time.Time `xml:",omitempty"`
	Description              string    `xml:",omitempty"`
	GrossDetailAmount        float64   `xml:",omitempty"`
	ItemID                   string    `xml:",omitempty"`
	Memo                     string    `xml:",omitempty"`
	NetDetailAmount          float64   `xml:",omitempty"`
	OrderLineItemID          string    `xml:",omitempty"`
	ReceivedTopRatedDiscount bool      `xml:",omitempty"`
	RefNumber                string    `xml:",omitempty"`
	Title                    string    `xml:",omitempty"`
	TransactionID            string    `xml:",omitempty"`
	VATPercent               float64   `xml:",omitempty"`
}

type AccountSummaryType struct {
	XMLName              xml.Name               `xml:"AccountSummary"`
	AccountState         string                 `xml:",omitempty"`
	AdditionalAccount    *AdditionalAccountType `xml:",omitempty"`
	AmountPastDue        float64                `xml:",omitempty"`
	BankAccountInfo      string                 `xml:",omitempty"`
	BankModifyDate       time.Time              `xml:",omitempty"`
	BillingCycleDate     int                    `xml:",omitempty"`
	CreditCardExpiration time.Time              `xml:",omitempty"`
	CreditCardInfo       string                 `xml:",omitempty"`
	CreditCardModifyDate time.Time              `xml:",omitempty"`
	CurrentBalance       float64                `xml:",omitempty"`
	InvoiceBalance       float64                `xml:",omitempty"`
	InvoiceCredit        float64                `xml:",omitempty"`
	InvoiceDate          time.Time              `xml:",omitempty"`
	InvoiceNewFee        float64                `xml:",omitempty"`
	InvoicePayment       float64                `xml:",omitempty"`
	LastAmountPaid       float64                `xml:",omitempty"`
	LastPaymentDate      time.Time              `xml:",omitempty"`
	PastDue              bool                   `xml:",omitempty"`
	PaymentMethod        string                 `xml:",omitempty"`
}

type AdditionalAccountType struct {
	XMLName     xml.Name `xml:"AdditionalAccount"`
	AccountCode string   `xml:",omitempty"`
	Balance     float64  `xml:",omitempty"`
	Currency    string   `xml:",omitempty"`
}

type PaginationResultType struct {
	XMLName              xml.Name `xml:"PaginationResult"`
	TotalNumberOfEntries int      `xml:",omitempty"`
	TotalNumberOfPages   int      `xml:",omitempty"`
}

type AdFormatLeadType struct {
	XMLName               xml.Name                         `xml:"AdFormatLead"`
	AdditionalInformation string                           `xml:",omitempty"`
	Address               *AddressType                     `xml:",omitempty"`
	Answer1               bool                             `xml:",omitempty"`
	Answer2               bool                             `xml:",omitempty"`
	BestTimeToCall        string                           `xml:",omitempty"`
	Email                 string                           `xml:",omitempty"`
	ExternalEmail         string                           `xml:",omitempty"`
	FinancingAnswer       bool                             `xml:",omitempty"`
	ItemID                string                           `xml:",omitempty"`
	ItemTitle             string                           `xml:",omitempty"`
	MemberMessage         []MemberMessageExchangeArrayType `xml:",omitempty"`
	PurchaseTimeFrame     string                           `xml:",omitempty"`
	Status                string                           `xml:",omitempty"`
	SubmittedTime         time.Time                        `xml:",omitempty"`
	TradeInMake           string                           `xml:",omitempty"`
	TradeInModel          string                           `xml:",omitempty"`
	TradeInYear           string                           `xml:",omitempty"`
	UserID                string                           `xml:",omitempty"`
}

type MemberMessageExchangeArrayType struct {
	XMLName               xml.Name                   `xml:"MemberMessageExchangeArray"`
	MemberMessageExchange *MemberMessageExchangeType `xml:",omitempty"`
}

type MemberMessageExchangeType struct {
	XMLName          xml.Name           `xml:"MemberMessageExchange"`
	Item             *ItemType          `xml:",omitempty"`
	LastModifiedDate time.Time          `xml:",omitempty"`
	MessageMedia     *MessageMediaType  `xml:",omitempty"`
	MessageStatus    string             `xml:",omitempty"`
	Question         *MemberMessageType `xml:",omitempty"`
	Response         string             `xml:",omitempty"`
}

type OfferArrayType struct {
	XMLName xml.Name   `xml:"OfferArray"`
	Offer   *OfferType `xml:",omitempty"`
}

type OfferType struct {
	XMLName     xml.Name `xml:"Offer"`
	Action      string   `xml:",omitempty"`
	BestOfferID string   `xml:",omitempty"`
	MaxBid      float64  `xml:",omitempty"`
	Message     string   `xml:",omitempty"`
	Quantity    int      `xml:",omitempty"`
	UserConsent bool     `xml:",omitempty"`
}

type BiddingSummaryType struct {
	XMLName                xml.Name            `xml:"BiddingSummary"`
	BidActivityWithSeller  int                 `xml:",omitempty"`
	BidRetractions         int                 `xml:",omitempty"`
	BidsToUniqueCategories int                 `xml:",omitempty"`
	BidsToUniqueSellers    int                 `xml:",omitempty"`
	ItemBidDetails         *ItemBidDetailsType `xml:",omitempty"`
	SummaryDays            int                 `xml:",omitempty"`
	TotalBids              int                 `xml:",omitempty"`
}

type ItemBidDetailsType struct {
	XMLName     xml.Name  `xml:"ItemBidDetails"`
	BidCount    int       `xml:",omitempty"`
	CategoryID  string    `xml:",omitempty"`
	ItemID      string    `xml:",omitempty"`
	LastBidTime time.Time `xml:",omitempty"`
	SellerID    string    `xml:",omitempty"`
}

type BuyerType struct {
	XMLName            xml.Name           `xml:"Buyer"`
	BuyerTaxIdentifier *TaxIdentifierType `xml:",omitempty"`
	ShippingAddress    *AddressType       `xml:",omitempty"`
}

type ApiAccessRuleType struct {
	XMLName               xml.Name  `xml:"ApiAccessRule"`
	CallName              string    `xml:",omitempty"`
	CountsTowardAggregate bool      `xml:",omitempty"`
	DailyHardLimit        float64   `xml:",omitempty"`
	DailySoftLimit        float64   `xml:",omitempty"`
	DailyUsage            float64   `xml:",omitempty"`
	HourlyHardLimit       float64   `xml:",omitempty"`
	HourlySoftLimit       float64   `xml:",omitempty"`
	HourlyUsage           float64   `xml:",omitempty"`
	ModTime               time.Time `xml:",omitempty"`
	Period                int       `xml:",omitempty"`
	PeriodicHardLimit     float64   `xml:",omitempty"`
	PeriodicSoftLimit     float64   `xml:",omitempty"`
	PeriodicStartDate     time.Time `xml:",omitempty"`
	PeriodicUsage         float64   `xml:",omitempty"`
	RuleCurrentStatus     string    `xml:",omitempty"`
	RuleStatus            string    `xml:",omitempty"`
}

type BestOfferArrayType struct {
	XMLName   xml.Name       `xml:"BestOfferArray"`
	BestOffer *BestOfferType `xml:",omitempty"`
}

type BestOfferType struct {
	XMLName    xml.Name `xml:"BestOffer"`
	CallStatus string   `xml:",omitempty"`
}

type ItemBestOffersArrayType struct {
	XMLName        xml.Name            `xml:"ItemBestOffersArray"`
	ItemBestOffers *ItemBestOffersType `xml:",omitempty"`
}

type ItemBestOffersType struct {
	XMLName        xml.Name             `xml:"ItemBestOffers"`
	BestOfferArray []BestOfferArrayType `xml:",omitempty"`
	Item           *ItemType            `xml:",omitempty"`
	Role           string               `xml:",omitempty"`
}

type ItemArrayType struct {
	XMLName xml.Name  `xml:"ItemArray"`
	Item    *ItemType `xml:",omitempty"`
}

type ReviseStatusType struct {
	XMLName         xml.Name `xml:"ReviseStatus"`
	BuyItNowAdded   bool     `xml:",omitempty"`
	BuyItNowLowered bool     `xml:",omitempty"`
	ItemRevised     bool     `xml:",omitempty"`
	ReserveLowered  bool     `xml:",omitempty"`
	ReserveRemoved  bool     `xml:",omitempty"`
}

type SellerType struct {
	XMLName                          xml.Name                              `xml:"Seller"`
	AllowPaymentEdit                 bool                                  `xml:",omitempty"`
	CharityAffiliationDetails        *CharityAffiliationDetailsType        `xml:",omitempty"`
	CharityRegistered                bool                                  `xml:",omitempty"`
	CheckoutEnabled                  bool                                  `xml:",omitempty"`
	CIPBankAccountStored             bool                                  `xml:",omitempty"`
	DomesticRateTable                bool                                  `xml:",omitempty"`
	FeatureEligibility               *FeatureEligibilityType               `xml:",omitempty"`
	GoodStanding                     bool                                  `xml:",omitempty"`
	IntegratedMerchantCreditCardInfo *IntegratedMerchantCreditCardInfoType `xml:",omitempty"`
	InternationalRateTable           bool                                  `xml:",omitempty"`
	PaisaPayEscrowEMIStatus          int                                   `xml:",omitempty"`
	PaisaPayStatus                   int                                   `xml:",omitempty"`
	PaymentMethod                    string                                `xml:",omitempty"`
	QualifiesForB2BVAT               bool                                  `xml:",omitempty"`
	RecoupmentPolicyConsent          *RecoupmentPolicyConsentType          `xml:",omitempty"`
	RegisteredBusinessSeller         bool                                  `xml:",omitempty"`
	SafePaymentExempt                bool                                  `xml:",omitempty"`
	SchedulingInfo                   *SchedulingInfoType                   `xml:",omitempty"`
	SellerBusinessType               string                                `xml:",omitempty"`
	SellerLevel                      string                                `xml:",omitempty"`
	SellerPaymentAddress             *AddressType                          `xml:",omitempty"`
	StoreOwner                       bool                                  `xml:",omitempty"`
	StoreSite                        string                                `xml:",omitempty"`
	StoreURL                         string                                `xml:",omitempty"`
	TopRatedSeller                   bool                                  `xml:",omitempty"`
	TopRatedSellerDetails            *TopRatedSellerDetailsType            `xml:",omitempty"`
	TransactionPercent               float64                               `xml:",omitempty"`
}

type SellingStatusType struct {
	XMLName               xml.Name               `xml:"SellingStatus"`
	ConvertedCurrentPrice float64                `xml:",omitempty"`
	CurrentPrice          float64                `xml:",omitempty"`
	HighBidder            *UserType              `xml:",omitempty"`
	MinimumToBid          float64                `xml:",omitempty"`
	ReserveMet            bool                   `xml:",omitempty"`
	SuggestedBidValues    *SuggestedBidValueType `xml:",omitempty"`
}

type PromotionalSaleDetailsType struct {
	XMLName       xml.Name  `xml:"PromotionalSaleDetails"`
	EndTime       time.Time `xml:",omitempty"`
	OriginalPrice float64   `xml:",omitempty"`
	StartTime     time.Time `xml:",omitempty"`
}

type TaxTableType struct {
	XMLName         xml.Name             `xml:"TaxTable"`
	TaxJurisdiction *TaxJurisdictionType `xml:",omitempty"`
}

type TaxJurisdictionType struct {
	XMLName               xml.Name `xml:"TaxJurisdiction"`
	JurisdictionID        string   `xml:",omitempty"`
	SalesTaxPercent       float64  `xml:",omitempty"`
	ShippingIncludedInTax bool     `xml:",omitempty"`
}

type CategoryArrayType struct {
	XMLName  xml.Name      `xml:"CategoryArray"`
	Category *CategoryType `xml:",omitempty"`
}

type CategoryFeatureType struct {
	XMLName                                   xml.Name                                 `xml:"CategoryFeature"`
	AdditionalCompatibilityEnabled            bool                                     `xml:",omitempty"`
	AdFormatEnabled                           string                                   `xml:",omitempty"`
	BestOfferAutoAcceptEnabled                bool                                     `xml:",omitempty"`
	BestOfferAutoDeclineEnabled               bool                                     `xml:",omitempty"`
	BestOfferCounterEnabled                   bool                                     `xml:",omitempty"`
	BestOfferEnabled                          bool                                     `xml:",omitempty"`
	BrandMPNIdentifierEnabled                 bool                                     `xml:",omitempty"`
	BuyerGuaranteeEnabled                     bool                                     `xml:",omitempty"`
	CategoryID                                string                                   `xml:",omitempty"`
	ClassifiedAdAutoAcceptEnabled             bool                                     `xml:",omitempty"`
	ClassifiedAdAutoDeclineEnabled            bool                                     `xml:",omitempty"`
	ClassifiedAdBestOfferEnabled              string                                   `xml:",omitempty"`
	ClassifiedAdCompanyNameEnabled            bool                                     `xml:",omitempty"`
	ClassifiedAdContactByAddressEnabled       bool                                     `xml:",omitempty"`
	ClassifiedAdContactByEmailEnabled         bool                                     `xml:",omitempty"`
	ClassifiedAdContactByPhoneEnabled         bool                                     `xml:",omitempty"`
	ClassifiedAdCounterOfferEnabled           bool                                     `xml:",omitempty"`
	ClassifiedAdPaymentMethodEnabled          string                                   `xml:",omitempty"`
	ClassifiedAdPayPerLeadEnabled             bool                                     `xml:",omitempty"`
	ClassifiedAdPhoneCount                    int                                      `xml:",omitempty"`
	ClassifiedAdShippingMethodEnabled         bool                                     `xml:",omitempty"`
	ClassifiedAdStreetCount                   int                                      `xml:",omitempty"`
	CompatibleVehicleType                     string                                   `xml:",omitempty"`
	ConditionEnabled                          string                                   `xml:",omitempty"`
	ConditionValues                           *ConditionValuesType                     `xml:",omitempty"`
	CrossBorderTradeAustraliaEnabled          bool                                     `xml:",omitempty"`
	CrossBorderTradeGBEnabled                 bool                                     `xml:",omitempty"`
	CrossBorderTradeNorthAmericaEnabled       bool                                     `xml:",omitempty"`
	DepositSupported                          bool                                     `xml:",omitempty"`
	DigitalGoodDeliveryEnabled                bool                                     `xml:",omitempty"`
	EANEnabled                                string                                   `xml:",omitempty"`
	eBayMotorsProAdFormatEnabled              string                                   `xml:",omitempty"`
	eBayMotorsProAutoAcceptEnabled            bool                                     `xml:",omitempty"`
	eBayMotorsProAutoDeclineEnabled           bool                                     `xml:",omitempty"`
	eBayMotorsProBestOfferEnabled             string                                   `xml:",omitempty"`
	eBayMotorsProCompanyNameEnabled           bool                                     `xml:",omitempty"`
	eBayMotorsProContactByAddressEnabled      bool                                     `xml:",omitempty"`
	eBayMotorsProContactByEmailEnabled        bool                                     `xml:",omitempty"`
	eBayMotorsProContactByPhoneEnabled        bool                                     `xml:",omitempty"`
	eBayMotorsProCounterOfferEnabled          bool                                     `xml:",omitempty"`
	eBayMotorsProPaymentMethodCheckOutEnabled string                                   `xml:",omitempty"`
	eBayMotorsProPhoneCount                   int                                      `xml:",omitempty"`
	eBayMotorsProSellerContactDetailsEnabled  bool                                     `xml:",omitempty"`
	eBayMotorsProShippingMethodEnabled        bool                                     `xml:",omitempty"`
	eBayMotorsProStreetCount                  int                                      `xml:",omitempty"`
	EpidSupported                             bool                                     `xml:",omitempty"`
	FreeGalleryPlusEnabled                    bool                                     `xml:",omitempty"`
	FreePicturePackEnabled                    bool                                     `xml:",omitempty"`
	GalleryFeaturedDurations                  *ListingEnhancementDurationReferenceType `xml:",omitempty"`
	GlobalShippingEnabled                     bool                                     `xml:",omitempty"`
	Group1MaxFlatShippingCost                 float64                                  `xml:",omitempty"`
	Group2MaxFlatShippingCost                 float64                                  `xml:",omitempty"`
	Group3MaxFlatShippingCost                 float64                                  `xml:",omitempty"`
	HandlingTimeEnabled                       bool                                     `xml:",omitempty"`
	HomePageFeaturedEnabled                   bool                                     `xml:",omitempty"`
	INEscrowWorkflowTimeline                  string                                   `xml:",omitempty"`
	ISBNEnabled                               string                                   `xml:",omitempty"`
	ItemCompatibilityEnabled                  string                                   `xml:",omitempty"`
	ItemSpecificsEnabled                      string                                   `xml:",omitempty"`
	KTypeSupported                            bool                                     `xml:",omitempty"`
	ListingDuration                           int                                      `xml:",omitempty"`
	LocalMarketAdFormatEnabled                string                                   `xml:",omitempty"`
	LocalMarketAutoAcceptEnabled              bool                                     `xml:",omitempty"`
	LocalMarketAutoDeclineEnabled             bool                                     `xml:",omitempty"`
	LocalMarketBestOfferEnabled               string                                   `xml:",omitempty"`
	LocalMarketCompanyNameEnabled             bool                                     `xml:",omitempty"`
	LocalMarketContactByAddressEnabled        bool                                     `xml:",omitempty"`
	LocalMarketContactByEmailEnabled          bool                                     `xml:",omitempty"`
	LocalMarketContactByPhoneEnabled          bool                                     `xml:",omitempty"`
	LocalMarketCounterOfferEnabled            bool                                     `xml:",omitempty"`
	LocalMarketNonSubscription                bool                                     `xml:",omitempty"`
	LocalMarketPaymentMethodCheckOutEnabled   string                                   `xml:",omitempty"`
	LocalMarketPhoneCount                     int                                      `xml:",omitempty"`
	LocalMarketPremiumSubscription            bool                                     `xml:",omitempty"`
	LocalMarketRegularSubscription            bool                                     `xml:",omitempty"`
	LocalMarketSellerContactDetailsEnabled    bool                                     `xml:",omitempty"`
	LocalMarketShippingMethodEnabled          bool                                     `xml:",omitempty"`
	LocalMarketSpecialitySubscription         bool                                     `xml:",omitempty"`
	LocalMarketStreetCount                    int                                      `xml:",omitempty"`
	MaxFlatShippingCost                       float64                                  `xml:",omitempty"`
	MaxGranularFitmentCount                   int                                      `xml:",omitempty"`
	MaxItemCompatibility                      int                                      `xml:",omitempty"`
	MinimumReservePrice                       float64                                  `xml:",omitempty"`
	MinItemCompatibility                      int                                      `xml:",omitempty"`
	NonSubscription                           string                                   `xml:",omitempty"`
	PaisaPayFullEscrowEnabled                 bool                                     `xml:",omitempty"`
	PaymentMethod                             string                                   `xml:",omitempty"`
	PaymentProfileCategoryGroup               string                                   `xml:",omitempty"`
	PayPalBuyerProtectionEnabled              bool                                     `xml:",omitempty"`
	PayPalRequired                            bool                                     `xml:",omitempty"`
	PickupDropOffEnabled                      bool                                     `xml:",omitempty"`
	PremiumSubscription                       string                                   `xml:",omitempty"`
	ProductCreationEnabled                    string                                   `xml:",omitempty"`
	ProPackEnabled                            bool                                     `xml:",omitempty"`
	ProPackPlusEnabled                        bool                                     `xml:",omitempty"`
	RegularSubscription                       string                                   `xml:",omitempty"`
	ReturnPolicyEnabled                       bool                                     `xml:",omitempty"`
	ReturnPolicyProfileCategoryGroup          string                                   `xml:",omitempty"`
	RevisePriceAllowed                        bool                                     `xml:",omitempty"`
	ReviseQuantityAllowed                     bool                                     `xml:",omitempty"`
	SafePaymentRequired                       bool                                     `xml:",omitempty"`
	SellerContactDetailsEnabled               bool                                     `xml:",omitempty"`
	SellerProvidedTitleSupported              bool                                     `xml:",omitempty"`
	ShippingProfileCategoryGroup              string                                   `xml:",omitempty"`
	ShippingTermsRequired                     bool                                     `xml:",omitempty"`
	SkypeMeNonTransactionalEnabled            bool                                     `xml:",omitempty"`
	SkypeMeTransactionalEnabled               bool                                     `xml:",omitempty"`
	SpecialitySubscription                    string                                   `xml:",omitempty"`
	StoreOwnerExtendedListingDurations        *StoreOwnerExtendedListingDurationsType  `xml:",omitempty"`
	StoreOwnerExtendedListingDurationsEnabled bool                                     `xml:",omitempty"`
	TransactionConfirmationRequestEnabled     bool                                     `xml:",omitempty"`
	UPCEnabled                                string                                   `xml:",omitempty"`
	UserConsentRequired                       bool                                     `xml:",omitempty"`
	ValueCategory                             bool                                     `xml:",omitempty"`
	ValuePackEnabled                          bool                                     `xml:",omitempty"`
	VariationsEnabled                         bool                                     `xml:",omitempty"`
	VINSupported                              bool                                     `xml:",omitempty"`
	VRMSupported                              bool                                     `xml:",omitempty"`
}

type ConditionValuesType struct {
	XMLName          xml.Name       `xml:"ConditionValues"`
	Condition        *ConditionType `xml:",omitempty"`
	ConditionHelpURL string         `xml:",omitempty"`
}

type ConditionType struct {
	XMLName     xml.Name `xml:"Condition"`
	DisplayName string   `xml:",omitempty"`
	ID          int      `xml:",omitempty"`
}

type ListingEnhancementDurationReferenceType struct {
	XMLName  xml.Name `xml:"ListingEnhancementDurationReference"`
	Duration string   `xml:",omitempty"`
}

type StoreOwnerExtendedListingDurationsType struct {
	XMLName  xml.Name `xml:"StoreOwnerExtendedListingDurations"`
	Duration string   `xml:",omitempty"`
}

type FeatureDefinitionsType struct {
	XMLName                                   xml.Name                        `xml:"FeatureDefinitions"`
	AdditionalCompatibilityEnabled            string                          `xml:",omitempty"`
	AdFormatEnabled                           string                          `xml:",omitempty"`
	BestOfferAutoAcceptEnabled                string                          `xml:",omitempty"`
	BestOfferAutoDeclineEnabled               string                          `xml:",omitempty"`
	BestOfferCounterEnabled                   string                          `xml:",omitempty"`
	BestOfferEnabled                          string                          `xml:",omitempty"`
	BrandMPNIdentifierEnabled                 string                          `xml:",omitempty"`
	BuyerGuaranteeEnabled                     string                          `xml:",omitempty"`
	ClassifiedAdAutoAcceptEnabled             string                          `xml:",omitempty"`
	ClassifiedAdAutoDeclineEnabled            string                          `xml:",omitempty"`
	ClassifiedAdBestOfferEnabled              string                          `xml:",omitempty"`
	ClassifiedAdCompanyNameEnabled            string                          `xml:",omitempty"`
	ClassifiedAdContactByAddressEnabled       string                          `xml:",omitempty"`
	ClassifiedAdContactByEmailEnabled         string                          `xml:",omitempty"`
	ClassifiedAdContactByPhoneEnabled         string                          `xml:",omitempty"`
	ClassifiedAdCounterOfferEnabled           string                          `xml:",omitempty"`
	ClassifiedAdPaymentMethodEnabled          string                          `xml:",omitempty"`
	ClassifiedAdPayPerLeadEnabled             string                          `xml:",omitempty"`
	ClassifiedAdPhoneCount                    string                          `xml:",omitempty"`
	ClassifiedAdShippingMethodEnabled         string                          `xml:",omitempty"`
	ClassifiedAdStreetCount                   string                          `xml:",omitempty"`
	CompatibleVehicleType                     string                          `xml:",omitempty"`
	ConditionEnabled                          string                          `xml:",omitempty"`
	ConditionValues                           string                          `xml:",omitempty"`
	CrossBorderTradeAustraliaEnabled          string                          `xml:",omitempty"`
	CrossBorderTradeGBEnabled                 string                          `xml:",omitempty"`
	CrossBorderTradeNorthAmericaEnabled       string                          `xml:",omitempty"`
	DepositSupported                          string                          `xml:",omitempty"`
	DigitalGoodDeliveryEnabled                string                          `xml:",omitempty"`
	EANEnabled                                string                          `xml:",omitempty"`
	eBayMotorsProAdFormatEnabled              string                          `xml:",omitempty"`
	eBayMotorsProAutoAcceptEnabled            string                          `xml:",omitempty"`
	eBayMotorsProAutoDeclineEnabled           string                          `xml:",omitempty"`
	eBayMotorsProBestOfferEnabled             string                          `xml:",omitempty"`
	eBayMotorsProCompanyNameEnabled           string                          `xml:",omitempty"`
	eBayMotorsProContactByAddressEnabled      string                          `xml:",omitempty"`
	eBayMotorsProContactByEmailEnabled        string                          `xml:",omitempty"`
	eBayMotorsProContactByPhoneEnabled        string                          `xml:",omitempty"`
	eBayMotorsProCounterOfferEnabled          string                          `xml:",omitempty"`
	eBayMotorsProPaymentMethodCheckOutEnabled string                          `xml:",omitempty"`
	eBayMotorsProPhoneCount                   string                          `xml:",omitempty"`
	eBayMotorsProSellerContactDetailsEnabled  string                          `xml:",omitempty"`
	eBayMotorsProShippingMethodEnabled        string                          `xml:",omitempty"`
	eBayMotorsProStreetCount                  string                          `xml:",omitempty"`
	EpidSupported                             string                          `xml:",omitempty"`
	FreeGalleryPlusEnabled                    string                          `xml:",omitempty"`
	FreePicturePackEnabled                    string                          `xml:",omitempty"`
	GalleryFeaturedDurations                  string                          `xml:",omitempty"`
	GlobalShippingEnabled                     string                          `xml:",omitempty"`
	Group1MaxFlatShippingCost                 string                          `xml:",omitempty"`
	Group2MaxFlatShippingCost                 string                          `xml:",omitempty"`
	Group3MaxFlatShippingCost                 string                          `xml:",omitempty"`
	HandlingTimeEnabled                       string                          `xml:",omitempty"`
	HomePageFeaturedEnabled                   string                          `xml:",omitempty"`
	INEscrowWorkflowTimeline                  string                          `xml:",omitempty"`
	ISBNEnabled                               string                          `xml:",omitempty"`
	ItemCompatibilityEnabled                  string                          `xml:",omitempty"`
	ItemSpecificsEnabled                      string                          `xml:",omitempty"`
	KTypeSupported                            string                          `xml:",omitempty"`
	ListingDurations                          *ListingDurationDefinitionsType `xml:",omitempty"`
	LocalListingDistancesNonSubscription      string                          `xml:",omitempty"`
	LocalListingDistancesRegular              string                          `xml:",omitempty"`
	LocalListingDistancesSpecialty            string                          `xml:",omitempty"`
	LocalMarketAdFormatEnabled                string                          `xml:",omitempty"`
	LocalMarketAutoAcceptEnabled              string                          `xml:",omitempty"`
	LocalMarketAutoDeclineEnabled             string                          `xml:",omitempty"`
	LocalMarketBestOfferEnabled               string                          `xml:",omitempty"`
	LocalMarketCompanyNameEnabled             string                          `xml:",omitempty"`
	LocalMarketContactByAddressEnabled        string                          `xml:",omitempty"`
	LocalMarketContactByEmailEnabled          string                          `xml:",omitempty"`
	LocalMarketContactByPhoneEnabled          string                          `xml:",omitempty"`
	LocalMarketCounterOfferEnabled            string                          `xml:",omitempty"`
	LocalMarketNonSubscription                string                          `xml:",omitempty"`
	LocalMarketPaymentMethodCheckOutEnabled   string                          `xml:",omitempty"`
	LocalMarketPhoneCount                     string                          `xml:",omitempty"`
	LocalMarketPremiumSubscription            string                          `xml:",omitempty"`
	LocalMarketRegularSubscription            string                          `xml:",omitempty"`
	LocalMarketSellerContactDetailsEnabled    string                          `xml:",omitempty"`
	LocalMarketShippingMethodEnabled          string                          `xml:",omitempty"`
	LocalMarketSpecialitySubscription         string                          `xml:",omitempty"`
	LocalMarketStreetCount                    string                          `xml:",omitempty"`
	MaxFlatShippingCost                       string                          `xml:",omitempty"`
	MaxFlatShippingCostCBTExempt              string                          `xml:",omitempty"`
	MaxGranularFitmentCount                   string                          `xml:",omitempty"`
	MaxItemCompatibility                      string                          `xml:",omitempty"`
	MinimumReservePrice                       string                          `xml:",omitempty"`
	MinItemCompatibility                      string                          `xml:",omitempty"`
	NonSubscription                           string                          `xml:",omitempty"`
	PaisaPayFullEscrowEnabled                 string                          `xml:",omitempty"`
	PaymentMethod                             string                          `xml:",omitempty"`
	PaymentProfileCategoryGroup               string                          `xml:",omitempty"`
	PayPalBuyerProtectionEnabled              string                          `xml:",omitempty"`
	PayPalRequired                            string                          `xml:",omitempty"`
	PickupDropOffEnabled                      string                          `xml:",omitempty"`
	PremiumSubscription                       string                          `xml:",omitempty"`
	ProductCreationEnabled                    string                          `xml:",omitempty"`
	ProPackEnabled                            string                          `xml:",omitempty"`
	ProPackPlusEnabled                        string                          `xml:",omitempty"`
	RegularSubscription                       string                          `xml:",omitempty"`
	ReturnPolicyEnabled                       string                          `xml:",omitempty"`
	ReturnPolicyProfileCategoryGroup          string                          `xml:",omitempty"`
	RevisePriceAllowed                        string                          `xml:",omitempty"`
	ReviseQuantityAllowed                     string                          `xml:",omitempty"`
	SafePaymentRequired                       string                          `xml:",omitempty"`
	SellerContactDetailsEnabled               string                          `xml:",omitempty"`
	SellerProvidedTitleSupported              string                          `xml:",omitempty"`
	ShippingProfileCategoryGroup              string                          `xml:",omitempty"`
	ShippingTermsRequired                     string                          `xml:",omitempty"`
	SkypeMeNonTransactionalEnabled            string                          `xml:",omitempty"`
	SkypeMeTransactionalEnabled               string                          `xml:",omitempty"`
	SpecialitySubscription                    string                          `xml:",omitempty"`
	StoreOwnerExtendedListingDurations        string                          `xml:",omitempty"`
	StoreOwnerExtendedListingDurationsEnabled string                          `xml:",omitempty"`
	TransactionConfirmationRequestEnabled     string                          `xml:",omitempty"`
	UPCEnabled                                string                          `xml:",omitempty"`
	ValueCategory                             string                          `xml:",omitempty"`
	ValuePackEnabled                          string                          `xml:",omitempty"`
	VariationsEnabled                         string                          `xml:",omitempty"`
	VINSupported                              string                          `xml:",omitempty"`
	VRMSupported                              string                          `xml:",omitempty"`
}

type ListingDurationDefinitionsType struct {
	XMLName         xml.Name                       `xml:"ListingDurationDefinitions"`
	ListingDuration *ListingDurationDefinitionType `xml:",omitempty"`
}

type ListingDurationDefinitionType struct {
	XMLName  xml.Name `xml:"ListingDurationDefinition"`
	Duration string   `xml:",omitempty"`
}

type SiteDefaultsType struct {
	XMLName                                   xml.Name                                 `xml:"SiteDefaults"`
	AdditionalCompatibilityEnabled            bool                                     `xml:",omitempty"`
	AdFormatEnabled                           string                                   `xml:",omitempty"`
	BasicUpgradePackEnabled                   bool                                     `xml:",omitempty"`
	BestOfferAutoAcceptEnabled                bool                                     `xml:",omitempty"`
	BestOfferAutoDeclineEnabled               bool                                     `xml:",omitempty"`
	BestOfferCounterEnabled                   bool                                     `xml:",omitempty"`
	BestOfferEnabled                          bool                                     `xml:",omitempty"`
	BrandMPNIdentifierEnabled                 bool                                     `xml:",omitempty"`
	BuyerGuaranteeEnabled                     bool                                     `xml:",omitempty"`
	ClassifiedAdAutoAcceptEnabled             bool                                     `xml:",omitempty"`
	ClassifiedAdAutoDeclineEnabled            bool                                     `xml:",omitempty"`
	ClassifiedAdBestOfferEnabled              string                                   `xml:",omitempty"`
	ClassifiedAdCompanyNameEnabled            bool                                     `xml:",omitempty"`
	ClassifiedAdContactByAddressEnabled       bool                                     `xml:",omitempty"`
	ClassifiedAdContactByEmailEnabled         bool                                     `xml:",omitempty"`
	ClassifiedAdContactByPhoneEnabled         bool                                     `xml:",omitempty"`
	ClassifiedAdCounterOfferEnabled           bool                                     `xml:",omitempty"`
	ClassifiedAdPaymentMethodEnabled          string                                   `xml:",omitempty"`
	ClassifiedAdPayPerLeadEnabled             bool                                     `xml:",omitempty"`
	ClassifiedAdPhoneCount                    int                                      `xml:",omitempty"`
	ClassifiedAdShippingMethodEnabled         bool                                     `xml:",omitempty"`
	ClassifiedAdStreetCount                   int                                      `xml:",omitempty"`
	CompatibleVehicleType                     string                                   `xml:",omitempty"`
	ConditionEnabled                          string                                   `xml:",omitempty"`
	ConditionValues                           *ConditionValuesType                     `xml:",omitempty"`
	CrossBorderTradeAustraliaEnabled          bool                                     `xml:",omitempty"`
	CrossBorderTradeGBEnabled                 bool                                     `xml:",omitempty"`
	CrossBorderTradeNorthAmericaEnabled       bool                                     `xml:",omitempty"`
	DepositSupported                          bool                                     `xml:",omitempty"`
	DigitalGoodDeliveryEnabled                bool                                     `xml:",omitempty"`
	DutchBINEnabled                           bool                                     `xml:",omitempty"`
	EANEnabled                                string                                   `xml:",omitempty"`
	EANIdentifierEnabled                      bool                                     `xml:",omitempty"`
	eBayMotorsProAdFormatEnabled              string                                   `xml:",omitempty"`
	eBayMotorsProAutoAcceptEnabled            bool                                     `xml:",omitempty"`
	eBayMotorsProAutoDeclineEnabled           bool                                     `xml:",omitempty"`
	eBayMotorsProBestOfferEnabled             string                                   `xml:",omitempty"`
	eBayMotorsProCompanyNameEnabled           bool                                     `xml:",omitempty"`
	eBayMotorsProContactByAddressEnabled      bool                                     `xml:",omitempty"`
	eBayMotorsProContactByEmailEnabled        bool                                     `xml:",omitempty"`
	eBayMotorsProContactByPhoneEnabled        bool                                     `xml:",omitempty"`
	eBayMotorsProCounterOfferEnabled          bool                                     `xml:",omitempty"`
	eBayMotorsProPaymentMethodCheckOutEnabled string                                   `xml:",omitempty"`
	eBayMotorsProPhoneCount                   int                                      `xml:",omitempty"`
	eBayMotorsProSellerContactDetailsEnabled  bool                                     `xml:",omitempty"`
	eBayMotorsProShippingMethodEnabled        bool                                     `xml:",omitempty"`
	eBayMotorsProStreetCount                  int                                      `xml:",omitempty"`
	EpidSupported                             bool                                     `xml:",omitempty"`
	FreeGalleryPlusEnabled                    bool                                     `xml:",omitempty"`
	FreePicturePackEnabled                    bool                                     `xml:",omitempty"`
	GalleryFeaturedDurations                  *ListingEnhancementDurationReferenceType `xml:",omitempty"`
	GlobalShippingEnabled                     bool                                     `xml:",omitempty"`
	Group1MaxFlatShippingCost                 float64                                  `xml:",omitempty"`
	Group2MaxFlatShippingCost                 float64                                  `xml:",omitempty"`
	Group3MaxFlatShippingCost                 float64                                  `xml:",omitempty"`
	HandlingTimeEnabled                       bool                                     `xml:",omitempty"`
	HomePageFeaturedEnabled                   bool                                     `xml:",omitempty"`
	INEscrowWorkflowTimeline                  string                                   `xml:",omitempty"`
	ISBNEnabled                               string                                   `xml:",omitempty"`
	ISBNIdentifierEnabled                     bool                                     `xml:",omitempty"`
	ItemCompatibilityEnabled                  string                                   `xml:",omitempty"`
	ItemSpecificsEnabled                      string                                   `xml:",omitempty"`
	KTypeSupported                            bool                                     `xml:",omitempty"`
	ListingDuration                           int                                      `xml:",omitempty"`
	LocalListingDistancesNonSubscription      string                                   `xml:",omitempty"`
	LocalListingDistancesRegular              string                                   `xml:",omitempty"`
	LocalListingDistancesSpecialty            string                                   `xml:",omitempty"`
	LocalMarketAdFormatEnabled                string                                   `xml:",omitempty"`
	LocalMarketAutoAcceptEnabled              bool                                     `xml:",omitempty"`
	LocalMarketAutoDeclineEnabled             bool                                     `xml:",omitempty"`
	LocalMarketBestOfferEnabled               string                                   `xml:",omitempty"`
	LocalMarketCompanyNameEnabled             bool                                     `xml:",omitempty"`
	LocalMarketContactByAddressEnabled        bool                                     `xml:",omitempty"`
	LocalMarketContactByEmailEnabled          bool                                     `xml:",omitempty"`
	LocalMarketContactByPhoneEnabled          bool                                     `xml:",omitempty"`
	LocalMarketCounterOfferEnabled            bool                                     `xml:",omitempty"`
	LocalMarketNonSubscription                bool                                     `xml:",omitempty"`
	LocalMarketPaymentMethodCheckOutEnabled   string                                   `xml:",omitempty"`
	LocalMarketPhoneCount                     int                                      `xml:",omitempty"`
	LocalMarketPremiumSubscription            bool                                     `xml:",omitempty"`
	LocalMarketRegularSubscription            bool                                     `xml:",omitempty"`
	LocalMarketSellerContactDetailsEnabled    bool                                     `xml:",omitempty"`
	LocalMarketShippingMethodEnabled          bool                                     `xml:",omitempty"`
	LocalMarketSpecialitySubscription         bool                                     `xml:",omitempty"`
	LocalMarketStreetCount                    int                                      `xml:",omitempty"`
	MaxFlatShippingCost                       float64                                  `xml:",omitempty"`
	MaxFlatShippingCostCBTExempt              bool                                     `xml:",omitempty"`
	MaxGranularFitmentCount                   int                                      `xml:",omitempty"`
	MaxItemCompatibility                      int                                      `xml:",omitempty"`
	MinimumReservePrice                       float64                                  `xml:",omitempty"`
	MinItemCompatibility                      int                                      `xml:",omitempty"`
	NonSubscription                           string                                   `xml:",omitempty"`
	PaisaPayFullEscrowEnabled                 bool                                     `xml:",omitempty"`
	PaymentMethod                             string                                   `xml:",omitempty"`
	PaymentOptionsGroup                       string                                   `xml:",omitempty"`
	PaymentProfileCategoryGroup               string                                   `xml:",omitempty"`
	PayPalBuyerProtectionEnabled              bool                                     `xml:",omitempty"`
	PayPalRequired                            bool                                     `xml:",omitempty"`
	PayPalRequiredForStoreOwner               bool                                     `xml:",omitempty"`
	PickupDropOffEnabled                      bool                                     `xml:",omitempty"`
	PremiumSubscription                       string                                   `xml:",omitempty"`
	ProductCreationEnabled                    string                                   `xml:",omitempty"`
	ProPackEnabled                            bool                                     `xml:",omitempty"`
	ProPackPlusEnabled                        bool                                     `xml:",omitempty"`
	RegularSubscription                       string                                   `xml:",omitempty"`
	ReturnPolicyEnabled                       bool                                     `xml:",omitempty"`
	ReturnPolicyProfileCategoryGroup          string                                   `xml:",omitempty"`
	RevisePriceAllowed                        bool                                     `xml:",omitempty"`
	ReviseQuantityAllowed                     bool                                     `xml:",omitempty"`
	SafePaymentRequired                       bool                                     `xml:",omitempty"`
	SellerContactDetailsEnabled               bool                                     `xml:",omitempty"`
	SellerProvidedTitleSupported              bool                                     `xml:",omitempty"`
	ShippingProfileCategoryGroup              string                                   `xml:",omitempty"`
	ShippingTermsRequired                     bool                                     `xml:",omitempty"`
	SkypeMeNonTransactionalEnabled            bool                                     `xml:",omitempty"`
	SkypeMeTransactionalEnabled               bool                                     `xml:",omitempty"`
	SpecialitySubscription                    string                                   `xml:",omitempty"`
	StoreOwnerExtendedListingDurations        *StoreOwnerExtendedListingDurationsType  `xml:",omitempty"`
	StoreOwnerExtendedListingDurationsEnabled bool                                     `xml:",omitempty"`
	TransactionConfirmationRequestEnabled     bool                                     `xml:",omitempty"`
	UPCEnabled                                string                                   `xml:",omitempty"`
	UPCIdentifierEnabled                      bool                                     `xml:",omitempty"`
	UserConsentRequired                       bool                                     `xml:",omitempty"`
	ValueCategory                             bool                                     `xml:",omitempty"`
	ValuePackEnabled                          bool                                     `xml:",omitempty"`
	VariationsEnabled                         bool                                     `xml:",omitempty"`
	VINSupported                              bool                                     `xml:",omitempty"`
	VRMSupported                              bool                                     `xml:",omitempty"`
}

type CategoryItemSpecificsType struct {
	XMLName       xml.Name                 `xml:"CategoryItemSpecifics"`
	CategoryID    string                   `xml:",omitempty"`
	ItemSpecifics []NameValueListArrayType `xml:",omitempty"`
}

type RecommendationsType struct {
	XMLName            xml.Name                `xml:"Recommendations"`
	CategoryID         string                  `xml:",omitempty"`
	NameRecommendation *NameRecommendationType `xml:",omitempty"`
	ProductIdentifiers *ProductIdentifiersType `xml:",omitempty"`
	Updated            bool                    `xml:",omitempty"`
}

type NameRecommendationType struct {
	XMLName             xml.Name                           `xml:"NameRecommendation"`
	HelpText            string                             `xml:",omitempty"`
	HelpURL             string                             `xml:",omitempty"`
	Name                string                             `xml:",omitempty"`
	ValidationRules     *RecommendationValidationRulesType `xml:",omitempty"`
	ValueRecommendation *ValueRecommendationType           `xml:",omitempty"`
}

type RecommendationValidationRulesType struct {
	XMLName      xml.Name                   `xml:"RecommendationValidationRules"`
	Confidence   int                        `xml:",omitempty"`
	Relationship *NameValueRelationshipType `xml:",omitempty"`
}

type NameValueRelationshipType struct {
	XMLName     xml.Name `xml:"NameValueRelationship"`
	ParentName  string   `xml:",omitempty"`
	ParentValue string   `xml:",omitempty"`
}

type ValueRecommendationType struct {
	XMLName         xml.Name                           `xml:"ValueRecommendation"`
	ValidationRules *RecommendationValidationRulesType `xml:",omitempty"`
	Value           string                             `xml:",omitempty"`
}

type ProductIdentifiersType struct {
	XMLName            xml.Name                  `xml:"ProductIdentifiers"`
	NameRecommendation *NameRecommendationType   `xml:",omitempty"`
	ValidationRules    *GroupValidationRulesType `xml:",omitempty"`
}

type GroupValidationRulesType struct {
	XMLName     xml.Name `xml:"GroupValidationRules"`
	MinRequired int      `xml:",omitempty"`
}

type CharityInfoType struct {
	XMLName        xml.Name `xml:"CharityInfo"`
	CharityDomain  int      `xml:",omitempty"`
	CharityRegion  int      `xml:",omitempty"`
	Description    string   `xml:",omitempty"`
	ExternalID     string   `xml:",omitempty"`
	LogoURL        string   `xml:",omitempty"`
	LogoURLSelling string   `xml:",omitempty"`
	Mission        string   `xml:",omitempty"`
	Name           string   `xml:",omitempty"`
}

type ContextSearchAssetType struct {
	XMLName  xml.Name      `xml:"ContextSearchAsset"`
	Category *CategoryType `xml:",omitempty"`
	Keyword  string        `xml:",omitempty"`
	Ranking  int           `xml:",omitempty"`
}

type DescriptionTemplateType struct {
	XMLName     xml.Name `xml:"DescriptionTemplate"`
	GroupID     int      `xml:",omitempty"`
	ID          int      `xml:",omitempty"`
	ImageURL    string   `xml:",omitempty"`
	Name        string   `xml:",omitempty"`
	TemplateXML string   `xml:",omitempty"`
	Type        string   `xml:",omitempty"`
}

type ThemeGroupType struct {
	XMLName    xml.Name `xml:"ThemeGroup"`
	GroupID    int      `xml:",omitempty"`
	GroupName  string   `xml:",omitempty"`
	ThemeID    int      `xml:",omitempty"`
	ThemeTotal int      `xml:",omitempty"`
}

type DisputeType struct {
	XMLName                  xml.Name               `xml:"Dispute"`
	DisputeCreatedTime       time.Time              `xml:",omitempty"`
	DisputeCreditEligibility string                 `xml:",omitempty"`
	DisputeExplanation       string                 `xml:",omitempty"`
	DisputeID                string                 `xml:",omitempty"`
	DisputeMessage           *DisputeMessageType    `xml:",omitempty"`
	DisputeModifiedTime      time.Time              `xml:",omitempty"`
	DisputeReason            string                 `xml:",omitempty"`
	DisputeRecordType        string                 `xml:",omitempty"`
	DisputeResolution        *DisputeResolutionType `xml:",omitempty"`
	DisputeState             string                 `xml:",omitempty"`
	DisputeStatus            string                 `xml:",omitempty"`
	Item                     *ItemType              `xml:",omitempty"`
	OrderLineItemID          string                 `xml:",omitempty"`
	OtherPartyName           string                 `xml:",omitempty"`
	OtherPartyRole           string                 `xml:",omitempty"`
	TransactionID            string                 `xml:",omitempty"`
	UserRole                 string                 `xml:",omitempty"`
}

type DisputeMessageType struct {
	XMLName             xml.Name  `xml:"DisputeMessage"`
	MessageCreationTime time.Time `xml:",omitempty"`
	MessageID           int       `xml:",omitempty"`
	MessageSource       string    `xml:",omitempty"`
	MessageText         string    `xml:",omitempty"`
}

type SiteBuyerRequirementDetailsType struct {
	XMLName                      xml.Name                                 `xml:"SiteBuyerRequirementDetails"`
	DetailVersion                string                                   `xml:",omitempty"`
	LinkedPayPalAccount          bool                                     `xml:",omitempty"`
	MaximumBuyerPolicyViolations *MaximumBuyerPolicyViolationsDetailsType `xml:",omitempty"`
	MaximumItemRequirements      *MaximumItemRequirementsDetailsType      `xml:",omitempty"`
	MaximumUnpaidItemStrikesInfo *MaximumUnpaidItemStrikesInfoDetailsType `xml:",omitempty"`
	MinimumFeedbackScore         *MinimumFeedbackScoreDetailsType         `xml:",omitempty"`
	ShipToRegistrationCountry    bool                                     `xml:",omitempty"`
	UpdateTime                   time.Time                                `xml:",omitempty"`
	VerifiedUserRequirements     *VerifiedUserRequirementsDetailsType     `xml:",omitempty"`
}

type MaximumBuyerPolicyViolationsDetailsType struct {
	XMLName                  xml.Name                             `xml:"MaximumBuyerPolicyViolationsDetails"`
	NumberOfPolicyViolations *NumberOfPolicyViolationsDetailsType `xml:",omitempty"`
	PolicyViolationDuration  *PolicyViolationDurationDetailsType  `xml:",omitempty"`
}

type NumberOfPolicyViolationsDetailsType struct {
	XMLName xml.Name `xml:"NumberOfPolicyViolationsDetails"`
	Count   int      `xml:",omitempty"`
}

type PolicyViolationDurationDetailsType struct {
	XMLName     xml.Name `xml:"PolicyViolationDurationDetails"`
	Description string   `xml:",omitempty"`
	Period      string   `xml:",omitempty"`
}

type MaximumItemRequirementsDetailsType struct {
	XMLName              xml.Name `xml:"MaximumItemRequirementsDetails"`
	MaximumItemCount     int      `xml:",omitempty"`
	MinimumFeedbackScore int      `xml:",omitempty"`
}

type MaximumUnpaidItemStrikesInfoDetailsType struct {
	XMLName                          xml.Name                                     `xml:"MaximumUnpaidItemStrikesInfoDetails"`
	MaximumUnpaidItemStrikesCount    *MaximumUnpaidItemStrikesCountDetailsType    `xml:",omitempty"`
	MaximumUnpaidItemStrikesDuration *MaximumUnpaidItemStrikesDurationDetailsType `xml:",omitempty"`
}

type MaximumUnpaidItemStrikesCountDetailsType struct {
	XMLName xml.Name `xml:"MaximumUnpaidItemStrikesCountDetails"`
	Count   int      `xml:",omitempty"`
}

type MaximumUnpaidItemStrikesDurationDetailsType struct {
	XMLName     xml.Name `xml:"MaximumUnpaidItemStrikesDurationDetails"`
	Description string   `xml:",omitempty"`
	Period      string   `xml:",omitempty"`
}

type MinimumFeedbackScoreDetailsType struct {
	XMLName       xml.Name `xml:"MinimumFeedbackScoreDetails"`
	FeedbackScore int      `xml:",omitempty"`
}

type VerifiedUserRequirementsDetailsType struct {
	XMLName       xml.Name `xml:"VerifiedUserRequirementsDetails"`
	FeedbackScore int      `xml:",omitempty"`
	VerifiedUser  bool     `xml:",omitempty"`
}

type CountryDetailsType struct {
	XMLName       xml.Name  `xml:"CountryDetails"`
	Country       string    `xml:",omitempty"`
	Description   string    `xml:",omitempty"`
	DetailVersion string    `xml:",omitempty"`
	UpdateTime    time.Time `xml:",omitempty"`
}

type CurrencyDetailsType struct {
	XMLName       xml.Name  `xml:"CurrencyDetails"`
	Currency      string    `xml:",omitempty"`
	Description   string    `xml:",omitempty"`
	DetailVersion string    `xml:",omitempty"`
	UpdateTime    time.Time `xml:",omitempty"`
}

type DispatchTimeMaxDetailsType struct {
	XMLName          xml.Name  `xml:"DispatchTimeMaxDetails"`
	Description      string    `xml:",omitempty"`
	DetailVersion    string    `xml:",omitempty"`
	DispatchTimeMax  int       `xml:",omitempty"`
	ExtendedHandling bool      `xml:",omitempty"`
	UpdateTime       time.Time `xml:",omitempty"`
}

type ExcludeShippingLocationDetailsType struct {
	XMLName       xml.Name  `xml:"ExcludeShippingLocationDetails"`
	Description   string    `xml:",omitempty"`
	DetailVersion string    `xml:",omitempty"`
	Location      string    `xml:",omitempty"`
	Region        string    `xml:",omitempty"`
	UpdateTime    time.Time `xml:",omitempty"`
}

type ItemSpecificDetailsType struct {
	XMLName                 xml.Name  `xml:"ItemSpecificDetails"`
	DetailVersion           string    `xml:",omitempty"`
	MaxCharactersPerName    int       `xml:",omitempty"`
	MaxCharactersPerValue   int       `xml:",omitempty"`
	MaxItemSpecificsPerItem int       `xml:",omitempty"`
	MaxValuesPerName        int       `xml:",omitempty"`
	UpdateTime              time.Time `xml:",omitempty"`
}

type ListingFeatureDetailsType struct {
	XMLName          xml.Name  `xml:"ListingFeatureDetails"`
	BoldTitle        string    `xml:",omitempty"`
	Border           string    `xml:",omitempty"`
	DetailVersion    string    `xml:",omitempty"`
	FeaturedFirst    string    `xml:",omitempty"`
	FeaturedPlus     string    `xml:",omitempty"`
	GiftIcon         string    `xml:",omitempty"`
	Highlight        string    `xml:",omitempty"`
	HomePageFeatured string    `xml:",omitempty"`
	ProPack          string    `xml:",omitempty"`
	UpdateTime       time.Time `xml:",omitempty"`
}

type ListingStartPriceDetailsType struct {
	XMLName                 xml.Name  `xml:"ListingStartPriceDetails"`
	Description             string    `xml:",omitempty"`
	DetailVersion           string    `xml:",omitempty"`
	ListingType             string    `xml:",omitempty"`
	MinBuyItNowPricePercent float64   `xml:",omitempty"`
	StartPrice              float64   `xml:",omitempty"`
	UpdateTime              time.Time `xml:",omitempty"`
}

type ProductDetailsType struct {
	XMLName                          xml.Name `xml:"ProductDetails"`
	ProductIdentifierUnavailableText string   `xml:",omitempty"`
}

type RecoupmentPolicyDetailsType struct {
	XMLName                    xml.Name  `xml:"RecoupmentPolicyDetails"`
	DetailVersion              string    `xml:",omitempty"`
	EnforcedOnListingSite      bool      `xml:",omitempty"`
	EnforcedOnRegistrationSite bool      `xml:",omitempty"`
	UpdateTime                 time.Time `xml:",omitempty"`
}

type ReturnPolicyDetailsType struct {
	XMLName            xml.Name                       `xml:"ReturnPolicyDetails"`
	Description        bool                           `xml:",omitempty"`
	DetailVersion      string                         `xml:",omitempty"`
	EAN                bool                           `xml:",omitempty"`
	Refund             *RefundDetailsType             `xml:",omitempty"`
	RestockingFeeValue *RestockingFeeValueDetailsType `xml:",omitempty"`
	ReturnsAccepted    *ReturnsAcceptedDetailsType    `xml:",omitempty"`
	ReturnsWithin      *ReturnsWithinDetailsType      `xml:",omitempty"`
	ShippingCostPaidBy *ShippingCostPaidByDetailsType `xml:",omitempty"`
	UpdateTime         time.Time                      `xml:",omitempty"`
	WarrantyDuration   *WarrantyDurationDetailsType   `xml:",omitempty"`
	WarrantyOffered    *WarrantyOfferedDetailsType    `xml:",omitempty"`
	WarrantyType       *WarrantyTypeDetailsType       `xml:",omitempty"`
}

type RefundDetailsType struct {
	XMLName      xml.Name `xml:"RefundDetails"`
	Description  string   `xml:",omitempty"`
	RefundOption string   `xml:",omitempty"`
}

type RestockingFeeValueDetailsType struct {
	XMLName                  xml.Name `xml:"RestockingFeeValueDetails"`
	Description              string   `xml:",omitempty"`
	RestockingFeeValueOption string   `xml:",omitempty"`
}

type ReturnsAcceptedDetailsType struct {
	XMLName               xml.Name `xml:"ReturnsAcceptedDetails"`
	Description           string   `xml:",omitempty"`
	ReturnsAcceptedOption string   `xml:",omitempty"`
}

type ReturnsWithinDetailsType struct {
	XMLName             xml.Name `xml:"ReturnsWithinDetails"`
	Description         string   `xml:",omitempty"`
	ReturnsWithinOption string   `xml:",omitempty"`
}

type ShippingCostPaidByDetailsType struct {
	XMLName                  xml.Name `xml:"ShippingCostPaidByDetails"`
	Description              string   `xml:",omitempty"`
	ShippingCostPaidByOption string   `xml:",omitempty"`
}

type WarrantyDurationDetailsType struct {
	XMLName                xml.Name `xml:"WarrantyDurationDetails"`
	Description            string   `xml:",omitempty"`
	WarrantyDurationOption string   `xml:",omitempty"`
}

type WarrantyOfferedDetailsType struct {
	XMLName               xml.Name `xml:"WarrantyOfferedDetails"`
	Description           string   `xml:",omitempty"`
	WarrantyOfferedOption string   `xml:",omitempty"`
}

type WarrantyTypeDetailsType struct {
	XMLName            xml.Name `xml:"WarrantyDetails"`
	Description        string   `xml:",omitempty"`
	WarrantyTypeOption string   `xml:",omitempty"`
}

type ShippingCarrierDetailsType struct {
	XMLName           xml.Name  `xml:"ShippingCarrierDetails"`
	Description       string    `xml:",omitempty"`
	DetailVersion     string    `xml:",omitempty"`
	ShippingCarrier   string    `xml:",omitempty"`
	ShippingCarrierID int       `xml:",omitempty"`
	UpdateTime        time.Time `xml:",omitempty"`
}

type ShippingCategoryDetailsType struct {
	XMLName          xml.Name  `xml:"ShippingCategoryDetails"`
	Description      string    `xml:",omitempty"`
	DetailVersion    string    `xml:",omitempty"`
	ShippingCategory string    `xml:",omitempty"`
	UpdateTime       time.Time `xml:",omitempty"`
}

type ShippingLocationDetailsType struct {
	XMLName          xml.Name  `xml:"ShippingLocationDetails"`
	Description      string    `xml:",omitempty"`
	DetailVersion    string    `xml:",omitempty"`
	ShippingLocation string    `xml:",omitempty"`
	UpdateTime       time.Time `xml:",omitempty"`
}

type ShippingPackageDetailsType struct {
	XMLName             xml.Name  `xml:"ShippingPackageDetails"`
	DefaultValue        bool      `xml:",omitempty"`
	Description         string    `xml:",omitempty"`
	DetailVersion       string    `xml:",omitempty"`
	DimensionsSupported bool      `xml:",omitempty"`
	PackageID           int       `xml:",omitempty"`
	ShippingPackage     string    `xml:",omitempty"`
	UpdateTime          time.Time `xml:",omitempty"`
}

type ShippingServiceDetailsType struct {
	XMLName                       xml.Name                           `xml:"ShippingServiceDetails"`
	CODService                    bool                               `xml:",omitempty"`
	CostGroupFlat                 string                             `xml:",omitempty"`
	DeprecationDetails            *AnnouncementMessageType           `xml:",omitempty"`
	Description                   string                             `xml:",omitempty"`
	DetailVersion                 string                             `xml:",omitempty"`
	DimensionsRequired            bool                               `xml:",omitempty"`
	ExpeditedService              bool                               `xml:",omitempty"`
	InternationalService          bool                               `xml:",omitempty"`
	MappedToShippingServiceID     int                                `xml:",omitempty"`
	ServiceType                   string                             `xml:",omitempty"`
	ShippingCarrier               string                             `xml:",omitempty"`
	ShippingCategory              string                             `xml:",omitempty"`
	ShippingPackage               string                             `xml:",omitempty"`
	ShippingService               string                             `xml:",omitempty"`
	ShippingServiceID             int                                `xml:",omitempty"`
	ShippingServicePackageDetails *ShippingServicePackageDetailsType `xml:",omitempty"`
	ShippingTimeMax               int                                `xml:",omitempty"`
	ShippingTimeMin               int                                `xml:",omitempty"`
	SurchargeApplicable           bool                               `xml:",omitempty"`
	UpdateTime                    time.Time                          `xml:",omitempty"`
	ValidForSellingFlow           bool                               `xml:",omitempty"`
	WeightRequired                bool                               `xml:",omitempty"`
}

type AnnouncementMessageType struct {
	XMLName               xml.Name  `xml:"AnnouncementMessage"`
	AnnouncementStartTime time.Time `xml:",omitempty"`
	EventTime             time.Time `xml:",omitempty"`
	MessageType           string    `xml:",omitempty"`
}

type ShippingServicePackageDetailsType struct {
	XMLName            xml.Name `xml:"ShippingServicePackageDetails"`
	DimensionsRequired bool     `xml:",omitempty"`
	Name               string   `xml:",omitempty"`
}

type SiteDetailsType struct {
	XMLName       xml.Name  `xml:"SiteDetails"`
	DetailVersion string    `xml:",omitempty"`
	Site          string    `xml:",omitempty"`
	SiteID        int       `xml:",omitempty"`
	UpdateTime    time.Time `xml:",omitempty"`
}

type TimeZoneDetailsType struct {
	XMLName                 xml.Name  `xml:"TimeZoneDetails"`
	DaylightSavingsInEffect bool      `xml:",omitempty"`
	DaylightSavingsLabel    string    `xml:",omitempty"`
	DaylightSavingsOffset   string    `xml:",omitempty"`
	DetailVersion           string    `xml:",omitempty"`
	StandardLabel           string    `xml:",omitempty"`
	StandardOffset          string    `xml:",omitempty"`
	TimeZoneID              string    `xml:",omitempty"`
	UpdateTime              time.Time `xml:",omitempty"`
}

type URLDetailsType struct {
	XMLName       xml.Name  `xml:"URLDetails"`
	DetailVersion string    `xml:",omitempty"`
	UpdateTime    time.Time `xml:",omitempty"`
	URL           string    `xml:",omitempty"`
	URLType       string    `xml:",omitempty"`
}

type VariationDetailsType struct {
	XMLName                               xml.Name  `xml:"VariationDetails"`
	DetailVersion                         string    `xml:",omitempty"`
	MaxNamesPerVariationSpecificsSet      int       `xml:",omitempty"`
	MaxValuesPerVariationSpecificsSetName int       `xml:",omitempty"`
	MaxVariationsPerItem                  int       `xml:",omitempty"`
	UpdateTime                            time.Time `xml:",omitempty"`
}

type FeedbackDetailArrayType struct {
	XMLName        xml.Name            `xml:"FeedbackDetailArray"`
	FeedbackDetail *FeedbackDetailType `xml:",omitempty"`
}

type FeedbackDetailType struct {
	XMLName             xml.Name  `xml:"FeedbackDetail"`
	CommentingUser      string    `xml:",omitempty"`
	CommentingUserScore int       `xml:",omitempty"`
	CommentReplaced     bool      `xml:",omitempty"`
	CommentText         string    `xml:",omitempty"`
	CommentTime         time.Time `xml:",omitempty"`
	CommentType         string    `xml:",omitempty"`
	Countable           bool      `xml:",omitempty"`
	FeedbackID          string    `xml:",omitempty"`
	FeedbackRatingStar  string    `xml:",omitempty"`
	FeedbackResponse    string    `xml:",omitempty"`
	FeedbackRevised     bool      `xml:",omitempty"`
	Followup            string    `xml:",omitempty"`
	FollowUpReplaced    bool      `xml:",omitempty"`
	ItemID              string    `xml:",omitempty"`
	ItemPrice           float64   `xml:",omitempty"`
	ItemTitle           string    `xml:",omitempty"`
	OrderLineItemID     string    `xml:",omitempty"`
	ResponseReplaced    bool      `xml:",omitempty"`
	Role                string    `xml:",omitempty"`
	TransactionID       string    `xml:",omitempty"`
}

type FeedbackSummaryType struct {
	XMLName                               xml.Name                       `xml:"FeedbackSummary"`
	BidRetractionFeedbackPeriodArray      []FeedbackPeriodArrayType      `xml:",omitempty"`
	BuyerRoleMetrics                      *BuyerRoleMetricsType          `xml:",omitempty"`
	NegativeFeedbackPeriodArray           []FeedbackPeriodArrayType      `xml:",omitempty"`
	NeutralCommentCountFromSuspendedUsers int                            `xml:",omitempty"`
	NeutralFeedbackPeriodArray            []FeedbackPeriodArrayType      `xml:",omitempty"`
	PositiveFeedbackPeriodArray           []FeedbackPeriodArrayType      `xml:",omitempty"`
	SellerRatingSummaryArray              []SellerRatingSummaryArrayType `xml:",omitempty"`
	SellerRoleMetrics                     *SellerRoleMetricsType         `xml:",omitempty"`
	TotalFeedbackPeriodArray              []FeedbackPeriodArrayType      `xml:",omitempty"`
	UniqueNegativeFeedbackCount           int                            `xml:",omitempty"`
	UniqueNeutralFeedbackCount            int                            `xml:",omitempty"`
	UniquePositiveFeedbackCount           int                            `xml:",omitempty"`
}

type FeedbackPeriodArrayType struct {
	XMLName        xml.Name            `xml:"FeedbackPeriodArray"`
	FeedbackPeriod *FeedbackPeriodType `xml:",omitempty"`
}

type FeedbackPeriodType struct {
	XMLName      xml.Name `xml:"FeedbackPeriod"`
	Count        int      `xml:",omitempty"`
	PeriodInDays int      `xml:",omitempty"`
}

type BuyerRoleMetricsType struct {
	XMLName                   xml.Name `xml:"BuyerRoleMetrics"`
	FeedbackLeftPercent       float64  `xml:",omitempty"`
	NegativeFeedbackLeftCount int      `xml:",omitempty"`
	NeutralFeedbackLeftCount  int      `xml:",omitempty"`
	PositiveFeedbackLeftCount int      `xml:",omitempty"`
}

type SellerRatingSummaryArrayType struct {
	XMLName              xml.Name                  `xml:"SellerRatingSummaryArray"`
	AverageRatingSummary *AverageRatingSummaryType `xml:",omitempty"`
}

type AverageRatingSummaryType struct {
	XMLName               xml.Name                  `xml:"AverageRatingSummary"`
	AverageRatingDetails  *AverageRatingDetailsType `xml:",omitempty"`
	FeedbackSummaryPeriod string                    `xml:",omitempty"`
}

type AverageRatingDetailsType struct {
	XMLName      xml.Name `xml:"AverageRatingDetails"`
	Rating       float64  `xml:",omitempty"`
	RatingCount  int      `xml:",omitempty"`
	RatingDetail string   `xml:",omitempty"`
}

type SellerRoleMetricsType struct {
	XMLName                       xml.Name `xml:"SellerRoleMetrics"`
	CrossBorderTransactionCount   int      `xml:",omitempty"`
	CrossBorderTransactionPercent float64  `xml:",omitempty"`
	FeedbackLeftPercent           float64  `xml:",omitempty"`
	NegativeFeedbackLeftCount     int      `xml:",omitempty"`
	NeutralFeedbackLeftCount      int      `xml:",omitempty"`
	PositiveFeedbackLeftCount     int      `xml:",omitempty"`
	RepeatBuyerCount              int      `xml:",omitempty"`
	RepeatBuyerPercent            float64  `xml:",omitempty"`
	TransactionPercent            float64  `xml:",omitempty"`
	UniqueBuyerCount              int      `xml:",omitempty"`
}

type BuyerProtectionDetailsType struct {
	XMLName               xml.Name `xml:"BuyerProtectionDetails"`
	BuyerProtectionSource string   `xml:",omitempty"`
	BuyerProtectionStatus string   `xml:",omitempty"`
}

type BusinessSellerDetailsType struct {
	XMLName                      xml.Name        `xml:"BusinessSellerDetails"`
	AdditionalContactInformation string          `xml:",omitempty"`
	Address                      *AddressType    `xml:",omitempty"`
	Email                        string          `xml:",omitempty"`
	Fax                          string          `xml:",omitempty"`
	LegalInvoice                 bool            `xml:",omitempty"`
	TermsAndConditions           string          `xml:",omitempty"`
	TradeRegistrationNumber      string          `xml:",omitempty"`
	VATDetails                   *VATDetailsType `xml:",omitempty"`
}

type ItemPolicyViolationType struct {
	XMLName    xml.Name `xml:"ItemPolicyViolation"`
	PolicyID   float64  `xml:",omitempty"`
	PolicyText string   `xml:",omitempty"`
}

type ExtendedPictureDetailsType struct {
	XMLName     xml.Name         `xml:"ExtendedPictureDetails"`
	PictureURLs *PictureURLsType `xml:",omitempty"`
}

type PictureURLsType struct {
	XMLName            xml.Name `xml:"PictureURLs"`
	eBayPictureURL     string   `xml:",omitempty"`
	ExternalPictureURL string   `xml:",omitempty"`
}

type CalculatedShippingDiscountType struct {
	XMLName         xml.Name             `xml:"CalculatedShippingDiscount"`
	DiscountName    string               `xml:",omitempty"`
	DiscountProfile *DiscountProfileType `xml:",omitempty"`
}

type DiscountProfileType struct {
	XMLName                  xml.Name `xml:"DiscountProfile"`
	DiscountProfileID        string   `xml:",omitempty"`
	DiscountProfileName      string   `xml:",omitempty"`
	EachAdditionalAmount     float64  `xml:",omitempty"`
	EachAdditionalAmountOff  float64  `xml:",omitempty"`
	EachAdditionalPercentOff float64  `xml:",omitempty"`
}

type FlatShippingDiscountType struct {
	XMLName         xml.Name             `xml:"FlatShippingDiscount"`
	DiscountName    string               `xml:",omitempty"`
	DiscountProfile *DiscountProfileType `xml:",omitempty"`
}

type PromotionalShippingDiscountDetailsType struct {
	XMLName      xml.Name `xml:"PromotionalShippingDiscountDetails"`
	DiscountName string   `xml:",omitempty"`
	ItemCount    int      `xml:",omitempty"`
	OrderAmount  float64  `xml:",omitempty"`
	ShippingCost float64  `xml:",omitempty"`
}

type PaginatedTransactionArrayType struct {
	XMLName          xml.Name               `xml:"PaginatedTransactionArray"`
	PaginationResult *PaginationResultType  `xml:",omitempty"`
	TransactionArray []TransactionArrayType `xml:",omitempty"`
}

type TaxIdentifierType struct {
	XMLName   xml.Name `xml:"TaxIdentifier"`
	Attribute string   `xml:",omitempty"`
	ID        string   `xml:",omitempty"`
	Type      string   `xml:",omitempty"`
}

type BuyerPackageEnclosuresType struct {
	XMLName               xml.Name `xml:"BuyerPackageEnclosures"`
	BuyerPackageEnclosure string   `xml:",omitempty"`
}

type CancelDetailType struct {
	XMLName             xml.Name  `xml:"CancelDetail"`
	CancelCompleteDate  time.Time `xml:",omitempty"`
	CancelIntiationDate time.Time `xml:",omitempty"`
	CancelIntiator      string    `xml:",omitempty"`
	CancelReason        string    `xml:",omitempty"`
	CancelReasonDetails string    `xml:",omitempty"`
}

type PaymentsInformationType struct {
	XMLName  xml.Name                `xml:"PaymentsInformation"`
	Payments *PaymentInformationType `xml:",omitempty"`
	Refunds  *RefundInformationType  `xml:",omitempty"`
}

type PaymentInformationType struct {
	XMLName xml.Name                `xml:"PaymentInformation"`
	Payment *PaymentTransactionType `xml:",omitempty"`
}

type PaymentTransactionType struct {
	XMLName            xml.Name  `xml:"PaymentTransaction"`
	FeeOrCreditAmount  float64   `xml:",omitempty"`
	Payee              string    `xml:",omitempty"`
	Payer              string    `xml:",omitempty"`
	PaymentAmount      float64   `xml:",omitempty"`
	PaymentReferenceID string    `xml:",omitempty"`
	PaymentStatus      string    `xml:",omitempty"`
	PaymentTime        time.Time `xml:",omitempty"`
	ReferenceID        string    `xml:",omitempty"`
}

type RefundInformationType struct {
	XMLName xml.Name                   `xml:"RefundInformation"`
	Refund  *RefundTransactionInfoType `xml:",omitempty"`
}

type RefundTransactionInfoType struct {
	XMLName           xml.Name  `xml:"RefundTransactionInfo"`
	FeeOrCreditAmount float64   `xml:",omitempty"`
	ReferenceID       string    `xml:",omitempty"`
	RefundAmount      float64   `xml:",omitempty"`
	RefundStatus      string    `xml:",omitempty"`
	RefundTime        time.Time `xml:",omitempty"`
	RefundTo          string    `xml:",omitempty"`
	RefundType        string    `xml:",omitempty"`
}

type MultiLegShippingDetailsType struct {
	XMLName                           xml.Name              `xml:"MultiLegShippingDetails"`
	SellerShipmentToLogisticsProvider *MultiLegShipmentType `xml:",omitempty"`
}

type MultiLegShipmentType struct {
	XMLName                xml.Name                     `xml:"MultiLegShipment"`
	ShippingServiceDetails *MultiLegShippingServiceType `xml:",omitempty"`
	ShippingTimeMax        int                          `xml:",omitempty"`
	ShippingTimeMin        int                          `xml:",omitempty"`
	ShipToAddress          *AddressType                 `xml:",omitempty"`
}

type MultiLegShippingServiceType struct {
	XMLName           xml.Name `xml:"MultiLegShippingService"`
	ShippingService   string   `xml:",omitempty"`
	TotalShippingCost float64  `xml:",omitempty"`
}

type PaymentHoldDetailType struct {
	XMLName                   xml.Name                        `xml:"PaymentHoldDetail"`
	ExpectedReleaseDate       time.Time                       `xml:",omitempty"`
	NumOfReqSellerActions     int                             `xml:",omitempty"`
	PaymentHoldReason         string                          `xml:",omitempty"`
	RequiredSellerActionArray []RequiredSellerActionArrayType `xml:",omitempty"`
}

type RequiredSellerActionArrayType struct {
	XMLName              xml.Name `xml:"RequiredSellerActionArray"`
	RequiredSellerAction string   `xml:",omitempty"`
}

type DigitalDeliverySelectedType struct {
	XMLName         xml.Name             `xml:"DigitalDeliverySelected"`
	DeliveryDetails *DeliveryDetailsType `xml:",omitempty"`
	DeliveryMethod  string               `xml:",omitempty"`
	DeliveryStatus  *DeliveryStatusType  `xml:",omitempty"`
}

type DeliveryDetailsType struct {
	XMLName   xml.Name                 `xml:"DeliveryDetails"`
	Recipient *DigitalDeliveryUserType `xml:",omitempty"`
	Sender    *DigitalDeliveryUserType `xml:",omitempty"`
}

type DigitalDeliveryUserType struct {
	XMLName xml.Name `xml:"DigitalDeliveryUser"`
	Email   string   `xml:",omitempty"`
	Name    string   `xml:",omitempty"`
}

type DeliveryStatusType struct {
	XMLName xml.Name `xml:"DeliveryStatus"`
	Email   string   `xml:",omitempty"`
}

type ExternalTransactionType struct {
	XMLName                 xml.Name  `xml:"ExternalTransaction"`
	ExternalTransactionID   string    `xml:",omitempty"`
	ExternalTransactionTime time.Time `xml:",omitempty"`
}

type GiftSummaryType struct {
	XMLName xml.Name `xml:"GiftSummary"`
	Message string   `xml:",omitempty"`
}

type PickupDetailsType struct {
	XMLName       xml.Name           `xml:"PickupDetails"`
	PickupOptions *PickupOptionsType `xml:",omitempty"`
}

type PickupOptionsType struct {
	XMLName        xml.Name `xml:"PickupOptions"`
	PickupMethod   string   `xml:",omitempty"`
	PickupPriority int      `xml:",omitempty"`
}

type PickupMethodSelectedType struct {
	XMLName               xml.Name  `xml:"PickupMethodSelected"`
	MerchantPickupCode    string    `xml:",omitempty"`
	PickupFulfillmentTime time.Time `xml:",omitempty"`
	PickupLocationUUID    string    `xml:",omitempty"`
	PickupMethod          string    `xml:",omitempty"`
	PickupStatus          string    `xml:",omitempty"`
	PickupStoreID         string    `xml:",omitempty"`
}

type SellerDiscountsType struct {
	XMLName                  xml.Name            `xml:"SellerDiscounts"`
	OriginalItemPrice        float64             `xml:",omitempty"`
	OriginalItemShippingCost float64             `xml:",omitempty"`
	OriginalShippingService  string              `xml:",omitempty"`
	SellerDiscount           *SellerDiscountType `xml:",omitempty"`
}

type SellerDiscountType struct {
	XMLName                xml.Name `xml:"SellerDiscount"`
	CampaignDisplayName    string   `xml:",omitempty"`
	CampaignID             float64  `xml:",omitempty"`
	ItemDiscountAmount     float64  `xml:",omitempty"`
	ShippingDiscountAmount float64  `xml:",omitempty"`
}

type ShippingPackageInfoType struct {
	XMLName                        xml.Name  `xml:"ShippingPackageInfo"`
	ActualDeliveryTime             time.Time `xml:",omitempty"`
	EstimatedDeliveryTimeMax       time.Time `xml:",omitempty"`
	EstimatedDeliveryTimeMin       time.Time `xml:",omitempty"`
	HandleByTime                   time.Time `xml:",omitempty"`
	MaxNativeEstimatedDeliveryTime time.Time `xml:",omitempty"`
	MinNativeEstimatedDeliveryTime time.Time `xml:",omitempty"`
	ScheduledDeliveryTimeMax       time.Time `xml:",omitempty"`
	ScheduledDeliveryTimeMin       time.Time `xml:",omitempty"`
	ShippingTrackingEvent          string    `xml:",omitempty"`
	StoreID                        string    `xml:",omitempty"`
}

type TransactionStatusType struct {
	XMLName                             xml.Name  `xml:"TransactionStatus"`
	BuyerSelectedShipping               bool      `xml:",omitempty"`
	CheckoutStatus                      string    `xml:",omitempty"`
	CompleteStatus                      string    `xml:",omitempty"`
	DigitalStatus                       string    `xml:",omitempty"`
	eBayPaymentStatus                   string    `xml:",omitempty"`
	InquiryStatus                       string    `xml:",omitempty"`
	IntegratedMerchantCreditCardEnabled bool      `xml:",omitempty"`
	LastTimeModified                    time.Time `xml:",omitempty"`
	PaymentHoldStatus                   string    `xml:",omitempty"`
	PaymentInstrument                   string    `xml:",omitempty"`
	PaymentMethodUsed                   string    `xml:",omitempty"`
	ReturnStatus                        string    `xml:",omitempty"`
}

type TaxesType struct {
	XMLName        xml.Name        `xml:"Taxes"`
	TaxDetails     *TaxDetailsType `xml:",omitempty"`
	TotalTaxAmount float64         `xml:",omitempty"`
}

type TaxDetailsType struct {
	XMLName             xml.Name `xml:"TaxDetails"`
	Imposition          string   `xml:",omitempty"`
	TaxAmount           float64  `xml:",omitempty"`
	TaxCode             string   `xml:",omitempty"`
	TaxDescription      string   `xml:",omitempty"`
	TaxOnHandlingAmount float64  `xml:",omitempty"`
	TaxOnShippingAmount float64  `xml:",omitempty"`
	TaxOnSubtotalAmount float64  `xml:",omitempty"`
}

type UnpaidItemType struct {
	XMLName xml.Name `xml:"UnpaidItem"`
	Status  string   `xml:",omitempty"`
	Type    string   `xml:",omitempty"`
}

type ASQPreferencesType struct {
	XMLName              xml.Name `xml:"ASQPreferences"`
	ResetDefaultSubjects bool     `xml:",omitempty"`
	Subject              string   `xml:",omitempty"`
}

type ItemListCustomizationType struct {
	XMLName        xml.Name        `xml:"ItemListCustomization"`
	DurationInDays int             `xml:",omitempty"`
	Include        bool            `xml:",omitempty"`
	IncludeNotes   bool            `xml:",omitempty"`
	Pagination     *PaginationType `xml:",omitempty"`
	Sort           string          `xml:",omitempty"`
}

type MyeBaySelectionType struct {
	XMLName                    xml.Name `xml:"MyeBaySelection"`
	Include                    bool     `xml:",omitempty"`
	IncludeFavoriteSellerCount bool     `xml:",omitempty"`
	IncludeItemCount           bool     `xml:",omitempty"`
	IncludeListContents        bool     `xml:",omitempty"`
	MaxResults                 int      `xml:",omitempty"`
	Sort                       string   `xml:",omitempty"`
	UserDefinedListName        string   `xml:",omitempty"`
}

type PaginatedItemArrayType struct {
	XMLName          xml.Name              `xml:"PaginatedItemArray"`
	ItemArray        []ItemArrayType       `xml:",omitempty"`
	PaginationResult *PaginationResultType `xml:",omitempty"`
}

type BiddingDetailsType struct {
	XMLName         xml.Name `xml:"BiddingDetails"`
	ConvertedMaxBid float64  `xml:",omitempty"`
	MaxBid          float64  `xml:",omitempty"`
	QuantityBid     int      `xml:",omitempty"`
	QuantityWon     int      `xml:",omitempty"`
}

type BuyingSummaryType struct {
	XMLName           xml.Name `xml:"BuyingSummary"`
	BestOfferCount    int      `xml:",omitempty"`
	BiddingCount      int      `xml:",omitempty"`
	TotalWinningCost  float64  `xml:",omitempty"`
	TotalWonCost      float64  `xml:",omitempty"`
	WinningCount      int      `xml:",omitempty"`
	WonCount          int      `xml:",omitempty"`
	WonDurationInDays int      `xml:",omitempty"`
}

type PaginatedOrderTransactionArrayType struct {
	XMLName               xml.Name                    `xml:"PaginatedOrderTransactionArray"`
	OrderTransactionArray []OrderTransactionArrayType `xml:",omitempty"`
	PaginationResult      *PaginationResultType       `xml:",omitempty"`
}

type OrderTransactionArrayType struct {
	XMLName          xml.Name              `xml:"OrderTransactionArray"`
	OrderTransaction *OrderTransactionType `xml:",omitempty"`
}

type OrderTransactionType struct {
	XMLName     xml.Name         `xml:"OrderTransaction"`
	Order       *OrderType       `xml:",omitempty"`
	Transaction *TransactionType `xml:",omitempty"`
}

type MyeBayFavoriteSearchListType struct {
	XMLName        xml.Name                  `xml:"MyeBayFavoriteSearchList"`
	FavoriteSearch *MyeBayFavoriteSearchType `xml:",omitempty"`
	TotalAvailable int                       `xml:",omitempty"`
}

type MyeBayFavoriteSearchType struct {
	XMLName            xml.Name  `xml:"MyeBayFavoriteSearch"`
	BidCountMax        int       `xml:",omitempty"`
	BidCountMin        int       `xml:",omitempty"`
	CategoryID         string    `xml:",omitempty"`
	Condition          string    `xml:",omitempty"`
	Currency           string    `xml:",omitempty"`
	EndTimeFrom        time.Time `xml:",omitempty"`
	EndTimeTo          time.Time `xml:",omitempty"`
	ItemsAvailableTo   string    `xml:",omitempty"`
	ItemsLocatedIn     string    `xml:",omitempty"`
	ItemSort           string    `xml:",omitempty"`
	ItemType           string    `xml:",omitempty"`
	MaxDistance        int       `xml:",omitempty"`
	PaymentMethod      string    `xml:",omitempty"`
	PostalCode         string    `xml:",omitempty"`
	PreferredLocation  string    `xml:",omitempty"`
	PriceMax           float64   `xml:",omitempty"`
	PriceMin           float64   `xml:",omitempty"`
	Quantity           int       `xml:",omitempty"`
	QuantityOperator   string    `xml:",omitempty"`
	QueryKeywords      string    `xml:",omitempty"`
	SearchFlag         string    `xml:",omitempty"`
	SearchName         string    `xml:",omitempty"`
	SearchQuery        string    `xml:",omitempty"`
	SellerBusinessType string    `xml:",omitempty"`
	SellerID           string    `xml:",omitempty"`
	SellerIDExclude    string    `xml:",omitempty"`
	SortOrder          string    `xml:",omitempty"`
}

type MyeBayFavoriteSellerListType struct {
	XMLName        xml.Name                  `xml:"MyeBayFavoriteSellerList"`
	FavoriteSeller *MyeBayFavoriteSellerType `xml:",omitempty"`
	TotalAvailable int                       `xml:",omitempty"`
}

type MyeBayFavoriteSellerType struct {
	XMLName   xml.Name `xml:"MyeBayFavoriteSeller"`
	StoreName string   `xml:",omitempty"`
	UserID    string   `xml:",omitempty"`
}

type UserDefinedListType struct {
	XMLName             xml.Name                      `xml:"UserDefinedList"`
	FavoriteSearches    *MyeBayFavoriteSearchListType `xml:",omitempty"`
	FavoriteSellerCount int                           `xml:",omitempty"`
	FavoriteSellers     *MyeBayFavoriteSellerListType `xml:",omitempty"`
	ItemArray           []ItemArrayType               `xml:",omitempty"`
	ItemCount           int                           `xml:",omitempty"`
	Name                string                        `xml:",omitempty"`
}

type ReminderCustomizationType struct {
	XMLName        xml.Name `xml:"ReminderCustomization"`
	DurationInDays int      `xml:",omitempty"`
	Include        bool     `xml:",omitempty"`
}

type RemindersType struct {
	XMLName                               xml.Name `xml:"Reminders"`
	DeclinedRTERequestCount               int      `xml:",omitempty"`
	DocsForCCProcessingToSendCount        int      `xml:",omitempty"`
	FeedbackToReceiveCount                int      `xml:",omitempty"`
	FeedbackToSendCount                   int      `xml:",omitempty"`
	ItemReceiptConfirmationToReceiveCount int      `xml:",omitempty"`
	ItemReceiptToConfirmCount             int      `xml:",omitempty"`
	PaymentToReceiveCount                 int      `xml:",omitempty"`
	PendingRTERequestCount                int      `xml:",omitempty"`
	RefundCancelledCount                  int      `xml:",omitempty"`
	RefundInitiatedCount                  int      `xml:",omitempty"`
	RefundOnHoldCount                     int      `xml:",omitempty"`
	RelistingNeededCount                  int      `xml:",omitempty"`
	RTEToProcessCount                     int      `xml:",omitempty"`
	SecondChanceOfferCount                int      `xml:",omitempty"`
	ShippingDetailsToBeProvidedCount      int      `xml:",omitempty"`
	ShippingNeededCount                   int      `xml:",omitempty"`
	TotalNewLeadsCount                    int      `xml:",omitempty"`
}

type SellingSummaryType struct {
	XMLName                  xml.Name `xml:"SellingSummary"`
	ActiveAuctionCount       int      `xml:",omitempty"`
	AuctionBidCount          int      `xml:",omitempty"`
	AuctionSellingCount      int      `xml:",omitempty"`
	SoldDurationInDays       int      `xml:",omitempty"`
	TotalAuctionSellingValue float64  `xml:",omitempty"`
	TotalSoldCount           int      `xml:",omitempty"`
	TotalSoldValue           float64  `xml:",omitempty"`
}

type MyeBaySellingSummaryType struct {
	XMLName                  xml.Name `xml:"MyeBaySellingSummary"`
	ActiveAuctionCount       int      `xml:",omitempty"`
	AmountLimitRemaining     float64  `xml:",omitempty"`
	AuctionBidCount          int      `xml:",omitempty"`
	AuctionSellingCount      int      `xml:",omitempty"`
	ClassifiedAdCount        int      `xml:",omitempty"`
	ClassifiedAdOfferCount   int      `xml:",omitempty"`
	QuantityLimitRemaining   float64  `xml:",omitempty"`
	SoldDurationInDays       int      `xml:",omitempty"`
	TotalAuctionSellingValue float64  `xml:",omitempty"`
	TotalLeadCount           int      `xml:",omitempty"`
	TotalListingsWithLeads   int      `xml:",omitempty"`
	TotalSoldCount           int      `xml:",omitempty"`
	TotalSoldValue           float64  `xml:",omitempty"`
}

type MyMessagesExternalMessageIDArrayType struct {
	XMLName           xml.Name `xml:"MyMessagesExternalMessageIDArray"`
	ExternalMessageID string   `xml:",omitempty"`
}

type MyMessagesMessageArrayType struct {
	XMLName xml.Name               `xml:"MyMessagesMessageArray"`
	Message *MyMessagesMessageType `xml:",omitempty"`
}

type MyMessagesMessageType struct {
	XMLName           xml.Name                       `xml:"MyMessagesMessage"`
	Content           string                         `xml:",omitempty"`
	ExpirationDate    time.Time                      `xml:",omitempty"`
	ExternalMessageID string                         `xml:",omitempty"`
	Flagged           bool                           `xml:",omitempty"`
	Folder            *MyMessagesFolderType          `xml:",omitempty"`
	HighPriority      bool                           `xml:",omitempty"`
	ItemEndTime       time.Time                      `xml:",omitempty"`
	ItemID            string                         `xml:",omitempty"`
	ItemTitle         string                         `xml:",omitempty"`
	ListingStatus     string                         `xml:",omitempty"`
	MessageID         string                         `xml:",omitempty"`
	MessageMedia      *MessageMediaType              `xml:",omitempty"`
	MessageType       string                         `xml:",omitempty"`
	QuestionType      string                         `xml:",omitempty"`
	Read              bool                           `xml:",omitempty"`
	ReceiveDate       time.Time                      `xml:",omitempty"`
	RecipientUserID   string                         `xml:",omitempty"`
	Replied           bool                           `xml:",omitempty"`
	ResponseDetails   *MyMessagesResponseDetailsType `xml:",omitempty"`
	Sender            string                         `xml:",omitempty"`
	SendToName        string                         `xml:",omitempty"`
	Subject           string                         `xml:",omitempty"`
	Text              string                         `xml:",omitempty"`
}

type MyMessagesFolderType struct {
	XMLName  xml.Name `xml:"MyMessagesFolder"`
	FolderID float64  `xml:",omitempty"`
}

type MyMessagesResponseDetailsType struct {
	XMLName         xml.Name `xml:"MyMessagesResponseDetails"`
	ResponseEnabled bool     `xml:",omitempty"`
	ResponseURL     string   `xml:",omitempty"`
}

type MyMessagesSummaryType struct {
	XMLName                xml.Name                     `xml:"MyMessagesSummary"`
	FlaggedMessageCount    int                          `xml:",omitempty"`
	FolderSummary          *MyMessagesFolderSummaryType `xml:",omitempty"`
	NewAlertCount          int                          `xml:",omitempty"`
	NewHighPriorityCount   int                          `xml:",omitempty"`
	NewMessageCount        int                          `xml:",omitempty"`
	TotalAlertCount        int                          `xml:",omitempty"`
	TotalHighPriorityCount int                          `xml:",omitempty"`
	TotalMessageCount      int                          `xml:",omitempty"`
	UnresolvedAlertCount   int                          `xml:",omitempty"`
}

type MyMessagesFolderSummaryType struct {
	XMLName                xml.Name `xml:"MyMessagesFolderSummary"`
	FolderID               float64  `xml:",omitempty"`
	FolderName             string   `xml:",omitempty"`
	NewAlertCount          int      `xml:",omitempty"`
	NewHighPriorityCount   int      `xml:",omitempty"`
	NewMessageCount        int      `xml:",omitempty"`
	TotalAlertCount        int      `xml:",omitempty"`
	TotalHighPriorityCount int      `xml:",omitempty"`
	TotalMessageCount      int      `xml:",omitempty"`
}

type ApplicationDeliveryPreferencesType struct {
	XMLName            xml.Name               `xml:"ApplicationDeliveryPreferences"`
	AlertEmail         string                 `xml:",omitempty"`
	AlertEnable        string                 `xml:",omitempty"`
	ApplicationEnable  string                 `xml:",omitempty"`
	ApplicationURL     string                 `xml:",omitempty"`
	DeliveryURLDetails *DeliveryURLDetailType `xml:",omitempty"`
	DeviceType         string                 `xml:",omitempty"`
	PayloadVersion     string                 `xml:",omitempty"`
}

type DeliveryURLDetailType struct {
	XMLName         xml.Name `xml:"DeliveryURLDetail"`
	DeliveryURL     string   `xml:",omitempty"`
	DeliveryURLName string   `xml:",omitempty"`
	Status          string   `xml:",omitempty"`
}

type NotificationEventPropertyType struct {
	XMLName   xml.Name `xml:"NotificationEventProperty"`
	EventType string   `xml:",omitempty"`
	Name      string   `xml:",omitempty"`
	Value     string   `xml:",omitempty"`
}

type NotificationUserDataType struct {
	XMLName          xml.Name                  `xml:"NotificationUserData"`
	ExternalUserData string                    `xml:",omitempty"`
	SMSSubscription  *SMSSubscriptionType      `xml:",omitempty"`
	SummarySchedule  *SummaryEventScheduleType `xml:",omitempty"`
}

type SMSSubscriptionType struct {
	XMLName           xml.Name `xml:"SMSSubscription"`
	CarrierID         string   `xml:",omitempty"`
	ErrorCode         string   `xml:",omitempty"`
	ItemToUnsubscribe string   `xml:",omitempty"`
	SMSPhone          string   `xml:",omitempty"`
	UserStatus        string   `xml:",omitempty"`
}

type SummaryEventScheduleType struct {
	XMLName       xml.Name `xml:"SummaryEventSchedule"`
	EventType     string   `xml:",omitempty"`
	Frequency     string   `xml:",omitempty"`
	SummaryPeriod string   `xml:",omitempty"`
}

type NotificationEnableArrayType struct {
	XMLName            xml.Name                `xml:"NotificationEnableArray"`
	NotificationEnable *NotificationEnableType `xml:",omitempty"`
}

type NotificationEnableType struct {
	XMLName     xml.Name `xml:"NotificationEnable"`
	EventEnable string   `xml:",omitempty"`
	EventType   string   `xml:",omitempty"`
}

type MarkUpMarkDownHistoryType struct {
	XMLName             xml.Name                 `xml:"MarkUpMarkDownHistory"`
	MarkUpMarkDownEvent *MarkUpMarkDownEventType `xml:",omitempty"`
}

type MarkUpMarkDownEventType struct {
	XMLName xml.Name  `xml:"MarkUpMarkDownEvent"`
	Reason  string    `xml:",omitempty"`
	Time    time.Time `xml:",omitempty"`
	Type    string    `xml:",omitempty"`
}

type NotificationDetailsArrayType struct {
	XMLName             xml.Name                 `xml:"NotificationDetailsArray"`
	NotificationDetails *NotificationDetailsType `xml:",omitempty"`
}

type NotificationDetailsType struct {
	XMLName         xml.Name  `xml:"NotificationDetails"`
	DeliveryStatus  string    `xml:",omitempty"`
	DeliveryTime    time.Time `xml:",omitempty"`
	DeliveryURL     string    `xml:",omitempty"`
	DeliveryURLName string    `xml:",omitempty"`
	ErrorMessage    string    `xml:",omitempty"`
	ExpirationTime  time.Time `xml:",omitempty"`
	NextRetryTime   time.Time `xml:",omitempty"`
	Retries         int       `xml:",omitempty"`
	Type            string    `xml:",omitempty"`
}

type NotificationStatisticsType struct {
	XMLName            xml.Name `xml:"NotificationStatistics"`
	DeliveredCount     int      `xml:",omitempty"`
	ErrorCount         int      `xml:",omitempty"`
	ExpiredCount       int      `xml:",omitempty"`
	QueuedNewCount     int      `xml:",omitempty"`
	QueuedPendingCount int      `xml:",omitempty"`
}

type OrderIDArrayType struct {
	XMLName xml.Name `xml:"OrderIDArray"`
	OrderID string   `xml:",omitempty"`
}

type OrderArrayType struct {
	XMLName xml.Name   `xml:"OrderArray"`
	Order   *OrderType `xml:",omitempty"`
}

type CheckoutStatusType struct {
	XMLName                             xml.Name  `xml:"CheckoutStatus"`
	eBayPaymentStatus                   string    `xml:",omitempty"`
	IntegratedMerchantCreditCardEnabled bool      `xml:",omitempty"`
	LastModifiedTime                    time.Time `xml:",omitempty"`
	PaymentInstrument                   string    `xml:",omitempty"`
	PaymentMethod                       string    `xml:",omitempty"`
	Status                              string    `xml:",omitempty"`
}

type ItemTransactionIDArrayType struct {
	XMLName           xml.Name               `xml:"ItemTransactionIDArray"`
	ItemTransactionID *ItemTransactionIDType `xml:",omitempty"`
}

type ItemTransactionIDType struct {
	XMLName         xml.Name `xml:"ItemTransactionID"`
	ItemID          string   `xml:",omitempty"`
	OrderLineItemID string   `xml:",omitempty"`
	SKU             string   `xml:",omitempty"`
	TransactionID   string   `xml:",omitempty"`
}

type PromotionalSaleArrayType struct {
	XMLName         xml.Name             `xml:"PromotionalSaleArray"`
	PromotionalSale *PromotionalSaleType `xml:",omitempty"`
}

type PromotionalSaleType struct {
	XMLName                  xml.Name  `xml:"PromotionalSale"`
	DiscountType             string    `xml:",omitempty"`
	DiscountValue            float64   `xml:",omitempty"`
	PromotionalSaleEndTime   time.Time `xml:",omitempty"`
	PromotionalSaleID        float64   `xml:",omitempty"`
	PromotionalSaleName      string    `xml:",omitempty"`
	PromotionalSaleStartTime time.Time `xml:",omitempty"`
	PromotionalSaleType      string    `xml:",omitempty"`
}

type ItemIDArrayType struct {
	XMLName xml.Name `xml:"ItemIDArray"`
	ItemID  string   `xml:",omitempty"`
}

type BuyerSatisfactionDashboardType struct {
	XMLName xml.Name                  `xml:"BuyerSatisfactionDashboard"`
	Alert   *SellerDashboardAlertType `xml:",omitempty"`
	Status  string                    `xml:",omitempty"`
}

type SellerDashboardAlertType struct {
	XMLName  xml.Name `xml:"SellerDashboardAlert"`
	Severity string   `xml:",omitempty"`
	Text     string   `xml:",omitempty"`
}

type PerformanceDashboardType struct {
	XMLName xml.Name                  `xml:"PerformanceDashboard"`
	Alert   *SellerDashboardAlertType `xml:",omitempty"`
	Site    string                    `xml:",omitempty"`
	Status  string                    `xml:",omitempty"`
}

type PowerSellerDashboardType struct {
	XMLName xml.Name                  `xml:"PowerSellerDashboard"`
	Alert   *SellerDashboardAlertType `xml:",omitempty"`
	Level   string                    `xml:",omitempty"`
}

type SearchStandingDashboardType struct {
	XMLName xml.Name `xml:"SearchStandingDashboard"`
	Status  string   `xml:",omitempty"`
}

type SellerAccountDashboardType struct {
	XMLName xml.Name                  `xml:"SellerAccountDashboard"`
	Alert   *SellerDashboardAlertType `xml:",omitempty"`
	Status  string                    `xml:",omitempty"`
}

type SellerFeeDiscountDashboardType struct {
	XMLName xml.Name `xml:"SellerFeeDiscountDashboard"`
	Percent float64  `xml:",omitempty"`
}

type UserIDArrayType struct {
	XMLName xml.Name `xml:"UserIDArray"`
	UserID  string   `xml:",omitempty"`
}

type SKUArrayType struct {
	XMLName xml.Name `xml:"SKUArray"`
	SKU     string   `xml:",omitempty"`
}

type SellingManagerAlertType struct {
	XMLName         xml.Name `xml:"SellingManagerAlert"`
	AlertType       string   `xml:",omitempty"`
	AutomationAlert string   `xml:",omitempty"`
	Count           float64  `xml:",omitempty"`
	DurationInDays  int      `xml:",omitempty"`
	GeneralAlert    string   `xml:",omitempty"`
	InventoryAlert  string   `xml:",omitempty"`
	PaisaPayAlert   string   `xml:",omitempty"`
	SoldAlert       string   `xml:",omitempty"`
}

type TimeRangeType struct {
	XMLName  xml.Name  `xml:"TimeRange"`
	TimeFrom time.Time `xml:",omitempty"`
	TimeTo   time.Time `xml:",omitempty"`
}

type SellingManagerEmailLogType struct {
	XMLName         xml.Name  `xml:"SellingManagerEmailLog"`
	CustomEmailName string    `xml:",omitempty"`
	EmailState      string    `xml:",omitempty"`
	EmailType       string    `xml:",omitempty"`
	EventTime       time.Time `xml:",omitempty"`
}

type SellingManagerSearchType struct {
	XMLName     xml.Name `xml:"SellingManagerSearch"`
	SearchType  string   `xml:",omitempty"`
	SearchValue string   `xml:",omitempty"`
}

type SellingManagerProductType struct {
	XMLName                              xml.Name                                  `xml:"SellingManagerProduct"`
	SellingManagerProductDetails         *SellingManagerProductDetailsType         `xml:",omitempty"`
	SellingManagerProductInventoryStatus *SellingManagerProductInventoryStatusType `xml:",omitempty"`
	SellingManagerProductSpecifics       *SellingManagerProductSpecificsType       `xml:",omitempty"`
	SellingManagerTemplateDetailsArray   []SellingManagerTemplateDetailsArrayType  `xml:",omitempty"`
}

type SellingManagerVendorDetailsType struct {
	XMLName           xml.Name `xml:"SellingManagerVendorDetails"`
	VendorContactInfo string   `xml:",omitempty"`
	VendorName        string   `xml:",omitempty"`
}

type SellingManagerProductInventoryStatusType struct {
	XMLName             xml.Name `xml:"SellingManagerProductInventoryStatus"`
	AverageSellingPrice float64  `xml:",omitempty"`
	QuantityActive      int      `xml:",omitempty"`
	QuantityScheduled   int      `xml:",omitempty"`
	QuantitySold        int      `xml:",omitempty"`
	QuantityUnsold      int      `xml:",omitempty"`
	SuccessPercent      float64  `xml:",omitempty"`
}

type SellingManagerTemplateDetailsArrayType struct {
	XMLName                       xml.Name                           `xml:"SellingManagerTemplateDetailsArray"`
	SellingManagerTemplateDetails *SellingManagerTemplateDetailsType `xml:",omitempty"`
}

type SellingManagerTemplateDetailsType struct {
	XMLName                      xml.Name                          `xml:"SellingManagerTemplateDetails"`
	SaleTemplateID               string                            `xml:",omitempty"`
	SaleTemplateName             string                            `xml:",omitempty"`
	SellingManagerProductDetails *SellingManagerProductDetailsType `xml:",omitempty"`
	SuccessPercent               float64                           `xml:",omitempty"`
	Template                     *ItemType                         `xml:",omitempty"`
}

type SellingManagerFolderDetailsType struct {
	XMLName  xml.Name `xml:"SellingManagerFolderDetails"`
	FolderID float64  `xml:",omitempty"`
}

type SellingManagerSoldOrderType struct {
	XMLName            xml.Name                       `xml:"SellingManagerSoldOrder"`
	ActualShippingCost float64                        `xml:",omitempty"`
	AdjustmentAmount   float64                        `xml:",omitempty"`
	CashOnDeliveryCost float64                        `xml:",omitempty"`
	ItemCost           float64                        `xml:",omitempty"`
	NotesToBuyer       string                         `xml:",omitempty"`
	NotesToSeller      string                         `xml:",omitempty"`
	OrderStatus        *SellingManagerOrderStatusType `xml:",omitempty"`
	ShippingAddress    *AddressType                   `xml:",omitempty"`
	ShippingDetails    *ShippingDetailsType           `xml:",omitempty"`
	VATRate            *VATRateType                   `xml:",omitempty"`
}

type SellingManagerOrderStatusType struct {
	XMLName           xml.Name `xml:"SellingManagerOrderStatus"`
	PaidStatus        string   `xml:",omitempty"`
	PaymentMethodUsed string   `xml:",omitempty"`
	ShippedStatus     string   `xml:",omitempty"`
}

type SellingManagerSoldTransactionType struct {
	XMLName               xml.Name `xml:"SellingManagerSoldTransaction"`
	CharityListing        bool     `xml:",omitempty"`
	CustomLabel           string   `xml:",omitempty"`
	ItemID                string   `xml:",omitempty"`
	ItemTitle             string   `xml:",omitempty"`
	ListedOn              string   `xml:",omitempty"`
	ListingType           string   `xml:",omitempty"`
	OrderLineItemID       string   `xml:",omitempty"`
	QuantitySold          int      `xml:",omitempty"`
	Relisted              bool     `xml:",omitempty"`
	ReservePrice          float64  `xml:",omitempty"`
	SaleRecordID          float64  `xml:",omitempty"`
	SecondChanceOfferSent bool     `xml:",omitempty"`
	SoldOn                string   `xml:",omitempty"`
	StartPrice            float64  `xml:",omitempty"`
	TransactionID         float64  `xml:",omitempty"`
	Variation             string   `xml:",omitempty"`
	WatchCount            int      `xml:",omitempty"`
}

type VATRateType struct {
	XMLName         xml.Name `xml:"VATRate"`
	ItemID          string   `xml:",omitempty"`
	OrderLineItemID string   `xml:",omitempty"`
	TransactionID   string   `xml:",omitempty"`
	VATPercent      float64  `xml:",omitempty"`
}

type CalculatedHandlingDiscountType struct {
	XMLName                  xml.Name `xml:"CalculatedHandlingDiscount"`
	DiscountName             string   `xml:",omitempty"`
	EachAdditionalAmount     float64  `xml:",omitempty"`
	EachAdditionalOffAmount  float64  `xml:",omitempty"`
	EachAdditionalPercentOff float64  `xml:",omitempty"`
	OrderHandlingAmount      float64  `xml:",omitempty"`
}

type ShippingInsuranceType struct {
	XMLName                    xml.Name                        `xml:"ShippingInsurance"`
	FlatRateInsuranceRangeCost *FlatRateInsuranceRangeCostType `xml:",omitempty"`
	InsuranceOption            string                          `xml:",omitempty"`
}

type FlatRateInsuranceRangeCostType struct {
	XMLName                xml.Name `xml:"FlatRateInsuranceRangeCost"`
	FlatRateInsuranceRange string   `xml:",omitempty"`
	InsuranceCost          float64  `xml:",omitempty"`
}

type StoreType struct {
	XMLName                xml.Name                      `xml:"Store"`
	CustomHeader           string                        `xml:",omitempty"`
	CustomHeaderLayout     string                        `xml:",omitempty"`
	CustomListingHeader    *StoreCustomListingHeaderType `xml:",omitempty"`
	Description            string                        `xml:",omitempty"`
	ExportListings         bool                          `xml:",omitempty"`
	HeaderStyle            string                        `xml:",omitempty"`
	HomePage               float64                       `xml:",omitempty"`
	ItemListLayout         string                        `xml:",omitempty"`
	ItemListSortOrder      string                        `xml:",omitempty"`
	LastOpenedTime         time.Time                     `xml:",omitempty"`
	Logo                   *StoreLogoType                `xml:",omitempty"`
	Name                   string                        `xml:",omitempty"`
	SubscriptionLevel      string                        `xml:",omitempty"`
	Theme                  *StoreThemeType               `xml:",omitempty"`
	TitleWithCompatibility bool                          `xml:",omitempty"`
	URL                    string                        `xml:",omitempty"`
	URLPath                string                        `xml:",omitempty"`
}

type StoreCustomCategoryArrayType struct {
	XMLName        xml.Name                 `xml:"StoreCustomCategoryArray"`
	CustomCategory *StoreCustomCategoryType `xml:",omitempty"`
}

type StoreCustomCategoryType struct {
	XMLName       xml.Name `xml:"StoreCustomCategory"`
	CategoryID    float64  `xml:",omitempty"`
	ChildCategory string   `xml:",omitempty"`
	Name          string   `xml:",omitempty"`
	Order         int      `xml:",omitempty"`
}

type StoreCustomListingHeaderType struct {
	XMLName                  xml.Name                          `xml:"StoreCustomListingHeader"`
	AddToFavoriteStores      bool                              `xml:",omitempty"`
	DisplayType              string                            `xml:",omitempty"`
	LinkToInclude            *StoreCustomListingHeaderLinkType `xml:",omitempty"`
	Logo                     bool                              `xml:",omitempty"`
	SearchBox                bool                              `xml:",omitempty"`
	SignUpForStoreNewsletter bool                              `xml:",omitempty"`
}

type StoreCustomListingHeaderLinkType struct {
	XMLName  xml.Name `xml:"StoreCustomListingHeaderLink"`
	LinkID   int      `xml:",omitempty"`
	LinkType string   `xml:",omitempty"`
	Order    int      `xml:",omitempty"`
}

type StoreLogoType struct {
	XMLName xml.Name `xml:"StoreLogo"`
	LogoID  int      `xml:",omitempty"`
	Name    string   `xml:",omitempty"`
	URL     string   `xml:",omitempty"`
}

type StoreThemeType struct {
	XMLName     xml.Name              `xml:"StoreTheme"`
	ColorScheme *StoreColorSchemeType `xml:",omitempty"`
	Name        string                `xml:",omitempty"`
	ThemeID     int                   `xml:",omitempty"`
}

type StoreColorSchemeType struct {
	XMLName       xml.Name        `xml:"StoreColorScheme"`
	Color         *StoreColorType `xml:",omitempty"`
	ColorSchemeID int             `xml:",omitempty"`
	Font          *StoreFontType  `xml:",omitempty"`
	Name          string          `xml:",omitempty"`
}

type StoreColorType struct {
	XMLName   xml.Name `xml:"StoreColor"`
	Accent    string   `xml:",omitempty"`
	Primary   string   `xml:",omitempty"`
	Secondary string   `xml:",omitempty"`
}

type StoreFontType struct {
	XMLName    xml.Name `xml:"StoreFont"`
	DescColor  string   `xml:",omitempty"`
	DescFace   string   `xml:",omitempty"`
	DescSize   string   `xml:",omitempty"`
	NameColor  string   `xml:",omitempty"`
	NameFace   string   `xml:",omitempty"`
	NameSize   string   `xml:",omitempty"`
	TitleColor string   `xml:",omitempty"`
	TitleFace  string   `xml:",omitempty"`
	TitleSize  string   `xml:",omitempty"`
}

type StoreCustomPageArrayType struct {
	XMLName    xml.Name             `xml:"StoreCustomPageArray"`
	CustomPage *StoreCustomPageType `xml:",omitempty"`
}

type StoreCustomPageType struct {
	XMLName        xml.Name `xml:"StoreCustomPage"`
	Content        string   `xml:",omitempty"`
	LeftNav        bool     `xml:",omitempty"`
	Name           string   `xml:",omitempty"`
	Order          int      `xml:",omitempty"`
	PageID         float64  `xml:",omitempty"`
	PreviewEnabled bool     `xml:",omitempty"`
	Status         string   `xml:",omitempty"`
	URLPath        string   `xml:",omitempty"`
}

type StoreThemeArrayType struct {
	XMLName xml.Name        `xml:"StoreThemeArray"`
	Theme   *StoreThemeType `xml:",omitempty"`
}

type StoreColorSchemeArrayType struct {
	XMLName     xml.Name              `xml:"StoreColorSchemeArray"`
	ColorScheme *StoreColorSchemeType `xml:",omitempty"`
}

type StoreLogoArrayType struct {
	XMLName xml.Name       `xml:"StoreLogoArray"`
	Logo    *StoreLogoType `xml:",omitempty"`
}

type StoreSubscriptionArrayType struct {
	XMLName      xml.Name               `xml:"StoreSubscriptionArray"`
	Subscription *StoreSubscriptionType `xml:",omitempty"`
}

type StoreSubscriptionType struct {
	XMLName xml.Name `xml:"StoreSubscription"`
	Fee     float64  `xml:",omitempty"`
	Level   string   `xml:",omitempty"`
}

type StorePreferencesType struct {
	XMLName             xml.Name                      `xml:"StorePreferences"`
	VacationPreferences *StoreVacationPreferencesType `xml:",omitempty"`
}

type StoreVacationPreferencesType struct {
	XMLName                       xml.Name  `xml:"StoreVacationPreferences"`
	DisplayMessageStoreCustomText bool      `xml:",omitempty"`
	MessageItem                   bool      `xml:",omitempty"`
	MessageStore                  bool      `xml:",omitempty"`
	MessageStoreCustomText        string    `xml:",omitempty"`
	OnVacation                    bool      `xml:",omitempty"`
	ReturnDate                    time.Time `xml:",omitempty"`
}

type SuggestedCategoryArrayType struct {
	XMLName           xml.Name               `xml:"SuggestedCategoryArray"`
	SuggestedCategory *SuggestedCategoryType `xml:",omitempty"`
}

type SuggestedCategoryType struct {
	XMLName          xml.Name      `xml:"SuggestedCategory"`
	Category         *CategoryType `xml:",omitempty"`
	PercentItemFound int           `xml:",omitempty"`
}

type TokenStatusType struct {
	XMLName        xml.Name  `xml:"TokenStatus"`
	EIASToken      string    `xml:",omitempty"`
	ExpirationTime time.Time `xml:",omitempty"`
	RevocationTime time.Time `xml:",omitempty"`
	Status         string    `xml:",omitempty"`
}

type MembershipDetailsType struct {
	XMLName xml.Name              `xml:"MembershipDetails"`
	Program *MembershipDetailType `xml:",omitempty"`
}

type MembershipDetailType struct {
	XMLName     xml.Name  `xml:"MembershipDetail"`
	ExpiryDate  time.Time `xml:",omitempty"`
	ProgramName string    `xml:",omitempty"`
	Site        string    `xml:",omitempty"`
}

type CharityAffiliationDetailsType struct {
	XMLName                  xml.Name                      `xml:"CharityAffiliationDetails"`
	CharityAffiliationDetail *CharityAffiliationDetailType `xml:",omitempty"`
}

type CharityAffiliationDetailType struct {
	XMLName         xml.Name  `xml:"CharityAffiliationDetail"`
	AffiliationType string    `xml:",omitempty"`
	CharityID       string    `xml:",omitempty"`
	LastUsedTime    time.Time `xml:",omitempty"`
}

type FeatureEligibilityType struct {
	XMLName                              xml.Name `xml:"FeatureEligibility"`
	QualifiedForAuctionOneDayDuration    bool     `xml:",omitempty"`
	QualifiedForFixedPriceOneDayDuration bool     `xml:",omitempty"`
	QualifiesForBuyItNow                 bool     `xml:",omitempty"`
	QualifiesForBuyItNowMultiple         bool     `xml:",omitempty"`
	QualifiesForVariations               bool     `xml:",omitempty"`
}

type IntegratedMerchantCreditCardInfoType struct {
	XMLName       xml.Name `xml:"IntegratedMerchantCreditCardInfo"`
	SupportedSite string   `xml:",omitempty"`
}

type RecoupmentPolicyConsentType struct {
	XMLName xml.Name `xml:"RecoupmentPolicyConsent"`
	Site    string   `xml:",omitempty"`
}

type SchedulingInfoType struct {
	XMLName             xml.Name `xml:"SchedulingInfo"`
	MaxScheduledItems   int      `xml:",omitempty"`
	MaxScheduledMinutes int      `xml:",omitempty"`
	MinScheduledMinutes int      `xml:",omitempty"`
}

type TopRatedSellerDetailsType struct {
	XMLName         xml.Name `xml:"TopRatedSellerDetails"`
	TopRatedProgram string   `xml:",omitempty"`
}

type DisputeArrayType struct {
	XMLName xml.Name     `xml:"DisputeArray"`
	Dispute *DisputeType `xml:",omitempty"`
}

type DisputeResolutionType struct {
	XMLName                     xml.Name  `xml:"DisputeResolution"`
	DisputeResolutionReason     string    `xml:",omitempty"`
	DisputeResolutionRecordType string    `xml:",omitempty"`
	ResolutionTime              time.Time `xml:",omitempty"`
}

type DisputeFilterCountType struct {
	XMLName           xml.Name `xml:"DisputeFilterCount"`
	DisputeFilterType string   `xml:",omitempty"`
	TotalAvailable    int      `xml:",omitempty"`
}

type BidderNoticePreferencesType struct {
	XMLName                                xml.Name `xml:"BidderNoticePreferences"`
	UnsuccessfulBidderNoticeIncludeMyItems bool     `xml:",omitempty"`
}

type CombinedPaymentPreferencesType struct {
	XMLName               xml.Name `xml:"CombinedPaymentPreferences"`
	CombinedPaymentOption string   `xml:",omitempty"`
}

type DispatchCutoffTimePreferencesType struct {
	XMLName    xml.Name  `xml:"DispatchCutoffTimePreferences"`
	CutoffTime time.Time `xml:",omitempty"`
}

type eBayPLUSPreferenceType struct {
	XMLName           xml.Name `xml:"eBayPLUSPreference"`
	Country           string   `xml:",omitempty"`
	ListingPreference bool     `xml:",omitempty"`
	OptInStatus       bool     `xml:",omitempty"`
}

type EndOfAuctionEmailPreferencesType struct {
	XMLName         xml.Name `xml:"EndOfAuctionEmailPreferences"`
	EmailCustomized bool     `xml:",omitempty"`
	LogoCustomized  bool     `xml:",omitempty"`
	LogoType        string   `xml:",omitempty"`
	LogoURL         string   `xml:",omitempty"`
	TemplateText    string   `xml:",omitempty"`
	TextCustomized  bool     `xml:",omitempty"`
}

type PurchaseReminderEmailPreferencesType struct {
	XMLName                          xml.Name `xml:"PurchaseReminderEmailPreferences"`
	PurchaseReminderEmailPreferences bool     `xml:",omitempty"`
}

type SellerExcludeShipToLocationPreferencesType struct {
	XMLName               xml.Name `xml:"SellerExcludeShipToLocationPreferences"`
	ExcludeShipToLocation string   `xml:",omitempty"`
}

type SellerFavoriteItemPreferencesType struct {
	XMLName         xml.Name `xml:"SellerFavoriteItemPreferences"`
	FavoriteItemID  string   `xml:",omitempty"`
	ListingType     string   `xml:",omitempty"`
	MaxPrice        float64  `xml:",omitempty"`
	MinPrice        float64  `xml:",omitempty"`
	SearchKeywords  string   `xml:",omitempty"`
	SearchSortOrder string   `xml:",omitempty"`
	StoreCategoryID float64  `xml:",omitempty"`
}

type SellerPaymentPreferencesType struct {
	XMLName                     xml.Name     `xml:"SellerPaymentPreferences"`
	AlwaysUseThisPaymentAddress bool         `xml:",omitempty"`
	DefaultPayPalEmailAddress   string       `xml:",omitempty"`
	DisplayPayNowButton         string       `xml:",omitempty"`
	FedExRateOption             string       `xml:",omitempty"`
	PayPalAlwaysOn              bool         `xml:",omitempty"`
	PayPalPreferred             bool         `xml:",omitempty"`
	SellerPaymentAddress        *AddressType `xml:",omitempty"`
	UPSRateOption               string       `xml:",omitempty"`
	USPSRateOption              string       `xml:",omitempty"`
}

type SellerProfilePreferencesType struct {
	XMLName                 xml.Name                     `xml:"SellerProfilePreferences"`
	SellerProfileOptedIn    bool                         `xml:",omitempty"`
	SupportedSellerProfiles *SupportedSellerProfilesType `xml:",omitempty"`
}

type SupportedSellerProfilesType struct {
	XMLName                xml.Name                    `xml:"SupportedSellerProfiles"`
	SupportedSellerProfile *SupportedSellerProfileType `xml:",omitempty"`
}

type SupportedSellerProfileType struct {
	XMLName       xml.Name           `xml:"SupportedSellerProfile"`
	CategoryGroup *CategoryGroupType `xml:",omitempty"`
	ProfileID     float64            `xml:",omitempty"`
	ProfileName   string             `xml:",omitempty"`
	ProfileType   string             `xml:",omitempty"`
	ShortSummary  string             `xml:",omitempty"`
}

type CategoryGroupType struct {
	XMLName   xml.Name `xml:"CategoryGroup"`
	IsDefault bool     `xml:",omitempty"`
	Name      string   `xml:",omitempty"`
}

type SellerReturnPreferencesType struct {
	XMLName xml.Name `xml:"SellerReturnPreferences"`
	OptedIn bool     `xml:",omitempty"`
}

type UnpaidItemAssistancePreferencesType struct {
	XMLName                   xml.Name `xml:"UnpaidItemAssistancePreferences"`
	AutoOptDonationRefund     bool     `xml:",omitempty"`
	AutoRelist                bool     `xml:",omitempty"`
	DelayBeforeOpeningDispute int      `xml:",omitempty"`
	ExcludedUser              string   `xml:",omitempty"`
	OptInStatus               bool     `xml:",omitempty"`
	RemoveAllExcludedUsers    bool     `xml:",omitempty"`
}

type VeROReasonCodeDetailsType struct {
	XMLName        xml.Name            `xml:"VeROReasonCodeDetails"`
	VeROSiteDetail *VeROSiteDetailType `xml:",omitempty"`
}

type VeROSiteDetailType struct {
	XMLName          xml.Name              `xml:"VeROSiteDetail"`
	ReasonCodeDetail *ReasonCodeDetailType `xml:",omitempty"`
	Site             string                `xml:",omitempty"`
}

type ReasonCodeDetailType struct {
	XMLName      xml.Name `xml:"ReasonCodeDetail"`
	BriefText    string   `xml:",omitempty"`
	DetailedText string   `xml:",omitempty"`
}

type VeROReportedItemDetailsType struct {
	XMLName      xml.Name              `xml:"VeROReportedItemDetails"`
	ReportedItem *VeROReportedItemType `xml:",omitempty"`
}

type VeROReportedItemType struct {
	XMLName              xml.Name `xml:"VeROReportedItem"`
	ItemID               string   `xml:",omitempty"`
	ItemReasonForFailure string   `xml:",omitempty"`
	ItemStatus           string   `xml:",omitempty"`
}

type ItemRatingDetailArrayType struct {
	XMLName           xml.Name               `xml:"ItemRatingDetailArray"`
	ItemRatingDetails *ItemRatingDetailsType `xml:",omitempty"`
}

type ItemRatingDetailsType struct {
	XMLName      xml.Name `xml:"ItemRatingDetails"`
	Rating       int      `xml:",omitempty"`
	RatingDetail string   `xml:",omitempty"`
}

type AffiliateTrackingDetailsType struct {
	XMLName               xml.Name `xml:"AffiliateTrackingDetails"`
	AffiliateUserID       string   `xml:",omitempty"`
	ApplicationDeviceType string   `xml:",omitempty"`
	TrackingID            string   `xml:",omitempty"`
	TrackingPartnerCode   string   `xml:",omitempty"`
}

type BotBlockRequestType struct {
	XMLName           xml.Name `xml:"BotBlockRequest"`
	BotBlockToken     string   `xml:",omitempty"`
	BotBlockUserInput string   `xml:",omitempty"`
}

type SuggestedBidValueType struct {
	XMLName  xml.Name `xml:"SuggestedBidValue"`
	BidValue float64  `xml:",omitempty"`
}

type BotBlockResponseType struct {
	XMLName          xml.Name `xml:"BotBlockResponse"`
	BotBlockAudioUrl string   `xml:",omitempty"`
	BotBlockToken    string   `xml:",omitempty"`
	BotBlockUrl      string   `xml:",omitempty"`
}

type ModifyNameArrayType struct {
	XMLName    xml.Name        `xml:"ModifyNameArray"`
	ModifyName *ModifyNameType `xml:",omitempty"`
}

type ModifyNameType struct {
	XMLName xml.Name `xml:"ModifyName"`
	Name    string   `xml:",omitempty"`
	NewName string   `xml:",omitempty"`
}

type InventoryStatusType struct {
	XMLName    xml.Name `xml:"InventoryStatus"`
	ItemID     string   `xml:",omitempty"`
	Quantity   int      `xml:",omitempty"`
	SKU        string   `xml:",omitempty"`
	StartPrice float64  `xml:",omitempty"`
}

type InventoryFeesType struct {
	XMLName xml.Name `xml:"InventoryFees"`
	Fee     *FeeType `xml:",omitempty"`
	ItemID  string   `xml:",omitempty"`
}

type FeedbackCommentArrayType struct {
	XMLName           xml.Name `xml:"FeedbackCommentArray"`
	StoredCommentText string   `xml:",omitempty"`
}

type SiteHostedPictureDetailsType struct {
	XMLName            xml.Name              `xml:"SiteHostedPictureDetails"`
	BaseURL            string                `xml:",omitempty"`
	ExternalPictureURL string                `xml:",omitempty"`
	FullURL            string                `xml:",omitempty"`
	PictureFormat      string                `xml:",omitempty"`
	PictureName        string                `xml:",omitempty"`
	PictureSet         string                `xml:",omitempty"`
	PictureSetMember   *PictureSetMemberType `xml:",omitempty"`
	UseByDate          time.Time             `xml:",omitempty"`
}

type PictureSetMemberType struct {
	XMLName       xml.Name `xml:"PictureSetMember"`
	MemberURL     string   `xml:",omitempty"`
	PictureHeight int      `xml:",omitempty"`
	PictureWidth  int      `xml:",omitempty"`
}

type VeROReportItemsType struct {
	XMLName    xml.Name            `xml:"VeROReportItems"`
	ReportItem *VeROReportItemType `xml:",omitempty"`
}

type VeROReportItemType struct {
	XMLName                xml.Name `xml:"VeROReportItem"`
	CopyEmailToRightsOwner bool     `xml:",omitempty"`
	Country                string   `xml:",omitempty"`
	DetailedMessage        string   `xml:",omitempty"`
	ItemID                 string   `xml:",omitempty"`
	MessageToSeller        string   `xml:",omitempty"`
	Patent                 string   `xml:",omitempty"`
	Region                 string   `xml:",omitempty"`
	VeROReasonCodeID       float64  `xml:",omitempty"`
}

type RequesterCredentialsType struct {
	EbayAuthToken string `xml:"eBayAuthToken"`
}
