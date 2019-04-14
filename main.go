package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env"
	"github.com/ddspog/colog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type config struct {
	LabelGroups []string `env:"POD_LABELS,required" envSeparator:";"`
}

func init() {
	colog.Register()
}

func main() {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("error:", err)
	}

	labelMaps := parseLabelGroups(cfg.LabelGroups)
	log.Println("info: looking for", labelMaps, "labels")

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalln("error:", err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for {
		if scanPods(clientset, labelMaps) {
			log.Println("info: found all the labels in running pods!")
			os.Exit(0)
		}

		log.Println("debug: labels not found in the running pods, sleeping and retrying")
		time.Sleep(2 * time.Second)
	}
}

func parseLabelGroups(lgs []string) []map[string]string {
	lms := make([]map[string]string, len(lgs))
	for i, lg := range lgs {
		lms[i] = make(map[string]string)
		ls := strings.Split(lg, ",")
		for _, l := range ls {
			kv := strings.Split(l, "=")
			lms[i][kv[0]] = kv[1]
		}
	}
	return lms
}

func mapInMap(src map[string]string, dst map[string]string) bool {
	var f int
	for sk, sv := range src {
		if dv, ok := dst[sk]; ok {
			if sv == dv {
				f++
				if f == len(src) {
					return true
				}
			}
		}
	}
	return false
}

func scanPods(clientset *kubernetes.Clientset, labelMaps []map[string]string) bool {
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("error:", err)
	}

	log.Printf("debug: there are currently %d pods in the cluster\n", len(pods.Items))

	var f int
	for _, p := range pods.Items {
		if p.Status.Phase == "Running" {
			for _, lm := range labelMaps {
				if mapInMap(lm, p.ObjectMeta.Labels) {
					f++
					if f == len(labelMaps) {
						return true
					}
				}
			}
		}
	}
	return false
}
