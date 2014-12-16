// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudformation provides a client for AWS CloudFormation.
package cloudformation

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

// CloudFormation is a client for AWS CloudFormation.
type CloudFormation struct {
	client *aws.QueryClient
}

// New returns a new CloudFormation client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *CloudFormation {
	if client == nil {
		client = http.DefaultClient
	}

	service := "cloudformation"
	endpoint, service, region := endpoints.Lookup("cloudformation", region)

	return &CloudFormation{
		client: &aws.QueryClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2010-05-15",
		},
	}
}

// CancelUpdateStack cancels an update on the specified stack. If the call
// completes successfully, the stack will roll back the update and revert
// to the previous stack configuration. Only stacks that are in the state
// can be canceled.
func (c *CloudFormation) CancelUpdateStack(req *CancelUpdateStackInput) (err error) {
	// NRE
	err = c.client.Do("CancelUpdateStack", "POST", "/", req, nil)
	return
}

// CreateStack creates a stack as specified in the template. After the call
// completes successfully, the stack creation starts. You can check the
// status of the stack via the DescribeStacks
func (c *CloudFormation) CreateStack(req *CreateStackInput) (resp *CreateStackResult, err error) {
	resp = &CreateStackResult{}
	err = c.client.Do("CreateStack", "POST", "/", req, resp)
	return
}

// DeleteStack deletes a specified stack. Once the call completes
// successfully, stack deletion starts. Deleted stacks do not show up in
// the DescribeStacks API if the deletion has been completed successfully.
func (c *CloudFormation) DeleteStack(req *DeleteStackInput) (err error) {
	// NRE
	err = c.client.Do("DeleteStack", "POST", "/", req, nil)
	return
}

// DescribeStackEvents returns all stack related events for a specified
// stack. For more information about a stack's event history, go to Stacks
// in the AWS CloudFormation User Guide. You can list events for stacks
// that have failed to create or have been deleted by specifying the unique
// stack identifier (stack
func (c *CloudFormation) DescribeStackEvents(req *DescribeStackEventsInput) (resp *DescribeStackEventsResult, err error) {
	resp = &DescribeStackEventsResult{}
	err = c.client.Do("DescribeStackEvents", "POST", "/", req, resp)
	return
}

// DescribeStackResource returns a description of the specified resource in
// the specified stack. For deleted stacks, DescribeStackResource returns
// resource information for up to 90 days after the stack has been deleted.
func (c *CloudFormation) DescribeStackResource(req *DescribeStackResourceInput) (resp *DescribeStackResourceResult, err error) {
	resp = &DescribeStackResourceResult{}
	err = c.client.Do("DescribeStackResource", "POST", "/", req, resp)
	return
}

// DescribeStackResources returns AWS resource descriptions for running and
// deleted stacks. If StackName is specified, all the associated resources
// that are part of the stack are returned. If PhysicalResourceId is
// specified, the associated resources of the stack that the resource
// belongs to are returned. Only the first 100 resources will be returned.
// If your stack has more resources than this, you should use
// ListStackResources instead. For deleted stacks, DescribeStackResources
// returns resource information for up to 90 days after the stack has been
// deleted. You must specify either StackName or PhysicalResourceId , but
// not both. In addition, you can specify LogicalResourceId to filter the
// returned result. For more information about resources, the
// LogicalResourceId and PhysicalResourceId , go to the AWS CloudFormation
// User Guide A ValidationError is returned if you specify both StackName
// and PhysicalResourceId in the same request.
func (c *CloudFormation) DescribeStackResources(req *DescribeStackResourcesInput) (resp *DescribeStackResourcesResult, err error) {
	resp = &DescribeStackResourcesResult{}
	err = c.client.Do("DescribeStackResources", "POST", "/", req, resp)
	return
}

// DescribeStacks returns the description for the specified stack; if no
// stack name was specified, then it returns the description for all the
// stacks created.
func (c *CloudFormation) DescribeStacks(req *DescribeStacksInput) (resp *DescribeStacksResult, err error) {
	resp = &DescribeStacksResult{}
	err = c.client.Do("DescribeStacks", "POST", "/", req, resp)
	return
}

