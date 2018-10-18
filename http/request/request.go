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

package request

import (
	"net/http"
	"encoding/json"
	"bytes"
	"log"
)

func Get(url string, result interface{}) (err error){
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(result)
}


func Post(url string, body interface{}, result interface{}) (err error, payload string, code int){
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(body)
	if err != nil {
		return
	}
	resp, err := http.Post(url, "application/json", b)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	payload = buf.String()
	code = resp.StatusCode
	if result != nil {
		err = json.NewDecoder(buf).Decode(result)
		if err != nil {
			b2 := new(bytes.Buffer)
			json.NewEncoder(b2).Encode(body)

			log.Println("json pars error on post to '"+url+"'\n", "send body:\n", b2.String(), "\nreceived body:\n", payload)
			log.Println(err)
		}
	}
	return
}


func Put(url string, body interface{}, result interface{}) (err error, payload string, code int){
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(body)
	if err != nil {
		return
	}
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, b)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	payload = buf.String()
	code = resp.StatusCode
	if result != nil {
		err = json.NewDecoder(buf).Decode(result)
		if err != nil {
			b2 := new(bytes.Buffer)
			json.NewEncoder(b2).Encode(body)

			log.Println("json pars error on put to '"+url+"'\n", "send body:\n", b2.String(), "\nreceived body:\n", payload)
			log.Println(err)
		}
	}
	return
}