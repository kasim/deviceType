package deviceType

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("input", `{"type":"button", "id":"1", "name":"b1", "action":"clicked", "value": -1.0}`)

	act.Eval(tc)

	deviceType := tc.GetOutput("type")
	deviceId := tc.GetOutput("id")
	name := tc.GetOutput("name")
	action := tc.GetOutput("action")
	value := tc.GetOutput("value")

	assert.Equal(t, "button", deviceType)
	assert.Equal(t, "1", deviceId)
	assert.Equal(t, "b1", name)
	assert.Equal(t, "clicked", action)
	assert.Equal(t, -1.0, value)

	//check result attr
}