// EstimateTemplateCost returns the estimated monthly cost of a template.
// The return value is an AWS Simple Monthly Calculator URL with a query
// string that describes the resources required to run the template.
func (c *CloudFormation) EstimateTemplateCost(req *EstimateTemplateCostInput) (resp *EstimateTemplateCostResult, err error) {
	resp = &EstimateTemplateCostResult{}
	err = c.client.Do("EstimateTemplateCost", "POST", "/", req, resp)
	return
}

// GetStackPolicy returns the stack policy for a specified stack. If a
// stack doesn't have a policy, a null value is returned.
func (c *CloudFormation) GetStackPolicy(req *GetStackPolicyInput) (resp *GetStackPolicyResult, err error) {
	resp = &GetStackPolicyResult{}
	err = c.client.Do("GetStackPolicy", "POST", "/", req, resp)
	return
}

// GetTemplate returns the template body for a specified stack. You can get
// the template for running or deleted stacks. For deleted stacks,
// GetTemplate returns the template for up to 90 days after the stack has
// been deleted. If the template does not exist, a ValidationError is
// returned.
func (c *CloudFormation) GetTemplate(req *GetTemplateInput) (resp *GetTemplateResult, err error) {
	resp = &GetTemplateResult{}
	err = c.client.Do("GetTemplate", "POST", "/", req, resp)
	return
}

// GetTemplateSummary returns information about a new or existing template.
// The GetTemplateSummary action is useful for viewing parameter
// information, such as default parameter values and parameter types,
// before you create or update a stack. You can use the GetTemplateSummary
// action when you submit a template, or you can get template information
// for a running or deleted stack. For deleted stacks, GetTemplateSummary
// returns the template information for up to 90 days after the stack has
// been deleted. If the template does not exist, a ValidationError is
// returned.
func (c *CloudFormation) GetTemplateSummary(req *GetTemplateSummaryInput) (resp *GetTemplateSummaryResult, err error) {
	resp = &GetTemplateSummaryResult{}
	err = c.client.Do("GetTemplateSummary", "POST", "/", req, resp)
	return
}

// ListStackResources returns descriptions of all resources of the
// specified stack. For deleted stacks, ListStackResources returns resource
// information for up to 90 days after the stack has been deleted.
func (c *CloudFormation) ListStackResources(req *ListStackResourcesInput) (resp *ListStackResourcesResult, err error) {
	resp = &ListStackResourcesResult{}
	err = c.client.Do("ListStackResources", "POST", "/", req, resp)
	return
}

// ListStacks returns the summary information for stacks whose status
// matches the specified StackStatusFilter. Summary information for stacks
// that have been deleted is kept for 90 days after the stack is deleted.
// If no StackStatusFilter is specified, summary information for all stacks
// is returned (including existing stacks and stacks that have been
// deleted).
func (c *CloudFormation) ListStacks(req *ListStacksInput) (resp *ListStacksResult, err error) {
	resp = &ListStacksResult{}
	err = c.client.Do("ListStacks", "POST", "/", req, resp)
	return
}

// SetStackPolicy is undocumented.
func (c *CloudFormation) SetStackPolicy(req *SetStackPolicyInput) (err error) {
	// NRE
	err = c.client.Do("SetStackPolicy", "POST", "/", req, nil)
	return
}

// SignalResource sends a signal to the specified resource with a success
// or failure status. You can use the SignalResource API in conjunction
// with a creation policy or update policy. AWS CloudFormation doesn't
// proceed with a stack creation or update until resources receive the
// required number of signals or the timeout period is exceeded. The
// SignalResource API is useful in cases where you want to send signals
// from anywhere other than an Amazon EC2 instance.
func (c *CloudFormation) SignalResource(req *SignalResourceInput) (err error) {
	// NRE
	err = c.client.Do("SignalResource", "POST", "/", req, nil)
	return
}

