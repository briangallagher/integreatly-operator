// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.BrokeredInfraConfig": schema_pkg_apis_enmasse_v1beta1_BrokeredInfraConfig(ref),
		"github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.StandardInfraConfig": schema_pkg_apis_enmasse_v1beta1_StandardInfraConfig(ref),
	}
}

func schema_pkg_apis_enmasse_v1beta1_BrokeredInfraConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BrokeredInfraConfig is the Schema for the brokeredinfraconfigs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.BrokeredInfraConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.BrokeredInfraConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.BrokeredInfraConfigSpec", "github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.BrokeredInfraConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_enmasse_v1beta1_StandardInfraConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StandardInfraConfig is the Schema for the standardinfraconfigs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.StandardInfraConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.StandardInfraConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.StandardInfraConfigSpec", "github.com/briangallagher/integreatly-operator/pkg/apis/enmasse/v1beta1.StandardInfraConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}
