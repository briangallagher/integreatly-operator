package containersecurity

import (
	"context"
	"fmt"

	integreatlyv1alpha1 "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/integr8ly/integreatly-operator/pkg/config"
	"github.com/integr8ly/integreatly-operator/pkg/resources"
	"github.com/integr8ly/integreatly-operator/pkg/resources/backup"
	"github.com/integr8ly/integreatly-operator/pkg/resources/constants"
	"github.com/integr8ly/integreatly-operator/pkg/resources/events"
	l "github.com/integr8ly/integreatly-operator/pkg/resources/logger"
	"github.com/integr8ly/integreatly-operator/pkg/resources/marketplace"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	//cs "github.com/quay/container-security-operator/apis/secscan/v1alpha1"
)

const (
	defaultInstallationNamespace = "container-security"
	manifestPackage              = "integreatly-container-security"
)

type Reconciler struct {
	*resources.Reconciler
	ConfigManager config.ConfigReadWriter
	Config        *config.ContainerSecurity
	installation  *integreatlyv1alpha1.RHMI
	mpm           marketplace.MarketplaceInterface
	log           l.Logger
	extraParams   map[string]string
	recorder      record.EventRecorder
}

func (r *Reconciler) GetPreflightObject(ns string) runtime.Object {
	return nil
}

func (r *Reconciler) VerifyVersion(installation *integreatlyv1alpha1.RHMI) bool {
	// TODO: product image is currently not versioned. see csv file
	return true
	//return version.VerifyProductAndOperatorVersion(
	//	installation.Status.Stages[integreatlyv1alpha1.ProductsStage].Products[integreatlyv1alpha1.ProductGrafana],
	//	string(integreatlyv1alpha1.VersionGrafana),
	//	string(integreatlyv1alpha1.OperatorVersionGrafana),
	//)
}

func NewReconciler(configManager config.ConfigReadWriter, installation *integreatlyv1alpha1.RHMI, mpm marketplace.MarketplaceInterface, recorder record.EventRecorder, logger l.Logger) (*Reconciler, error) {
	ns := installation.Spec.NamespacePrefix + defaultInstallationNamespace
	config, err := configManager.ReadContainerSecurity()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve container security config: %w", err)
	}
	config.SetNamespace(ns + "-operator")
	config.SetOperatorNamespace(ns + "-operator")
	configManager.WriteConfig(config)

	return &Reconciler{
		ConfigManager: configManager,
		Config:        config,
		installation:  installation,
		mpm:           mpm,
		log:           logger,
		Reconciler:    resources.NewReconciler(mpm),
		recorder:      recorder,
	}, nil
}

func (r *Reconciler) Reconcile(ctx context.Context, installation *integreatlyv1alpha1.RHMI, product *integreatlyv1alpha1.RHMIProductStatus, client k8sclient.Client) (integreatlyv1alpha1.StatusPhase, error) {
	r.log.Info("Start Container Security reconcile")

	operatorNamespace := r.Config.GetOperatorNamespace()
	productNamespace := r.Config.GetOperatorNamespace()

	phase, err := r.ReconcileFinalizer(ctx, client, installation, string(r.Config.GetProductName()), func() (integreatlyv1alpha1.StatusPhase, error) {
		phase, err := resources.RemoveNamespace(ctx, installation, client, productNamespace, r.log)
		if err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
			return phase, err
		}

		phase, err = resources.RemoveNamespace(ctx, installation, client, operatorNamespace, r.log)
		if err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
			return phase, err
		}

		return integreatlyv1alpha1.PhaseCompleted, nil
	}, r.log)
	if err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
		events.HandleError(r.recorder, installation, phase, "Failed to reconcile finalizer", err)
		return phase, err
	}

	phase, err = r.ReconcileNamespace(ctx, operatorNamespace, installation, client, r.log)
	if err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
		events.HandleError(r.recorder, installation, phase, fmt.Sprintf("Failed to reconcile %s ns", operatorNamespace), err)
		return phase, err
	}

	phase, err = r.reconcileSubscription(ctx, client, installation, productNamespace, operatorNamespace)
	if err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
		events.HandleError(r.recorder, installation, phase, fmt.Sprintf("Failed to reconcile %s subscription", constants.ThreeScaleSubscriptionName), err)
		return phase, err
	}

	phase, err = r.reconcileComponents(ctx, client, installation)
	if err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
		events.HandleError(r.recorder, installation, phase, "Failed to create components", err)
		return phase, err
	}
	// TODO:
	//phase, err = r.reconcileHost(ctx, client)
	//if err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
	//	events.HandleError(r.recorder, installation, phase, "Failed to reconcile host", err)
	//	return phase, err
	//}

	// TODO:
	//if string(r.Config.GetProductVersion()) != string(integreatlyv1alpha1.VersionGrafana) {
	//	r.Config.SetProductVersion(string(integreatlyv1alpha1.VersionGrafana))
	//	r.ConfigManager.WriteConfig(r.Config)
	//}

	// TODO:
	//alertsReconciler := r.newAlertReconciler(r.log)
	//if phase, err := alertsReconciler.ReconcileAlerts(ctx, client); err != nil || phase != integreatlyv1alpha1.PhaseCompleted {
	//	events.HandleError(r.recorder, installation, phase, "Failed to reconcile grafana alerts", err)
	//	return phase, err
	//}

	product.Host = r.Config.GetHost()
	product.Version = r.Config.GetProductVersion()
	product.OperatorVersion = r.Config.GetOperatorVersion()

	events.HandleProductComplete(r.recorder, installation, integreatlyv1alpha1.ProductsStage, r.Config.GetProductName())
	r.log.Info("Reconciled successfully")
	return integreatlyv1alpha1.PhaseCompleted, nil
}

