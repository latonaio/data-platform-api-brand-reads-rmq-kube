package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"APIType"`
	Header            Header   `json:"Brand"`
	Headers           Headers  `json:"Brands"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	Brand                         int       `json:"Brand"`
	BrandType                     *string   `json:"BrandType"`
	BrandOwner                    *int      `json:"BrandOwner"`
	BrandOwnerBusinessPartnerRole *string   `json:"BrandOwnerBusinessPartnerRole"`
	PersonResponsible             *string   `json:"PersonResponsible"`
	ValidityStartDate             *string   `json:"ValidityStartDate"`
	ValidityStartTime             *string   `json:"ValidityStartTime"`
	ValidityEndDate               *string   `json:"ValidityEndDate"`
	ValidityEndTime               *string   `json:"ValidityEndTime"`
	Description                   *string   `json:"Description"`
	LongText                      *string   `json:"LongText"`
	Introduction                  *string   `json:"Introduction"`
	Project                       *int      `json:"Project"`
	WBSElement                    *int      `json:"WBSElement"`
	Tag1                          *string   `json:"Tag1"`
	Tag2                          *string   `json:"Tag2"`
	Tag3                          *string   `json:"Tag3"`
	Tag4                          *string   `json:"Tag4"`
	CreationDate                  *string   `json:"CreationDate"`
	CreationTime                  *string   `json:"CreationTime"`
	LastChangeDate                *string   `json:"LastChangeDate"`
	LastChangeTime                *string   `json:"LastChangeTime"`
	IsReleased					  *bool	   `json:"IsReleased"`
	IsMarkedForDeletion           *bool     `json:"IsMarkedForDeletion"`
	Partner                       []Partner `json:"Partner"`
}

type Headers []struct {
	Brand                		  int     `json:"Brand"`
	BrandType            		  *string `json:"BrandType"`
	BrandOwner           		  *int    `json:"BrandOwner"`
	BrandOwnerBusinessPartnerRole *string   `json:"BrandOwnerBusinessPartnerRole"`
	PersonResponsible   		  *int    `json:"PersonResponsible"`
	ValidityStartDate			  *string `json:"ValidityStartDate"`
	ValidityStartTime   		  *string `json:"ValidityStartTime"`
	ValidityEndDate     		  *string `json:"ValidityEndDate"`
	ValidityEndTime     		  *string `json:"ValidityEndTime"`
	Description         		  *string `json:"Description"`
	LongText            		  *string `json:"LongText"`
	Introduction				  *string `json:"Introduction"`
	Project             		  *int    `json:"Project"`
	WBSElement          		  *int    `json:"WBSElement"`
	Tag1                		  *string `json:"Tag1"`
	Tag2                		  *string `json:"Tag2"`
	Tag3                		  *string `json:"Tag3"`
	Tag4                		  *string `json:"Tag4"`
	CreationDate                  *string `json:"CreationDate"`
	CreationTime                  *string `json:"CreationTime"`
	LastChangeDate                *string `json:"LastChangeDate"`
	LastChangeTime                *string `json:"LastChangeTime"`
	IsReleased					  *bool	 `json:"IsReleased"`
	IsMarkedForDeletion 		  *bool   `json:"IsMarkedForDeletion"`
	Partner                       []Partner `json:"Partner"`
}

type Partner struct {
	Brand                    int     `json:"Brand"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}
