package common

import (
	goctx "context"
	"github.com/integr8ly/integreatly-operator/pkg/apis-products/enmasse/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sync"
	"testing"
	"time"
)

const (
	RetryInterval = time.Second * 5
	Timeout       = time.Second * 60 * 5
	RetryCount    = 3
)

type CrWrapper interface {
	getCr() (runtime.Object, v1.ObjectMeta)
	modifyReservedField()
	verifyModifyOnReservedField() bool
	//verifyModifyOnOpenField() bool			// TODO:
	//verifyDeleteOnReservedField() bool		// TODO:
	//verifyDeleteOnOpenField() bool			// TODO:
	//verifyAddField() bool						// TODO:
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

func (apw AddressPlanCrWrapper) verifyModifyOnReservedField() bool {
	// Only checking one field
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

func testAMQOnlineCrs(t *testing.T, ctx *TestingContext, wg *sync.WaitGroup) {
	testCrUpdates(t, ctx, getAddressSpacePlanCr(ctx))

	//testAddressPlan(wg, t, ctx)
	//testAuthenticationServiceCr(wg, t, ctx)
	//testBrokeredInfraConfigCr(wg, t, ctx)
	//testStandardInfraConfigCr(wg, t, ctx)
	//testRoleCr(wg, t, ctx)
	//testRoleBindingCr(wg, t, ctx)
}

// Resource Version is on v1.ObjectMeta so can be generic method
func getResourceVersion(crWrapper CrWrapper) string {
	_, meta := crWrapper.getCr()
	return meta.GetResourceVersion()
}

func testCrUpdates(t *testing.T, ctx *TestingContext, crWrapper CrWrapper) {
	testModifyReservedField(t, ctx, crWrapper)
	// testModifyOpenField(t, ctx, crWrapper)  		// TODO:
	// testDeleteReservedField(t, ctx, crWrapper) 	// TODO:
	// testDeleteOpenField(t, ctx, crWrapper) 		// TODO:
	// addNewCRValues(t, ctx, rt, crData)			// TODO:
}

func testModifyReservedField(t *testing.T, ctx *TestingContext, crWrapper CrWrapper) {
	crWrapper.modifyReservedField()
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
