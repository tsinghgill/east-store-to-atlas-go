package main

import (
	"context"
	"log"

	// Dependencies of Turbine
	"github.com/meroxa/turbine-go/pkg/turbine"
	"github.com/meroxa/turbine-go/pkg/turbine/cmd"
)

func main() {
	cmd.Start(App{})
}

var _ turbine.App = (*App)(nil)

type App struct{}

func (a App) Run(v turbine.Turbine) error {

	source, err := v.Resources("east-store-mongo")
	if err != nil {
		return err
	}

	rr, err := source.RecordsWithContext(context.Background(), "medicine", turbine.ConnectionOptions{
		{Field: "transforms", Value: "unwrap"},
		{Field: "transforms.unwrap.type", Value: "io.debezium.connector.mongodb.transforms.ExtractNewDocumentState"},
	})
	if err != nil {
		return err
	}

	res, err := v.Process(rr, Anonymize{})
	if err != nil {
		return err
	}

	dest, err := v.Resources("meroxa-atlas")
	if err != nil {
		return err
	}

	err = dest.WriteWithConfig(res, "aggregated_medicine", turbine.ConnectionOptions{
		{Field: "max.batch.size", Value: "1"},
	})
	if err != nil {
		return err
	}

	return nil
}

type Anonymize struct{}

func (f Anonymize) Process(stream []turbine.Record) []turbine.Record {
	for i, record := range stream {
		// log.Printf("Processing record %d: %+v\n", i, record) // Logging the record details
		log.Printf("Payload: \n%s\n", record.Payload) // Logging the payload

		log.Printf("Setting StoreID for East Store to 002")
		err := record.Payload.Set("storeId", "002")
		if err != nil {
			log.Println("error setting value: ", err)
			continue
		}

		stream[i] = record
	}
	return stream
}
