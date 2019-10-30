// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseAutonomousDatabaseRegionalWalletManagementDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	return GetSingularDataSourceItemSchema(DatabaseAutonomousDatabaseRegionalWalletManagementResource(), fieldMap, readSingularDatabaseAutonomousDatabaseRegionalWalletManagement)
}

func readSingularDatabaseAutonomousDatabaseRegionalWalletManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.AutonomousDatabaseWallet
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseRegionalWalletRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDatabaseRegionalWallet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseWallet
	return nil
}

func (s *DatabaseAutonomousDatabaseRegionalWalletManagementDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeRotated != nil {
		s.D.Set("time_rotated", s.Res.TimeRotated.String())
	}

	return nil
}