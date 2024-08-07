// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"alpinskiy/statshouse-test/gen2/basictl"
)

var _ = basictl.NatWrite

type True struct {
}

func (True) TLName() string { return "true" }
func (True) TLTag() uint32  { return 0x3fedd339 }

func (item *True) Reset() {}

func (item *True) Read(w []byte) (_ []byte, err error) { return w, nil }

// This method is general version of Write, use it instead!
func (item *True) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *True) Write(w []byte) []byte {
	return w
}

func (item *True) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x3fedd339); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *True) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *True) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x3fedd339)
	return item.Write(w)
}