// UpdateStack updates a stack as specified in the template. After the call
// completes successfully, the stack update starts. You can check the
// status of the stack via the DescribeStacks action. To get a copy of the
// template for an existing stack, you can use the GetTemplate action. Tags
// that were associated with this stack during creation time will still be
// associated with the stack after an UpdateStack operation. For more
// information about creating an update template, updating a stack, and
// monitoring the progress of the update, see Updating a Stack
func (c *CloudFormation) UpdateStack(req *UpdateStackInput) (resp *UpdateStackResult, err error) {
	resp = &UpdateStackResult{}
	err = c.client.Do("UpdateStack", "POST", "/", req, resp)
	return
}

// ValidateTemplate is undocumented.
func (c *CloudFormation) ValidateTemplate(req *ValidateTemplateInput) (resp *ValidateTemplateResult, err error) {
	resp = &ValidateTemplateResult{}
	err = c.client.Do("ValidateTemplate", "POST", "/", req, resp)
	return
}

// CancelUpdateStackInput is undocumented.
type CancelUpdateStackInput struct {
	StackName aws.StringValue `xml:"StackName"`
}

// CreateStackInput is undocumented.
type CreateStackInput struct {
	Capabilities     []string         `xml:"Capabilities>member"`
	DisableRollback  aws.BooleanValue `xml:"DisableRollback"`
	NotificationARNs []string         `xml:"NotificationARNs>member"`
	OnFailure        aws.StringValue  `xml:"OnFailure"`
	Parameters       []Parameter      `xml:"Parameters>member"`
	StackName        aws.StringValue  `xml:"StackName"`
	StackPolicyBody  aws.StringValue  `xml:"StackPolicyBody"`
	StackPolicyURL   aws.StringValue  `xml:"StackPolicyURL"`
	Tags             []Tag            `xml:"Tags>member"`
	TemplateBody     aws.StringValue  `xml:"TemplateBody"`
	TemplateURL      aws.StringValue  `xml:"TemplateURL"`
	TimeoutInMinutes aws.IntegerValue `xml:"TimeoutInMinutes"`
}

// CreateStackOutput is undocumented.
type CreateStackOutput struct {
	StackID aws.StringValue `xml:"CreateStackResult>StackId"`
}

// DeleteStackInput is undocumented.
type DeleteStackInput struct {
	StackName aws.StringValue `xml:"StackName"`
}

// DescribeStackEventsInput is undocumented.
type DescribeStackEventsInput struct {
	NextToken aws.StringValue `xml:"NextToken"`
	StackName aws.StringValue `xml:"StackName"`
}

// DescribeStackEventsOutput is undocumented.
type DescribeStackEventsOutput struct {
	NextToken   aws.StringValue `xml:"DescribeStackEventsResult>NextToken"`
	StackEvents []StackEvent    `xml:"DescribeStackEventsResult>StackEvents>member"`
}

// DescribeStackResourceInput is undocumented.
type DescribeStackResourceInput struct {
	LogicalResourceID aws.StringValue `xml:"LogicalResourceId"`
	StackName         aws.StringValue `xml:"StackName"`
}

// DescribeStackResourceOutput is undocumented.
type DescribeStackResourceOutput struct {
	StackResourceDetail *StackResourceDetail `xml:"DescribeStackResourceResult>StackResourceDetail"`
}

// DescribeStackResourcesInput is undocumented.
type DescribeStackResourcesInput struct {
	LogicalResourceID  aws.StringValue `xml:"LogicalResourceId"`
	PhysicalResourceID aws.StringValue `xml:"PhysicalResourceId"`
	StackName          aws.StringValue `xml:"StackName"`
}

// DescribeStackResourcesOutput is undocumented.
type DescribeStackResourcesOutput struct {
	StackResources []StackResource `xml:"DescribeStackResourcesResult>StackResources>member"`
}

