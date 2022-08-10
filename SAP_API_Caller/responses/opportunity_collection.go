package responses

type OpportunityCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			OpportunityAccountTeamPartyDeprecated    struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityAccountTeamPartyDeprecated"`
			OpportunityAttachmentFolder struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityAttachmentFolder"`
			OpportunityBusinessTransactionDocumentReference struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityBusinessTransactionDocumentReference"`
			OpportunityCompetitorParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityCompetitorParty"`
			OpportunityEndBuyerPartyContactPartyDeprecated struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityEndBuyerPartyContactPartyDeprecated"`
			OpportunityExternalParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityExternalParty"`
			OpportunityHistoricalVersion struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityHistoricalVersion"`
			OpportunityInstalledObject struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityInstalledObject"`
			OpportunityItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityItem"`
			OpportunityOtherPartyDeprecated struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityOtherPartyDeprecated"`
			OpportunityParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityParty"`
			OpportunityPayerPartyContactPartyDeprecated struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityPayerPartyContactPartyDeprecated"`
			OpportunityProductRecipientPartyContactPartyDeprecated struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityProductRecipientPartyContactPartyDeprecated"`
			OpportunityProspectPartyContactParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityProspectPartyContactParty"`
			OpportunityRevenuePlanHeader struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityRevenuePlanHeader"`
			OpportunitySalesEmployeePartyDeprecated struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunitySalesEmployeePartyDeprecated"`
			OpportunitySalesPartnerPartyDeprecated struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunitySalesPartnerPartyDeprecated"`
			OpportunitySalesTeamParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunitySalesTeamParty"`
			OpportunityTextCollection struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"OpportunityTextCollection"`
		} `json:"results"`
	} `json:"d"`
}
