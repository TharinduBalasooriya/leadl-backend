package test

import (
	fclLib "github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
	"testing"
)

func connectQL(t *testing.T){

	Tests := fclLib.NewFCLWrapper();
    Tests.RunELInterpretter("D:\\Tracified\\LogAnalyzer\\Leedl-backend\\src\\debug_env\\Defs.txt");
}




