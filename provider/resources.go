// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudru

import (
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/sbercloud-terraform/terraform-provider-sbercloud/sbercloud" // Import the upstream provider

	"github.com/sbercloud-terraform/pulumi-cloudru/provider/pkg/version"
)

const (
	mainPkg          = "sbercloud"
	mainMod          = "index"
	dedicatedApigMod = "DedicatedApig" // ?
	sharedApigMod    = "SharedApig"    // ?
	aomMod           = "Aom"
	asMod            = "As"
	cbrMod           = "Cbr"
	cceMod           = "Cce"
	cdmMod           = "Cdm"
	cesMod           = "Ces"
	cfwMod           = "Cfw"
	cssMod           = "Css"
	ctsMod           = "Cts"
	dewMod           = "Dew" // ?
	disMod           = "Dis"
	dliMod           = "Dli"
	drsMod           = "Drs"
	dwsMod           = "Dws"
	dcsMod           = "Dcs"
	dmsMod           = "Dms"
	ddsMod           = "Dds"
	dnsMod           = "Dns"
	ecsMod           = "Ecs"
	eipMod           = "Eip"
	elbMod           = "Elb" // ?
	evsMod           = "Evs"
	epsMod           = "Eps"
	erMod            = "Er"
	functionGraphMod = "FunctionGraph"
	iamMod           = "Iam"
	imsMod           = "Ims"
	ltsMod           = "Lts"
	mrsMod           = "Mrs"
	natMod           = "Nat"
	networkACLMod    = "NetworkAcl"
	obsMod           = "Obs"
	rdsMod           = "Rds"
	sfsMod           = "Sfs"
	vpcepMod         = "Vpcep"
	vpcMod           = "Vpc"
)

