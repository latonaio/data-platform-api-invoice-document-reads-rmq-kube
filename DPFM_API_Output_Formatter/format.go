package dpfm_api_output_formatter

import (
	"data-platform-api-invoice-document-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-invoice-document-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToHeader(sdc *api_input_reader.SDC, rows *sql.Rows) (*Header, error) {
	pm := &requests.Header{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.InvoiceDocumentDate,
			&pm.InvoiceDocumentTime,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.AccountingPostingDate,
			&pm.InvoiceDocumentIsCancelled,
			&pm.CancelledInvoiceDocument,
			&pm.IsExportImportDelivery,
			&pm.HeaderBillingIsConfirmed,
			&pm.HeaderBillingConfStatus,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.HeaderPaymentBlockStatus,
			&pm.ExternalReferenceDocument,
			&pm.DocumentHeaderText,
			&pm.HeaderPaymentRequisitionIsCreated,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	header := &Header{
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
		NetPaymentDays:                    data.NetPaymentDays,
		PaymentMethod:                     data.PaymentMethod,
		HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
		ExternalReferenceDocument:         data.ExternalReferenceDocument,
		DocumentHeaderText:                data.DocumentHeaderText,
		HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
	}
	return header, nil
}

func ConvertToHeaderPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartner, error) {
	pm := &requests.HeaderPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartner := &HeaderPartner{
		InvoiceDocument:         data.InvoiceDocument,
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
	return headerPartner, nil
}

func ConvertToHeaderPartnerContact(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartnerContact, error) {
	pm := &requests.HeaderPartnerContact{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.ContactID,
			&pm.ContactPersonName,
			&pm.EmailAddress,
			&pm.PhoneNumber,
			&pm.MobilePhoneNumber,
			&pm.FaxNumber,
			&pm.ContactTag1,
			&pm.ContactTag2,
			&pm.ContactTag3,
			&pm.ContactTag4,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartnerContact := &HeaderPartnerContact{
		InvoiceDocument:   data.InvoiceDocument,
		PartnerFunction:   data.PartnerFunction,
		BusinessPartner:   data.BusinessPartner,
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
	return headerPartnerContact, nil
}

func ConvertToAddress(sdc *api_input_reader.SDC, rows *sql.Rows) (*Address, error) {
	pm := &requests.Address{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	address := &Address{
		InvoiceDocument: data.InvoiceDocument,
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
	return address, nil
}

func ConvertToItem(sdc *api_input_reader.SDC, rows *sql.Rows) (*Item, error) {
	pm := &requests.Item{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.InvoiceDocumentItem,
			&pm.DocumentItemCategory,
			&pm.InvoiceDocumentItemText,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.ItemBillingIsConfirmed,
			&pm.ItemBillingConfStatus,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.ProductStandardID,
			&pm.Batch,
			&pm.Product,
			&pm.ProductGroup,
			&pm.ProductionPlantPartnerFunction,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.IssuingPlantPartnerFunction,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ReceivingPlantBusinessPartner,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantStorageLocation,
			&pm.ServicesRenderedDate,
			&pm.InvoiceQuantity,
			&pm.InvoiceQuantityUnit,
			&pm.InvoiceQuantityInBaseUnit,
			&pm.BaseUnit,
			&pm.ActualGoodsIssueDate,
			&pm.ActualGoodsIssueTime,
			&pm.ActualGoodsReceiptDate,
			&pm.ActualGoodsReceiptTime,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.GoodsIssueOrReceiptSlipNumber,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.ProductTaxClassification,
			&pm.Project,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderType,
			&pm.ContractType,
			&pm.OrderVaridityStartDate,
			&pm.OrderVaridityEndDate,
			&pm.InvoiceScheduleStartDate,
			&pm.InvoiceScheduleEndDate,
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.OriginDocument,
			&pm.OriginDocumentItem,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.ReferenceDocumentType,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	item := &Item{
		InvoiceDocument:                data.InvoiceDocument,
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
	return item, nil
}

func ConvertToItemPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*ItemPartner, error) {
	pm := &requests.ItemPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.InvoiceDocumentItem,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	itemPartner := &ItemPartner{
		InvoiceDocument:     data.InvoiceDocument,
		InvoiceDocumentItem: data.InvoiceDocumentItem,
		PartnerFunction:     data.PartnerFunction,
		BusinessPartner:     data.BusinessPartner,
	}
	return itemPartner, nil
}

func ConvertToItemPricingElement(sdc *api_input_reader.SDC, rows *sql.Rows) (*ItemPricingElement, error) {
	pm := &requests.ItemPricingElement{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.InvoiceDocumentItem,
			&pm.PricingProcedureCounter,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.ConditionQuantityUnit,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	itemPricingElement := &ItemPricingElement{
		InvoiceDocument:            data.InvoiceDocument,
		InvoiceDocumentItem:        data.InvoiceDocumentItem,
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
	return itemPricingElement, nil
}
