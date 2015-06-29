package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Doc struct {
	XMLName         xml.Name        `xml:"Comprobante"`
	Tipo            string          `xml:"tipoDeComprobante,attr"`
	Version         string          `xml:"version,attr"`
	Serie           string          `xml:"serie,attr"`
	Folio           string          `xml:"folio,attr"`
	Fecha           string          `xml:"fecha,attr"`
	Moneda          string          `xml:"Moneda,attr"`
	TipoCambio      string          `xml:"TipoCambio,attr"`
	Total           string          `xml:"total,attr"`
	SubTotal        string          `xml:"subTotal,attr"`
	MetodoDePago    string          `xml:"metodoDePago,attr"`
	LugarExpedicion string          `xml:"LugarExpedicion,attr"`
	Emisor          CFDIEmisor      `xml:"Emisor"`
	Receptor        CFDIReceptor    `xml:"Receptor"`
	Conceptos       []CFDIConcepto  `xml:"Conceptos>Concepto"`
	Impuestos       CFDIImpuestos   `xml:"Impuestos"`
	Complemento     CFDIComplemento `xml:"Complemento"`
	Addenda         CFDIAddenda     `xml:"Addenda"`
}

type CFDIImpuestos struct {
	XMLName   xml.Name      `xml:"Impuestos"`
	Total     string        `xml:"totalImpuestosTrasladados,attr"`
	Traslados CFDITraslados `xml:"Traslados"`
}

type CFDITraslados struct {
	XMLName  xml.Name     `xml:"Traslados"`
	Traslado CFDITraslado `xml:"Traslado"`
}

type CFDITraslado struct {
	XMLName xml.Name `xml:"Traslado"`
	Importe string   `xml:"importe,attr"`
}

type CFDIAddenda struct {
	XMLName            xml.Name               `xml:"Addenda"`
	AddendaBuzonFiscal AddendaBuzonFiscalNode `xml:"AddendaBuzonFiscal"`
}

type AddendaBuzonFiscalNode struct {
	XMLName xml.Name `xml:"AddendaBuzonFiscal"`
	CFD     CFDNode  `xml:"CFD"`
}

type CFDNode struct {
	XMLName xml.Name `xml:"CFD"`
	RefID   string   `xml:"refID,attr"`
}

type CFDIEmisor struct {
	XMLName xml.Name `xml:"Emisor"`
	RFC     string   `xml:"rfc,attr"`
}

type CFDIReceptor struct {
	XMLName xml.Name `xml:"Receptor"`
	RFC     string   `xml:"rfc,attr"`
	Nombre  string   `xml:"nombre,attr"`
}

type CFDIConcepto struct {
	XMLName          xml.Name `xml:"Concepto"`
	Descripcion      string   `xml:"descripcion,attr"`
	NoIdentificacion string   `xml:"noIdentificacion,attr"`
	Cantidad         string   `xml:"cantidad,attr"`
	Unidad           string   `xml:"unidad,attr"`
	ValorUnitario    string   `xml:"valorUnitario,attr"`
	Importe          string   `xml:"importe,attr"`
}

type CFDIComplemento struct {
	XMLName             xml.Name               `xml:"Complemento"`
	TimbreFiscalDigital TFDTimbreFiscalDigital `xml:"TimbreFiscalDigital"`
	Nomina              NominaNomina           `xml:"Nomina"`
}

type TFDTimbreFiscalDigital struct {
	XMLName           xml.Name `xml:"TimbreFiscalDigital"`
	NumeroCertificado string   `xml:"noCertificadoSAT,attr"`
	FechaTimbrado     string   `xml:"FechaTimbrado,attr"`
	UUID              string   `xml:"UUID,attr"`
}

type NominaNomina struct {
	XMLName          xml.Name `xml:"Nomina"`
	FechaInicialPago string   `xml:"FechaInicialPago,attr"`
	FechaFinalPago   string   `xml:"FechaFinalPago,attr"`
}

func (d Doc) NumeroDeFactura() string {
	return fmt.Sprintf("%s-%s", d.Serie, d.Folio)
}

func parseXml(doc []byte) Doc {
	var query Doc
	xml.Unmarshal(doc, &query)
	return query
}

func EncodeAsRows(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	rawContent, _ := ioutil.ReadAll(file)
	cfdi := parseXml(rawContent)

	var records []string
	var record = []string{cfdi.Complemento.TimbreFiscalDigital.UUID}
	records = append(records, strings.Join(record, "\t"))
	return records
}

func EncodeHeaders() string {
	var headerList = []string{
		"UUID",
	}
	return strings.Join(headerList, "\t")
}