//go:embed cmd/pulumi-resource-cloudru/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		//nolint:lll
		P:           shimv2.NewProvider(sbercloud.Provider()),
		Name:        "sbercloud",
		Version:     version.Version,
		DisplayName: "Cloudru",
		Publisher:   "Cloud.ru",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an PNG logo (100x100) for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g. https://github.com/org/pulumi-provider-name/releases/download/v${VERSION}/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing cloud.ru cloud resources.",
		Keywords:          []string{"cloudru", "cloud", "cloud.ru", "cloud-ru", "sbercloud", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://cloud.ru/advanced",
		Repository:        "https://github.com/sbercloud-terraform/pulumi-cloudru",
		GitHubOrg:         "sbercloud-terraform",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		//Config: map[string]*tfbridge.SchemaInfo{
		//	"region": {
		//		Type: "sbercloud:region/region:Region",
		//	},
		//},
		Resources: map[string]*tfbridge.ResourceInfo{
			"sbercloud_aom_service_discovery_rule": {Tok: tfbridge.MakeResource(mainPkg, aomMod, "ServiceDiscoveryRule")},

			"sbercloud_api_gateway_api":   {Tok: tfbridge.MakeResource(mainPkg, sharedApigMod, "Api")},
			"sbercloud_api_gateway_group": {Tok: tfbridge.MakeResource(mainPkg, sharedApigMod, "Group")},

			"sbercloud_apig_api":                         {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Api")},
			"sbercloud_apig_api_publishment":             {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ApiPublishment")},
			"sbercloud_apig_instance":                    {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Instance")},
			"sbercloud_apig_application":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Application")},
			"sbercloud_apig_custom_authorizer":           {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "CustomAuthorizer")},
			"sbercloud_apig_environment":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Environment")},
			"sbercloud_apig_group":                       {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Group")},
			"sbercloud_apig_response":                    {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Response")},
			"sbercloud_apig_throttling_policy_associate": {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ThrottlingPolicyAssociate")},
			"sbercloud_apig_throttling_policy":           {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ThrottlingPolicy")},

			"sbercloud_as_configuration":    {Tok: tfbridge.MakeResource(mainPkg, asMod, "Configuration")},
			"sbercloud_as_group":            {Tok: tfbridge.MakeResource(mainPkg, asMod, "Group")},
			"sbercloud_as_policy":           {Tok: tfbridge.MakeResource(mainPkg, asMod, "Policy")},
			"sbercloud_as_bandwidth_policy": {Tok: tfbridge.MakeResource(mainPkg, asMod, "BandwidthPolicy")},

			"sbercloud_cbr_policy": {Tok: tfbridge.MakeResource(mainPkg, cbrMod, "Policy")},
			"sbercloud_cbr_vault":  {Tok: tfbridge.MakeResource(mainPkg, cbrMod, "Vault")},

			"sbercloud_cce_cluster":     {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Cluster")},
			"sbercloud_cce_node":        {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Node")},
			"sbercloud_cce_node_attach": {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodeAttach")},
			"sbercloud_cce_addon":       {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Addon")},
			"sbercloud_cce_node_pool":   {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodePool")},
			"sbercloud_cce_namespace":   {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Namespace")},
			"sbercloud_cce_pvc":         {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Pvc")},

			"sbercloud_cts_tracker":      {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "Tracker")},
			"sbercloud_cts_data_tracker": {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "DataTracker")},
			"sbercloud_cts_notification": {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "Notification")},

			"sbercloud_cdm_cluster": {Tok: tfbridge.MakeResource(mainPkg, cdmMod, "Cluster")},

			"sbercloud_ces_alarmrule": {Tok: tfbridge.MakeResource(mainPkg, cesMod, "Alarmrule")},

			"sbercloud_compute_instance":         {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Instance")},
			"sbercloud_compute_interface_attach": {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "InterfaceAttach")},
			"sbercloud_compute_keypair":          {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Keypair")},
			"sbercloud_compute_servergroup":      {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Servergroup")},
			"sbercloud_compute_eip_associate":    {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "EipAssociate")},
			"sbercloud_compute_volume_attach":    {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "VolumeAttach")},

			"sbercloud_cfw_acl_rule":             {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AclRule")},
			"sbercloud_cfw_address_group":        {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AddressGroup")},
			"sbercloud_cfw_address_group_member": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AddressGroupMember")},
			"sbercloud_cfw_alarm_config":         {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AlarmConfig")},
			"sbercloud_cfw_anti_virus":           {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AntiVirus")},
			"sbercloud_cfw_black_white_list":     {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "BlackWhiteList")},
			"sbercloud_cfw_capture_task":         {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "CaptureTask")},
			"sbercloud_cfw_dns_resolution":       {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "DnsResolution")},
			"sbercloud_cfw_domain_name_group":    {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "DomainNameGroup")},
			"sbercloud_cfw_eip_protection":       {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "EipProtection")},
			"sbercloud_cfw_firewall":             {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "Firewall")},
			"sbercloud_cfw_ips_rule_mode_change": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "IpsRuleModeChange")},
			"sbercloud_cfw_lts_log":              {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "LtsLog")},
			"sbercloud_cfw_service_group":        {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "ServiceGroup")},
			"sbercloud_cfw_service_group_member": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "ServiceGroupMember")},

			"sbercloud_css_cluster": {Tok: tfbridge.MakeResource(mainPkg, cssMod, "Cluster")},

			"sbercloud_dcs_instance": {Tok: tfbridge.MakeResource(mainPkg, dcsMod, "Instance")},
			"sbercloud_dds_instance": {Tok: tfbridge.MakeResource(mainPkg, ddsMod, "Instance")},

			"sbercloud_dis_stream": {Tok: tfbridge.MakeResource(mainPkg, disMod, "Stream")},

			"sbercloud_dli_database":  {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Database")},
			"sbercloud_dli_package":   {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Package")},
			"sbercloud_dli_queue":     {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Queue")},
			"sbercloud_dli_spark_job": {Tok: tfbridge.MakeResource(mainPkg, dliMod, "SparkJob")},

			"sbercloud_dms_kafka_user":        {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaUser")},
			"sbercloud_dms_kafka_permissions": {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaPermissions")},
			"sbercloud_dms_kafka_instance":    {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaInstance")},
			"sbercloud_dms_kafka_topic":       {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaTopic")},
			"sbercloud_dms_rabbitmq_instance": {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "RabbitmqInstance")},

			"sbercloud_dns_recordset": {Tok: tfbridge.MakeResource(mainPkg, dnsMod, "Recordset")},
			"sbercloud_dns_zone":      {Tok: tfbridge.MakeResource(mainPkg, dnsMod, "Zone")},

			"sbercloud_drs_job":     {Tok: tfbridge.MakeResource(mainPkg, drsMod, "Job")},
			"sbercloud_dws_cluster": {Tok: tfbridge.MakeResource(mainPkg, dwsMod, "Cluster")},

			"sbercloud_enterprise_project": {Tok: tfbridge.MakeResource(mainPkg, epsMod, "Project")},

			"sbercloud_er_association":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "Association")},
			"sbercloud_er_attachment_accepter": {Tok: tfbridge.MakeResource(mainPkg, erMod, "AttachmentAccepter")},
			"sbercloud_er_flow_log":            {Tok: tfbridge.MakeResource(mainPkg, erMod, "FlowLog")},
			"sbercloud_er_instance":            {Tok: tfbridge.MakeResource(mainPkg, erMod, "Instance")},
			"sbercloud_er_propagation":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "Propagation")},
			"sbercloud_er_route_table":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "RouteTable")},
			"sbercloud_er_static_route":        {Tok: tfbridge.MakeResource(mainPkg, erMod, "StaticRoute")},
			"sbercloud_er_vpc_attachment":      {Tok: tfbridge.MakeResource(mainPkg, erMod, "VpcAttachment")},

			"sbercloud_evs_snapshot": {Tok: tfbridge.MakeResource(mainPkg, evsMod, "Snapshot")},
			"sbercloud_evs_volume":   {Tok: tfbridge.MakeResource(mainPkg, evsMod, "Volume")},

			"sbercloud_fgs_dependency": {Tok: tfbridge.MakeResource(mainPkg, functionGraphMod, "Dependency")},
			"sbercloud_fgs_function":   {Tok: tfbridge.MakeResource(mainPkg, functionGraphMod, "Function")},

			"sbercloud_identity_access_key":            {Tok: tfbridge.MakeResource(mainPkg, iamMod, "AccessKey")},
			"sbercloud_identity_acl":                   {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Acl")},
			"sbercloud_identity_agency":                {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Agency")},
			"sbercloud_identity_group":                 {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Group")},
			"sbercloud_identity_group_membership":      {Tok: tfbridge.MakeResource(mainPkg, iamMod, "GroupMembership")},
			"sbercloud_identity_project":               {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Project")},
			"sbercloud_identity_role":                  {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Role")},
			"sbercloud_identity_role_assignment":       {Tok: tfbridge.MakeResource(mainPkg, iamMod, "RoleAssignment")},
			"sbercloud_identity_user":                  {Tok: tfbridge.MakeResource(mainPkg, iamMod, "User")},
			"sbercloud_identity_provider":              {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Provider")},
			"sbercloud_identity_group_role_assignment": {Tok: tfbridge.MakeResource(mainPkg, iamMod, "GroupRoleAssignment")},
			"sbercloud_identity_provider_conversion":   {Tok: tfbridge.MakeResource(mainPkg, iamMod, "ProviderConversion")},

			"sbercloud_images_image": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "Image")},

			"sbercloud_kms_key":     {Tok: tfbridge.MakeResource(mainPkg, dewMod, "Key")},
			"sbercloud_kps_keypair": {Tok: tfbridge.MakeResource(mainPkg, dewMod, "Keypair")},

			"sbercloud_lb_certificate":  {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Certificate")},
			"sbercloud_lb_l7policy":     {Tok: tfbridge.MakeResource(mainPkg, elbMod, "L7policy")},
			"sbercloud_lb_l7rule":       {Tok: tfbridge.MakeResource(mainPkg, elbMod, "L7rule")},
			"sbercloud_lb_listener":     {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Listener")},
			"sbercloud_lb_loadbalancer": {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Loadbalancer")},
			"sbercloud_lb_member":       {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Member")},
			"sbercloud_lb_monitor":      {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Monitor")},
			"sbercloud_lb_pool":         {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Pool")},
			"sbercloud_lb_whitelist":    {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Whitelist")},

			"sbercloud_lts_group":  {Tok: tfbridge.MakeResource(mainPkg, ltsMod, "Group")},
			"sbercloud_lts_stream": {Tok: tfbridge.MakeResource(mainPkg, ltsMod, "Stream")},

			"sbercloud_mapreduce_cluster": {Tok: tfbridge.MakeResource(mainPkg, mrsMod, "cluster")},
			"sbercloud_mapreduce_job":     {Tok: tfbridge.MakeResource(mainPkg, mrsMod, "Job")},

			"sbercloud_nat_dnat_rule": {Tok: tfbridge.MakeResource(mainPkg, natMod, "DnatRule")},
			"sbercloud_nat_gateway":   {Tok: tfbridge.MakeResource(mainPkg, natMod, "Gateway")},
			"sbercloud_nat_snat_rule": {Tok: tfbridge.MakeResource(mainPkg, natMod, "SnatRule")},

			"sbercloud_network_acl":      {Tok: tfbridge.MakeResource(mainPkg, networkACLMod, "Acl")},
			"sbercloud_network_acl_rule": {Tok: tfbridge.MakeResource(mainPkg, networkACLMod, "AclRule")},

			"sbercloud_obs_bucket":        {Tok: tfbridge.MakeResource(mainPkg, obsMod, "Bucket")},
			"sbercloud_obs_bucket_object": {Tok: tfbridge.MakeResource(mainPkg, obsMod, "BucketObject")},
			"sbercloud_obs_bucket_policy": {Tok: tfbridge.MakeResource(mainPkg, obsMod, "BucketPolicy")},

			"sbercloud_rds_instance":                     {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Instance")},
			"sbercloud_rds_parametergroup":               {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Parametergroup")},
			"sbercloud_rds_read_replica_instance":        {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "ReadReplicaInstance")},
			"sbercloud_rds_backup":                       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Backup")},
			"sbercloud_rds_instance_eip_associate":       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "InstanceEipAssociate")},
			"sbercloud_rds_mysql_account":                {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlAccount")},
			"sbercloud_rds_mysql_binlog":                 {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlBinlog")},
			"sbercloud_rds_mysql_database":               {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabase")},
			"sbercloud_rds_mysql_database_privilege":     {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabasePrivilege")},
			"sbercloud_rds_mysql_database_table_restore": {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabaseTableRestore")},
			"sbercloud_rds_pg_account":                   {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgAccount")},
			"sbercloud_rds_pg_account_roles":             {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgAccountRoles")},
			"sbercloud_rds_pg_database":                  {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgDatabase")},
			"sbercloud_rds_pg_hba":                       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgHba")},
			"sbercloud_rds_pg_plugin":                    {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPlugin")},
			"sbercloud_rds_pg_plugin_parameter":          {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPluginParameter")},
			"sbercloud_rds_pg_plugin_update":             {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPluginUpdate")},
			"sbercloud_rds_pg_sql_limit":                 {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgSqlLimit")},
			"sbercloud_rds_sql_audit":                    {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlAudit")},
			"sbercloud_rds_sqlserver_account":            {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverAccount")},
			"sbercloud_rds_sqlserver_database":           {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverDatabase")},
			"sbercloud_rds_sqlserver_database_privilege": {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverDatabasePrivilege")},

			"sbercloud_sfs_access_rule": {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "AccessRule")},
			"sbercloud_sfs_file_system": {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "FileSystem")},
			"sbercloud_sfs_turbo":       {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "Turbo")},

			"sbercloud_vpc_bandwidth": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Bandwidth")},
			"sbercloud_vpc_eip":       {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Eip")},

			"sbercloud_networking_secgroup":      {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Secgroup")},
			"sbercloud_networking_secgroup_rule": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "SecgroupRule")},
			"sbercloud_networking_vip":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Vip")},
			"sbercloud_networking_vip_associate": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "VipAssociate")},

			"sbercloud_vpc_peering_connection":          {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "PeeringConnection")},
			"sbercloud_vpc_peering_connection_accepter": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "PeeringConnectionAccepter")},
			"sbercloud_vpc_route_table":                 {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "RouteTable")},
			"sbercloud_vpc":                             {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Vpc")},
			"sbercloud_vpc_route":                       {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Route")},
			"sbercloud_vpc_subnet":                      {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Subnet")},
			"sbercloud_vpc_address_group":               {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "AddressGroup")},
			"sbercloud_vpcep_endpoint":                  {Tok: tfbridge.MakeResource(mainPkg, vpcepMod, "Endpoint")},
			"sbercloud_vpcep_service":                   {Tok: tfbridge.MakeResource(mainPkg, vpcepMod, "Service")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"sbercloud_apig_environments": {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedApigMod, "getEnvironments")},

			"sbercloud_cbr_vaults": {Tok: tfbridge.MakeDataSource(mainPkg, cbrMod, "getVaults")},

			"sbercloud_cce_addon_template":      {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getAddonTemplate")},
			"sbercloud_cce_cluster":             {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getCluster")},
			"sbercloud_cce_clusters":            {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getClusters")},
			"sbercloud_cce_node":                {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNode")},
			"sbercloud_cce_nodes":               {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNodes")},
			"sbercloud_cce_node_pool":           {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNodePool")},
			"sbercloud_cce_cluster_certificate": {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getClusterCertificates")},

			"sbercloud_cfw_access_control_logs":       {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAccessControlLogs")},
			"sbercloud_cfw_address_group_members":     {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAddressGroupMembers")},
			"sbercloud_cfw_address_groups":            {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAddressGroups")},
			"sbercloud_cfw_attack_logs":               {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAttackLogs")},
			"sbercloud_cfw_black_white_lists":         {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getBlackWhiteLists")},
			"sbercloud_cfw_capture_task_results":      {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getCaptureTaskResults")},
			"sbercloud_cfw_capture_tasks":             {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getCaptureTasks")},
			"sbercloud_cfw_domain_name_groups":        {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getDomainNameGroups")},
			"sbercloud_cfw_domain_name_parse_ip_list": {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getDomainNameParseIpList")},
			"sbercloud_cfw_firewalls":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getFirewalls")},
			"sbercloud_cfw_flow_logs":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getFlowLogs")},
			"sbercloud_cfw_ips_custom_rules":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsCustomRules")},
			"sbercloud_cfw_ips_rule_details":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsRuleDetails")},
			"sbercloud_cfw_ips_rules":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsRules")},
			"sbercloud_cfw_protection_rules":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getProtectionRules")},
			"sbercloud_cfw_regions":                   {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getRegions")},
			"sbercloud_cfw_resource_tags":             {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getResourceTags")},
			"sbercloud_cfw_service_group_members":     {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getServiceGroupMembers")},
			"sbercloud_cfw_service_groups":            {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getServiceGroups")},
			"sbercloud_cfw_tags":                      {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getTags")},

			"sbercloud_compute_flavors":      {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getFlavors")},
			"sbercloud_compute_instance":     {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getInstance")},
			"sbercloud_compute_instances":    {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getInstances")},
			"sbercloud_compute_servergroups": {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getServergroups")},

			"sbercloud_css_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, cssMod, "getFlavors")},

			"sbercloud_dcs_flavors":        {Tok: tfbridge.MakeDataSource(mainPkg, dcsMod, "getFlavors")},
			"sbercloud_dcs_maintainwindow": {Tok: tfbridge.MakeDataSource(mainPkg, dcsMod, "getMaintainwindow")},

			"sbercloud_dds_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, ddsMod, "getFlavors")},

			"sbercloud_dms_kafka_flavors":   {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getFlavors")},
			"sbercloud_dms_kafka_instances": {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getInstances")},
			"sbercloud_dms_product":         {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getProduct")},
			"sbercloud_dms_maintainwindow":  {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getMaintainwindow")},

			"sbercloud_enterprise_project": {Tok: tfbridge.MakeDataSource(mainPkg, epsMod, "getProject")},

			"sbercloud_er_associations":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAssociations")},
			"sbercloud_er_attachments":        {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAttachments")},
			"sbercloud_er_availability_zones": {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAvailabilityZones")},
			"sbercloud_er_available_routes":   {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAvailableRoutes")},
			"sbercloud_er_flow_logs":          {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getFlowLogs")},
			"sbercloud_er_instances":          {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getInstances")},
			"sbercloud_er_propagations":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getPropagations")},
			"sbercloud_er_quotas":             {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getQuotas")},
			"sbercloud_er_resource_tags":      {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getResourceTags")},
			"sbercloud_er_route_tables":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getRouteTables")},
			"sbercloud_er_tags":               {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getTags")},

			"sbercloud_evs_volumes": {Tok: tfbridge.MakeDataSource(mainPkg, evsMod, "getVolumes")},

			"sbercloud_fgs_dependencies": {Tok: tfbridge.MakeDataSource(mainPkg, functionGraphMod, "getDependencies")},

			"sbercloud_identity_role":        {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getRole")},
			"sbercloud_identity_custom_role": {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getCustomRole")},
			"sbercloud_identity_group":       {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getGroup")},
			"sbercloud_identity_projects":    {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getProjects")},
			"sbercloud_identity_users":       {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getUsers")},

			"sbercloud_images_image":  {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getImage")},
			"sbercloud_images_images": {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getImages")},

			"sbercloud_kms_key":      {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getKey")},
			"sbercloud_kms_data_key": {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getDataKey")},
			"sbercloud_kps_keypairs": {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getKeypairs")},

			"sbercloud_lb_listeners":    {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getListeners")},
			"sbercloud_lb_loadbalancer": {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getLoadbalancer")},
			"sbercloud_lb_certificate":  {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getCertificate")},
			"sbercloud_lb_pools":        {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getPools")},

			"sbercloud_nat_gateway": {Tok: tfbridge.MakeDataSource(mainPkg, natMod, "getGateway")},

			"sbercloud_obs_buckets":       {Tok: tfbridge.MakeDataSource(mainPkg, obsMod, "getBuckets")},
			"sbercloud_obs_bucket_object": {Tok: tfbridge.MakeDataSource(mainPkg, obsMod, "getBucketObject")},

			"sbercloud_rds_flavors":                         {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getFlavors")},
			"sbercloud_rds_engine_versions":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getEngineVersions")},
			"sbercloud_rds_instances":                       {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getInstances")},
			"sbercloud_rds_backups":                         {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getBackups")},
			"sbercloud_rds_pg_accounts":                     {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgAccounts")},
			"sbercloud_rds_pg_databases":                    {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgDatabases")},
			"sbercloud_rds_pg_plugin_parameter_value_range": {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPluginParameterValueRange")},
			"sbercloud_rds_pg_plugin_parameter_values":      {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPluginParameterValues")},
			"sbercloud_rds_pg_plugins":                      {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPlugins")},
			"sbercloud_rds_pg_roles":                        {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgRoles")},
			"sbercloud_rds_pg_sql_limits":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgSqlLimits")},
			"sbercloud_rds_storage_types":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getStorageTypes")},

			"sbercloud_sfs_file_system": {Tok: tfbridge.MakeDataSource(mainPkg, sfsMod, "getFileSystem")},
			"sbercloud_sfs_turbos":      {Tok: tfbridge.MakeDataSource(mainPkg, sfsMod, "getTurbos")},

			"sbercloud_vpc_bandwidth": {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getBandwidth")},
			"sbercloud_vpc_eip":       {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getEip")},
			"sbercloud_vpc_eips":      {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getEips")},

			"sbercloud_networking_port":      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getPort")},
			"sbercloud_networking_secgroup":  {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSecgroup")},
			"sbercloud_networking_secgroups": {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSecgroups")},

			"sbercloud_vpc":                    {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getVpc")},
			"sbercloud_vpcs":                   {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getVpcs")},
			"sbercloud_vpc_ids":                {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getIds")},
			"sbercloud_vpc_peering_connection": {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getPeeringConnection")},
			"sbercloud_vpc_route_table":        {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getRouteTable")},
			"sbercloud_vpc_subnet":             {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnet")},
			"sbercloud_vpc_subnets":            {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnets")},
			"sbercloud_vpc_subnet_ids":         {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnetIds")},
			"sbercloud_vpc_address_groups":     {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getAddressGroups")},
			"sbercloud_vpc_routes":             {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getRoutes")},

			"sbercloud_vpcep_public_services": {Tok: tfbridge.MakeDataSource(mainPkg, vpcepMod, "getPublicServices")},

			"sbercloud_dws_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, dwsMod, "getFlaovrs")},
		},
		//ExtraTypes: map[string]schema.ComplexTypeSpec{
		//	"sbercloud:region/region:Region": {
		//		ObjectTypeSpec: schema.ObjectTypeSpec{
		//			Type: "string",
		//		},
		//	},
		//},
		//
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "pulumi-cloudru",
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			PackageName:          "pulumi_cloudru",
			RespectSchemaVersion: true,
			PyProject:            struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				"github.com/sbercloud-terraform/pulumi-cloudru/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"cloudru",
				//mainPkg,
			),
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("sbercloud_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
