package config

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime"

	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
)

type ContainerSecurity struct {
	config ProductConfig
}

func NewContainerSecurity(config ProductConfig) *ContainerSecurity {
	return &ContainerSecurity{config: config}
}

func (a *ContainerSecurity) GetWatchableCRDs() []runtime.Object {
	// TODO:
	return []runtime.Object{}
	//return []runtime.Object{
	//	&enmassev1beta2.AddressPlan{
	//		TypeMeta: metav1.TypeMeta{
	//			Kind:       "AddressPlan",
	//			APIVersion: enmassev1beta2.SchemeGroupVersion.String(),
	//		},
	//	},
	//	&enmassev1beta2.AddressSpacePlan{
	//		TypeMeta: metav1.TypeMeta{
	//			Kind:       "AddressSpacePlan",
	//			APIVersion: enmassev1beta2.SchemeGroupVersion.String(),
	//		},
	//	},
	//	&enmassev1beta1.BrokeredInfraConfig{
	//		TypeMeta: metav1.TypeMeta{
	//			Kind:       "BrokeredInfraConfig",
	//			APIVersion: enmassev1beta1.SchemeGroupVersion.String(),
	//		},
	//	},
	//	&enmasseadminv1beta1.AuthenticationService{
	//		TypeMeta: metav1.TypeMeta{
	//			Kind:       "AuthenticationService",
	//			APIVersion: enmasseadminv1beta1.SchemeGroupVersion.String(),
	//		},
	//	},
	//}
}

func (a *ContainerSecurity) GetHost() string {
	return a.config["HOST"]
}

func (a *ContainerSecurity) SetHost(newHost string) {
	a.config["HOST"] = newHost
}

func (a *ContainerSecurity) GetNamespace() string {
	return a.config["NAMESPACE"]
}

func (a *ContainerSecurity) GetOperatorNamespace() string {
	return a.config["OPERATOR_NAMESPACE"]
}

func (a *ContainerSecurity) SetOperatorNamespace(newNamespace string) {
	a.config["OPERATOR_NAMESPACE"] = newNamespace
}

func (a *ContainerSecurity) GetLabelSelector() string {
	return "middleware"
}

func (a *ContainerSecurity) SetNamespace(newNamespace string) {
	a.config["NAMESPACE"] = newNamespace
}

func (a *ContainerSecurity) Read() ProductConfig {
	return a.config
}

func (a *ContainerSecurity) GetProductName() integreatlyv1alpha1.ProductName {
	return integreatlyv1alpha1.ProductContainerSecurity
}

func (a *ContainerSecurity) GetProductVersion() integreatlyv1alpha1.ProductVersion {
	return integreatlyv1alpha1.VersionConatinerSecurity
}

func (a *ContainerSecurity) GetOperatorVersion() integreatlyv1alpha1.OperatorVersion {
	return integreatlyv1alpha1.OperatorVersionContainerSecurity
}

func (a *ContainerSecurity) Validate() error {
	if a.GetNamespace() == "" {
		return errors.New("config namespace is not defined")
	}
	if a.GetProductName() == "" {
		return errors.New("config product name is not defined")
	}
	if a.GetHost() == "" {
		return errors.New("config host is not defined")
	}
	return nil
}
