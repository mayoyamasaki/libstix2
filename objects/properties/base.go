// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
BaseProperties - This type includes all of the common properties that are used
by all STIX objects

ObjectIDProperty - A property used by one or more STIX objects that captures
the unique object ID. This is not included in the JSON serialization, but is
used for writing to the database.

SpecVersionProperty - A property used by one or more STIX objects that
captures the STIX specification version.

CreatedByRefProperty - A property used by one or more STIX objects that
captures the STIX identifier of the identity that created this object.
*/
type BaseProperties struct {
	ObjectID int `json:"-"`
	TypeProperty
	SpecVersion string `json:"spec_version,omitempty"`
	IDProperty
	CreatedByRef string `json:"created_by_ref,omitempty"`
	BaseTimestampProperties
}

// ----------------------------------------------------------------------
// Public Methods - BaseProperties
// ----------------------------------------------------------------------

/*
SetObjectID - This method takes in a int64 representing an object ID and
updates the Version property.
*/
func (p *BaseProperties) SetObjectID(i int) error {
	p.ObjectID = i
	return nil
}

/*
GetObjectID - This method returns the object ID value as a int64.
*/
func (p *BaseProperties) GetObjectID() int {
	return p.ObjectID
}

/*
SetSpecVersion20 - This method will set the specification version to 2.0.
*/
func (p *BaseProperties) SetSpecVersion20() error {
	p.SpecVersion = "2.0"
	return nil
}

/*
SetSpecVersion20 - This method will set the specification version to 2.1.
*/
func (p *BaseProperties) SetSpecVersion21() error {
	p.SpecVersion = "2.1"
	return nil
}

/*
SetSpecVersion - This method takes in a string representing a STIX specification
version and updates the Version property.
*/
func (p *BaseProperties) SetSpecVersion(s string) error {
	p.SpecVersion = s
	return nil
}

/*
GetSpecVersion - This method returns the version value as a string.
*/
func (p *BaseProperties) GetSpecVersion() string {
	return p.SpecVersion
}

/*
SetCreatedByRef - This method takes in a string value representing a STIX
identifier and updates the Created By Ref property.
*/
func (p *BaseProperties) SetCreatedByRef(s string) error {
	p.CreatedByRef = s
	return nil
}

/*
GetCreatedByRef - This method returns the STIX identifier for the identity
that created this object.
*/
func (p *BaseProperties) GetCreatedByRef() string {
	return p.CreatedByRef
}

// ----------------------------------------------------------------------
// Other Types
// ----------------------------------------------------------------------

/*
ConfidenceProperty - A property used by one or more STIX objects that
captures the STIX confidence score, which is a value from 0-100.
*/
type ConfidenceProperty struct {
	Confidence int `json:"confidence,omitempty"`
}

/*
DescriptionProperty - A property used by one or more STIX objects that
captures the description for the object as a string.
*/
type DescriptionProperty struct {
	Description string `json:"description,omitempty"`
}

/*
LangProperty - A property used by one or more STIX objects that
captures the lang string as defined in RFC 5646. This is used to record the
language that a given object is using.
*/
type LangProperty struct {
	Lang string `json:"lang,omitempty"`
}

/*
NameProperty - A property used by one or more STIX objects that captures a
vanity name for the STIX object in string format.
*/
type NameProperty struct {
	Name string `json:"name,omitempty"`
}

/*
RevokedProperty - A property used by one or more STIX objects that
captures whether or not this STIX object has been revoked by the object
// creator.
*/
type RevokedProperty struct {
	Revoked bool `json:"revoked,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ConfidenceProperty
// ----------------------------------------------------------------------

/*
SetConfidence - This method takes in an integer representing a STIX
confidence level 0-100 and updates the Confidence property.
*/
func (p *ConfidenceProperty) SetConfidence(i int) error {
	p.Confidence = i
	return nil
}

/*
GetConfidence - This method returns the confidence value as an integer.
*/
func (p *ConfidenceProperty) GetConfidence() int {
	return p.Confidence
}

// ----------------------------------------------------------------------
// Public Methods - DescriptionProperty
// ----------------------------------------------------------------------

/*
SetDescription - This method takes in a string value representing a text
description and updates the description property.
*/
func (p *DescriptionProperty) SetDescription(s string) error {
	p.Description = s
	return nil
}

/*
GetDescription - This method returns the description for an object as a string.
*/
func (p *DescriptionProperty) GetDescription() string {
	return p.Description
}

// ----------------------------------------------------------------------
// Public Methods - LangProperty
// ----------------------------------------------------------------------

/*
SetLang - This method takes in a string value representing an ISO 639-2
encoded language code as defined in RFC 5646 and updates the lang property.
*/
func (p *LangProperty) SetLang(s string) error {
	p.Lang = s
	return nil
}

/*
GetLang - This method returns the current language code for a given object.
*/
func (p *LangProperty) GetLang() string {
	return p.Lang
}

// ----------------------------------------------------------------------
// Public Methods - NameProperty
// ----------------------------------------------------------------------

/*
SetName - This method takes in a string value representing a name of the object
and updates the name property.
*/
func (p *NameProperty) SetName(s string) error {
	p.Name = s
	return nil
}

/*
GetName - This method returns the current name of the object.
*/
func (p *NameProperty) GetName() string {
	return p.Name
}

// ----------------------------------------------------------------------
// Public Methods - RevokedProperty
// ----------------------------------------------------------------------

/*
SetRevoked - This method sets the revoked boolean to true
*/
func (p *RevokedProperty) SetRevoked() error {
	p.Revoked = true
	return nil
}

/*
GetRevoked - This method returns the current value of the revoked property.
*/
func (p *RevokedProperty) GetRevoked() bool {
	return p.Revoked
}