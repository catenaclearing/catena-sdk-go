# RefHosRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HosRulesetCode** | **string** | Stable identifier for the HOS Cycle Ruleset (normalized slug). | 
**HosRulesetName** | **string** | Human-readable ruleset rule label (display; not for joins). | 
**JurisdictionType** | **string** | Scope of authority for the rule (US: Interstate/Intrastate; CA: Federal/Provincial). | 
**IssueAuthority** | **string** | Regulator/authority issuing the rule (e.g., FMCSA, Transport Canada, TX DOT). | 
**Region** | **string** | Friendly region label (e.g., US Interstate, US Texas Intrastate, Canada Federal). | 
**RegionCode** | **string** | ISO code for geography (ISO-3166), e.g., US, US-TX, CA, CA-AB. | 
**CycleDays** | Pointer to **NullableInt32** |  | [optional] 
**MaxCycleHours** | Pointer to **NullableInt32** |  | [optional] 
**MaxDriveHoursPerDay** | Pointer to **NullableInt32** |  | [optional] 
**MaxWorkHoursPerDay** | Pointer to **NullableInt32** |  | [optional] 
**MinOffdutyPerDay** | Pointer to **NullableInt32** |  | [optional] 
**IsBreakRequired** | Pointer to **NullableBool** |  | [optional] 
**BreakIntervalHours** | Pointer to **NullableInt32** |  | [optional] 
**IsRestartAllowed** | Pointer to **NullableBool** |  | [optional] 
**RestartHours** | Pointer to **NullableInt32** |  | [optional] 
**RestartFrequencyDays** | Pointer to **NullableInt32** |  | [optional] 
**IsSleeperSplitAllowed** | Pointer to **NullableBool** |  | [optional] 
**SleeperSplitNotes** | Pointer to **NullableString** |  | [optional] 
**IsTeamRules** | Pointer to **NullableBool** |  | [optional] 
**IsBigDayAllowed** | Pointer to **NullableBool** |  | [optional] 
**IsShortHaulExemption** | Pointer to **NullableBool** |  | [optional] 
**ShortHaulAirmiles** | Pointer to **NullableInt32** |  | [optional] 
**ShortHaulMaxDutyHours** | Pointer to **NullableInt32** |  | [optional] 
**IsPassenger** | Pointer to **NullableBool** |  | [optional] 
**IsRailroadExemption** | Pointer to **NullableBool** |  | [optional] 
**IsOilfield** | Pointer to **NullableBool** |  | [optional] 
**IsFarmProduct** | Pointer to **NullableBool** |  | [optional] 
**IsFlammable** | Pointer to **NullableBool** |  | [optional] 
**IsSchoolPupil** | Pointer to **NullableBool** |  | [optional] 
**IsSeasonalExemption** | Pointer to **NullableBool** |  | [optional] 
**IsSpecialExemption** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRefHosRuleset

`func NewRefHosRuleset(hosRulesetCode string, hosRulesetName string, jurisdictionType string, issueAuthority string, region string, regionCode string, ) *RefHosRuleset`

NewRefHosRuleset instantiates a new RefHosRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefHosRulesetWithDefaults

`func NewRefHosRulesetWithDefaults() *RefHosRuleset`

NewRefHosRulesetWithDefaults instantiates a new RefHosRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosRulesetCode

`func (o *RefHosRuleset) GetHosRulesetCode() string`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *RefHosRuleset) GetHosRulesetCodeOk() (*string, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *RefHosRuleset) SetHosRulesetCode(v string)`

SetHosRulesetCode sets HosRulesetCode field to given value.


### GetHosRulesetName

`func (o *RefHosRuleset) GetHosRulesetName() string`

GetHosRulesetName returns the HosRulesetName field if non-nil, zero value otherwise.

### GetHosRulesetNameOk

`func (o *RefHosRuleset) GetHosRulesetNameOk() (*string, bool)`

GetHosRulesetNameOk returns a tuple with the HosRulesetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetName

`func (o *RefHosRuleset) SetHosRulesetName(v string)`

SetHosRulesetName sets HosRulesetName field to given value.


### GetJurisdictionType

`func (o *RefHosRuleset) GetJurisdictionType() string`

GetJurisdictionType returns the JurisdictionType field if non-nil, zero value otherwise.

### GetJurisdictionTypeOk

`func (o *RefHosRuleset) GetJurisdictionTypeOk() (*string, bool)`

GetJurisdictionTypeOk returns a tuple with the JurisdictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionType

`func (o *RefHosRuleset) SetJurisdictionType(v string)`

SetJurisdictionType sets JurisdictionType field to given value.


### GetIssueAuthority

`func (o *RefHosRuleset) GetIssueAuthority() string`

GetIssueAuthority returns the IssueAuthority field if non-nil, zero value otherwise.

### GetIssueAuthorityOk

`func (o *RefHosRuleset) GetIssueAuthorityOk() (*string, bool)`

GetIssueAuthorityOk returns a tuple with the IssueAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAuthority

`func (o *RefHosRuleset) SetIssueAuthority(v string)`

SetIssueAuthority sets IssueAuthority field to given value.


### GetRegion

`func (o *RefHosRuleset) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RefHosRuleset) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RefHosRuleset) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRegionCode

