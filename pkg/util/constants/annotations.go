package constants

// Names and values in Kubernetes annotation for services, statefulsets and volumes
const (
	CosmicRocksDNSNameAnnotation       = "external-dns.alpha.kubernetes.io/hostname"
	ElbTimeoutAnnotationName           = "service.beta.kubernetes.io/aws-load-balancer-connection-idle-timeout"
	ElbTimeoutAnnotationValue          = "3600"
	KubeIAmAnnotation                  = "iam.amazonaws.com/role"
	VolumeStorateProvisionerAnnotation = "pv.kubernetes.io/provisioned-by"
	PostgresqlControllerAnnotationKey  = "acid.cosmic.rocks/controller"
)
