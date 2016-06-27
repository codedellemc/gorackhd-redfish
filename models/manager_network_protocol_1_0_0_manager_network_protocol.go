package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*ManagerNetworkProtocol100ManagerNetworkProtocol This resource is used to obtain or modify the network services managed by a given manager.

swagger:model ManagerNetworkProtocol.1.0.0_ManagerNetworkProtocol
*/
type ManagerNetworkProtocol100ManagerNetworkProtocol struct {

	/* at odata context
	 */
	AtOdataContext Odata400Context `json:"@odata.context,omitempty"`

	/* at odata id
	 */
	AtOdataID Odata400ID `json:"@odata.id,omitempty"`

	/* at odata type
	 */
	AtOdataType Odata400Type `json:"@odata.type,omitempty"`

	/* description
	 */
	Description ResourceDescription `json:"Description,omitempty"`

	/* This is the fully qualified domain name for the manager obtained by DNS including the host name and top-level domain name.

	Read Only: true
	*/
	FQDN string `json:"FQDN,omitempty"`

	/* Settings for this Manager's HTTP protocol support

	Read Only: true
	*/
	HTTP *ManagerNetworkProtocol100Protocol `json:"HTTP,omitempty"`

	/* Settings for this Manager's HTTPS protocol support

	Read Only: true
	*/
	HTTPS *ManagerNetworkProtocol100Protocol `json:"HTTPS,omitempty"`

	/* The DNS Host Name of this manager, without any domain information

	Read Only: true
	*/
	HostName string `json:"HostName,omitempty"`

	/* Settings for this Manager's IPMI-over-LAN protocol support

	Read Only: true
	*/
	IPMI *ManagerNetworkProtocol100Protocol `json:"IPMI,omitempty"`

	/* Id
	 */
	ID ResourceID `json:"Id,omitempty"`

	/* Settings for this Manager's KVM-IP protocol support

	Read Only: true
	*/
	KVMIP *ManagerNetworkProtocol100Protocol `json:"KVMIP,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* Settings for this Manager's SNMP support

	Read Only: true
	*/
	SNMP *ManagerNetworkProtocol100Protocol `json:"SNMP,omitempty"`

	/* Settings for this Manager's SSDP support

	Read Only: true
	*/
	SSDP *ManagerNetworkProtocol100SSDProtocol `json:"SSDP,omitempty"`

	/* Settings for this Manager's SSH (Secure Shell) protocol support

	Read Only: true
	*/
	SSH *ManagerNetworkProtocol100Protocol `json:"SSH,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`

	/* Settings for this Manager's Telnet protocol support

	Read Only: true
	*/
	Telnet *ManagerNetworkProtocol100Protocol `json:"Telnet,omitempty"`

	/* Settings for this Manager's Virtual Media support

	Read Only: true
	*/
	VirtualMedia *ManagerNetworkProtocol100Protocol `json:"VirtualMedia,omitempty"`
}

// Validate validates this manager network protocol 1 0 0 manager network protocol
func (m *ManagerNetworkProtocol100ManagerNetworkProtocol) Validate(formats strfmt.Registry) error {
	return nil
}