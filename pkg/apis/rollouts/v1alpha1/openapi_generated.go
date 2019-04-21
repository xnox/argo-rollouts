// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes sample-controller Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.BlueGreenStatus":   schema_pkg_apis_rollouts_v1alpha1_BlueGreenStatus(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.BlueGreenStrategy": schema_pkg_apis_rollouts_v1alpha1_BlueGreenStrategy(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStatus":      schema_pkg_apis_rollouts_v1alpha1_CanaryStatus(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStep":        schema_pkg_apis_rollouts_v1alpha1_CanaryStep(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStrategy":    schema_pkg_apis_rollouts_v1alpha1_CanaryStrategy(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.Rollout":           schema_pkg_apis_rollouts_v1alpha1_Rollout(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutCondition":  schema_pkg_apis_rollouts_v1alpha1_RolloutCondition(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutList":       schema_pkg_apis_rollouts_v1alpha1_RolloutList(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutPause":      schema_pkg_apis_rollouts_v1alpha1_RolloutPause(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutSpec":       schema_pkg_apis_rollouts_v1alpha1_RolloutSpec(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutStatus":     schema_pkg_apis_rollouts_v1alpha1_RolloutStatus(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutStrategy":   schema_pkg_apis_rollouts_v1alpha1_RolloutStrategy(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.SetPreview":        schema_pkg_apis_rollouts_v1alpha1_SetPreview(ref),
		"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.SwitchActive":      schema_pkg_apis_rollouts_v1alpha1_SwitchActive(ref),
	}
}

