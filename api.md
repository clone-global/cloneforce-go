# V1

## Clones

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneListResponse">V1CloneListResponse</a>

Methods:

- <code title="get /public/v1/clones">client.V1.Clones.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneListResponse">V1CloneListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Profile

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#CloneHeadshot">CloneHeadshot</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#CloneProfile">CloneProfile</a>

Methods:

- <code title="get /public/v1/clones/{cloneId}/profile">client.V1.Clones.Profile.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#CloneProfile">CloneProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /public/v1/clones/{cloneId}/profile">client.V1.Clones.Profile.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneProfileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneProfileUpdateParams">V1CloneProfileUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#CloneProfile">CloneProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Voice

Params Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#GenerateRequestParam">GenerateRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#GenerationStatus">GenerationStatus</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/voice/generate">client.V1.Clones.Voice.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneVoiceService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneVoiceGenerateParams">V1CloneVoiceGenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#GenerationStatus">GenerationStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Headshot

Methods:

- <code title="post /public/v1/clones/{cloneId}/headshot/generate">client.V1.Clones.Headshot.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneHeadshotService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneHeadshotGenerateParams">V1CloneHeadshotGenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#GenerationStatus">GenerationStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Skills

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#SkillSummary">SkillSummary</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillListResponse">V1CloneSkillListResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillDeleteResponse">V1CloneSkillDeleteResponse</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/skills">client.V1.Clones.Skills.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillNewParams">V1CloneSkillNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#SkillSummary">SkillSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /public/v1/clones/{cloneId}/skills/{skillName}">client.V1.Clones.Skills.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, skillName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillUpdateParams">V1CloneSkillUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#SkillSummary">SkillSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/skills">client.V1.Clones.Skills.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillListParams">V1CloneSkillListParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillListResponse">V1CloneSkillListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /public/v1/clones/{cloneId}/skills/{skillName}">client.V1.Clones.Skills.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, skillName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillDeleteParams">V1CloneSkillDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillDeleteResponse">V1CloneSkillDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#SkillConnectionInfo">SkillConnectionInfo</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillConnectionListResponse">V1CloneSkillConnectionListResponse</a>

Methods:

- <code title="put /public/v1/clones/{cloneId}/skills/{skillName}/connections/{settingName}">client.V1.Clones.Skills.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillConnectionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, settingName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillConnectionUpdateParams">V1CloneSkillConnectionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#SkillConnectionInfo">SkillConnectionInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/skills/{skillName}/connections">client.V1.Clones.Skills.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, skillName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillConnectionListParams">V1CloneSkillConnectionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneSkillConnectionListResponse">V1CloneSkillConnectionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Tasks

Params Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#TaskRecurrenceParam">TaskRecurrenceParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#TaskRecurrence">TaskRecurrence</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#TaskSummary">TaskSummary</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskListResponse">V1CloneTaskListResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskDeleteResponse">V1CloneTaskDeleteResponse</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/tasks">client.V1.Clones.Tasks.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskNewParams">V1CloneTaskNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#TaskSummary">TaskSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/tasks/{taskId}">client.V1.Clones.Tasks.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, taskID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskGetParams">V1CloneTaskGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#TaskSummary">TaskSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /public/v1/clones/{cloneId}/tasks/{taskId}">client.V1.Clones.Tasks.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, taskID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskUpdateParams">V1CloneTaskUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#TaskSummary">TaskSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/tasks">client.V1.Clones.Tasks.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskListParams">V1CloneTaskListParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskListResponse">V1CloneTaskListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /public/v1/clones/{cloneId}/tasks/{taskId}">client.V1.Clones.Tasks.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, taskID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskDeleteParams">V1CloneTaskDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneTaskDeleteResponse">V1CloneTaskDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Files

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#KBFileSummary">KBFileSummary</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileListResponse">V1CloneFileListResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileDeleteResponse">V1CloneFileDeleteResponse</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/files">client.V1.Clones.Files.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileNewParams">V1CloneFileNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#KBFileSummary">KBFileSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/files/{fileId}">client.V1.Clones.Files.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileGetParams">V1CloneFileGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#KBFileSummary">KBFileSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/files">client.V1.Clones.Files.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileListResponse">V1CloneFileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /public/v1/clones/{cloneId}/files/{fileId}">client.V1.Clones.Files.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileDeleteParams">V1CloneFileDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneFileDeleteResponse">V1CloneFileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Gallery

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#GalleryItemSummary">GalleryItemSummary</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryListResponse">V1CloneGalleryListResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryDeleteResponse">V1CloneGalleryDeleteResponse</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/gallery">client.V1.Clones.Gallery.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryNewParams">V1CloneGalleryNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#GalleryItemSummary">GalleryItemSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/gallery/{itemId}">client.V1.Clones.Gallery.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryGetParams">V1CloneGalleryGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#GalleryItemSummary">GalleryItemSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/gallery">client.V1.Clones.Gallery.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryListParams">V1CloneGalleryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryListResponse">V1CloneGalleryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /public/v1/clones/{cloneId}/gallery/{itemId}">client.V1.Clones.Gallery.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryDeleteParams">V1CloneGalleryDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneGalleryDeleteResponse">V1CloneGalleryDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Integrations

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#IntegrationSummary">IntegrationSummary</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationListResponse">V1CloneIntegrationListResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationDeleteResponse">V1CloneIntegrationDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationNewPhoneResponse">V1CloneIntegrationNewPhoneResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationGetSetupURLResponse">V1CloneIntegrationGetSetupURLResponse</a>

Methods:

