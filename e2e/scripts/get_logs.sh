#!/bin/bash
kubectl logs $(kubectl get pods -l name=scdl8 --field-selector status.phase=Running -o jsonpath='{.items..metadata.name}')
