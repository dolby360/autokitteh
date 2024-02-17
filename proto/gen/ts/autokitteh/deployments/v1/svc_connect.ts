// @generated by protoc-gen-connect-es v1.1.4 with parameter "target=ts"
// @generated from file autokitteh/deployments/v1/svc.proto (package autokitteh.deployments.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ActivateRequest, ActivateResponse, CreateRequest, CreateResponse, DeactivateRequest, DeactivateResponse, DrainRequest, DrainResponse, GetRequest, GetResponse, ListRequest, ListResponse, TestRequest, TestResponse } from "./svc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service autokitteh.deployments.v1.DeploymentsService
 */
export const DeploymentsService = {
  typeName: "autokitteh.deployments.v1.DeploymentsService",
  methods: {
    /**
     * @generated from rpc autokitteh.deployments.v1.DeploymentsService.Create
     */
    create: {
      name: "Create",
      I: CreateRequest,
      O: CreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Activate a deployment, automatically drains all others.
     *
     * @generated from rpc autokitteh.deployments.v1.DeploymentsService.Activate
     */
    activate: {
      name: "Activate",
      I: ActivateRequest,
      O: ActivateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Drain a deployment.
     *
     * @generated from rpc autokitteh.deployments.v1.DeploymentsService.Drain
     */
    drain: {
      name: "Drain",
      I: DrainRequest,
      O: DrainResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Deactivate a deployment - forcefully stops all sessions associated
     * with the deployment.
     *
     * @generated from rpc autokitteh.deployments.v1.DeploymentsService.Deactivate
     */
    deactivate: {
      name: "Deactivate",
      I: DeactivateRequest,
      O: DeactivateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.deployments.v1.DeploymentsService.Test
     */
    test: {
      name: "Test",
      I: TestRequest,
      O: TestResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.deployments.v1.DeploymentsService.List
     */
    list: {
      name: "List",
      I: ListRequest,
      O: ListResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.deployments.v1.DeploymentsService.Get
     */
    get: {
      name: "Get",
      I: GetRequest,
      O: GetResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

