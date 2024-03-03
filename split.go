package main

func fonction(s string, sep string) []string {
	//création d'un tableau
	tbl := []string(nil)
	startindex := 0
	Lsep := len(sep)

	//boucle jusqu'a len - len du sep pour eviter runtime
	for i := 0; i < len(s)-Lsep+1; i++ {

		//BANGER
		if s[i:i+Lsep] == sep {
			//CLASSICO
			tbl = append(tbl, s[startindex:i])
			startindex = i + Lsep
			i = i + Lsep - 1
		}
	}
	tbl = append(tbl, s[startindex:])
	return tbl
}

/*
	s = "HelloHAjeHAsuiHAclement"
	sep = "HA"
	return = {"Hello", "je", "suis", "clement"}
*/
