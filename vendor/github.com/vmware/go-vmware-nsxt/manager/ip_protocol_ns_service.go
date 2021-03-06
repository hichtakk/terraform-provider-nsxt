/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type IpProtocolNsServiceEntry struct {
	ResourceType string `json:"resource_type"`

	ProtocolNumber int64 `json:"protocol_number"`
}

type IpProtocolNsService struct {
	NsService

	NsserviceElement IpProtocolNsServiceEntry `json:"nsservice_element"`
}
