package go_obs

// Indicates a scene change.
type SwitchScenesEvent struct {
	eventData
	// The new scene.
	SceneName string `json:"scene-name"`
	// List of scene items in the new scene. Same specification as
	// [`GetCurrentScene`](#getcurrentscene).
	Sources []SceneItem `json:"sources"`
}

//   Note: This event is not fired when the scenes are reordered.
type ScenesChangedEvent struct {
	eventData
	// Scenes list.
	Scenes []Scene `json:"scenes"`
}

// Triggered when switching to another scene collection or when renaming the
// current scene collection.
type SceneCollectionChangedEvent struct {
	eventData
	// Name of the new current scene collection.
	SceneCollection string `json:"sceneCollection"`
}

// Triggered when a scene collection is created, added, renamed, or removed.
type SceneCollectionListChangedEvent struct {
	eventData
	// Scene collections list.
	SceneCollections []struct {
		// Scene collection name.
		Name string `json:"name"`
	} `json:"sceneCollections"`
}

// Triggered when switching to another profile or when renaming the current
// profile.
type ProfileChangedEvent struct {
	eventData
	// Name of the new current profile.
	Profile string `json:"profile"`
}

// Triggered when a profile is created, added, renamed, or removed.
type ProfileListChangedEvent struct {
	eventData
	// Profiles list.
	Profiles []struct {
		// Profile name.
		Name string `json:"name"`
	} `json:"profiles"`
}

// Virtual cam started successfully.
type VirtualCamStartedEvent struct {
	eventData
}

// Virtual cam stopped successfully.
type VirtualCamStoppedEvent struct {
	eventData
}

// Emitted every 2 seconds after enabling it by calling SetHeartbeat.
type HeartbeatEvent struct {
	eventData
	// Toggles between every JSON message as an "I am alive" indicator.
	Pulse bool `json:"pulse"`
	// Current active profile.
	CurrentProfile string `json:"current-profile,omitempty"`
	// Current active scene.
	CurrentScene string `json:"current-scene,omitempty"`
	// Current streaming state.
	Streaming *bool `json:"streaming,omitempty"`
	// Total time (in seconds) since the stream started.
	TotalStreamTime *int `json:"total-stream-time,omitempty"`
	// Total bytes sent since the stream started.
	TotalStreamBytes *int `json:"total-stream-bytes,omitempty"`
	// Total frames streamed since the stream started.
	TotalStreamFrames *int `json:"total-stream-frames,omitempty"`
	// Current recording state.
	Recording *bool `json:"recording,omitempty"`
	// Total time (in seconds) since recording started.
	TotalRecordTime *int `json:"total-record-time,omitempty"`
	// Total bytes recorded since the recording started.
	TotalRecordBytes *int `json:"total-record-bytes,omitempty"`
	// Total frames recorded since the recording started.
	TotalRecordFrames *int `json:"total-record-frames,omitempty"`
	// OBS Stats
	Stats OBSStats `json:"stats"`
}

// A custom broadcast message, sent by the server, requested by one of the
// websocket clients.
type BroadcastCustomMessageEvent struct {
	eventData
	// Identifier provided by the sender
	Realm string `json:"realm"`
	// User-defined data
	Data interface{} `json:"data"`
}

// Scene items within a scene have been reordered.
type SourceOrderChangedEvent struct {
	eventData
	// Name of the scene where items have been reordered.
	SceneName string `json:"scene-name"`
	// Ordered list of scene items
	SceneItems []struct {
		// Item source name
		SourceName string `json:"source-name"`
		// Scene item unique ID
		ItemId int `json:"item-id"`
	} `json:"scene-items"`
}

// A scene item has been added to a scene.
type SceneItemAddedEvent struct {
	eventData
	// Name of the scene.
	SceneName string `json:"scene-name"`
	// Name of the item added to the scene.
	ItemName string `json:"item-name"`
	// Scene item ID
	ItemId int `json:"item-id"`
}

