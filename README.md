# cpe
## Install
    go get -u github.com/tongchengbin/cpe
## Example Usage
``` 
    main() {
	cpeString := "cpe:2.3:a:appserv_open_project:appserv:*:*:*:*:*:*:*:*"

	cpe, err := Parse(cpeString)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	fmt.Println("Part:", cpe.Part)
	fmt.Println("Vendor:", cpe.Vendor)
	fmt.Println("Product:", cpe.Product)
	fmt.Println("Version:", cpe.Version)
	fmt.Println("Update:", cpe.Update)
	fmt.Println("Edition:", cpe.Edition)
	fmt.Println("Language:", cpe.Language)
	fmt.Println("SwEdition:", cpe.SwEdition)
	fmt.Println("TargetSw:", cpe.TargetSw)
	fmt.Println("TargetHw:", cpe.TargetHw)
	fmt.Println("Other:", cpe.Other)
}
