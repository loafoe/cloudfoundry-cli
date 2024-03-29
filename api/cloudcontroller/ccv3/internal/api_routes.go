package internal

import "net/http"

// Naming convention:
//
// HTTP method + non-parameter parts of the path + "Request"
//
// If the request returns a single entity by GUID, use the singular (for example
// /v2/organizations/:organization_guid is GetOrganization).
const (
	DeleteApplicationProcessInstanceRequest                     = "DeleteApplicationProcessInstance"
	DeleteApplicationRequest                                    = "DeleteApplication"
	DeleteSharedOrgFromDomainRequest                            = "DeleteSharedOrgFromDomain"
	DeleteBuildpackRequest                                      = "DeleteBuildpack"
	DeleteDomainRequest                                         = "DeleteDomainRequest"
	DeleteIsolationSegmentRelationshipOrganizationRequest       = "DeleteIsolationSegmentRelationshipOrganization"
	DeleteIsolationSegmentRequest                               = "DeleteIsolationSegment"
	DeleteServiceInstanceRelationshipsSharedSpaceRequest        = "DeleteServiceInstanceRelationshipsSharedSpace"
	GetApplicationDropletCurrentRequest                         = "GetApplicationDropletCurrent"
	GetApplicationEnvRequest                                    = "GetApplicationEnv"
	GetApplicationManifestRequest                               = "GetApplicationManifest"
	GetApplicationProcessesRequest                              = "GetApplicationProcesses"
	GetApplicationProcessRequest                                = "GetApplicationProcess"
	GetApplicationsRequest                                      = "GetApplications"
	GetApplicationTasksRequest                                  = "GetApplicationTasks"
	GetBuildpacksRequest                                        = "GetBuildpacks"
	GetBuildRequest                                             = "GetBuild"
	GetDeploymentRequest                                        = "GetDeployment"
	GetDeploymentsRequest                                       = "GetDeployments"
	GetDomainsRequest                                           = "GetDomains"
	GetDropletRequest                                           = "GetDroplet"
	GetDropletsRequest                                          = "GetDroplets"
	GetFeatureFlagRequest                                       = "GetFeatureFlag"
	GetFeatureFlagsRequest                                      = "GetFeatureFlags"
	GetIsolationSegmentOrganizationsRequest                     = "GetIsolationSegmentOrganizations"
	GetIsolationSegmentRequest                                  = "GetIsolationSegment"
	GetIsolationSegmentsRequest                                 = "GetIsolationSegments"
	GetOrganizationDomainsRequest                               = "GetOrganizationDomains"
	GetOrganizationRelationshipDefaultIsolationSegmentRequest   = "GetOrganizationRelationshipDefaultIsolationSegment"
	GetOrganizationsRequest                                     = "GetOrganizations"
	GetPackageRequest                                           = "GetPackage"
	GetPackagesRequest                                          = "GetPackages"
	GetProcessStatsRequest                                      = "GetProcessStats"
	GetServiceInstancesRequest                                  = "GetServiceInstances"
	GetSpaceRelationshipIsolationSegmentRequest                 = "GetSpaceRelationshipIsolationSegment"
	GetSpacesRequest                                            = "GetSpaces"
	GetStacksRequest                                            = "GetStacks"
	PatchApplicationCurrentDropletRequest                       = "PatchApplicationCurrentDroplet"
	PatchApplicationEnvironmentVariablesRequest                 = "PatchApplicationEnvironmentVariables"
	PatchApplicationRequest                                     = "PatchApplication"
	PatchBuildpackRequest                                       = "PatchBuildpack"
	PatchFeatureFlagRequest                                     = "PatchFeatureFlag"
	PatchOrganizationRequest                                    = "PatchOrganization"
	PatchOrganizationRelationshipDefaultIsolationSegmentRequest = "PatchOrganizationRelationshipDefaultIsolationSegment"
	PatchProcessRequest                                         = "PatchProcess"
	PatchSpaceRelationshipIsolationSegmentRequest               = "PatchSpaceRelationshipIsolationSegment"
	PostApplicationActionApplyManifest                          = "PostApplicationActionApplyM"
	PostApplicationActionRestartRequest                         = "PostApplicationActionRestart"
	PostApplicationActionStartRequest                           = "PostApplicationActionStart"
	PostApplicationActionStopRequest                            = "PostApplicationActionStop"
	PostApplicationDeploymentActionCancelRequest                = "PostApplicationDeploymentActionCancel"
	PostApplicationDeploymentRequest                            = "PostApplicationDeployment"
	PostApplicationProcessActionScaleRequest                    = "PostApplicationProcessActionScale"
	PostApplicationRequest                                      = "PostApplication"
	PostApplicationTasksRequest                                 = "PostApplicationTasks"
	PostBuildRequest                                            = "PostBuild"
	PostBuildpackBitsRequest                                    = "PostBuildpackBits"
	PostBuildpackRequest                                        = "PostBuildpack"
	PostDomainRequest                                           = "PostDomain"
	PostIsolationSegmentRelationshipOrganizationsRequest        = "PostIsolationSegmentRelationshipOrganizations"
	PostIsolationSegmentsRequest                                = "PostIsolationSegments"
	PostPackageRequest                                          = "PostPackage"
	PostResourceMatchesRequest                                  = "PostResourceMatches"
	PostServiceInstanceRelationshipsSharedSpacesRequest         = "PostServiceInstanceRelationshipsSharedSpaces"
	PostSpaceActionApplyManifestRequest                         = "PostSpaceActionApplyManifest"
	PostSpaceDiffManifestRequest                                = "PostSpaceDiffManifest"
	PutTaskCancelRequest                                        = "PutTaskCancel"
	SharePrivateDomainRequest                                   = "SharePrivateDomainRequest"

	GetOrganizationRequest    = "GetOrganization"
	DeleteOrganizationRequest = "DeleteOrganization"
	PostOrganizationRequest   = "PostOrganization"
	GetDefaultDomainRequest   = "GetDefaultDomain"

	GetServiceOfferingsRequest   = "GetServiceOfferings"
	DeleteServiceOfferingRequest = "DeleteServiceOffering"

	GetServicePlansRequest = "GetServicePlans"

	PostRouteRequest   = "PostRoute"
	DeleteRouteRequest = "DeleteRouteRequest"
	PatchRouteRequest  = "PatchRoute"

	DeleteOrphanedRoutesRequest = "DeleteOrphanedRoutes"
	GetApplicationRoutesRequest = "GetApplicationRoutes"
	GetRouteDestinationsRequest = "GetRouteDestinations"
	GetRoutesRequest            = "GetRoutes"
	MapRouteRequest             = "MapRoute"
	UnmapRouteRequest           = "UnmapRoute"

	// v3 droplet
	PostDropletRequest        = "PostDroplet"
	GetPackageDropletsRequest = "GetPackageDroplets"
	PostDropletBitsRequest    = "PostDropletBits"
	GetDropletBitsRequest     = "GetDropletBits"

	// v3 package
	PostPackageBitsRequest = "PostPackageBits"

	// v3 service credential binding
	PostServiceCredentialBindingRequest       = "PostServiceCredentialBinding"
	GetServiceCredentialBindingsRequest       = "GetServiceCredentialBindings"
	DeleteServiceCredentialBindingRequest     = "DeleteServiceCredentialBinding"
	GetServiceCredentialBindingDetailsRequest = "GetServiceCredentialBindingDetails"

	// service_instance
	GetServiceInstanceParametersRequest                = "GetServiceInstanceParameters"
	GetServiceInstanceCredentialsRequest               = "GetServiceInstanceCredentails"
	PostServiceInstanceRequest                         = "PostServiceInstance"
	PatchServiceInstanceRequest                        = "PatchServiceInstance"
	DeleteServiceInstanceRequest                       = "DeleteServiceInstance"
	GetServiceInstanceRelationshipsSharedSpacesRequest = "GetServiceInstanceRelationshipSharedSpacesRequest"
	GetServiceInstanceSharedSpacesUsageSummaryRequest  = "GetServiceInstanceSharedSpacesUsageSummaryRequest"

	// v3 process add missing endpoints
	GetProcessRequest   = "GetProcess"
	GetProcessesRequest = "GetProcesses"

	// v3 application feature
	GetApplicationFeaturesRequest   = "GetApplicationFeatures"
	GetSSHEnabled                   = "GetSSHEnabled"
	PatchApplicationFeaturesRequest = "PatchApplicationFeatures"

	// v3 environment variable
	GetEnvironmentVariableGroupRequest   = "GetEnvironmentVariableGroup"
	PatchEnvironmentVariableGroupRequest = "PatchEnvironmentVariableGroup"

	// v3 domain add missing endpoints
	GetDomainRouteReservationsRequest = "GetDomainRouteReservations"
	GetDomainRequest                  = "GetDomain"

	// v3 space add missing
	PostSpaceRequest   = "PostSpace"
	DeleteSpaceRequest = "DeleteSpace"
	PatchSpaceRequest  = "PatchSpace"

	// v3 revision
	GetApplicationRevisionsRequest         = "GetApplicationRevisions"
	GetApplicationRevisionsDeployedRequest = "GetApplicationRevisionsDeployed"

	// v3 organization quota
	PostOrganizationQuotaApplyRequest = "PostOrganizationQuotaApply"
	PostOrganizationQuotaRequest      = "PostOrganizationQuota"
	DeleteOrganizationQuotaRequest    = "DeleteOrganizationQuota"
	GetOrganizationQuotaRequest       = "GetOrganizationQuota"
	GetOrganizationQuotasRequest      = "GetOrganizationQuotas"
	PatchOrganizationQuotaRequest     = "PatchOrganizationQuota"
)

