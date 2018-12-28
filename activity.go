package deviceType 

import (
	//"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-deviceType")

const (
	ivType = "type"
	ivID = "id"
	ivName = "name"
	ivAction = "action"
	ivValue = "value"
)

// MyActivity is a stub for your Activity implementation
type Output struct {
	Name string `json:"name"`
	Action string `json:"action"`
	Value float64 `json:"value"`
}
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {
	input := context.GetInput("input")
	//output := Output{}
	//json.Unmarshal([]byte(input), &output)

	//output := make(map[string]interface{})
	//json.Unmarshal(input, &output)
	deviceType, ok := input.(map[string]interface{})[ivType].(string)
	if !ok { log.Error("No type!") }
	deviceId, ok := input.(map[string]interface{})[ivID].(string)
	if !ok { log.Error("No id!") }
	name, ok := input.(map[string]interface{})[ivName].(string)
	if !ok { log.Error("No name!") }
	action, ok := input.(map[string]interface{})[ivAction].(string)
	if !ok { log.Error("No action!") }
	value, ok := input.(map[string]interface{})[ivValue].(float64) 
	if !ok { log.Error("No value!") }

	log.Infof("input: %+v", input.(map[string]interface{}))
	//log.Infof("[%s]", input["name"].(string))
	//log.Infof("%+v", output)

	context.SetOutput(ivType, deviceType)
	context.SetOutput(ivID, deviceId)
	context.SetOutput(ivName, name)
	context.SetOutput(ivAction, action)
	context.SetOutput(ivValue, value)
	// do eval

	return true, nil
}