// DescribeStacksInput is undocumented.
type DescribeStacksInput struct {
	NextToken aws.StringValue `xml:"NextToken"`
	StackName aws.StringValue `xml:"StackName"`
}

// DescribeStacksOutput is undocumented.
type DescribeStacksOutput struct {
	NextToken aws.StringValue `xml:"DescribeStacksResult>NextToken"`
	Stacks    []Stack         `xml:"DescribeStacksResult>Stacks>member"`
}

// EstimateTemplateCostInput is undocumented.
type EstimateTemplateCostInput struct {
	Parameters   []Parameter     `xml:"Parameters>member"`
	TemplateBody aws.StringValue `xml:"TemplateBody"`
	TemplateURL  aws.StringValue `xml:"TemplateURL"`
}

// EstimateTemplateCostOutput is undocumented.
type EstimateTemplateCostOutput struct {
	URL aws.StringValue `xml:"EstimateTemplateCostResult>Url"`
}

// GetStackPolicyInput is undocumented.
type GetStackPolicyInput struct {
	StackName aws.StringValue `xml:"StackName"`
}

// GetStackPolicyOutput is undocumented.
type GetStackPolicyOutput struct {
	StackPolicyBody aws.StringValue `xml:"GetStackPolicyResult>StackPolicyBody"`
}

// GetTemplateInput is undocumented.
type GetTemplateInput struct {
	StackName aws.StringValue `xml:"StackName"`
}

// GetTemplateOutput is undocumented.
type GetTemplateOutput struct {
	TemplateBody aws.StringValue `xml:"GetTemplateResult>TemplateBody"`
}

// GetTemplateSummaryInput is undocumented.
type GetTemplateSummaryInput struct {
	StackName    aws.StringValue `xml:"StackName"`
	TemplateBody aws.StringValue `xml:"TemplateBody"`
	TemplateURL  aws.StringValue `xml:"TemplateURL"`
}

// GetTemplateSummaryOutput is undocumented.
type GetTemplateSummaryOutput struct {
	Capabilities       []string               `xml:"GetTemplateSummaryResult>Capabilities>member"`
	CapabilitiesReason aws.StringValue        `xml:"GetTemplateSummaryResult>CapabilitiesReason"`
	Description        aws.StringValue        `xml:"GetTemplateSummaryResult>Description"`
	Parameters         []ParameterDeclaration `xml:"GetTemplateSummaryResult>Parameters>member"`
	Version            aws.StringValue        `xml:"GetTemplateSummaryResult>Version"`
}

// ListStackResourcesInput is undocumented.
type ListStackResourcesInput struct {
	NextToken aws.StringValue `xml:"NextToken"`
	StackName aws.StringValue `xml:"StackName"`
}

// ListStackResourcesOutput is undocumented.
type ListStackResourcesOutput struct {
	NextToken              aws.StringValue        `xml:"ListStackResourcesResult>NextToken"`
	StackResourceSummaries []StackResourceSummary `xml:"ListStackResourcesResult>StackResourceSummaries>member"`
}

// ListStacksInput is undocumented.
type ListStacksInput struct {
	NextToken         aws.StringValue `xml:"NextToken"`
	StackStatusFilter []string        `xml:"StackStatusFilter>member"`
}

// ListStacksOutput is undocumented.
type ListStacksOutput struct {
	NextToken      aws.StringValue `xml:"ListStacksResult>NextToken"`
	StackSummaries []StackSummary  `xml:"ListStacksResult>StackSummaries>member"`
}

// Output is undocumented.
type Output struct {
	Description aws.StringValue `xml:"Description"`
	OutputKey   aws.StringValue `xml:"OutputKey"`
	OutputValue aws.StringValue `xml:"OutputValue"`
}

// Parameter is undocumented.
type Parameter struct {
	ParameterKey     aws.StringValue  `xml:"ParameterKey"`
	ParameterValue   aws.StringValue  `xml:"ParameterValue"`
	UsePreviousValue aws.BooleanValue `xml:"UsePreviousValue"`
}

