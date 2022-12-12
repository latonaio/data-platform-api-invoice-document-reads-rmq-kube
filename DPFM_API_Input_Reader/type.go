package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Header            Header   `json:"InvoiceDocument"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	InvoiceDocument                   int           `json:"InvoiceDocument"`
	CreationDate                      *string       `json:"CreationDate"`
	LastChangeDate                    *string       `json:"LastChangeDate"`
	BillToParty                       *int          `json:"BillToParty"`
	BillFromParty                     *int          `json:"BillFromParty"`
	BillToCountry                     *string       `json:"BillToCountry"`
	BillFromCountry                   *string       `json:"BillFromCountry"`
	InvoiceDocumentDate               *string       `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime               *string       `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate            *string       `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              *string       `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string       `json:"AccountingPostingDate"`
	InvoiceDocumentIsCancelled        *bool         `json:"InvoiceDocumentIsCancelled"`
	CancelledInvoiceDocument          *int          `json:"CancelledInvoiceDocument"`
	IsExportImportDelivery            *bool         `json:"IsExportImportDelivery"`
	HeaderBillingIsConfirmed          *bool         `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus           *string       `json:"HeaderBillingConfStatus"`
	TotalNetAmount                    *float32      `json:"TotalNetAmount"`
	TotalTaxAmount                    *float32      `json:"TotalTaxAmount"`
	TotalGrossAmount                  *float32      `json:"TotalGrossAmount"`
	TransactionCurrency               *string       `json:"TransactionCurrency"`
	Incoterms                         *string       `json:"Incoterms"`
	PaymentTerms                      *string       `json:"PaymentTerms"`
	DueCalculationBaseDate            *string       `json:"DueCalculationBaseDate"`
	PaymentDueDate                    *string       `json:"PaymentDueDate"`
	NetPaymentDays                    *int          `json:"NetPaymentDays"`
	PaymentMethod                     *string       `json:"PaymentMethod"`
	HeaderPaymentBlockStatus          *bool         `json:"HeaderPaymentBlockStatus"`
	ExternalReferenceDocument         *string       `json:"ExternalReferenceDocument"`
	DocumentHeaderText                *string       `json:"DocumentHeaderText"`
	HeaderPaymentRequisitionIsCreated *bool         `json:"HeaderPaymentRequisitionIsCreated"`
	HeaderPDF                         HeaderPDF     `json:"HeaderPDF"`
	HeaderPartner                     HeaderPartner `json:"HeaderPartner"`
	Item                              Item          `json:"Item"`
	Address                           Address       `json:"Address"`
}

type HeaderPDF struct {
	InvoiceDocument          int     `json:"InvoiceDocument"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FileName                 *string `json:"FileName"`
}

type HeaderPartner struct {
	InvoiceDocument         int                  `json:"InvoiceDocument"`
	PartnerFunction         string               `json:"PartnerFunction"`
	BusinessPartner         int                  `json:"BusinessPartner"`
	BusinessPartnerFullName *string              `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string              `json:"BusinessPartnerName"`
	Organization            *string              `json:"Organization"`
	Country                 *string              `json:"Country"`
	Language                *string              `json:"Language"`
	Currency                *string              `json:"Currency"`
	ExternalDocumentID      *string              `json:"ExternalDocumentID"`
	AddressID               *int                 `json:"AddressID"`
	HeaderPartnerContact    HeaderPartnerContact `json:"HeaderPartnerContact"`
}

type HeaderPartnerContact struct {
	InvoiceDocument   int     `json:"InvoiceDocument"`
	PartnerFunction   string  `json:"PartnerFunction"`
	BusinessPartner   int     `json:"BusinessPartner"`
	ContactID         int     `json:"ContactID"`
	ContactPersonName *string `json:"ContactPersonName"`
	EmailAddress      *string `json:"EmailAddress"`
	PhoneNumber       *string `json:"PhoneNumber"`
	MobilePhoneNumber *string `json:"MobilePhoneNumber"`
	FaxNumber         *string `json:"FaxNumber"`
	ContactTag1       *string `json:"ContactTag1"`
	ContactTag2       *string `json:"ContactTag2"`
	ContactTag3       *string `json:"ContactTag3"`
	ContactTag4       *string `json:"ContactTag4"`
}

