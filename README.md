# NeMo Guardrails Controller

A standalone Kubernetes controller for managing NeMo Guardrails instances. This controller can be imported and used by multiple operators.

## Overview

The `nemo-guardrails-controller` provides a complete controller implementation for the `NemoGuardrails` Custom Resource Definition (CRD). It handles:

- Deployment of NeMo Guardrails servers
- Configuration management via ConfigMaps
- CA bundle mounting for custom certificates
- Authentication proxy setup (kube-rbac-proxy)
- OpenShift Route management
- Status reporting and conditions

## Features

- **Multi-Config Support**: Mount multiple NeMo configurations from different ConfigMaps
- **CA Bundle Management**: Automatic mounting of custom CA certificates from ConfigMaps
- **Authentication**: Optional kube-rbac-proxy integration for secure access
- **OpenShift Integration**: Automatic Route creation on OpenShift clusters
- **Status Tracking**: Comprehensive status conditions and CA bundle status reporting

## Installation

Add this controller to your operator:

```bash
go get github.com/trustyai-explainability/nemo-guardrails-controller@v0.1.0
```

## Usage

### Import into Your Operator

```go
import (
    nemosetup "github.com/trustyai-explainability/nemo-guardrails-controller/pkg/setup"
)

// In your main.go or controller setup
if err := nemosetup.SetupWithManager(mgr, namespace, configMapName, recorder); err != nil {
    setupLog.Error(err, "unable to create controller", "controller", "NemoGuardrails")
    os.Exit(1)
}
```

### Required ConfigMap

Your operator must provide a ConfigMap with image references:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-operator-config
  namespace: my-operator-system
data:
  nemo-guardrails-image: "quay.io/trustyai/nemo-guardrails:latest"
  kube-rbac-proxy: "gcr.io/kubebuilder/kube-rbac-proxy:v0.8.0"
```

### Example NemoGuardrails CR

```yaml
apiVersion: trustyai.opendatahub.io/v1alpha1
kind: NemoGuardrails
metadata:
  name: my-guardrails
  namespace: default
spec:
  nemoConfigs:
    - name: config1
      configMaps:
        - nemo-config-1
      default: true
    - name: config2
      configMaps:
        - nemo-config-2
  caBundleConfig:
    configMapName: my-ca-bundle
    configMapKeys:
      - ca-bundle.crt
  env:
    - name: LOG_LEVEL
      value: "INFO"
```

## Architecture

```
nemo-guardrails-controller/
├── api/
│   └── v1alpha1/              # API definitions
│       ├── nemoguardrails_types.go
│       └── groupversion_info.go
├── controllers/                # Controller logic
│   ├── nemoguardrail_controller.go
│   ├── deployment.go
│   ├── ca.go
│   ├── status.go
│   └── templates/             # Resource templates
├── config/
│   ├── crd/                   # CRD manifests
│   ├── rbac/                  # RBAC manifests
│   └── samples/               # Example CRs
└── pkg/
    └── setup/                 # Setup utilities for importing
```

## API Reference

### NemoGuardrailsSpec

| Field | Type | Description |
|-------|------|-------------|
| `nemoConfigs` | `[]NemoConfig` | List of NeMo configuration mounts |
| `caBundleConfig` | `*CABundleConfig` | CA bundle configuration (optional) |
| `env` | `[]corev1.EnvVar` | Environment variables for the main container (optional) |

### NemoConfig

| Field | Type | Description |
|-------|------|-------------|
| `name` | `string` | Configuration name (used as directory name) |
| `configMaps` | `[]string` | List of ConfigMap names to mount |
| `default` | `bool` | Whether this is the default configuration |

### CABundleConfig

| Field | Type | Description |
|-------|------|-------------|
| `configMapName` | `string` | Name of the ConfigMap containing CA certificates |
| `configMapNamespace` | `string` | Namespace of the ConfigMap (optional, defaults to CR namespace) |
| `configMapKeys` | `[]string` | Keys within the ConfigMap containing CA data (optional, defaults to `["ca-bundle.crt"]`) |

## Development

### Prerequisites

- Go 1.23 or later
- Kubernetes 1.29 or later
- controller-runtime v0.17.0

### Building

```bash
make build
```

### Running Tests

```bash
make test
```

### Generating CRDs

```bash
make manifests
```

## License

Apache License 2.0

## Contributing

Contributions are welcome! Please ensure all tests pass and CRDs are regenerated before submitting a pull request.
