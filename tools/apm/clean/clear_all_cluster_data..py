#!/usr/bin/env python
#---------------------------------------------------------------------
# Licensed Materials - Property of IBM
# 5737-A40
#
# (C) Copyright IBM Corporation 2018. All Rights Reserved.
# US Government Users Restricted Rights - Use, duplication
# or disclosure restricted by GSA ADP Schedule Contract
# with IBM Corp.
#---------------------------------------------------------------------

"""
Simple script to clear cluster data from a given tenant
"""

import sys
import os
import json
import math
import time

from pathos.multiprocessing import ProcessingPool as Pool
from future.utils import raise_with_traceback
import requests
from requests.packages.urllib3.exceptions import InsecureRequestWarning
requests.packages.urllib3.disable_warnings(InsecureRequestWarning) # Suppress


# CONFIGURATION
# NOTE/TODO: would be nice to integrate click here so this can be passed from command line invocation
provider_id = 'ba9b0e584f4661384e6c9d9983384e2d'
server_url = 'https://icp-proxy.cp4mcp-demo-002-a376efc1170b9b8ace6422196c51e491-0000.us-south.containers.appdomain.cloud'
tenant_id = 'id-mycluster-account'

SERVER_URL_ENV='IBM_APM_SERVER_INGRESS_URL'
NUM_PROC=8 # Used to simulate batch adds/deletes/updates to topology service and ingress
REQUEST_TIMEOUT_SECONDS=60 # request timeout for third party HTTP requests
TENANT_ID_ENV='APM_TENANT_ID'
PROVIDER_ID_ENV='PROVIDER_ID'
CLUSTER_NAME_ENV='CLUSTER_NAME'
SERVER_URL_ENV='IBM_APM_SERVER_INGRESS_URL'
RESOURCE_LIMIT=400 #1000000000 # max num of resources to return from topology service api
AUTHORIZATION_HEADER_KEY='Authorization'
BEARER_PREFIX='Bearer '
TENANT_HEADER_KEY='X-TenantID'
ACCEPT_HEADER_KEY='Accept'
PROVIDER_HEADER_KEY='Provider'
CONTENT_TYPE_HEADER_KEY='Content-Type'
TOPOLOGY_SERVICE_API_PREFIX='/applicationmgmt/0.9/'
SUBSCRIPTION_ENTITY='k8sSubscription'
# Kuberentes entity types
PROVIDER_ENTITY='k8sProvider'
NODE_ENTITY='k8sNode'
POD_ENTITY='k8sPod'
CONTAINER_ENTITY='k8sContainer'
IMAGE_ENTITY='k8sImage'
SERVICE_ENTITY='k8sService'
CLUSTER_ENTITY='k8sCluster'
NAMESPACE_ENTITY='k8sNamespace'
INGRESS_ENTITY='k8sIngress'
DEPLOYMENT_ENTITY='k8sDeployment' # NOTE: Openshift 'DeploymentConfig' will be saved as 'k8sDeployment'
REPLICA_SET_ENTITY='k8sReplicaSet'
DAEMON_SET_ENTITY='k8sDaemonSet'
JOB_ENTITY='k8sJob'
STATEFUL_SET_ENTITY='k8sStatefulSet'
POD_VOLUME_ENTITY='k8sPodVolume'
REPLICATION_CONTROLLER_ENTITY='k8sReplicationController'
APPLICATION_ENTITY='k8sApplication'
CRON_JOB_ENTITY='k8sCronJob'
# Openshift types
#PROJECT_ENTITY='k8sProject'
ROUTE_ENTITY='k8sRoute'
#OKD_APPLICATION_ENTITY='OkdApplication'
RESOURCE_QUOTA_ENTITY='k8sResourceQuota'
CLUSTER_RESOURCE_QUOTA_ENTITY='k8sClusterResourceQuota'
ALL_K8_TYPES=[SUBSCRIPTION_ENTITY,CLUSTER_ENTITY,NODE_ENTITY,SERVICE_ENTITY,POD_ENTITY,IMAGE_ENTITY,NAMESPACE_ENTITY,CONTAINER_ENTITY,INGRESS_ENTITY,APPLICATION_ENTITY,DEPLOYMENT_ENTITY,DAEMON_SET_ENTITY,REPLICATION_CONTROLLER_ENTITY,REPLICA_SET_ENTITY,JOB_ENTITY,STATEFUL_SET_ENTITY,CRON_JOB_ENTITY]
ALL_OKD_TYPES=[ROUTE_ENTITY,RESOURCE_QUOTA_ENTITY,CLUSTER_RESOURCE_QUOTA_ENTITY]


