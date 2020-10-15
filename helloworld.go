package main

import (
	"context"
	"github.com/codenotary/immudb/pkg/api/schema"
	immusdk "github.com/codenotary/immudb/pkg/client"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	immuclient, err := immusdk.NewImmuClient(immusdk.DefaultOptions().WithAddress("127.0.0.1").WithPort(3322))
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
	if err != nil {
		log.Fatal(err)
	}

	md = metadata.Pairs("authorization", "Bearer "+respUse.Token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	i, err := immuclient.Set(ctx, []byte(`aaa`), []byte(`bbb`))
	if err != nil {
		log.Fatal(err)
	}
	println(i)


}

