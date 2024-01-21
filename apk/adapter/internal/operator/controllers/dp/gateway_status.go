package dp

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func setGatewayAccepted(gw *gwapiv1b1.Gateway, accepted bool, msg string) *gwapiv1b1.Gateway {
	gw.Status.Conditions = merge(gw.Status.Conditions, gatewayStatusAcceptedCondition(gw, accepted, msg))
	return gw
}

// setGatewayProgrammed inserts or updates the Programmed condition for the provided Gateway resource.
func setGatewayProgrammed(gw *gwapiv1b1.Gateway, ready bool, msg string) *gwapiv1b1.Gateway {
	gw.Status.Conditions = merge(gw.Status.Conditions, gatewayStatusProgrammedCondition(gw, ready, msg))
	return gw
}

func gatewayStatusProgrammedCondition(gw *gwapiv1b1.Gateway, scheduled bool, msg string) metav1.Condition {
	switch scheduled {
	case true:
		return metav1.Condition{
			Type:               string(gwapiv1b1.GatewayConditionProgrammed),
			Status:             metav1.ConditionTrue,
			Reason:             string(gwapiv1b1.GatewayReasonProgrammed),
			Message:            msg,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	default:
		return metav1.Condition{
			Type:               string(gwapiv1b1.GatewayConditionProgrammed),
			Status:             metav1.ConditionFalse,
			Reason:             string(gwapiv1b1.GatewayReasonListenersNotReady),
			Message:            msg,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	}
}

func gatewayStatusAcceptedCondition(gw *gwapiv1b1.Gateway, accepted bool, msg string) metav1.Condition {
	switch accepted {
	case true:
		return metav1.Condition{
			Type:               string(gwapiv1b1.GatewayConditionAccepted),
			Status:             metav1.ConditionTrue,
			Reason:             string(gwapiv1b1.GatewayConditionAccepted),
			Message:            msg,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	default:
		return metav1.Condition{
			Type:               string(gwapiv1b1.GatewayConditionAccepted),
			Status:             metav1.ConditionFalse,
			Reason:             string(gwapiv1b1.GatewayReasonNoResources),
			Message:            msg,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	}
}

func gatewayStatusReadyCondition(gw *gwapiv1b1.Gateway, scheduled bool, msg string) metav1.Condition {
	switch scheduled {
	case true:
		return metav1.Condition{
			Type:               string(gwapiv1b1.GatewayConditionReady),
			Status:             metav1.ConditionTrue,
			Reason:             string(gwapiv1b1.GatewayReasonReady),
			Message:            msg,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	default:
		return metav1.Condition{
			Type:               string(gwapiv1b1.GatewayConditionReady),
			Status:             metav1.ConditionFalse,
			Reason:             string(gwapiv1b1.GatewayReasonListenersNotReady),
			Message:            msg,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	}
}

func gatewayListenerProgrammedCondition(gw *gwapiv1b1.Gateway, ready bool, msg string) metav1.Condition {
	switch ready {
	case true:
		return metav1.Condition{
			Type:               string(gwapiv1b1.ListenerConditionProgrammed),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
			Reason:             string(gwapiv1b1.ListenerConditionProgrammed),
			Message:            msg,
		}
	default:
		return metav1.Condition{
			Type:               string(gwapiv1b1.ListenerConditionProgrammed),
			Status:             metav1.ConditionFalse,
			Reason:             string(gwapiv1b1.ListenerReasonPending),
			Message:            msg,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	}
}

func gatewayListenerAcceptedCondition(gw *gwapiv1b1.Gateway, ready bool, msg string) metav1.Condition {
	switch ready {
	case true:
		return metav1.Condition{
			Type:               string(gwapiv1b1.ListenerConditionAccepted),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
			Reason:             string(gwapiv1b1.ListenerConditionAccepted),
			Message:            msg,
		}
	default:
		return metav1.Condition{
			Type:               string(gwapiv1b1.ListenerConditionAccepted),
			Status:             metav1.ConditionFalse,
			Reason:             string(gwapiv1b1.ListenerReasonPending),
			Message:            msg,
			ObservedGeneration: gw.GetGeneration(),
			LastTransitionTime: metav1.NewTime(time.Now()),
		}
	}
}

func gatewayListenerInvalidRouteKinds(gw *gwapiv1b1.Gateway, msg string) metav1.Condition {
	return metav1.Condition{
		Type:               string(gwapiv1b1.ListenerConditionResolvedRefs),
		Status:             metav1.ConditionFalse,
		Reason:             string(gwapiv1b1.ListenerReasonInvalidRouteKinds),
		Message:            msg,
		ObservedGeneration: gw.GetGeneration(),
		LastTransitionTime: metav1.NewTime(time.Now()),
	}
}
