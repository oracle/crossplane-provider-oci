# Route Table with Multiple Route Rules

## Introduction
This is a basic composite resource example of a route table that is deployed using Crossplane and its OCI provider. This is not meant for production use, it is only for demonstration purposes.

The example references an existing VCN and compartment whose Oracle Cloud Identifiers (OCIDs) and other parameters are provided in the claim.

Three resources are composed.

| Kind             | Description                                                           |
|------------------|-----------------------------------------------------------------------|
| `RouteTable`     | A route table with two rules that reference NAT and Service gateways. |
| `NATGateway`     | A NAT gateway.                                                        |
| `ServiceGateway` | A Service gateway.                                                    |

## Prerequisites
- OCIDs of an existing compartment and VCN

## Set up role binding for crossplane-system account
You may run this script once for demonstration of all composition examples.
```
kubectl apply -f examples/oke_sample/crossplane-cluster-role-binding.yml
```

## Apply the example
1. Apply the Composite Resource Definition (XRD) and Composition.
    ```
    cd examples/routetable/composition
    kubectl apply -f definition_route_table.yaml
    kubectl apply -f composition_route_table.yaml
    ```
1. Copy `claim-route_table.yml` to `claim.yml` and enter the correct values for your environment, then apply it.
    ```
    kubectl apply -f claim.yml
    ```
1. Check composed resources to verify that the composite resource `routetabletest` has created three managed resources.
    ```
    $ kubectl get managed | grep routetabletest -B 1
    
    NAME                                                      READY   SYNCED   EXTERNAL-NAME                                                                                         AGE
    servicegateway.core.oci.upbound.io/routetabletest-hrznl   True    True     ocid1.servicegateway.oc1.ca-montreal-1.xxx   134m
    --
    NAME                                                  READY   SYNCED   EXTERNAL-NAME                                                                                     AGE
    natgateway.core.oci.upbound.io/routetabletest-zd47k   True    True     ocid1.natgateway.oc1.ca-montreal-1.xxx   134m
    --
    NAME                                                  READY   SYNCED   EXTERNAL-NAME                                                                                     AGE
    routetable.core.oci.upbound.io/routetabletest-4xcgb   True    True     ocid1.routetable.oc1.ca-montreal-1.xxx   134m```
    ```

## Contributing
This project is open source.  Please submit your contributions by forking this repository and submitting a pull request!  Oracle appreciates any contributions that are made by the open source community.

## License
Copyright (c) 2022 Oracle and its affiliates.

Licensed under the Universal Permissive License (UPL), Version 1.0.

See [LICENSE](LICENSE) for more details.

ORACLE AND ITS AFFILIATES DO NOT PROVIDE ANY WARRANTY WHATSOEVER, EXPRESS OR IMPLIED, FOR ANY SOFTWARE, MATERIAL OR CONTENT OF ANY KIND CONTAINED OR PRODUCED WITHIN THIS REPOSITORY, AND IN PARTICULAR SPECIFICALLY DISCLAIM ANY AND ALL IMPLIED WARRANTIES OF TITLE, NON-INFRINGEMENT, MERCHANTABILITY, AND FITNESS FOR A PARTICULAR PURPOSE.  FURTHERMORE, ORACLE AND ITS AFFILIATES DO NOT REPRESENT THAT ANY CUSTOMARY SECURITY REVIEW HAS BEEN PERFORMED WITH RESPECT TO ANY SOFTWARE, MATERIAL OR CONTENT CONTAINED OR PRODUCED WITHIN THIS REPOSITORY. IN ADDITION, AND WITHOUT LIMITING THE FOREGOING, THIRD PARTIES MAY HAVE POSTED SOFTWARE, MATERIAL OR CONTENT TO THIS REPOSITORY WITHOUT ANY REVIEW. USE AT YOUR OWN RISK. 

