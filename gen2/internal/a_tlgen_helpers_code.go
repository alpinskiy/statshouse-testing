// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"fmt"
	"strconv"

	"github.com/mailru/easyjson/jlexer"
)

type UnionElement struct {
	TLTag    uint32
	TLName   string
	TLString string
}

func ErrorClientWrite(typeName string, err error) error {
	return fmt.Errorf("failed to serialize %s request: %w", typeName, err)
}

func ErrorClientDo(typeName string, network string, actorID int64, address string, err error) error {
	return fmt.Errorf("%s request to %s://%d@%s failed: %w", typeName, network, actorID, address, err)
}

func ErrorClientReadResult(typeName string, network string, actorID int64, address string, err error) error {
	return fmt.Errorf("failed to deserialize %s response from %s://%d@%s: %w", typeName, network, actorID, address, err)
}

func ErrorServerHandle(typeName string, err error) error {
	return fmt.Errorf("failed to handle %s: %w", typeName, err)
}

func ErrorServerRead(typeName string, err error) error {
	return fmt.Errorf("failed to deserialize %s request: %w", typeName, err)
}

func ErrorServerWriteResult(typeName string, err error) error {
	return fmt.Errorf("failed to serialize %s response: %w", typeName, err)
}

func ErrorInvalidEnumTag(typeName string, tag uint32) error {
	return fmt.Errorf("invalid enum %q tag: 0x%x", typeName, tag)
}

func ErrorInvalidUnionTag(typeName string, tag uint32) error {
	return fmt.Errorf("invalid union %q tag: 0x%x", typeName, tag)
}

func ErrorWrongSequenceLength(typeName string, actual int, expected uint32) error {
	return fmt.Errorf("wrong sequence %q length: %d expected: %d", typeName, actual, expected)
}

func ErrorInvalidJSONWithDuplicatingKeys(typeName string, field string) error {
	return fmt.Errorf("invalid json for type %q: %q repeats several times", typeName, field)
}

func ErrorInvalidJSONExcessElement(typeName string, key string) error {
	return fmt.Errorf("invalid json object key %q", key)
}

func Json2ReadUint32(in *jlexer.Lexer, dst *uint32) error {
	if in == nil {
		*dst = 0
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		value, err := strconv.ParseUint(src, 10, 32)
		if err != nil {
			return err
		}
		*dst = uint32(value)
	case jlexer.TokenNumber:
		*dst = in.Uint32()
	default:
		return fmt.Errorf("invalid json for uint32")
	}
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func ErrorInvalidJSON(typeName string, msg string) error {
	return fmt.Errorf("invalid json for type %q - %s", typeName, msg)
}

func Json2ReadString(in *jlexer.Lexer, dst *string) error {
	if in == nil {
		*dst = ""
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		*dst = in.String()
	case jlexer.TokenDelim:
		var findValue = false

		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "base64":
				if findValue {
					return fmt.Errorf("base64 repeats several times")
				}
				*dst = string(in.Bytes())
				findValue = true
			default:
				return fmt.Errorf("unexpected field \"" + key + "\"")
			}

			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}

		if !findValue {
			return fmt.Errorf("base64 is absent")
		}
	default:
		return fmt.Errorf("invalid json for string")
	}
	return nil
}

func Json2ReadFloat64(in *jlexer.Lexer, dst *float64) error {
	if in == nil {
		*dst = 0
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		value, err := strconv.ParseFloat(src, 64)
		if err != nil {
			return err
		}
		*dst = value
	case jlexer.TokenNumber:
		*dst = in.Float64()
	default:
		return fmt.Errorf("invalid json for float64")
	}
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func Json2ReadInt64(in *jlexer.Lexer, dst *int64) error {
	if in == nil {
		*dst = 0
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		src := in.UnsafeString()
		value, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return err
		}
		*dst = value
	case jlexer.TokenNumber:
		*dst = in.Int64()
	default:
		return fmt.Errorf("invalid json for int64")
	}
	if !in.Ok() {
		return in.Error()
	}
	return nil
}

func Json2ReadStringBytes(in *jlexer.Lexer, dst *[]byte) error {
	if in == nil {
		*dst = nil
		return nil
	}

	switch in.CurrentToken() {
	case jlexer.TokenString:
		*dst = append((*dst)[:0], in.String()...)
	case jlexer.TokenDelim:
		var findValue = false

		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "base64":
				if findValue {
					return fmt.Errorf("base64 repeats several times")
				}
				*dst = in.Bytes()
				findValue = true
			default:
				return fmt.Errorf("unexpected field \"" + key + "\"")
			}

			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}

		if !findValue {
			return fmt.Errorf("base64 is absent")
		}
	default:
		return fmt.Errorf("invalid json for string")
	}
	return nil
}
