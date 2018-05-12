# kdisco

kidsco: a kubernetes service which simply lists pods with a given label

    $ curl http://kdisco/pods?label=name\&value=my-awesome-app
    my-awesome-app-67823478-3498
    my-awesome-app-67823478-8424
    my-awesome-app-67823478-5412

This is useful if you ever want to contact every pod in a service. For example, to trigger a data-refresh

# Example

A minimal kubernetes yaml file for kdisco might look like this, but you probably want to use a Service and a Deployment with 2+ replicas

```
---
apiVersion: v1
kind: Pod
metadata:
  name: kdisco
  labels:
    name: kdisco
spec:
  containers:
  - name: kdisco
    image: laher/kdisco:v0.0.1
    imagePullPolicy: IfNotPresent
    resources:
      requests:
        cpu: 50m
        memory: 20Mi
      limits:
        cpu: 50m
        memory: 64Mi
    command: 
      - "/kdisco"
      - "-namespace=my-namespace"
      - "--"
```
