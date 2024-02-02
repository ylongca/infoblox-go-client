// Code generated by "infoblox-go-client-generator"; DO NOT EDIT.

package ibclient_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

// This test checks if common methods for each
// generated Infoblox object works as expected
var _ = DescribeTable("Common methods of WAPI Objects",
	func(obj ibclient.IBObject, expectedReturnFields []string, expectedObjectType string) {
		// Check if default return fields are valid
		Expect(obj.ReturnFields()).To(Equal(expectedReturnFields))
		// Check if ObjectType is valid
		Expect(obj.ObjectType()).To(Equal(expectedObjectType))
	},
	Entry(
		"AdAuthService",
		&ibclient.AdAuthService{},
		[]string{"name"},
		"ad_auth_service",
	),
	Entry(
		"Admingroup",
		&ibclient.Admingroup{},
		[]string{"comment", "name"},
		"admingroup",
	),
	Entry(
		"Adminrole",
		&ibclient.Adminrole{},
		[]string{"comment", "name"},
		"adminrole",
	),
	Entry(
		"Adminuser",
		&ibclient.Adminuser{},
		[]string{"admin_groups", "comment", "name"},
		"adminuser",
	),
	Entry(
		"Allendpoints",
		&ibclient.Allendpoints{},
		[]string{},
		"allendpoints",
	),
	Entry(
		"Allnsgroup",
		&ibclient.Allnsgroup{},
		[]string{"name", "type"},
		"allnsgroup",
	),
	Entry(
		"Allrecords",
		&ibclient.Allrecords{},
		[]string{"comment", "name", "type", "view", "zone"},
		"allrecords",
	),
	Entry(
		"Allrpzrecords",
		&ibclient.Allrpzrecords{},
		[]string{"comment", "name", "type", "view", "zone"},
		"allrpzrecords",
	),
	Entry(
		"Approvalworkflow",
		&ibclient.Approvalworkflow{},
		[]string{"approval_group", "submitter_group"},
		"approvalworkflow",
	),
	Entry(
		"Authpolicy",
		&ibclient.Authpolicy{},
		[]string{"default_group", "usage_type"},
		"authpolicy",
	),
	Entry(
		"Awsrte53taskgroup",
		&ibclient.Awsrte53taskgroup{},
		[]string{"account_id", "comment", "disabled", "name", "sync_status"},
		"awsrte53taskgroup",
	),
	Entry(
		"Awsuser",
		&ibclient.Awsuser{},
		[]string{"access_key_id", "account_id", "name"},
		"awsuser",
	),
	Entry(
		"Bfdtemplate",
		&ibclient.Bfdtemplate{},
		[]string{"name"},
		"bfdtemplate",
	),
	Entry(
		"Bulkhost",
		&ibclient.Bulkhost{},
		[]string{"comment", "prefix"},
		"bulkhost",
	),
	Entry(
		"Bulkhostnametemplate",
		&ibclient.Bulkhostnametemplate{},
		[]string{"is_grid_default", "template_format", "template_name"},
		"bulkhostnametemplate",
	),
	Entry(
		"Cacertificate",
		&ibclient.Cacertificate{},
		[]string{"distinguished_name", "issuer", "serial", "used_by", "valid_not_after", "valid_not_before"},
		"cacertificate",
	),
	Entry(
		"CapacityReport",
		&ibclient.CapacityReport{},
		[]string{"name", "percent_used", "role"},
		"capacityreport",
	),
	Entry(
		"Captiveportal",
		&ibclient.Captiveportal{},
		[]string{"name"},
		"captiveportal",
	),
	Entry(
		"CertificateAuthservice",
		&ibclient.CertificateAuthservice{},
		[]string{"name"},
		"certificate:authservice",
	),
	Entry(
		"CiscoiseEndpoint",
		&ibclient.CiscoiseEndpoint{},
		[]string{"address", "disable", "resolved_address", "type", "version"},
		"ciscoise:endpoint",
	),
	Entry(
		"Csvimporttask",
		&ibclient.Csvimporttask{},
		[]string{"action", "admin_name", "end_time", "file_name", "file_size", "import_id", "lines_failed", "lines_processed", "lines_warning", "on_error", "operation", "separator", "start_time", "status", "update_method"},
		"csvimporttask",
	),
	Entry(
		"DbObjects",
		&ibclient.DbObjects{},
		[]string{"last_sequence_id", "object", "object_type", "unique_id"},
		"db_objects",
	),
	Entry(
		"Dbsnapshot",
		&ibclient.Dbsnapshot{},
		[]string{"comment", "timestamp"},
		"dbsnapshot",
	),
	Entry(
		"DdnsPrincipalcluster",
		&ibclient.DdnsPrincipalcluster{},
		[]string{"comment", "group", "name", "principals"},
		"ddns:principalcluster",
	),
	Entry(
		"DdnsPrincipalclusterGroup",
		&ibclient.DdnsPrincipalclusterGroup{},
		[]string{"comment", "name"},
		"ddns:principalcluster:group",
	),
	Entry(
		"DeletedObjects",
		&ibclient.DeletedObjects{},
		[]string{"object_type"},
		"deleted_objects",
	),
	Entry(
		"DhcpStatistics",
		&ibclient.DhcpStatistics{},
		[]string{"dhcp_utilization", "dhcp_utilization_status", "dynamic_hosts", "static_hosts", "total_hosts"},
		"dhcp:statistics",
	),
	Entry(
		"Dhcpfailover",
		&ibclient.Dhcpfailover{},
		[]string{"name"},
		"dhcpfailover",
	),
	Entry(
		"Dhcpoptiondefinition",
		&ibclient.Dhcpoptiondefinition{},
		[]string{"code", "name", "type"},
		"dhcpoptiondefinition",
	),
	Entry(
		"Dhcpoptionspace",
		&ibclient.Dhcpoptionspace{},
		[]string{"comment", "name"},
		"dhcpoptionspace",
	),
	Entry(
		"Discovery",
		&ibclient.Discovery{},
		[]string{},
		"discovery",
	),
	Entry(
		"DiscoveryCredentialgroup",
		&ibclient.DiscoveryCredentialgroup{},
		[]string{"name"},
		"discovery:credentialgroup",
	),
	Entry(
		"DiscoveryDevice",
		&ibclient.DiscoveryDevice{},
		[]string{"address", "name", "network_view"},
		"discovery:device",
	),
	Entry(
		"DiscoveryDevicecomponent",
		&ibclient.DiscoveryDevicecomponent{},
		[]string{"component_name", "description", "model", "serial", "type"},
		"discovery:devicecomponent",
	),
	Entry(
		"DiscoveryDeviceinterface",
		&ibclient.DiscoveryDeviceinterface{},
		[]string{"name", "type"},
		"discovery:deviceinterface",
	),
	Entry(
		"DiscoveryDeviceneighbor",
		&ibclient.DiscoveryDeviceneighbor{},
		[]string{"address", "address_ref", "mac", "name"},
		"discovery:deviceneighbor",
	),
	Entry(
		"DiscoveryDevicesupportbundle",
		&ibclient.DiscoveryDevicesupportbundle{},
		[]string{"author", "integrated_ind", "name", "version"},
		"discovery:devicesupportbundle",
	),
	Entry(
		"DiscoveryDiagnostictask",
		&ibclient.DiscoveryDiagnostictask{},
		[]string{"ip_address", "network_view", "task_id"},
		"discovery:diagnostictask",
	),
	Entry(
		"DiscoveryGridproperties",
		&ibclient.DiscoveryGridproperties{},
		[]string{"grid_name"},
		"discovery:gridproperties",
	),
	Entry(
		"DiscoveryMemberproperties",
		&ibclient.DiscoveryMemberproperties{},
		[]string{"discovery_member"},
		"discovery:memberproperties",
	),
	Entry(
		"DiscoverySdnnetwork",
		&ibclient.DiscoverySdnnetwork{},
		[]string{"name", "network_view", "source_sdn_config"},
		"discovery:sdnnetwork",
	),
	Entry(
		"DiscoveryStatus",
		&ibclient.DiscoveryStatus{},
		[]string{"address", "name", "network_view", "status"},
		"discovery:status",
	),
	Entry(
		"DiscoveryVrf",
		&ibclient.DiscoveryVrf{},
		[]string{"device", "name", "network_view", "route_distinguisher"},
		"discovery:vrf",
	),
	Entry(
		"Discoverytask",
		&ibclient.Discoverytask{},
		[]string{"discovery_task_oid", "member_name"},
		"discoverytask",
	),
	Entry(
		"Distributionschedule",
		&ibclient.Distributionschedule{},
		[]string{"active", "start_time", "time_zone"},
		"distributionschedule",
	),
	Entry(
		"Dns64group",
		&ibclient.Dns64group{},
		[]string{"comment", "disable", "name"},
		"dns64group",
	),
	Entry(
		"Dtc",
		&ibclient.Dtc{},
		[]string{},
		"dtc",
	),
	Entry(
		"DtcAllrecords",
		&ibclient.DtcAllrecords{},
		[]string{"comment", "dtc_server", "type"},
		"dtc:allrecords",
	),
	Entry(
		"DtcCertificate",
		&ibclient.DtcCertificate{},
		[]string{},
		"dtc:certificate",
	),
	Entry(
		"DtcLbdn",
		&ibclient.DtcLbdn{},
		[]string{"comment", "name"},
		"dtc:lbdn",
	),
	Entry(
		"DtcMonitor",
		&ibclient.DtcMonitor{},
		[]string{"comment", "name", "type"},
		"dtc:monitor",
	),
	Entry(
		"DtcMonitorHttp",
		&ibclient.DtcMonitorHttp{},
		[]string{"comment", "name"},
		"dtc:monitor:http",
	),
	Entry(
		"DtcMonitorIcmp",
		&ibclient.DtcMonitorIcmp{},
		[]string{"comment", "name"},
		"dtc:monitor:icmp",
	),
	Entry(
		"DtcMonitorPdp",
		&ibclient.DtcMonitorPdp{},
		[]string{"comment", "name"},
		"dtc:monitor:pdp",
	),
	Entry(
		"DtcMonitorSip",
		&ibclient.DtcMonitorSip{},
		[]string{"comment", "name"},
		"dtc:monitor:sip",
	),
	Entry(
		"DtcMonitorSnmp",
		&ibclient.DtcMonitorSnmp{},
		[]string{"comment", "name"},
		"dtc:monitor:snmp",
	),
	Entry(
		"DtcMonitorTcp",
		&ibclient.DtcMonitorTcp{},
		[]string{"comment", "name"},
		"dtc:monitor:tcp",
	),
	Entry(
		"DtcObject",
		&ibclient.DtcObject{},
		[]string{"abstract_type", "comment", "display_type", "name", "status"},
		"dtc:object",
	),
	Entry(
		"DtcPool",
		&ibclient.DtcPool{},
		[]string{"comment", "name"},
		"dtc:pool",
	),
	Entry(
		"DtcRecordA",
		&ibclient.DtcRecordA{},
		[]string{"dtc_server", "ipv4addr"},
		"dtc:record:a",
	),
	Entry(
		"DtcRecordAaaa",
		&ibclient.DtcRecordAaaa{},
		[]string{"dtc_server", "ipv6addr"},
		"dtc:record:aaaa",
	),
	Entry(
		"DtcRecordCname",
		&ibclient.DtcRecordCname{},
		[]string{"canonical", "dtc_server"},
		"dtc:record:cname",
	),
	Entry(
		"DtcRecordNaptr",
		&ibclient.DtcRecordNaptr{},
		[]string{"dtc_server", "order", "preference", "regexp", "replacement", "services"},
		"dtc:record:naptr",
	),
	Entry(
		"DtcRecordSrv",
		&ibclient.DtcRecordSrv{},
		[]string{"dtc_server", "name", "port", "priority", "target", "weight"},
		"dtc:record:srv",
	),
	Entry(
		"DtcServer",
		&ibclient.DtcServer{},
		[]string{"comment", "host", "name"},
		"dtc:server",
	),
	Entry(
		"DtcTopology",
		&ibclient.DtcTopology{},
		[]string{"comment", "name"},
		"dtc:topology",
	),
	Entry(
		"DtcTopologyLabel",
		&ibclient.DtcTopologyLabel{},
		[]string{"field", "label"},
		"dtc:topology:label",
	),
	Entry(
		"DtcTopologyRule",
		&ibclient.DtcTopologyRule{},
		[]string{},
		"dtc:topology:rule",
	),
	Entry(
		"DxlEndpoint",
		&ibclient.DxlEndpoint{},
		[]string{"disable", "name", "outbound_member_type"},
		"dxl:endpoint",
	),
	Entry(
		"EADefinition",
		&ibclient.EADefinition{},
		[]string{"comment", "default_value", "name", "type"},
		"extensibleattributedef",
	),
	Entry(
		"Fileop",
		&ibclient.Fileop{},
		[]string{},
		"fileop",
	),
	Entry(
		"Filterfingerprint",
		&ibclient.Filterfingerprint{},
		[]string{"comment", "name"},
		"filterfingerprint",
	),
	Entry(
		"Filtermac",
		&ibclient.Filtermac{},
		[]string{"comment", "name"},
		"filtermac",
	),
	Entry(
		"Filternac",
		&ibclient.Filternac{},
		[]string{"comment", "name"},
		"filternac",
	),
	Entry(
		"Filteroption",
		&ibclient.Filteroption{},
		[]string{"comment", "name"},
		"filteroption",
	),
	Entry(
		"Filterrelayagent",
		&ibclient.Filterrelayagent{},
		[]string{"comment", "name"},
		"filterrelayagent",
	),
	Entry(
		"Fingerprint",
		&ibclient.Fingerprint{},
		[]string{"comment", "device_class", "name"},
		"fingerprint",
	),
	Entry(
		"Ipv4FixedAddress",
		&ibclient.Ipv4FixedAddress{},
		[]string{"ipv4addr", "network_view"},
		"fixedaddress",
	),
	Entry(
		"Fixedaddresstemplate",
		&ibclient.Fixedaddresstemplate{},
		[]string{"comment", "name"},
		"fixedaddresstemplate",
	),
	Entry(
		"Ftpuser",
		&ibclient.Ftpuser{},
		[]string{"username"},
		"ftpuser",
	),
	Entry(
		"Grid",
		&ibclient.Grid{},
		[]string{},
		"grid",
	),
	Entry(
		"GridCloudapi",
		&ibclient.GridCloudapi{},
		[]string{"allow_api_admins", "allowed_api_admins", "enable_recycle_bin"},
		"grid:cloudapi",
	),
	Entry(
		"GridCloudapiCloudstatistics",
		&ibclient.GridCloudapiCloudstatistics{},
		[]string{"allocated_available_ratio", "allocated_ip_count", "available_ip_count", "fixed_ip_count", "floating_ip_count", "tenant_count", "tenant_ip_count", "tenant_vm_count"},
		"grid:cloudapi:cloudstatistics",
	),
	Entry(
		"GridCloudapiTenant",
		&ibclient.GridCloudapiTenant{},
		[]string{"comment", "id", "name"},
		"grid:cloudapi:tenant",
	),
	Entry(
		"GridCloudapiVm",
		&ibclient.GridCloudapiVm{},
		[]string{"comment", "id", "name"},
		"grid:cloudapi:vm",
	),
	Entry(
		"GridCloudapiVmaddress",
		&ibclient.GridCloudapiVmaddress{},
		[]string{"address", "is_ipv4", "network_view", "port_id", "vm_name"},
		"grid:cloudapi:vmaddress",
	),
	Entry(
		"GridDashboard",
		&ibclient.GridDashboard{},
		[]string{"analytics_tunneling_event_critical_threshold", "analytics_tunneling_event_warning_threshold", "atp_critical_event_critical_threshold", "atp_critical_event_warning_threshold", "atp_major_event_critical_threshold", "atp_major_event_warning_threshold", "atp_warning_event_critical_threshold", "atp_warning_event_warning_threshold", "rpz_blocked_hit_critical_threshold", "rpz_blocked_hit_warning_threshold", "rpz_passthru_event_critical_threshold", "rpz_passthru_event_warning_threshold", "rpz_substituted_hit_critical_threshold", "rpz_substituted_hit_warning_threshold"},
		"grid:dashboard",
	),
	Entry(
		"GridDhcpproperties",
		&ibclient.GridDhcpproperties{},
		[]string{"disable_all_nac_filters", "grid"},
		"grid:dhcpproperties",
	),
	Entry(
		"GridDns",
		&ibclient.GridDns{},
		[]string{},
		"grid:dns",
	),
	Entry(
		"GridFiledistribution",
		&ibclient.GridFiledistribution{},
		[]string{"allow_uploads", "current_usage", "global_status", "name", "storage_limit"},
		"grid:filedistribution",
	),
	Entry(
		"GridLicensePool",
		&ibclient.GridLicensePool{},
		[]string{"type"},
		"grid:license_pool",
	),
	Entry(
		"GridLicensePoolContainer",
		&ibclient.GridLicensePoolContainer{},
		[]string{},
		"grid:license_pool_container",
	),
	Entry(
		"GridMaxminddbinfo",
		&ibclient.GridMaxminddbinfo{},
		[]string{"binary_major_version", "binary_minor_version", "build_time", "database_type", "deployment_time", "member", "topology_type"},
		"grid:maxminddbinfo",
	),
	Entry(
		"GridMemberCloudapi",
		&ibclient.GridMemberCloudapi{},
		[]string{"allow_api_admins", "allowed_api_admins", "enable_service", "member", "status"},
		"grid:member:cloudapi",
	),
	Entry(
		"GridServicerestartGroup",
		&ibclient.GridServicerestartGroup{},
		[]string{"comment", "name", "service"},
		"grid:servicerestart:group",
	),
	Entry(
		"GridServicerestartGroupOrder",
		&ibclient.GridServicerestartGroupOrder{},
		[]string{},
		"grid:servicerestart:group:order",
	),
	Entry(
		"GridServicerestartRequest",
		&ibclient.GridServicerestartRequest{},
		[]string{"error", "group", "result", "state"},
		"grid:servicerestart:request",
	),
	Entry(
		"GridServicerestartRequestChangedobject",
		&ibclient.GridServicerestartRequestChangedobject{},
		[]string{"action", "changed_properties", "changed_time", "object_name", "object_type", "user_name"},
		"grid:servicerestart:request:changedobject",
	),
	Entry(
		"GridServicerestartStatus",
		&ibclient.GridServicerestartStatus{},
		[]string{"failures", "finished", "grouped", "needed_restart", "no_restart", "parent", "pending", "pending_restart", "processing", "restarting", "success", "timeouts"},
		"grid:servicerestart:status",
	),
	Entry(
		"GridThreatanalytics",
		&ibclient.GridThreatanalytics{},
		[]string{"enable_auto_download", "enable_scheduled_download", "module_update_policy", "name"},
		"grid:threatanalytics",
	),
	Entry(
		"GridThreatprotection",
		&ibclient.GridThreatprotection{},
		[]string{"grid_name"},
		"grid:threatprotection",
	),
	Entry(
		"GridX509certificate",
		&ibclient.GridX509certificate{},
		[]string{"issuer", "serial", "subject"},
		"grid:x509certificate",
	),
	Entry(
		"Hostnamerewritepolicy",
		&ibclient.Hostnamerewritepolicy{},
		[]string{"name", "replacement_character", "valid_characters"},
		"hostnamerewritepolicy",
	),
	Entry(
		"HsmAllgroups",
		&ibclient.HsmAllgroups{},
		[]string{"groups"},
		"hsm:allgroups",
	),
	Entry(
		"HsmSafenetgroup",
		&ibclient.HsmSafenetgroup{},
		[]string{"comment", "hsm_version", "name"},
		"hsm:safenetgroup",
	),
	Entry(
		"HsmThalesgroup",
		&ibclient.HsmThalesgroup{},
		[]string{"comment", "key_server_ip", "name"},
		"hsm:thalesgroup",
	),
	Entry(
		"IpamStatistics",
		&ibclient.IpamStatistics{},
		[]string{"cidr", "network", "network_view"},
		"ipam:statistics",
	),
	Entry(
		"IPv4Address",
		&ibclient.IPv4Address{},
		[]string{"dhcp_client_identifier", "ip_address", "is_conflict", "lease_state", "mac_address", "names", "network", "network_view", "objects", "status", "types", "usage", "username"},
		"ipv4address",
	),
	Entry(
		"IPv6Address",
		&ibclient.IPv6Address{},
		[]string{"duid", "ip_address", "is_conflict", "lease_state", "names", "network", "network_view", "objects", "status", "types", "usage"},
		"ipv6address",
	),
	Entry(
		"Ipv6dhcpoptiondefinition",
		&ibclient.Ipv6dhcpoptiondefinition{},
		[]string{"code", "name", "type"},
		"ipv6dhcpoptiondefinition",
	),
	Entry(
		"Ipv6dhcpoptionspace",
		&ibclient.Ipv6dhcpoptionspace{},
		[]string{"comment", "enterprise_number", "name"},
		"ipv6dhcpoptionspace",
	),
	Entry(
		"Ipv6filteroption",
		&ibclient.Ipv6filteroption{},
		[]string{"comment", "name"},
		"ipv6filteroption",
	),
	Entry(
		"Ipv6FixedAddress",
		&ibclient.Ipv6FixedAddress{},
		[]string{"duid", "ipv6addr", "network_view"},
		"ipv6fixedaddress",
	),
	Entry(
		"Ipv6fixedaddresstemplate",
		&ibclient.Ipv6fixedaddresstemplate{},
		[]string{"comment", "name"},
		"ipv6fixedaddresstemplate",
	),
	Entry(
		"Ipv6Network",
		&ibclient.Ipv6Network{},
		[]string{"comment", "network", "network_view"},
		"ipv6network",
	),
	Entry(
		"Ipv6NetworkContainer",
		&ibclient.Ipv6NetworkContainer{},
		[]string{"comment", "network", "network_view"},
		"ipv6networkcontainer",
	),
	Entry(
		"IPv6NetworkTemplate",
		&ibclient.IPv6NetworkTemplate{},
		[]string{"comment", "name"},
		"ipv6networktemplate",
	),
	Entry(
		"IPv6Range",
		&ibclient.IPv6Range{},
		[]string{"comment", "end_addr", "network", "network_view", "start_addr"},
		"ipv6range",
	),
	Entry(
		"Ipv6rangetemplate",
		&ibclient.Ipv6rangetemplate{},
		[]string{"comment", "name", "number_of_addresses", "offset"},
		"ipv6rangetemplate",
	),
	Entry(
		"IPv6SharedNetwork",
		&ibclient.IPv6SharedNetwork{},
		[]string{"comment", "name", "network_view", "networks"},
		"ipv6sharednetwork",
	),
	Entry(
		"Kerberoskey",
		&ibclient.Kerberoskey{},
		[]string{"domain", "enctype", "in_use", "principal", "version"},
		"kerberoskey",
	),
	Entry(
		"LdapAuthService",
		&ibclient.LdapAuthService{},
		[]string{"comment", "disable", "ldap_user_attribute", "mode", "name"},
		"ldap_auth_service",
	),
	Entry(
		"Lease",
		&ibclient.Lease{},
		[]string{"address", "network_view"},
		"lease",
	),
	Entry(
		"LicenseGridwide",
		&ibclient.LicenseGridwide{},
		[]string{"type"},
		"license:gridwide",
	),
	Entry(
		"LocaluserAuthservice",
		&ibclient.LocaluserAuthservice{},
		[]string{"comment", "disabled", "name"},
		"localuser:authservice",
	),
	Entry(
		"MACFilterAddress",
		&ibclient.MACFilterAddress{},
		[]string{"authentication_time", "comment", "expiration_time", "filter", "guest_custom_field1", "guest_custom_field2", "guest_custom_field3", "guest_custom_field4", "guest_email", "guest_first_name", "guest_last_name", "guest_middle_name", "guest_phone", "is_registered_user", "mac", "never_expires", "reserved_for_infoblox", "username"},
		"macfilteraddress",
	),
	Entry(
		"Mastergrid",
		&ibclient.Mastergrid{},
		[]string{"address", "enable", "port"},
		"mastergrid",
	),
	Entry(
		"Member",
		&ibclient.Member{},
		[]string{"config_addr_type", "host_name", "platform", "service_type_configuration"},
		"member",
	),
	Entry(
		"MemberDHCPProperties",
		&ibclient.MemberDHCPProperties{},
		[]string{"host_name", "ipv4addr", "ipv6addr"},
		"member:dhcpproperties",
	),
	Entry(
		"MemberDns",
		&ibclient.MemberDns{},
		[]string{"host_name", "ipv4addr", "ipv6addr"},
		"member:dns",
	),
	Entry(
		"MemberFiledistribution",
		&ibclient.MemberFiledistribution{},
		[]string{"host_name", "ipv4_address", "ipv6_address", "status"},
		"member:filedistribution",
	),
	Entry(
		"MemberLicense",
		&ibclient.MemberLicense{},
		[]string{"type"},
		"member:license",
	),
	Entry(
		"MemberParentalcontrol",
		&ibclient.MemberParentalcontrol{},
		[]string{"enable_service", "name"},
		"member:parentalcontrol",
	),
	Entry(
		"MemberThreatanalytics",
		&ibclient.MemberThreatanalytics{},
		[]string{"host_name", "ipv4_address", "ipv6_address", "status"},
		"member:threatanalytics",
	),
	Entry(
		"MemberThreatprotection",
		&ibclient.MemberThreatprotection{},
		[]string{},
		"member:threatprotection",
	),
	Entry(
		"Memberdfp",
		&ibclient.Memberdfp{},
		[]string{},
		"memberdfp",
	),
	Entry(
		"Msserver",
		&ibclient.Msserver{},
		[]string{"address"},
		"msserver",
	),
	Entry(
		"MsserverAdsitesDomain",
		&ibclient.MsserverAdsitesDomain{},
		[]string{"name", "netbios", "network_view"},
		"msserver:adsites:domain",
	),
	Entry(
		"MsserverAdsitesSite",
		&ibclient.MsserverAdsitesSite{},
		[]string{"domain", "name"},
		"msserver:adsites:site",
	),
	Entry(
		"MsserverDhcp",
		&ibclient.MsserverDhcp{},
		[]string{"address"},
		"msserver:dhcp",
	),
	Entry(
		"MsserverDns",
		&ibclient.MsserverDns{},
		[]string{"address"},
		"msserver:dns",
	),
	Entry(
		"Mssuperscope",
		&ibclient.Mssuperscope{},
		[]string{"disable", "name", "network_view"},
		"mssuperscope",
	),
	Entry(
		"Namedacl",
		&ibclient.Namedacl{},
		[]string{"comment", "name"},
		"namedacl",
	),
	Entry(
		"Natgroup",
		&ibclient.Natgroup{},
		[]string{"comment", "name"},
		"natgroup",
	),
	Entry(
		"Ipv4Network",
		&ibclient.Ipv4Network{},
		[]string{"comment", "network", "network_view"},
		"network",
	),
	Entry(
		"NetworkDiscovery",
		&ibclient.NetworkDiscovery{},
		[]string{},
		"network_discovery",
	),
	Entry(
		"Ipv4NetworkContainer",
		&ibclient.Ipv4NetworkContainer{},
		[]string{"comment", "network", "network_view"},
		"networkcontainer",
	),
	Entry(
		"NetworkTemplate",
		&ibclient.NetworkTemplate{},
		[]string{"comment", "name"},
		"networktemplate",
	),
	Entry(
		"Networkuser",
		&ibclient.Networkuser{},
		[]string{"address", "domainname", "name", "network_view", "user_status"},
		"networkuser",
	),
	Entry(
		"NetworkView",
		&ibclient.NetworkView{},
		[]string{"comment", "is_default", "name"},
		"networkview",
	),
	Entry(
		"NotificationRestEndpoint",
		&ibclient.NotificationRestEndpoint{},
		[]string{"name", "outbound_member_type", "uri"},
		"notification:rest:endpoint",
	),
	Entry(
		"NotificationRestTemplate",
		&ibclient.NotificationRestTemplate{},
		[]string{"content", "name"},
		"notification:rest:template",
	),
	Entry(
		"NotificationRule",
		&ibclient.NotificationRule{},
		[]string{"event_type", "name", "notification_action", "notification_target"},
		"notification:rule",
	),
	Entry(
		"Nsgroup",
		&ibclient.Nsgroup{},
		[]string{"comment", "name"},
		"nsgroup",
	),
	Entry(
		"NsgroupDelegation",
		&ibclient.NsgroupDelegation{},
		[]string{"delegate_to", "name"},
		"nsgroup:delegation",
	),
	Entry(
		"NsgroupForwardingmember",
		&ibclient.NsgroupForwardingmember{},
		[]string{"forwarding_servers", "name"},
		"nsgroup:forwardingmember",
	),
	Entry(
		"NsgroupForwardstubserver",
		&ibclient.NsgroupForwardstubserver{},
		[]string{"external_servers", "name"},
		"nsgroup:forwardstubserver",
	),
	Entry(
		"NsgroupStubmember",
		&ibclient.NsgroupStubmember{},
		[]string{"name"},
		"nsgroup:stubmember",
	),
	Entry(
		"Orderedranges",
		&ibclient.Orderedranges{},
		[]string{"network", "ranges"},
		"orderedranges",
	),
	Entry(
		"Orderedresponsepolicyzones",
		&ibclient.Orderedresponsepolicyzones{},
		[]string{"view"},
		"orderedresponsepolicyzones",
	),
	Entry(
		"OutboundCloudclient",
		&ibclient.OutboundCloudclient{},
		[]string{"enable", "interval"},
		"outbound:cloudclient",
	),
	Entry(
		"ParentalcontrolAvp",
		&ibclient.ParentalcontrolAvp{},
		[]string{"name", "type", "value_type"},
		"parentalcontrol:avp",
	),
	Entry(
		"ParentalcontrolBlockingpolicy",
		&ibclient.ParentalcontrolBlockingpolicy{},
		[]string{"name", "value"},
		"parentalcontrol:blockingpolicy",
	),
	Entry(
		"ParentalcontrolSubscriber",
		&ibclient.ParentalcontrolSubscriber{},
		[]string{"alt_subscriber_id", "local_id", "subscriber_id"},
		"parentalcontrol:subscriber",
	),
	Entry(
		"ParentalcontrolSubscriberrecord",
		&ibclient.ParentalcontrolSubscriberrecord{},
		[]string{"accounting_session_id", "ip_addr", "ipsd", "localid", "prefix", "site", "subscriber_id"},
		"parentalcontrol:subscriberrecord",
	),
	Entry(
		"ParentalcontrolSubscribersite",
		&ibclient.ParentalcontrolSubscribersite{},
		[]string{"block_size", "dca_sub_bw_list", "dca_sub_query_count", "first_port", "name", "stop_anycast", "strict_nat"},
		"parentalcontrol:subscribersite",
	),
	Entry(
		"Permission",
		&ibclient.Permission{},
		[]string{"group", "permission", "resource_type", "role"},
		"permission",
	),
	Entry(
		"PxgridEndpoint",
		&ibclient.PxgridEndpoint{},
		[]string{"address", "disable", "name", "outbound_member_type"},
		"pxgrid:endpoint",
	),
	Entry(
		"RadiusAuthservice",
		&ibclient.RadiusAuthservice{},
		[]string{"comment", "disable", "name"},
		"radius:authservice",
	),
	Entry(
		"Range",
		&ibclient.Range{},
		[]string{"comment", "end_addr", "network", "network_view", "start_addr"},
		"range",
	),
	Entry(
		"Rangetemplate",
		&ibclient.Rangetemplate{},
		[]string{"comment", "name", "number_of_addresses", "offset"},
		"rangetemplate",
	),
	Entry(
		"RecordA",
		&ibclient.RecordA{},
		[]string{"ipv4addr", "name", "view"},
		"record:a",
	),
	Entry(
		"RecordAAAA",
		&ibclient.RecordAAAA{},
		[]string{"ipv6addr", "name", "view"},
		"record:aaaa",
	),
	Entry(
		"RecordAlias",
		&ibclient.RecordAlias{},
		[]string{"name", "target_name", "target_type", "view"},
		"record:alias",
	),
	Entry(
		"RecordCaa",
		&ibclient.RecordCaa{},
		[]string{"name", "view"},
		"record:caa",
	),
	Entry(
		"RecordCNAME",
		&ibclient.RecordCNAME{},
		[]string{"canonical", "name", "view"},
		"record:cname",
	),
	Entry(
		"RecordDhcid",
		&ibclient.RecordDhcid{},
		[]string{"name", "view"},
		"record:dhcid",
	),
	Entry(
		"RecordDname",
		&ibclient.RecordDname{},
		[]string{"name", "target", "view"},
		"record:dname",
	),
	Entry(
		"RecordDnskey",
		&ibclient.RecordDnskey{},
		[]string{"name", "view"},
		"record:dnskey",
	),
	Entry(
		"RecordDs",
		&ibclient.RecordDs{},
		[]string{"name", "view"},
		"record:ds",
	),
	Entry(
		"RecordDtclbdn",
		&ibclient.RecordDtclbdn{},
		[]string{"comment", "name", "view", "zone"},
		"record:dtclbdn",
	),
	Entry(
		"HostRecord",
		&ibclient.HostRecord{},
		[]string{"ipv4addrs", "ipv6addrs", "name", "view"},
		"record:host",
	),
	Entry(
		"HostRecordIpv4Addr",
		&ibclient.HostRecordIpv4Addr{},
		[]string{"configure_for_dhcp", "host", "ipv4addr", "mac"},
		"record:host_ipv4addr",
	),
	Entry(
		"HostRecordIpv6Addr",
		&ibclient.HostRecordIpv6Addr{},
		[]string{"configure_for_dhcp", "duid", "host", "ipv6addr"},
		"record:host_ipv6addr",
	),
	Entry(
		"RecordMX",
		&ibclient.RecordMX{},
		[]string{"mail_exchanger", "name", "preference", "view"},
		"record:mx",
	),
	Entry(
		"RecordNaptr",
		&ibclient.RecordNaptr{},
		[]string{"name", "order", "preference", "regexp", "replacement", "services", "view"},
		"record:naptr",
	),
	Entry(
		"RecordNS",
		&ibclient.RecordNS{},
		[]string{"name", "nameserver", "view"},
		"record:ns",
	),
	Entry(
		"RecordNsec",
		&ibclient.RecordNsec{},
		[]string{"name", "view"},
		"record:nsec",
	),
	Entry(
		"RecordNsec3",
		&ibclient.RecordNsec3{},
		[]string{"name", "view"},
		"record:nsec3",
	),
	Entry(
		"RecordNsec3param",
		&ibclient.RecordNsec3param{},
		[]string{"name", "view"},
		"record:nsec3param",
	),
	Entry(
		"RecordPTR",
		&ibclient.RecordPTR{},
		[]string{"ptrdname", "view"},
		"record:ptr",
	),
	Entry(
		"RecordRpzA",
		&ibclient.RecordRpzA{},
		[]string{"ipv4addr", "name", "view"},
		"record:rpz:a",
	),
	Entry(
		"RecordRpzAIpaddress",
		&ibclient.RecordRpzAIpaddress{},
		[]string{"ipv4addr", "name", "view"},
		"record:rpz:a:ipaddress",
	),
	Entry(
		"RecordRpzAaaa",
		&ibclient.RecordRpzAaaa{},
		[]string{"ipv6addr", "name", "view"},
		"record:rpz:aaaa",
	),
	Entry(
		"RecordRpzAaaaIpaddress",
		&ibclient.RecordRpzAaaaIpaddress{},
		[]string{"ipv6addr", "name", "view"},
		"record:rpz:aaaa:ipaddress",
	),
	Entry(
		"RecordRpzCname",
		&ibclient.RecordRpzCname{},
		[]string{"canonical", "name", "view"},
		"record:rpz:cname",
	),
	Entry(
		"RecordRpzCnameClientipaddress",
		&ibclient.RecordRpzCnameClientipaddress{},
		[]string{"canonical", "name", "view"},
		"record:rpz:cname:clientipaddress",
	),
	Entry(
		"RecordRpzCnameClientipaddressdn",
		&ibclient.RecordRpzCnameClientipaddressdn{},
		[]string{"canonical", "name", "view"},
		"record:rpz:cname:clientipaddressdn",
	),
	Entry(
		"RecordRpzCnameIpaddress",
		&ibclient.RecordRpzCnameIpaddress{},
		[]string{"canonical", "name", "view"},
		"record:rpz:cname:ipaddress",
	),
	Entry(
		"RecordRpzCnameIpaddressdn",
		&ibclient.RecordRpzCnameIpaddressdn{},
		[]string{"canonical", "name", "view"},
		"record:rpz:cname:ipaddressdn",
	),
	Entry(
		"RecordRpzMx",
		&ibclient.RecordRpzMx{},
		[]string{"mail_exchanger", "name", "preference", "view"},
		"record:rpz:mx",
	),
	Entry(
		"RecordRpzNaptr",
		&ibclient.RecordRpzNaptr{},
		[]string{"name", "order", "preference", "regexp", "replacement", "services", "view"},
		"record:rpz:naptr",
	),
	Entry(
		"RecordRpzPtr",
		&ibclient.RecordRpzPtr{},
		[]string{"ptrdname", "view"},
		"record:rpz:ptr",
	),
	Entry(
		"RecordRpzSrv",
		&ibclient.RecordRpzSrv{},
		[]string{"name", "port", "priority", "target", "view", "weight"},
		"record:rpz:srv",
	),
	Entry(
		"RecordRpzTxt",
		&ibclient.RecordRpzTxt{},
		[]string{"name", "text", "view"},
		"record:rpz:txt",
	),
	Entry(
		"RecordRrsig",
		&ibclient.RecordRrsig{},
		[]string{"name", "view"},
		"record:rrsig",
	),
	Entry(
		"RecordSRV",
		&ibclient.RecordSRV{},
		[]string{"name", "port", "priority", "target", "view", "weight"},
		"record:srv",
	),
	Entry(
		"RecordTlsa",
		&ibclient.RecordTlsa{},
		[]string{"name", "view"},
		"record:tlsa",
	),
	Entry(
		"RecordTXT",
		&ibclient.RecordTXT{},
		[]string{"name", "text", "view"},
		"record:txt",
	),
	Entry(
		"RecordUnknown",
		&ibclient.RecordUnknown{},
		[]string{"name", "view"},
		"record:unknown",
	),
	Entry(
		"Recordnamepolicy",
		&ibclient.Recordnamepolicy{},
		[]string{"is_default", "name", "regex"},
		"recordnamepolicy",
	),
	Entry(
		"Restartservicestatus",
		&ibclient.Restartservicestatus{},
		[]string{"dhcp_status", "dns_status", "member", "reporting_status"},
		"restartservicestatus",
	),
	Entry(
		"Rir",
		&ibclient.Rir{},
		[]string{"communication_mode", "email", "name", "url"},
		"rir",
	),
	Entry(
		"RirOrganization",
		&ibclient.RirOrganization{},
		[]string{"id", "maintainer", "name", "rir", "sender_email"},
		"rir:organization",
	),
	Entry(
		"RoamingHost",
		&ibclient.RoamingHost{},
		[]string{"address_type", "name", "network_view"},
		"roaminghost",
	),
	Entry(
		"Ruleset",
		&ibclient.Ruleset{},
		[]string{"comment", "disabled", "name", "type"},
		"ruleset",
	),
	Entry(
		"SamlAuthservice",
		&ibclient.SamlAuthservice{},
		[]string{"name"},
		"saml:authservice",
	),
	Entry(
		"Scavengingtask",
		&ibclient.Scavengingtask{},
		[]string{"action", "associated_object", "status"},
		"scavengingtask",
	),
	Entry(
		"ScheduledTask",
		&ibclient.ScheduledTask{},
		[]string{"approval_status", "execution_status", "task_id"},
		"scheduledtask",
	),
	Entry(
		"Search",
		&ibclient.Search{},
		[]string{},
		"search",
	),
	Entry(
		"SharedNetwork",
		&ibclient.SharedNetwork{},
		[]string{"comment", "name", "network_view", "networks"},
		"sharednetwork",
	),
	Entry(
		"SharedRecordA",
		&ibclient.SharedRecordA{},
		[]string{"ipv4addr", "name", "shared_record_group"},
		"sharedrecord:a",
	),
	Entry(
		"SharedRecordAAAA",
		&ibclient.SharedRecordAAAA{},
		[]string{"ipv6addr", "name", "shared_record_group"},
		"sharedrecord:aaaa",
	),
	Entry(
		"SharedrecordCname",
		&ibclient.SharedrecordCname{},
		[]string{"canonical", "name", "shared_record_group"},
		"sharedrecord:cname",
	),
	Entry(
		"SharedRecordMX",
		&ibclient.SharedRecordMX{},
		[]string{"mail_exchanger", "name", "preference", "shared_record_group"},
		"sharedrecord:mx",
	),
	Entry(
		"SharedrecordSrv",
		&ibclient.SharedrecordSrv{},
		[]string{"name", "port", "priority", "shared_record_group", "target", "weight"},
		"sharedrecord:srv",
	),
	Entry(
		"SharedRecordTXT",
		&ibclient.SharedRecordTXT{},
		[]string{"name", "shared_record_group", "text"},
		"sharedrecord:txt",
	),
	Entry(
		"Sharedrecordgroup",
		&ibclient.Sharedrecordgroup{},
		[]string{"comment", "name"},
		"sharedrecordgroup",
	),
	Entry(
		"SmartfolderChildren",
		&ibclient.SmartfolderChildren{},
		[]string{"resource", "value", "value_type"},
		"smartfolder:children",
	),
	Entry(
		"SmartfolderGlobal",
		&ibclient.SmartfolderGlobal{},
		[]string{"comment", "name"},
		"smartfolder:global",
	),
	Entry(
		"SmartfolderPersonal",
		&ibclient.SmartfolderPersonal{},
		[]string{"comment", "is_shortcut", "name"},
		"smartfolder:personal",
	),
	Entry(
		"SNMPUser",
		&ibclient.SNMPUser{},
		[]string{"comment", "name"},
		"snmpuser",
	),
	Entry(
		"Superhost",
		&ibclient.Superhost{},
		[]string{"comment", "name"},
		"superhost",
	),
	Entry(
		"Superhostchild",
		&ibclient.Superhostchild{},
		[]string{"comment", "data", "name", "network_view", "parent", "record_parent", "type", "view"},
		"superhostchild",
	),
	Entry(
		"SyslogEndpoint",
		&ibclient.SyslogEndpoint{},
		[]string{"name", "outbound_member_type"},
		"syslog:endpoint",
	),
	Entry(
		"TacacsplusAuthservice",
		&ibclient.TacacsplusAuthservice{},
		[]string{"comment", "disable", "name"},
		"tacacsplus:authservice",
	),
	Entry(
		"Taxii",
		&ibclient.Taxii{},
		[]string{"ipv4addr", "ipv6addr", "name"},
		"taxii",
	),
	Entry(
		"Tftpfiledir",
		&ibclient.Tftpfiledir{},
		[]string{"directory", "name", "type"},
		"tftpfiledir",
	),
	Entry(
		"ThreatanalyticsAnalyticsWhitelist",
		&ibclient.ThreatanalyticsAnalyticsWhitelist{},
		[]string{"version"},
		"threatanalytics:analytics_whitelist",
	),
	Entry(
		"ThreatanalyticsModuleset",
		&ibclient.ThreatanalyticsModuleset{},
		[]string{"version"},
		"threatanalytics:moduleset",
	),
	Entry(
		"ThreatanalyticsWhitelist",
		&ibclient.ThreatanalyticsWhitelist{},
		[]string{"comment", "disable", "fqdn"},
		"threatanalytics:whitelist",
	),
	Entry(
		"ThreatinsightCloudclient",
		&ibclient.ThreatinsightCloudclient{},
		[]string{"enable", "interval"},
		"threatinsight:cloudclient",
	),
	Entry(
		"ThreatprotectionGridRule",
		&ibclient.ThreatprotectionGridRule{},
		[]string{"name", "ruleset", "sid"},
		"threatprotection:grid:rule",
	),
	Entry(
		"ThreatprotectionProfile",
		&ibclient.ThreatprotectionProfile{},
		[]string{"comment", "name"},
		"threatprotection:profile",
	),
	Entry(
		"ThreatprotectionProfileRule",
		&ibclient.ThreatprotectionProfileRule{},
		[]string{"profile", "rule"},
		"threatprotection:profile:rule",
	),
	Entry(
		"ThreatprotectionRule",
		&ibclient.ThreatprotectionRule{},
		[]string{"member", "rule"},
		"threatprotection:rule",
	),
	Entry(
		"ThreatprotectionRulecategory",
		&ibclient.ThreatprotectionRulecategory{},
		[]string{"name", "ruleset"},
		"threatprotection:rulecategory",
	),
	Entry(
		"ThreatprotectionRuleset",
		&ibclient.ThreatprotectionRuleset{},
		[]string{"add_type", "version"},
		"threatprotection:ruleset",
	),
	Entry(
		"ThreatprotectionRuletemplate",
		&ibclient.ThreatprotectionRuletemplate{},
		[]string{"name", "ruleset", "sid"},
		"threatprotection:ruletemplate",
	),
	Entry(
		"ThreatprotectionStatistics",
		&ibclient.ThreatprotectionStatistics{},
		[]string{"member", "stat_infos"},
		"threatprotection:statistics",
	),
	Entry(
		"Upgradegroup",
		&ibclient.Upgradegroup{},
		[]string{"comment", "name"},
		"upgradegroup",
	),
	Entry(
		"Upgradeschedule",
		&ibclient.Upgradeschedule{},
		[]string{"active", "start_time", "time_zone"},
		"upgradeschedule",
	),
	Entry(
		"UpgradeStatus",
		&ibclient.UpgradeStatus{},
		[]string{"alternate_version", "comment", "current_version", "distribution_version", "element_status", "grid_state", "group_state", "ha_status", "hotfixes", "ipv4_address", "ipv6_address", "member", "message", "pnode_role", "reverted", "status_value", "status_value_update_time", "steps", "steps_completed", "steps_total", "type", "upgrade_group", "upgrade_state", "upgrade_test_status", "upload_version"},
		"upgradestatus",
	),
	Entry(
		"UserProfile",
		&ibclient.UserProfile{},
		[]string{"name"},
		"userprofile",
	),
	Entry(
		"Vdiscoverytask",
		&ibclient.Vdiscoverytask{},
		[]string{"name", "state"},
		"vdiscoverytask",
	),
	Entry(
		"View",
		&ibclient.View{},
		[]string{"comment", "is_default", "name"},
		"view",
	),
	Entry(
		"Vlan",
		&ibclient.Vlan{},
		[]string{"id", "name", "parent"},
		"vlan",
	),
	Entry(
		"Vlanrange",
		&ibclient.Vlanrange{},
		[]string{"end_vlan_id", "name", "start_vlan_id", "vlan_view"},
		"vlanrange",
	),
	Entry(
		"Vlanview",
		&ibclient.Vlanview{},
		[]string{"end_vlan_id", "name", "start_vlan_id"},
		"vlanview",
	),
	Entry(
		"ZoneAuth",
		&ibclient.ZoneAuth{},
		[]string{"fqdn", "view"},
		"zone_auth",
	),
	Entry(
		"ZoneAuthDiscrepancy",
		&ibclient.ZoneAuthDiscrepancy{},
		[]string{"description", "severity", "timestamp", "zone"},
		"zone_auth_discrepancy",
	),
	Entry(
		"ZoneDelegated",
		&ibclient.ZoneDelegated{},
		[]string{"delegate_to", "fqdn", "view"},
		"zone_delegated",
	),
	Entry(
		"ZoneForward",
		&ibclient.ZoneForward{},
		[]string{"forward_to", "fqdn", "view"},
		"zone_forward",
	),
	Entry(
		"ZoneRp",
		&ibclient.ZoneRp{},
		[]string{"fqdn", "view"},
		"zone_rp",
	),
	Entry(
		"ZoneStub",
		&ibclient.ZoneStub{},
		[]string{"fqdn", "stub_from", "view"},
		"zone_stub",
	),
)

var _ = Describe("WAPI_VERSION metadata", func() {
	It("Should be equal to 2.12.1", func() {
		Expect(ibclient.WAPI_VERSION).To(Equal("2.12.1"))
	})
})