// A scene item has been removed from a scene.
type SceneItemRemovedEvent struct {
	eventData
	// Name of the scene.
	SceneName string `json:"scene-name"`
	// Name of the item removed from the scene.
	ItemName string `json:"item-name"`
	// Scene item ID
	ItemId int `json:"item-id"`
}

// A scene item's visibility has been toggled.
type SceneItemVisibilityChangedEvent struct {
	eventData
	// Name of the scene.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	ItemName string `json:"item-name"`
	// Scene item ID
	ItemId int `json:"item-id"`
	// New visibility state of the item.
	ItemVisible bool `json:"item-visible"`
}

// A scene item's locked status has been toggled.
type SceneItemLockChangedEvent struct {
	eventData
	// Name of the scene.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	ItemName string `json:"item-name"`
	// Scene item ID
	ItemId int `json:"item-id"`
	// New locked state of the item.
	ItemLocked bool `json:"item-locked"`
}

// A scene item's transform has been changed.
type SceneItemTransformChangedEvent struct {
	eventData
	// Name of the scene.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	ItemName string `json:"item-name"`
	// Scene item ID
	ItemId int `json:"item-id"`
	// Scene item transform properties
	Transform SceneItemTransform `json:"transform"`
}

// A scene item is selected.
type SceneItemSelectedEvent struct {
	eventData
	// Name of the scene.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	ItemName string `json:"item-name"`
	// Name of the item in the scene.
	ItemId int `json:"item-id"`
}

// A scene item is deselected.
type SceneItemDeselectedEvent struct {
	eventData
	// Name of the scene.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	ItemName string `json:"item-name"`
	// Name of the item in the scene.
	ItemId int `json:"item-id"`
}

// The active transition has been changed.
type SwitchTransitionEvent struct {
	eventData
	// The name of the new active transition.
	TransitionName string `json:"transition-name"`
}

// The list of available transitions has been modified. Transitions have been
// added, removed, or renamed.
type TransitionListChangedEvent struct {
	eventData
	// Transitions list.
	Transitions []struct {
		// Transition name.
		Name string `json:"name"`
	} `json:"transitions"`
}

// The active transition duration has been changed.
type TransitionDurationChangedEvent struct {
	eventData
	// New transition duration.
	NewDuration int `json:"new-duration"`
}

// A transition (other than "cut") has begun.
type TransitionBeginEvent struct {
	eventData
	// Transition name.
	Name string `json:"name"`
	// Transition type.
	Type string `json:"type"`
	// Transition duration (in milliseconds). Will be -1 for any transition with a
	// fixed duration, such as a Stinger, due to limitations of the OBS API.
	Duration int `json:"duration"`
	// Source scene of the transition
	FromScene string `json:"from-scene,omitempty"`
	// Destination scene of the transition
	ToScene string `json:"to-scene"`
}

// A transition (other than "cut") has ended. Note: The `from-scene` field is
// not available in TransitionEnd.
type TransitionEndEvent struct {
	eventData
	// Transition name.
	Name string `json:"name"`
	// Transition type.
	Type string `json:"type"`
	// Transition duration (in milliseconds).
	Duration int `json:"duration"`
	// Destination scene of the transition
	ToScene string `json:"to-scene"`
}

// A stinger transition has finished playing its video.
type TransitionVideoEndEvent struct {
	eventData
	// Transition name.
	Name string `json:"name"`
	// Transition type.
	Type string `json:"type"`
	// Transition duration (in milliseconds).
	Duration int `json:"duration"`
	// Source scene of the transition
	FromScene string `json:"from-scene,omitempty"`
	// Destination scene of the transition
	ToScene string `json:"to-scene"`
}

// A request to start streaming has been issued.
type StreamStartingEvent struct {
	eventData
	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`
}

// Streaming started successfully.
type StreamStartedEvent struct {
	eventData
}

// A request to stop streaming has been issued.
type StreamStoppingEvent struct {
	eventData
	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`
}

// Streaming stopped successfully.
type StreamStoppedEvent struct {
	eventData
}

