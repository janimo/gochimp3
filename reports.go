package gochimp3

import (
	"errors"
	"fmt"
)

const (
	reports_path       = "/reports"
	single_report_path = reports_path + "/%s"
)

type ReportQueryParams struct {
	ExtendedQueryParams

	Type           string
	BeforeSendTime string
	SinceSendTime  string
}

func (q ReportQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()
	m["type"] = q.Type
	m["before_send_time"] = q.BeforeSendTime
	m["since_send_time"] = q.SinceSendTime
	return m
}

type ReportResponseBounces struct {
	HardBounces  int `json:"hard_bounces"`
	SoftBounces  int `json:"soft_bounces"`
	SyntaxErrors int `json:"syntax_errors"`
}

type ReportResponseForwards struct {
	ForwardsCount int `json:"forwards_count"`
	ForwardsOpens int `json:"forwards_opens"`
}

type ReportResponseOpens struct {
	OpensTotal  int     `json:"opens_total"`
	UniqueOpens int     `json:"unique_opens`
	OpenRate    float64 `json:"open_rate"`
	LastOpen    string  `json:"last_open"`
}

type ReportResponseClicks struct {
	ClicksTotal            int     `json:"clicks_total"`
	UniqueClicks           int     `json:"unique_clicks`
	UniqueSubscriberClicks int     `json:"unique_subscriber_clicks`
	ClickRate              float64 `json:"click_rate"`
	LastClick              string  `json:"last_click"`
}

type ReportFacebookLikes struct {
	RecipientLikes int `json:"recipient_likes"`
	UniqueLikes    int `json:"unique_likes"`
	FacebookLikes  int `json:"facebook_likes"`
}

type ReportIndustryStats struct {
	Type       string  `json:"type"`
	OpenRate   float64 `json:"open_rate"`
	ClickRate  float64 `json:"click_rate"`
	BounceRate float64 `json:"bounce_rate"`
	UnopenRate float64 `json:"unopen_rate"`
	UnsubRate  float64 `json:"unsub_rate"`
	AbuseRate  float64 `json:"abuse_rate"`
}

type ReportEcommerce struct {
	TotalOrders  int    `json:"total_orders"`
	TotalSpent   int    `json:"total_spent"`
	TotalRevenue int    `json:"total_revenue"`
	CurrencyCode string `json:"currency_code"`
}

type ReportDeliveryStatus struct {
	Enabled        bool   `json:"enabled"`
	CanCancel      bool   `json:"can_cancel"`
	Status         string `json:"status"`
	EmailsSent     int    `json:"emails_sent"`
	EmailsCanceled int    `json:"emails_canceled"`
}

type ReportListStats struct {
	SubRate   float64 `json:"sub_rate"`
	UnsubRate float64 `json:"unsub_rate"`
	OpenRate  float64 `json:"open_rate"`
	ClickRate float64 `json:"click_rate"`
}

type ReportResponse struct {
	withLinks

	ID             string                 `json:"id"`
	CampaignTitle  string                 `json:"campaign_title"`
	Type           string                 `json:"type"`
	ListID         string                 `json:"list_id"`
	ListIsActive   bool                   `json:"list_is_active"`
	ListName       string                 `json:"list_name"`
	SubjectLine    string                 `json:"subject_line"`
	PreviewText    string                 `json:"preview_text"`
	EmailsSent     int                    `json:"emails_sent"`
	AbuseReports   int                    `json:"abuse_reports"`
	Unsubscribed   int                    `json:"unsubscribed"`
	SendTime       string                 `json:"send_time"`
	RSSLastSend    string                 `json:"rss_last_send"`
	Bounces        ReportResponseBounces  `json:"bounces"`
	Forwards       ReportResponseForwards `json:"forwards"`
	Opens          ReportResponseOpens    `json:"opens"`
	Clicks         ReportResponseClicks   `json:"clicks"`
	FacebookLikes  ReportFacebookLikes    `json:"facebook_likes"`
	IndustryStats  ReportIndustryStats    `json:"industry_stats"`
	ListStats      ReportListStats        `json:"list_stats"`
	Ecommerce      ReportEcommerce        `json:"ecommerce"`
	DeliveryStatus ReportDeliveryStatus   `json:"delivery_status"`
	api            *API
}

type ListOfReports struct {
	baseList
	Reports []ReportResponse `json:"reports"`
}

func (report ReportResponse) CanMakeRequest() error {
	if report.ID == "" {
		return errors.New("No ID provided on report")
	}

	return nil
}
func (api *API) GetReports(params *ReportQueryParams) (*ListOfReports, error) {
	response := new(ListOfReports)

	err := api.Request("GET", reports_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, l := range response.Reports {
		l.api = api
	}

	return response, nil
}

func (api *API) GetReport(id string, params *BasicQueryParams) (*ReportResponse, error) {
	endpoint := fmt.Sprintf(single_report_path, id)

	response := new(ReportResponse)
	response.api = api

	return response, api.Request("GET", endpoint, params, nil, response)
}
