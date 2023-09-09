package dpfm_api_output_formatter

import (
	"data-platform-api-invoice-document-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.InvoiceDocumentDate,
			&pm.InvoiceDocumentTime,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.AccountingPostingDate,
			&pm.IsExportImport,
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
			&pm.ExternalReferenceDocument,
			&pm.DocumentHeaderText,
			&pm.HeaderIsCleared,
			&pm.HeaderPaymentBlockStatus,
			&pm.HeaderPaymentRequisitionIsCreated,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			InvoiceDocument:                   data.InvoiceDocument,
			SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID:  data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID:  data.SupplyChainRelationshipPaymentID,
			Buyer:                       	   data.Buyer,
			Seller:		                       data.Seller,
			BillToParty:                       data.BillToParty,
			BillFromParty:                     data.BillFromParty,
			BillToCountry:                     data.BillToCountry,
			BillFromCountry:                   data.BillFromCountry,
			Payer:                             data.Payer,
			Payee:                             data.Payee,
			InvoiceDocumentDate:               data.InvoiceDocumentDate,
			InvoiceDocumentTime:               data.InvoiceDocumentTime,
			InvoicePeriodStartDate:            data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:              data.InvoicePeriodEndDate,
			AccountingPostingDate:             data.AccountingPostingDate,
			IsExportImport:                    data.IsExportImport,
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
			ExternalReferenceDocument:         data.ExternalReferenceDocument,
			DocumentHeaderText:                data.DocumentHeaderText,
			HeaderIsCleared:                   data.HeaderIsCleared,
			HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
			HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
			CreationDate:                      data.CreationDate,
			CreationTime:                      data.CreationTime,
			LastChangeDate:                    data.LastChangeDate,
			LastChangeTime:                    data.LastChangeTime,
			IsCancelled:                       data.IsCancelled,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToPartner(rows *sql.Rows) (*[]Partner, error) {
	defer rows.Close()
	partner := make([]Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Partner{}

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
			return &partner, err
		}

		data := pm
		partner = append(partner, Partner{
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
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &partner, nil
	}

	return &partner, nil
}

func ConvertToAddress(rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	address := make([]Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

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
			return &address, err
		}

		data := pm
		address = append(address, Address{InvoiceDocument: data.InvoiceDocument,
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &address, nil
	}

	return &address, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.InvoiceDocumentItem,
			&pm.InvoiceDocumentItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.InvoiceDocumentItemText,
			&pm.InvoiceDocumentItemTextByBuyer,
			&pm.InvoiceDocumentItemTextBySeller,
			&pm.Product,
			&pm.ProductGroup,
			&pm.ProductStandardID,
			&pm.ItemBillingIsConfirmed,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.ServicesRenderedDate,
			&pm.InvoiceQuantity,
			&pm.InvoiceQuantityUnit,
			&pm.InvoiceQuantityInBaseUnit,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
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
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.Project,
			&pm.WBSElement,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderType,
			&pm.ContractType,
			&pm.OrderVaridityStartDate,
			&pm.OrderVaridityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.OriginDocument,
			&pm.OriginDocumentItem,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.ExternalReferenceDocument,
			&pm.ExternalReferenceDocumentItem,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.Equipment,
			&pm.ItemPaymentRequisitionIsCreated,
			&pm.ItemIsCleared,
			&pm.ItemPaymentBlockStatus,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, nil
		}

		data := pm
		item = append(item, Item{
			InvoiceDocument:                         data.InvoiceDocument,
			InvoiceDocumentItem:                     data.InvoiceDocumentItem,
			InvoiceDocumentItemCategory:             data.InvoiceDocumentItemCategory,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:       data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:  data.SupplyChainRelationshipDeliveryPlantID,
			InvoiceDocumentItemText:                 data.InvoiceDocumentItemText,
			InvoiceDocumentItemTextByBuyer:          data.InvoiceDocumentItemTextByBuyer,
			InvoiceDocumentItemTextBySeller:         data.InvoiceDocumentItemTextBySeller,
			Product:                                 data.Product,
			ProductGroup:                            data.ProductGroup,
			ProductStandardID:                       data.ProductStandardID,
			ItemBillingIsConfirmed:                  data.ItemBillingIsConfirmed,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			DeliverToParty:                          data.DeliverToParty,
			DeliverFromParty:                        data.DeliverFromParty,
			DeliverToPlant:                          data.DeliverToPlant,
			DeliverToPlantStorageLocation:           data.DeliverToPlantStorageLocation,
			DeliverFromPlant:                        data.DeliverFromPlant,
			DeliverFromPlantStorageLocation:         data.DeliverFromPlantStorageLocation,
			ProductionPlantBusinessPartner:          data.ProductionPlantBusinessPartner,
			ProductionPlant:                         data.ProductionPlant,
			ProductionPlantStorageLocation:          data.ProductionPlantStorageLocation,
			ServicesRenderedDate:                    data.ServicesRenderedDate,
			InvoiceQuantity:                         data.InvoiceQuantity,
			InvoiceQuantityUnit:                     data.InvoiceQuantityUnit,
			InvoiceQuantityInBaseUnit:               data.InvoiceQuantityInBaseUnit,
			BaseUnit:                                data.BaseUnit,
			DeliveryUnit:                            data.DeliveryUnit,
			ActualGoodsIssueDate:                    data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:                    data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:                  data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:                  data.ActualGoodsReceiptTime,
			ItemGrossWeight:                         data.ItemGrossWeight,
			ItemNetWeight:                           data.ItemNetWeight,
			ItemWeightUnit:                          data.ItemWeightUnit,
			NetAmount:                               data.NetAmount,
			TaxAmount:                               data.TaxAmount,
			GrossAmount:                             data.GrossAmount,
			GoodsIssueOrReceiptSlipNumber:           data.GoodsIssueOrReceiptSlipNumber,
			TransactionCurrency:                     data.TransactionCurrency,
			PricingDate:                             data.PricingDate,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                data.DefinedTaxClassification,
			Project:                                 data.Project,
			WBSElement:                              data.WBSElement,
			OrderID:                                 data.OrderID,
			OrderItem:                               data.OrderItem,
			OrderType:                               data.OrderType,
			ContractType:                            data.ContractType,
			OrderVaridityStartDate:                  data.OrderVaridityStartDate,
			OrderVaridityEndDate:                    data.OrderVaridityEndDate,
			InvoicePeriodStartDate:                  data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:                    data.InvoicePeriodEndDate,
			DeliveryDocument:                        data.DeliveryDocument,
			DeliveryDocumentItem:                    data.DeliveryDocumentItem,
			OriginDocument:                          data.OriginDocument,
			OriginDocumentItem:                      data.OriginDocumentItem,
			ReferenceDocument:                       data.ReferenceDocument,
			ReferenceDocumentItem:                   data.ReferenceDocumentItem,
			ExternalReferenceDocument:               data.ExternalReferenceDocument,
			ExternalReferenceDocumentItem:           data.ExternalReferenceDocumentItem,
			TaxCode:                                 data.TaxCode,
			TaxRate:                                 data.TaxRate,
			CountryOfOrigin:                         data.CountryOfOrigin,
			CountryOfOriginLanguage:                 data.CountryOfOriginLanguage,
			Equipment:                         		 data.Equipment,
			ItemPaymentRequisitionIsCreated:         data.ItemPaymentRequisitionIsCreated,
			ItemIsCleared:                           data.ItemIsCleared,
			ItemPaymentBlockStatus:                  data.ItemPaymentBlockStatus,
			CreationDate:                            data.CreationDate,
			CreationTime:                            data.CreationTime,
			LastChangeDate:                          data.LastChangeDate,
			LastChangeTime:                          data.LastChangeTime,
			IsCancelled:                             data.IsCancelled,
		})
	}

	return &item, nil
}

func ConvertToItemPricingElement(rows *sql.Rows) (*[]ItemPricingElement, error) {
	defer rows.Close()
	itemPricingElement := make([]ItemPricingElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemPricingElement{}

		err := rows.Scan(
			&pm.InvoiceDocument,
			&pm.InvoiceDocumentItem,
			&pm.PricingProcedureCounter,
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemPricingElement, err
		}

		data := pm
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			InvoiceDocument:            data.InvoiceDocument,
			InvoiceDocumentItem:        data.InvoiceDocumentItem,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:  					data.Buyer,
			Seller:  					data.Seller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionRateValueUnit:     data.ConditionRateValueUnit,
			ConditionScaleQuantity:     data.ConditionScaleQuantity,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
			CreationDate:               data.CreationDate,
			CreationTime:               data.CreationTime,
			LastChangeDate:             data.LastChangeDate,
			LastChangeTime:             data.LastChangeTime,
			IsCancelled:                data.IsCancelled,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemPricingElement, nil
	}

	return &itemPricingElement, nil
}
