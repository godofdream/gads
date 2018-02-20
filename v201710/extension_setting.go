package v201710

import (
	"encoding/xml"
	"fmt"
)

// https://developers.google.com/adwords/api/docs/reference/v201710/AdGroupExtensionSettingService.ExtensionSetting
// A setting specifying when and which extensions should serve at a given level (customer, campaign, or ad group).
type ExtensionSetting struct {
	PlatformRestrictions ExtensionSettingPlatform `xml:"platformRestrictions,omitempty"`

	Extensions Extension `xml:"https://adwords.google.com/api/adwords/cm/v201710 extensions,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201710/AdGroupExtensionSettingService.ExtensionSetting.Platform
// Different levels of platform restrictions
// DESKTOP, MOBILE, NONE
type ExtensionSettingPlatform string

type Extension interface{}

// https://developers.google.com/adwords/api/docs/reference/v201710/AdGroupExtensionSettingService.ExtensionFeedItem
// Contains base extension feed item data for an extension in an extension feed managed by AdWords.
type ExtensionFeedItem struct {
	XMLName xml.Name `json:"-" xml:"extensions"`

	FeedId                  int64                      `xml:"https://adwords.google.com/api/adwords/cm/v201710 feedId,omitempty"`
	FeedItemId              int64                      `xml:"https://adwords.google.com/api/adwords/cm/v201710 feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"https://adwords.google.com/api/adwords/cm/v201710 status,omitempty"`
	FeedType                *FeedType                  `xml:"https://adwords.google.com/api/adwords/cm/v201710 feedType,omitempty"`
	StartTime               string                     `xml:"https://adwords.google.com/api/adwords/cm/v201710 startTime,omitempty"` //  special value "00000101 000000" may be used to clear an existing start time.
	EndTime                 string                     `xml:"https://adwords.google.com/api/adwords/cm/v201710 endTime,omitempty"`   //  special value "00000101 000000" may be used to clear an existing end time.
	DevicePreference        *FeedItemDevicePreference  `xml:"https://adwords.google.com/api/adwords/cm/v201710 devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"https://adwords.google.com/api/adwords/cm/v201710 scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"https://adwords.google.com/api/adwords/cm/v201710 campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"https://adwords.google.com/api/adwords/cm/v201710 adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"https://adwords.google.com/api/adwords/cm/v201710 keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"https://adwords.google.com/api/adwords/cm/v201710 geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"https://adwords.google.com/api/adwords/cm/v201710 geoTargetingRestriction,omitempty"`
	PolicyData              *[]FeedItemPolicyData      `xml:"https://adwords.google.com/api/adwords/cm/v201710 policyData,omitempty"`

	ExtensionFeedItemType string `xml:"https://adwords.google.com/api/adwords/cm/v201710 ExtensionFeedItem.Type,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201710/AdGroupExtensionSettingService.CallFeedItem
// Represents a Call extension.
type CallFeedItem struct {
	ExtensionFeedItem

	CallPhoneNumber               string             `xml:"https://adwords.google.com/api/adwords/cm/v201710 callPhoneNumber,omitempty"`
	CallCountryCode               string             `xml:"https://adwords.google.com/api/adwords/cm/v201710 callCountryCode,omitempty"`
	CallTracking                  bool               `xml:"https://adwords.google.com/api/adwords/cm/v201710 callTracking,omitempty"`
	CallConversionType            CallConversionType `xml:"https://adwords.google.com/api/adwords/cm/v201710 callConversionType,omitempty"`
	DisableCallConversionTracking bool               `xml:"https://adwords.google.com/api/adwords/cm/v201710 disableCallConversionTracking,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201710/AdGroupExtensionSettingService.CalloutFeedItem
// Represents a Callout extension.
type CalloutFeedItem struct {
	ExtensionFeedItem
	CalloutText string `xml:"https://adwords.google.com/api/adwords/cm/v201710 calloutText,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201710/AdGroupExtensionSettingService.SitelinkFeedItem
// Represents a Sitelink extension.
type SitelinkFeedItem struct {
	ExtensionFeedItem
	SitelinkText                string            `xml:"https://adwords.google.com/api/adwords/cm/v201710 sitelinkText,omitempty"`
	SitelinkURL                 string            `xml:"https://adwords.google.com/api/adwords/cm/v201710 sitelinkUrl,omitempty"`
	SitelinkLine2               string            `xml:"https://adwords.google.com/api/adwords/cm/v201710 sitelinkLine2,omitempty"`
	SitelinkLine3               string            `xml:"https://adwords.google.com/api/adwords/cm/v201710 sitelinkLine3,omitempty"`
	SitelinkFinalUrls           []string          `xml:"https://adwords.google.com/api/adwords/cm/v201710 sitelinkFinalUrls,omitempty"`
	SitelinkFinalMobileUrls     []string          `xml:"https://adwords.google.com/api/adwords/cm/v201710 sitelinkFinalMobileUrls,omitempty"`
	SitelinkTrackingUrlTemplate string            `xml:"https://adwords.google.com/api/adwords/cm/v201710 sitelinkTrackingUrlTemplate,omitempty"`
	SitelinkUrlCustomParameters *CustomParameters `xml:"https://adwords.google.com/api/adwords/cm/v201710 sitelinkUrlCustomParameters,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201710/AdGroupExtensionSettingService.StructuredSnippetFeedItem
// Represents a StructuredSnippet extension.
type StructuredSnippetFeedItem struct {
	ExtensionFeedItem
	Header string   `xml:"https://adwords.google.com/api/adwords/cm/v201710 header,omitempty"`
	Values []string `xml:"https://adwords.google.com/api/adwords/cm/v201710 values,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201710/AdGroupExtensionSettingService.PriceFeedItem
// Represents a PriceFeed extension.
type PriceFeedItem struct {
	ExtensionFeedItem
	// BRANDS EVENTS LOCATIONS NEIGHBORHOODS PRODUCT_CATEGORIES PRODUCT_TIERS SERVICES SERVICE_CATEGORIES SERVICE_TIERS
	PriceExtensionType string `xml:"https://adwords.google.com/api/adwords/cm/v201710 priceExtensionType,omitempty"`
	//UNKOWN FROM UP_TO AVERAGE NONE
	PriceQualifier      string          `xml:"https://adwords.google.com/api/adwords/cm/v201710 priceQualifier,omitempty"`
	TrackingUrlTemplate string          `xml:"https://adwords.google.com/api/adwords/cm/v201710 trackingUrlTemplate,omitempty"`
	Language            string          `xml:"https://adwords.google.com/api/adwords/cm/v201710 language,omitempty"`
	TableRows           []PriceTableRow `xml:"https://adwords.google.com/api/adwords/cm/v201710 tableRows,omitempty"`
	Values              []string        `xml:"https://adwords.google.com/api/adwords/cm/v201710 values,omitempty"`
}

type PriceTableRow struct {
	Header          string            `xml:"https://adwords.google.com/api/adwords/cm/v201710 header,omitempty"`
	Description     string            `xml:"https://adwords.google.com/api/adwords/cm/v201710 description,omitempty"`
	FinalUrls       []string          `xml:"https://adwords.google.com/api/adwords/cm/v201710 finalUrls,omitempty"`
	FinalMobileUrls []string          `xml:"https://adwords.google.com/api/adwords/cm/v201710 finalMobileUrls,omitempty"`
	Price           MoneyWithCurrency `xml:"https://adwords.google.com/api/adwords/cm/v201710 price,omitempty"`
	// UNKNOWN PER_HOUR PER_DAY PER_WEEK PER_MONTH PER_YEAR PER_NIGHT
	PriceUnit string `xml:"https://adwords.google.com/api/adwords/cm/v201710 priceUnit,omitempty"`
}

type MoneyWithCurrency struct {
	Amount       int64  `xml:"https://adwords.google.com/api/adwords/cm/v201710 money>microAmount,omitempty"`
	CurrencyCode string `xml:"https://adwords.google.com/api/adwords/cm/v201710 currencyCode,omitempty"`
}

func extensionsUnmarshalXML(dec *xml.Decoder, start xml.StartElement) (ext interface{}, err error) {
	extensionsType, err := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
	if err != nil {
		return
	}
	switch extensionsType {
	case "CallFeedItem":
		c := CallFeedItem{}
		err = dec.DecodeElement(&c, &start)
		ext = c
	case "PriceFeedItem":
		c := PriceFeedItem{}
		err = dec.DecodeElement(&c, &start)
		ext = c
	case "StructuredSnippetFeedItem":
		c := StructuredSnippetFeedItem{}
		err = dec.DecodeElement(&c, &start)
		ext = c
	case "SitelinkFeedItem":
		c := SitelinkFeedItem{}
		err = dec.DecodeElement(&c, &start)
		ext = c
	case "CalloutFeedItem":
		c := CalloutFeedItem{}
		err = dec.DecodeElement(&c, &start)
		ext = c

	default:
		err = fmt.Errorf("unknown Extensions type %#v", extensionsType)
	}
	return
}

func (s ExtensionSetting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if s.PlatformRestrictions != "NONE" {
		e.EncodeElement(&s.PlatformRestrictions, xml.StartElement{Name: xml.Name{
			"https://adwords.google.com/api/adwords/cm/v201710",
			"platformRestrictions"}})
	}
	switch extType := s.Extensions.(type) {
	case []CallFeedItem:
		e.EncodeElement(s.Extensions.([]CallFeedItem), xml.StartElement{
			xml.Name{baseUrl, "extensions"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "CallFeedItem"},
			},
		})
	default:
		return fmt.Errorf("unknown extension type %#v\n", extType)

	}

	e.EncodeToken(start.End())
	return nil
}

func (s *ExtensionSetting) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {
	s.Extensions = []interface{}{}

	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			switch start.Name.Local {
			case "platformRestrictions":
				if err := dec.DecodeElement(&s.PlatformRestrictions, &start); err != nil {
					return err
				}
			case "extensions":
				extension, err := extensionsUnmarshalXML(dec, start)
				if err != nil {
					return err
				}
				s.Extensions = append(s.Extensions.([]interface{}), extension)
			}
		}
	}
	return nil
}
