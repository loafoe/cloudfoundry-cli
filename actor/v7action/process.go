package v7action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/resources"
)

// Process represents a V3 actor process.
type Process resources.Process

// GetProcessByTypeAndApplication returns a process for the given application
// and type.
func (actor Actor) GetProcessByTypeAndApplication(processType string, appGUID string) (Process, Warnings, error) {
	process, warnings, err := actor.CloudControllerClient.GetApplicationProcessByType(appGUID, processType)
	if _, ok := err.(ccerror.ProcessNotFoundError); ok {
		return Process{}, Warnings(warnings), actionerror.ProcessNotFoundError{ProcessType: processType}
	}
	return Process(process), Warnings(warnings), err
}

func (actor Actor) ScaleProcessByApplication(appGUID string, process Process) (Warnings, error) {
	_, warnings, err := actor.CloudControllerClient.CreateApplicationProcessScale(appGUID, resources.Process(process))
	allWarnings := Warnings(warnings)
	if err != nil {
		if _, ok := err.(ccerror.ProcessNotFoundError); ok {
			return allWarnings, actionerror.ProcessNotFoundError{ProcessType: process.Type}
		}
		return allWarnings, err
	}

	return allWarnings, nil
}

func (actor Actor) UpdateProcessByTypeAndApplication(processType string, appGUID string, updatedProcess Process) (Warnings, error) {
	if updatedProcess.HealthCheckType != constant.HTTP {
		if updatedProcess.HealthCheckEndpoint != constant.ProcessHealthCheckEndpointDefault && updatedProcess.HealthCheckEndpoint != "" {
			return nil, actionerror.HTTPHealthCheckInvalidError{}
		}

		updatedProcess.HealthCheckEndpoint = ""
	}

	process, warnings, err := actor.GetProcessByTypeAndApplication(processType, appGUID)
	allWarnings := warnings
	if err != nil {
		return allWarnings, err
	}

	updatedProcess.GUID = process.GUID
	_, updateWarnings, err := actor.CloudControllerClient.UpdateProcess(resources.Process(updatedProcess))
	allWarnings = append(allWarnings, Warnings(updateWarnings)...)
	if err != nil {
		return allWarnings, err
	}

	return allWarnings, nil
}
