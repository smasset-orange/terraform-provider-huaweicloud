// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product CC
// ---------------------------------------------------------------

package cc

import (
	"context"
	"regexp"
	"strings"

	"github.com/chnsz/golangsdk"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/jmespath/go-jmespath"
)

func ResourceNetworkInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkInstanceCreate,
		UpdateContext: resourceNetworkInstanceUpdate,
		ReadContext:   resourceNetworkInstanceRead,
		DeleteContext: resourceNetworkInstanceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Type of the network instance to be loaded to the cloud connection.`,
				ValidateFunc: validation.StringInSlice([]string{
					"vpc", "vgw",
				}, false),
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the VPC or virtual gateway to be loaded to the cloud connection.`,
			},
			"cidrs": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				Description: `List of routes advertised by the VPC or virtual gateway.`,
			},
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Project ID of the VPC or virtual gateway.`,
			},
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Region ID of the VPC or virtual gateway.`,
			},
			"cloud_connection_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Cloud connection ID.`,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The network instance name.`,
				ValidateFunc: validation.All(
					validation.StringMatch(regexp.MustCompile(`^[\x{4E00}-\x{9FFC}A-Za-z-_0-9.]*$`),
						"the input is invalid"),
					validation.StringLenBetween(1, 64),
				),
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The description about the network instance.`,
				ValidateFunc: validation.All(
					validation.StringMatch(regexp.MustCompile(`^[^<>]+$`),
						"the input is invalid"),
					validation.StringLenBetween(0, 255),
				),
			},
			"instance_domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Account ID of the VPC or virtual gateway.`,
			},
			"domain_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Account ID.`,
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Network instance status.`,
			},
		},
	}
}

func resourceNetworkInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	// createNetworkInstance: create a Cloud Connect.
	var (
		createNetworkInstanceHttpUrl = "v3/{domain_id}/ccaas/network-instances"
		createNetworkInstanceProduct = "cc"
	)
	createNetworkInstanceClient, err := config.NewServiceClient(createNetworkInstanceProduct, region)
	if err != nil {
		return diag.Errorf("error creating NetworkInstance Client: %s", err)
	}

	createNetworkInstancePath := createNetworkInstanceClient.Endpoint + createNetworkInstanceHttpUrl
	createNetworkInstancePath = strings.Replace(createNetworkInstancePath, "{domain_id}", config.DomainID, -1)

	createNetworkInstanceOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			201,
		},
	}
	createNetworkInstanceOpt.JSONBody = utils.RemoveNil(buildCreateNetworkInstanceBodyParams(d, config))
	createNetworkInstanceResp, err := createNetworkInstanceClient.Request("POST", createNetworkInstancePath, &createNetworkInstanceOpt)
	if err != nil {
		return diag.Errorf("error creating NetworkInstance: %s", err)
	}

	createNetworkInstanceRespBody, err := utils.FlattenResponse(createNetworkInstanceResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := jmespath.Search("network_instance.id", createNetworkInstanceRespBody)
	if err != nil {
		return diag.Errorf("error creating NetworkInstance: ID is not found in API response")
	}
	d.SetId(id.(string))

	return resourceNetworkInstanceRead(ctx, d, meta)
}

func buildCreateNetworkInstanceBodyParams(d *schema.ResourceData, config *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"network_instance": buildCreateNetworkInstanceNetworkInstanceChildBody(d),
	}
	return bodyParams
}

func buildCreateNetworkInstanceNetworkInstanceChildBody(d *schema.ResourceData) map[string]interface{} {
	params := map[string]interface{}{
		"name":                utils.ValueIngoreEmpty(d.Get("name")),
		"description":         utils.ValueIngoreEmpty(d.Get("description")),
		"type":                utils.ValueIngoreEmpty(d.Get("type")),
		"instance_id":         utils.ValueIngoreEmpty(d.Get("instance_id")),
		"instance_domain_id":  utils.ValueIngoreEmpty(d.Get("instance_domain_id")),
		"project_id":          utils.ValueIngoreEmpty(d.Get("project_id")),
		"region_id":           utils.ValueIngoreEmpty(d.Get("region_id")),
		"cloud_connection_id": utils.ValueIngoreEmpty(d.Get("cloud_connection_id")),
		"cidrs":               utils.ValueIngoreEmpty(d.Get("cidrs")),
	}
	return params
}

func resourceNetworkInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	var mErr *multierror.Error

	// getNetworkInstance: Query the Network instance
	var (
		getNetworkInstanceHttpUrl = "v3/{domain_id}/ccaas/network-instances/{id}"
		getNetworkInstanceProduct = "cc"
	)
	getNetworkInstanceClient, err := config.NewServiceClient(getNetworkInstanceProduct, region)
	if err != nil {
		return diag.Errorf("error creating NetworkInstance Client: %s", err)
	}

	getNetworkInstancePath := getNetworkInstanceClient.Endpoint + getNetworkInstanceHttpUrl
	getNetworkInstancePath = strings.Replace(getNetworkInstancePath, "{domain_id}", config.DomainID, -1)
	getNetworkInstancePath = strings.Replace(getNetworkInstancePath, "{id}", d.Id(), -1)

	getNetworkInstanceOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	getNetworkInstanceResp, err := getNetworkInstanceClient.Request("GET", getNetworkInstancePath, &getNetworkInstanceOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving NetworkInstance")
	}

	getNetworkInstanceRespBody, err := utils.FlattenResponse(getNetworkInstanceResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("name", utils.PathSearch("network_instance.name", getNetworkInstanceRespBody, nil)),
		d.Set("description", utils.PathSearch("network_instance.description", getNetworkInstanceRespBody, nil)),
		d.Set("domain_id", utils.PathSearch("network_instance.domain_id", getNetworkInstanceRespBody, nil)),
		d.Set("status", utils.PathSearch("network_instance.status", getNetworkInstanceRespBody, nil)),
		d.Set("type", utils.PathSearch("network_instance.type", getNetworkInstanceRespBody, nil)),
		d.Set("cloud_connection_id", utils.PathSearch("network_instance.cloud_connection_id", getNetworkInstanceRespBody, nil)),
		d.Set("instance_id", utils.PathSearch("network_instance.instance_id", getNetworkInstanceRespBody, nil)),
		d.Set("instance_domain_id", utils.PathSearch("network_instance.instance_domain_id", getNetworkInstanceRespBody, nil)),
		d.Set("region_id", utils.PathSearch("network_instance.region_id", getNetworkInstanceRespBody, nil)),
		d.Set("project_id", utils.PathSearch("network_instance.project_id", getNetworkInstanceRespBody, nil)),
		d.Set("cidrs", utils.PathSearch("network_instance.cidrs", getNetworkInstanceRespBody, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func resourceNetworkInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	updateNetworkInstancehasChanges := []string{
		"name",
		"description",
		"cidrs",
	}

	if d.HasChanges(updateNetworkInstancehasChanges...) {
		// updateNetworkInstance: update the Network instance
		var (
			updateNetworkInstanceHttpUrl = "v3/{domain_id}/ccaas/network-instances/{id}"
			updateNetworkInstanceProduct = "cc"
		)
		updateNetworkInstanceClient, err := config.NewServiceClient(updateNetworkInstanceProduct, region)
		if err != nil {
			return diag.Errorf("error creating NetworkInstance Client: %s", err)
		}

		updateNetworkInstancePath := updateNetworkInstanceClient.Endpoint + updateNetworkInstanceHttpUrl
		updateNetworkInstancePath = strings.Replace(updateNetworkInstancePath, "{domain_id}", config.DomainID, -1)
		updateNetworkInstancePath = strings.Replace(updateNetworkInstancePath, "{id}", d.Id(), -1)

		updateNetworkInstanceOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			OkCodes: []int{
				200,
			},
		}
		updateNetworkInstanceOpt.JSONBody = utils.RemoveNil(buildUpdateNetworkInstanceBodyParams(d, config))
		_, err = updateNetworkInstanceClient.Request("PUT", updateNetworkInstancePath, &updateNetworkInstanceOpt)
		if err != nil {
			return diag.Errorf("error updating NetworkInstance: %s", err)
		}
	}
	return resourceNetworkInstanceRead(ctx, d, meta)
}

func buildUpdateNetworkInstanceBodyParams(d *schema.ResourceData, config *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"network_instance": buildUpdateNetworkInstanceNetworkInstanceChildBody(d),
	}
	return bodyParams
}

func buildUpdateNetworkInstanceNetworkInstanceChildBody(d *schema.ResourceData) map[string]interface{} {
	params := map[string]interface{}{
		"name":        utils.ValueIngoreEmpty(d.Get("name")),
		"description": utils.ValueIngoreEmpty(d.Get("description")),
		"cidrs":       utils.ValueIngoreEmpty(d.Get("cidrs")),
	}
	return params
}

func resourceNetworkInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	// deleteNetworkInstance: missing operation notes
	var (
		deleteNetworkInstanceHttpUrl = "v3/{domain_id}/ccaas/network-instances/{id}"
		deleteNetworkInstanceProduct = "cc"
	)
	deleteNetworkInstanceClient, err := config.NewServiceClient(deleteNetworkInstanceProduct, region)
	if err != nil {
		return diag.Errorf("error creating NetworkInstance Client: %s", err)
	}

	deleteNetworkInstancePath := deleteNetworkInstanceClient.Endpoint + deleteNetworkInstanceHttpUrl
	deleteNetworkInstancePath = strings.Replace(deleteNetworkInstancePath, "{domain_id}", config.DomainID, -1)
	deleteNetworkInstancePath = strings.Replace(deleteNetworkInstancePath, "{id}", d.Id(), -1)

	deleteNetworkInstanceOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			204,
		},
	}
	_, err = deleteNetworkInstanceClient.Request("DELETE", deleteNetworkInstancePath, &deleteNetworkInstanceOpt)
	if err != nil {
		return diag.Errorf("error deleting NetworkInstance: %s", err)
	}

	return nil
}