// Emitted every 2 seconds when stream is active.
type StreamStatusEvent struct {
	eventData
	// Current streaming state.
	Streaming bool `json:"streaming"`
	// Current recording state.
	Recording bool `json:"recording"`
	// Replay Buffer status
	ReplayBufferActive bool `json:"replay-buffer-active"`
	// Amount of data per second (in bytes) transmitted by the stream encoder.
	BytesPerSec int `json:"bytes-per-sec"`
	// Amount of data per second (in kilobits) transmitted by the stream encoder.
	KbitsPerSec int `json:"kbits-per-sec"`
	// Percentage of dropped frames.
	Strain float64 `json:"strain"`
	// Total time (in seconds) since the stream started.
	TotalStreamTime int `json:"total-stream-time"`
	// Total number of frames transmitted since the stream started.
	NumTotalFrames int `json:"num-total-frames"`
	// Number of frames dropped by the encoder since the stream started.
	NumDroppedFrames int `json:"num-dropped-frames"`
	// Current framerate.
	Fps float64 `json:"fps"`
	// Number of frames rendered
	RenderTotalFrames int `json:"render-total-frames"`
	// Number of frames missed due to rendering lag
	RenderMissedFrames int `json:"render-missed-frames"`
	// Number of frames outputted
	OutputTotalFrames int `json:"output-total-frames"`
	// Number of frames skipped due to encoding lag
	OutputSkippedFrames int `json:"output-skipped-frames"`
	// Average frame time (in milliseconds)
	AverageFrameTime float64 `json:"average-frame-time"`
	// Current CPU usage (percentage)
	CpuUsage float64 `json:"cpu-usage"`
	// Current RAM usage (in megabytes)
	MemoryUsage float64 `json:"memory-usage"`
	// Free recording disk space (in megabytes)
	FreeDiskSpace float64 `json:"free-disk-space"`
	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`
}

//   Note: `recordingFilename` is not provided in this event because this
// information is not available at the time this event is emitted.
type RecordingStartingEvent struct {
	eventData
}

// Recording started successfully.
type RecordingStartedEvent struct {
	eventData
	// Absolute path to the file of the current recording.
	RecordingFilename string `json:"recordingFilename"`
}

// A request to stop recording has been issued.
type RecordingStoppingEvent struct {
	eventData
	// Absolute path to the file of the current recording.
	RecordingFilename string `json:"recordingFilename"`
}

// Recording stopped successfully.
type RecordingStoppedEvent struct {
	eventData
	// Absolute path to the file of the current recording.
	RecordingFilename string `json:"recordingFilename"`
}

// Current recording paused
type RecordingPausedEvent struct {
	eventData
}

// Current recording resumed
type RecordingResumedEvent struct {
	eventData
}

// A request to start the replay buffer has been issued.
type ReplayStartingEvent struct {
	eventData
}

// Replay Buffer started successfully
type ReplayStartedEvent struct {
	eventData
}

// A request to stop the replay buffer has been issued.
type ReplayStoppingEvent struct {
	eventData
}

// Replay Buffer stopped successfully
type ReplayStoppedEvent struct {
	eventData
}

// OBS is exiting.
type ExitingEvent struct {
	eventData
}

// A source has been created. A source can be an input, a scene or a transition.
type SourceCreatedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Source type. Can be "input", "scene", "transition" or "filter".
	SourceType string `json:"sourceType"`
	// Source kind.
	SourceKind string `json:"sourceKind"`
	// Source settings
	SourceSettings interface{} `json:"sourceSettings"`
}

// A source has been destroyed/removed. A source can be an input, a scene or a
// transition.
type SourceDestroyedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Source type. Can be "input", "scene", "transition" or "filter".
	SourceType string `json:"sourceType"`
	// Source kind.
	SourceKind string `json:"sourceKind"`
}

// The volume of a source has changed.
type SourceVolumeChangedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Source volume
	Volume float32 `json:"volume"`
	// Source volume in Decibel
	VolumeDb float32 `json:"volumeDb"`
}

// A source has been muted or unmuted.
type SourceMuteStateChangedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Mute status of the source
	Muted bool `json:"muted"`
}

// A source has removed audio.
type SourceAudioDeactivatedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
}

// A source has added audio.
type SourceAudioActivatedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
}

// The audio sync offset of a source has changed.
type SourceAudioSyncOffsetChangedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Audio sync offset of the source (in nanoseconds)
	SyncOffset int `json:"syncOffset"`
}

// Audio mixer routing changed on a source.
type SourceAudioMixersChangedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Routing status of the source for each audio mixer (array of 6 values)
	Mixers []struct {
		// Mixer number
		Id int `json:"id"`
		// Routing status
		Enabled bool `json:"enabled"`
	} `json:"mixers"`
	// Raw mixer flags (little-endian, one bit per mixer) as an hexadecimal value
	HexMixersValue string `json:"hexMixersValue"`
}

// A source has been renamed.
type SourceRenamedEvent struct {
	eventData
	// Previous source name
	PreviousName string `json:"previousName"`
	// New source name
	NewName string `json:"newName"`
	// Type of source (input, scene, filter, transition)
	SourceType string `json:"sourceType"`
}

// A filter was added to a source.
type SourceFilterAddedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Filter name
	FilterName string `json:"filterName"`
	// Filter type
	FilterType string `json:"filterType"`
	// Filter settings
	FilterSettings interface{} `json:"filterSettings"`
}

// A filter was removed from a source.
type SourceFilterRemovedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Filter name
	FilterName string `json:"filterName"`
	// Filter type
	FilterType string `json:"filterType"`
}

// The visibility/enabled state of a filter changed
type SourceFilterVisibilityChangedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Filter name
	FilterName string `json:"filterName"`
	// New filter state
	FilterEnabled bool `json:"filterEnabled"`
}

// Filters in a source have been reordered.
type SourceFiltersReorderedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// Ordered Filters list
	Filters []struct {
		// Filter name
		Name string `json:"name"`
		// Filter type
		Type string `json:"type"`
		// Filter visibility status
		Enabled bool `json:"enabled"`
	} `json:"filters"`
}

//   Note: This event is only emitted when something actively controls the
// media/VLC source. In other words, the source will never emit this on its own
// naturally.
type MediaPlayingEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind"`
}

