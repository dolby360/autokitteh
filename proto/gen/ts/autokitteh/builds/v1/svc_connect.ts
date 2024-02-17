// @generated by protoc-gen-connect-es v1.1.4 with parameter "target=ts"
// @generated from file autokitteh/builds/v1/svc.proto (package autokitteh.builds.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DownloadRequest, DownloadResponse, GetRequest, GetResponse, ListRequest, ListResponse, RemoveRequest, RemoveResponse, SaveRequest, SaveResponse } from "./svc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service autokitteh.builds.v1.BuildsService
 */
export const BuildsService = {
  typeName: "autokitteh.builds.v1.BuildsService",
  methods: {
    /**
     * @generated from rpc autokitteh.builds.v1.BuildsService.Get
     */
    get: {
      name: "Get",
      I: GetRequest,
      O: GetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.builds.v1.BuildsService.List
     */
    list: {
      name: "List",
      I: ListRequest,
      O: ListResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.builds.v1.BuildsService.Save
     */
    save: {
      name: "Save",
      I: SaveRequest,
      O: SaveResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.builds.v1.BuildsService.Remove
     */
    remove: {
      name: "Remove",
      I: RemoveRequest,
      O: RemoveResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.builds.v1.BuildsService.Download
     */
    download: {
      name: "Download",
      I: DownloadRequest,
      O: DownloadResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

