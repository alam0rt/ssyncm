/*

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

package controllers

import (
	"context"

	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ssmv1alpha1 "github.com/alam0rt/ssyncm/api/v1alpha1"
)

var (
	// Ssm is the AWS service object
	Ssm *ssm.SSM
	// Session is the active AWS session
	Session *session.Session
)

// SsmSyncReconciler reconciles a SsmSync object
type SsmSyncReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=ssm.aws.afterpaytouch.dev,resources=ssmsyncs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ssm.aws.afterpaytouch.dev,resources=ssmsyncs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets/status,verbs=get

// Reconcile is the heart of the controller - it ensures the cluster is always
// in the desired state.
func (r *SsmSyncReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("ssmsync", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *SsmSyncReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ssmv1alpha1.SsmSync{}).
		Complete(r)
}

func init() {
	Session, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		fmt.Print("test")
	}
	Ssm = ssm.New(Session)

}
