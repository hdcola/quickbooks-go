// Copyright (c) 2018, Randy Westlund. All rights reserved.
// This code is under the BSD-2-Clause license.

package quickbooks

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Invoice represents a QuickBooks Invoice object.
type Invoice struct {
	Id                           string        `json:"Id,omitempty"`
	SyncToken                    string        `json:"SyncToken,omitempty"`
	MetaData                     MetaData      `json:"MetaData,omitempty"`
	CustomField                  []CustomField `json:"CustomField,omitempty"`
	DocNumber                    string        `json:"DocNumber,omitempty"`
	TxnDate                      Date          `json:"TxnDate,omitempty"`
	DepartmentRef                ReferenceType `json:"DepartmentRef,omitempty"`
	PrivateNote                  string        `json:"PrivateNote,omitempty"`
	LinkedTxn                    []LinkedTxn   `json:"LinkedTxn,omitempty"`
	Line                         []Line        `json:"Line"`
	TxnTaxDetail                 TxnTaxDetail  `json:"TxnTaxDetail,omitempty"`
	CustomerRef                  ReferenceType `json:"CustomerRef"`
	CustomerMemo                 MemoRef       `json:"CustomerMemo,omitempty"`
	BillAddr                     PhysicalAddress `json:"BillAddr,omitempty"`
	ShipAddr                     PhysicalAddress `json:"ShipAddr,omitempty"`
	ShipFromAddr                 PhysicalAddress `json:"ShipFromAddr,omitempty"`
	ClassRef                     ReferenceType   `json:"ClassRef,omitempty"`
	SalesTermRef                 ReferenceType   `json:"SalesTermRef,omitempty"`
	DueDate                      Date            `json:"DueDate,omitempty"`
	GlobalTaxCalculation         string          `json:"GlobalTaxCalculation,omitempty"`
	ShipMethodRef                ReferenceType   `json:"ShipMethodRef,omitempty"`
	ShipDate                     Date            `json:"ShipDate,omitempty"`
	TrackingNum                  string          `json:"TrackingNum,omitempty"`
	TotalAmt                     json.Number     `json:"TotalAmt,omitempty"`
	CurrencyRef                  ReferenceType   `json:"CurrencyRef,omitempty"`
	ExchangeRate                 json.Number     `json:"ExchangeRate,omitempty"`
	HomeTotalAmt                 json.Number     `json:"HomeTotalAmt,omitempty"`
	HomeBalance                  json.Number     `json:"HomeBalance,omitempty"`
	ApplyTaxAfterDiscount        bool            `json:"ApplyTaxAfterDiscount,omitempty"`
	PrintStatus                  string          `json:"PrintStatus,omitempty"`
	EmailStatus                  string          `json:"EmailStatus,omitempty"`
	BillEmail                    EmailAddress    `json:"BillEmail,omitempty"`
	BillEmailCc                  EmailAddress    `json:"BillEmailCc,omitempty"`
	BillEmailBcc                 EmailAddress    `json:"BillEmailBcc,omitempty"`
	DeliveryInfo                 any             `json:"DeliveryInfo,omitempty"`
	Balance                      json.Number     `json:"Balance,omitempty"`
	TxnSource                    string          `json:"TxnSource,omitempty"`
	AllowOnlineCreditCardPayment bool            `json:"AllowOnlineCreditCardPayment,omitempty"`
	AllowOnlineACHPayment        bool            `json:"AllowOnlineACHPayment,omitempty"`
	Deposit                      json.Number     `json:"Deposit,omitempty"`
	DepositToAccountRef          ReferenceType   `json:"DepositToAccountRef,omitempty"`
	TransactionLocationType      string          `json:"TransactionLocationType,omitempty"`
	ProjectRef                   ReferenceType   `json:"ProjectRef,omitempty"`
	InvoiceLink                  string          `json:"InvoiceLink,omitempty"`
	RecurDataRef                 ReferenceType   `json:"RecurDataRef,omitempty"`
	TaxExemptionRef              ReferenceType   `json:"TaxExemptionRef,omitempty"`
	FreeFormAddress              bool            `json:"FreeFormAddress,omitempty"`
	AllowOnlinePayment           bool            `json:"AllowOnlinePayment,omitempty"`
	AllowIPNPayment              bool            `json:"AllowIPNPayment,omitempty"`
}

// DeliveryInfo represents email delivery information returned when a request
// has been made to deliver email with the send operation.
// The specific fields are not documented in the API, so this can be any structure.
type DeliveryInfo = any

type LinkedTxn struct {
	TxnID     string `json:"TxnId"`
	TxnType   string `json:"TxnType"`
	TxnLineId string `json:"TxnLineId,omitempty"`
}

