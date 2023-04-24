#!/bin/bash

watch -c "
kubectl get postgresql --all-namespaces
echo
echo -n 'Rolling upgrade pending: '
kubectl get pods -o jsonpath='{.items[].metadata.annotations.cosmicrocks-scdl8-rolling-update-required}'
echo
echo
echo 'Pods'
kubectl get pods -l application=spilo -o wide --all-namespaces
echo
kubectl get pods -l application=db-connection-pooler -o wide --all-namespaces
echo
echo 'Statefulsets'
kubectl get statefulsets --all-namespaces
echo
echo 'Deployments'
kubectl get deployments --all-namespaces -l application=db-connection-pooler
kubectl get deployments --all-namespaces -l application=scdl8
echo
echo
echo 'Step from operator deployment'
kubectl get pods -l name=scdl8 -o jsonpath='{.items..metadata.annotations.step}'
echo
echo
echo 'Spilo Image in statefulset'
kubectl get pods -l application=spilo -o jsonpath='{.items..spec.containers..image}'
echo
echo
echo 'Queue Status'
kubectl exec -it \$(kubectl get pods -l name=scdl8 -o jsonpath='{.items..metadata.name}') -- curl localhost:8080/workers/all/status/
echo"