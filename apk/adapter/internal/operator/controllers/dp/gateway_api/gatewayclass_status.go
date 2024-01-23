package gateway_api

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

const (
	gatewayClassAcceptedMessage    = "Valid GatewayClass"
	gatewayClassNotAcceptedMessage = "Invalid GatewayClass"
)

func setGatewayClassAccepted(gwc *gwapiv1b1.GatewayClass, accepted bool) *gwapiv1b1.GatewayClass {
	gwc.Status.Conditions = merge(gwc.Status.Conditions, gatewayClassAcceptedCondition(gwc, accepted))
	return gwc
}

func gatewayClassAcceptedCondition(gwc *gwapiv1b1.GatewayClass, accepted bool) metav1.Condition {
	switch accepted {
	case true:
		return metav1.Condition{
			Type:               string(gwapiv1b1.GatewayClassConditionStatusAccepted),
			Status:             metav1.ConditionTrue,
			Reason:             string(gwapiv1b1.GatewayClassReasonAccepted),
			Message:            gatewayClassAcceptedMessage,
			ObservedGeneration: gwc.Generation,
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	default:
		return metav1.Condition{
			Type:               string(gwapiv1b1.GatewayClassConditionStatusAccepted),
			Status:             metav1.ConditionFalse,
			Reason:             string(gwapiv1b1.GatewayClassReasonInvalidParameters),
			Message:            gatewayClassNotAcceptedMessage,
			ObservedGeneration: gwc.Generation,
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	}
}
