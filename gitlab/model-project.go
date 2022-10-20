package gitlab

import (
	"io"
	"time"
)

type VisibilityValue string

const (
	PrivateVisibility  VisibilityValue = "private"
	InternalVisibility VisibilityValue = "internal"
	PublicVisibility   VisibilityValue = "public"
)

type ForkProjectOption struct {
	Description                   string          `json:"description,omitempty" description:"project description" required:"false"`
	MergeRequestDefaultTargetSelf bool            `json:"mr_default_target_self,omitempty" description:"true merge-reuqest this project; false upstream project" required:"false"`
	Name                          string          `json:"name,omitempty" description:"new project name" required:"false"`
	NamespaceID                   int             `json:"namespace_id,omitempty" description:"project namespace id" required:"false"`
	NamespacePath                 string          `json:"namespace_path,omitempty" description:"project namespace path" required:"false"`
	Path                          string          `json:"path,omitempty" description:"fork project path" required:"false"`
	Visibility                    VisibilityValue `json:"visibility,omitempty" description:"project visibility" required:"false"`
}

type ForkProject struct {
	ID                                        int             `json:"id"`
	Description                               string          `json:"description"`
	DefaultBranch                             string          `json:"default_branch"`
	Public                                    bool            `json:"public"`
	Visibility                                VisibilityValue `json:"visibility"`
	SSHURLToRepo                              string          `json:"ssh_url_to_repo"`
	HTTPURLToRepo                             string          `json:"http_url_to_repo"`
	WebURL                                    string          `json:"web_url"`
	ReadmeURL                                 string          `json:"readme_url"`
	TagList                                   []string        `json:"tag_list"`
	Topics                                    []string        `json:"topics"`
	Name                                      string          `json:"name"`
	NameWithNamespace                         string          `json:"name_with_namespace"`
	Path                                      string          `json:"path"`
	PathWithNamespace                         string          `json:"path_with_namespace"`
	IssuesEnabled                             bool            `json:"issues_enabled"`
	OpenIssuesCount                           int             `json:"open_issues_count"`
	MergeRequestsEnabled                      bool            `json:"merge_requests_enabled"`
	ApprovalsBeforeMerge                      int             `json:"approvals_before_merge"`
	JobsEnabled                               bool            `json:"jobs_enabled"`
	WikiEnabled                               bool            `json:"wiki_enabled"`
	SnippetsEnabled                           bool            `json:"snippets_enabled"`
	ResolveOutdatedDiffDiscussions            bool            `json:"resolve_outdated_diff_discussions"`
	ContainerRegistryEnabled                  bool            `json:"container_registry_enabled"`
	ContainerRegistryImagePrefix              string          `json:"container_registry_image_prefix,omitempty"`
	CreatedAt                                 *time.Time      `json:"created_at,omitempty"`
	LastActivityAt                            *time.Time      `json:"last_activity_at,omitempty"`
	CreatorID                                 int             `json:"creator_id"`
	ImportStatus                              string          `json:"import_status"`
	ImportError                               string          `json:"import_error"`
	EmptyRepo                                 bool            `json:"empty_repo"`
	Archived                                  bool            `json:"archived"`
	AvatarURL                                 string          `json:"avatar_url"`
	LicenseURL                                string          `json:"license_url"`
	SharedRunnersEnabled                      bool            `json:"shared_runners_enabled"`
	ForksCount                                int             `json:"forks_count"`
	StarCount                                 int             `json:"star_count"`
	RunnersToken                              string          `json:"runners_token"`
	PublicBuilds                              bool            `json:"public_builds"`
	AllowMergeOnSkippedPipeline               bool            `json:"allow_merge_on_skipped_pipeline"`
	OnlyAllowMergeIfPipelineSucceeds          bool            `json:"only_allow_merge_if_pipeline_succeeds"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool            `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool            `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool            `json:"printing_merge_request_link_enabled"`
	LFSEnabled                                bool            `json:"lfs_enabled"`
	RepositoryStorage                         string          `json:"repository_storage"`
	RequestAccessEnabled                      bool            `json:"request_access_enabled"`
	Mirror                                    bool            `json:"mirror"`
	MirrorUserID                              int             `json:"mirror_user_id"`
	MirrorTriggerBuilds                       bool            `json:"mirror_trigger_builds"`
	OnlyMirrorProtectedBranches               bool            `json:"only_mirror_protected_branches"`
	MirrorOverwritesDivergedBranches          bool            `json:"mirror_overwrites_diverged_branches"`
	PackagesEnabled                           bool            `json:"packages_enabled"`
	ServiceDeskEnabled                        bool            `json:"service_desk_enabled"`
	ServiceDeskAddress                        string          `json:"service_desk_address"`
	AutocloseReferencedIssues                 bool            `json:"autoclose_referenced_issues"`
	SuggestionCommitMessage                   string          `json:"suggestion_commit_message"`
	AutoCancelPendingPipelines                string          `json:"auto_cancel_pending_pipelines"`
	CIForwardDeploymentEnabled                bool            `json:"ci_forward_deployment_enabled"`
	SharedWithGroups                          []struct {
		GroupID          int    `json:"group_id"`
		GroupName        string `json:"group_name"`
		GroupAccessLevel int    `json:"group_access_level"`
	} `json:"shared_with_groups"`
	CIConfigPath                             string   `json:"ci_config_path"`
	CIDefaultGitDepth                        int      `json:"ci_default_git_depth"`
	ComplianceFrameworks                     []string `json:"compliance_frameworks"`
	BuildCoverageRegex                       string   `json:"build_coverage_regex"`
	BuildTimeout                             int      `json:"build_timeout"`
	IssuesTemplate                           string   `json:"issues_template"`
	MergeRequestsTemplate                    string   `json:"merge_requests_template"`
	KeepLatestArtifact                       bool     `json:"keep_latest_artifact"`
	MergePipelinesEnabled                    bool     `json:"merge_pipelines_enabled"`
	MergeTrainsEnabled                       bool     `json:"merge_trains_enabled"`
	RestrictUserDefinedVariables             bool     `json:"restrict_user_defined_variables"`
	MergeCommitTemplate                      string   `json:"merge_commit_template"`
	SquashCommitTemplate                     string   `json:"squash_commit_template"`
	AutoDevopsDeployStrategy                 string   `json:"auto_devops_deploy_strategy"`
	AutoDevopsEnabled                        bool     `json:"auto_devops_enabled"`
	BuildGitStrategy                         string   `json:"build_git_strategy"`
	EmailsDisabled                           bool     `json:"emails_disabled"`
	ExternalAuthorizationClassificationLabel string   `json:"external_authorization_classification_label"`
}

type ProjectAvatar struct {
	Filename string
	Image    io.Reader
}

type EditProjectOptions struct {
	AllowMergeOnSkippedPipeline               *bool                                `json:"allow_merge_on_skipped_pipeline,omitempty"`
	Avatar                                    *ProjectAvatar                       `url:"-" json:"-"`
	AnalyticsAccessLevel                      *string                              `json:"analytics_access_level,omitempty"`
	ApprovalsBeforeMerge                      *int                                 `json:"approvals_before_merge,omitempty"`
	AutoCancelPendingPipelines                *string                              `json:"auto_cancel_pending_pipelines,omitempty"`
	AutoDevopsDeployStrategy                  *string                              `json:"auto_devops_deploy_strategy,omitempty"`
	AutoDevopsEnabled                         *bool                                `json:"auto_devops_enabled,omitempty"`
	AutocloseReferencedIssues                 *bool                                `json:"autoclose_referenced_issues,omitempty"`
	BuildCoverageRegex                        *string                              `json:"build_coverage_regex,omitempty"`
	BuildGitStrategy                          *string                              `json:"build_git_strategy,omitempty"`
	BuildTimeout                              *int                                 `json:"build_timeout,omitempty"`
	BuildsAccessLevel                         *string                              `json:"builds_access_level,omitempty"`
	CIConfigPath                              *string                              `json:"ci_config_path,omitempty"`
	CIDefaultGitDepth                         *int                                 `json:"ci_default_git_depth,omitempty"`
	ContainerExpirationPolicyAttributes       *ContainerExpirationPolicyAttributes `json:"container_expiration_policy_attributes,omitempty"`
	ContainerRegistryAccessLevel              *string                              `json:"container_registry_access_level,omitempty"`
	DefaultBranch                             *string                              `json:"default_branch,omitempty" description:"默认分支"`
	Description                               *string                              `json:"description,omitempty" description:"项目描述"`
	EmailsDisabled                            *bool                                `json:"emails_disabled,omitempty" description:"是否禁用电子邮件"`
	ExternalAuthorizationClassificationLabel  *string                              `json:"external_authorization_classification_label,omitempty"`
	ForkingAccessLevel                        *string                              `json:"forking_access_level,omitempty"`
	ImportURL                                 *string                              `json:"import_url,omitempty"`
	IssuesAccessLevel                         *string                              `json:"issues_access_level,omitempty"`
	KeepLatestArtifact                        *bool                                `json:"keep_latest_artifact,omitempty"`
	LFSEnabled                                *bool                                `json:"lfs_enabled,omitempty"`
	MergeCommitTemplate                       *string                              `json:"merge_commit_template,omitempty"`
	MergeMethod                               *string                              `json:"merge_method,omitempty"`
	MergePipelinesEnabled                     *bool                                `json:"merge_pipelines_enabled,omitempty"`
	MergeRequestsAccessLevel                  *string                              `json:"merge_requests_access_level,omitempty"`
	MergeTrainsEnabled                        *bool                                `json:"merge_trains_enabled,omitempty"`
	Mirror                                    *bool                                `json:"mirror,omitempty"`
	MirrorOverwritesDivergedBranches          *bool                                `json:"mirror_overwrites_diverged_branches,omitempty"`
	MirrorTriggerBuilds                       *bool                                `json:"mirror_trigger_builds,omitempty"`
	MirrorUserID                              *int                                 `json:"mirror_user_id,omitempty"`
	Name                                      *string                              `json:"name,omitempty" description:"项目名称"`
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool                                `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          *bool                                `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	OnlyMirrorProtectedBranches               *bool                                `json:"only_mirror_protected_branches,omitempty"`
	OperationsAccessLevel                     *string                              `json:"operations_access_level,omitempty"`
	PackagesEnabled                           *bool                                `json:"packages_enabled,omitempty"`
	PagesAccessLevel                          *string                              `json:"pages_access_level,omitempty"`
	Path                                      *string                              `json:"path,omitempty"`
	PublicBuilds                              *bool                                `json:"public_builds,omitempty"`
	RemoveSourceBranchAfterMerge              *bool                                `json:"remove_source_branch_after_merge,omitempty"`
	PrintingMergeRequestLinkEnabled           *bool                                `json:"printing_merge_request_link_enabled,omitempty"`
	RepositoryAccessLevel                     *string                              `json:"repository_access_level,omitempty"`
	RepositoryStorage                         *string                              `json:"repository_storage,omitempty"`
	RequestAccessEnabled                      *bool                                `json:"request_access_enabled,omitempty"`
	RequirementsAccessLevel                   *string                              `json:"requirements_access_level,omitempty"`
	ResolveOutdatedDiffDiscussions            *bool                                `json:"resolve_outdated_diff_discussions,omitempty"`
	RestrictUserDefinedVariables              *bool                                `json:"restrict_user_defined_variables,omitempty"`
	SecurityAndComplianceAccessLevel          *string                              `json:"security_and_compliance_access_level,omitempty"`
	SharedRunnersEnabled                      *bool                                `json:"shared_runners_enabled,omitempty"`
	ShowDefaultAwardEmojis                    *bool                                `json:"show_default_award_emojis,omitempty"`
	SnippetsAccessLevel                       *string                              `json:"snippets_access_level,omitempty"`
	SquashCommitTemplate                      *string                              `json:"squash_commit_template,omitempty"`
	SquashOption                              *string                              `json:"squash_option,omitempty"`
	SuggestionCommitMessage                   *string                              `json:"suggestion_commit_message,omitempty"`
	Topics                                    *[]string                            `json:"topics,omitempty"`
	Visibility                                *VisibilityValue                     `json:"visibility,omitempty" description:"项目可见性"`
	WikiAccessLevel                           *string                              `json:"wiki_access_level,omitempty"`
}

type ContainerExpirationPolicyAttributes struct {
	Cadence         *string `url:"cadence,omitempty" json:"cadence,omitempty"`
	KeepN           *int    `url:"keep_n,omitempty" json:"keep_n,omitempty"`
	OlderThan       *string `url:"older_than,omitempty" json:"older_than,omitempty"`
	NameRegexDelete *string `url:"name_regex_delete,omitempty" json:"name_regex_delete,omitempty"`
	NameRegexKeep   *string `url:"name_regex_keep,omitempty" json:"name_regex_keep,omitempty"`
	Enabled         *bool   `url:"enabled,omitempty" json:"enabled,omitempty"`
}

type CreateProjectOptions struct {
	AllowMergeOnSkippedPipeline               *bool                                `url:"allow_merge_on_skipped_pipeline,omitempty" json:"allow_merge_on_skipped_pipeline,omitempty"`
	AnalyticsAccessLevel                      *string                              `url:"analytics_access_level,omitempty" json:"analytics_access_level,omitempty"`
	ApprovalsBeforeMerge                      *int                                 `url:"approvals_before_merge,omitempty" json:"approvals_before_merge,omitempty"`
	AutoCancelPendingPipelines                *string                              `url:"auto_cancel_pending_pipelines,omitempty" json:"auto_cancel_pending_pipelines,omitempty"`
	AutoDevopsDeployStrategy                  *string                              `url:"auto_devops_deploy_strategy,omitempty" json:"auto_devops_deploy_strategy,omitempty"`
	AutoDevopsEnabled                         *bool                                `url:"auto_devops_enabled,omitempty" json:"auto_devops_enabled,omitempty"`
	AutocloseReferencedIssues                 *bool                                `url:"autoclose_referenced_issues,omitempty" json:"autoclose_referenced_issues,omitempty"`
	BuildCoverageRegex                        *string                              `url:"build_coverage_regex,omitempty" json:"build_coverage_regex,omitempty"`
	BuildGitStrategy                          *string                              `url:"build_git_strategy,omitempty" json:"build_git_strategy,omitempty"`
	BuildTimeout                              *int                                 `url:"build_timeout,omitempty" json:"build_timeout,omitempty"`
	BuildsAccessLevel                         *string                              `url:"builds_access_level,omitempty" json:"builds_access_level,omitempty"`
	CIConfigPath                              *string                              `url:"ci_config_path,omitempty" json:"ci_config_path,omitempty"`
	ContainerExpirationPolicyAttributes       *ContainerExpirationPolicyAttributes `url:"container_expiration_policy_attributes,omitempty" json:"container_expiration_policy_attributes,omitempty"`
	ContainerRegistryAccessLevel              *string                              `url:"container_registry_access_level,omitempty" json:"container_registry_access_level,omitempty"`
	DefaultBranch                             *string                              `url:"default_branch,omitempty" json:"default_branch,omitempty"`
	Description                               *string                              `url:"description,omitempty" json:"description,omitempty"`
	EmailsDisabled                            *bool                                `url:"emails_disabled,omitempty" json:"emails_disabled,omitempty"`
	ExternalAuthorizationClassificationLabel  *string                              `url:"external_authorization_classification_label,omitempty" json:"external_authorization_classification_label,omitempty"`
	ForkingAccessLevel                        *string                              `url:"forking_access_level,omitempty" json:"forking_access_level,omitempty"`
	GroupWithProjectTemplatesID               *int                                 `url:"group_with_project_templates_id,omitempty" json:"group_with_project_templates_id,omitempty"`
	ImportURL                                 *string                              `url:"import_url,omitempty" json:"import_url,omitempty"`
	InitializeWithReadme                      *bool                                `url:"initialize_with_readme,omitempty" json:"initialize_with_readme,omitempty"`
	IssuesAccessLevel                         *string                              `url:"issues_access_level,omitempty" json:"issues_access_level,omitempty"`
	LFSEnabled                                *bool                                `url:"lfs_enabled,omitempty" json:"lfs_enabled,omitempty"`
	MergeCommitTemplate                       *string                              `url:"merge_commit_template,omitempty" json:"merge_commit_template,omitempty"`
	MergeMethod                               *string                              `url:"merge_method,omitempty" json:"merge_method,omitempty"`
	MergePipelinesEnabled                     *bool                                `url:"merge_pipelines_enabled,omitempty" json:"merge_pipelines_enabled,omitempty"`
	MergeRequestsAccessLevel                  *string                              `url:"merge_requests_access_level,omitempty" json:"merge_requests_access_level,omitempty"`
	MergeTrainsEnabled                        *bool                                `url:"merge_trains_enabled,omitempty" json:"merge_trains_enabled,omitempty"`
	Mirror                                    *bool                                `url:"mirror,omitempty" json:"mirror,omitempty"`
	MirrorTriggerBuilds                       *bool                                `url:"mirror_trigger_builds,omitempty" json:"mirror_trigger_builds,omitempty"`
	Name                                      *string                              `url:"name,omitempty" json:"name,omitempty"`
	NamespaceID                               *int                                 `url:"namespace_id,omitempty" json:"namespace_id,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool                                `url:"only_allow_merge_if_all_discussions_are_resolved,omitempty" json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          *bool                                `url:"only_allow_merge_if_pipeline_succeeds,omitempty" json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	OperationsAccessLevel                     *string                              `url:"operations_access_level,omitempty" json:"operations_access_level,omitempty"`
	PackagesEnabled                           *bool                                `url:"packages_enabled,omitempty" json:"packages_enabled,omitempty"`
	PagesAccessLevel                          *string                              `url:"pages_access_level,omitempty" json:"pages_access_level,omitempty"`
	Path                                      *string                              `url:"path,omitempty" json:"path,omitempty"`
	PublicBuilds                              *bool                                `url:"public_builds,omitempty" json:"public_builds,omitempty"`
	RemoveSourceBranchAfterMerge              *bool                                `url:"remove_source_branch_after_merge,omitempty" json:"remove_source_branch_after_merge,omitempty"`
	PrintingMergeRequestLinkEnabled           *bool                                `url:"printing_merge_request_link_enabled,omitempty" json:"printing_merge_request_link_enabled,omitempty"`
	RepositoryAccessLevel                     *string                              `url:"repository_access_level,omitempty" json:"repository_access_level,omitempty"`
	RepositoryStorage                         *string                              `url:"repository_storage,omitempty" json:"repository_storage,omitempty"`
	RequestAccessEnabled                      *bool                                `url:"request_access_enabled,omitempty" json:"request_access_enabled,omitempty"`
	RequirementsAccessLevel                   *string                              `url:"requirements_access_level,omitempty" json:"requirements_access_level,omitempty"`
	ResolveOutdatedDiffDiscussions            *bool                                `url:"resolve_outdated_diff_discussions,omitempty" json:"resolve_outdated_diff_discussions,omitempty"`
	SecurityAndComplianceAccessLevel          *string                              `url:"security_and_compliance_access_level,omitempty" json:"security_and_compliance_access_level,omitempty"`
	SharedRunnersEnabled                      *bool                                `url:"shared_runners_enabled,omitempty" json:"shared_runners_enabled,omitempty"`
	ShowDefaultAwardEmojis                    *bool                                `url:"show_default_award_emojis,omitempty" json:"show_default_award_emojis,omitempty"`
	SnippetsAccessLevel                       *string                              `url:"snippets_access_level,omitempty" json:"snippets_access_level,omitempty"`
	SquashCommitTemplate                      *string                              `url:"squash_commit_template,omitempty" json:"squash_commit_template,omitempty"`
	SquashOption                              *string                              `url:"squash_option,omitempty" json:"squash_option,omitempty"`
	SuggestionCommitMessage                   *string                              `url:"suggestion_commit_message,omitempty" json:"suggestion_commit_message,omitempty"`
	TemplateName                              *string                              `url:"template_name,omitempty" json:"template_name,omitempty"`
	TemplateProjectID                         *int                                 `url:"template_project_id,omitempty" json:"template_project_id,omitempty"`
	Topics                                    *[]string                            `url:"topics,omitempty" json:"topics,omitempty"`
	UseCustomTemplate                         *bool                                `url:"use_custom_template,omitempty" json:"use_custom_template,omitempty"`
	Visibility                                *VisibilityValue                     `url:"visibility,omitempty" json:"visibility,omitempty"`
	WikiAccessLevel                           *string                              `url:"wiki_access_level,omitempty" json:"wiki_access_level,omitempty"`
}
