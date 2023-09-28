# Observability OpenTelemetry Golang

- OpenTelemetry Golang

```bash
# Adding Tracing and defining Spans
go get go.opentelemetry.io/otel go.opentelemetry.io/otel/trace

# Setup Exporter
go get go.opentelemetry.io/otel/sdk \
       go.opentelemetry.io/otel/exporters/stdout/stdouttrace

```

- Figure out how to get traces to sidecar 

- https://github.com/open-telemetry/opentelemetry-go/discussions/3726

- https://github.com/open-telemetry/opentelemetry-collector/issues/4829

- We add traces to our functions to establish a structure

- Create the Exporter to send Telemetry Data to the Console or backend , etc.

- Then create a resource so Exporter knows where data is coming from

- We then use a Trace Provider to connect the Exporter and Resource

- Kube 

- Create Image for go app 

- We deploy the OTEL Collector as a Sidecar with the app deployment. 
- Within go code - we setup an Exporter to export Traces to the Collector
- Our Collector is configured to send Traces to Tempo
- Grafana uses Tempo as the backend

- Tempo 

- Setting up Tempo Backend 

- Tempo config - local storage 

```yaml
global:
  clusterDomain: "cluster.local"
gateway:
  enabled: true
storage:
  trace:
    backend: local
traces:
  otlp:
    grpc:
      enabled: true
    http:
      enabled: true
  zipkin:
    enabled: false
  jaeger:
    thriftHttp:
      enabled: false
  opencensus:
    enabled: false
distributor:
  config:
    log_received_spans:
      enabled: true


```

```bash
helm upgrade --install tempo grafana/tempo-distributed -f tempo-config.yaml


```
- Grafana Setup with Tempo defined 

```yaml
env:
  GF_AUTH_ANONYMOUS_ENABLED: true
  GF_AUTH_ANONYMOUS_ORG_ROLE: "Admin"
  GF_AUTH_DISABLE_LOGIN_FORM: true

datasources:
  datasources.yaml:
    apiVersion: 1
    datasources:
      - name: Tempo
        type: tempo
        access: proxy
        orgId: 1
        url: http://tempo-gateway
        basicAuth: false
        isDefault: true
        version: 1
        editable: false
        apiVersion: 1
        uid: tempo

```

```bash
helm upgrade --install grafana grafana/grafana -f grafana-config.yaml
```

- Config Map for OTel Collector Config

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: otel-collector-conf
  labels:
    app: opentelemetry
    component: otel-collector-conf
data:
  otel-collector-config.yaml: |
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: 0.0.0.0:4317
          http:
            endpoint: 0.0.0.0:4318
    processors:
    extensions:
      health_check: {}
    exporters:
      otlp:
        endpoint: http://tempo-gateway/otlp
        tls:
          insecure: true
          insecure_skip_verify: true
    service:
      extensions: [health_check]
      pipelines:
        traces:
          receivers: [otlp]
          processors: []
          exporters: [otlp]


```

```bash
kubectl apply -f otel-collector-config.yaml

kubectl apply -f go-otel-dep.yaml

kubectl delete deployment go-app-deployment
kubectl delete configmap otel-collector-conf

# Getting Tempo values applied 
⮞  kubectl get svc -l app.kubernetes.io/name=tempo

# Test curl 

kubectl port-forward pod/go-app-deployment-5f45b89f66-mgp74 8080:8080

kubectl logs go-app-deployment-5f45b89f66-mgp74  -c otel-collector

kubectl logs go-app-deployment-5f45b89f66-mgp74   

kubectl get pods   

curl http://localhost:8080/fibonacci/12
curl http://localhost:8080/fibonacci/5
curl http://localhost:8080/fibonacci/6 
curl http://localhost:8080/fibonacci/12
curl http://localhost:8080/fibonacci/12
curl http://localhost:8080/fibonacci/5
curl http://localhost:8080/fibonacci/5
curl http://localhost:8080/fibonacci/5
curl http://localhost:8080/fibonacci/5
curl http://localhost:8080/fibonacci/5
curl http://localhost:8080/fibonacci/5
curl http://localhost:8080/fibonacci/5
curl -X POST http://localhost:8080/fibonacci/5
curl -X POST http://localhost:8080/fibonacci/5
curl http://localhost:8080/fibonacci/5abc
curl http://localhost:8080/fibonacci/5qqqq
curl http://localhost:8080/fibonacci/12
curl http://localhost:8080/stop   


kubectl exec -it go-app-deployment-5f45b89f66-jtml6 -- /bin/sh

# Running Processes 
ps aux

# Send ctrl-c to kill the process 1 
kill -2 1 


```
