package controllerruntime

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

func Success() (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func Fail(e error) (ctrl.Result, error) {
	return ctrl.Result{}, e
}