__author__ = "Abigail Johnson"
__version__ = "0.8.0"
__maintainer__ = "Abigail Johnson"
__email__ = "abigail.m.johnson@ibm.com"
__status__ = "Development"


class TopologyServiceLite(object):
    """
    Lightweight topo wrapper class
    """

    def __init__(self):
        """
        Initialize Topology Service client API
        """
        try:
            self.api_prefix = os.environ[SERVER_URL_ENV] + TOPOLOGY_SERVICE_API_PREFIX
            self.headers = {  TENANT_HEADER_KEY : os.environ[TENANT_ID_ENV],
                              ACCEPT_HEADER_KEY: 'application/json',
                              CONTENT_TYPE_HEADER_KEY : 'application/json',
                              PROVIDER_HEADER_KEY : os.environ[PROVIDER_ID_ENV]}
        except Exception as e:
            exc_type, exc_obj, exc_tb = sys.exc_info()
            fname = os.path.split(exc_tb.tb_frame.f_code.co_filename)[1]
            print(exc_type, fname, exc_tb.tb_lineno)
            print(str(e))


    def get_resources(self, filter_str='', alt_limit=None):
        """
        Gets all resources currently in IBM topology service
        Returns list of TopologyResource objects
        """
        r_limit = RESOURCE_LIMIT
        # Alternate limit for case when alot of data being requested
        if alt_limit:
            r_limit = alt_limit
        try:
            # Manually calculate pagination offset
            offset = 0
            all_resources = []
            resources_remaining = True
            # Paginate until empty resource list returned
            while resources_remaining:
                query_url = self.api_prefix + 'resources?_field=uid&_limit='+ str(r_limit) + '&_offset=' + str(offset) + filter_str
                r = requests.get(query_url, headers=self.headers, timeout=REQUEST_TIMEOUT_SECONDS,verify=False).json()
                # No more resources - stop pagination
                if len(r['_items']) == 0:
                    resources_remaining = False
                all_resources.extend( r['_items'] )
                offset = offset + len(r['_items'])
            sys.stdout.write('Fetched ' + str(len(all_resources)) + ' resources from IBM Topology Service\n')
            return all_resources
        except Exception as e:
            sys.stdout.write('ERROR retrieving resources from TopologyService: ' + str(e) + "\n")


    def delete_resources(self, id_list):
        """
        Batch Deletes stale resources from IBM topology service via DELETE
        Expects list of _ids that need to be deleted
        """
        def delete_single_resource(resource_id):
            """
            Multiprocess Worker Method to delete single resource from IBM topology service
            """
            r = requests.delete(self.api_prefix + 'resources/' + resource_id,headers=self.headers,verify=False)
        sys.stdout.write('Deleting ' + str(len(id_list)) + ' resources from IBM Topology Service\n')
        pool = Pool(NUM_PROC)
        pool.map(delete_single_resource, id_list)


os.environ[TENANT_ID_ENV] = tenant_id
os.environ[SERVER_URL_ENV] = server_url
os.environ[PROVIDER_ID_ENV] = provider_id

print('Begin cluster cleanup for provider ' + provider_id + ' on server ' + server_url + ' under tenant ' + tenant_id)
ts = TopologyServiceLite()
all_monitored_types = []
all_monitored_types.extend(ALL_K8_TYPES)
all_monitored_types.extend(ALL_OKD_TYPES)
all_filter_str = [ '&_field=uid' + '&_type=' + k8s_type for k8s_type in all_monitored_types ]
k8s_resources = []
# TODO: do in parrallel
for filter_str in all_filter_str:
    k8s_resource_list = ts.get_resources(filter_str=filter_str, alt_limit=200)
    k8s_resources.extend(k8s_resource_list)
print(str(len(k8s_resources)) + ' total Kubernetes resources in Topology Service')
all_cluster_data = [r['uid'] for r in k8s_resources if 'uid' in r.keys()]
print(str(len(all_cluster_data)) + ' total Kubernetes resources in Topology Service for provider ' + provider_id)
ts.delete_resources(all_cluster_data)
