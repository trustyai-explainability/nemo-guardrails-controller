package setup

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	nemoguardrailsv1alpha1 "github.com/trustyai-explainability/nemo-guardrails-controller/api/v1alpha1"
	"github.com/trustyai-explainability/nemo-guardrails-controller/controllers"
)

// ControllerName is the unique identifier for this controller
const ControllerName = "NEMO_GUARDRAILS"

// SetupController registers the NemoGuardrails controller with the manager
// Parameters:
//   - mgr: The controller-runtime manager
//   - namespace: The namespace to watch for operator config (empty for all namespaces)
//   - configMap: Name of the operator config map containing image references
//   - recorder: Event recorder for Kubernetes events
func SetupController(mgr manager.Manager, namespace, configMap string, recorder record.EventRecorder) error {
	return controllers.ControllerSetUp(mgr, namespace, configMap, recorder)
}

// RegisterScheme adds the NemoGuardrails API types to the scheme
func RegisterScheme(scheme *runtime.Scheme) error {
	return nemoguardrailsv1alpha1.AddToScheme(scheme)
}

// SetupWithManager is a convenience function that registers the scheme and controller
func SetupWithManager(mgr manager.Manager, namespace, configMap string, recorder record.EventRecorder) error {
	if err := RegisterScheme(mgr.GetScheme()); err != nil {
		return err
	}
	return SetupController(mgr, namespace, configMap, recorder)
}