// ParameterDeclaration is undocumented.
type ParameterDeclaration struct {
	DefaultValue  aws.StringValue  `xml:"DefaultValue"`
	Description   aws.StringValue  `xml:"Description"`
	NoEcho        aws.BooleanValue `xml:"NoEcho"`
	ParameterKey  aws.StringValue  `xml:"ParameterKey"`
	ParameterType aws.StringValue  `xml:"ParameterType"`
}

// SetStackPolicyInput is undocumented.
type SetStackPolicyInput struct {
	StackName       aws.StringValue `xml:"StackName"`
	StackPolicyBody aws.StringValue `xml:"StackPolicyBody"`
	StackPolicyURL  aws.StringValue `xml:"StackPolicyURL"`
}

// SignalResourceInput is undocumented.
type SignalResourceInput struct {
	LogicalResourceID aws.StringValue `xml:"LogicalResourceId"`
	StackName         aws.StringValue `xml:"StackName"`
	Status            aws.StringValue `xml:"Status"`
	UniqueID          aws.StringValue `xml:"UniqueId"`
}

// Stack is undocumented.
type Stack struct {
	Capabilities      []string         `xml:"Capabilities>member"`
	CreationTime      time.Time        `xml:"CreationTime"`
	Description       aws.StringValue  `xml:"Description"`
	DisableRollback   aws.BooleanValue `xml:"DisableRollback"`
	LastUpdatedTime   time.Time        `xml:"LastUpdatedTime"`
	NotificationARNs  []string         `xml:"NotificationARNs>member"`
	Outputs           []Output         `xml:"Outputs>member"`
	Parameters        []Parameter      `xml:"Parameters>member"`
	StackID           aws.StringValue  `xml:"StackId"`
	StackName         aws.StringValue  `xml:"StackName"`
	StackStatus       aws.StringValue  `xml:"StackStatus"`
	StackStatusReason aws.StringValue  `xml:"StackStatusReason"`
	Tags              []Tag            `xml:"Tags>member"`
	TimeoutInMinutes  aws.IntegerValue `xml:"TimeoutInMinutes"`
}

// StackEvent is undocumented.
type StackEvent struct {
	EventID              aws.StringValue `xml:"EventId"`
	LogicalResourceID    aws.StringValue `xml:"LogicalResourceId"`
	PhysicalResourceID   aws.StringValue `xml:"PhysicalResourceId"`
	ResourceProperties   aws.StringValue `xml:"ResourceProperties"`
	ResourceStatus       aws.StringValue `xml:"ResourceStatus"`
	ResourceStatusReason aws.StringValue `xml:"ResourceStatusReason"`
	ResourceType         aws.StringValue `xml:"ResourceType"`
	StackID              aws.StringValue `xml:"StackId"`
	StackName            aws.StringValue `xml:"StackName"`
	Timestamp            time.Time       `xml:"Timestamp"`
}

// StackResource is undocumented.
type StackResource struct {
	Description          aws.StringValue `xml:"Description"`
	LogicalResourceID    aws.StringValue `xml:"LogicalResourceId"`
	PhysicalResourceID   aws.StringValue `xml:"PhysicalResourceId"`
	ResourceStatus       aws.StringValue `xml:"ResourceStatus"`
	ResourceStatusReason aws.StringValue `xml:"ResourceStatusReason"`
	ResourceType         aws.StringValue `xml:"ResourceType"`
	StackID              aws.StringValue `xml:"StackId"`
	StackName            aws.StringValue `xml:"StackName"`
	Timestamp            time.Time       `xml:"Timestamp"`
}