type TxnTaxDetail struct {
	TxnTaxCodeRef ReferenceType `json:"TxnTaxCodeRef,omitempty"`
	TotalTax      json.Number   `json:"TotalTax,omitempty"`
	TaxLine       []Line        `json:"TaxLine,omitempty"`
}

type AccountBasedExpenseLineDetail struct {
	AccountRef ReferenceType `json:"AccountRef"`
	TaxAmount  json.Number   `json:"TaxAmount,omitempty"`
	// TaxInclusiveAmt json.Number              `json:"TaxInclusiveAmt,omitempty"`
	// ClassRef        ReferenceType `json:"ClassRef,omitempty"`
	// TaxCodeRef      ReferenceType `json:"TaxCodeRef,omitempty"`
	// MarkupInfo MarkupInfo `json:"MarkupInfo,omitempty"`
	// BillableStatus BillableStatusEnum       `json:"BillableStatus,omitempty"`
	// CustomerRef    ReferenceType `json:"CustomerRef,omitempty"`
}

type Line struct {
	Id                            string `json:"Id,omitempty"`
	LineNum                       int    `json:"LineNum,omitempty"`
	Description                   string `json:"Description,omitempty"`
	Amount                        json.Number `json:"Amount"`
	DetailType                    string `json:"DetailType"`
	SalesItemLineDetail           SalesItemLineDetail           `json:"SalesItemLineDetail,omitempty"`
	DiscountLineDetail            DiscountLineDetail            `json:"DiscountLineDetail,omitempty"`
	TaxLineDetail                 TaxLineDetail                 `json:"TaxLineDetail,omitempty"`
	SubTotalLineDetail            SubTotalLineDetail            `json:"SubTotalLineDetail,omitempty"`
	DescriptionOnlyLineDetail     DescriptionOnlyLineDetail     `json:"DescriptionOnlyLineDetail,omitempty"`
	GroupLineDetail               GroupLineDetail               `json:"GroupLineDetail,omitempty"`
	AccountBasedExpenseLineDetail AccountBasedExpenseLineDetail `json:"AccountBasedExpenseLineDetail,omitempty"`
}

// TaxLineDetail ...
type TaxLineDetail struct {
	PercentBased     bool        `json:"PercentBased,omitempty"`
	NetAmountTaxable json.Number `json:"NetAmountTaxable,omitempty"`
	// TaxInclusiveAmount json.Number `json:"TaxInclusiveAmount,omitempty"`
	// OverrideDeltaAmount
	TaxPercent json.Number   `json:"TaxPercent,omitempty"`
	TaxRateRef ReferenceType `json:"TaxRateRef"`
}

// SalesItemLineDetail ...
type SalesItemLineDetail struct {
	ItemRef   ReferenceType `json:"ItemRef,omitempty"`
	ClassRef  ReferenceType `json:"ClassRef,omitempty"`
	UnitPrice json.Number   `json:"UnitPrice,omitempty"`
	// MarkupInfo
	Qty             float32       `json:"Qty,omitempty"`
	ItemAccountRef  ReferenceType `json:"ItemAccountRef,omitempty"`
	TaxCodeRef      ReferenceType `json:"TaxCodeRef,omitempty"`
	ServiceDate     Date          `json:"ServiceDate,omitempty"`
	TaxInclusiveAmt json.Number   `json:"TaxInclusiveAmt,omitempty"`
	DiscountRate    json.Number   `json:"DiscountRate,omitempty"`
	DiscountAmt     json.Number   `json:"DiscountAmt,omitempty"`
}

// DiscountLineDetail ...
type DiscountLineDetail struct {
	PercentBased    bool    `json:"PercentBased"`
	DiscountPercent float32 `json:"DiscountPercent,omitempty"`
}

// SubTotalLineDetail represents a subtotal line on an invoice
type SubTotalLineDetail struct{}

// DescriptionOnlyLineDetail represents a description-only line on an invoice
type DescriptionOnlyLineDetail struct{}

// GroupLineDetail represents a group line on an invoice
type GroupLineDetail struct {
	GroupName  string  `json:"GroupName,omitempty"`
	LineItems  []Line  `json:"LineItems,omitempty"`
	Quantity   float32 `json:"Quantity,omitempty"`
	UnitPrice  json.Number `json:"UnitPrice,omitempty"`
	ItemRef    ReferenceType `json:"ItemRef,omitempty"`
}

