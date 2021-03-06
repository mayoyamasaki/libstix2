// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
LangProperty - A property used by one or more STIX objects that
captures the lang string as defined in RFC 5646. This is used to record the
language that a given object is using.
*/
type LangProperty struct {
	Lang string `json:"lang,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LangProperty
// ----------------------------------------------------------------------

/*
SetLang - This method takes in a string value representing an ISO 639-2
encoded language code as defined in RFC 5646 and updates the lang property.
*/
func (o *LangProperty) SetLang(s string) error {
	o.Lang = s
	return nil
}

/*
GetLang - This method returns the current language code for a given object.
*/
func (o *LangProperty) GetLang() string {
	return o.Lang
}
