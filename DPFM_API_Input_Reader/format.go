package dpfm_api_input_reader

import (
	"data-platform-api-invoice-document-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		InvoiceDocument:                   data.InvoiceDocument,
		CreationDate:                      data.CreationDate,
		LastChangeDate:                    data.LastChangeDate,
		BillToParty:                       data.BillToParty,
		BillFromParty:                     data.BillFromParty,
		BillToCountry:                     data.BillToCountry,
		BillFromCountry:                   data.BillFromCountry,
		InvoiceDocumentDate:               data.InvoiceDocumentDate,
		InvoiceDocumentTime:               data.InvoiceDocumentTime,
		InvoicePeriodStartDate:            data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:              data.InvoicePeriodEndDate,
		AccountingPostingDate:             data.AccountingPostingDate,
		InvoiceDocumentIsCancelled:        data.InvoiceDocumentIsCancelled,
		CancelledInvoiceDocument:          data.CancelledInvoiceDocument,
		IsExportImportDelivery:            data.IsExportImportDelivery,
		HeaderBillingIsConfirmed:          data.HeaderBillingIsConfirmed,
		HeaderBillingConfStatus:           data.HeaderBillingConfStatus,
		TotalNetAmount:                    data.TotalNetAmount,
		TotalTaxAmount:                    data.TotalTaxAmount,
		TotalGrossAmount:                  data.TotalGrossAmount,
		TransactionCurrency:               data.TransactionCurrency,
		Incoterms:                         data.Incoterms,
		PaymentTerms:                      data.PaymentTerms,
		DueCalculationBaseDate:            data.DueCalculationBaseDate,
		PaymentDueDate:                    data.PaymentDueDate,
		NetPaymentDays:                    data.NetPaymentDays,
		PaymentMethod:                     data.PaymentMethod,
		HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
		ExternalReferenceDocument:         data.ExternalReferenceDocument,
		DocumentHeaderText:                data.DocumentHeaderText,
		HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
	}
}

func (sdc *SDC) ConvertToHeaderPDF() *requests.HeaderPDF {
	dataHeader := sdc.Header
	data := sdc.Header.HeaderPDF
	return &requests.HeaderPDF{
		InvoiceDocument:          dataHeader.InvoiceDocument,
		DocType:                  data.DocType,
		DocVersionID:             data.DocVersionID,
		DocID:                    data.DocID,
		DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		FileName:                 data.FileName,
	}
}

