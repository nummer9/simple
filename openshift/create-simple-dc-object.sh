#! /bin/bash

if [[ -z $1 ]] 
then
    echo "No variable provided! Please do provide variable."
else
    echo "custom variable is: ${1}"
    echo "creating deployment-config resource / object from template..."
    oc process -f simple-dc-template.yaml -p CUSTOM_VARIABLE=${1} | oc create -f -
fi

