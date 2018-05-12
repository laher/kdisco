package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

func main() {
	var (
		listen    = flag.String("listen", ":8080", "listen address")
		namespace = flag.String("namespace-restrict", "", "Restrict namespace (if empty 'namespace' should be sent as a query parameter)")
	)
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		//for security purposes the namespace can be fixed at startup
		ns := *namespace
		if ns == "" {
			ns = q.Get("namespace")
		}
		pods, err := listBy(ns, q.Get("label"), q.Get("value"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		for _, p := range pods {
			fmt.Fprintln(w, p)
		}
	})

	log.Printf("Running on: %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))

}

func listBy(namespace string, label string, value string) ([]string, error) {
	var (
		ctx      = context.Background()
		pods     corev1.PodList
		selector = new(k8s.LabelSelector)
		ret      = []string{}
	)
	selector.In(label, value)

	client, err := k8s.NewInClusterClient()
	if err != nil {
		return nil, err
	}
	if err := client.List(ctx, namespace, &pods, selector.Selector()); err != nil {
		return nil, err
	}
	for _, p := range pods.Items {
		ret = append(ret, *p.Metadata.Name)
	}
	return ret, nil
}
