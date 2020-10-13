package main

import (
	"context"
	"fmt"
	"github.com/codenotary/immudb/pkg/api/schema"
	immusdk "github.com/codenotary/immudb/pkg/client"
	sdk "github.com/vchain-us/ledger-compliance-go/grpcclient"
	"google.golang.org/grpc/metadata"

	"log"
	"time"
)

func main() {
	benchmarkWork := 1000
	log.Printf("Ledger compliance vs Immudb SDK. Each loop are %d safeset", benchmarkWork)
	for loop:= 0; loop < 5; loop ++ {

		client := sdk.NewLcClient(sdk.ApiKey("eeicdeqliaxfpiezgtiqljfhdhvbjnrgyvml"), sdk.Host("localhost"), sdk.Port(3324))
		err := client.Connect()
		if err != nil {
			log.Fatal(err)
		}
		start := time.Now()

		for i:=0; i<=benchmarkWork;i++{
			_, err := client.SafeSet(context.Background(), []byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("%d", i)))
			if err != nil {
				log.Fatal(err)
			}
		}
		elapsed := time.Since(start)
		log.Printf("Elapsed %s on lc \t\tloop %d", elapsed, loop)

		//log.Printf("Immudb starting on loop %d", loop)
		startImmu := time.Now()
		immuclient, err := immusdk.NewImmuClient(immusdk.DefaultOptions())
		if err != nil {
			log.Fatal(err)
		}
		resp, err := immuclient.Login(context.Background(), []byte(`immudb`), []byte(`immudb`))
		if err != nil {
			log.Fatal(err)
		}
		md := metadata.Pairs("authorization", "Bearer "+resp.Token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		respUse, err := immuclient.UseDatabase(ctx, &schema.Database{
			Databasename: "defaultdb",
		})
		md = metadata.Pairs("authorization", "Bearer "+respUse.Token)
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		for i:=0; i<=benchmarkWork;i++{
			_, err := immuclient.SafeSet(ctx, []byte(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("%d", i)))
			if err != nil {
				log.Fatal(err)
			}
		}
		elapsedImmu := time.Since(startImmu)
		log.Printf("Elapsed %s on Immudb \tloop %d", elapsedImmu, loop)
	}

}

