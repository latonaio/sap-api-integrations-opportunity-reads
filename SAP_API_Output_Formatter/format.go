package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-opportunity-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToOpportunityCollection(raw []byte, l *logger.Logger) ([]OpportunityCollection, error) {
	pm := &responses.OpportunityCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to OpportunityCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	opportunityCollection := make([]OpportunityCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		opportunityCollection = append(opportunityCollection, OpportunityCollection{
			ObjectID:                                 data.ObjectID,
			ProcessingTypeCode:                       data.ProcessingTypeCode,
			ID:                                       data.ID,
			ExternalID:                               data.ExternalID,
			UUID:                                     data.UUID,
			ProspectPartyID:                          data.ProspectPartyID,
			Name:                                     data.Name,
			PrimaryContactPartyID:                    data.PrimaryContactPartyID,
			PriorityCode:                             data.PriorityCode,
			OriginTypeCode:                           data.OriginTypeCode,
			LifeCycleStatusCode:                      data.LifeCycleStatusCode,
			ExternalUserStatusCode:                   data.ExternalUserStatusCode,
			ResultReasonCode:                         data.ResultReasonCode,
			SalesCycleCode:                           data.SalesCycleCode,
			SalesCyclePhaseCode:                      data.SalesCyclePhaseCode,
			ProcessStatusValidSinceDate:              data.ProcessStatusValidSinceDate,
			SalesCyclePhaseStartDate:                 data.SalesCyclePhaseStartDate,
			ProbabilityPercent:                       data.ProbabilityPercent,
			HeaderRevenueSchedule:                    data.HeaderRevenueSchedule,
			SalesRevenueForecastRelevanceIndicator:   data.SalesRevenueForecastRelevanceIndicator,
			ExpectedRevenueAmount:                    data.ExpectedRevenueAmount,
			ExpectedRevenueAmountCurrencyCode:        data.ExpectedRevenueAmountCurrencyCode,
			TotalExpectedNetAmount:                   data.TotalExpectedNetAmount,
			TotalExpectedNetAmountAmountCurrencyCode: data.TotalExpectedNetAmountAmountCurrencyCode,
			WeightedExpectedNetAmount:                data.WeightedExpectedNetAmount,
			WeightedExpectedNetAmountCurrencyCode:    data.WeightedExpectedNetAmountCurrencyCode,
			ExpectedProcessingStartDate:              data.ExpectedProcessingStartDate,
			ExpectedProcessingEndDate:                data.ExpectedProcessingEndDate,
			ExpectedRevenueStartDate:                 data.ExpectedRevenueStartDate,
			ExpectedRevenueEndDate:                   data.ExpectedRevenueEndDate,
			SalesForecastCategoryCode:                data.SalesForecastCategoryCode,
			GroupCode:                                data.GroupCode,
			SalesUnitPartyID:                         data.SalesUnitPartyID,
			SalesOrganisationID:                      data.SalesOrganisationID,
			DistributionChannelCode:                  data.DistributionChannelCode,
			DivisionCode:                             data.DivisionCode,
			SalesOfficeID:                            data.SalesOfficeID,
			SalesGroupID:                             data.SalesGroupID,
			SalesTerritoryID:                         data.SalesTerritoryID,
			MainEmployeeResponsiblePartyID:           data.MainEmployeeResponsiblePartyID,
			EndBuyerPartyID:                          data.EndBuyerPartyID,
			ProductRecepientPartyID:                  data.ProductRecepientPartyID,
			ApproverPartyID:                          data.ApproverPartyID,
			PayerPartyID:                             data.PayerPartyID,
			BillToPartyID:                            data.BillToPartyID,
			SellerPartyID:                            data.SellerPartyID,
			PhaseProgressEvaluationStatusCode:        data.PhaseProgressEvaluationStatusCode,
			ProspectBudgetAmount:                     data.ProspectBudgetAmount,
			ProspectBudgetAmountCurrencyCode:         data.ProspectBudgetAmountCurrencyCode,
			Score:                                    data.Score,
			MainEmployeeResponsiblePartyName:         data.MainEmployeeResponsiblePartyName,
			SalesUnityPartyName:                      data.SalesUnityPartyName,
			BillToPartyName:                          data.BillToPartyName,
			ProductRecipientPartyName:                data.ProductRecipientPartyName,
			SellerPartyName:                          data.SellerPartyName,
			PayerPartyName:                           data.PayerPartyName,
			EndBuyerPartyName:                        data.EndBuyerPartyName,
			PrimaryContactPartyName:                  data.PrimaryContactPartyName,
			ProspectPartyName:                        data.ProspectPartyName,
			ApproverPartyName:                        data.ApproverPartyName,
			SalesOrganisationName:                    data.SalesOrganisationName,
			SalesOfficeName:                          data.SalesOfficeName,
			SalesGroupName:                           data.SalesGroupName,
			SalesTerritoryName:                       data.SalesTerritoryName,
			ConsistencyStatusCode:                    data.ConsistencyStatusCode,
			ApprovalStatusCode:                       data.ApprovalStatusCode,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			CreationDateTime:                         data.CreationDateTime,
			LastChangeDateTime:                       data.LastChangeDateTime,
			CreatedBy:                                data.CreatedBy,
			LastChangedBy:                            data.LastChangedBy,
			CreatedByUUID:                            data.CreatedByUUID,
			LastChangedByUUID:                        data.LastChangedByUUID,
			EntityLastChangedOn:                      data.EntityLastChangedOn,
			ETag:                                     data.ETag,
			BestConnectedColleague:                   data.BestConnectedColleague,
			DealScore:                                data.DealScore,
			DealScoreReason:                          data.DealScoreReason,
			FirstEmailReceivedOn:                     data.FirstEmailReceivedOn,
			FirstEmailSentOn:                         data.FirstEmailSentOn,
			LastDataHugSyncDateTime:                  data.LastDataHugSyncDateTime,
			LastEmailReceivedFrom:                    data.LastEmailReceivedFrom,
			LastEmailReceivedOn:                      data.LastEmailReceivedOn,
			LastEmailSentBy:                          data.LastEmailSentBy,
			LastEmailSentOn:                          data.LastEmailSentOn,
			LastMeetingOn:                            data.LastMeetingOn,
			NextMeetingOn:                            data.NextMeetingOn,
			NumberOfColleagues:                       data.NumberOfColleagues,
			NumberOfMeetings:                         data.NumberOfMeetings,
			NumberOfOtherPeopleAtCompany:             data.NumberOfOtherPeopleAtCompany,
			NumberOfReceivedEmails:                   data.NumberOfReceivedEmails,
			NumberOfSentEmails:                       data.NumberOfSentEmails,
			AutoCreateContacts:                       data.AutoCreateContacts,
			HugRank:                                  data.HugRank,
			ExternalProbabilityPercent:               data.ExternalProbabilityPercent,
		})
	}

	return opportunityCollection, nil
}
