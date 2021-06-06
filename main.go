package main

import (
	"fmt"
	"github.com/spongeprojects/client-go/api/spongeprojects.com/v1alpha1"
	"github.com/spongeprojects/client-go/client/clientset/versioned"
	"github.com/spongeprojects/client-go/client/informers/externalversions"
	"github.com/spongeprojects/magicconch"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"os"
)

// main just demonstrates usage of this package
func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	magicconch.Must(err)

	clientset, err := versioned.NewForConfig(config)
	magicconch.Must(err)

	informerFactory := externalversions.NewSharedInformerFactory(clientset, 0)
	channelsInformer := informerFactory.Spongeprojects().V1alpha1().Channels().Informer()
	channelsLister := informerFactory.Spongeprojects().V1alpha1().Channels().Lister()

	channelsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c, ok := obj.(*v1alpha1.Channel)
			if !ok {
				return
			}
			fmt.Printf("channel added: %s\n", c.Name)
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	klog.Info("start syncing")

	go informerFactory.Start(stopCh)

	klog.Info("waiting for cache to sync...")

	cache.WaitForCacheSync(stopCh, channelsInformer.HasSynced)

	klog.Info("cache synced")

	channels, err := channelsLister.List(labels.Everything())
	magicconch.Must(err)

	fmt.Printf("listing %d channels:\n", len(channels))

	for _, channel := range channels {
		fmt.Printf("  %s\n", channel.Name)
	}

	<-stopCh
}
