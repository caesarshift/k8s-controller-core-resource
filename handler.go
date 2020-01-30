package main

import (
	log "github.com/sirupsen/logrus"
    //core_v1 "k8s.io/api/core/v1"
	myresource_v1 "github.com/caesarshift/k8s-controller-core-resource/pkg/apis/myresource/v1"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) {
	log.Info("TestHandler.ObjectCreated")
	// assert the type to a Pod object to pull out relevant data
	myr := obj.(*myresource_v1.MyResource)
	log.Infof("    Message: %s", myr.Spec.Message)
	log.Infof("    SomeValue: %d", *myr.Spec.SomeValue)
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}
