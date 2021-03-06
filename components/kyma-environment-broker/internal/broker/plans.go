package broker

import (
	"encoding/json"

	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/runtime/components"

	"github.com/pivotal-cf/brokerapi/v7/domain"
)

const (
	AllPlansSelector = "all_plans"

	GCPPlanID         = "ca6e5357-707f-4565-bbbd-b3ab732597c6"
	GCPPlanName       = "gcp"
	AzurePlanID       = "4deee563-e5ec-4731-b9b1-53b42d855f0c"
	AzurePlanName     = "azure"
	AzureLitePlanID   = "8cb22518-aa26-44c5-91a0-e669ec9bf443"
	AzureLitePlanName = "azure_lite"
	TrialPlanID       = "7d55d31d-35ae-4438-bf13-6ffdfa107d9f"
	TrialPlanName     = "trial"
)

var PlanNamesMapping = map[string]string{
	GCPPlanID:       GCPPlanName,
	AzurePlanID:     AzurePlanName,
	AzureLitePlanID: AzureLitePlanName,
	TrialPlanID:     TrialPlanName,
}

var PlanIDsMapping = map[string]string{
	AzurePlanName:     AzurePlanID,
	AzureLitePlanName: AzureLitePlanID,
	GCPPlanName:       GCPPlanID,
	TrialPlanName:     TrialPlanID,
}

type TrialCloudRegion string

const (
	Europe TrialCloudRegion = "europe"
	Us     TrialCloudRegion = "us"
	Asia   TrialCloudRegion = "asia"
)

func AzureRegions() []string {
	return []string{
		"centralus",
		"eastus",
		"westus2",
		"northeurope",
		"uksouth",
		"japaneast",
		"southeastasia",
		"westeurope",
	}
}

type Type struct {
	Type            string        `json:"type"`
	Minimum         int           `json:"minimum,omitempty"`
	Enum            []interface{} `json:"enum,omitempty"`
	Items           []Type        `json:"items,omitempty"`
	AdditionalItems *bool         `json:"additionalItems,omitempty"`
	UniqueItems     *bool         `json:"uniqueItems,omitempty"`
}

type RootSchema struct {
	Schema string `json:"$schema"`
	Type
	Properties interface{} `json:"properties"`
	Required   []string    `json:"required"`
}

type ProvisioningProperties struct {
	Components     Type `json:"components"`
	Name           Type `json:"name"`
	DiskType       Type `json:"diskType"`
	VolumeSizeGb   Type `json:"volumeSizeGb"`
	MachineType    Type `json:"machineType"`
	Region         Type `json:"region"`
	Zones          Type `json:"zones"`
	AutoScalerMin  Type `json:"autoScalerMin"`
	AutoScalerMax  Type `json:"autoScalerMax"`
	MaxSurge       Type `json:"maxSurge"`
	MaxUnavailable Type `json:"maxUnavailable"`
}

func GCPSchema(machineTypes []string) []byte {
	f := new(bool)
	*f = false
	t := new(bool)
	*t = true

	rs := RootSchema{
		Schema: "http://json-schema.org/draft-04/schema#",
		Type: Type{
			Type: "object",
		},
		Properties: ProvisioningProperties{
			Components: Type{
				Type: "array",
				Items: []Type{{
					Type: "string",
					Enum: ToInterfaceSlice([]string{components.Kiali, components.Tracing}),
				}},
				AdditionalItems: f,
				UniqueItems:     t,
			},
			Name: Type{
				Type: "string",
			},
			DiskType: Type{Type: "string"},
			VolumeSizeGb: Type{
				Type: "integer",
			},
			MachineType: Type{
				Type: "string",
				Enum: ToInterfaceSlice(machineTypes),
			},
			Region: Type{
				Type: "string",
				Enum: ToInterfaceSlice([]string{
					"asia-south1", "asia-southeast1",
					"asia-east2", "asia-east1",
					"asia-northeast1", "asia-northeast2", "asia-northeast-3",
					"australia-southeast1",
					"europe-west2", "europe-west4", "europe-west5", "europe-west6", "europe-west3",
					"europe-north1",
					"us-west1", "us-west2", "us-west3",
					"us-central1",
					"us-east4",
					"northamerica-northeast1", "southamerica-east1"}),
			},
			Zones: Type{
				Type: "array",
				Items: []Type{{
					Type: "string",
					Enum: ToInterfaceSlice([]string{
						"asia-south1-a", "asia-south1-b", "asia-south1-c",
						"asia-southeast1-a", "asia-southeast1-b", "asia-southeast1-c",
						"asia-east2-a", "asia-east2-b", "asia-east2-c",
						"asia-east1-a", "asia-east1-b", "asia-east1-c",
						"asia-northeast1-a", "asia-northeast1-b", "asia-northeast1-c",
						"asia-northeast2-a", "asia-northeast2-b", "asia-northeast2-c",
						"asia-northeast-3-a", "asia-northeast-3-b", "asia-northeast-3-c",
						"australia-southeast1-a", "australia-southeast1-b", "australia-southeast1-c",
						"europe-west2-a", "europe-west2-b", "europe-west2-c",
						"europe-west4-a", "europe-west4-b", "europe-west4-c",
						"europe-west5-a", "europe-west5-b", "europe-west5-c",
						"europe-west6-a", "europe-west6-b", "europe-west6-c",
						"europe-west3-a", "europe-west3-b", "europe-west3-c",
						"europe-north1-a", "europe-north1-b", "europe-north1-c",
						"us-west1-a", "us-west1-b", "us-west1-c",
						"us-west2-a", "us-west2-b", "us-west2-c",
						"us-west3-a", "us-west3-b", "us-west3-c",
						"us-central1-a", "us-central1-b", "us-central1-c",
						"us-east4-a", "us-east4-b", "us-east4-c",
						"northamerica-northeast1-a", "northamerica-northeast1-b", "northamerica-northeast1-c",
						"southamerica-east1-a", "southamerica-east1-b", "southamerica-east1-c"}),
				}},
			},
			AutoScalerMin: Type{
				Type: "integer",
			},
			AutoScalerMax: Type{
				Type: "integer",
			},
			MaxSurge: Type{
				Type: "integer",
			},
			MaxUnavailable: Type{
				Type: "integer",
			},
		},
		Required: []string{"name"},
	}

	bytes, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	return bytes
}

