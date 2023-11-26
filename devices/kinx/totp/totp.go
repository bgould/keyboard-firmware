/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// basic, TinyGo-compatible remix of https://github.com/pquerna/otp
package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
	"strconv"
	"strings"
	"time"
)

const debug = false

type errorString string

var DefaultOpts = ValidateOpts{
	Period:    30,
	Skew:      1,
	Digits:    DigitsSix,
	Algorithm: AlgorithmSHA1,
}

func (s errorString) Error() string {
	return string(s)
}

// ValidateOpts provides options for ValidateCustom().
type ValidateOpts struct {
	// Number of seconds a TOTP hash is valid for. Defaults to 30 seconds.
	Period uint
	// Periods before or after the current time to allow.  Value of 1 allows up to Period
	// of either side of the specified time.  Defaults to 0 allowed skews.  Values greater
	// than 1 are likely sketchy.
	Skew uint
	// Digits as part of the input. Defaults to 6.
	Digits Digits
	// Algorithm to use for HMAC. Defaults to SHA1.
	Algorithm Algorithm
}

// GenerateCode creates a TOTP token using the current time.
// A shortcut for GenerateCodeCustom, GenerateCode uses a configuration
// that is compatible with Google-Authenticator and most clients.
func GenerateCode(secret string, t time.Time) (string, error) {
	return GenerateCodeCustom(secret, t, DefaultOpts)
}

func TimeBasedCounter(t time.Time, period uint) uint64 {
	return uint64(math.Floor(float64(t.Unix()) / float64(period)))
}

// GenerateCodeCustom takes a timepoint and produces a passcode using a
// secret and the provided opts. (Under the hood, this is making an adapted
// call to hotp.GenerateCodeCustom)
func GenerateCodeCustom(secret string, t time.Time, opts ValidateOpts) (passcode string, err error) {
	if opts.Period == 0 {
		opts.Period = 30
	}
	counter := TimeBasedCounter(t, opts.Period)
	passcode, err = generateCodeHTOP(secret, counter, opts)
	if err != nil {
		return "", err
	}
	return passcode, nil
}

// GenerateCodeCustom uses a counter and secret value and options struct to
// create a passcode.
func generateCodeHTOP(secret string, counter uint64, opts ValidateOpts) (passcode string, err error) {
	//Set default value
	if opts.Digits == 0 {
		opts.Digits = DigitsSix
	}
	// As noted in issue #10 and #17 this adds support for TOTP secrets that are
	// missing their padding.
	secret = strings.TrimSpace(secret)
	if n := len(secret) % 8; n != 0 {
		secret = secret + strings.Repeat("=", 8-n)
	}

	// As noted in issue #24 Google has started producing base32 in lower case,
	// but the StdEncoding (and the RFC), expect a dictionary of only upper case letters.
	secret = strings.ToUpper(secret)

	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", errorString("invalid base32")
	}

	buf := make([]byte, 8)
	mac := hmac.New(opts.Algorithm.Hash, secretBytes)
	binary.BigEndian.PutUint64(buf, counter)
	if debug {
		fmt.Printf("counter=%v\n", counter)
		fmt.Printf("buf=%v\n", buf)
	}

	mac.Write(buf)
	sum := mac.Sum(nil)

	// "Dynamic truncation" in RFC 4226
	// http://tools.ietf.org/html/rfc4226#section-5.4
	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))

	l := opts.Digits.Length()
	mod := int32(value % int64(math.Pow10(l)))

	if debug {
		fmt.Printf("offset=%v\n", offset)
		fmt.Printf("value=%v\n", value)
		fmt.Printf("mod'ed=%v\n", mod)
	}

	return opts.Digits.Format(mod), nil
}

// Digits represents the number of digits present in the
// user's OTP passcode. Six and Eight are the most common values.
type Digits int

const (
	DigitsSix   Digits = 6
	DigitsEight Digits = 8
)

const zeroes = "00000000"

// Format converts an integer into the zero-filled size for this Digits.
func (d Digits) Format(in int32) string {
	intstr := strconv.Itoa(int(in))
	return zeroes[0:(d.Length()-len(intstr))] + intstr
}

// Length returns the number of characters for this Digits.
func (d Digits) Length() int {
	return int(d)
}

func (d Digits) String() string {
	return fmt.Sprintf("%d", d)
}

// Algorithm represents the hashing function to use in the HMAC
// operation needed for OTPs.
type Algorithm int

const (
	// AlgorithmSHA1 should be used for compatibility with Google Authenticator.
	//
	// See https://github.com/pquerna/otp/issues/55 for additional details.
	AlgorithmSHA1 Algorithm = iota
	// AlgorithmSHA256
	// AlgorithmSHA512
	// AlgorithmMD5
)

func (a Algorithm) String() string {
	switch a {
	case AlgorithmSHA1:
		return "SHA1"
		// case AlgorithmSHA256:
		// 	return "SHA256"
		// case AlgorithmSHA512:
		// 	return "SHA512"
		// case AlgorithmMD5:
		// 	return "MD5"
	}
	panic("unreached")
}

func (a Algorithm) Hash() hash.Hash {
	switch a {
	case AlgorithmSHA1:
		return sha1.New()
		// case AlgorithmSHA256:
		// 	return sha256.New()
		// case AlgorithmSHA512:
		// 	return sha512.New()
		// case AlgorithmMD5:
		// 	return md5.New()
	}
	panic("unreached")
}
