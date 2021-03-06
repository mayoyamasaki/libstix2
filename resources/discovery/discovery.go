// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package discovery

import (
	"github.com/freetaxii/libstix2/resources/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
Discovery - This type implements the TAXII 2 Discovery Resource and defines
all of the properties and methods needed to create and work with the TAXII Discovery
Resource. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the TAXII 2 specification documents.

This Endpoint provides general information about a TAXII Server, including the
advertised API Roots. It's a common entry point for TAXII Clients into the data
and services provided by a TAXII Server. For example, clients auto-discovering
TAXII Servers via the DNS SRV record defined in section 1.4.1 will be able to
automatically retrieve a discovery response for that server by requesting the
/taxii/ path on that domain.

Discovery API responses MAY advertise any TAXII API Root that they have
permission to advertise, included those hosted on other servers.

The discovery resource contains information about a TAXII Server, such as a
human-readable title, description, and contact information, as well as a list of
API Roots that it is advertising. It also has an indication of which API Root it
considers the default, or the one to use in the absence of other
information/user choice.
*/
type Discovery struct {
	properties.TitleProperty
	properties.DescriptionProperty
	Contact  string   `json:"contact,omitempty"`
	Default  string   `json:"default,omitempty"`
	APIRoots []string `json:"api_roots,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new TAXII Discovery object and return it as a
pointer.
*/
func New() *Discovery {
	var obj Discovery
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Discovery
// ----------------------------------------------------------------------

/*
SetContact - This methods takes in a string value representing contact
information and updates the contact property.
*/
func (r *Discovery) SetContact(s string) error {
	r.Contact = s
	return nil
}

/*
GetContact - This method returns the contact information from the contact
property.
*/
func (r *Discovery) GetContact() string {
	return r.Contact
}

/*
SetDefault - This methods takes in a string value representing a default
api-root and updates the default property.
*/
func (r *Discovery) SetDefault(s string) error {
	r.Default = s
	return nil
}

/*
GetDefault - This methods returns the default api-root.
*/
func (r *Discovery) GetDefault() string {
	return r.Default
}

/*
AddAPIRoot - This method takes in a string value that represents an api-root
and adds it to the list in the APIRoots property.
*/
func (r *Discovery) AddAPIRoot(s string) error {
	if r.APIRoots == nil {
		a := make([]string, 0)
		r.APIRoots = a
	}
	r.APIRoots = append(r.APIRoots, s)
	return nil
}
