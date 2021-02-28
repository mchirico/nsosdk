package controllers

import (
	"context"
	topv1alpha1 "github.com/mchirico/ns-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

// TODO-mmc:  Need to finish setup of tests ...
//  Need to write stubs.  Currently this will call Reconcile
var _ = Describe("Namespecial controller", func() {

	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		jobName      = "stuff"
		jobNamespace = "default"
		JobName      = "test-job"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When Deleting Namespace `stuff`", func() {
		It("Should act ", func() {
			By("By creating a new namespace")

			ctx := context.Background()

			// TODO-mmc:  Appears to work but clean up
			nsJob := &corev1.Namespace{}
			nsJob.Namespace = "stuff"
			k8sClient.Create(ctx, nsJob)
			Expect(k8sClient.Delete(ctx, nsJob)).ShouldNot(Succeed())

			cronjobLookupKey := types.NamespacedName{Name: jobName, Namespace: jobNamespace}
			var instance topv1alpha1.Namespecial

			// We'll need to retry getting this, given that creation may not immediately happen.
			Eventually(func() bool {
				err := k8sClient.Get(ctx, cronjobLookupKey, &instance)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeFalse()) // TODO-mmc: put in the correct value and set to true

			// TODO-mmc:  Need to finish setkup of tests and call reconcile
			//Expect(k8sClient.Create(ctx, nsJob)).Should(Succeed())

		})
	})
})
