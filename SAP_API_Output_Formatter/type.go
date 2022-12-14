package sap_api_output_formatter

type Opportunity struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	OpportunityCode string `json:"opportunity_code"`
	Deleted         bool   `json:"deleted"`
}

type OpportunityCollection struct {
	ObjectID                                 string `json:"ObjectID"`
	ProcessingTypeCode                       string `json:"ProcessingTypeCode"`
	ID                                       string `json:"ID"`
	ExternalID                               string `json:"ExternalID"`
	UUID                                     string `json:"UUID"`
	ProspectPartyID                          string `json:"ProspectPartyID"`
	Name                                     string `json:"Name"`
	PrimaryContactPartyID                    string `json:"PrimaryContactPartyID"`
	PriorityCode                             string `json:"PriorityCode"`
	OriginTypeCode                           string `json:"OriginTypeCode"`
	LifeCycleStatusCode                      string `json:"LifeCycleStatusCode"`
	ExternalUserStatusCode                   string `json:"ExternalUserStatusCode"`
	ResultReasonCode                         string `json:"ResultReasonCode"`
	SalesCycleCode                           string `json:"SalesCycleCode"`
	SalesCyclePhaseCode                      string `json:"SalesCyclePhaseCode"`
	ProcessStatusValidSinceDate              string `json:"ProcessStatusValidSinceDate"`
	SalesCyclePhaseStartDate                 string `json:"SalesCyclePhaseStartDate"`
	ProbabilityPercent                       string `json:"ProbabilityPercent"`
	HeaderRevenueSchedule                    bool   `json:"HeaderRevenueSchedule"`
	SalesRevenueForecastRelevanceIndicator   bool   `json:"SalesRevenueForecastRelevanceIndicator"`
	ExpectedRevenueAmount                    string `json:"ExpectedRevenueAmount"`
	ExpectedRevenueAmountCurrencyCode        string `json:"ExpectedRevenueAmountCurrencyCode"`
	TotalExpectedNetAmount                   string `json:"TotalExpectedNetAmount"`
	TotalExpectedNetAmountAmountCurrencyCode string `json:"TotalExpectedNetAmountAmountCurrencyCode"`
	WeightedExpectedNetAmount                string `json:"WeightedExpectedNetAmount"`
	WeightedExpectedNetAmountCurrencyCode    string `json:"WeightedExpectedNetAmountCurrencyCode"`
	ExpectedProcessingStartDate              string `json:"ExpectedProcessingStartDate"`
	ExpectedProcessingEndDate                string `json:"ExpectedProcessingEndDate"`
	ExpectedRevenueStartDate                 string `json:"ExpectedRevenueStartDate"`
	ExpectedRevenueEndDate                   string `json:"ExpectedRevenueEndDate"`
	SalesForecastCategoryCode                string `json:"SalesForecastCategoryCode"`
	GroupCode                                string `json:"GroupCode"`
	SalesUnitPartyID                         string `json:"SalesUnitPartyID"`
	SalesOrganisationID                      string `json:"SalesOrganisationID"`
	DistributionChannelCode                  string `json:"DistributionChannelCode"`
	DivisionCode                             string `json:"DivisionCode"`
	SalesOfficeID                            string `json:"SalesOfficeID"`
	SalesGroupID                             string `json:"SalesGroupID"`
	SalesTerritoryID                         string `json:"SalesTerritoryID"`
	MainEmployeeResponsiblePartyID           string `json:"MainEmployeeResponsiblePartyID"`
	EndBuyerPartyID                          string `json:"EndBuyerPartyID"`
	ProductRecepientPartyID                  string `json:"ProductRecepientPartyID"`
	ApproverPartyID                          string `json:"ApproverPartyID"`
	PayerPartyID                             string `json:"PayerPartyID"`
	BillToPartyID                            string `json:"BillToPartyID"`
	SellerPartyID                            string `json:"SellerPartyID"`
	PhaseProgressEvaluationStatusCode        string `json:"PhaseProgressEvaluationStatusCode"`
	ProspectBudgetAmount                     string `json:"ProspectBudgetAmount"`
	ProspectBudgetAmountCurrencyCode         string `json:"ProspectBudgetAmountCurrencyCode"`
	Score                                    int    `json:"Score"`
	MainEmployeeResponsiblePartyName         string `json:"MainEmployeeResponsiblePartyName"`
	SalesUnityPartyName                      string `json:"SalesUnityPartyName"`
	BillToPartyName                          string `json:"BillToPartyName"`
	ProductRecipientPartyName                string `json:"ProductRecipientPartyName"`
	SellerPartyName                          string `json:"SellerPartyName"`
	PayerPartyName                           string `json:"PayerPartyName"`
	EndBuyerPartyName                        string `json:"EndBuyerPartyName"`
	PrimaryContactPartyName                  string `json:"PrimaryContactPartyName"`
	ProspectPartyName                        string `json:"ProspectPartyName"`
	ApproverPartyName                        string `json:"ApproverPartyName"`
	SalesOrganisationName                    string `json:"SalesOrganisationName"`
	SalesOfficeName                          string `json:"SalesOfficeName"`
	SalesGroupName                           string `json:"SalesGroupName"`
	SalesTerritoryName                       string `json:"SalesTerritoryName"`
	ConsistencyStatusCode                    string `json:"ConsistencyStatusCode"`
	ApprovalStatusCode                       string `json:"ApprovalStatusCode"`
	CreationDate                             string `json:"CreationDate"`
	LastChangeDate                           string `json:"LastChangeDate"`
	CreationDateTime                         string `json:"CreationDateTime"`
	LastChangeDateTime                       string `json:"LastChangeDateTime"`
	CreatedBy                                string `json:"CreatedBy"`
	LastChangedBy                            string `json:"LastChangedBy"`
	CreatedByUUID                            string `json:"CreatedByUUID"`
	LastChangedByUUID                        string `json:"LastChangedByUUID"`
	EntityLastChangedOn                      string `json:"EntityLastChangedOn"`
	ETag                                     string `json:"ETag"`
	BestConnectedColleague                   string `json:"BestConnectedColleague"`
	DealScore                                string `json:"DealScore"`
	DealScoreReason                          string `json:"DealScoreReason"`
	FirstEmailReceivedOn                     string `json:"FirstEmailReceivedOn"`
	FirstEmailSentOn                         string `json:"FirstEmailSentOn"`
	LastDataHugSyncDateTime                  string `json:"LastDataHugSyncDateTime"`
	LastEmailReceivedFrom                    string `json:"LastEmailReceivedFrom"`
	LastEmailReceivedOn                      string `json:"LastEmailReceivedOn"`
	LastEmailSentBy                          string `json:"LastEmailSentBy"`
	LastEmailSentOn                          string `json:"LastEmailSentOn"`
	LastMeetingOn                            string `json:"LastMeetingOn"`
	NextMeetingOn                            string `json:"NextMeetingOn"`
	NumberOfColleagues                       int    `json:"NumberOfColleagues"`
	NumberOfMeetings                         int    `json:"NumberOfMeetings"`
	NumberOfOtherPeopleAtCompany             int    `json:"NumberOfOtherPeopleAtCompany"`
	NumberOfReceivedEmails                   int    `json:"NumberOfReceivedEmails"`
	NumberOfSentEmails                       int    `json:"NumberOfSentEmails"`
	AutoCreateContacts                       bool   `json:"AutoCreateContacts"`
	HugRank                                  int    `json:"HugRank"`
	ExternalProbabilityPercent               string `json:"ExternalProbabilityPercent"`
}
