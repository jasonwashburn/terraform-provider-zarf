package provider

import (
	"context"
	"fmt"

	"terraform-provider-zarf/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &packageDataSource{}
	_ datasource.DataSourceWithConfigure = &packageDataSource{}
)

func NewPackageDataSource() datasource.DataSource {
	return &packageDataSource{}
}

type packageDataSource struct {
	client client.Client
}

type packageDataSourceModel struct {
	Source   types.String     `tfsdk:"source"`
	Metadata *packageMetadata `tfsdk:"metadata"`
	Digest   types.String     `tfsdk:"digest"`
}
type packageMetadata struct {
	Name         types.String `tfsdk:"name"`
	Description  types.String `tfsdk:"description"`
	Version      types.String `tfsdk:"version"`
	URL          types.String `tfsdk:"url"`
	Architecture types.String `tfsdk:"architecture"`
}

func (d *packageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_package"
}

func (d *packageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"source": schema.StringAttribute{
				Required:    true,
				Description: "The source of the package to inspect.",
			},
			"metadata": schema.SingleNestedAttribute{
				Computed:    true,
				Description: "The metadata of the package.",
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed:    true,
						Description: "The name of the package.",
					},
					"description": schema.StringAttribute{
						Computed:    true,
						Description: "The description of the package.",
					},
					"version": schema.StringAttribute{
						Computed:    true,
						Description: "The version of the package.",
					},
					"url": schema.StringAttribute{
						Computed:    true,
						Description: "The URL of the package.",
					},
					"architecture": schema.StringAttribute{
						Computed:    true,
						Description: "The architecture of the package.",
					},
				},
			},
			"digest": schema.StringAttribute{
				Computed:    true,
				Description: "The digest of the package.",
			},
		},
	}
}

func (d *packageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data packageDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	packageData, err := d.client.InspectPackage(ctx, data.Source.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Package",
			fmt.Sprintf("An unexpected error was encountered when reading the package: %s", err.Error()),
		)
		return
	}

	metadata := packageMetadata{
		Name:         types.StringValue(packageData.Metadata.Name),
		Description:  types.StringValue(packageData.Metadata.Description),
		Version:      types.StringValue(packageData.Metadata.Version),
		URL:          types.StringValue(packageData.Metadata.URL),
		Architecture: types.StringValue(packageData.Metadata.Architecture),
	}

	data.Digest = types.StringValue(packageData.Digest)
	data.Metadata = &metadata

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *packageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}