// APIRoutes is a list of routes used by the router to construct request URLs.
var APIRoutes = []Route{
	// v3 droplet
	{Resource: DropletsResource, Path: "/", Method: http.MethodPost, Name: PostDropletRequest},
	{Resource: PackagesResource, Path: "/:package_guid/droplets", Method: http.MethodGet, Name: GetPackageDropletsRequest},
	{Resource: DropletsResource, Path: "/:droplet_guid/upload", Method: http.MethodPost, Name: PostDropletBitsRequest},
	{Resource: DropletsResource, Path: "/:droplet_guid/download", Method: http.MethodGet, Name: GetDropletBitsRequest},

	// v3 package
	{Resource: PackagesResource, Path: "/:package_guid/upload", Method: http.MethodPost, Name: PostPackageBitsRequest},

	// v3 service credential binding
	{Resource: ServiceCredentialBindingsResource, Path: "/", Method: http.MethodGet, Name: GetServiceCredentialBindingsRequest},
	{Resource: ServiceCredentialBindingsResource, Path: "/:service_credential_binding_guid/details", Method: http.MethodGet, Name: GetServiceCredentialBindingDetailsRequest},
	{Resource: ServiceCredentialBindingsResource, Path: "/:service_credential_binding_guid", Method: http.MethodDelete, Name: DeleteServiceCredentialBindingRequest},
	{Resource: ServiceCredentialBindingsResource, Path: "/", Method: http.MethodPost, Name: PostServiceCredentialBindingRequest},

	// v3 service instance
	{Resource: ServiceInstancesResource, Path: "/:service_instance_guid/parameters", Method: http.MethodGet, Name: GetServiceInstanceParametersRequest},
	{Resource: ServiceInstancesResource, Path: "/:service_instance_guid/credentials", Method: http.MethodGet, Name: GetServiceInstanceCredentialsRequest},
	{Resource: ServiceInstancesResource, Path: "/", Method: http.MethodPost, Name: PostServiceInstanceRequest},
	{Resource: ServiceInstancesResource, Path: "/:service_instance_guid", Method: http.MethodPatch, Name: PatchServiceInstanceRequest},
	{Resource: ServiceInstancesResource, Path: "/:service_instance_guid", Method: http.MethodDelete, Name: DeleteServiceInstanceRequest},
	{Resource: ServiceInstancesResource, Path: "/:service_instance_guid/relationships/shared_spaces", Method: http.MethodGet, Name: GetServiceInstanceRelationshipsSharedSpacesRequest},
	{Resource: ServiceInstancesResource, Path: "/:service_instance_guid/relationships/shared_spaces/usage_summary", Method: http.MethodGet, Name: GetServiceInstanceSharedSpacesUsageSummaryRequest},

	// v3 process add missing endpoints
	{Resource: ProcessesResource, Path: "/", Method: http.MethodGet, Name: GetProcessesRequest},
	{Resource: ProcessesResource, Path: "/:process_guid", Method: http.MethodGet, Name: GetProcessRequest},

	{Resource: RoutesResource, Path: "/:route_guid/destinations/:destination_guid", Method: http.MethodDelete, Name: UnmapRouteRequest},
	{Resource: RoutesResource, Path: "/:route_guid/destinations", Method: http.MethodPost, Name: MapRouteRequest},
	{Resource: RoutesResource, Path: "/", Method: http.MethodGet, Name: GetRoutesRequest},
	{Resource: RoutesResource, Path: "/:route_guid/destinations", Method: http.MethodGet, Name: GetRouteDestinationsRequest},
	{Resource: AppsResource, Path: "/:app_guid/routes", Method: http.MethodGet, Name: GetApplicationRoutesRequest},
	{Resource: SpacesResource, Path: "/:space_guid/routes", Method: http.MethodDelete, Name: DeleteOrphanedRoutesRequest},
	{Resource: RoutesResource, Path: "/", Method: http.MethodPost, Name: PostRouteRequest},
	{Resource: RoutesResource, Path: "/:route_guid", Method: http.MethodDelete, Name: DeleteRouteRequest},
	{Resource: RoutesResource, Path: "/:route_guid", Method: http.MethodPatch, Name: PatchRouteRequest},

	// v3 application feature
	{Resource: AppsResource, Path: "/:app_guid/ssh_enabled", Method: http.MethodGet, Name: GetSSHEnabled},
	{Resource: AppsResource, Path: "/:app_guid/features/:name", Method: http.MethodGet, Name: GetApplicationFeaturesRequest},
	{Resource: AppsResource, Path: "/:app_guid/features/:name", Method: http.MethodPatch, Name: PatchApplicationFeaturesRequest},

	// v3 environment varirable
	{Resource: EnvironmentVariableGroupsResource, Path: "/:group_name", Method: http.MethodGet, Name: GetEnvironmentVariableGroupRequest},
	{Resource: EnvironmentVariableGroupsResource, Path: "/:group_name", Method: http.MethodPatch, Name: PatchEnvironmentVariableGroupRequest},

	// v3 domain add missing
	{Resource: DomainsResource, Path: "/:domain_guid/route_reservations", Method: http.MethodGet, Name: GetDomainRouteReservationsRequest},
	{Resource: DomainsResource, Path: "/:domain_guid", Method: http.MethodGet, Name: GetDomainRequest},

	// v3 space add missing
	{Resource: SpacesResource, Path: "/", Method: http.MethodPost, Name: PostSpaceRequest},
	{Resource: SpacesResource, Path: "/:space_guid", Method: http.MethodDelete, Name: DeleteSpaceRequest},
	{Resource: SpacesResource, Path: "/:space_guid", Method: http.MethodPatch, Name: PatchSpaceRequest},

	// v3 org add missing
	{Resource: OrgsResource, Path: "/:organization_guid", Method: http.MethodGet, Name: GetOrganizationRequest},

	// v3 revision
	{Resource: AppsResource, Path: "/:app_guid/revisions", Method: http.MethodGet, Name: GetApplicationRevisionsRequest},
	{Resource: AppsResource, Path: "/:app_guid/revisions/deployed", Method: http.MethodGet, Name: GetApplicationRevisionsDeployedRequest},

	// v3 organization quota
	{Resource: OrgQuotasResource, Path: "/:quota_guid/relationships/organizations", Method: http.MethodPost, Name: PostOrganizationQuotaApplyRequest},
	{Resource: OrgQuotasResource, Path: "/", Method: http.MethodPost, Name: PostOrganizationQuotaRequest},
	{Resource: OrgQuotasResource, Path: "/:quota_guid", Method: http.MethodDelete, Name: DeleteOrganizationQuotaRequest},
	{Resource: OrgQuotasResource, Path: "/:quota_guid", Method: http.MethodGet, Name: GetOrganizationQuotaRequest},
	{Resource: OrgQuotasResource, Path: "/", Method: http.MethodGet, Name: GetOrganizationQuotasRequest},
	{Resource: OrgQuotasResource, Path: "/:quota_guid", Method: http.MethodPatch, Name: PatchOrganizationQuotaRequest},

	{Resource: ServicePlansResource, Path: "/", Method: http.MethodGet, Name: GetServicePlansRequest},

	{Resource: ServiceOfferingsResource, Path: "/", Method: http.MethodGet, Name: GetServiceOfferingsRequest},
	{Resource: ServiceOfferingsResource, Path: "/:service_offering_guid", Method: http.MethodDelete, Name: DeleteServiceOfferingRequest},

	{Resource: OrgsResource, Path: "/", Method: http.MethodGet, Name: GetOrganizationsRequest},
	{Resource: OrgsResource, Path: "/:organization_guid/", Method: http.MethodDelete, Name: DeleteOrganizationRequest},
	{Resource: OrgsResource, Path: "/", Method: http.MethodPost, Name: PostOrganizationRequest},
	{Resource: OrgsResource, Path: "/:organization_guid/domains/default", Method: http.MethodGet, Name: GetDefaultDomainRequest},

	{Resource: AppsResource, Path: "/", Method: http.MethodGet, Name: GetApplicationsRequest},
	{Resource: AppsResource, Path: "/", Method: http.MethodPost, Name: PostApplicationRequest},
	{Resource: AppsResource, Path: "/:app_guid", Method: http.MethodDelete, Name: DeleteApplicationRequest},
	{Resource: AppsResource, Path: "/:app_guid", Method: http.MethodPatch, Name: PatchApplicationRequest},
	{Resource: AppsResource, Path: "/:app_guid/actions/apply_manifest", Method: http.MethodPost, Name: PostApplicationActionApplyManifest},
	{Resource: AppsResource, Path: "/:app_guid/actions/restart", Method: http.MethodPost, Name: PostApplicationActionRestartRequest},
	{Resource: AppsResource, Path: "/:app_guid/actions/start", Method: http.MethodPost, Name: PostApplicationActionStartRequest},
	{Resource: AppsResource, Path: "/:app_guid/actions/stop", Method: http.MethodPost, Name: PostApplicationActionStopRequest},
	{Resource: AppsResource, Path: "/:app_guid/droplets/current", Method: http.MethodGet, Name: GetApplicationDropletCurrentRequest},
	{Resource: AppsResource, Path: "/:app_guid/env", Method: http.MethodGet, Name: GetApplicationEnvRequest},
	{Resource: AppsResource, Path: "/:app_guid/environment_variables", Method: http.MethodPatch, Name: PatchApplicationEnvironmentVariablesRequest},
	{Resource: AppsResource, Path: "/:app_guid/manifest", Method: http.MethodGet, Name: GetApplicationManifestRequest},
	{Resource: AppsResource, Path: "/:app_guid/processes", Method: http.MethodGet, Name: GetApplicationProcessesRequest},
	{Resource: AppsResource, Path: "/:app_guid/processes/:type", Method: http.MethodGet, Name: GetApplicationProcessRequest},
	{Resource: AppsResource, Path: "/:app_guid/processes/:type/actions/scale", Method: http.MethodPost, Name: PostApplicationProcessActionScaleRequest},
	{Resource: AppsResource, Path: "/:app_guid/processes/:type/instances/:index", Method: http.MethodDelete, Name: DeleteApplicationProcessInstanceRequest},
	{Resource: AppsResource, Path: "/:app_guid/relationships/current_droplet", Method: http.MethodPatch, Name: PatchApplicationCurrentDropletRequest},
	{Resource: AppsResource, Path: "/:app_guid/tasks", Method: http.MethodGet, Name: GetApplicationTasksRequest},
	{Resource: AppsResource, Path: "/:app_guid/tasks", Method: http.MethodPost, Name: PostApplicationTasksRequest},
	{Resource: BuildpacksResource, Path: "/", Method: http.MethodGet, Name: GetBuildpacksRequest},
	{Resource: BuildpacksResource, Path: "/", Method: http.MethodPost, Name: PostBuildpackRequest},
	{Resource: BuildpacksResource, Path: "/:buildpack_guid", Method: http.MethodPatch, Name: PatchBuildpackRequest},
	{Resource: BuildpacksResource, Path: "/:buildpack_guid/upload", Method: http.MethodPost, Name: PostBuildpackBitsRequest},
	{Resource: BuildpacksResource, Path: "/:buildpack_guid", Method: http.MethodDelete, Name: DeleteBuildpackRequest},
	{Resource: BuildsResource, Path: "/", Method: http.MethodPost, Name: PostBuildRequest},
	{Resource: BuildsResource, Path: "/:build_guid", Method: http.MethodGet, Name: GetBuildRequest},
	{Resource: DeploymentsResource, Path: "/", Method: http.MethodGet, Name: GetDeploymentsRequest},
	{Resource: DeploymentsResource, Path: "/", Method: http.MethodPost, Name: PostApplicationDeploymentRequest},
	{Resource: DeploymentsResource, Path: "/:deployment_guid", Method: http.MethodGet, Name: GetDeploymentRequest},
	{Resource: DeploymentsResource, Path: "/:deployment_guid/actions/cancel", Method: http.MethodPost, Name: PostApplicationDeploymentActionCancelRequest},
	{Resource: DomainsResource, Path: "/", Method: http.MethodPost, Name: PostDomainRequest},
	{Resource: DomainsResource, Path: "/", Method: http.MethodGet, Name: GetDomainsRequest},
	{Resource: DomainsResource, Path: "/:domain_guid", Method: http.MethodDelete, Name: DeleteDomainRequest},
	{Resource: DomainsResource, Path: "/:domain_guid/relationships/shared_organizations", Method: http.MethodPost, Name: SharePrivateDomainRequest},
	{Resource: DomainsResource, Path: "/:domain_guid/relationships/shared_organizations/:org_guid", Method: http.MethodDelete, Name: DeleteSharedOrgFromDomainRequest},
	{Resource: DropletsResource, Path: "/", Method: http.MethodGet, Name: GetDropletsRequest},
	{Resource: DropletsResource, Path: "/:droplet_guid", Method: http.MethodGet, Name: GetDropletRequest},
	{Resource: FeatureFlagsResource, Path: "/:name", Method: http.MethodGet, Name: GetFeatureFlagRequest},
	{Resource: FeatureFlagsResource, Path: "/:name", Method: http.MethodPatch, Name: PatchFeatureFlagRequest},
	{Resource: FeatureFlagsResource, Path: "/", Method: http.MethodGet, Name: GetFeatureFlagsRequest},
	{Resource: IsolationSegmentsResource, Path: "/", Method: http.MethodGet, Name: GetIsolationSegmentsRequest},
	{Resource: IsolationSegmentsResource, Path: "/", Method: http.MethodPost, Name: PostIsolationSegmentsRequest},
	{Resource: IsolationSegmentsResource, Path: "/:isolation_segment_guid", Method: http.MethodDelete, Name: DeleteIsolationSegmentRequest},
	{Resource: IsolationSegmentsResource, Path: "/:isolation_segment_guid", Method: http.MethodGet, Name: GetIsolationSegmentRequest},
	{Resource: IsolationSegmentsResource, Path: "/:isolation_segment_guid/organizations", Method: http.MethodGet, Name: GetIsolationSegmentOrganizationsRequest},
	{Resource: IsolationSegmentsResource, Path: "/:isolation_segment_guid/relationships/organizations", Method: http.MethodPost, Name: PostIsolationSegmentRelationshipOrganizationsRequest},
	{Resource: IsolationSegmentsResource, Path: "/:isolation_segment_guid/relationships/organizations/:organization_guid", Method: http.MethodDelete, Name: DeleteIsolationSegmentRelationshipOrganizationRequest},
	{Resource: OrgsResource, Path: "/", Method: http.MethodGet, Name: GetOrganizationsRequest},
	{Resource: OrgsResource, Path: "/:organization_guid/domains", Method: http.MethodGet, Name: GetOrganizationDomainsRequest},
	{Resource: OrgsResource, Path: "/:organization_guid/relationships/default_isolation_segment", Method: http.MethodGet, Name: GetOrganizationRelationshipDefaultIsolationSegmentRequest},
	{Resource: OrgsResource, Path: "/:organization_guid/relationships/default_isolation_segment", Method: http.MethodPatch, Name: PatchOrganizationRelationshipDefaultIsolationSegmentRequest},
	{Resource: OrgsResource, Path: "/:organization_guid/", Method: http.MethodPatch, Name: PatchOrganizationRequest},
	{Resource: PackagesResource, Path: "/", Method: http.MethodGet, Name: GetPackagesRequest},
	{Resource: PackagesResource, Path: "/", Method: http.MethodPost, Name: PostPackageRequest},
	{Resource: PackagesResource, Path: "/:package_guid", Method: http.MethodGet, Name: GetPackageRequest},
	{Resource: ProcessesResource, Path: "/:process_guid", Method: http.MethodPatch, Name: PatchProcessRequest},
	{Resource: ProcessesResource, Path: "/:process_guid/stats", Method: http.MethodGet, Name: GetProcessStatsRequest},
	{Resource: ResourceMatches, Path: "/", Method: http.MethodPost, Name: PostResourceMatchesRequest},
	{Resource: ServiceInstancesResource, Path: "/", Method: http.MethodGet, Name: GetServiceInstancesRequest},
	{Resource: ServiceInstancesResource, Path: "/:service_instance_guid/relationships/shared_spaces", Method: http.MethodPost, Name: PostServiceInstanceRelationshipsSharedSpacesRequest},
	{Resource: ServiceInstancesResource, Path: "/:service_instance_guid/relationships/shared_spaces/:space_guid", Method: http.MethodDelete, Name: DeleteServiceInstanceRelationshipsSharedSpaceRequest},
	{Resource: SpacesResource, Path: "/", Method: http.MethodGet, Name: GetSpacesRequest},
	{Resource: SpacesResource, Path: "/:space_guid/relationships/isolation_segment", Method: http.MethodGet, Name: GetSpaceRelationshipIsolationSegmentRequest},
	{Resource: SpacesResource, Path: "/:space_guid/relationships/isolation_segment", Method: http.MethodPatch, Name: PatchSpaceRelationshipIsolationSegmentRequest},
	{Resource: SpacesResource, Path: "/:space_guid/actions/apply_manifest", Method: http.MethodPost, Name: PostSpaceActionApplyManifestRequest},
	{Resource: StacksResource, Path: "/", Method: http.MethodGet, Name: GetStacksRequest},
	{Resource: TasksResource, Path: "/:task_guid/cancel", Method: http.MethodPut, Name: PutTaskCancelRequest},
}
