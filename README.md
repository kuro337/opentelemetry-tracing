```
 ██████╗ ████████╗███████╗██╗         ████████╗██████╗  █████╗  ██████╗██╗███╗   ██╗ ██████╗ 
██╔═══██╗╚══██╔══╝██╔════╝██║         ╚══██╔══╝██╔══██╗██╔══██╗██╔════╝██║████╗  ██║██╔════╝ 
██║   ██║   ██║   █████╗  ██║            ██║   ██████╔╝███████║██║     ██║██╔██╗ ██║██║  ███╗
██║   ██║   ██║   ██╔══╝  ██║            ██║   ██╔══██╗██╔══██║██║     ██║██║╚██╗██║██║   ██║
╚██████╔╝   ██║   ███████╗███████╗       ██║   ██║  ██║██║  ██║╚██████╗██║██║ ╚████║╚██████╔╝
 ╚═════╝    ╚═╝   ╚══════╝╚══════╝       ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝╚═╝  ╚═══╝ ╚═════╝ 
                                                                                             

```

## OpenTelemetry Distributing Tracing 


- This  repo shows how to create a `Kubernetes` Cluster with an `OpenTelemtry Collector` running as a `Sidecar` and a `Go` application that uses the `OpenTelemetry Go SDK` to send `traces` to the Collector. 
`
- The Collector is configured to send traces to `Tempo` - an Open Source backend to receive telemetry data from various formats.

- `Grafana` is used to visualize the Trace Data. 