// StackResourceDetail is undocumented.
type StackResourceDetail struct {
	Description          aws.StringValue `xml:"Description"`
	LastUpdatedTimestamp time.Time       `xml:"LastUpdatedTimestamp"`
	LogicalResourceID    aws.StringValue `xml:"LogicalResourceId"`
	Metadata             aws.StringValue `xml:"Metadata"`
	PhysicalResourceID   aws.StringValue `xml:"PhysicalResourceId"`
	ResourceStatus       aws.StringValue `xml:"ResourceStatus"`
	ResourceStatusReason aws.StringValue `xml:"ResourceStatusReason"`
	ResourceType         aws.StringValue `xml:"ResourceType"`
	StackID              aws.StringValue `xml:"StackId"`
	StackName            aws.StringValue `xml:"StackName"`
}

// StackResourceSummary is undocumented.
type StackResourceSummary struct {
	LastUpdatedTimestamp time.Time       `xml:"LastUpdatedTimestamp"`
	LogicalResourceID    aws.StringValue `xml:"LogicalResourceId"`
	PhysicalResourceID   aws.StringValue `xml:"PhysicalResourceId"`
	ResourceStatus       aws.StringValue `xml:"ResourceStatus"`
	ResourceStatusReason aws.StringValue `xml:"ResourceStatusReason"`
	ResourceType         aws.StringValue `xml:"ResourceType"`
}

// StackSummary is undocumented.
type StackSummary struct {
	CreationTime        time.Time       `xml:"CreationTime"`
	DeletionTime        time.Time       `xml:"DeletionTime"`
	LastUpdatedTime     time.Time       `xml:"LastUpdatedTime"`
	StackID             aws.StringValue `xml:"StackId"`
	StackName           aws.StringValue `xml:"StackName"`
	StackStatus         aws.StringValue `xml:"StackStatus"`
	StackStatusReason   aws.StringValue `xml:"StackStatusReason"`
	TemplateDescription aws.StringValue `xml:"TemplateDescription"`
}

// Tag is undocumented.
type Tag struct {
	Key   aws.StringValue `xml:"Key"`
	Value aws.StringValue `xml:"Value"`
}

// TemplateParameter is undocumented.
type TemplateParameter struct {
	DefaultValue aws.StringValue  `xml:"DefaultValue"`
	Description  aws.StringValue  `xml:"Description"`
	NoEcho       aws.BooleanValue `xml:"NoEcho"`
	ParameterKey aws.StringValue  `xml:"ParameterKey"`
}

// UpdateStackInput is undocumented.
type UpdateStackInput struct {
	Capabilities                []string         `xml:"Capabilities>member"`
	NotificationARNs            []string         `xml:"NotificationARNs>member"`
	Parameters                  []Parameter      `xml:"Parameters>member"`
	StackName                   aws.StringValue  `xml:"StackName"`
	StackPolicyBody             aws.StringValue  `xml:"StackPolicyBody"`
	StackPolicyDuringUpdateBody aws.StringValue  `xml:"StackPolicyDuringUpdateBody"`
	StackPolicyDuringUpdateURL  aws.StringValue  `xml:"StackPolicyDuringUpdateURL"`
	StackPolicyURL              aws.StringValue  `xml:"StackPolicyURL"`
	TemplateBody                aws.StringValue  `xml:"TemplateBody"`
	TemplateURL                 aws.StringValue  `xml:"TemplateURL"`
	UsePreviousTemplate         aws.BooleanValue `xml:"UsePreviousTemplate"`
}

// UpdateStackOutput is undocumented.
type UpdateStackOutput struct {
	StackID aws.StringValue `xml:"UpdateStackResult>StackId"`
}

// ValidateTemplateInput is undocumented.
type ValidateTemplateInput struct {
	TemplateBody aws.StringValue `xml:"TemplateBody"`
	TemplateURL  aws.StringValue `xml:"TemplateURL"`
}

// ValidateTemplateOutput is undocumented.
type ValidateTemplateOutput struct {
	Capabilities       []string            `xml:"ValidateTemplateResult>Capabilities>member"`
	CapabilitiesReason aws.StringValue     `xml:"ValidateTemplateResult>CapabilitiesReason"`
	Description        aws.StringValue     `xml:"ValidateTemplateResult>Description"`
	Parameters         []TemplateParameter `xml:"ValidateTemplateResult>Parameters>member"`
}

