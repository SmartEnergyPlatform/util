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

import "net/http"

type ResponseHandler struct{
	writer http.ResponseWriter
}

//order is important!!
//if code() and mimeType() is used, mimeType has to be called as the first one
func (this *ResponseHandler) MimeType(mime string) *ResponseHandler{
	MimeType(this.writer, mime)
	return this
}

//order is important!!
//if code() and mimeType() is used, mimeType has to be called as the first one
//if Code() is used the MimeType() has to be called to get a different mime-type as text/pain (settings of Text() and Json() will be ignored)
func (this *ResponseHandler) Code(code int) *ResponseHandler{
	StatusCode(this.writer, code)
	return this
}

func (this *ResponseHandler) Text(message string) *ResponseHandler{
	Text(this.writer, message)
	return this
}

func (this *ResponseHandler) Json(message interface{}) *ResponseHandler{
	Json(this.writer, message)
	return this
}

func To(res http.ResponseWriter) *ResponseHandler{
	return &ResponseHandler{writer:res}
}


