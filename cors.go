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
	"net/url"
	"strings"

	"github.com/gorilla/handlers"
)

// ---------------------------------------------------------------------------------------
//  public functions
// ---------------------------------------------------------------------------------------

// CORS creates a gorilla CORS adapter.
func CORS(origins ...string) Adapter {
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PATCH", "DELETE"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	validator := handlers.AllowedOriginValidator(func(requestOrigin string) bool {
		originUrl, err := url.Parse(requestOrigin)
		if err != nil {
			return false
		}

		for _, origin := range origins {
			if strings.HasSuffix(originUrl.Hostname(), origin) {
				return true
			}
		}

		return false
	})

	return handlers.CORS(methods, headers, validator, handlers.AllowCredentials())
}
