package main

import "github.com/orbs-network/pbparser"

func fileHasServiceMethods(file *pbparser.ProtoFile) bool {
	for _, service := range file.Services {
		if len(service.RPCs) > 0 {
			return true
		}
	}
	return false
}
