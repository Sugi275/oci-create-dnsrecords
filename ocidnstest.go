package main

import (
	"context"
	"fmt"

	"github.com/Sugi275/oci-env-configprovider/envprovider"
	"github.com/oracle/oci-go-sdk/dns"
)

func main() {
	fmt.Println("vim-go")

	zn := "test.enc"
	dn := "_acme-challenge.test.enc"

	client, err := dns.NewDnsClientWithConfigurationProvider(envprovider.GetEnvConfigProvider())
	if err != nil {
		panic(err)
	}

	// DNSのレコードを作成するパラメータを生成
	txttype := "TXT"
	falseFlg := false
	rdata := "testdayo"
	ttl := 30

	recordOperation := dns.RecordOperation{
		Rtype:       &txttype,
		Domain:      &dn,
		IsProtected: &falseFlg,
		Rdata:       &rdata,
		Ttl:         &ttl,
	}

	var recordOperationList []dns.RecordOperation
	recordOperationList = append(recordOperationList, recordOperation)
	fmt.Println("debug1: ", recordOperationList)

	patchDomainRecordsDetails := dns.PatchDomainRecordsDetails{
		Items: recordOperationList,
	}

	request := dns.PatchDomainRecordsRequest{
		ZoneNameOrId:              &zn,
		Domain:                    &dn,
		PatchDomainRecordsDetails: patchDomainRecordsDetails,
	}

	ctx := context.Background()
	fmt.Println("debug2: ", request)
	response, err := client.PatchDomainRecords(ctx, request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