func AzureSchema(machineTypes []string) []byte {
	f := new(bool)
	*f = false
	t := new(bool)
	*t = true
	rs := RootSchema{
		Schema: "http://json-schema.org/draft-04/schema#",
		Type: Type{
			Type: "object",
		},
		Properties: ProvisioningProperties{
			Components: Type{
				Type: "array",
				Items: []Type{{
					Type: "string",
					Enum: ToInterfaceSlice([]string{components.Kiali, components.Tracing}),
				}},
				AdditionalItems: f,
				UniqueItems:     t,
			},
			Name: Type{
				Type: "string",
			},
			DiskType: Type{Type: "string"},
			VolumeSizeGb: Type{
				Type:    "integer",
				Minimum: 50,
			},
			MachineType: Type{
				Type: "string",
				Enum: ToInterfaceSlice(machineTypes),
			},
			Region: Type{
				Type: "string",
				Enum: ToInterfaceSlice(AzureRegions()),
			},
			Zones: Type{
				Type: "array",
				Items: []Type{{
					Type: "string",
					//TODO: add enum for zones
				}},
			},
			AutoScalerMin: Type{
				Type: "integer",
			},
			AutoScalerMax: Type{
				Type: "integer",
			},
			MaxSurge: Type{
				Type: "integer",
			},
			MaxUnavailable: Type{
				Type: "integer",
			},
		},
		Required: []string{"name"},
	}

	bytes, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	return bytes
}

func TrialSchema() []byte {
	schema := `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    },
    "region": {
      "type": "string",
      "enum": [
        "europe",
        "us"
      ]
    },
    "provider": {
      "type": "string",
      "enum": [
        "Azure",
        "GCP"
      ]
    }
  },
  "required": [
    "name"
  ]
}`

	bytes := []byte(schema)
	return bytes
}

func ToInterfaceSlice(input []string) []interface{} {
	interfaces := make([]interface{}, len(input))
	for i, item := range input {
		interfaces[i] = item
	}
	return interfaces
}

// plans is designed to hold plan defaulting logic
// keep internal/hyperscaler/azure/config.go in sync with any changes to available zones
var Plans = map[string]struct {
	PlanDefinition        domain.ServicePlan
	provisioningRawSchema []byte
}{
	GCPPlanID: {
		PlanDefinition: domain.ServicePlan{
			ID:          GCPPlanID,
			Name:        GCPPlanName,
			Description: "GCP",
			Metadata: &domain.ServicePlanMetadata{
				DisplayName: "GCP",
			},
			Schemas: &domain.ServiceSchemas{
				Instance: domain.ServiceInstanceSchema{
					Create: domain.Schema{
						Parameters: make(map[string]interface{}),
					},
				},
			},
		},
		provisioningRawSchema: GCPSchema([]string{"n1-standard-2", "n1-standard-4", "n1-standard-8", "n1-standard-16", "n1-standard-32", "n1-standard-64"}),
	},
	AzurePlanID: {
		PlanDefinition: domain.ServicePlan{
			ID:          AzurePlanID,
			Name:        AzurePlanName,
			Description: "Azure",
			Metadata: &domain.ServicePlanMetadata{
				DisplayName: "Azure",
			},
			Schemas: &domain.ServiceSchemas{
				Instance: domain.ServiceInstanceSchema{
					Create: domain.Schema{
						Parameters: make(map[string]interface{}),
					},
				},
			},
		},
		provisioningRawSchema: AzureSchema([]string{"Standard_D8_v3"}),
	},
	AzureLitePlanID: {
		PlanDefinition: domain.ServicePlan{
			ID:          AzureLitePlanID,
			Name:        AzureLitePlanName,
			Description: "Azure Lite",
			Metadata: &domain.ServicePlanMetadata{
				DisplayName: "Azure Lite",
			},
			Schemas: &domain.ServiceSchemas{
				Instance: domain.ServiceInstanceSchema{
					Create: domain.Schema{
						Parameters: make(map[string]interface{}),
					},
				},
			},
		},
		provisioningRawSchema: AzureSchema([]string{"Standard_D4_v3"}),
	},
	TrialPlanID: {
		PlanDefinition: domain.ServicePlan{
			ID:          TrialPlanID,
			Name:        TrialPlanName,
			Description: "Trial",
			Metadata: &domain.ServicePlanMetadata{
				DisplayName: "Trial",
			},
			Schemas: &domain.ServiceSchemas{
				Instance: domain.ServiceInstanceSchema{
					Create: domain.Schema{
						Parameters: make(map[string]interface{}),
					},
				},
			},
		},
		provisioningRawSchema: TrialSchema(),
	},
}

func IsTrialPlan(planID string) bool {
	switch planID {
	case TrialPlanID:
		return true
	default:
		return false
	}
}
