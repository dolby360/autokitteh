// @generated by protoc-gen-connect-es v1.1.4 with parameter "target=ts"
// @generated from file autokitteh/integration_provider/v1/svc.proto (package autokitteh.integration_provider.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CallRequest, CallResponse, GetRequest, GetResponse, ListRequest, ListResponse } from "./svc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Implemented by integration providers, to respond to autokitteh.
 * This service may be in charge of more than one integration, in which
 * case it's responsible for managing integration IDs and connection mappings.
 * In addition, autokitteh also implements this service - but as a router
 * for all registered integrations, instead of an actual integration.
 *
 * @generated from service autokitteh.integration_provider.v1.IntegrationProviderService
 */
export const IntegrationProviderService = {
  typeName: "autokitteh.integration_provider.v1.IntegrationProviderService",
  methods: {
    /**
     * Static declaration(s) of functions and values exposed to autokitteh.
     *
     * @generated from rpc autokitteh.integration_provider.v1.IntegrationProviderService.Get
     */
    get: {
      name: "Get",
      I: GetRequest,
      O: GetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.integration_provider.v1.IntegrationProviderService.List
     */
    list: {
      name: "List",
      I: ListRequest,
      O: ListResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Relay an API method call from the autokitteh runtime engine,
     * via the autokitteh connection manager, to the integration
     * provider, and then relay back the API's response.
     *
     * @generated from rpc autokitteh.integration_provider.v1.IntegrationProviderService.Call
     */
    call: {
      name: "Call",
      I: CallRequest,
      O: CallResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

