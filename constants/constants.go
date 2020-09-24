package constants

var (
	DESKTOP        = "/Desktop" // Desktop;
	AuditDirectory = "/audit"   // File which will be created;

	SavedFileDIRECTORY   = "/new-audits/"   // Path to downloaded audits;
	ParsedDataDirectory  = "/policy-info/"  // Path to our json;
	CustomAuditDirectory = "/custom-audit/" // Path to custom audits;

	Policy       = "policy"        // part of name in our file;
	ParsedPolicy = "parsed-policy" //part of name for parsed audit;
	CustomAudit  = "custom-audit"  // part of file name for custom created audit;

	AuditFormat      = ".audit"
	ParsedFileFormat = ".json"

	LinkToDownloadFrom = "https://www.tenable.com/downloads/api/v1/public/pages/cis-compliance-audit-policies/downloads/10758/download?i_agree_to_tenable_license_agreement=true"
	CustomItemStart    = "<custom_item>"
	CustomItemEnd      = "</custom_item>"

	ShellToUse = "bash"
)