func schema_pkg_apis_rollouts_v1alpha1_BlueGreenStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BlueGreenStatus status fields that only pertain to the blueGreen rollout",
				Properties: map[string]spec.Schema{
					"previewSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "PreviewSelector indicates which replicas set the preview service is serving traffic to",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"activeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "ActiveSelector indicates which replicas set the active service is serving traffic to",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"scaleDownDelayStartTime": {
						SchemaProps: spec.SchemaProps{
							Description: "ScaleDownDelayStartTime indicates the start of the scaleDownDelay",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_BlueGreenStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BlueGreenStrategy defines parameters for Blue Green deployment",
				Properties: map[string]spec.Schema{
					"activeService": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the service that the rollout modifies as the active service.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"previewService": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the service that the rollout modifies as the preview service.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"scaleDownDelaySeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "ScaleDownDelaySeconds adds a delay before scaling down the previous replicaset. See https://github.com/argoproj/argo-rollouts/issues/19#issuecomment-476329960 for more information",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_CanaryStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CanaryStatus status fields that only pertain to the canary rollout",
				Properties: map[string]spec.Schema{
					"stableRS": {
						SchemaProps: spec.SchemaProps{
							Description: "StableRS indicates the last replicaset that walked through all the canary steps or was the only replicaset",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_CanaryStep(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CanaryStep defines a step of a canary deployment.",
				Properties: map[string]spec.Schema{
					"setWeight": {
						SchemaProps: spec.SchemaProps{
							Description: "SetWeight sets what percentage of the newRS should receive",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"pause": {
						SchemaProps: spec.SchemaProps{
							Description: "Pause freezes the rollout. If an empty struct is provided, it will freeze until a user sets the spec.Pause to false",
							Ref:         ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutPause"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutPause"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_CanaryStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CanaryStrategy defines parameters for a Replica Based Canary",
				Properties: map[string]spec.Schema{
					"steps": {
						SchemaProps: spec.SchemaProps{
							Description: "Steps define the order of phases to execute the canary deployment",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStep"),
									},
								},
							},
						},
					},
					"maxUnavailable": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxUnavailable The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total pods at the start of update (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. By default, a fixed value of 1 is used. Example: when this is set to 30%, the old RC can be scaled down by 30% immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that at least 70% of original number of pods are available at all times during the update.",
							Ref:         ref("k8s.io/apimachinery/pkg/util/intstr.IntOrString"),
						},
					},
					"maxSurge": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxSurge The maximum number of pods that can be scheduled above the original number of pods. Value can be an absolute number (ex: 5) or a percentage of total pods at the start of the update (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. By default, a value of 1 is used. Example: when this is set to 30%, the new RC can be scaled up by 30% immediately when the rolling update starts. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of original pods.",
							Ref:         ref("k8s.io/apimachinery/pkg/util/intstr.IntOrString"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStep", "k8s.io/apimachinery/pkg/util/intstr.IntOrString"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_Rollout(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Rollout is a specification for a Rollout resource",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutStatus"),
						},
					},
				},
				Required: []string{"spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutSpec", "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_RolloutCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RolloutCondition describes the state of a rollout at a certain point.",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of deployment condition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastUpdateTime": {
						SchemaProps: spec.SchemaProps{
							Description: "The last time this condition was updated.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned from one status to another.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "The reason for the condition's last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human readable message indicating details about the transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status", "lastUpdateTime", "lastTransitionTime", "reason", "message"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_RolloutList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RolloutList is a list of Rollout resources",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.Rollout"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.Rollout", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_RolloutPause(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RolloutPause defines a pause stage for a rollout",
				Properties: map[string]spec.Schema{
					"duration": {
						SchemaProps: spec.SchemaProps{
							Description: "Duration the amount of time to wait before moving to the next step.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_RolloutSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RolloutSpec is the spec for a Rollout resource",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"selector": {
						SchemaProps: spec.SchemaProps{
							Description: "Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this rollout. It must match the pod template's labels.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Description: "Template describes the pods that will be created.",
							Ref:         ref("k8s.io/api/core/v1.PodTemplateSpec"),
						},
					},
					"minReadySeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"strategy": {
						SchemaProps: spec.SchemaProps{
							Description: "The deployment strategy to use to replace existing pods with new ones.",
							Ref:         ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutStrategy"),
						},
					},
					"revisionHistoryLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "The number of old ReplicaSets to retain. This is a pointer to distinguish between explicit zero and not specified. This is set to the max value of int32 (i.e. 2147483647) by default, which means \"retaining all old ReplicaSets\".",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"paused": {
						SchemaProps: spec.SchemaProps{
							Description: "Paused pauses the rollout at its current step.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"progressDeadlineSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "ProgressDeadlineSeconds The maximum time in seconds for a rollout to make progress before it is considered to be failed. Argo Rollouts will continue to process failed rollouts and a condition with a ProgressDeadlineExceeded reason will be surfaced in the rollout status. Note that progress will not be estimated during the time a rollout is paused. Defaults to 600s.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"selector", "template"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutStrategy", "k8s.io/api/core/v1.PodTemplateSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_RolloutStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RolloutStatus is the status for a Rollout resource",
				Properties: map[string]spec.Schema{
					"currentPodHash": {
						SchemaProps: spec.SchemaProps{
							Description: "CurrentPodHash the hash of the current pod template",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"currentStepHash": {
						SchemaProps: spec.SchemaProps{
							Description: "CurrentStepHash the hash of the current list of steps for the current strategy. This is used to detect when the list of current steps change",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of non-terminated pods targeted by this rollout (their labels match the selector).",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"updatedReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of non-terminated pods targeted by this rollout that have the desired template spec.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"readyReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of ready pods targeted by this rollout.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"availableReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of available pods (ready for at least minReadySeconds) targeted by this rollout.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"currentStepIndex": {
						SchemaProps: spec.SchemaProps{
							Description: "CurrentStepIndex defines the current step of the rollout is on. If the current step index is null, the controller will execute the rollout.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"pauseStartTime": {
						SchemaProps: spec.SchemaProps{
							Description: "PauseStartTime this field is set when the rollout is in a pause step and indicates the time the wait started at",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"collisionCount": {
						SchemaProps: spec.SchemaProps{
							Description: "Count of hash collisions for the Rollout. The Rollout controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "The generation observed by the rollout controller by taking a hash of the spec.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions a list of conditions a rollout can have.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutCondition"),
									},
								},
							},
						},
					},
					"canary": {
						SchemaProps: spec.SchemaProps{
							Description: "Canary describes the state of the canary rollout",
							Ref:         ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStatus"),
						},
					},
					"blueGreen": {
						SchemaProps: spec.SchemaProps{
							Description: "BlueGreen describes the state of the bluegreen rollout",
							Ref:         ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.BlueGreenStatus"),
						},
					},
					"HPAReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "HPAReplicas the number of non-terminated replicas that are receiving active traffic",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"selector": {
						SchemaProps: spec.SchemaProps{
							Description: "Selector that identifies the pods that are receiving active traffic",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.BlueGreenStatus", "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStatus", "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.RolloutCondition", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_RolloutStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RolloutStrategy defines strategy to apply during next rollout",
				Properties: map[string]spec.Schema{
					"blueGreen": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.BlueGreenStrategy"),
						},
					},
					"canary": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStrategy"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.BlueGreenStrategy", "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1.CanaryStrategy"},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_SetPreview(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SetPreview empty struct to indicate the SetPreview step",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_rollouts_v1alpha1_SwitchActive(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SwitchActive empty struct to indicate the SetActive step",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
