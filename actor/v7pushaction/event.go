package v7pushaction

type Event string

const (
	BoundRoutes                     Event = "bound routes"
	BoundServices                   Event = "bound services"
	ConfiguringServices             Event = "configuring services"
	CreatedApplication              Event = "created application"
	CreatedRoutes                   Event = "created routes"
	CreatingAndMappingRoutes        Event = "creating and mapping routes"
	CreatingArchive                 Event = "creating archive"
	CreatingPackage                 Event = "creating package"
	PollingBuild                    Event = "polling build"
	ReadingArchive                  Event = "reading archive"
	ResourceMatching                Event = "resource matching"
	RetryUpload                     Event = "retry upload"
	ScaleWebProcess                 Event = "scaling the web process"
	ScaleWebProcessComplete         Event = "scaling the web process complete"
	SetDropletComplete              Event = "set droplet complete"
	SetHealthCheck                  Event = "setting the health check type on the process"
	SetHealthCheckComplete          Event = "completed setting the health check type on the process"
	SettingDroplet                  Event = "setting droplet"
	SettingUpApplication            Event = "setting up application"
	SkippingApplicationCreation     Event = "skipping creation"
	StagingComplete                 Event = "staging complete"
	StartingStaging                 Event = "starting staging"
	UnmappingRoutes                 Event = "unmapping routes"
	UpdatedApplication              Event = "updated application"
	UploadDropletComplete           Event = "upload droplet complete"
	UploadingApplication            Event = "uploading application"
	UploadingApplicationWithArchive Event = "uploading application with archive"
	UploadingDroplet                Event = "uploading droplet"
	UploadWithArchiveComplete       Event = "upload complete"
	Complete                        Event = "complete"
)