func (r *Reconciler) reconcileComponents(ctx context.Context, client k8sclient.Client, installation *integreatlyv1alpha1.RHMI) (integreatlyv1alpha1.StatusPhase, error) {
	// TODO: Not sure if this is even required. It may just create one by default
	r.log.Info("reconciling container security custom resource")

	//grafana := &grafanav1alpha1.Grafana{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      "grafana",
	//		Namespace: r.Config.GetOperatorNamespace(),
	//	},
	//}
	//
	//status, err := controllerutil.CreateOrUpdate(ctx, client, grafana, func() error {
	//	owner.AddIntegreatlyOwnerAnnotations(grafana, r.installation)
	//	grafana.Spec = grafanav1alpha1.GrafanaSpec{
	//		Config: grafanav1alpha1.GrafanaConfig{
	//			Log: &grafanav1alpha1.GrafanaConfigLog{
	//				Mode:  "console",
	//				Level: "warn",
	//			},
	//		},
	//	}
	//	return nil
	//})
	//
	//if err != nil {
	//	return integreatlyv1alpha1.PhaseFailed, err
	//}
	//
	//r.log.Infof("Grafana CR: ", l.Fields{"status": status})
	//
	//r.log.Infof("Grafana datasource: ", l.Fields{"status": status})

	// if there are no errors, the phase is complete
	return integreatlyv1alpha1.PhaseCompleted, nil
}

func (r *Reconciler) reconcileSubscription(ctx context.Context, serverClient k8sclient.Client, inst *integreatlyv1alpha1.RHMI, productNamespace string, operatorNamespace string) (integreatlyv1alpha1.StatusPhase, error) {
	r.log.Info("reconciling subscription")

	target := marketplace.Target{
		Pkg:       constants.ContainerSecuritySubscriptionName,
		Namespace: operatorNamespace,
		Channel:   marketplace.IntegreatlyChannel,
	}
	catalogSourceReconciler := marketplace.NewConfigMapCatalogSourceReconciler(
		manifestPackage,
		serverClient,
		operatorNamespace,
		marketplace.CatalogSourceName,
	)
	return r.Reconciler.ReconcileSubscription(
		ctx,
		target,
		//[]string{productNamespace},
		[]string{}, // allnamespaces
		r.preUpgradeBackupExecutor(),
		serverClient,
		catalogSourceReconciler,
		r.log,
	)
}

func (r *Reconciler) preUpgradeBackupExecutor() backup.BackupExecutor {
	return backup.NewNoopBackupExecutor()
}

//func (r *Reconciler) reconcileHost(ctx context.Context, serverClient k8sclient.Client) (integreatlyv1alpha1.StatusPhase, error) {
//	grafanaRoute := &routev1.Route{}
//
//	err := serverClient.Get(ctx, k8sclient.ObjectKey{Name: defaultRoutename, Namespace: r.Config.GetOperatorNamespace()}, grafanaRoute)
//	if err != nil {
//		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("Failed to get route for Grafana: %w", err)
//	}
//
//	r.Config.SetHost("https://" + grafanaRoute.Spec.Host)
//	err = r.ConfigManager.WriteConfig(r.Config)
//	if err != nil {
//		return integreatlyv1alpha1.PhaseFailed, fmt.Errorf("Could not set Grafana route: %w", err)
//	}
//
//	return integreatlyv1alpha1.PhaseCompleted, nil
//}
