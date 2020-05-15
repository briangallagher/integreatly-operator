package common

import (
	goctx "context"
	"github.com/integr8ly/integreatly-operator/pkg/apis-products/enmasse/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
	"time"
)

// TODO: Further enhancements: Make all the changes at the same time and just call one update on the CR
// Then verify all changes

// NOTES: We only need to test one CR per type
// We only need to check modification patterns and not all fields. Meaning, we check that a reserved field
// will be reconciled as apposed to checking all reserved fields.

const (
	RetryInterval = time.Second * 5
	Timeout       = time.Second * 60 * 5
	RetryCount    = 3
)

type CrWrapper interface {
	getCr() (runtime.Object, v1.ObjectMeta)
	modifyReservedField()
	verifyModifyOnReservedField() bool
	addNewField()
	verifyAddNewField() bool
	deleteReservedField()
	verifyDeleteReservedField() bool
}

// ************************************************************************************
// CR Specific type
// ************************************************************************************
type AddressPlanCrWrapper struct {
	// Add real CR as type
	cr *v1beta2.AddressPlan
}

func (apw AddressPlanCrWrapper) getCr() (runtime.Object, v1.ObjectMeta) {
	return apw.cr, apw.cr.ObjectMeta
}

func (apw AddressPlanCrWrapper) modifyReservedField() {
	apw.cr.Spec.AddressType = "Changed-Valued"
}

// Add a field not under the control of the reconciler
func (apw AddressPlanCrWrapper) addNewField() {
	apw.cr.Spec.Partitions = 1234
}

func (apw AddressPlanCrWrapper) verifyAddNewField() bool {
	return apw.cr.Spec.Partitions == 1234
}

// TODO: is this the same as modify? Can not set a string to nil
func (apw AddressPlanCrWrapper) deleteReservedField() {
	apw.cr.Spec.DisplayName = ""
}

func (apw AddressPlanCrWrapper) verifyDeleteReservedField() bool {
	return apw.cr.Spec.DisplayName != ""
}

func (apw AddressPlanCrWrapper) verifyModifyOnReservedField() bool {
	// NOTE: Only checking one field
	return apw.cr.Spec.AddressType != "Changed-Valued"
}

func getAddressSpacePlanCr(ctx *TestingContext) AddressPlanCrWrapper {
	addressPlan := &v1beta2.AddressPlan{}
	// NOTE: only using one specific CR
	ctx.Client.Get(goctx.TODO(), types.NamespacedName{Name: "brokered-queue", Namespace: "redhat-rhmi-amq-online"}, addressPlan)
	return AddressPlanCrWrapper{cr: addressPlan}
}

// ************************************************************************************
// Generic Functions
// ************************************************************************************

func TestAMQOnlineCrs(t *testing.T, ctx *TestingContext) {
	// TODO: test using go routines
	//go testCrUpdates(t, ctx, getAddressSpacePlanCr(ctx))
	testCrUpdates(t, ctx, getAddressSpacePlanCr(ctx))

	// TODO:
	//go testCrUpdates(t, ctx, getAuthenticationServiceCr(ctx))
	//go testCrUpdates(t, ctx, getBrokeredInfraConfigCr(ctx))
	// etc
}

// Resource Version is on v1.ObjectMeta so can be generic method
func getResourceVersion(crWrapper CrWrapper) string {
	_, meta := crWrapper.getCr()
	return meta.GetResourceVersion()
}

func testCrUpdates(t *testing.T, ctx *TestingContext, crWrapper CrWrapper) {
	testModifyReservedField(t, ctx, crWrapper)
	addNewCRValues(t, ctx, crWrapper)
	testDeleteReservedField(t, ctx, crWrapper)
}

func addNewCRValues(t *testing.T, ctx *TestingContext, crWrapper CrWrapper) {
	crWrapper.addNewField()
	resourceVersion := getResourceVersion(crWrapper)
	updateCr(t, ctx, crWrapper)
	verifyUpdatedCr(t, ctx, crWrapper, crWrapper.verifyModifyOnReservedField, resourceVersion, RetryCount)
}

func testModifyReservedField(t *testing.T, ctx *TestingContext, crWrapper CrWrapper) {
	crWrapper.modifyReservedField()
	resourceVersion := getResourceVersion(crWrapper)
	updateCr(t, ctx, crWrapper)
	verifyUpdatedCr(t, ctx, crWrapper, crWrapper.verifyModifyOnReservedField, resourceVersion, RetryCount)
}

func testDeleteReservedField(t *testing.T, ctx *TestingContext, crWrapper CrWrapper) {
	crWrapper.deleteReservedField()
	resourceVersion := getResourceVersion(crWrapper)
	updateCr(t, ctx, crWrapper)
	verifyUpdatedCr(t, ctx, crWrapper, crWrapper.verifyModifyOnReservedField, resourceVersion, RetryCount)
}

func verifyUpdatedCr(t *testing.T, ctx *TestingContext, crWrapper CrWrapper, verify func() bool, resourceVersion string, retry int) error {

	err := getUpdatedResourceVersion(t, ctx, crWrapper, resourceVersion)
	if err != nil && retry > 0 {
		// recursive call
		verifyUpdatedCr(t, ctx, crWrapper, verify, resourceVersion, retry-1)
	} else if err != nil {
		t.Fatal("Failed to Verify ...... TODO:::")
	} else {
		if verify() {
			return nil // success
		} else if retry > 0 {
			verifyUpdatedCr(t, ctx, crWrapper, verify, resourceVersion, retry-1)
		} else {
			t.Fatal("Ran out of retries ...... TODO:::")
		}
	}
	return nil
}

func getUpdatedResourceVersion(t *testing.T, ctx *TestingContext, crWrapper CrWrapper, resourceVersion string) error {
	cr, crMeta := crWrapper.getCr()
	return wait.Poll(RetryInterval, Timeout, func() (done bool, err error) {
		err = ctx.Client.Get(goctx.TODO(), k8sclient.ObjectKey{Name: crMeta.GetName(), Namespace: crMeta.GetNamespace()}, cr)
		if err != nil {
			t.Fatalf("%s : Fail to get the cr", err)
		}
		if resourceVersion != crMeta.GetResourceVersion() {
			return true, nil
		}
		return false, nil
	})
}

func updateCr(t *testing.T, ctx *TestingContext, crWrapper CrWrapper) error {
	// TODO: can we shortcut this?
	cr, _ := crWrapper.getCr()
	err := ctx.Client.Update(goctx.TODO(), cr)
	if err != nil {
		t.Fatalf("%s : Fail to update the cr", err)
	}
	return nil
}
