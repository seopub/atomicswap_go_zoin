// Copyright (c) 2015-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package cfgutil

import (
	"strconv"
	"strings"

	xzcutil "seopub/xzcutil"
)

// AmountFlag embeds a xzcutil.Amount and implements the flags.Marshaler and
// Unmarshaler interfaces so it can be used as a config struct field.
type AmountFlag struct {
	xzcutil.Amount
}

// NewAmountFlag creates an AmountFlag with a default xzcutil.Amount.
func NewAmountFlag(defaultValue xzcutil.Amount) *AmountFlag {
	return &AmountFlag{defaultValue}
}

// MarshalFlag satisifes the flags.Marshaler interface.
func (a *AmountFlag) MarshalFlag() (string, error) {
	return a.Amount.String(), nil
}

// UnmarshalFlag satisifes the flags.Unmarshaler interface.
func (a *AmountFlag) UnmarshalFlag(value string) error {
	value = strings.TrimSuffix(value, " BTC")
	valueF64, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	amount, err := xzcutil.NewAmount(valueF64)
	if err != nil {
		return err
	}
	a.Amount = amount
	return nil
}
