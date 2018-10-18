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

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

//order is important!!
//if StatusCode() and MimeType() is used, MimeType has to be called as the first one
func MimeType(res http.ResponseWriter, mime string){
	res.Header().Set("Content-Type", mime)
}

//order is important!!
//if StatusCode() and MimeType() is used, MimeType has to be called as the first one
//if StatusCode() is used the MimeType() has to be called to get a different mime-type as text/pain (settings of Text() and Json() will be ignored)
func StatusCode(res http.ResponseWriter, code int){
	res.WriteHeader(code)
}

func Json(res http.ResponseWriter, structure interface{}){
	value, err := json.Marshal(structure)
	if err != nil {
		log.Println("error on json response: ", err)
	}else{
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(res, string(value))
	}
}

func Text(res http.ResponseWriter, value string){
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(res, value)
}

