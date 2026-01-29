package controllers

const (
	// Version is the version of the nemo-guardrails-controller
	Version = "0.1.0"

	ServiceName                    = "NEMO_GUARDRAILS"
	nemoGuardrailsImageKey         = "nemo-guardrails-image"
	configMapKubeRBACProxyImageKey = "kube-rbac-proxy"
	finalizerName                  = "trustyai.opendatahub.io/nemo-guardrails-finalizer"
)
