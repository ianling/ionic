package events

const (
	// AnalysisFailed occurs when an analysis has failed
	AnalysisFailed = "analysis_failed"
	// AnalysisFinished occurs when an analysis has finished
	AnalysisFinished = "analysis_finished"
	// AnalysisPassed occurs when an analysis has passed
	AnalysisPassed = "analysis_passed"
	// DeliveryFailed occurs when a delivery has failed
	DeliveryFailed = "delivery_failed"
	// ArtifactDeliveryFailed occurs when an artifact delivery has failed
	ArtifactDeliveryFailed = "artifact_delivery_failed"
	// ReportDeliveryFailed occurs when a report delivery has failed
	ReportDeliveryFailed = "report_delivery_failed"
	// SevaDeliveryFailed occurs when seva delivery has failed
	SevaDeliveryFailed = "seva_delivery_failed"
	// DeliveryFinished occurs when a delivery has finished
	DeliveryFinished = "delivery_finished"
	// ArtifactDeliveryFinished occurs when an artifact delivery has finished
	ArtifactDeliveryFinished = "artifact_delivery_finished"
	// ReportDeliveryFinished occurs when a report delivery has finished
	ReportDeliveryFinished = "report_delivery_finished"
	// SevaDeliveryFinished occurs when a seva delivery has finished
	SevaDeliveryFinished = "seva_delivery_finished"
	// DeliveryCanceled occurs when a delivery has been canceled
	DeliveryCanceled = "delivery_canceled"
	// ArtifactDeliveryCanceled occurs when an artifact delivery has been canceled
	ArtifactDeliveryCanceled = "artifact_delivery_canceled"
	// ReportDeliveryCanceled occurs when a report delivery has been canceled
	ReportDeliveryCanceled = "report_delivery_canceled"
	// SevaDeliveryCanceled occurs when a seva delivery has been canceled
	SevaDeliveryCanceled = "seva_delivery_canceled"
	// ProjectAdded occurs when a project is added
	ProjectAdded = "project_added"
	// VersionAdded occurs when a version is added
	VersionAdded = "version_added"
	// AccountCreated occurs when an account is created
	AccountCreated = "account_created"
	// ForgotPassword occurs when a user triggers a password reset
	ForgotPassword = "forgot_password"
	// PasswordChanged occurs when a user's password is changed
	PasswordChanged = "password_changed"
	// UserSignup occurs when a user signs up
	UserSignup = "user_signup"
	// UserSignupStarted occurs when a user begins the signup process
	UserSignupStarted = "user_signup_started"
	// VulnerabilityAdded occurs when a vulnerability is added
	VulnerabilityAdded = "vulnerability_added"
	// VulnerabilityUpdated occurs when a vulnerability is updated
	VulnerabilityUpdated = "vulnerability_updated"
	// ProjectFlipped occurs when a project's latest analysis has flipped from passing to failing
	ProjectFlipped = "project_flipped"
)