func (sdc *SDC) ConvertToHeaderPartner() *requests.HeaderPartner {
	dataHeader := sdc.Header
	data := sdc.Header.HeaderPartner
	return &requests.HeaderPartner{
		InvoiceDocument:         dataHeader.InvoiceDocument,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
}

func (sdc *SDC) ConvertToHeaderPartnerContact() *requests.HeaderPartnerContact {
	dataHeader := sdc.Header
	dataHeaderPartner := sdc.Header.HeaderPartner
	data := sdc.Header.HeaderPartner.HeaderPartnerContact
	return &requests.HeaderPartnerContact{
		InvoiceDocument:   dataHeader.InvoiceDocument,
		PartnerFunction:   dataHeaderPartner.PartnerFunction,
		BusinessPartner:   dataHeaderPartner.BusinessPartner,
		ContactID:         data.ContactID,
		ContactPersonName: data.ContactPersonName,
		EmailAddress:      data.EmailAddress,
		PhoneNumber:       data.PhoneNumber,
		MobilePhoneNumber: data.MobilePhoneNumber,
		FaxNumber:         data.FaxNumber,
		ContactTag1:       data.ContactTag1,
		ContactTag2:       data.ContactTag2,
		ContactTag3:       data.ContactTag3,
		ContactTag4:       data.ContactTag4,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataHeader := sdc.Header
	data := sdc.Header.Item
	return &requests.Item{
		InvoiceDocument:                dataHeader.InvoiceDocument,
		InvoiceDocumentItem:            data.InvoiceDocumentItem,
		DocumentItemCategory:           data.DocumentItemCategory,
		InvoiceDocumentItemText:        data.InvoiceDocumentItemText,
		CreationDate:                   data.CreationDate,
		CreationTime:                   data.CreationTime,
		ItemBillingIsConfirmed:         data.ItemBillingIsConfirmed,
		ItemBillingConfStatus:          data.ItemBillingConfStatus,
		Buyer:                          data.Buyer,
		Seller:                         data.Seller,
		DeliverToParty:                 data.DeliverToParty,
		DeliverFromParty:               data.DeliverFromParty,
		ProductStandardID:              data.ProductStandardID,
		Batch:                          data.Batch,
		Product:                        data.Product,
		ProductGroup:                   data.ProductGroup,
		ProductionPlantPartnerFunction: data.ProductionPlantPartnerFunction,
		ProductionPlantBusinessPartner: data.ProductionPlantBusinessPartner,
		ProductionPlant:                data.ProductionPlant,
		ProductionPlantStorageLocation: data.ProductionPlantStorageLocation,
		IssuingPlantPartnerFunction:    data.IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner:    data.IssuingPlantBusinessPartner,
		IssuingPlant:                   data.IssuingPlant,
		IssuingPlantStorageLocation:    data.IssuingPlantStorageLocation,
		ReceivingPlantPartnerFunction:  data.ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner:  data.ReceivingPlantBusinessPartner,
		ReceivingPlant:                 data.ReceivingPlant,
		ReceivingPlantStorageLocation:  data.ReceivingPlantStorageLocation,
		ServicesRenderedDate:           data.ServicesRenderedDate,
		InvoiceQuantity:                data.InvoiceQuantity,
		InvoiceQuantityUnit:            data.InvoiceQuantityUnit,
		InvoiceQuantityInBaseUnit:      data.InvoiceQuantityInBaseUnit,
		BaseUnit:                       data.BaseUnit,
		ActualGoodsIssueDate:           data.ActualGoodsIssueDate,
		ActualGoodsIssueTime:           data.ActualGoodsIssueTime,
		ActualGoodsReceiptDate:         data.ActualGoodsReceiptDate,
		ActualGoodsReceiptTime:         data.ActualGoodsReceiptTime,
		ItemGrossWeight:                data.ItemGrossWeight,
		ItemNetWeight:                  data.ItemNetWeight,
		ItemWeightUnit:                 data.ItemWeightUnit,
		NetAmount:                      data.NetAmount,
		TaxAmount:                      data.TaxAmount,
		GrossAmount:                    data.GrossAmount,
		GoodsIssueOrReceiptSlipNumber:  data.GoodsIssueOrReceiptSlipNumber,
		TransactionCurrency:            data.TransactionCurrency,
		PricingDate:                    data.PricingDate,
		ProductTaxClassification:       data.ProductTaxClassification,
		Project:                        data.Project,
		OrderID:                        data.OrderID,
		OrderItem:                      data.OrderItem,
		OrderType:                      data.OrderType,
		ContractType:                   data.ContractType,
		OrderVaridityStartDate:         data.OrderVaridityStartDate,
		OrderVaridityEndDate:           data.OrderVaridityEndDate,
		InvoiceScheduleStartDate:       data.InvoiceScheduleStartDate,
		InvoiceScheduleEndDate:         data.InvoiceScheduleEndDate,
		DeliveryDocument:               data.DeliveryDocument,
		DeliveryDocumentItem:           data.DeliveryDocumentItem,
		OriginDocument:                 data.OriginDocument,
		OriginDocumentItem:             data.OriginDocumentItem,
		ReferenceDocument:              data.ReferenceDocument,
		ReferenceDocumentItem:          data.ReferenceDocumentItem,
		ReferenceDocumentType:          data.ReferenceDocumentType,
		ExternalReferenceDocument:      data.ExternalReferenceDocument,
		ExternalReferenceDocumentItem:  data.ExternalReferenceDocumentItem,
		TaxCode:                        data.TaxCode,
		TaxRate:                        data.TaxRate,
		CountryOfOrigin:                data.CountryOfOrigin,
	}
}

func (sdc *SDC) ConvertToItemPartner() *requests.ItemPartner {
	dataHeader := sdc.Header
	dataItem := sdc.Header.Item
	data := sdc.Header.Item.ItemPartner
	return &requests.ItemPartner{
		InvoiceDocument:     dataHeader.InvoiceDocument,
		InvoiceDocumentItem: dataItem.InvoiceDocumentItem,
		PartnerFunction:     data.PartnerFunction,
		BusinessPartner:     data.BusinessPartner,
	}
}

func (sdc *SDC) ConvertToItemPricingElement() *requests.ItemPricingElement {
	dataHeader := sdc.Header
	dataItem := sdc.Header.Item
	data := sdc.Header.Item.ItemPricingElement
	return &requests.ItemPricingElement{
		InvoiceDocument:            dataHeader.InvoiceDocument,
		InvoiceDocumentItem:        dataItem.InvoiceDocumentItem,
		PricingProcedureCounter:    data.PricingProcedureCounter,
		ConditionType:              data.ConditionType,
		PricingDate:                data.PricingDate,
		ConditionRateValue:         data.ConditionRateValue,
		ConditionCurrency:          data.ConditionCurrency,
		ConditionQuantity:          data.ConditionQuantity,
		ConditionQuantityUnit:      data.ConditionQuantityUnit,
		ConditionRecord:            data.ConditionRecord,
		ConditionSequentialNumber:  data.ConditionSequentialNumber,
		TaxCode:                    data.TaxCode,
		ConditionAmount:            data.ConditionAmount,
		TransactionCurrency:        data.TransactionCurrency,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}
}

func (sdc *SDC) ConvertToAddress() *requests.Address {
	dataHeader := sdc.Header
	data := sdc.Header.Address
	return &requests.Address{
		InvoiceDocument: dataHeader.InvoiceDocument,
		AddressID:       data.AddressID,
		PostalCode:      data.PostalCode,
		LocalRegion:     data.LocalRegion,
		Country:         data.Country,
		District:        data.District,
		StreetName:      data.StreetName,
		CityName:        data.CityName,
		Building:        data.Building,
		Floor:           data.Floor,
		Room:            data.Room,
	}
}
