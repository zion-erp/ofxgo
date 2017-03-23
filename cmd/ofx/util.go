package main

import (
	"github.com/aclindsa/ofxgo"
)

func NewRequest() (*ofxgo.Client, *ofxgo.Request) {
	var client = ofxgo.Client{
		AppId:       appId,
		AppVer:      appVer,
		SpecVersion: ofxVersion,
		NoIndent:    noIndentRequests,
	}

	var query ofxgo.Request
	query.URL = serverURL
	query.Signon.ClientUID = ofxgo.UID(clientUID)
	query.Signon.UserId = ofxgo.String(username)
	query.Signon.UserPass = ofxgo.String(password)
	query.Signon.Org = ofxgo.String(org)
	query.Signon.Fid = ofxgo.String(fid)

	return &client, &query
}
