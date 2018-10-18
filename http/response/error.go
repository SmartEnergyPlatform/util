/*
 * Copyright 2018 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package response

import "log"

const ERROR_GENERIC = "GENERIC_ERROR"
const ERROR_DEPENDENT_DEVICE_TYPE = "DEPENDENT_DEVICE_TYPE_ERROR"
const ERROR_DEPENDENT_DEVICE_INSTANCE = "DEPENDENT_DEVICE_INSTANCE_ERROR"
const ERROR_INCONSISTENT_NEW_ELEMENT = "INCONSISTENT_NEW_ELEMENT_ERROR"

// Default Error-Message
// swagger:response ErrorMessage
type ErrorMessage struct {
	StatusCode int      `json:"status_code,omitempty"`
	Message    string   `json:"message"`
	ErrorCode  string   `json:"error_code,omitempty"`
	Detail     []string `json:"detail,omitempty"`
}

func (this *ResponseHandler) DefaultError(message string, code int) *ResponseHandler {
	return this.Error(ErrorMessage{StatusCode: code, Message: message, ErrorCode: ERROR_GENERIC})
}

func (this *ResponseHandler) Error(msg ErrorMessage) *ResponseHandler {
	log.Println("ERROR: ", msg)
	this.MimeType("application/json; charset=utf-8").Code(msg.StatusCode).Json(msg)
	return this
}