// CreateInvoice creates the given Invoice on the QuickBooks server, returning
// the resulting Invoice object.
func (c *Client) CreateInvoice(invoice *Invoice) (*Invoice, error) {
	var resp struct {
		Invoice Invoice
		Time    Date
	}

	if err := c.post("invoice", invoice, &resp, nil); err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// DeleteInvoice deletes the invoice
//
// If the invoice was already deleted, QuickBooks returns 400 :(
// The response looks like this:
// {"Fault":{"Error":[{"Message":"Object Not Found","Detail":"Object Not Found : Something you're trying to use has been made inactive. Check the fields with accounts, invoices, items, vendors or employees.","code":"610","element":""}],"type":"ValidationFault"},"time":"2018-03-20T20:15:59.571-07:00"}
//
// This is slightly horrifying and not documented in their API. When this
// happens we just return success; the goal of deleting it has been
// accomplished, just not by us.
func (c *Client) DeleteInvoice(invoice *Invoice) error {
	if invoice.Id == "" || invoice.SyncToken == "" {
		return errors.New("missing id/sync token")
	}

	return c.post("invoice", invoice, nil, map[string]string{"operation": "delete"})
}

// FindInvoices gets the full list of Invoices in the QuickBooks account.
func (c *Client) FindInvoices() ([]Invoice, error) {
	var resp struct {
		QueryResponse struct {
			Invoices      []Invoice `json:"Invoice"`
			MaxResults    int
			StartPosition int
			TotalCount    int
		}
	}

	if err := c.query("SELECT COUNT(*) FROM Invoice", &resp); err != nil {
		return nil, err
	}

	if resp.QueryResponse.TotalCount == 0 {
		return nil, errors.New("no invoices could be found")
	}

	invoices := make([]Invoice, 0, resp.QueryResponse.TotalCount)

	for i := 0; i < resp.QueryResponse.TotalCount; i += queryPageSize {
		query := "SELECT * FROM Invoice ORDERBY Id STARTPOSITION " + strconv.Itoa(i+1) + " MAXRESULTS " + strconv.Itoa(queryPageSize)

		if err := c.query(query, &resp); err != nil {
			return nil, err
		}

		if resp.QueryResponse.Invoices == nil {
			return nil, errors.New("no invoices could be found")
		}

		invoices = append(invoices, resp.QueryResponse.Invoices...)
	}

	return invoices, nil
}

// FindInvoiceById finds the invoice by the given id
func (c *Client) FindInvoiceById(id string) (*Invoice, error) {
	var resp struct {
		Invoice Invoice
		Time    Date
	}

	if err := c.get("invoice/"+id, &resp, nil); err != nil {
		return nil, err
	}

	return &resp.Invoice, nil
}

// QueryInvoices accepts an SQL query and returns all invoices found using it
func (c *Client) QueryInvoices(query string) ([]Invoice, error) {
	var resp struct {
		QueryResponse struct {
			Invoices      []Invoice `json:"Invoice"`
			StartPosition int
			MaxResults    int
		}
	}

	if err := c.query(query, &resp); err != nil {
		return nil, err
	}

	if resp.QueryResponse.Invoices == nil {
		return nil, errors.New("could not find any invoices")
	}

	return resp.QueryResponse.Invoices, nil
}

// SendInvoice sends the invoice to the Invoice.BillEmail if emailAddress is left empty
func (c *Client) SendInvoice(invoiceId string, emailAddress string) error {
	queryParameters := make(map[string]string)

	if emailAddress != "" {
		queryParameters["sendTo"] = emailAddress
	}

	return c.post("invoice/"+invoiceId+"/send", nil, nil, queryParameters)
}

// UpdateInvoice updates the invoice
func (c *Client) UpdateInvoice(invoice *Invoice) (*Invoice, error) {
	if invoice.Id == "" {
		return nil, errors.New("missing invoice id")
	}

	existingInvoice, err := c.FindInvoiceById(invoice.Id)
	if err != nil {
		return nil, err
	}

	invoice.SyncToken = existingInvoice.SyncToken

	payload := struct {
		*Invoice
		Sparse bool `json:"sparse"`
	}{
		Invoice: invoice,
		Sparse:  true,
	}

	var invoiceData struct {
		Invoice Invoice
		Time    Date
	}

	if err = c.post("invoice", payload, &invoiceData, nil); err != nil {
		return nil, err
	}

	return &invoiceData.Invoice, err
}

func (c *Client) VoidInvoice(invoice Invoice) error {
	if invoice.Id == "" {
		return errors.New("missing invoice id")
	}

	existingInvoice, err := c.FindInvoiceById(invoice.Id)
	if err != nil {
		return err
	}

	invoice.SyncToken = existingInvoice.SyncToken

	return c.post("invoice", invoice, nil, map[string]string{"operation": "void"})
}
