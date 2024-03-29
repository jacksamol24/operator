// Copyright (c) 2020-2021 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package components

var (
	EnterpriseRelease string = "v3.11.0"

	ComponentAPIServer = component{
		Version: "v3.11.0",
		Image:   "tigera/cnx-apiserver",
	}

	ComponentComplianceBenchmarker = component{
		Version: "v3.11.0",
		Image:   "tigera/compliance-benchmarker",
	}

	ComponentComplianceController = component{
		Version: "v3.11.0",
		Image:   "tigera/compliance-controller",
	}

	ComponentComplianceReporter = component{
		Version: "v3.11.0",
		Image:   "tigera/compliance-reporter",
	}

	ComponentComplianceServer = component{
		Version: "v3.11.0",
		Image:   "tigera/compliance-server",
	}

	ComponentComplianceSnapshotter = component{
		Version: "v3.11.0",
		Image:   "tigera/compliance-snapshotter",
	}

	ComponentDeepPacketInspection = component{
		Version: "v3.11.0",
		Image:   "tigera/deep-packet-inspection",
	}

	ComponentEckElasticsearch = component{
		Version: "7.11.2",
		Image:   "tigera/elasticsearch",
	}

	ComponentEckKibana = component{
		Version: "7.11.2",
		Image:   "tigera/kibana",
	}

	ComponentElasticTseeInstaller = component{
		Version: "v3.11.0",
		Image:   "tigera/intrusion-detection-job-installer",
	}

	ComponentElasticsearch = component{
		Version: "v3.11.0",
		Image:   "tigera/elasticsearch",
	}

	ComponentElasticsearchOperator = component{
		Version: "1.7.1",
		Image:   "tigera/eck-operator",
	}

	ComponentEsCurator = component{
		Version: "v3.11.0",
		Image:   "tigera/es-curator",
	}

	ComponentEsProxy = component{
		Version: "v3.11.0",
		Image:   "tigera/es-proxy",
	}

	ComponentESGateway = component{
		Version: "v3.11.0",
		Image:   "tigera/es-gateway",
	}

	ComponentFluentd = component{
		Version: "v3.11.0",
		Image:   "tigera/fluentd",
	}

	ComponentFluentdWindows = component{
		Version: "v3.11.0",
		Image:   "tigera/fluentd-windows",
	}

	ComponentGuardian = component{
		Version: "v3.11.0",
		Image:   "tigera/guardian",
	}

	ComponentIntrusionDetectionController = component{
		Version: "v3.11.0",
		Image:   "tigera/intrusion-detection-controller",
	}

	ComponentKibana = component{
		Version: "v3.11.0",
		Image:   "tigera/kibana",
	}

	ComponentManager = component{
		Version: "v3.11.0",
		Image:   "tigera/cnx-manager",
	}

	ComponentDex = component{
		Version: "v3.11.0",
		Image:   "tigera/dex",
	}

	ComponentManagerProxy = component{
		Version: "v3.11.0",
		Image:   "tigera/voltron",
	}

	ComponentPacketCapture = component{
		Version: "v3.11.0",
		Image:   "tigera/packetcapture-api",
	}

	ComponentPrometheus = component{
		Version: "v2.17.2",
		Image:   "tigera/prometheus",
	}

	ComponentTigeraPrometheusService = component{
		Version: "v3.11.0",
		Image:   "tigera/prometheus-service",
	}

	ComponentPrometheusAlertmanager = component{
		Version: "v0.20.0",
		Image:   "tigera/alertmanager",
	}

	ComponentQueryServer = component{
		Version: "v3.11.0",
		Image:   "tigera/cnx-queryserver",
	}

	ComponentTigeraKubeControllers = component{
		Version: "v3.11.0",
		Image:   "tigera/kube-controllers",
	}

	ComponentTigeraNode = component{
		Version: "v3.11.0",
		Image:   "tigera/cnx-node",
	}

	ComponentTigeraTypha = component{
		Version: "v3.11.0",
		Image:   "tigera/typha",
	}

	ComponentTigeraCNI = component{
		Version: "v3.11.0",
		Image:   "tigera/cni",
	}

	ComponentCloudControllers = component{
		Version: "v3.11.0",
		Image:   "tigera/cloud-controllers",
	}

	ComponentElasticsearchMetrics = component{
		Version: "v3.11.0",
		Image:   "tigera/elasticsearch-metrics",
	}

	ComponentTigeraWindows = component{
		Version: "v3.11.0",
		Image:   "tigera/calico-windows-upgrade",
	}
	EnterpriseComponents = []component{
		ComponentAPIServer,
		ComponentComplianceBenchmarker,
		ComponentComplianceController,
		ComponentComplianceReporter,
		ComponentComplianceServer,
		ComponentComplianceSnapshotter,
		ComponentDeepPacketInspection,
		ComponentEckElasticsearch,
		ComponentEckKibana,
		ComponentElasticTseeInstaller,
		ComponentElasticsearch,
		ComponentElasticsearchOperator,
		ComponentEsCurator,
		ComponentEsProxy,
		ComponentFluentd,
		ComponentFluentdWindows,
		ComponentGuardian,
		ComponentIntrusionDetectionController,
		ComponentKibana,
		ComponentManager,
		ComponentDex,
		ComponentManagerProxy,
		ComponentPacketCapture,
		ComponentPrometheus,
		ComponentTigeraPrometheusService,
		ComponentPrometheusAlertmanager,
		ComponentQueryServer,
		ComponentTigeraKubeControllers,
		ComponentTigeraNode,
		ComponentTigeraTypha,
		ComponentTigeraCNI,
		ComponentCloudControllers,
		ComponentElasticsearchMetrics,
		ComponentESGateway,
		ComponentTigeraWindows,
	}
)