type Item struct {
	InvoiceDocument                int                `json:"InvoiceDocument"`
	InvoiceDocumentItem            int                `json:"InvoiceDocumentItem"`
	DocumentItemCategory           *string            `json:"DocumentItemCategory"`
	InvoiceDocumentItemText        *string            `json:"InvoiceDocumentItemText"`
	CreationDate                   *string            `json:"CreationDate"`
	CreationTime                   *string            `json:"CreationTime"`
	ItemBillingIsConfirmed         *bool              `json:"ItemBillingIsConfirmed"`
	ItemBillingConfStatus          *string            `json:"ItemBillingConfStatus"`
	Buyer                          *int               `json:"Buyer"`
	Seller                         *int               `json:"Seller"`
	DeliverToParty                 *int               `json:"DeliverToParty"`
	DeliverFromParty               *int               `json:"DeliverFromParty"`
	ProductStandardID              *string            `json:"ProductStandardID"`
	Batch                          *string            `json:"Batch"`
	Product                        *string            `json:"Product"`
	ProductGroup                   *string            `json:"ProductGroup"`
	ProductionPlantPartnerFunction *string            `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner *string            `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                *string            `json:"ProductionPlant"`
	ProductionPlantStorageLocation *string            `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction    *string            `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner    *string            `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                   *string            `json:"IssuingPlant"`
	IssuingPlantStorageLocation    *string            `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction  *string            `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner  *string            `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                 *string            `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation  *string            `json:"ReceivingPlantStorageLocation"`
	ServicesRenderedDate           *string            `json:"ServicesRenderedDate"`
	InvoiceQuantity                *float32           `json:"InvoiceQuantity"`
	InvoiceQuantityUnit            *string            `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit      *float32           `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                       *string            `json:"BaseUnit"`
	ActualGoodsIssueDate           *string            `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime           *string            `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate         *string            `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime         *string            `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight                *float32           `json:"ItemGrossWeight"`
	ItemNetWeight                  *float32           `json:"ItemNetWeight"`
	ItemWeightUnit                 *string            `json:"ItemWeightUnit"`
	NetAmount                      *float32           `json:"NetAmount"`
	TaxAmount                      *float32           `json:"TaxAmount"`
	GrossAmount                    *float32           `json:"GrossAmount"`
	GoodsIssueOrReceiptSlipNumber  *string            `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency            *string            `json:"TransactionCurrency"`
	PricingDate                    *string            `json:"PricingDate"`
	ProductTaxClassification       *string            `json:"ProductTaxClassification"`
	Project                        *string            `json:"Project"`
	OrderID                        *int               `json:"OrderID"`
	OrderItem                      *int               `json:"OrderItem"`
	OrderType                      *string            `json:"OrderType"`
	ContractType                   *string            `json:"ContractType"`
	OrderVaridityStartDate         *string            `json:"OrderVaridityStartDate"`
	OrderVaridityEndDate           *string            `json:"OrderVaridityEndDate"`
	InvoiceScheduleStartDate       *string            `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate         *string            `json:"InvoiceScheduleEndDate"`
	DeliveryDocument               *int               `json:"DeliveryDocument"`
	DeliveryDocumentItem           *int               `json:"DeliveryDocumentItem"`
	OriginDocument                 *int               `json:"OriginDocument"`
	OriginDocumentItem             *int               `json:"OriginDocumentItem"`
	ReferenceDocument              *int               `json:"ReferenceDocument"`
	ReferenceDocumentItem          *int               `json:"ReferenceDocumentItem"`
	ReferenceDocumentType          *string            `json:"ReferenceDocumentType"`
	ExternalReferenceDocument      *string            `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem  *string            `json:"ExternalReferenceDocumentItem"`
	TaxCode                        *string            `json:"TaxCode"`
	TaxRate                        *float32           `json:"TaxRate"`
	CountryOfOrigin                *string            `json:"CountryOfOrigin"`
	ItemPartner                    ItemPartner        `json:"ItemPartner"`
	ItemPricingElement             ItemPricingElement `json:"ItemPricingElement"`
}

type ItemPartner struct {
	InvoiceDocument     int    `json:"InvoiceDocument"`
	InvoiceDocumentItem int    `json:"InvoiceDocumentItem"`
	PartnerFunction     string `json:"PartnerFunction"`
	BusinessPartner     int    `json:"BusinessPartner"`
}

type ItemPricingElement struct {
	InvoiceDocument            int      `json:"InvoiceDocument"`
	InvoiceDocumentItem        int      `json:"InvoiceDocumentItem"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type Address struct {
	InvoiceDocument int     `json:"InvoiceDocument"`
	AddressID       int     `json:"AddressID"`
	PostalCode      *string `json:"PostalCode"`
	LocalRegion     *string `json:"LocalRegion"`
	Country         *string `json:"Country"`
	District        *string `json:"District"`
	StreetName      *string `json:"StreetName"`
	CityName        *string `json:"CityName"`
	Building        *string `json:"Building"`
	Floor           *int    `json:"Floor"`
	Room            *int    `json:"Room"`
}
