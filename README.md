# sap-api-integrations-opportunity-reads  
sap-api-integrations-opportunity-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API オポチュニティデータを取得するマイクロサービスです。  
sap-api-integrations-opportunity-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-opportunity-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/opportunity/overview  

## 動作環境
sap-api-integrations-opportunity-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-opportunity-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-opportunity-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/opportunity/overview
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-opportunity-reads には、次の API をコールするためのリソースが含まれています。  

* OpportunityCollection（オポチュニティ - オポチュニティ）


## API への 値入力条件 の 初期値
sap-api-integrations-opportunity-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.OpportunityCollection.ID（ID）


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"OpportunityCollection" が指定されています。    
  
```
	"api_schema": "Opportunity",
	"accepter": ["OpportunityCollection"],
	"opportunity_code": "31927",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "Opportunity",
	"accepter": ["All"],
	"opportunity_code": "31927",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetOpportunity(iD string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "OpportunityCollection":
			func() {
				c.OpportunityCollection(iD)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP の オポチュニティデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "ExternalProbabilityPercent" は、/SAP_API_Output_Formatter/type.go 内 の Type OpportunityCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-opportunity-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-opportunity-reads/SAP_API_Caller.(*SAPAPICaller).OpportunityCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E04B6021EE2B3C2A1CEA2FDAFE1",
			"ProcessingTypeCode": "OPPT",
			"ID": "31927",
			"ExternalID": "",
			"UUID": "00163E04-B602-1EE2-B3C2-A1CEA2FDAFE1",
			"ProspectPartyID": "10005",
			"Name": "Trade",
			"PrimaryContactPartyID": "1000086",
			"PriorityCode": "3",
			"OriginTypeCode": "003",
			"LifeCycleStatusCode": "2",
			"ExternalUserStatusCode": "",
			"ResultReasonCode": "",
			"SalesCycleCode": "001",
			"SalesCyclePhaseCode": "004",
			"ProcessStatusValidSinceDate": "2013-06-05T09:00:00+09:00",
			"SalesCyclePhaseStartDate": "2019-11-12T09:00:00+09:00",
			"ProbabilityPercent": "60.000000",
			"HeaderRevenueSchedule": false,
			"SalesRevenueForecastRelevanceIndicator": true,
			"ExpectedRevenueAmount": "99999600.000000",
			"ExpectedRevenueAmountCurrencyCode": "USD",
			"TotalExpectedNetAmount": "1200.000000",
			"TotalExpectedNetAmountAmountCurrencyCode": "USD",
			"WeightedExpectedNetAmount": "59999760.000000",
			"WeightedExpectedNetAmountCurrencyCode": "USD",
			"ExpectedProcessingStartDate": "2014-11-18T09:00:00+09:00",
			"ExpectedProcessingEndDate": "2017-12-19T09:00:00+09:00",
			"ExpectedRevenueStartDate": "2017-12-21T09:00:00+09:00",
			"ExpectedRevenueEndDate": "2018-03-01T09:00:00+09:00",
			"SalesForecastCategoryCode": "",
			"GroupCode": "",
			"SalesUnitPartyID": "US1100",
			"SalesOrganisationID": "US1100",
			"DistributionChannelCode": "01",
			"DivisionCode": "",
			"SalesOfficeID": "",
			"SalesGroupID": "",
			"SalesTerritoryID": "3",
			"MainEmployeeResponsiblePartyID": "8000000002",
			"EndBuyerPartyID": "",
			"ProductRecepientPartyID": "10005",
			"ApproverPartyID": "",
			"PayerPartyID": "",
			"BillToPartyID": "10005",
			"SellerPartyID": "1000",
			"PhaseProgressEvaluationStatusCode": "3",
			"ProspectBudgetAmount": "0.000000",
			"ProspectBudgetAmountCurrencyCode": "USD",
			"Score": 0,
			"MainEmployeeResponsiblePartyName": "Phil Hughes",
			"SalesUnityPartyName": "Sales Unit US",
			"BillToPartyName": "FutureVision",
			"ProductRecipientPartyName": "FutureVision",
			"SellerPartyName": "Almika",
			"PayerPartyName": "",
			"EndBuyerPartyName": "",
			"PrimaryContactPartyName": "Matthias Brunner",
			"ProspectPartyName": "FutureVision",
			"ApproverPartyName": "",
			"SalesOrganisationName": "Sales Unit US",
			"SalesOfficeName": "",
			"SalesGroupName": "",
			"SalesTerritoryName": "US Central",
			"ConsistencyStatusCode": "2",
			"ApprovalStatusCode": "1",
			"CreationDate": "2013-06-05T09:00:00+09:00",
			"LastChangeDate": "2020-07-10T09:00:00+09:00",
			"CreationDateTime": "2013-06-06T04:21:08+09:00",
			"LastChangeDateTime": "2020-07-10T11:05:13+09:00",
			"CreatedBy": "Eddie Smoke",
			"LastChangedBy": "",
			"CreatedByUUID": "00163E03-A070-1EE2-88BA-39BD20F290B5",
			"LastChangedByUUID": "00163E03-A06A-1EE2-87C4-A680EDD73F8D",
			"EntityLastChangedOn": "2020-07-10T11:05:13+09:00",
			"ETag": "2020-07-10T11:05:13+09:00",
			"BestConnectedColleague": "",
			"DealScore": "0.000000",
			"DealScoreReason": "",
			"FirstEmailReceivedOn": "",
			"FirstEmailSentOn": "",
			"LastDataHugSyncDateTime": "",
			"LastEmailReceivedFrom": "",
			"LastEmailReceivedOn": "",
			"LastEmailSentBy": "",
			"LastEmailSentOn": "",
			"LastMeetingOn": "",
			"NextMeetingOn": "",
			"NumberOfColleagues": 0,
			"NumberOfMeetings": 0,
			"NumberOfOtherPeopleAtCompany": 0,
			"NumberOfReceivedEmails": 0,
			"NumberOfSentEmails": 0,
			"AutoCreateContacts": false,
			"HugRank": 0,
			"ExternalProbabilityPercent": "0.000000"
		}
	],
	"time": "2022-08-10T18:38:32+09:00"
}

```