- <code title="get /public/v1/clones/{cloneId}/integrations/{integrationId}">client.V1.Clones.Integrations.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, integrationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationGetParams">V1CloneIntegrationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#IntegrationSummary">IntegrationSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/integrations">client.V1.Clones.Integrations.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationListParams">V1CloneIntegrationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationListResponse">V1CloneIntegrationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /public/v1/clones/{cloneId}/integrations/{integrationId}">client.V1.Clones.Integrations.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, integrationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationDeleteParams">V1CloneIntegrationDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationDeleteResponse">V1CloneIntegrationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /public/v1/clones/{cloneId}/integrations/phone">client.V1.Clones.Integrations.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationService.NewPhone">NewPhone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationNewPhoneParams">V1CloneIntegrationNewPhoneParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationNewPhoneResponse">V1CloneIntegrationNewPhoneResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/integrations/{integrationId}/setup">client.V1.Clones.Integrations.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationService.GetSetupURL">GetSetupURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, integrationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationGetSetupURLParams">V1CloneIntegrationGetSetupURLParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationGetSetupURLResponse">V1CloneIntegrationGetSetupURLResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Slack

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#SlackIntegration">SlackIntegration</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/integrations/slack">client.V1.Clones.Integrations.Slack.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationSlackService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#SlackIntegration">SlackIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /public/v1/clones/{cloneId}/integrations/slack/{integrationId}">client.V1.Clones.Integrations.Slack.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationSlackService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, integrationID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationSlackUpdateParams">V1CloneIntegrationSlackUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#SlackIntegration">SlackIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Msteams

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#MsTeamsTeamRef">MsTeamsTeamRef</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationMsteamTeamsResponse">V1CloneIntegrationMsteamTeamsResponse</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/integrations/msteams/{integrationId}/teams">client.V1.Clones.Integrations.Msteams.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationMsteamService.Teams">Teams</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, integrationID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationMsteamTeamsParams">V1CloneIntegrationMsteamTeamsParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneIntegrationMsteamTeamsResponse">V1CloneIntegrationMsteamTeamsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Activity

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityGetResponse">V1CloneActivityGetResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityListResponse">V1CloneActivityListResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityDeleteResponse">V1CloneActivityDeleteResponse</a>

Methods:

- <code title="get /public/v1/clones/{cloneId}/activity/{activityId}">client.V1.Clones.Activity.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, activityID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityGetParams">V1CloneActivityGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityGetResponse">V1CloneActivityGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/clones/{cloneId}/activity">client.V1.Clones.Activity.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityListResponse">V1CloneActivityListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /public/v1/clones/{cloneId}/activity/{activityId}">client.V1.Clones.Activity.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, activityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityDeleteParams">V1CloneActivityDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneActivityDeleteResponse">V1CloneActivityDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chats

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ChatCompletionResponse">ChatCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#CreateChatResponse">CreateChatResponse</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/chats">client.V1.Clones.Chats.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cloneID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneChatNewParams">V1CloneChatNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#CreateChatResponse">CreateChatResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ChatCompletionChunk">ChatCompletionChunk</a>

Methods:

- <code title="post /public/v1/clones/{cloneId}/chats/{chatId}/completions">client.V1.Clones.Chats.Completions.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1CloneChatCompletionNewParams">V1CloneChatCompletionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ChatCompletionResponse">ChatCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Skills

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1SkillGetResponse">V1SkillGetResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1SkillSearchResponse">V1SkillSearchResponse</a>

Methods:

- <code title="get /public/v1/skills/{skillId}">client.V1.Skills.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1SkillService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, skillID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1SkillGetParams">V1SkillGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1SkillGetResponse">V1SkillGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/skills/search">client.V1.Skills.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1SkillService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1SkillSearchParams">V1SkillSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1SkillSearchResponse">V1SkillSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Integrations

### Phone

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1IntegrationPhoneListAvailableResponse">V1IntegrationPhoneListAvailableResponse</a>

Methods:

- <code title="get /public/v1/integrations/phone/available">client.V1.Integrations.Phone.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1IntegrationPhoneService.ListAvailable">ListAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1IntegrationPhoneListAvailableParams">V1IntegrationPhoneListAvailableParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1IntegrationPhoneListAvailableResponse">V1IntegrationPhoneListAvailableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ConnectionDetail">ConnectionDetail</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ConnectionStatus">ConnectionStatus</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#OAuthProvision">OAuthProvision</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionListResponse">V1ConnectionListResponse</a>
- <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionDeleteResponse">V1ConnectionDeleteResponse</a>

Methods:

- <code title="post /public/v1/connections">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionNewParams">V1ConnectionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ConnectionDetail">ConnectionDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/connections/{connectionId}">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ConnectionDetail">ConnectionDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /public/v1/connections/{connectionId}">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionUpdateParams">V1ConnectionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ConnectionDetail">ConnectionDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/connections">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionListParams">V1ConnectionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionListResponse">V1ConnectionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /public/v1/connections/{connectionId}">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionDeleteResponse">V1ConnectionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /public/v1/connections/oauth">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.NewOAuth">NewOAuth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionNewOAuthParams">V1ConnectionNewOAuthParams</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#OAuthProvision">OAuthProvision</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /public/v1/connections/{connectionId}/status">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ConnectionStatus">ConnectionStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /public/v1/connections/{connectionId}/refresh">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.Refresh">Refresh</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#ConnectionStatus">ConnectionStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /public/v1/connections/{connectionId}/reprovision">client.V1.Connections.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#V1ConnectionService.ReprovisionOAuth">ReprovisionOAuth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go">cloneforce</a>.<a href="https://pkg.go.dev/github.com/clone-global/cloneforce-go#OAuthProvision">OAuthProvision</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
