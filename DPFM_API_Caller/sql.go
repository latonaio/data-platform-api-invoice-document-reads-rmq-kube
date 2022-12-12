package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-invoice-document-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var headerPartner *dpfm_api_output_formatter.HeaderPartner
	var headerPartnerContact *dpfm_api_output_formatter.HeaderPartnerContact
	var address *dpfm_api_output_formatter.Address
	var item *dpfm_api_output_formatter.Item
	var itemPartner *dpfm_api_output_formatter.ItemPartner
	var itemPricingElement *dpfm_api_output_formatter.ItemPricingElement
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeaderPartner":
			func() {
				headerPartner = c.HeaderPartner(mtx, input, output, errs, log)
			}()
		case "HeaderPartnerContact":
			func() {
				headerPartnerContact = c.HeaderPartnerContact(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "ItemPartner":
			func() {
				itemPartner = c.ItemPartner(mtx, input, output, errs, log)
			}()
		case "ItemPricingElement":
			func() {
				itemPricingElement = c.ItemPricingElement(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:               header,
		HeaderPartner:        headerPartner,
		HeaderPartnerContact: headerPartnerContact,
		Address:              address,
		Item:                 item,
		ItemPartner:          itemPartner,
		ItemPricingElement:   itemPricingElement,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	invoiceDocument := input.Header.InvoiceDocument

	rows, err := c.db.Query(
		`SELECT InvoiceDocument, CreationDate, LastChangeDate, BillToParty, BillFromParty, BillToCountry, 
		BillFromCountry, InvoiceDocumentDate, InvoiceDocumentTime, InvoicePeriodStartDate, InvoicePeriodEndDate, 
		AccountingPostingDate, InvoiceDocumentIsCancelled, CancelledInvoiceDocument, IsExportImportDelivery, HeaderBillingIsConfirmed, 
		HeaderBillingConfStatus, TotalNetAmount, TotalTaxAmount, TotalGrossAmount, TransactionCurrency, Incoterms, PaymentTerms, DueCalculationBaseDate, 
		PaymentDueDate, NetPaymentDays, PaymentMethod, HeaderPaymentBlockStatus, ExternalReferenceDocument, DocumentHeaderText, HeaderPaymentRequisitionIsCreated
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_data
		WHERE InvoiceDocument = ?;`, invoiceDocument,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartner {
	invoiceDocument := input.Header.InvoiceDocument
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT InvoiceDocument, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, 
		Organization, Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_partner_data
		WHERE (InvoiceDocument, PartnerFunction, BusinessPartner) = (?, ?, ?);`, invoiceDocument, partnerFunction, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartnerContact(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartnerContact {
	invoiceDocument := input.Header.InvoiceDocument
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner
	contactID := input.Header.HeaderPartner.HeaderPartnerContact.ContactID

	rows, err := c.db.Query(
		`SELECT InvoiceDocument, PartnerFunction, BusinessPartner, ContactID, ContactPersonName, EmailAddress, PhoneNumber, 
		MobilePhoneNumber, FaxNumber, ContactTag1, ContactTag2, ContactTag3, ContactTag4
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_partner_contact_data
		WHERE (InvoiceDocument, PartnerFunction, BusinessPartner, ContactID) = (?, ?, ?, ?);`, invoiceDocument, partnerFunction, businessPartner, contactID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartnerContact(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Address {
	invoiceDocument := input.Header.InvoiceDocument
	addressID := input.Header.Address.AddressID

	rows, err := c.db.Query(
		`SELECT InvoiceDocument, AddressID, PostalCode, LocalRegion, Country, District, StreetName, CityName, 
		Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_address_data
		WHERE (InvoiceDocument, AddressID) = (?, ?);`, invoiceDocument, addressID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAddress(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Item {
	invoiceDocument := input.Header.InvoiceDocument
	invoiceDocumentItem := input.Header.Item.InvoiceDocumentItem

	rows, err := c.db.Query(
		`SELECT InvoiceDocument, InvoiceDocumentItem, DocumentItemCategory, InvoiceDocumentItemText, CreationDate, 
		CreationTime, ItemBillingIsConfirmed, ItemBillingConfStatus, Buyer, Seller, DeliverToParty, DeliverFromParty, 
		ProductStandardID, Batch, Product, ProductGroup, ProductionPlantPartnerFunction, ProductionPlantBusinessPartner, 
		ProductionPlant, ProductionPlantStorageLocation, IssuingPlantPartnerFunction, IssuingPlantBusinessPartner, 
		IssuingPlant, IssuingPlantStorageLocation, ReceivingPlantPartnerFunction, ReceivingPlantBusinessPartner, 
		ReceivingPlant, ReceivingPlantStorageLocation, ServicesRenderedDate, InvoiceQuantity, InvoiceQuantityUnit, 
		InvoiceQuantityInBaseUnit, BaseUnit, ActualGoodsIssueDate, ActualGoodsIssueTime, ActualGoodsReceiptDate, 
		ActualGoodsReceiptTime, ItemGrossWeight, ItemNetWeight, ItemWeightUnit, NetAmount, TaxAmount, GrossAmount, 
		GoodsIssueOrReceiptSlipNumber, TransactionCurrency, PricingDate, ProductTaxClassification, Project, OrderID, 
		OrderItem, OrderType, ContractType, OrderVaridityStartDate, OrderVaridityEndDate, InvoiceScheduleStartDate, 
		InvoiceScheduleEndDate, DeliveryDocument, DeliveryDocumentItem, OriginDocument, OriginDocumentItem, ReferenceDocument, 
		ReferenceDocumentItem, ReferenceDocumentType, ExternalReferenceDocument, ExternalReferenceDocumentItem, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_item_data
		WHERE (InvoiceDocument, InvoiceDocumentItem) = (?, ?);`, invoiceDocument, invoiceDocumentItem,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItem(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPartner {
	invoiceDocument := input.Header.InvoiceDocument
	invoiceDocumentItem := input.Header.Item.InvoiceDocumentItem
	partnerFunction := input.Header.Item.ItemPartner.PartnerFunction
	businessPartner := input.Header.Item.ItemPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT InvoiceDocument, InvoiceDocumentItem, PartnerFunction, BusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_item_partner_data
		WHERE (InvoiceDocument, InvoiceDocumentItem, PartnerFunction, BusinessPartner) = (?, ?, ?, ?);`, invoiceDocument, invoiceDocumentItem, partnerFunction, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPricingElement(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPricingElement {
	invoiceDocument := input.Header.InvoiceDocument
	invoiceDocumentItem := input.Header.Item.InvoiceDocumentItem
	pricingProcedureCounter := input.Header.Item.ItemPricingElement.PricingProcedureCounter

	rows, err := c.db.Query(
		`SELECT InvoiceDocument, InvoiceDocumentItem, PricingProcedureCounter, ConditionType, PricingDate, 
		ConditionRateValue, ConditionCurrency, ConditionQuantity, ConditionQuantityUnit, ConditionRecord, 
		ConditionSequentialNumber, TaxCode, ConditionAmount, TransactionCurrency, ConditionIsManuallyChanged
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_item_pricing_element_data
		WHERE (InvoiceDocument, InvoiceDocumentItem, PricingProcedureCounter) = (?, ?, ?);`, invoiceDocument, invoiceDocumentItem, pricingProcedureCounter,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
