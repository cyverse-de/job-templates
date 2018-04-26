package jobs

import (
	"testing"
)

func TestGenerateOSGSubmit(t *testing.T) {
	cfg := InitConfig(t)
	s := InitTestsFromFile(t, cfg, "osg_submission.json")
	actual, err := generateFileContents(osgSubmissionTemplate, s)
	if err != nil {
		t.Error(err)
	}
	expected := `universe = vanilla
executable = wrapper
requirements = HAS_SINGULARITY == TRUE

output = script-output.log
error = script-error.log
log = condor.log

+SingularityImage = "/cvmfs/singularity.opensciencegrid.org/discoenv/osg-word-count"
+SingularityBindCVMFS = True

+IpcUuid = "2256dd6d-d984-4d3a-ad71-ab1ff341f636"
+IpcJobId = "generated_script"
+IpcUsername = "sarahr"

should_transfer_files = YES
transfer_executable = False
transfer_input_files = iplant.cmd,config.json,input_ticket.list,output_ticket.list
when_to_transfer_output = NEVER
notification = NEVER

queue
`
	if actual.String() != expected {
		t.Errorf("GenerateCondorSubmit() returned:\n\n%s\n\ninstead of:\n\n%s", actual, expected)
	}
}