`func (o *RefHosRuleset) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *RefHosRuleset) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *RefHosRuleset) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetCycleDays

`func (o *RefHosRuleset) GetCycleDays() int32`

GetCycleDays returns the CycleDays field if non-nil, zero value otherwise.

### GetCycleDaysOk

`func (o *RefHosRuleset) GetCycleDaysOk() (*int32, bool)`

GetCycleDaysOk returns a tuple with the CycleDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleDays

`func (o *RefHosRuleset) SetCycleDays(v int32)`

SetCycleDays sets CycleDays field to given value.

### HasCycleDays

`func (o *RefHosRuleset) HasCycleDays() bool`

HasCycleDays returns a boolean if a field has been set.

### SetCycleDaysNil

`func (o *RefHosRuleset) SetCycleDaysNil(b bool)`

 SetCycleDaysNil sets the value for CycleDays to be an explicit nil

### UnsetCycleDays
`func (o *RefHosRuleset) UnsetCycleDays()`

UnsetCycleDays ensures that no value is present for CycleDays, not even an explicit nil
### GetMaxCycleHours

`func (o *RefHosRuleset) GetMaxCycleHours() int32`

GetMaxCycleHours returns the MaxCycleHours field if non-nil, zero value otherwise.

### GetMaxCycleHoursOk

`func (o *RefHosRuleset) GetMaxCycleHoursOk() (*int32, bool)`

GetMaxCycleHoursOk returns a tuple with the MaxCycleHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCycleHours

`func (o *RefHosRuleset) SetMaxCycleHours(v int32)`

SetMaxCycleHours sets MaxCycleHours field to given value.

### HasMaxCycleHours

`func (o *RefHosRuleset) HasMaxCycleHours() bool`

HasMaxCycleHours returns a boolean if a field has been set.

### SetMaxCycleHoursNil

`func (o *RefHosRuleset) SetMaxCycleHoursNil(b bool)`

 SetMaxCycleHoursNil sets the value for MaxCycleHours to be an explicit nil

### UnsetMaxCycleHours
`func (o *RefHosRuleset) UnsetMaxCycleHours()`

UnsetMaxCycleHours ensures that no value is present for MaxCycleHours, not even an explicit nil
### GetMaxDriveHoursPerDay

`func (o *RefHosRuleset) GetMaxDriveHoursPerDay() int32`

GetMaxDriveHoursPerDay returns the MaxDriveHoursPerDay field if non-nil, zero value otherwise.

### GetMaxDriveHoursPerDayOk

`func (o *RefHosRuleset) GetMaxDriveHoursPerDayOk() (*int32, bool)`

GetMaxDriveHoursPerDayOk returns a tuple with the MaxDriveHoursPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDriveHoursPerDay

`func (o *RefHosRuleset) SetMaxDriveHoursPerDay(v int32)`

SetMaxDriveHoursPerDay sets MaxDriveHoursPerDay field to given value.

### HasMaxDriveHoursPerDay

`func (o *RefHosRuleset) HasMaxDriveHoursPerDay() bool`

HasMaxDriveHoursPerDay returns a boolean if a field has been set.

### SetMaxDriveHoursPerDayNil

`func (o *RefHosRuleset) SetMaxDriveHoursPerDayNil(b bool)`

 SetMaxDriveHoursPerDayNil sets the value for MaxDriveHoursPerDay to be an explicit nil

### UnsetMaxDriveHoursPerDay
`func (o *RefHosRuleset) UnsetMaxDriveHoursPerDay()`

UnsetMaxDriveHoursPerDay ensures that no value is present for MaxDriveHoursPerDay, not even an explicit nil
### GetMaxWorkHoursPerDay

`func (o *RefHosRuleset) GetMaxWorkHoursPerDay() int32`

GetMaxWorkHoursPerDay returns the MaxWorkHoursPerDay field if non-nil, zero value otherwise.

### GetMaxWorkHoursPerDayOk

`func (o *RefHosRuleset) GetMaxWorkHoursPerDayOk() (*int32, bool)`

GetMaxWorkHoursPerDayOk returns a tuple with the MaxWorkHoursPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkHoursPerDay

`func (o *RefHosRuleset) SetMaxWorkHoursPerDay(v int32)`

SetMaxWorkHoursPerDay sets MaxWorkHoursPerDay field to given value.

### HasMaxWorkHoursPerDay

`func (o *RefHosRuleset) HasMaxWorkHoursPerDay() bool`

HasMaxWorkHoursPerDay returns a boolean if a field has been set.

### SetMaxWorkHoursPerDayNil

`func (o *RefHosRuleset) SetMaxWorkHoursPerDayNil(b bool)`

 SetMaxWorkHoursPerDayNil sets the value for MaxWorkHoursPerDay to be an explicit nil

### UnsetMaxWorkHoursPerDay
`func (o *RefHosRuleset) UnsetMaxWorkHoursPerDay()`

UnsetMaxWorkHoursPerDay ensures that no value is present for MaxWorkHoursPerDay, not even an explicit nil
### GetMinOffdutyPerDay

`func (o *RefHosRuleset) GetMinOffdutyPerDay() int32`

GetMinOffdutyPerDay returns the MinOffdutyPerDay field if non-nil, zero value otherwise.

### GetMinOffdutyPerDayOk

`func (o *RefHosRuleset) GetMinOffdutyPerDayOk() (*int32, bool)`

GetMinOffdutyPerDayOk returns a tuple with the MinOffdutyPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOffdutyPerDay

`func (o *RefHosRuleset) SetMinOffdutyPerDay(v int32)`

SetMinOffdutyPerDay sets MinOffdutyPerDay field to given value.

### HasMinOffdutyPerDay

`func (o *RefHosRuleset) HasMinOffdutyPerDay() bool`

HasMinOffdutyPerDay returns a boolean if a field has been set.

### SetMinOffdutyPerDayNil

`func (o *RefHosRuleset) SetMinOffdutyPerDayNil(b bool)`

 SetMinOffdutyPerDayNil sets the value for MinOffdutyPerDay to be an explicit nil

### UnsetMinOffdutyPerDay
`func (o *RefHosRuleset) UnsetMinOffdutyPerDay()`

UnsetMinOffdutyPerDay ensures that no value is present for MinOffdutyPerDay, not even an explicit nil
### GetIsBreakRequired

`func (o *RefHosRuleset) GetIsBreakRequired() bool`

GetIsBreakRequired returns the IsBreakRequired field if non-nil, zero value otherwise.

### GetIsBreakRequiredOk

`func (o *RefHosRuleset) GetIsBreakRequiredOk() (*bool, bool)`

GetIsBreakRequiredOk returns a tuple with the IsBreakRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBreakRequired

`func (o *RefHosRuleset) SetIsBreakRequired(v bool)`

SetIsBreakRequired sets IsBreakRequired field to given value.

### HasIsBreakRequired

`func (o *RefHosRuleset) HasIsBreakRequired() bool`

HasIsBreakRequired returns a boolean if a field has been set.

### SetIsBreakRequiredNil

`func (o *RefHosRuleset) SetIsBreakRequiredNil(b bool)`

 SetIsBreakRequiredNil sets the value for IsBreakRequired to be an explicit nil

### UnsetIsBreakRequired
`func (o *RefHosRuleset) UnsetIsBreakRequired()`

UnsetIsBreakRequired ensures that no value is present for IsBreakRequired, not even an explicit nil
### GetBreakIntervalHours

`func (o *RefHosRuleset) GetBreakIntervalHours() int32`

GetBreakIntervalHours returns the BreakIntervalHours field if non-nil, zero value otherwise.

### GetBreakIntervalHoursOk

`func (o *RefHosRuleset) GetBreakIntervalHoursOk() (*int32, bool)`

GetBreakIntervalHoursOk returns a tuple with the BreakIntervalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakIntervalHours

`func (o *RefHosRuleset) SetBreakIntervalHours(v int32)`

SetBreakIntervalHours sets BreakIntervalHours field to given value.

### HasBreakIntervalHours

`func (o *RefHosRuleset) HasBreakIntervalHours() bool`

HasBreakIntervalHours returns a boolean if a field has been set.

### SetBreakIntervalHoursNil

`func (o *RefHosRuleset) SetBreakIntervalHoursNil(b bool)`

 SetBreakIntervalHoursNil sets the value for BreakIntervalHours to be an explicit nil

### UnsetBreakIntervalHours
`func (o *RefHosRuleset) UnsetBreakIntervalHours()`

UnsetBreakIntervalHours ensures that no value is present for BreakIntervalHours, not even an explicit nil
### GetIsRestartAllowed

`func (o *RefHosRuleset) GetIsRestartAllowed() bool`

GetIsRestartAllowed returns the IsRestartAllowed field if non-nil, zero value otherwise.

### GetIsRestartAllowedOk

`func (o *RefHosRuleset) GetIsRestartAllowedOk() (*bool, bool)`

GetIsRestartAllowedOk returns a tuple with the IsRestartAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestartAllowed

`func (o *RefHosRuleset) SetIsRestartAllowed(v bool)`

SetIsRestartAllowed sets IsRestartAllowed field to given value.

### HasIsRestartAllowed

`func (o *RefHosRuleset) HasIsRestartAllowed() bool`

HasIsRestartAllowed returns a boolean if a field has been set.

### SetIsRestartAllowedNil

`func (o *RefHosRuleset) SetIsRestartAllowedNil(b bool)`

 SetIsRestartAllowedNil sets the value for IsRestartAllowed to be an explicit nil

### UnsetIsRestartAllowed
`func (o *RefHosRuleset) UnsetIsRestartAllowed()`

UnsetIsRestartAllowed ensures that no value is present for IsRestartAllowed, not even an explicit nil
### GetRestartHours

`func (o *RefHosRuleset) GetRestartHours() int32`

GetRestartHours returns the RestartHours field if non-nil, zero value otherwise.

### GetRestartHoursOk

`func (o *RefHosRuleset) GetRestartHoursOk() (*int32, bool)`

GetRestartHoursOk returns a tuple with the RestartHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartHours

`func (o *RefHosRuleset) SetRestartHours(v int32)`

SetRestartHours sets RestartHours field to given value.

### HasRestartHours

`func (o *RefHosRuleset) HasRestartHours() bool`

HasRestartHours returns a boolean if a field has been set.

### SetRestartHoursNil

`func (o *RefHosRuleset) SetRestartHoursNil(b bool)`

 SetRestartHoursNil sets the value for RestartHours to be an explicit nil

### UnsetRestartHours
`func (o *RefHosRuleset) UnsetRestartHours()`

UnsetRestartHours ensures that no value is present for RestartHours, not even an explicit nil
### GetRestartFrequencyDays

`func (o *RefHosRuleset) GetRestartFrequencyDays() int32`

GetRestartFrequencyDays returns the RestartFrequencyDays field if non-nil, zero value otherwise.

### GetRestartFrequencyDaysOk

`func (o *RefHosRuleset) GetRestartFrequencyDaysOk() (*int32, bool)`

GetRestartFrequencyDaysOk returns a tuple with the RestartFrequencyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartFrequencyDays

`func (o *RefHosRuleset) SetRestartFrequencyDays(v int32)`

SetRestartFrequencyDays sets RestartFrequencyDays field to given value.

### HasRestartFrequencyDays

`func (o *RefHosRuleset) HasRestartFrequencyDays() bool`

HasRestartFrequencyDays returns a boolean if a field has been set.

### SetRestartFrequencyDaysNil

`func (o *RefHosRuleset) SetRestartFrequencyDaysNil(b bool)`

 SetRestartFrequencyDaysNil sets the value for RestartFrequencyDays to be an explicit nil

### UnsetRestartFrequencyDays
`func (o *RefHosRuleset) UnsetRestartFrequencyDays()`

UnsetRestartFrequencyDays ensures that no value is present for RestartFrequencyDays, not even an explicit nil
### GetIsSleeperSplitAllowed

`func (o *RefHosRuleset) GetIsSleeperSplitAllowed() bool`

GetIsSleeperSplitAllowed returns the IsSleeperSplitAllowed field if non-nil, zero value otherwise.

### GetIsSleeperSplitAllowedOk

`func (o *RefHosRuleset) GetIsSleeperSplitAllowedOk() (*bool, bool)`

GetIsSleeperSplitAllowedOk returns a tuple with the IsSleeperSplitAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSleeperSplitAllowed

`func (o *RefHosRuleset) SetIsSleeperSplitAllowed(v bool)`

SetIsSleeperSplitAllowed sets IsSleeperSplitAllowed field to given value.

### HasIsSleeperSplitAllowed

`func (o *RefHosRuleset) HasIsSleeperSplitAllowed() bool`

HasIsSleeperSplitAllowed returns a boolean if a field has been set.

### SetIsSleeperSplitAllowedNil

`func (o *RefHosRuleset) SetIsSleeperSplitAllowedNil(b bool)`

 SetIsSleeperSplitAllowedNil sets the value for IsSleeperSplitAllowed to be an explicit nil

### UnsetIsSleeperSplitAllowed
`func (o *RefHosRuleset) UnsetIsSleeperSplitAllowed()`

UnsetIsSleeperSplitAllowed ensures that no value is present for IsSleeperSplitAllowed, not even an explicit nil
### GetSleeperSplitNotes

`func (o *RefHosRuleset) GetSleeperSplitNotes() string`

GetSleeperSplitNotes returns the SleeperSplitNotes field if non-nil, zero value otherwise.

### GetSleeperSplitNotesOk

`func (o *RefHosRuleset) GetSleeperSplitNotesOk() (*string, bool)`

GetSleeperSplitNotesOk returns a tuple with the SleeperSplitNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleeperSplitNotes

`func (o *RefHosRuleset) SetSleeperSplitNotes(v string)`

SetSleeperSplitNotes sets SleeperSplitNotes field to given value.

### HasSleeperSplitNotes

`func (o *RefHosRuleset) HasSleeperSplitNotes() bool`

HasSleeperSplitNotes returns a boolean if a field has been set.

### SetSleeperSplitNotesNil

`func (o *RefHosRuleset) SetSleeperSplitNotesNil(b bool)`

 SetSleeperSplitNotesNil sets the value for SleeperSplitNotes to be an explicit nil

### UnsetSleeperSplitNotes
`func (o *RefHosRuleset) UnsetSleeperSplitNotes()`

UnsetSleeperSplitNotes ensures that no value is present for SleeperSplitNotes, not even an explicit nil
### GetIsTeamRules

`func (o *RefHosRuleset) GetIsTeamRules() bool`

GetIsTeamRules returns the IsTeamRules field if non-nil, zero value otherwise.

### GetIsTeamRulesOk

`func (o *RefHosRuleset) GetIsTeamRulesOk() (*bool, bool)`

GetIsTeamRulesOk returns a tuple with the IsTeamRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTeamRules

`func (o *RefHosRuleset) SetIsTeamRules(v bool)`

SetIsTeamRules sets IsTeamRules field to given value.

### HasIsTeamRules

`func (o *RefHosRuleset) HasIsTeamRules() bool`

HasIsTeamRules returns a boolean if a field has been set.

### SetIsTeamRulesNil

`func (o *RefHosRuleset) SetIsTeamRulesNil(b bool)`

 SetIsTeamRulesNil sets the value for IsTeamRules to be an explicit nil

### UnsetIsTeamRules
`func (o *RefHosRuleset) UnsetIsTeamRules()`

UnsetIsTeamRules ensures that no value is present for IsTeamRules, not even an explicit nil
### GetIsBigDayAllowed

`func (o *RefHosRuleset) GetIsBigDayAllowed() bool`

GetIsBigDayAllowed returns the IsBigDayAllowed field if non-nil, zero value otherwise.

### GetIsBigDayAllowedOk

`func (o *RefHosRuleset) GetIsBigDayAllowedOk() (*bool, bool)`

GetIsBigDayAllowedOk returns a tuple with the IsBigDayAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBigDayAllowed

`func (o *RefHosRuleset) SetIsBigDayAllowed(v bool)`

SetIsBigDayAllowed sets IsBigDayAllowed field to given value.

### HasIsBigDayAllowed

`func (o *RefHosRuleset) HasIsBigDayAllowed() bool`

HasIsBigDayAllowed returns a boolean if a field has been set.

### SetIsBigDayAllowedNil

`func (o *RefHosRuleset) SetIsBigDayAllowedNil(b bool)`

 SetIsBigDayAllowedNil sets the value for IsBigDayAllowed to be an explicit nil

### UnsetIsBigDayAllowed
`func (o *RefHosRuleset) UnsetIsBigDayAllowed()`

UnsetIsBigDayAllowed ensures that no value is present for IsBigDayAllowed, not even an explicit nil
### GetIsShortHaulExemption

`func (o *RefHosRuleset) GetIsShortHaulExemption() bool`

GetIsShortHaulExemption returns the IsShortHaulExemption field if non-nil, zero value otherwise.

### GetIsShortHaulExemptionOk

`func (o *RefHosRuleset) GetIsShortHaulExemptionOk() (*bool, bool)`

GetIsShortHaulExemptionOk returns a tuple with the IsShortHaulExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShortHaulExemption

`func (o *RefHosRuleset) SetIsShortHaulExemption(v bool)`

SetIsShortHaulExemption sets IsShortHaulExemption field to given value.

### HasIsShortHaulExemption

`func (o *RefHosRuleset) HasIsShortHaulExemption() bool`

HasIsShortHaulExemption returns a boolean if a field has been set.

### SetIsShortHaulExemptionNil

`func (o *RefHosRuleset) SetIsShortHaulExemptionNil(b bool)`

 SetIsShortHaulExemptionNil sets the value for IsShortHaulExemption to be an explicit nil

### UnsetIsShortHaulExemption
`func (o *RefHosRuleset) UnsetIsShortHaulExemption()`

UnsetIsShortHaulExemption ensures that no value is present for IsShortHaulExemption, not even an explicit nil
### GetShortHaulAirmiles

`func (o *RefHosRuleset) GetShortHaulAirmiles() int32`

GetShortHaulAirmiles returns the ShortHaulAirmiles field if non-nil, zero value otherwise.

### GetShortHaulAirmilesOk

`func (o *RefHosRuleset) GetShortHaulAirmilesOk() (*int32, bool)`

GetShortHaulAirmilesOk returns a tuple with the ShortHaulAirmiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortHaulAirmiles

`func (o *RefHosRuleset) SetShortHaulAirmiles(v int32)`

SetShortHaulAirmiles sets ShortHaulAirmiles field to given value.

### HasShortHaulAirmiles

`func (o *RefHosRuleset) HasShortHaulAirmiles() bool`

HasShortHaulAirmiles returns a boolean if a field has been set.

### SetShortHaulAirmilesNil

`func (o *RefHosRuleset) SetShortHaulAirmilesNil(b bool)`

 SetShortHaulAirmilesNil sets the value for ShortHaulAirmiles to be an explicit nil

### UnsetShortHaulAirmiles
`func (o *RefHosRuleset) UnsetShortHaulAirmiles()`

UnsetShortHaulAirmiles ensures that no value is present for ShortHaulAirmiles, not even an explicit nil
### GetShortHaulMaxDutyHours

`func (o *RefHosRuleset) GetShortHaulMaxDutyHours() int32`

GetShortHaulMaxDutyHours returns the ShortHaulMaxDutyHours field if non-nil, zero value otherwise.

### GetShortHaulMaxDutyHoursOk

`func (o *RefHosRuleset) GetShortHaulMaxDutyHoursOk() (*int32, bool)`

GetShortHaulMaxDutyHoursOk returns a tuple with the ShortHaulMaxDutyHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortHaulMaxDutyHours

`func (o *RefHosRuleset) SetShortHaulMaxDutyHours(v int32)`

SetShortHaulMaxDutyHours sets ShortHaulMaxDutyHours field to given value.

### HasShortHaulMaxDutyHours

`func (o *RefHosRuleset) HasShortHaulMaxDutyHours() bool`

HasShortHaulMaxDutyHours returns a boolean if a field has been set.

### SetShortHaulMaxDutyHoursNil

`func (o *RefHosRuleset) SetShortHaulMaxDutyHoursNil(b bool)`

 SetShortHaulMaxDutyHoursNil sets the value for ShortHaulMaxDutyHours to be an explicit nil

### UnsetShortHaulMaxDutyHours
`func (o *RefHosRuleset) UnsetShortHaulMaxDutyHours()`

UnsetShortHaulMaxDutyHours ensures that no value is present for ShortHaulMaxDutyHours, not even an explicit nil
### GetIsPassenger

`func (o *RefHosRuleset) GetIsPassenger() bool`

GetIsPassenger returns the IsPassenger field if non-nil, zero value otherwise.

### GetIsPassengerOk

`func (o *RefHosRuleset) GetIsPassengerOk() (*bool, bool)`

GetIsPassengerOk returns a tuple with the IsPassenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassenger

`func (o *RefHosRuleset) SetIsPassenger(v bool)`

SetIsPassenger sets IsPassenger field to given value.

### HasIsPassenger

`func (o *RefHosRuleset) HasIsPassenger() bool`

HasIsPassenger returns a boolean if a field has been set.

### SetIsPassengerNil

`func (o *RefHosRuleset) SetIsPassengerNil(b bool)`

 SetIsPassengerNil sets the value for IsPassenger to be an explicit nil

### UnsetIsPassenger
`func (o *RefHosRuleset) UnsetIsPassenger()`

UnsetIsPassenger ensures that no value is present for IsPassenger, not even an explicit nil
### GetIsRailroadExemption

`func (o *RefHosRuleset) GetIsRailroadExemption() bool`

GetIsRailroadExemption returns the IsRailroadExemption field if non-nil, zero value otherwise.

### GetIsRailroadExemptionOk

`func (o *RefHosRuleset) GetIsRailroadExemptionOk() (*bool, bool)`

GetIsRailroadExemptionOk returns a tuple with the IsRailroadExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRailroadExemption

`func (o *RefHosRuleset) SetIsRailroadExemption(v bool)`

SetIsRailroadExemption sets IsRailroadExemption field to given value.

### HasIsRailroadExemption

`func (o *RefHosRuleset) HasIsRailroadExemption() bool`

HasIsRailroadExemption returns a boolean if a field has been set.

### SetIsRailroadExemptionNil

`func (o *RefHosRuleset) SetIsRailroadExemptionNil(b bool)`

 SetIsRailroadExemptionNil sets the value for IsRailroadExemption to be an explicit nil

### UnsetIsRailroadExemption
`func (o *RefHosRuleset) UnsetIsRailroadExemption()`

UnsetIsRailroadExemption ensures that no value is present for IsRailroadExemption, not even an explicit nil
### GetIsOilfield

`func (o *RefHosRuleset) GetIsOilfield() bool`

GetIsOilfield returns the IsOilfield field if non-nil, zero value otherwise.

### GetIsOilfieldOk

`func (o *RefHosRuleset) GetIsOilfieldOk() (*bool, bool)`

GetIsOilfieldOk returns a tuple with the IsOilfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOilfield

`func (o *RefHosRuleset) SetIsOilfield(v bool)`

SetIsOilfield sets IsOilfield field to given value.

### HasIsOilfield

`func (o *RefHosRuleset) HasIsOilfield() bool`

HasIsOilfield returns a boolean if a field has been set.

### SetIsOilfieldNil

`func (o *RefHosRuleset) SetIsOilfieldNil(b bool)`

 SetIsOilfieldNil sets the value for IsOilfield to be an explicit nil

### UnsetIsOilfield
`func (o *RefHosRuleset) UnsetIsOilfield()`

UnsetIsOilfield ensures that no value is present for IsOilfield, not even an explicit nil
### GetIsFarmProduct

`func (o *RefHosRuleset) GetIsFarmProduct() bool`

GetIsFarmProduct returns the IsFarmProduct field if non-nil, zero value otherwise.

### GetIsFarmProductOk

`func (o *RefHosRuleset) GetIsFarmProductOk() (*bool, bool)`

GetIsFarmProductOk returns a tuple with the IsFarmProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFarmProduct

`func (o *RefHosRuleset) SetIsFarmProduct(v bool)`

SetIsFarmProduct sets IsFarmProduct field to given value.

### HasIsFarmProduct

`func (o *RefHosRuleset) HasIsFarmProduct() bool`

HasIsFarmProduct returns a boolean if a field has been set.

### SetIsFarmProductNil

`func (o *RefHosRuleset) SetIsFarmProductNil(b bool)`

 SetIsFarmProductNil sets the value for IsFarmProduct to be an explicit nil

### UnsetIsFarmProduct
`func (o *RefHosRuleset) UnsetIsFarmProduct()`

UnsetIsFarmProduct ensures that no value is present for IsFarmProduct, not even an explicit nil
### GetIsFlammable

`func (o *RefHosRuleset) GetIsFlammable() bool`

GetIsFlammable returns the IsFlammable field if non-nil, zero value otherwise.

### GetIsFlammableOk

`func (o *RefHosRuleset) GetIsFlammableOk() (*bool, bool)`

GetIsFlammableOk returns a tuple with the IsFlammable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlammable

`func (o *RefHosRuleset) SetIsFlammable(v bool)`

SetIsFlammable sets IsFlammable field to given value.

### HasIsFlammable

`func (o *RefHosRuleset) HasIsFlammable() bool`

HasIsFlammable returns a boolean if a field has been set.

### SetIsFlammableNil

`func (o *RefHosRuleset) SetIsFlammableNil(b bool)`

 SetIsFlammableNil sets the value for IsFlammable to be an explicit nil

### UnsetIsFlammable
`func (o *RefHosRuleset) UnsetIsFlammable()`

UnsetIsFlammable ensures that no value is present for IsFlammable, not even an explicit nil
### GetIsSchoolPupil

`func (o *RefHosRuleset) GetIsSchoolPupil() bool`

GetIsSchoolPupil returns the IsSchoolPupil field if non-nil, zero value otherwise.

### GetIsSchoolPupilOk

`func (o *RefHosRuleset) GetIsSchoolPupilOk() (*bool, bool)`

GetIsSchoolPupilOk returns a tuple with the IsSchoolPupil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSchoolPupil

`func (o *RefHosRuleset) SetIsSchoolPupil(v bool)`

SetIsSchoolPupil sets IsSchoolPupil field to given value.

### HasIsSchoolPupil

`func (o *RefHosRuleset) HasIsSchoolPupil() bool`

HasIsSchoolPupil returns a boolean if a field has been set.

### SetIsSchoolPupilNil

`func (o *RefHosRuleset) SetIsSchoolPupilNil(b bool)`

 SetIsSchoolPupilNil sets the value for IsSchoolPupil to be an explicit nil

### UnsetIsSchoolPupil
`func (o *RefHosRuleset) UnsetIsSchoolPupil()`

UnsetIsSchoolPupil ensures that no value is present for IsSchoolPupil, not even an explicit nil
### GetIsSeasonalExemption

`func (o *RefHosRuleset) GetIsSeasonalExemption() bool`

GetIsSeasonalExemption returns the IsSeasonalExemption field if non-nil, zero value otherwise.

### GetIsSeasonalExemptionOk

`func (o *RefHosRuleset) GetIsSeasonalExemptionOk() (*bool, bool)`

GetIsSeasonalExemptionOk returns a tuple with the IsSeasonalExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeasonalExemption

`func (o *RefHosRuleset) SetIsSeasonalExemption(v bool)`

SetIsSeasonalExemption sets IsSeasonalExemption field to given value.

### HasIsSeasonalExemption

`func (o *RefHosRuleset) HasIsSeasonalExemption() bool`

HasIsSeasonalExemption returns a boolean if a field has been set.

### SetIsSeasonalExemptionNil

`func (o *RefHosRuleset) SetIsSeasonalExemptionNil(b bool)`

 SetIsSeasonalExemptionNil sets the value for IsSeasonalExemption to be an explicit nil

### UnsetIsSeasonalExemption
`func (o *RefHosRuleset) UnsetIsSeasonalExemption()`

UnsetIsSeasonalExemption ensures that no value is present for IsSeasonalExemption, not even an explicit nil
### GetIsSpecialExemption

`func (o *RefHosRuleset) GetIsSpecialExemption() string`

GetIsSpecialExemption returns the IsSpecialExemption field if non-nil, zero value otherwise.

### GetIsSpecialExemptionOk

`func (o *RefHosRuleset) GetIsSpecialExemptionOk() (*string, bool)`

GetIsSpecialExemptionOk returns a tuple with the IsSpecialExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialExemption

`func (o *RefHosRuleset) SetIsSpecialExemption(v string)`

SetIsSpecialExemption sets IsSpecialExemption field to given value.

### HasIsSpecialExemption

`func (o *RefHosRuleset) HasIsSpecialExemption() bool`

HasIsSpecialExemption returns a boolean if a field has been set.

### SetIsSpecialExemptionNil

`func (o *RefHosRuleset) SetIsSpecialExemptionNil(b bool)`

 SetIsSpecialExemptionNil sets the value for IsSpecialExemption to be an explicit nil

### UnsetIsSpecialExemption
`func (o *RefHosRuleset) UnsetIsSpecialExemption()`

UnsetIsSpecialExemption ensures that no value is present for IsSpecialExemption, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


