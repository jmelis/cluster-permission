/*
Copyright 2023.

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

package v1alpha1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ConditionTypeAppliedRBACManifestWork string = "AppliedRBACManifestWork"
	ConditionTypeValidation              string = "Validation"
)

// ClusterPermissionSpec defines the desired state of ClusterPermission
type ClusterPermissionSpec struct {
	// ClusterRole represents the ClusterRole that is being created on the managed cluster
	// +optional
	ClusterRole *ClusterRole `json:"clusterRole,omitempty"`

	// ClusterRoleBinding represents the ClusterRoleBinding that is being created on the managed cluster
	// +optional
	ClusterRoleBinding *ClusterRoleBinding `json:"clusterRoleBinding,omitempty"`

	// Roles represents roles that are being created on the managed cluster
	// +optional
	Roles *[]Role `json:"roles,omitempty"`

	// RoleBindings represents RoleBindings that are being created on the managed cluster
	// +optional
	RoleBindings *[]RoleBinding `json:"roleBindings,omitempty"`
}

// ClusterRole represents the ClusterRole that is being created on the managed cluster
type ClusterRole struct {
	// Rules holds all the PolicyRules for this ClusterRole
	// +required
	Rules []rbacv1.PolicyRule `json:"rules"`
}

// ClusterRoleBinding represents the ClusterRoleBinding that is being created on the managed cluster
type ClusterRoleBinding struct {
	// Subject contains a reference to the object or user identities a ClusterPermission binding applies to.
	// Besides the typical subject for a binding, a ManagedServiceAccount can be used as a subject as well.
	// +required
	rbacv1.Subject `json:"subject"`

	// RoleRef contains information that points to the role being used
	// +optional
	RoleRef `json:"roleRef"`
}

// Role represents the Role that is being created on the managed cluster
type Role struct {
	// Namespace of the Role for that is being created on the managed cluster
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,4,opt,name=namespace"`

	// NamespaceSelector define the general labelSelector which namespace to apply the rules to
	// Note: the namespace must exists on the hub cluster
	// +optional
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`

	// Rules holds all the PolicyRules for this Role
	// +required
	Rules []rbacv1.PolicyRule `json:"rules"`
}

// RoleBinding represents the RoleBinding that is being created on the managed cluster
type RoleBinding struct {
	// Subject contains a reference to the object or user identities a ClusterPermission binding applies to.
	// Besides the typical subject for a binding, a ManagedServiceAccount can be used as a subject as well.
	// +required
	rbacv1.Subject `json:"subject"`

	// RoleRef contains information that points to the role being used
	// +required
	RoleRef `json:"roleRef"`

	// Namespace of the Role for that is being created on the managed cluster
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,4,opt,name=namespace"`

	// NamespaceSelector define the general labelSelector which namespace to apply the rules to
	// Note: the namespace must exists on the hub cluster
	// +optional
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
}

// RoleRef contains information that points to the role being used
type RoleRef struct {
	// Kind is the type of resource being referenced
	// +required
	Kind string `json:"kind"`

	// Name is the name of the existing role to be referenced
	// +optional
	Name string `json:"name"`
}

// ClusterPermissionStatus defines the observed state of ClusterPermission
type ClusterPermissionStatus struct {
	// Conditions is the condition list.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClusterPermission is the Schema for the clusterpermissions API
type ClusterPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterPermissionSpec   `json:"spec,omitempty"`
	Status ClusterPermissionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterPermissionList contains a list of ClusterPermission
type ClusterPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterPermission `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterPermission{}, &ClusterPermissionList{})
}