//   Note: This event is only emitted when something actively controls the
// media/VLC source. In other words, the source will never emit this on its own
// naturally.
type MediaPausedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind"`
}

//   Note: This event is only emitted when something actively controls the
// media/VLC source. In other words, the source will never emit this on its own
// naturally.
type MediaRestartedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind"`
}

//   Note: This event is only emitted when something actively controls the
// media/VLC source. In other words, the source will never emit this on its own
// naturally.
type MediaStoppedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind"`
}

//   Note: This event is only emitted when something actively controls the
// media/VLC source. In other words, the source will never emit this on its own
// naturally.
type MediaNextEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind"`
}

//   Note: This event is only emitted when something actively controls the
// media/VLC source. In other words, the source will never emit this on its own
// naturally.
type MediaPreviousEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind"`
}

//   Note: These events are emitted by the OBS sources themselves. For example
// when the media file starts playing. The behavior depends on the type of media
// source being used.
type MediaStartedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind"`
}

//   Note: These events are emitted by the OBS sources themselves. For example
// when the media file ends. The behavior depends on the type of media source
// being used.
type MediaEndedEvent struct {
	eventData
	// Source name
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind"`
}

// The selected preview scene has changed (only available in Studio Mode).
type PreviewSceneChangedEvent struct {
	eventData
	// Name of the scene being previewed.
	SceneName string `json:"scene-name"`
	// List of sources composing the scene. Same specification as
	// [`GetCurrentScene`](#getcurrentscene).
	Sources []SceneItem `json:"sources"`
}

// Studio Mode has been enabled or disabled.
type StudioModeSwitchedEvent struct {
	eventData
	// The new enabled state of Studio Mode.
	NewState bool `json:"new-state"`
}
