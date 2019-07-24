package handlers

// handlers
// Copyright (C) 2019 Maximilian Pachl

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
    "net/http/httptest"
    "strings"
    "testing"
)

// ---------------------------------------------------------------------------------------
//  constants
// ---------------------------------------------------------------------------------------

const (
    BodyPassed = "passed"
)


// ---------------------------------------------------------------------------------------
//  tests
// ---------------------------------------------------------------------------------------

func TestRecaptchaNoKey(t *testing.T) {
    rr := httptest.NewRecorder()
    req, err := http.NewRequest(http.MethodGet, "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    Recaptcha("")(GetTestHandler()).ServeHTTP(rr, req)
    if rr.Code != http.StatusOK {
        t.Errorf("excepted http status to be %d but is %d", http.StatusOK, rr.Code)
        return
    }

    if rr.Body.String() != BodyPassed {
        t.Errorf("excepted body \"%s\" but got \"%s\"",
            BodyPassed, rr.Body.String())
        return
    }
}

func TestRecaptchaInvalidResponse(t *testing.T) {
    bodies := []string{
        "",
        "{}",
        "{\"test\": 12}",
        "{\"g-recaptcha-response\": \"test\"}",
    }

    for _, body := range bodies {
        rr := httptest.NewRecorder()
        req, err := http.NewRequest(http.MethodGet, "/", strings.NewReader(body))
        if err != nil {
            t.Fatal(err)
        }

        Recaptcha("test")(GetTestHandler()).ServeHTTP(rr, req)
        if rr.Code != http.StatusBadRequest {
            t.Errorf("excepted http status to be %d but is %d: middleware did not prevent handler exeution",
                http.StatusBadRequest, rr.Code)
            return
        }

        if rr.Body.String() == BodyPassed {
            t.Errorf("excepted body \"%s\" but got \"%s\": middleware did not prevent handler exeution",
                BodyPassed, rr.Body.String())
            return
        }
    }
}

// ---------------------------------------------------------------------------------------
//  helpers
// ---------------------------------------------------------------------------------------

// GetTestHandler returns a http.HandlerFunc for testing http middleware
func GetTestHandler() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, _ = w.Write([]byte(BodyPassed))
    })
}