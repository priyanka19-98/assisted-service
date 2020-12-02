// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterCreateParams cluster create params
//
// swagger:model cluster-create-params
type ClusterCreateParams struct {

	// NTP source going to be added to all the hosts.
	AdditionalNtpSource *string `json:"additional_ntp_source,omitempty"`

	// Base domain of the cluster. All DNS records must be sub-domains of this base and include the cluster name.
	BaseDNSDomain string `json:"base_dns_domain,omitempty"`

	// IP address block from which Pod IPs are allocated. This block must not overlap with existing physical networks. These IP addresses are used for the Pod network, and if you need to access the Pods from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$
	ClusterNetworkCidr *string `json:"cluster_network_cidr,omitempty"`

	// The subnet prefix length to assign to each individual node. For example, if clusterNetworkHostPrefix is set to 23, then each node is assigned a /23 subnet out of the given cidr (clusterNetworkCIDR), which allows for 510 (2^(32 - 23) - 2) pod IPs addresses. If you are required to provide access to nodes from an external network, configure load balancers and routers to manage the traffic.
	// Maximum: 128
	// Minimum: 1
	ClusterNetworkHostPrefix int64 `json:"cluster_network_host_prefix,omitempty"`

	// A proxy URL to use for creating HTTP connections outside the cluster.
	// http://\<username\>:\<pswd\>@\<ip\>:\<port\>
	//
	HTTPProxy *string `json:"http_proxy,omitempty"`

	// A proxy URL to use for creating HTTPS connections outside the cluster.
	// http://\<username\>:\<pswd\>@\<ip\>:\<port\>
	//
	HTTPSProxy *string `json:"https_proxy,omitempty"`

	// The virtual IP used for cluster ingress traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}))$
	IngressVip string `json:"ingress_vip,omitempty"`

	// Name of the OpenShift cluster.
	// Required: true
	// Max Length: 54
	// Min Length: 1
	Name *string `json:"name"`

	// An "*" or a comma-separated list of destination domain names, domains, IP addresses, or other network CIDRs to exclude from proxying.
	NoProxy *string `json:"no_proxy,omitempty"`

	// Version of the OpenShift cluster.
	// Required: true
	// Enum: [4.5 4.6 4.7]
	OpenshiftVersion *string `json:"openshift_version"`

	// operators
	Operators *Operators `json:"operators,omitempty"`

	// The pull secret obtained from Red Hat OpenShift Cluster Manager at cloud.redhat.com/openshift/install/pull-secret.
	// Required: true
	PullSecret *string `json:"pull_secret"`

	// The IP address pool to use for service IP addresses. You can enter only one IP address pool. If you need to access the services from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$
	ServiceNetworkCidr *string `json:"service_network_cidr,omitempty"`

	// SSH public key for debugging OpenShift nodes.
	SSHPublicKey string `json:"ssh_public_key,omitempty"`

	// Indicate if the networking is managed by the user.
	UserManagedNetworking *bool `json:"user-managed-networking,omitempty"`

	// Indicate if virtual IP DHCP allocation mode is enabled.
	VipDhcpAllocation *bool `json:"vip_dhcp_allocation,omitempty"`
}

// Validate validates this cluster create params
func (m *ClusterCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetworkHostPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressVip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenshiftVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreateParams) validateClusterNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("cluster_network_cidr", "body", string(*m.ClusterNetworkCidr), `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateClusterNetworkHostPrefix(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetworkHostPrefix) { // not required
		return nil
	}

	if err := validate.MinimumInt("cluster_network_host_prefix", "body", int64(m.ClusterNetworkHostPrefix), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cluster_network_host_prefix", "body", int64(m.ClusterNetworkHostPrefix), 128, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateIngressVip(formats strfmt.Registry) error {

	if swag.IsZero(m.IngressVip) { // not required
		return nil
	}

	if err := validate.Pattern("ingress_vip", "body", string(m.IngressVip), `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3})|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,}))$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 54); err != nil {
		return err
	}

	return nil
}

var clusterCreateParamsTypeOpenshiftVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["4.5","4.6","4.7"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterCreateParamsTypeOpenshiftVersionPropEnum = append(clusterCreateParamsTypeOpenshiftVersionPropEnum, v)
	}
}

const (

	// ClusterCreateParamsOpenshiftVersionNr45 captures enum value "4.5"
	ClusterCreateParamsOpenshiftVersionNr45 string = "4.5"

	// ClusterCreateParamsOpenshiftVersionNr46 captures enum value "4.6"
	ClusterCreateParamsOpenshiftVersionNr46 string = "4.6"

	// ClusterCreateParamsOpenshiftVersionNr47 captures enum value "4.7"
	ClusterCreateParamsOpenshiftVersionNr47 string = "4.7"
)

// prop value enum
func (m *ClusterCreateParams) validateOpenshiftVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterCreateParamsTypeOpenshiftVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterCreateParams) validateOpenshiftVersion(formats strfmt.Registry) error {

	if err := validate.Required("openshift_version", "body", m.OpenshiftVersion); err != nil {
		return err
	}

	// value enum
	if err := m.validateOpenshiftVersionEnum("openshift_version", "body", *m.OpenshiftVersion); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateOperators(formats strfmt.Registry) error {

	if swag.IsZero(m.Operators) { // not required
		return nil
	}

	if m.Operators != nil {
		if err := m.Operators.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operators")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreateParams) validatePullSecret(formats strfmt.Registry) error {

	if err := validate.Required("pull_secret", "body", m.PullSecret); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateParams) validateServiceNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("service_network_cidr", "body", string(*m.ServiceNetworkCidr), `^(?:(?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\/(?:(?:[0-9])|(?:[1-2][0-9])|(?:3[0-2])))|(?:(?:[0-9a-fA-F]*:[0-9a-fA-F]*){2,})/(?:(?:[0-9])|(?:[1-9][0-9])|(?:1[0-1][0-9])|(?:12[0-8])))$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreateParams) UnmarshalBinary(b []byte) error {
	var res ClusterCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
