# Supply Chain Templates

The package provides several templates to design paths to production on Kubernetes using Cartographer.

### Source (Flux)

* `flux-source-template`. Using Flux Source Controller, this template consumes an _application_ Git repository via a Git server or OCI registry, and passes it down to the rest of the supply chain on every code change.

### Image (kpack)

* `kpack-template`: it uses kpack, Cloud Native Buildpacks, and Paketo to transform application source code into a production-ready container image.

### Test (Tekton)

* `tekton-test-source-template`: it runs an instance of a Tekton pipeline to test the application source code.

### Scan (Grype and Trivy)

* `tekton-scan-image-template`: it provides a template to scan container images with Tekton and the configured vulnerability scanner.
* `tekton-scan-source-template`: it provides a template to scan application source code with Tekton and the configured vulnerability scanner.

### Conventions (Cartographer)

* `convention-template`: it applies configuration and best-practices to workloads at runtime by understanding the developer's intent, using Cartographer Conventions.

### Configuration (Carvel)

* `knative-config-template`: it uses Carvel `kapp` to package and configure the application as a Knative Service.

### Promotion (Tekton)

* `tekton-write-config-template`: it provides a template to publish deployment configuration to a container registry or Git repository for promotion to a specific environment.
* `tekton-write-config-and-pr-template`: it provides a template to publish deployment configuration to a Git repository for promotion to a specific environment via a pull request.

### Deliverable (Carvel and Flux)

* `deliverable-carvel-app-template`: it uses Carvel to generate a deliverable resource (`App`) for deploying the application on a Kubernetes cluster.
* `deliverable-flux-kustomization-template`: it uses Flux to generate a deliverable resource (`GitRepository` and `Kustomization`) for deploying the application on a Kubernetes cluster.

### Deploy (Carvel)

* `app-local-deployment-template`: it runs an application packaged as a Carvel `App` from local configuration.
* `app-gitops-deployment-template`: it runs an application packaged as a Carvel `App` from remote configuration (Git or OCI registry), based on either Carvel or Flux.
