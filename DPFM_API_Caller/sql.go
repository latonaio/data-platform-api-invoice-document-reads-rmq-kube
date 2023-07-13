package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-invoice-document-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
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
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemPricingElement *[]dpfm_api_output_formatter.ItemPricingElement
	var address *[]dpfm_api_output_formatter.Address
	var partner *[]dpfm_api_output_formatter.Partner
	var headerDoc *[]dpfm_api_output_formatter.HeaderDoc
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
//		case "Headers":
//			func() {
//				header = c.Headers(mtx, input, output, errs, log)
//			}()
		case "HeadersByBillToParty":
			func() {
				header = c.HeadersByBillToParty(mtx, input, output, errs, log)
			}()
		case "HeadersByBillFromParty":
			func() {
				header = c.HeadersByBillFromParty(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "HeaderDoc":
			func() {
				headerDoc = c.HeaderDoc(mtx, input, output, errs, log)
			}()
		case "Partner":
			func() {
				partner = c.Partner(mtx, input, output, errs, log)
			}()
		case "ItemPricingElement":
			func() {
				itemPricingElement = c.ItemPricingElement(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Item:               item,
		ItemPricingElement: itemPricingElement,
		Address:            address,
		Partner:            partner,
		HeaderDoc:          headerDoc,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	invoiceDocument := input.Header.InvoiceDocument

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_data
		WHERE InvoiceDocument = ? 
		ORDER BY IsCancelled ASC, InvoiceDocument DESC;`, invoiceDocument,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByBillToParty(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) []dpfm_api_output_formatter.Header {
	billToParty 					:= input.Header.BillToParty
	isCancelled						:= input.Header.IsCancelled
	isMarkedForDeletion 			:= input.Header.IsMarkedForDeletion

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_data
		WHERE (
		       BillToParty,
		       IsCancelled,
		       IsMarkedForDeletion
		) = (?, ?, ?);`,
		billToParty,
		isCancelled,
		isMarkedForDeletion,
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

func (c *DPFMAPICaller) HeadersByBillFromParty(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) []dpfm_api_output_formatter.Header {
	billFromParty 					:= input.Header.BillFromParty
	isCancelled						:= input.Header.IsCancelled
	isMarkedForDeletion 			:= input.Header.IsMarkedForDeletion

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_data
		WHERE (
		       BillFromParty,
		       IsCancelled,
		       IsMarkedForDeletion
		) = (?, ?, ?);`,
		billFromParty,
		isCancelled,
		isMarkedForDeletion,
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

//func (c *DPFMAPICaller) Headers(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *[]dpfm_api_output_formatter.Header {
//	where := "WHERE 1 = 1"
//	if input.Header.Payee != nil {
//		where = fmt.Sprintf("%s\nAND Payee = %d", where, *input.Header.Payee)
//	}
//	if input.Header.Payer != nil {
//		where = fmt.Sprintf("%s\nAND Payer = %d", where, *input.Header.Payer)
//	}
//	if input.Header.BillFromParty != nil {
//		where = fmt.Sprintf("%s\nAND BillFromParty = %d", where, *input.Header.BillFromParty)
//	}
//	if input.Header.BillToParty != nil {
//		where = fmt.Sprintf("%s\nAND BillToParty = %d", where, *input.Header.BillToParty)
//	}
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_data
//		` + where + ` ORDER BY IsCancelled ASC, InvoiceDocument DESC;`,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	invoiceDocument := input.Header.InvoiceDocument
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		args = append(args, invoiceDocument, v.InvoiceDocumentItem)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_item_data
		WHERE (InvoiceDocument, InvoiceDocumentItem) IN ( `+repeat+` ) 
		ORDER BY IsCancelled ASC, InvoiceDocument ASC, InvoiceDocumentItem ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	invoiceDocument := input.Header.InvoiceDocument

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_item_data
		WHERE InvoiceDocument = ? 
		ORDER BY IsCancelled ASC, InvoiceDocument ASC, InvoiceDocumentItem ASC;`, invoiceDocument,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
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
) *[]dpfm_api_output_formatter.ItemPricingElement {
	var args []interface{}
	invoiceDocument := input.Header.InvoiceDocument
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		itemPricingElement := v.ItemPricingElement
		for _, w := range itemPricingElement {
			args = append(args, invoiceDocument, v.InvoiceDocumentItem, w.PricingProcedureCounter)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_item_pricing_element_data
		WHERE (InvoiceDocument, InvoiceDocumentItem, PricingProcedureCounter) IN ( `+repeat+` ) 
		ORDER BY InvoiceDocument DESC, InvoiceDocumentItem ASC, PricingProcedureCounter ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	invoiceDocument := input.Header.InvoiceDocument
	partner := input.Header.Partner

	cnt := 0
	for _, v := range partner {
		args = append(args, invoiceDocument, v.PartnerFunction, v.BusinessPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_partner_data
		WHERE (InvoiceDocument, PartnerFunction, BusinessPartner) IN ( `+repeat+` ) 
		ORDER BY InvoiceDocument DESC, BusinessPartner DESC, AddressID DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
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
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	invoiceDocument := input.Header.InvoiceDocument
	address := input.Header.Address

	cnt := 0
	for _, v := range address {
		args = append(args, invoiceDocument, v.AddressID)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_address_data
		WHERE (InvoiceDocument, AddressID) IN ( `+repeat+` ) 
		ORDER BY InvoiceDocument DESC, AddressID DESC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderDoc(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.HeaderDoc {
	var args []interface{}
	invoiceDocument := input.Header.InvoiceDocument
	headerDoc := input.Header.HeaderDoc

	cnt := 0
	for _, v := range headerDoc {
		args = append(args, invoiceDocument, v.DocType, v.DocVersionID, v.DocID)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_invoice_document_header_doc_data
		WHERE (InvoiceDocument, DocType, DocVersionID, DocID) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderDoc(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
