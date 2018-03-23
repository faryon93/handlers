package handlers

// handlers
// Copyright (C) 2018 Maximilian Pachl

// MIT License
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// ---------------------------------------------------------------------------------------
//  imports
// ---------------------------------------------------------------------------------------

import (
	"net/http"

	"github.com/spf13/viper"
)

// ---------------------------------------------------------------------------------------
//  public functions
// ---------------------------------------------------------------------------------------

// Keyed restricts the execution of fn() to requests which
// have the query parameter "key" matched with the user supplied key.
func Keyed(key string) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestKey := r.URL.Query().Get("key")
			if requestKey != key {
				http.Error(w, "forbidded", http.StatusForbidden)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}

// Enabled denys access to fn() if the value of the given confKey is false.
// The value is resolved via viper.GetBool(confKey).
func Enabled(confKey string) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if confKey != "" && !viper.GetBool(confKey) {
				http.Error(w, "handler not enabled", http.StatusNotFound)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}