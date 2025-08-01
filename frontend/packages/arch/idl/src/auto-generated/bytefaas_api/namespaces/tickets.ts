/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
 
// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
/* eslint-disable */
/* tslint:disable */
// @ts-nocheck

import * as common from './common';

export type Int64 = string | number;

export interface ListPipelineTemplatesRequest {
  /** JWT token for authentication */
  'X-Jwt-Token'?: string;
}

export interface ListPipelineTemplatesResponse {
  code?: number;
  data?: Array<common.PipelineTemplateLegacy>;
  error?: string;
}

export interface QueryPipelineTemplateByTypeRequest {
  /** Template type */
  template_type: string;
  /** JWT token for authentication */
  'X-Jwt-Token'?: string;
}

export interface QueryPipelineTemplateByTypeResponse {
  code?: number;
  data?: common.PipelineTemplateLegacy;
  error?: string;
}
/* eslint-enable */
