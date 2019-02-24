package main

import (
	"context"
	"fmt"

	"github.com/Sugi275/oci-env-configprovider/envprovider"
	"github.com/oracle/oci-go-sdk/dns"
)

func main() {
	zn := "test.enc"
	dn := "_acme-challenge.test.enc"

	client, err := dns.NewDnsClientWithConfigurationProvider(envprovider.GetEnvConfigProvider())
	if err != nil {
		panic(err)
	}

	compartmentid, err := envprovider.GetCompartmentID()
	if err != nil {
		panic(err)
	}

	// DNSのレコードを作成するパラメータを生成
	txttype := "TXT"
	falseFlg := false
	rdata := "testdayo"
	ttl := 30

	recordDetails := dns.RecordDetails{
		Domain:      &dn,
		Rdata:       &rdata,
		Rtype:       &txttype,
		Ttl:         &ttl,
		IsProtected: &falseFlg,
	}

	var recordDetailsList []dns.RecordDetails
	recordDetailsList = append(recordDetailsList, recordDetails)

	updateDomainRecordsDetails := dns.UpdateDomainRecordsDetails{
		Items: recordDetailsList,
	}

	request := dns.UpdateDomainRecordsRequest{
		ZoneNameOrId:               &zn,
		Domain:                     &dn,
		UpdateDomainRecordsDetails: updateDomainRecordsDetails,
		CompartmentId:              &compartmentid,
	}

	ctx := context.Background()
	response, err := client.UpdateDomainRecords(ctx, request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