// CreateStackResult is a wrapper for CreateStackOutput.
type CreateStackResult struct {
	StackID aws.StringValue `xml:"CreateStackResult>StackId"`
}

// DescribeStackEventsResult is a wrapper for DescribeStackEventsOutput.
type DescribeStackEventsResult struct {
	NextToken   aws.StringValue `xml:"DescribeStackEventsResult>NextToken"`
	StackEvents []StackEvent    `xml:"DescribeStackEventsResult>StackEvents>member"`
}

// DescribeStackResourceResult is a wrapper for DescribeStackResourceOutput.
type DescribeStackResourceResult struct {
	StackResourceDetail *StackResourceDetail `xml:"DescribeStackResourceResult>StackResourceDetail"`
}

// DescribeStackResourcesResult is a wrapper for DescribeStackResourcesOutput.
type DescribeStackResourcesResult struct {
	StackResources []StackResource `xml:"DescribeStackResourcesResult>StackResources>member"`
}

// DescribeStacksResult is a wrapper for DescribeStacksOutput.
type DescribeStacksResult struct {
	NextToken aws.StringValue `xml:"DescribeStacksResult>NextToken"`
	Stacks    []Stack         `xml:"DescribeStacksResult>Stacks>member"`
}

// EstimateTemplateCostResult is a wrapper for EstimateTemplateCostOutput.
type EstimateTemplateCostResult struct {
	URL aws.StringValue `xml:"EstimateTemplateCostResult>Url"`
}

// GetStackPolicyResult is a wrapper for GetStackPolicyOutput.
type GetStackPolicyResult struct {
	StackPolicyBody aws.StringValue `xml:"GetStackPolicyResult>StackPolicyBody"`
}

// GetTemplateResult is a wrapper for GetTemplateOutput.
type GetTemplateResult struct {
	TemplateBody aws.StringValue `xml:"GetTemplateResult>TemplateBody"`
}

// GetTemplateSummaryResult is a wrapper for GetTemplateSummaryOutput.
type GetTemplateSummaryResult struct {
	Capabilities       []string               `xml:"GetTemplateSummaryResult>Capabilities>member"`
	CapabilitiesReason aws.StringValue        `xml:"GetTemplateSummaryResult>CapabilitiesReason"`
	Description        aws.StringValue        `xml:"GetTemplateSummaryResult>Description"`
	Parameters         []ParameterDeclaration `xml:"GetTemplateSummaryResult>Parameters>member"`
	Version            aws.StringValue        `xml:"GetTemplateSummaryResult>Version"`
}

// ListStackResourcesResult is a wrapper for ListStackResourcesOutput.
type ListStackResourcesResult struct {
	NextToken              aws.StringValue        `xml:"ListStackResourcesResult>NextToken"`
	StackResourceSummaries []StackResourceSummary `xml:"ListStackResourcesResult>StackResourceSummaries>member"`
}

// ListStacksResult is a wrapper for ListStacksOutput.
type ListStacksResult struct {
	NextToken      aws.StringValue `xml:"ListStacksResult>NextToken"`
	StackSummaries []StackSummary  `xml:"ListStacksResult>StackSummaries>member"`
}

// UpdateStackResult is a wrapper for UpdateStackOutput.
type UpdateStackResult struct {
	StackID aws.StringValue `xml:"UpdateStackResult>StackId"`
}

// ValidateTemplateResult is a wrapper for ValidateTemplateOutput.
type ValidateTemplateResult struct {
	Capabilities       []string            `xml:"ValidateTemplateResult>Capabilities>member"`
	CapabilitiesReason aws.StringValue     `xml:"ValidateTemplateResult>CapabilitiesReason"`
	Description        aws.StringValue     `xml:"ValidateTemplateResult>Description"`
	Parameters         []TemplateParameter `xml:"ValidateTemplateResult>Parameters>member"`
}

// avoid errors if the packages aren't referenced
var _ time.Time