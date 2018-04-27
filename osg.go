package jobs

import (
	"github.com/spf13/viper"
	"gopkg.in/cyverse-de/model.v2"
)

// OSGJobSubmissionBuilder is responsible for writing out the iplant.cmd, config.json,
// input_ticket.list, and output_ticket.list files for jobs that are sent to OSG.
type OSGJobSubmissionBuilder struct {
	cfg *viper.Viper
}

// Build is where the the files are actually written out for submissions to OSG.
func (b OSGJobSubmissionBuilder) Build(submission *model.Job, dirPath string) (string, error) {
	var err error

	templateFields := OtherTemplateFields{TicketPathListHeader: b.cfg.GetString("tickets_path_list.file_identifier")}
	templateModel := TemplatesModel{
		submission,
		templateFields,
	}

	submission.OutputTicketFile, err = generateOutputTicketList(dirPath, templateModel)
	if err != nil {
		return "", err
	}

	submission.InputTicketsFile, err = generateInputTicketList(dirPath, templateModel)
	if err != nil {
		return "", err
	}

	// Generate the submission file.
	submitFilePath, err := generateFile(dirPath, "iplant.cmd", osgSubmissionTemplate, submission)
	if err != nil {
		return "", err
	}

	return submitFilePath, nil
}

func newOSGJobSubmissionBuilder(cfg *viper.Viper) JobSubmissionBuilder {
	return OSGJobSubmissionBuilder{cfg: cfg}
}
